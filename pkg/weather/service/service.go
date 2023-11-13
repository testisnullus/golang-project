package service

import (
	"context"

	"github.com/testisnullus/golang-project/pkg/models"
	"github.com/testisnullus/golang-project/pkg/weather/repository"
)

type WeatherService struct {
	repo *repository.Repository
}

func NewWeatherService(repo *repository.Repository) *WeatherService {
	return &WeatherService{repo: repo}
}

func (s *WeatherService) Create(ctx context.Context, weather *models.CurrentWeather) error {
	err := s.repo.InsertWeather(ctx, weather)
	if err != nil {
		return err
	}

	return nil
}

func (s *WeatherService) Get(ctx context.Context) ([]*models.CurrentWeather, error) {
	result, err := s.repo.FetchWeather(ctx)
	if err != nil {
		return nil, err
	}

	return result, nil
}
