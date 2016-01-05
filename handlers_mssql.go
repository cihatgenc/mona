package main

import (
	"encoding/json"
	//"fmt"
	//"github.com/gorilla/mux"
	"net/http"
)

// Return version
func mssqlIndex(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintln(w, "Welcome!")
	myobject := Versioning{
		Name:    "Mona Microsoft SQL Server",
		Version: versionNumber,
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(myobject); err != nil {
		panic(err)
	}
}

// Return all MSSQL instances
func mssqlAllInstances(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintln(w, "Welcome!")
	myobject := ListAllInstances()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(myobject); err != nil {
		panic(err)
	}
}
