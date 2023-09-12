package main

import (
	"fmt"
	"net/http"
	"log"

	"minicache/minicache"
)

var db = map[string]string{
	"Tom":  "630",
	"Jack": "589",
	"Sam":  "567",
}

func main() {
	minicache.NewGroup("scores", 2<<10, minicache.GetterFunc(
		func(key string) ([]byte, error) {
			log.Println("[SlowDB] search key", key)
			if v, ok := db[key]; ok {
				return []byte(v), nil
			}
			return nil, fmt.Errorf("%s not exist", key)
		}))

	addr := "localhost:9999"
	peers := minicache.NewHTTPPool(addr)
	log.Println("minicache is running at", addr)
	log.Fatal(http.ListenAndServe(addr, peers))
}
