package main

import (
	"fmt"
	"golang.org/x/sys/windows/registry"
	"log"
	"unsafe"
)

// ListAllInstances - Return all installed instances
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
		var entry = Kv{"instancename", name}
		instances = append(instances, entry)
	}

	return instances
}

// ListAllConnections - Return connection parameters for all instances
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

// ListAllActiveConnections - Return connection parameters for all active instances
func ListAllActiveConnections() SqlConnections {
	fmt.Printf("Executing ListAllActiveConnections\n")
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

		var servicename string

		if name != "MSSQLSERVER" {
			servicename = fmt.Sprintf("%s%s", "MSSQL$", name)
		} else {
			servicename = name
		}

		state, err := GetServiceStatus(servicename)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("State for %s is %s\n", servicename, state)

		if state != "running" {
			fmt.Printf("%s is not active: %s\n", servicename, state)
			continue
		}

		var instancename string

		if name == "MSSQLSERVER" {
			instancename = host
		} else {
			instancename = fmt.Sprintf("%s\\%s", host, name)
		}

		var entry = SqlConnection{instancename, port}
		connections = append(connections, entry)
	}

	if fmt.Sprintln(unsafe.Sizeof(connections)) == "0" {
		return nil
	}
	return connections
}

func MssqlServicesStatus() {

}
