package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/testisnullus/golang-project/pkg/models"
)

type DB struct {
	*sqlx.DB
}

func NewWeatherRepository(db *sqlx.DB) *DB {
	return &DB{db}
}

func (r *DB) InsertWeather(ctx context.Context, weather *models.CurrentWeather) error {
	err := r.DB.QueryRowContext(ctx, "INSERT INTO current_weather  (city, description, time, speed, temp, humidity) VALUES ($1,$2,$3,$4,$5,$6)", weather.City, weather.Description, weather.Time, weather.Speed, weather.Temp, weather.Humidity)
	if err != nil {
		return err.Err()
	}

	return nil
}

func (r *DB) FetchWeather(ctx context.Context) ([]*models.CurrentWeather, error) {
	CurrentWeather := make([]*models.CurrentWeather, 0)

	query := `SELECT * FROM current_weather`

	err := sqlx.SelectContext(ctx, r.DB, &CurrentWeather, query)
	if err != nil {
		return nil, err
	}

	return CurrentWeather, nil
}
