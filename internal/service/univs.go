package service

import (
	"fmt"
	"helloladies/internal/model"
	"helloladies/internal/repository"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

const (
	errCreateUniv     = "error while creating university"
	errListUnivs      = "error while listing universities"
	errUpdateUniv     = "error while updating university"
	errDeleteUniv     = "error while deleting university"
	errGetUnivsByCity = "error while getting univs by city"
)

type UnivsServiceImpl struct {
	univsRepo  repository.UnivsRepository
	citiesRepo repository.CitiesRepository
	log        *logrus.Logger
}

func NewUnivsService(univsRepo repository.UnivsRepository, citiesRepo repository.CitiesRepository, log *logrus.Logger) *UnivsServiceImpl {
	return &UnivsServiceImpl{
		univsRepo:  univsRepo,
		citiesRepo: citiesRepo,
		log:        log,
	}
}

func (s *UnivsServiceImpl) CreateUniv(cityCode string, univ model.Univ) error {
	cityId, err := s.citiesRepo.GetIdByCode(cityCode)
	if err != nil {
		s.log.Printf("s.citiesRepo.GetIdByCode: %s", err.Error())
		return fmt.Errorf(errIncorrectCode)
	}

	univDto := model.UnivDto{
		Id:        uuid.NewString(),
		Code:      generateCode(univ.Title),
		Title:     univ.Title,
		CityId:    cityId,
		Confirmed: true,
	}
	if err = s.univsRepo.CreateUniv(univDto); err != nil {
		s.log.Printf("s.univsRepo.CreateUniv: %s", err.Error())
		return fmt.Errorf(errCreateUniv)
	}
	return nil
}

func (s *UnivsServiceImpl) List() ([]model.Univ, error) {
	univDtos, err := s.univsRepo.List()
	if err != nil {
		s.log.Printf("s.citiesRepo.List: %s", err.Error())
		return []model.Univ{}, fmt.Errorf(errListUnivs)
	}

	var univs []model.Univ
	for _, univDto := range univDtos {
		univs = append(univs, model.UnivDtoToUniv(univDto))
	}
	return univs, nil
}

func (s *UnivsServiceImpl) UpdateUniv(cityCode string, univ model.Univ) (model.Univ, error) {

	univId, err := s.univsRepo.GetIdByCode(univ.Code)
	if err != nil {
		s.log.Printf("s.univsRepo.GetIdByCode: %s", err.Error())
		return model.Univ{}, fmt.Errorf(errIncorrectCode)
	}

	univDto := model.UnivDto{
		Id:        univId,
		Code:      univ.Code,
		Title:     univ.Title,
		Confirmed: true,
	}
	if _, err := s.univsRepo.UpdateUniv(univDto); err != nil {
		s.log.Printf("s.univsRepo.UpdateUniv: %s", err.Error())
		return model.Univ{}, fmt.Errorf(errUpdateUniv)
	}
	return univ, nil
}

func (s *UnivsServiceImpl) DeleteUniv(univCode string) error {
	univId, err := s.univsRepo.GetIdByCode(univCode)
	if err != nil {
		s.log.Printf("s.univsRepo.GetIdByCode: %s", err.Error())
		return fmt.Errorf(errIncorrectCode)
	}

	if err := s.univsRepo.DeleteUniv(univId); err != nil {
		s.log.Printf("s.univsRepo.DeleteUniv: %s", err.Error())
		return fmt.Errorf(errDeleteUniv)
	}
	return nil
}

func (s *UnivsServiceImpl) GetUnivsByCity(cityCode string) ([]model.Univ, error) {
	univDtos, err := s.univsRepo.GetUnivsByCity(cityCode)
	if err != nil {
		s.log.Printf("s.univsRepo.GetUnivsByCity: %s", err.Error())
		return []model.Univ{}, fmt.Errorf(errGetUnivsByCity)
	}

	var univs []model.Univ
	for _, univDto := range univDtos {
		univs = append(univs, model.UnivDtoToUniv(univDto))
	}
	return univs, nil
}
