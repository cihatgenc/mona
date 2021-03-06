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

type ServiceStatus struct {
	ServiceName string `json:"servicename"`
	Status      string `json:"status"`
}

type ServicesStatus []ServiceStatus

type SensuMessage struct{
    Status  string  `json:"status"`
    Message string  `json:"message"`
}
