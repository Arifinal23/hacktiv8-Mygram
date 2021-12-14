package repositorysocialmedia

import (
	"testing"

	"github.com/arfan21/golang-mygram/config/configdb"
	"github.com/arfan21/golang-mygram/entity"
	"github.com/arfan21/golang-mygram/repository/repositoryuser"
	"github.com/jinzhu/copier"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

const ENV_TEST_PATH = "../../.env.test"

type RepoSocialMediaTestSuite struct {
	suite.Suite
	db             *gorm.DB
	repo           RepositorySocialMedia
	defaultPayload entity.SocialMedia
}

func TestRepositorySocialMedia(t *testing.T) {
	err := godotenv.Load(ENV_TEST_PATH)
	assert.NoError(t, err)
	//init db
	db, err := configdb.New()
	assert.NoError(t, err)

	db.AutoMigrate(entity.SocialMedia{})
	clearDb(db)

	repo := New(db)

	initialUser := entity.User{
		Username: "test",
		Email:    "test@example.com",
		Password: "123123123",
		Age:      8,
	}

	userRepo := repositoryuser.New(db)
	createdUser, err := userRepo.Create(initialUser)
	assert.NoError(t, err)

	defaultPayload := entity.SocialMedia{
		Name:           "sosmed",
		SocialMediaUrl: "https://socialmedia.com/sosmed",
		UserID:         createdUser.ID,
	}

	testSuite := &RepoSocialMediaTestSuite{
		db:             db,
		repo:           repo,
		defaultPayload: defaultPayload,
	}
	suite.Run(t, testSuite)
}

func (suite *RepoSocialMediaTestSuite) Test_A_CreateSocialMedia() {
	suite.T().Run("Create Social Media Success", func(t *testing.T) {
		createdSocialMedia, err := suite.repo.Create(suite.defaultPayload)
		assert.NoError(t, err)
		assert.NotEmpty(t, createdSocialMedia.ID)
		suite.defaultPayload.ID = createdSocialMedia.ID
	})

	suite.T().Run("Create Social Media Error Empty", func(t *testing.T) {
		_, err := suite.repo.Create(entity.SocialMedia{})
		assert.Error(t, err)
	})
}

func (suite *RepoSocialMediaTestSuite) Test_B_GetSocialMedias() {
	suite.T().Run("Get Social Medias Success", func(t *testing.T) {

		socialMedias, err := suite.repo.GetList()
		assert.NoError(t, err)
		assert.GreaterOrEqual(t, len(socialMedias), 1)
	})

}

func (suite *RepoSocialMediaTestSuite) Test_B_UpdateByID() {
	suite.T().Run("Update Social Media Success", func(t *testing.T) {
		tempPayload := entity.SocialMedia{}
		copier.Copy(&tempPayload, &suite.defaultPayload)
		tempPayload.Name = "updated"
		tempPayload.SocialMediaUrl = "https://updated.com/updated"
		updatedSocialMedia, err := suite.repo.UpdateByID(tempPayload)
		assert.NoError(t, err)
		assert.Equal(t, updatedSocialMedia.ID, suite.defaultPayload.ID)
		assert.NotEqual(t, updatedSocialMedia.Name, suite.defaultPayload.Name)
		assert.NotEqual(t, updatedSocialMedia.SocialMediaUrl, suite.defaultPayload.SocialMediaUrl)
	})
}

func (suite *RepoSocialMediaTestSuite) Test_C_DeleteByID() {
	suite.T().Run("Delete Social Media Success", func(t *testing.T) {
		err := suite.repo.DeleteByID(suite.defaultPayload.ID)
		assert.NoError(t, err)
	})

	suite.T().Run("Delete Social Media Error Not Found", func(t *testing.T) {
		err := suite.repo.DeleteByID(0)
		assert.Error(t, err)
	})
}

func clearDb(db *gorm.DB) {
	db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(entity.SocialMedia{})
	db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(entity.User{})
}
© 2021 GitHub, Inc.
Terms
Privacy
