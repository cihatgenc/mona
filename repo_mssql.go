package main

import (
	"fmt"
	"golang.org/x/sys/windows/registry"
	"log"
	//"syscall"
)

func ListAllInstances() Kvs {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Microsoft SQL Server\Instance Names\SQL`, registry.QUERY_VALUE)
	if err != nil {
		log.Fatal(err)
	}
	defer k.Close()

	// Get registry key stats
	stat, err := GetKeyStats(k)
	if err != nil {
		log.Fatal(err)
	}
	// Get the key names
	names, err := GetValueNames(k, int(stat.ValueCount))
	if err != nil {
		log.Fatal(err)
	}
	// Fetch all values for the keynames
	var instances Kvs

	for _, name := range names {
		fmt.Printf("Registry Key Name: %s\n", name)

		s, _, err := k.GetStringValue(name)
		if err != nil {
			log.Fatal(err)
		}

		var entry = Kv{name, s}
		instances = append(instances, entry)
	}

	return instances
}

func GetKeyStats(key registry.Key) (*registry.KeyInfo, error) {
	fmt.Printf("Executing GetKeyStats\n")
	stat, err := key.Stat()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return stat, nil
}

func GetValueNames(key registry.Key, count int) ([]string, error) {
	fmt.Printf("Executing GetValueNames\n")
	n, err := key.ReadValueNames(count)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return n, nil
}
