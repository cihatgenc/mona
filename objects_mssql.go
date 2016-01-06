package main

import (
// "net/http"
// "time"
)

type Kv struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Kvs []Kv

type SqlConnection struct {
	ServerName string `json:"servername"`
	Port       string `json:"port"`
}

type SqlConnections []SqlConnection
