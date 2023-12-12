package service

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"helloladies/internal/model"
	"helloladies/internal/repository"
)

const (
	errCreateTeacher     = "error while creating teacher"
	errListTeachers      = "error while listing teachers"
	errUpdateTeacher     = "error while updating teacher"
	errDeleteTeacher     = "error while deleting teacher"
	errGetTeachersByUniv = "error while getting teachers by university"
)

type TeachersServiceImpl struct {
	univsRepo    repository.UnivsRepository
	teachersRepo repository.TeachersRepository
	log          *logrus.Logger
}

func NewTeachersService(univsRepo repository.UnivsRepository, teachersRepo repository.TeachersRepository, log *logrus.Logger) *TeachersServiceImpl {
	return &TeachersServiceImpl{
		univsRepo:    univsRepo,
		teachersRepo: teachersRepo,
		log:          log,
	}
}

func (s *TeachersServiceImpl) CreateTeacher(univCode string, teacher model.Teacher) error {

	univId, err := s.univsRepo.GetIdByCode(univCode)
	if err != nil {
		s.log.Printf("s.univsRepo.GetIdByCode: %s", err.Error())
		return fmt.Errorf(errIncorrectCode)
	}

	teacherDto := model.TeacherDto{
		Id:         uuid.NewString(),
		Code:       generateCode(teacher.Surname),
		Name:       teacher.Name,
		Surname:    teacher.Surname,
		Patronymic: teacher.Patronymic,
		Degree:     teacher.Degree,
		UnivId:     univId,
		Confirmed:  true,
	}
	if err := s.teachersRepo.CreateTeacher(teacherDto); err != nil {
		s.log.Printf("s.teacherRepo.CreateTeacher: %s", err.Error())
		return fmt.Errorf(errCreateTeacher)
	}
	return nil
}

func (s *TeachersServiceImpl) List() ([]model.Teacher, error) {
	teacherDtos, err := s.teachersRepo.List()
	if err != nil {
		s.log.Printf("s.teachersRepo.List: %s", err.Error())
		return []model.Teacher{}, fmt.Errorf(errListTeachers)
	}

	var teachers []model.Teacher
	for _, teacherDto := range teacherDtos {
		teachers = append(teachers, model.TeacherDtoToTeacher(teacherDto))
	}
	return teachers, nil
}

func (s *TeachersServiceImpl) UpdateTeacher(univCode string, teacher model.Teacher) (model.Teacher, error) {
	univId, err := s.univsRepo.GetIdByCode(univCode)
	if err != nil {
		s.log.Printf("s.univsRepo.GetIdByCode: %s", err.Error())
		return model.Teacher{}, fmt.Errorf(errIncorrectCode)
	}

	teacherId, err := s.teachersRepo.GetIdByCode(teacher.Code)
	if err != nil {
		s.log.Printf("s.teachersRepo.GetIdByCode: %s", err.Error())
		return model.Teacher{}, fmt.Errorf(errIncorrectCode)
	}

	teacherDto := model.TeacherDto{
		Id:         teacherId,
		Code:       teacher.Code,
		Name:       teacher.Name,
		Surname:    teacher.Surname,
		Patronymic: teacher.Patronymic,
		Degree:     teacher.Degree,
		UnivId:     univId,
		Confirmed:  true,
	}
	if _, err := s.teachersRepo.UpdateTeacher(teacherDto); err != nil {
		s.log.Printf("s.teachersRepo.UpdateTeacher: %s", err.Error())
		return model.Teacher{}, fmt.Errorf(errUpdateTeacher)
	}
	return teacher, nil
}

func (s *TeachersServiceImpl) DeleteTeacher(teacherCode string) error {
	teacherId, err := s.teachersRepo.GetIdByCode(teacherCode)
	if err != nil {
		s.log.Printf("s.teachersRepo.GetIdByCode: %s", err.Error())
		return fmt.Errorf(errIncorrectCode)
	}

	if err := s.teachersRepo.DeleteTeacher(teacherId); err != nil {
		s.log.Printf("s.DeleteTeacher: %s", err.Error())
		return fmt.Errorf(errDeleteTeacher)
	}
	return nil
}

func (s *TeachersServiceImpl) GetTeachersByUniv(univCode string) ([]model.Teacher, error) {
	teacherDtos, err := s.teachersRepo.GetTeachersByUniv(univCode)
	if err != nil {
		s.log.Printf("s.teachersRepo.GetTeachersByUniv: %s", err.Error())
		return []model.Teacher{}, fmt.Errorf(errGetTeachersByUniv)
	}

	var teachers []model.Teacher
	for _, teacherDto := range teacherDtos {
		teachers = append(teachers, model.TeacherDtoToTeacher(teacherDto))
	}
	return teachers, nil
}
