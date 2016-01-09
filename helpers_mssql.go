package main

import (
	"fmt"
	"golang.org/x/sys/windows/registry"
	"golang.org/x/sys/windows/svc/mgr"
	"log"
	"os"
)

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

// GetKeyStrValues - Get the key values in string
func GetKeyStrValues(key registry.Key, keyname string) (string, error) {
	fmt.Printf("Executing GetKeyValues\n")
	n, _, err := key.GetStringValue(keyname)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	return n, nil
}

// GetKeyIntValues - Get the key values in integer
func GetKeyIntValues(key registry.Key, keyname string) (int, error) {
	fmt.Printf("Executing GetKeyValues\n")
	n, _, err := key.GetIntegerValue(keyname)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}

	return int(n), nil
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
	default:
		status = "Huh?"
	case 0:
		status = "unknown"
	case 1:
		status = "stopped"
	case 2:
		status = "start_pending"
	case 3:
		status = "stop_pending"
	case 4:
		status = "running"
	case 5:
		status = "continue_pending"
	case 6:
		status = "pause_pending"
	case 7:
		status = "paused"
	case 8:
		status = "service_not_found"
	case 9:
		status = "server_not_found"
	}

	fmt.Printf("State returned is: %v\n", status)
	return status, nil
}

// GetServiceStartType - Return startup mode for service
func GetServiceStartType(name string) (string, error) {
	fmt.Printf("Executing GetServiceStartType for: %s\n", name)
	var starttype string

	var keypath = fmt.Sprintf("SYSTEM\\CurrentControlSet\\services\\%s", name)

	k, err := registry.OpenKey(registry.LOCAL_MACHINE, keypath, registry.QUERY_VALUE)
	if err != nil {
		log.Fatal(err)
	}
	defer k.Close()

	//var s string
	s, err := GetKeyIntValues(k, "Start")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v", s)

	switch s {
	default:
		starttype = "unknown"
	case 2:
		starttype = "automatic"
	case 3:
		starttype = "manual"
	case 4:
		starttype = "disabled"
	}

	return starttype, nil
}

// GetSQLServiceNames - Return all sql service names
func GetSQLServiceNames() ([]string, []string, error) {
	fmt.Printf("Executing GetSQLServiceNames\n")
	var servicenames []string
	var instancenames []string

	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Microsoft SQL Server\Instance Names\SQL`, registry.QUERY_VALUE)
	if err != nil {
		fmt.Printf("No installed SQL Service found\n")
		//log.Warn(err)
		return servicenames, instancenames, nil
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
		if name != "MSSQLSERVER" {
			servicenames = append(servicenames, fmt.Sprintf("%s%s", "MSSQL$", name))
		} else {
			servicenames = append(servicenames, name)
		}

		if name == "MSSQLSERVER" {
			instancenames = append(instancenames, host)
		} else {
			instancenames = append(instancenames, fmt.Sprintf("%s\\%s", host, name))
		}
	}

	return servicenames, instancenames, nil
}

// GetSQLAgentNames - Return all sql agent names
func GetSQLAgentNames() {
	fmt.Printf("Executing GetSQLAgentNames")

}

// GetSQLRSNames - Return all sql agent names
func GetSQLRSNames() {
	fmt.Printf("Executing GetSQLRSNames")

}

// GetSQLASNames - Return all sql agent names
func GetSQLASNames() {
	fmt.Printf("Executing GetSQLASNames")

}

// GetSQLISNames - Return all sql agent names
func GetSQLISNames() {
	fmt.Printf("Executing GetSQLISNames")

}
