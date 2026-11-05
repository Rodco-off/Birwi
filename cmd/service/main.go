package main

import (
	"log"

	"github.com/Rodco-off/Birwi/internal/config"
	"github.com/Rodco-off/Birwi/internal/db"
	"github.com/Rodco-off/Birwi/internal/repository"
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

	companyRepo := repository.NewCompanyRepository(conn)
	companyReportRepo := repository.NewCompanyReportRepository(conn)

}
