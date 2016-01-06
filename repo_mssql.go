package main

import (
	"fmt"
	"golang.org/x/sys/windows/registry"
	"log"
	"os"
	//"syscall"
)

// Return all installed instances
func ListAllInstances() Kvs {
	fmt.Printf("Executing ListAllInstances\n")
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Microsoft SQL Server\Instance Names\SQL`, registry.QUERY_VALUE)
	if err != nil {
		log.Fatal(err)
	}
	defer k.Close()

	stat, err := GetKeyStats(k)
	if err != nil {
		log.Fatal(err)
	}

	names, err := GetKeyNames(k, int(stat.ValueCount))
	if err != nil {
		log.Fatal(err)
	}
	// Fetch all values for the keynames
	var instances Kvs

	for _, name := range names {
		//fmt.Printf("Registry Key Name: %s\n", name)

		// s, err := GetKeyValues(k, name)
		// if err != nil {
		// 	log.Fatal(err)
		// }

		// var entry = Kv{name, s}
		var entry = Kv{"instancename", name}
		instances = append(instances, entry)
	}

	return instances
}

func ListAllConnections() SqlConnections {
	fmt.Printf("Executing ListAllConnections\n")
	var connections SqlConnections

	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Microsoft SQL Server\Instance Names\SQL`, registry.QUERY_VALUE)
	if err != nil {
		log.Fatal(err)
	}
	defer k.Close()

	stat, err := GetKeyStats(k)
	if err != nil {
		log.Fatal(err)
	}

	names, err := GetKeyNames(k, int(stat.ValueCount))
	if err != nil {
		log.Fatal(err)
	}

	host, err := GetHostName()
	if err != nil {
		log.Fatal(err)
	}

	for _, name := range names {
		s, err := GetKeyValues(k, name)
		if err != nil {
			log.Fatal(err)
		}

		var keypath = fmt.Sprintf("SOFTWARE\\Microsoft\\Microsoft SQL Server\\%s\\MSSQLServer\\SuperSocketNetLib\\Tcp\\IPAll", s)
		fmt.Printf("%s", keypath)

		regport, err := registry.OpenKey(registry.LOCAL_MACHINE, keypath, registry.QUERY_VALUE)
		if err != nil {
			log.Fatal(err)
		}
		defer regport.Close()
		// Try to get the fixed port
		var port string
		port, err = GetKeyValues(regport, "TcpPort")
		if err != nil {
			log.Fatal(err)
		}
		// If fix port is not set, get the dynamic port
		if port == "" {
			port, err = GetKeyValues(regport, "TcpDynamicPorts")
			if err != nil {
				log.Fatal(err)
			}
		}

		if name == "MSSQLSERVER" {
			name = host
		} else {
			name = fmt.Sprintf("%s\\%s", host, name)
		}

		var entry = SqlConnection{name, port}
		connections = append(connections, entry)
	}

	return connections
}

// Get hostname
func GetHostName() (string, error) {
	fmt.Printf("Executing GetHostName\n")
	name, err := os.Hostname()

	if err != nil {
		log.Fatal(err)
		return "", err
	}

	return name, nil
}

// Get SQL port number
func GetSqlPortNumber(key registry.Key, keyname string) (string, error) {
	fmt.Printf("Executing GetSqlPortNumber\n")

	return "", nil
}

// Get registry key stats
func GetKeyStats(key registry.Key) (*registry.KeyInfo, error) {
	fmt.Printf("Executing GetKeyStats\n")
	stat, err := key.Stat()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return stat, nil
}

// Get the key names
func GetKeyNames(key registry.Key, count int) ([]string, error) {
	fmt.Printf("Executing GetKeyNames\n")
	n, err := key.ReadValueNames(count)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return n, nil
}

// Get the key names
func GetKeyValues(key registry.Key, keyname string) (string, error) {
	fmt.Printf("Executing GetKeyValues\n")
	n, _, err := key.GetStringValue(keyname)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	return n, nil
}
