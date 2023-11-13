package service

import (
	"context"
	"github.com/testisnullus/golang-project/pkg/models"
	"github.com/testisnullus/golang-project/pkg/weather/repository"
)

type Weather interface {
	Create(ctx context.Context, weather *models.CurrentWeather) error
	Get(ctx context.Context) ([]*models.CurrentWeather, error)
}

type Service struct {
	Weather
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Weather: NewWeatherService(repos),
	}
}
