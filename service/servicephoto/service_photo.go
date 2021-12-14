package servicephoto

import (
	"github.com/arfan21/golang-mygram/entity"
	"github.com/arfan21/golang-mygram/model/modelphoto"
	"github.com/arfan21/golang-mygram/repository/repositoryphoto"
	"github.com/arfan21/golang-mygram/validation"
	"github.com/jinzhu/copier"
)

type ServicePhoto interface {
	Create(data modelphoto.Request) (modelphoto.Response, error)
	GetPhotos() ([]modelphoto.ResponseGet, error)
	Update(data modelphoto.Request, photoID int) (modelphoto.ResponseUpdate, error)
	Delete(photoID int) error
}

func New(photoRepo repositoryphoto.RepositoryPhoto) ServicePhoto {
	return &service{RepositoryPhoto: photoRepo}
}

type service struct {
	RepositoryPhoto repositoryphoto.RepositoryPhoto
}

func (service *service) Update(data modelphoto.Request, photoID int) (modelphoto.ResponseUpdate, error) {
	// validate update request
	err := validation.ValidatePhotoCreate(data)
	if err != nil {
		return modelphoto.ResponseUpdate{}, err
	}

	entityPhoto := entity.Photo{}
	copier.Copy(&entityPhoto, &data)
	entityPhoto.ID = uint(photoID)

	// call repository method to update Photo
	update, err := service.RepositoryPhoto.Update(entityPhoto)
	if err != nil {
		return modelphoto.ResponseUpdate{}, err
	}
	resp := modelphoto.ResponseUpdate{}
	copier.Copy(&resp, &update)
	return resp, nil
}

func (service *service) Delete(photoID int) error {
	err := service.RepositoryPhoto.Delete(photoID)
	if err != nil {
		return err
	}
	return nil
}

func (service *service) GetPhotos() ([]modelphoto.ResponseGet, error) {
	resPhotos, err := service.RepositoryPhoto.GetPhotos()

	if err != nil {
		return []modelphoto.ResponseGet{}, nil
	}

	var response []modelphoto.ResponseGet
	for _, photo := range resPhotos {
		tempResp := modelphoto.ResponseGet{}
		tempResp.ID = photo.ID
		tempResp.Title = photo.Title
		tempResp.Caption = photo.Caption
		tempResp.PhotoURL = photo.PhotoURL
		tempResp.CreatedAt = photo.CreatedAt
		tempResp.UpdatedAt = photo.UpdatedAt
		tempResp.User.Username = photo.User.Username
		tempResp.User.Email = photo.User.Email
		response = append(response, tempResp)
	}

	return response, nil
}

func (service *service) Create(data modelphoto.Request) (modelphoto.Response, error) {
	// validation input
	err := validation.ValidatePhotoCreate(data)
	if err != nil {
		return modelphoto.Response{}, err
	}

	entityPhoto := new(entity.Photo)

	copier.Copy(&entityPhoto, &data)

	create, err := service.RepositoryPhoto.Create(*entityPhoto)
	if err != nil {
		return modelphoto.Response{}, err
	}

	response := modelphoto.Response{}
	copier.Copy(&response, &create)

	return response, nil
}
