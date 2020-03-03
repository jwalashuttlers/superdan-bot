package main

import "log"

func init() {
	initConfig()
}

func main() {
	srv, err := NewServer()
	if err != nil {
		log.Fatalf("couldn't initialize the server: %v", err)
	}

	srv.Run()
}
