package user

import (
	"github.com/og11423074s/gocourse_domain/domain"
	"log"
)

type (
	Filters struct {
		FirstName string
		LastName  string
	}

	Service interface {
		Create(firstName, lastName, email, phone string) (*domain.User, error)
		Get(id string) (*domain.User, error)
		GetAll(filters Filters, offset, limit int) ([]domain.User, error)
		Delete(id string) error
		Update(id string, firstName *string, lastName *string, email *string, phone *string) error
		Count(filters Filters) (int, error)
	}
	service struct {
		log  *log.Logger
		repo Repository
	}
)

func NewService(log *log.Logger, repo Repository) Service {
	return &service{
		log:  log,
		repo: repo,
	}
}

func (s service) Create(firstName, lastName, email, phone string) (*domain.User, error) {
	s.log.Println("Create user service")

	user := domain.User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Phone:     phone,
	}

	if err := s.repo.Create(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s service) GetAll(filters Filters, offset, limit int) ([]domain.User, error) {
	s.log.Println("Get all user service")

	users, err := s.repo.GetAll(filters, offset, limit)

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s service) Get(id string) (*domain.User, error) {
	s.log.Println("Get user service")

	user, err := s.repo.Get(id)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s service) Delete(id string) error {

	_, err := s.repo.Get(id)

	if err != nil {
		return err
	}

	return s.repo.DeleteById(id)
}

func (s service) Update(id string, firstName *string, lastName *string, email *string, phone *string) error {
	return s.repo.Update(id, firstName, lastName, email, phone)
}

func (s service) Count(filters Filters) (int, error) {
	return s.repo.Count(filters)
}
