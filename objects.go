package main

import (
	"net/http"
	"time"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

type Todo struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

type Todos []Todo

// Semantic Versioning
type Versioning struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	// Major   int    `json:"major"`
	// Minor   int    `json:"minor"`
	// Patch   int    `json:"patch"`
}
