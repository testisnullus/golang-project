package main

import (
	"github.com/jmoiron/sqlx"
	"github.com/testisnullus/golang-project/pkg/config"
	"github.com/testisnullus/golang-project/pkg/users/handlers"
	"github.com/testisnullus/golang-project/pkg/users/repository"
	"github.com/testisnullus/golang-project/pkg/users/service"
	"log"
	"net/http"

	_ "github.com/jackc/pgx/stdlib"
)

func main() {
	yaml, err := config.GetConfig("./cmd/users/config.yaml")
	if err != nil {
		log.Fatalf("Can't load config file, err: %s", err.Error())
	}

	postgresURL := config.PostgresURL(yaml)

	db, err := sqlx.Open("pgx", postgresURL)
	if err != nil {
		log.Fatalf("Cant up postgres connection, err: %s", err.Error())
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Can't ping the database, err: %s", err.Error())
	}

	repository := repository.NewRepository(db)
	services := service.NewService(repository, yaml)
	handlers := handlers.NewHandlers(services)

	err = http.ListenAndServe(yaml.ServerPort, handlers.Router())
	if err != nil {
		log.Fatalf("Cant't run http server: %s", err.Error())
	}
}
