package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
	"time"
)

type Result struct {
	Domain string
	IPs    []string
	Err    error
}

var wg sync.WaitGroup

func normalize_domain(domain string) string {
	domain = strings.TrimSpace(domain)
	domain = strings.ToLower(domain)
	domain = strings.TrimSuffix(domain, ".")

	return domain
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Invalid argument count\nUsage: ./main <domains>")
		return
	}

	host_addresses := os.Args[1:]
	seen := make(map[string]bool)
	domains := []string{}

	for _, host_address := range host_addresses {
		host_address = normalize_domain(host_address)

		if len(host_address) == 0 {
			continue
		}

		if seen[host_address] {
			continue
		}

		domains = append(domains, host_address)
		seen[host_address] = true
	}

	results := make(chan Result)

	dns_records, err := LoadRecord("data.json")

	if err != nil {
		fmt.Println("Unexpected error.", err)
	}

	for _, v := range domains {
		wg.Add(1)

		go func(host string) {
			defer wg.Done()

			ctx, cancel := context.WithTimeout(
				context.Background(),
				3*time.Second,
			)
			defer cancel()

			resolve := net.DefaultResolver
			ips, err := resolve.LookupIP(ctx, "ip4", host)

			ipv4 := []string{}

			for _, ip := range ips {
				ipv4 = append(ipv4, ip.String())
			}

			results <- Result{
				Domain: host,
				IPs:    ipv4,
				Err:    err,
			}
		}(v)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for result := range results {
		if result.Err != nil {
			fmt.Println("Unexpected error. ", result.Err)
			continue
		}

		if len(result.IPs) == 0 {
			fmt.Println("[WARNING] No IPv4 addresses for: ", result.Domain)
			continue
		}

		fmt.Println("Domain:", result.Domain)

		old_ips, exists := dns_records[result.Domain]

		if exists {
			added, removed, unchanged := CompareIP(old_ips, result.IPs)

			for _, v := range added {
				fmt.Println("[ADDED]", v)
			}
			for _, v := range unchanged {
				fmt.Println("[UNCHANGED]", v)
			}
			for _, v := range removed {
				fmt.Println("[REMOVED]", v)
			}
		} else {
			for _, v := range result.IPs {
				fmt.Println("[ADDED]", v)
			}
		}

		dns_records[result.Domain] = result.IPs
	}

	err = SaveRecord("data.json", dns_records)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}
