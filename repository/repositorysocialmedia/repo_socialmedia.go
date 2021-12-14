package repositorysocialmedia

import (
	"github.com/arfan21/golang-mygram/entity"
	"gorm.io/gorm"
)

type RepositorySocialMedia interface {
	Create(data entity.SocialMedia) (entity.SocialMedia, error)
	GetList() ([]entity.SocialMedia, error)
	UpdateByID(data entity.SocialMedia) (entity.SocialMedia, error)
	DeleteByID(id uint) error
}

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) RepositorySocialMedia {
	return &repository{db: db}
}

func (r *repository) Create(data entity.SocialMedia) (entity.SocialMedia, error) {
	err := r.db.Create(&data).Error
	if err != nil {
		return entity.SocialMedia{}, err
	}
	return data, nil
}

func (r *repository) GetList() ([]entity.SocialMedia, error) {
	var socialMedia []entity.SocialMedia
	err := r.db.Preload("User").Find(&socialMedia).Error
	if err != nil {
		return nil, err
	}
	return socialMedia, nil
}

func (r *repository) UpdateByID(data entity.SocialMedia) (entity.SocialMedia, error) {
	err := r.db.Model(&data).Updates(&data).First(&data).Error
	if err != nil {
		return entity.SocialMedia{}, err
	}
	return data, nil
}

func (r *repository) DeleteByID(id uint) error {
	socialMedia := new(entity.SocialMedia)
	socialMedia.ID = id
	return r.db.First(&socialMedia).Where("id = ?", socialMedia.ID).Delete(&socialMedia).Error
}
