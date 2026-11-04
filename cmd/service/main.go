package main

import (
	"log"

	"github.com/Rodco-off/Birwi/internal/config"
	"github.com/Rodco-off/Birwi/internal/db"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	conn, err := db.Connect(cfg)
	if err != nil {
		log.Fatal(err)
	}
}
