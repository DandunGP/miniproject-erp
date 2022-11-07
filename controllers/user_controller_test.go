package controllers

import (
	"bytes"
	"encoding/json"
	"erp/config"
	"erp/models"
	"errors"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/suite"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type userSuite struct {
	suite.Suite
	mock sqlmock.Sqlmock
}

func TestSuiteUsers(t *testing.T) {
	suite.Run(t, new(userSuite))
}

func (s *userSuite) SetupSuite() {
	db, mock, err := sqlmock.New()
	s.NoError(err)

	var gormDB *gorm.DB
	gormDB, err = gorm.Open(mysql.New(mysql.Config{
		SkipInitializeWithVersion: true,
		Conn:                      db,
	}), &gorm.Config{})
	s.NoError(err)

	config.DB = gormDB
	s.mock = mock
}

func (s *userSuite) TearDownSuite() {
	config.DB = nil
	s.mock = nil
}

// Valid
func (s *userSuite) TestUserController_GetUser() {

	test := []struct {
		name       string
		path       string
		expectCode int
		Body       models.User
	}{
		{
			name:       "get user by id",
			path:       "users/:id",
			expectCode: http.StatusOK,
			Body: models.User{
				Username: "Admin",
				Password: "Admin123",
				Status:   "Admin",
			},
		},
	}
	s.T().Run(test[0].name, func(t *testing.T) {

		expectedResult := sqlmock.NewRows([]string{"username", "password", "status"}).
			AddRow("Admin", "Admin123", "Admin")
		s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE id = ? AND `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT 1")).
			WithArgs(1).
			WillReturnRows(expectedResult)

		res, _ := json.Marshal(test[0].Body)

		r := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(res))
		w := httptest.NewRecorder()
		e := echo.New()
		ctx := e.NewContext(r, w)
		ctx.SetPath(test[0].path)
		ctx.SetParamNames("id")
		ctx.SetParamValues("1")

		if s.NoError(GetUserController(ctx)) {
			body := w.Body.Bytes()

			type Response struct {
				User models.User `json:"users"`
			}

			var user Response
			if err := json.Unmarshal(body, &user); err != nil {
				s.Error(err, "error unmarshalling")
			}

			s.Equal(test[0].expectCode, w.Result().StatusCode)
			s.Equal(test[0].Body, user.User)
		}
	})
}

func (s *userSuite) TestUserController_GetAllUser() {

	test := []struct {
		name       string
		path       string
		expectCode int
		Body       []models.User
	}{
		{
			name:       "get all user",
			path:       "users/",
			expectCode: http.StatusOK,
			Body: []models.User{{
				Username: "Admin",
				Password: "Admin123",
				Status:   "Admin",
			}, {
				Username: "Admin2",
				Password: "Admin1233",
				Status:   "Admin",
			},
			},
		},
	}
	for _, v := range test {
		s.T().Run(v.name, func(t *testing.T) {

			expectedResult := sqlmock.NewRows([]string{"username", "password", "status"}).
				AddRow("Admin", "Admin123", "Admin").
				AddRow("Admin2", "Admin1233", "Admin")
			s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL")).
				WillReturnRows(expectedResult)

			res, _ := json.Marshal(v.Body)

			r := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(res))
			w := httptest.NewRecorder()
			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath(v.path)

			if s.NoError(GetUsersController(ctx)) {
				body := w.Body.Bytes()

				type Response struct {
					User []models.User `json:"users"`
				}

				var user Response
				if err := json.Unmarshal(body, &user); err != nil {
					s.Error(err, "error unmarshalling")
				}

				s.Equal(v.expectCode, w.Result().StatusCode)
				s.Equal(v.Body, user.User)
			}
		})
	}
}

func (s *userSuite) TestUserController_CreateUser() {

	test := []struct {
		name       string
		path       string
		message    string
		expectCode int
		Body       models.User
	}{
		{
			name:       "create user",
			path:       "users/",
			message:    "success create new user",
			expectCode: http.StatusOK,
			Body: models.User{
				Username: "Admin",
				Password: "Admin123",
				Status:   "Admin",
			},
		},
	}
	s.T().Run(test[0].name, func(t *testing.T) {

		s.mock.ExpectBegin()
		s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `users` (`created_at`,`updated_at`,`deleted_at`,`username`,`password`,`status`) VALUES (?,?,?,?,?,?)")).
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), nil, "", "", "").
			WillReturnResult(sqlmock.NewResult(1, 1))
		s.mock.ExpectCommit()

		res, _ := json.Marshal(test[0].Body)

		r := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(res))
		w := httptest.NewRecorder()
		e := echo.New()
		ctx := e.NewContext(r, w)
		ctx.SetPath(test[0].path)
		ctx.Request().Header.Set("Content-Type", "application/json")

		if s.NoError(CreateUserController(ctx)) {
			body := w.Body.Bytes()

			type Response struct {
				Message string      `json:"message"`
				User    models.User `json:"user"`
			}

			var user Response
			if err := json.Unmarshal(body, &user); err != nil {
				s.Error(err, "error unmarshalling")
			}

			s.Equal(test[0].expectCode, w.Result().StatusCode)
			s.Equal(test[0].message, user.Message)
		}
	})
}

func (s *userSuite) TestUserController_UpdateUser() {

	test := []struct {
		name       string
		path       string
		message    string
		expectCode int
		Body       models.User
	}{
		{
			name:       "update user",
			path:       "users/:id",
			message:    "update success",
			expectCode: http.StatusOK,
			Body:       models.User{},
		},
	}
	s.T().Run(test[0].name, func(t *testing.T) {

		s.mock.ExpectBegin()
		s.mock.ExpectExec(regexp.QuoteMeta("UPDATE `users` SET `updated_at`=? WHERE id = ? AND `users`.`deleted_at` IS NULL")).
			WithArgs(sqlmock.AnyArg(), 1).
			WillReturnResult(sqlmock.NewResult(1, 1))
		s.mock.ExpectCommit()

		res, _ := json.Marshal(test[0].Body)

		r := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(res))
		w := httptest.NewRecorder()
		e := echo.New()
		ctx := e.NewContext(r, w)
		ctx.SetPath(test[0].path)
		ctx.SetParamNames("id")
		ctx.SetParamValues("1")

		if s.NoError(UpdateUserController(ctx)) {
			body := w.Body.Bytes()

			type Response struct {
				Message string `json:"message"`
			}

			var user Response
			if err := json.Unmarshal(body, &user); err != nil {
				s.Error(err, "error unmarshalling")
			}

			s.Equal(test[0].expectCode, w.Result().StatusCode)
			s.Equal(test[0].message, user.Message)
		}
	})
}

func (s *userSuite) TestUserController_DeleteUser() {

	test := []struct {
		name       string
		path       string
		message    string
		expectCode int
	}{
		{
			name:       "delete user",
			path:       "users/:id",
			message:    "success delete user",
			expectCode: http.StatusOK,
		},
	}
	s.T().Run(test[0].name, func(t *testing.T) {

		s.mock.ExpectBegin()
		s.mock.ExpectExec(regexp.QuoteMeta("UPDATE `users` SET `deleted_at`=? WHERE `users`.`id` = ? AND `users`.`deleted_at` IS NULL")).
			WithArgs(sqlmock.AnyArg(), 1).
			WillReturnResult(sqlmock.NewResult(1, 1))
		s.mock.ExpectCommit()

		r := httptest.NewRequest(http.MethodDelete, "/", nil)
		w := httptest.NewRecorder()
		e := echo.New()
		ctx := e.NewContext(r, w)
		ctx.SetParamNames("id")
		ctx.SetParamValues("1")
		ctx.SetPath(test[0].path)

		if s.NoError(DeleteUserController(ctx)) {
			body := w.Body.Bytes()

			type Response struct {
				Message string `json:"message"`
			}

			var user Response
			if err := json.Unmarshal(body, &user); err != nil {
				s.Error(err, "error unmarshalling")
			}

			s.Equal(test[0].expectCode, w.Result().StatusCode)
			s.Equal(test[0].message, user.Message)
		}
	})
}

// Invalid
func (s *userSuite) TestUserController_InvalidGetUser() {

	test := []struct {
		name       string
		path       string
		expectCode string
	}{
		{
			name:       "get user by id error",
			path:       "users/:id",
			expectCode: echo.NewHTTPError(http.StatusBadRequest, "Record not found!").Error(),
		},
	}
	s.T().Run(test[0].name, func(t *testing.T) {

		expectedResult := sqlmock.NewRows([]string{"username", "password", "status"}).
			AddRow("Admin", "Admin123", "Admin")
		s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE id = ? AND `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT 1")).
			WithArgs(1).
			WillReturnRows(expectedResult).
			WillReturnError(errors.New("Record not found!"))

		r := httptest.NewRequest(http.MethodGet, "/", nil)
		w := httptest.NewRecorder()
		e := echo.New()
		ctx := e.NewContext(r, w)
		ctx.SetPath(test[0].path)
		ctx.SetParamNames("id")
		ctx.SetParamValues("1")

		err := GetUserController(ctx).Error()

		s.Equal(test[0].expectCode, err)
	})
}

func (s *userSuite) TestUserController_InvalidGetAllUser() {

	test := []struct {
		name       string
		path       string
		expectCode string
	}{
		{
			name:       "get all user error",
			path:       "users/",
			expectCode: echo.NewHTTPError(http.StatusBadRequest, "Record not found!").Error(),
		},
	}
	s.T().Run(test[0].name, func(t *testing.T) {

		expectedResult := sqlmock.NewRows([]string{"username", "password", "email"}).
			AddRow("Admin", "Admin123", "Admin").
			AddRow("Admin2", "Admin1233", "Admin")
		s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL")).
			WillReturnRows(expectedResult).
			WillReturnError(errors.New("Record not found!"))

		r := httptest.NewRequest(http.MethodGet, "/", nil)
		w := httptest.NewRecorder()
		e := echo.New()
		ctx := e.NewContext(r, w)
		ctx.SetPath(test[0].path)

		err := GetUsersController(ctx).Error()

		s.Equal(test[0].expectCode, err)
	})
}

func (s *userSuite) TestUserController_InvalidCreateUser() {

	var test = []struct {
		name           string
		path           string
		expectedResult string
	}{
		{
			name:           "create user error",
			path:           "/users",
			expectedResult: echo.NewHTTPError(http.StatusBadRequest, "Record not found!").Error(),
		},
	}

	s.T().Run(test[0].name, func(t *testing.T) {

		s.mock.ExpectBegin()
		s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `users` (`created_at`,`updated_at`,`deleted_at`,`username`,`password`,`status`) VALUES (?,?,?,?,?,?)")).
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), nil, "", "", "").
			WillReturnError(errors.New("Record not found!"))
		s.mock.ExpectRollback()

		r := httptest.NewRequest(http.MethodPost, "/", nil)
		w := httptest.NewRecorder()
		e := echo.New()
		ctx := e.NewContext(r, w)
		ctx.SetPath(test[0].path)
		ctx.Request().Header.Set("Content-Type", "application/json")

		err := CreateUserController(ctx).Error()
		s.Equal(test[0].expectedResult, err)
	})
}

func (s *userSuite) TestUserController_InvalidUpdateUser() {

	var test = []struct {
		name           string
		path           string
		expectedResult string
	}{
		{
			name:           "update user error",
			path:           "/users/:id",
			expectedResult: echo.NewHTTPError(http.StatusBadRequest, "Record not found!").Error(),
		},
	}

	s.T().Run(test[0].name, func(t *testing.T) {

		s.mock.ExpectBegin()
		s.mock.ExpectExec(regexp.QuoteMeta("UPDATE `users` SET `updated_at`=? WHERE id = ? AND `users`.`deleted_at` IS NULL")).
			WithArgs(sqlmock.AnyArg(), 1).
			WillReturnError(errors.New("Record not found!"))
		s.mock.ExpectRollback()

		r := httptest.NewRequest(http.MethodPut, "/", nil)
		w := httptest.NewRecorder()
		e := echo.New()
		ctx := e.NewContext(r, w)
		ctx.SetPath(test[0].path)
		ctx.SetParamNames("id")
		ctx.SetParamValues("1")

		err := UpdateUserController(ctx).Error()
		s.Equal(test[0].expectedResult, err)
	})
}

func (s *userSuite) TestUserController_InvalidDeleteUser() {

	var test = []struct {
		name           string
		path           string
		expectedResult string
	}{
		{
			name:           "update user error",
			path:           "/users/:id",
			expectedResult: echo.NewHTTPError(http.StatusBadRequest, "Record not found!").Error(),
		},
	}

	s.T().Run(test[0].name, func(t *testing.T) {

		s.mock.ExpectBegin()
		s.mock.ExpectExec(regexp.QuoteMeta("UPDATE `users` SET `deleted_at`=? WHERE `users`.`id` = ? AND `users`.`deleted_at` IS NULL")).
			WithArgs(sqlmock.AnyArg(), 1).
			WillReturnError(errors.New("Record not found!"))
		s.mock.ExpectRollback()

		r := httptest.NewRequest(http.MethodDelete, "/", nil)
		w := httptest.NewRecorder()
		e := echo.New()
		ctx := e.NewContext(r, w)
		ctx.SetPath(test[0].path)
		ctx.SetParamNames("id")
		ctx.SetParamValues("1")

		err := DeleteUserController(ctx).Error()
		s.Equal(test[0].expectedResult, err)
	})
}
