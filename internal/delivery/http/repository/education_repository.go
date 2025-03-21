package repository

import (
	"jobseeker/internal/entity"

	"gorm.io/gorm"
)

type (
	EducationRepository interface {
		Create(tx *gorm.DB, education *entity.Education) error
		Update(tx *gorm.DB, id string, education *entity.Education) error
		FindByID(tx *gorm.DB, id string, education *entity.Education) error
		FindByIDAndUserID(tx *gorm.DB, id, userID string, education *entity.Education) error
		ExistsByIDAndUserID(tx *gorm.DB, id, userID string) (bool, error)
		DeleteByID(tx *gorm.DB, id string) error
		GetAllByUserID(tx *gorm.DB, userID string) ([]entity.Education, error)
	}

	educationRepository struct {
		db *gorm.DB
	}
)

func NewEducationRepository(db *gorm.DB) EducationRepository {
	return &educationRepository{db}
}

func (r *educationRepository) Create(tx *gorm.DB, education *entity.Education) error {
	if tx == nil {
		tx = r.db
	}
	return tx.Create(education).Error
}

func (r *educationRepository) FindByID(tx *gorm.DB, id string, education *entity.Education) error {
	if tx == nil {
		tx = r.db
	}
	return tx.Where("id = ?", id).First(education).Error
}

func (r *educationRepository) Update(tx *gorm.DB, id string, education *entity.Education) error {
	if tx == nil {
		tx = r.db
	}
	return tx.Model(education).Where("id = ?", id).Updates(education).Error
}

func (r *educationRepository) FindByIDAndUserID(tx *gorm.DB, id, userID string, education *entity.Education) error {
	if tx == nil {
		tx = r.db
	}
	return tx.Where("id = ? AND user_id = ?", id, userID).First(education).Error
}

func (r *educationRepository) ExistsByIDAndUserID(tx *gorm.DB, id, userID string) (bool, error) {
	if tx == nil {
		tx = r.db
	}
	var count int64
	err := tx.Model(&entity.Education{}).Where("id = ? AND user_id = ?", id, userID).Count(&count).Error
	return count > 0, err
}

func (r *educationRepository) DeleteByID(tx *gorm.DB, id string) error {
	if tx == nil {
		tx = r.db
	}
	return tx.Where("id = ?", id).Delete(&entity.Education{}).Error
}

func (r *educationRepository) GetAllByUserID(tx *gorm.DB, userID string) ([]entity.Education, error) {
	if tx == nil {
		tx = r.db
	}
	var educations []entity.Education
	err := tx.Where("user_id = ?", userID).Find(&educations).Error
	return educations, err
}
