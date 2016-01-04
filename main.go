/*
 MonA - Monitoring API
 REST API for monitoring systems and applications. Idea is to make it pluggable.
*/

package main

import (
	"fmt"
	"log"
	"net/http"
)

// compile passing -ldflags "-X main.versionNumber <build>"
var versionNumber = "0.0.1"

func main() {
	fmt.Printf("Mona Version: %s\n", versionNumber)

	router := NewRouter()

	// Port number must be fetched by KV store like consul
	log.Fatal(http.ListenAndServe(":8080", router))
}
