package servicesocialmedia

import (
	"errors"
	"testing"
	"time"

	"github.com/arfan21/golang-mygram/entity"
	"github.com/arfan21/golang-mygram/model/modelsocialmedia"
	"github.com/arfan21/golang-mygram/repository/repositoryphoto"
	"github.com/arfan21/golang-mygram/repository/repositorysocialmedia"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/jinzhu/copier"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

const ENV_TEST_PATH = "../../.env.test"

type ServiceUserTestSuite struct {
	suite.Suite
	repo           *repositorysocialmedia.RepositorySocialMediaMock
	repoPhoto      *repositoryphoto.RepositoryPhotoMock
	srv            ServiceSocialMedia
	defaultPayload modelsocialmedia.Request
}

func TestServiceUser(t *testing.T) {
	err := godotenv.Load(ENV_TEST_PATH)
	assert.NoError(t, err)
	repo := repositorysocialmedia.RepositorySocialMediaMock{}
	repoPhoto := repositoryphoto.RepositoryPhotoMock{}

	defaultPayload := modelsocialmedia.Request{
		Name:           "Test",
		SocialMediaUrl: "https://www.instagram.com/test",
	}

	srv := New(&repo, &repoPhoto)

	testSuite := &ServiceUserTestSuite{
		repo:           &repo,
		repoPhoto:      &repoPhoto,
		srv:            srv,
		defaultPayload: defaultPayload,
	}

	suite.Run(t, testSuite)
}

func (suite *ServiceUserTestSuite) Test_A_CreateSocialMedia() {
	suite.T().Run("Test Create Social Media Success", func(t *testing.T) {
		repoReturn := entity.SocialMedia{
			ID:             1,
			Name:           "Test",
			SocialMediaUrl: "https://www.instagram.com/test",
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
			UserID:         1,
		}

		suite.repo.On("Create", mock.Anything).Return(repoReturn, nil)

		createdSocialMedia, err := suite.srv.Create(suite.defaultPayload)
		assert.NoError(t, err)
		assert.NotNil(t, createdSocialMedia)
		assert.Equal(t, createdSocialMedia.ID, repoReturn.ID)
		assert.Equal(t, createdSocialMedia.Name, repoReturn.Name)
		assert.Equal(t, createdSocialMedia.SocialMediaUrl, repoReturn.SocialMediaUrl)
		assert.Equal(t, createdSocialMedia.UserID, repoReturn.UserID)
		suite.defaultPayload.ID = repoReturn.ID
	})

	suite.T().Run("Test Create Social Media Failed Validation", func(t *testing.T) {
		suite.repo.On("Create", mock.Anything).Return(entity.SocialMedia{}, errors.New("erro"))

		_, err := suite.srv.Create(modelsocialmedia.Request{})
		assert.Error(t, err)
		_, ok := err.(validation.Errors)
		assert.True(t, ok)
	})
}

func (suite *ServiceUserTestSuite) Test_B_GetList() {
	suite.T().Run("Test Get List Success", func(t *testing.T) {
		repoReturn := []entity.SocialMedia{
			{
				ID:             1,
				Name:           "Test",
				SocialMediaUrl: "https://www.instagram.com/test",
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
				UserID:         1,
				User: entity.User{
					ID:       1,
					Username: "test",
				},
			},
		}

		repoPhotoReturn := entity.Photo{
			ID:        1,
			Title:     "Test",
			Caption:   "Test",
			PhotoURL:  "https://www.photo.com/test.jpg",
			UserID:    1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		suite.repo.On("GetList", mock.Anything).Return(repoReturn, nil)
		suite.repoPhoto.On("GetPhotoByUserID", uint(1)).Return(repoPhotoReturn, nil)

		listSocialMedia, err := suite.srv.GetList()
		assert.NoError(t, err)
		assert.NotNil(t, listSocialMedia)
		assert.GreaterOrEqual(t, len(listSocialMedia.SocialMedias), 1)
		assert.Equal(t, listSocialMedia.SocialMedias[0].ID, repoReturn[0].ID)
		assert.Equal(t, listSocialMedia.SocialMedias[0].Name, repoReturn[0].Name)
		assert.Equal(t, listSocialMedia.SocialMedias[0].SocialMediaUrl, repoReturn[0].SocialMediaUrl)
		assert.Equal(t, listSocialMedia.SocialMedias[0].UserID, repoReturn[0].UserID)
		assert.Equal(t, listSocialMedia.SocialMedias[0].User.ProfileImageUrl, repoPhotoReturn.PhotoURL)
		assert.Equal(t, listSocialMedia.SocialMedias[0].User.Username, repoReturn[0].User.Username)
		assert.Equal(t, listSocialMedia.SocialMedias[0].User.ID, repoReturn[0].User.ID)

	})
}

func (suite *ServiceUserTestSuite) Test_C_UpdateByID() {
	suite.T().Run("Test Update By ID Success", func(t *testing.T) {
		repoReturn := entity.SocialMedia{
			ID:             1,
			Name:           "Test Update",
			SocialMediaUrl: "https://www.instagram.com/testupdate",
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
			UserID:         1,
		}

		tempPayload := modelsocialmedia.Request{}
		copier.Copy(&tempPayload, &suite.defaultPayload)
		tempPayload.Name = "Test Update"
		tempPayload.SocialMediaUrl = "https://www.instagram.com/testupdate"

		suite.repo.On("UpdateByID", mock.Anything).Return(repoReturn, nil).Once()

		updatedSocialMedia, err := suite.srv.UpdateByID(tempPayload)
		assert.NoError(t, err)
		assert.NotNil(t, updatedSocialMedia)
		assert.Equal(t, updatedSocialMedia.ID, repoReturn.ID)
		assert.Equal(t, updatedSocialMedia.Name, repoReturn.Name)
		assert.Equal(t, updatedSocialMedia.SocialMediaUrl, repoReturn.SocialMediaUrl)
		assert.Equal(t, updatedSocialMedia.UserID, repoReturn.UserID)
	})

	suite.T().Run("Test Update By ID failed, not found", func(t *testing.T) {
		tempPayload := modelsocialmedia.Request{}
		copier.Copy(&tempPayload, &suite.defaultPayload)
		tempPayload.ID = 22

		suite.repo.On("UpdateByID", mock.Anything).Return(entity.SocialMedia{}, errors.New("not found")).Once()

		_, err := suite.srv.UpdateByID(tempPayload)
		assert.Error(t, err)
	})

	suite.T().Run("Test Update By ID failed, validation failed", func(t *testing.T) {
		suite.repo.On("UpdateByID", mock.Anything).Return(entity.SocialMedia{}, errors.New("not found")).Once()

		_, err := suite.srv.UpdateByID(modelsocialmedia.Request{})
		assert.Error(t, err)
		_, ok := err.(validation.Errors)
		assert.True(t, ok)
	})
}

func (suite *ServiceUserTestSuite) Test_D_DeleteByID() {
	suite.T().Run("Test Delete By ID Success", func(t *testing.T) {
		suite.repo.On("DeleteByID", uint(1)).Return(nil)
		err := suite.srv.DeleteByID(suite.defaultPayload.ID)
		assert.NoError(t, err)
	})

	suite.T().Run("Test Delete By ID failed, not found", func(t *testing.T) {
		suite.repo.On("DeleteByID", uint(2)).Return(errors.New("not found"))
		err := suite.srv.DeleteByID(2)
		assert.Error(t, err)
	})
}
