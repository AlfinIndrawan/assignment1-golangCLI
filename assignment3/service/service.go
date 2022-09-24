package service

import (
	"assignment3/helpers"
	"assignment3/models"
	"assignment3/repository"
	"time"
)

type Service interface {
	GetData() (models.Response, error)
	UpdateData() error
}

type service struct {
	repository repository.Repository
}

func NewService(r repository.Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetData() (models.Response, error) {
	return s.repository.GetData()
}

func (s *service) UpdateData() error {
	var response models.Response
	response.Data.Water = helpers.GenerateRandom(1, 30)
	response.Data.Air = helpers.GenerateRandom(1, 30)

	response.Status = models.DataStatus{
		WaterStatus: helpers.CheckWater(response.Data.Water),
		AirStatus:   helpers.CheckAir(response.Data.Air),
	}

	response.UpdatedAt = time.Now()

	return s.repository.UpdateData(response)
}
