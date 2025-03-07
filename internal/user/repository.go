package user

import (
	"fmt"
	"github.com/og11423074s/gocourse_domain/domain"
	"gorm.io/gorm"
	"log"
	"strings"
)

type Repository interface {
	Create(user *domain.User) error
	GetAll(filters Filters, offset, limit int) ([]domain.User, error)
	Get(id string) (*domain.User, error)
	DeleteById(id string) error
	Update(id string, firstName *string, lastName *string, email *string, phone *string) error
	Count(filters Filters) (int, error)
}

type repo struct {
	log *log.Logger
	db  *gorm.DB
}

func NewRepo(log *log.Logger, db *gorm.DB) Repository {
	return &repo{
		log: log,
		db:  db,
	}
}

func (r *repo) Create(user *domain.User) error {

	if err := r.db.Create(user); err.Error != nil {
		r.log.Println(err.Error)
		return err.Error
	}
	r.log.Println("User created with id: ", user.ID)
	return nil
}

func (r *repo) GetAll(filters Filters, offset, limit int) ([]domain.User, error) {
	var users []domain.User

	tx := r.db.Model(&users)
	tx = applyFilters(tx, filters)
	tx = tx.Offset(offset).Limit(limit)

	result := tx.Order("created_at desc").Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

func (r *repo) Get(id string) (*domain.User, error) {
	user := domain.User{ID: id}

	err := r.db.First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *repo) DeleteById(id string) error {
	user := domain.User{ID: id}

	err := r.db.Delete(&user).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) Update(id string, firstName *string, lastName *string, email *string, phone *string) error {

	values := make(map[string]interface{})

	if firstName != nil {
		values["first_name"] = *firstName
	}

	if lastName != nil {
		values["last_name"] = *lastName
	}

	if email != nil {
		values["email"] = *email
	}

	if phone != nil {
		values["phone"] = *phone
	}

	if err := r.db.Model(&domain.User{}).Where("id = ?", id).Updates(values).Error; err != nil {
		return err
	}

	return nil

}

func (r *repo) Count(filters Filters) (int, error) {
	var count int64
	tx := r.db.Model(&domain.User{})
	tx = applyFilters(tx, filters)
	result := tx.Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}
	return int(count), nil
}

func applyFilters(tx *gorm.DB, filters Filters) *gorm.DB {
	if filters.FirstName != "" {
		filters.FirstName = fmt.Sprintf("%%%s%%", strings.ToLower(filters.FirstName))
		tx = tx.Where("lower(first_name) like ?", filters.FirstName)
	}

	if filters.LastName != "" {
		filters.LastName = fmt.Sprintf("%%%s%%", strings.ToLower(filters.LastName))
		tx = tx.Where("lower(last_name) like ?", filters.LastName)
	}

	return tx
}
