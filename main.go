package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Invalid argument count\nUsage: ./main <domain>")
		return
	}

	host_address := os.Args[1]
	dns_records := make(map[string][]string)
	ipv4 := []string{}

	ips, err := net.LookupIP(host_address)
	if err != nil {
		fmt.Println("Unexpected error.", err)
		return
	}

	for _, ip := range ips {
		if ip.To4() != nil {
			ipv4 = append(ipv4, ip.String())
		}
	}

	if len(ipv4) == 0 {
		fmt.Println("No IPv4 addresses.")
		return
	}

	fmt.Println("Domain: ", host_address)

	dns_records, err = LoadRecord("data.json")

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	old_ips, exists := dns_records[host_address]

	if exists {
		added, removed, unchanged := CompareIP(old_ips, ipv4)

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
		for _, v := range ipv4 {
			fmt.Println("[ADDED]", v)
		}
	}

	dns_records[host_address] = ipv4

	err = SaveRecord("data.json", dns_records)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}
