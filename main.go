package main

import (
	"fmt"
	"log"
)

func main() {
	store, err := NewPostgressStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", store)
	server := newApiServer(":8080", store)
	server.Run()
}
