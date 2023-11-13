package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/testisnullus/golang-project/pkg/models"
)

type Weather interface {
	InsertWeather(ctx context.Context, weather *models.CurrentWeather) error
	FetchWeather(ctx context.Context) ([]*models.CurrentWeather, error)
}

type Repository struct {
	Weather
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Weather: NewWeatherRepository(db),
	}
}
