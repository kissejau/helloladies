package service

import (
	"fmt"
	"helloladies/internal/model"
	"helloladies/internal/repository"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

const (
	errCreateCity = "error while creating city"
	errListCities = "error while listing cities"
	errUpdateCity = "error while updating city"
	errDeleteCity = "error while deleting city"
)

type CitiesServiceImpl struct {
	citiesRepo repository.CitiesRepository
	log        *logrus.Logger
}

func NewCitiesService(citiesRepo repository.CitiesRepository, log *logrus.Logger) *CitiesServiceImpl {
	return &CitiesServiceImpl{
		citiesRepo: citiesRepo,
		log:        log,
	}
}

func (s *CitiesServiceImpl) CreateCity(city model.City) error {
	cityDto := model.CityDto{
		Id:        uuid.NewString(),
		Code:      generateCode(city.Title),
		Title:     city.Title,
		Confirmed: true,
	}

	if err := s.citiesRepo.CreateCity(cityDto); err != nil {
		s.log.Printf("s.citiesRepo.CreateCity: %s", err.Error())
		return fmt.Errorf(errCreateCity)
	}

	return nil
}

func (s *CitiesServiceImpl) List() ([]model.City, error) {
	cityDtos, err := s.citiesRepo.List()
	if err != nil {
		s.log.Printf("s.citiesRepo.List: %s", err.Error())
		return []model.City{}, fmt.Errorf(errListCities)
	}

	var cities []model.City
	for _, cityDto := range cityDtos {
		cities = append(cities, model.CityDtoToCity(cityDto))
	}
	return cities, nil
}

func (s *CitiesServiceImpl) UpdateCity(code string, city model.City) (model.City, error) {
	id, err := s.citiesRepo.GetIdByCode(code)
	if err != nil {
		s.log.Printf("s.citiesRepo.GetIdByCode: %s", err.Error())
		return model.City{}, fmt.Errorf(errIncorrectCode)
	}

	cityDto := model.CityDto{
		Id:        id,
		Code:      code,
		Title:     city.Title,
		Confirmed: true,
	}

	if _, err = s.citiesRepo.UpdateCity(id, cityDto); err != nil {
		return model.City{}, fmt.Errorf(errUpdateCity)
	}
	return city, nil
}

func (s *CitiesServiceImpl) DeleteCity(code string) error {
	id, err := s.citiesRepo.GetIdByCode(code)
	if err != nil {
		s.log.Printf("s.citiesRepo.GetIdByCode: %s", err.Error())
		return fmt.Errorf(errIncorrectCode)
	}

	if err = s.citiesRepo.DeleteCity(id); err != nil {
		return fmt.Errorf(errDeleteCity)
	}
	return nil
}
