package serviceuser

import (
	"errors"
	"testing"

	"github.com/arfan21/golang-mygram/entity"
	"github.com/arfan21/golang-mygram/model/modeluser"
	"github.com/arfan21/golang-mygram/repository/repositoryuser"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/jinzhu/copier"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const ENV_TEST_PATH = "../../.env.test"

type ServiceUserTestSuite struct {
	suite.Suite
	repo           *repositoryuser.RepositoryUserMock
	srv            ServiceUser
	defaultPayload modeluser.Request
}

func TestServiceUser(t *testing.T) {
	err := godotenv.Load(ENV_TEST_PATH)
	assert.NoError(t, err)
	repo := repositoryuser.RepositoryUserMock{}

	defaultPayload := modeluser.Request{
		Username: "test",
		Email:    "test@example.com",
		Password: "A1231232221111",
		Age:      8,
	}

	srv := New(&repo)

	testSuite := &ServiceUserTestSuite{
		repo:           &repo,
		srv:            srv,
		defaultPayload: defaultPayload,
	}

	suite.Run(t, testSuite)
}

func (suite *ServiceUserTestSuite) Test_A_CreateUser() {
	suite.T().Run("Test create user success", func(t *testing.T) {
		repoReturn := entity.User{}
		copier.Copy(&repoReturn, &suite.defaultPayload)
		repoReturn.ID = 1
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(repoReturn.Password), bcrypt.DefaultCost)
		assert.NoError(t, err)
		repoReturn.Password = string(hashedPassword)

		suite.repo.On("IsEmailExist", mock.Anything).Return(nil).Once()
		suite.repo.On("Create", mock.Anything).Return(repoReturn, nil).Once()
		resp, err := suite.srv.Create(suite.defaultPayload)
		if assert.NoError(t, err) {
			assert.Equal(t, resp.ID, repoReturn.ID)
			assert.NotEmpty(t, resp)
			suite.defaultPayload.ID = resp.ID
		}
	})

	suite.T().Run("Test create user email duplicate", func(t *testing.T) {
		repoReturn := entity.User{}
		copier.Copy(&repoReturn, &suite.defaultPayload)
		repoReturn.ID = 1

		suite.repo.On("IsEmailExist", mock.Anything).Return(errors.New("email already exists")).Once()
		suite.repo.On("Create", mock.Anything).Return(repoReturn, nil).Once()
		_, err := suite.srv.Create(suite.defaultPayload)
		if assert.Error(t, err) {
			_, ok := err.(validation.Errors)
			assert.True(t, ok)
		}
	})

	suite.T().Run("Test create user payload empty", func(t *testing.T) {
		repoReturn := entity.User{}

		suite.repo.On("IsEmailExist", mock.Anything).Return(nil).Once()
		suite.repo.On("Create", mock.Anything).Return(repoReturn, nil).Once()
		_, err := suite.srv.Create(modeluser.Request{})
		if assert.Error(t, err) {
			_, ok := err.(validation.Errors)
			assert.True(t, ok)
		}
	})
}

func (suite *ServiceUserTestSuite) Test_B_UserLogin() {
	suite.T().Run("User login success", func(t *testing.T) {
		repoReturn := entity.User{}
		copier.Copy(&repoReturn, &suite.defaultPayload)
		repoReturn.ID = 1
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(repoReturn.Password), bcrypt.DefaultCost)
		assert.NoError(t, err)
		repoReturn.Password = string(hashedPassword)

		payloadLogin := modeluser.RequestLogin{}
		copier.Copy(&payloadLogin, &suite.defaultPayload)

		suite.repo.On("Login", mock.Anything).Return(repoReturn, nil).Once()

		resp, err := suite.srv.Login(payloadLogin)

		assert.NoError(t, err)
		assert.NotEmpty(t, resp)
	})

	suite.T().Run("user login password not match", func(t *testing.T) {
		repoReturn := entity.User{}
		copier.Copy(&repoReturn, &suite.defaultPayload)
		repoReturn.ID = 1
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(repoReturn.Password), bcrypt.DefaultCost)
		assert.NoError(t, err)
		repoReturn.Password = string(hashedPassword)

		payloadLogin := modeluser.RequestLogin{}
		copier.Copy(&payloadLogin, &suite.defaultPayload)

		suite.repo.On("Login", mock.Anything).Return(repoReturn, nil).Once()

		payloadLogin.Password = "asdadjoqw0ue092e"

		_, err = suite.srv.Login(payloadLogin)

		assert.Error(t, err)
	})

	suite.T().Run("user login invalid email or password", func(t *testing.T) {
		repoReturn := entity.User{}
		copier.Copy(&repoReturn, &suite.defaultPayload)
		repoReturn.ID = 1
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(repoReturn.Password), bcrypt.DefaultCost)
		assert.NoError(t, err)
		repoReturn.Password = string(hashedPassword)

		payloadLogin := modeluser.RequestLogin{}

		suite.repo.On("Login", mock.Anything).Return(repoReturn, nil).Once()

		_, err = suite.srv.Login(payloadLogin)

		assert.Error(t, err)
		_, ok := err.(validation.Errors)
		assert.True(t, ok)
	})
}

func (suite *ServiceUserTestSuite) Test_C_Update() {
	suite.T().Run("Update user succes", func(t *testing.T) {
		updatePayload := modeluser.Request{}
		copier.Copy(&updatePayload, &suite.defaultPayload)
		updatePayload.Username = "update"
		updatePayload.Email = "update@example.com"

		repoReturn := entity.User{}
		copier.Copy(&repoReturn, &updatePayload)
		repoReturn.ID = 1

		suite.repo.On("Update", mock.Anything).Return(repoReturn, nil).Once()

		updatedUser, err := suite.srv.Update(updatePayload)

		assert.NoError(t, err)
		assert.Equal(t, updatedUser.ID, repoReturn.ID)
		assert.Equal(t, updatedUser.Username, repoReturn.Username)
		assert.Equal(t, updatedUser.Email, repoReturn.Email)
	})

	suite.T().Run("update failed, user did not entered id ", func(t *testing.T) {
		updatePayload := modeluser.Request{}
		copier.Copy(&updatePayload, &suite.defaultPayload)
		updatePayload.ID = 0
		updatePayload.Username = "update"
		updatePayload.Email = "update@example.com"

		suite.repo.On("Update", mock.Anything).Return(entity.User{}, gorm.ErrMissingWhereClause).Once()

		_, err := suite.srv.Update(updatePayload)

		assert.Error(t, err)
		assert.True(t, errors.Is(err, gorm.ErrMissingWhereClause))
	})

	suite.T().Run("Update user failed, email & username missing", func(t *testing.T) {
		updatePayload := modeluser.Request{}
		copier.Copy(&updatePayload, &suite.defaultPayload)
		updatePayload.Username = ""
		updatePayload.Email = ""

		repoReturn := entity.User{}
		copier.Copy(&repoReturn, &updatePayload)
		repoReturn.ID = 1

		suite.repo.On("Update", mock.Anything).Return(repoReturn, nil).Once()

		_, err := suite.srv.Update(updatePayload)

		assert.Error(t, err)
	})

}

func (suite *ServiceUserTestSuite) Test_D_Delete() {
	suite.T().Run("Delete user success", func(t *testing.T) {
		suite.repo.On("DeleteByID", mock.Anything).Return(nil).Once()

		err := suite.srv.DeleteByID(suite.defaultPayload.ID)

		assert.NoError(t, err)
	})

	suite.T().Run("Delete user failed, user not found", func(t *testing.T) {
		suite.repo.On("DeleteByID", mock.Anything).Return(gorm.ErrRecordNotFound).Once()

		err := suite.srv.DeleteByID(0)

		assert.Error(t, err)
		assert.True(t, errors.Is(err, gorm.ErrRecordNotFound))
	})
}
