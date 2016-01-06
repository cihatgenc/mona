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
	// func (s *Service) Query() (svc.Status, error)
}
