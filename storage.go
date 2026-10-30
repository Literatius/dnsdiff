package main

import (
	"encoding/json"
	"os"
)

func LoadRecord(jsonname string) (map[string][]string, error) {
	read_data, err := os.ReadFile(jsonname)

	dns_records := make(map[string][]string)

	if os.IsNotExist(err) {
		return dns_records, nil
	}
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(read_data, &dns_records)
	if err != nil {
		return nil, err
	}

	return dns_records, nil
}

func SaveRecord(jsonname string, records map[string][]string) error {
	data, err := json.MarshalIndent(records, "", " ")

	if err != nil {
		return err
	}

	err = os.WriteFile(jsonname, data, 0644)

	if err != nil {
		return err
	}

	return nil
}
