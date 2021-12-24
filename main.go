package main

import (
	"exporter/controllers"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {

	var username = flag.String("username", "", "The 'username' option value")

	flag.Parse()

	if *username == "" {
		err := fmt.Errorf("username is required")
		log.Fatal(err)
	}

	action := os.Args[3]

	watchlistController := new(controllers.WatchlistController)

	if action == "json" {
		watchlistController.Create(*username)
	}

	if action == "history" {
		watchlistController.History(*username)
	}

	if action == "watchlist" {
		watchlistController.Watchlist(*username)
	}

	if action == "excel" {
		watchlistController.Excel(*username)
	}

	if action == "help" {
		fmt.Println("Usage:\n--username    The Username option value (default: null)\nHelp Options:\n--username <username> help     Show this help message\nApplication Options:\n --username <username> json     Save watchlist to json\n --username <username> history     10 recent actions\n --username <username> watchlist     Full watchlist\n --username <username> excel     Save watchlist to excel file")
	}
}
