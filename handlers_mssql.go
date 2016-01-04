package main

import (
	"encoding/json"
	//"fmt"
	//"github.com/gorilla/mux"
	"net/http"
)

func mssqlIndex(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintln(w, "Welcome!")
	myversion := Versioning{
		Name:    "Mona",
		Version: versionNumber,
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(myversion); err != nil {
		panic(err)
	}
}
