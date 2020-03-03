package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/jwalashuttle/superdan-bot/splitwise"
)

func init() {
	initConfig()
}

func main() {
	getAccessToken := flag.Bool("access-token", false, "If enabled, will get access token and prints to the screen")
	flag.Parse()

	// If getAccessToken is true, get token, token secret then print and exit
	if *getAccessToken {
		// Get new splitwise client
		spClient, _ := splitwise.NewClient(splitwise.Opts{
			ConsumerKey:    cfg.Splitwise.ConsumerKey,
			ConsumerSecret: cfg.Splitwise.ConsumerSecret,
		})

		// Get splitwise access token
		token, tokenSecret, err := spClient.GetAccessToken()
		if err != nil {
			log.Fatalln("couldn't get access token")
		}

		fmt.Printf("Token: %s\nToken Secret: %s\n", token, tokenSecret)
		os.Exit(0)
	}

	srv, err := NewServer()
	if err != nil {
		log.Fatalf("couldn't initialize the server: %v", err)
	}

	srv.Run()
}
