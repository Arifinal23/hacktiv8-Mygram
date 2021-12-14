package repositoryuser

import (
	"testing"

	"github.com/arfan21/golang-mygram/config/configdb"
	"github.com/arfan21/golang-mygram/entity"
	"github.com/jinzhu/copier"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

const ENV_TEST_PATH = "../../.env.test"

type RepoUserTestSuite struct {
	suite.Suite
	db             *gorm.DB
	repo           RepositoryUser
	defaultPayload entity.User
}

func TestRepositoryUser(t *testing.T) {
	err := godotenv.Load(ENV_TEST_PATH)
	assert.NoError(t, err)
	//init db
	db, err := configdb.New()
	assert.NoError(t, err)

	db.AutoMigrate(entity.User{})
	clearDbUser(db)

	repo := New(db)

	defaultPayload := entity.User{
		Username: "test",
		Email:    "test@example.com",
		Password: "123123123",
		Age:      8,
	}

	testSuite := &RepoUserTestSuite{
		db:             db,
		repo:           repo,
		defaultPayload: defaultPayload,
	}
	suite.Run(t, testSuite)
}

func (suite *RepoUserTestSuite) Test_A_CreateUser() {
	suite.T().Run("Create User Success", func(t *testing.T) {

		createdUser, err := suite.repo.Create(suite.defaultPayload)
		assert.NoError(t, err)
		assert.NotEmpty(t, createdUser.ID)
		suite.defaultPayload.ID = createdUser.ID
	})

	suite.T().Run("Create User Error Duplicate", func(t *testing.T) {
		_, err := suite.repo.Create(suite.defaultPayload)
		assert.Error(t, err)
	})
}

func (suite *RepoUserTestSuite) Test_B_IsEmailExist() {

	suite.T().Run("Is Email Exist Success", func(t *testing.T) {

		err := suite.repo.IsEmailExist(suite.defaultPayload.Email)
		if assert.Error(t, err) {
			assert.Equal(t, err.Error(), "email already exists")
		}
	})

	suite.T().Run("Is Email Not Found", func(t *testing.T) {

		err := suite.repo.IsEmailExist("test2@example.com")
		assert.NoError(t, err)
	})
}

func (suite *RepoUserTestSuite) Test_C_Login() {
	suite.T().Run("Login success", func(t *testing.T) {
		data, err := suite.repo.Login(suite.defaultPayload.Email)
		assert.NoError(t, err)
		assert.NotEmpty(t, data)
		assert.Equal(t, data.Email, suite.defaultPayload.Email)
	})

	suite.T().Run("Login failed", func(t *testing.T) {
		data, err := suite.repo.Login("test2@example.com")
		if assert.Error(t, err) {
			assert.Empty(t, data)
		}

	})
}

func (suite *RepoUserTestSuite) Test_D_Update() {
	suite.T().Run("Update success", func(t *testing.T) {
		updatePayload := entity.User{}
		copier.Copy(&updatePayload, &suite.defaultPayload)
		updatePayload.Username = "update"
		updatePayload.Email = "update@example.com"
		updatedUser, err := suite.repo.Update(updatePayload)
		assert.NoError(t, err)
		assert.NotEmpty(t, suite.defaultPayload.ID)
		assert.NotEqual(t, updatedUser.Username, suite.defaultPayload.Username)
		assert.NotEqual(t, updatedUser.UpdatedAt, suite.defaultPayload.UpdatedAt)
	})

	suite.T().Run("Update Failed", func(t *testing.T) {
		updatePayload := entity.User{}
		copier.Copy(&updatePayload, &suite.defaultPayload)
		updatePayload.Username = "updatefail"
		updatePayload.Email = "updatefail@example.com"
		updatePayload.ID = 0
		_, err := suite.repo.Update(updatePayload)
		assert.Error(t, err)
	})
}

func (suite *RepoUserTestSuite) Test_E_Delete() {
	suite.T().Run("Delete Success", func(t *testing.T) {
		err := suite.repo.DeleteByID(suite.defaultPayload.ID)
		assert.NoError(t, err)
	})

	suite.T().Run("Delete Failed", func(t *testing.T) {
		err := suite.repo.DeleteByID(0)
		assert.Error(t, err)
	})
}

func clearDbUser(db *gorm.DB) {
	db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(entity.User{})
}
