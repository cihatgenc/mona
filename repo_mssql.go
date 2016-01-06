package main

import (
	"fmt"
	"golang.org/x/sys/windows/registry"
	"golang.org/x/sys/windows/svc/mgr"
	"log"
	"os"
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

// GetHostName - Get hostname
func GetHostName() (string, error) {
	fmt.Printf("Executing GetHostName\n")
	name, err := os.Hostname()

	if err != nil {
		log.Fatal(err)
		return "", err
	}

	return name, nil
}

// GetSQLPortNumber - Get SQL port number
func GetSQLPortNumber(key registry.Key, keyname string) (string, error) {
	fmt.Printf("Executing GetSqlPortNumber\n")

	return "", nil
}

// GetKeyStats - Get registry key stats
func GetKeyStats(key registry.Key) (*registry.KeyInfo, error) {
	fmt.Printf("Executing GetKeyStats\n")
	stat, err := key.Stat()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return stat, nil
}

// GetKeyNames - Get the key names
func GetKeyNames(key registry.Key, count int) ([]string, error) {
	fmt.Printf("Executing GetKeyNames\n")
	n, err := key.ReadValueNames(count)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return n, nil
}

// GetKeyValues - Get the key names
func GetKeyValues(key registry.Key, keyname string) (string, error) {
	fmt.Printf("Executing GetKeyValues\n")
	n, _, err := key.GetStringValue(keyname)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	return n, nil
}

// GetServiceStatus - Return status of a service
func GetServiceStatus(name string) (string, error) {
	fmt.Printf("Executing GetServiceStatus for: %s\n", name)

    var status string
	m, err := mgr.Connect()
	if err != nil {
		return "", err
	}
	defer m.Disconnect()
	s, err := m.OpenService(name)
	if err != nil {
		return "", err
	}
	defer s.Close()

	state, err := s.Query()
	if err != nil {
		return "", err
	}

    switch state.State {
        default: status="Huh?"
        case 0 : status="unknown"
        case 1 : status="stopped"
        case 2 : status="start_pending"
        case 3 : status="stop_pending"
        case 4 : status="running"
        case 5 : status="continue_pending"
        case 6 : status="pause_pending"
        case 7 : status="paused"
        case 8 : status="service_not_found"
        case 9 : status="server_not_found"
    }
    
	fmt.Printf("State returned is: %v\n", status)
	return status, nil
	// func (s *Service) Query() (svc.Status, error)
}

/*
0 = MOM_SERVICE_UNKNOWN_STATE
1 = MOM_SERVICE_STOPPED
2 = MOM_SERVICE_START_PENDING
3 = MOM_SERVICE_STOP_PENDING
4 = MOM_SERVICE_RUNNING
5 = MOM_SERVICE_CONTINUE_PENDING
6 = MOM_SERVICE_PAUSE_PENDING
7 = MOM_SERVICE_PAUSED
8 = MOM_SERVICE_NOT_FOUND
9 = MOM_SERVER_NOT_FOUND
*/
