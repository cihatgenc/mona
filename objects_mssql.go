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
