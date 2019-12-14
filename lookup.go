package main

import (
	"fmt"
	"net/http"
	//"sync"
	//"log"
)

var lookupMap map[string]string

func lookup_init() {
	lookupMap = make(map[string]string)
	lookupMap["localhost"] = "127.0.0.1"
}

func lookup_debug_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%d entries in map\n", len(lookupMap))

	if len(lookupMap) > 0 {
		for key, value := range lookupMap {
			fmt.Fprintf(w, "%s => %s\n", key, value)
		}
	}
}

//LATER: locking, per https://blog.golang.org/go-maps-in-action

func lookup_get(hostname string) (string) {
	return lookupMap[hostname]
}

func lookup_set(hostname string, address string) {
	lookupMap[hostname] = address
}
