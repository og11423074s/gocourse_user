package user

import (
	"context"
	"fmt"
	"github.com/og11423074s/gocourse_domain/domain"
	"gorm.io/gorm"
	"log"
	"strings"
)

type Repository interface {
	Create(ctx context.Context, user *domain.User) error
	GetAll(ctx context.Context, filters Filters, offset, limit int) ([]domain.User, error)
	Get(ctx context.Context, id string) (*domain.User, error)
	DeleteById(ctx context.Context, id string) error
	Update(ctx context.Context, id string, firstName *string, lastName *string, email *string, phone *string) error
	Count(ctx context.Context, filters Filters) (int, error)
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

func (r *repo) Create(ctx context.Context, user *domain.User) error {

	if err := r.db.WithContext(ctx).Create(user); err.Error != nil {
		r.log.Println(err.Error)
		return err.Error
	}
	r.log.Println("User created with id: ", user.ID)
	return nil
}

func (r *repo) GetAll(ctx context.Context, filters Filters, offset, limit int) ([]domain.User, error) {
	var users []domain.User

	tx := r.db.WithContext(ctx).Model(&users)
	tx = applyFilters(tx, filters)
	tx = tx.Offset(offset).Limit(limit)

	result := tx.Order("created_at desc").Find(&users)

	if result.Error != nil {
		r.log.Println(result.Error)
		return nil, result.Error
	}

	return users, nil
}

func (r *repo) Get(ctx context.Context, id string) (*domain.User, error) {
	user := domain.User{ID: id}

	err := r.db.WithContext(ctx).First(&user).Error
	if err != nil {
		r.log.Println(err)
		return nil, err
	}

	return &user, nil
}

func (r *repo) DeleteById(ctx context.Context, id string) error {
	user := domain.User{ID: id}

	result := r.db.WithContext(ctx).Delete(&user)

	if result.Error != nil {
		r.log.Println(result.Error)
		return result.Error
	}

	if result.RowsAffected == 0 {
		r.log.Printf("user %s doesn't exists", id)
		return fmt.Errorf("user %s doesn't exists", id)
	}

	return nil
}

func (r *repo) Update(ctx context.Context, id string, firstName *string, lastName *string, email *string, phone *string) error {

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

	result := r.db.WithContext(ctx).Model(&domain.User{}).Where("id = ?", id).Updates(values)

	if result.Error != nil {
		r.log.Println(result.Error)
		return result.Error
	}

	if result.RowsAffected == 0 {
		r.log.Printf("user %s doesn't exists", id)
		return fmt.Errorf("user %s doesn't exists", id)
	}

	return nil

}

func (r *repo) Count(ctx context.Context, filters Filters) (int, error) {
	var count int64
	tx := r.db.WithContext(ctx).Model(&domain.User{})
	tx = applyFilters(tx, filters)
	if err := tx.Count(&count).Error; err != nil {
		r.log.Println(err)
		return 0, err
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
