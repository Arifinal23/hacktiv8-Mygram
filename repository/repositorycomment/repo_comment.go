package repositorycomment

import (
	"github.com/arfan21/golang-mygram/entity"
	"gorm.io/gorm"
)

type RepositoryComment interface {
	Create(data entity.Comment) (entity.Comment, error)
	Get() ([]entity.Comment, error)
	Update(data entity.Comment) (entity.Comment, error)
	Delete(commentID uint) error
}

type repository struct {
	db *gorm.DB
}

func (r repository) Create(data entity.Comment) (entity.Comment, error) {
	err := r.db.Create(&data).Error
	if err != nil {
		return entity.Comment{}, err
	}
	return data, nil
}

func (r repository) Get() ([]entity.Comment, error) {
	var comments []entity.Comment
	err := r.db.Preload("User").Preload("Photo").Find(&comments).Error
	if err != nil {
		return []entity.Comment{}, err
	}
	return comments, nil
}

func (r *repository) Update(data entity.Comment) (entity.Comment, error) {
	err := r.db.Updates(&data).First(&data).Error
	if err != nil {
		return entity.Comment{}, err
	}
	return data, nil
}

// Delete comment by comment id return a error or nil
func (r repository) Delete(commentID uint) error {
	var comment entity.Comment
	err := r.db.Where("id = ?", commentID).Delete(&comment).Error
	if err != nil {
		return err
	}
	return nil
}

func New(db *gorm.DB) RepositoryComment {
	return &repository{db: db}
}
