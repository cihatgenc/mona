package main

import (
	"encoding/json"
	"net/http"
    "fmt"
    "unsafe"
)

// Return all connection strings for active SQL Servers
func checkSQLServices(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
    
    myobject := statusSQLServices()
    if fmt.Sprintln(unsafe.Sizeof(myobject)) == "0" {
		w.WriteHeader(http.StatusNoContent)
	} else {
        w.WriteHeader(http.StatusOK)

        if err := json.NewEncoder(w).Encode(myobject); err != nil {
            panic(err)
        }
    }
}