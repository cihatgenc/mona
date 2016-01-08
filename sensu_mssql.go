package main

import (
	"fmt"
	"log"
    //"golang.org/x/sys/windows/registry"
	//"golang.org/x/sys/windows/svc/mgr"
)

var sensuOk = "0"
var sensuWarning = "1"
var sensuCritical = "2"
var sensuUnknown = "3"

// statusSQLServices - Check SQL Server Services
func statusSQLServices()  (SensuMessage) {
	fmt.Printf("Executing checkSQLServices\n")
    var message SensuMessage
    var allok = true

	servicenames, instancenames, err := GetSQLServiceNames()
	if err != nil {
		log.Fatal(err)
	}   

    for i, servicename := range servicenames {
  		
        starttype, err := GetServiceStartType(servicename)
		if err != nil {
			log.Fatal(err)
		}      

		state, err := GetServiceStatus(servicename)
		if err != nil {
			log.Fatal(err)
            allok = false
            message.Message += fmt.Sprintf("%s=%s,", instancenames[i], "unknown")
            continue
		}

		fmt.Printf("State for %s is %s\n", instancenames[i], state)

		if state != "running" && starttype == "automatic" {
			allok = false
		}
        
        message.Message += fmt.Sprintf("%s=%s,", instancenames[i], state)
	}

	if message.Message == "" {
        message.Message = "NO DATA"
		message.Status = sensuUnknown
        return message
	}
    
    message.Message = message.Message[0:(len(message.Message)-1)]
    if allok == false{
        message.Status = sensuCritical
    } else {
        message.Status = sensuOk
    }
    
	return message
}