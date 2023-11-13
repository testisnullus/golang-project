package main

import (
	"context"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/testisnullus/golang-project/pkg/config"
	"github.com/testisnullus/golang-project/pkg/weather/handlers"
	"github.com/testisnullus/golang-project/pkg/weather/repository"
	"github.com/testisnullus/golang-project/pkg/weather/service"
	"log"
	"net/http"
	"time"
)

var cities = []string{"Dnipro", "New York", "Paris"}

func main() {
	yaml, err := config.GetConfig("./cmd/weather/config.yaml")
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
	serv := service.NewService(repository)
	handler := handlers.NewHandlers(serv, yaml)

	go func() {
		for {
			for _, city := range cities {
				data, err := handlers.GetCurrentWeather(yaml.ApiKey, city, "", "", "")
				if err != nil {
					panic(err)
				}

				repository.InsertWeather(context.Background(), data)
			}
			time.Sleep(time.Hour * 1)
		}
	}()

	err = http.ListenAndServe(yaml.ServerPort, handler.Router())
	if err != nil {
		log.Fatalf("Cant't run http server: %s", err.Error())
	}
}
