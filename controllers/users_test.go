package controllers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"user-kel-6/config"
	"user-kel-6/models"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var user []models.User

func InitEcho() *echo.Echo {
	// Setup
	config.InitDB()
	e := echo.New()
	return e
}

func TestGetUser(t *testing.T) {
	t.Run("Test case 1, valid user id", func(t *testing.T) {
		e := InitEcho()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath("/users/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")

		// Assertions
		if assert.NoError(t, GetUserController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})
	t.Run("Test case 2, invalid user id", func(t *testing.T) {
		e := InitEcho()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath("/users/:id")
		c.SetParamNames("id")
		c.SetParamValues("999")

		// Assertions
		if assert.NoError(t, GetUserController(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})
}

func TestLoginUser(t *testing.T) {
	t.Run("Test case 1, valid authentication user", func(t *testing.T) {
		e := InitEcho()
		userJSON := `{"email":"azka@alterra.com","password":"1234567"}`
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, LoginUsersController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}

	})
}

func TestCreateUser(t *testing.T) {
	t.Run("Test case 1, valid register", func(t *testing.T) {
		test := "testing"
		config.DB.Where("email = ?", test).Delete(&user)
		e := InitEcho()
		userJSON := `{"name":"testing","username":"testing","email":"testing","phone":"012345","password":"testing","address":"malang"}`
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, CreateUserController(c)) {
			assert.Equal(t, http.StatusCreated, rec.Code)
		}
	})
	t.Run("Test case 2, uncompleted email", func(t *testing.T) {
		test := "testing"
		config.DB.Where("email = ?", test).Delete(&user)
		e := InitEcho()
		userJSON := `{"name":"testing","username":"testing","email":"","phone":"","password":"","address":"malang"}`
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, CreateUserController(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})
	t.Run("Test case 2.1, uncompleted name", func(t *testing.T) {
		test := "testing"
		config.DB.Where("email = ?", test).Delete(&user)
		e := InitEcho()
		userJSON := `{"name":"","username":"testing","email":"","phone":"","password":"","address":"malang"}`
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, CreateUserController(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})

	t.Run("Test case 2.2, uncompleted phone", func(t *testing.T) {
		test := "testing"
		config.DB.Where("email = ?", test).Delete(&user)
		e := InitEcho()
		userJSON := `{"name":"testing","username":"testing","email":"testing","phone":"","password":"","address":"malang"}`
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, CreateUserController(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})

	t.Run("Test case 2.3, uncompleted password", func(t *testing.T) {
		test := "testing"
		config.DB.Where("email = ?", test).Delete(&user)
		e := InitEcho()
		userJSON := `{"name":"testing","username":"testing","email":"testing","phone":"0","password":"","address":"malang"}`
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, CreateUserController(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})

	t.Run("Test case 2.4, uncompleted address", func(t *testing.T) {
		test := "testing"
		config.DB.Where("email = ?", test).Delete(&user)
		e := InitEcho()
		userJSON := `{"name":"testing","username":"testing","email":"testing","phone":"0","password":"asas","address":""}`
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, CreateUserController(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})

	t.Run("Test case 2.5, uncompleted username", func(t *testing.T) {
		test := "testing"
		config.DB.Where("email = ?", test).Delete(&user)
		e := InitEcho()
		userJSON := `{"name":"testing","username":"","email":"testing","phone":"0","password":"asas","address":""}`
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, CreateUserController(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})

	t.Run("Test case 3, registered email", func(t *testing.T) {
		e := InitEcho()
		userJSON := `{"name":"testing","username":"testing","email":"azka@alterra.com","phone":"012121212","password":"testing","address":"malang"}`
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, CreateUserController(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})

	t.Run("Test case 4, registered phone", func(t *testing.T) {
		test := "testing"
		config.DB.Where("email = ?", test).Delete(&user)
		e := InitEcho()
		userJSON := `{"name":"testing","username":"testing","email":"testing","phone":"081334304990","password":"testing","address":"malang"}`
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, CreateUserController(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})

	t.Run("Test case 5, registered username", func(t *testing.T) {
		test := "testing"
		config.DB.Where("email = ?", test).Delete(&user)
		e := InitEcho()
		userJSON := `{"name":"testing","username":"testing","email":"testing","phone":"081334304990","password":"testing","address":"malang"}`
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, CreateUserController(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})
}

func TestUpdateUser(t *testing.T) {
	t.Run("Test case 1, valid register", func(t *testing.T) {
		test := "testing"
		config.DB.Where("email = ?", test).Delete(&user)
		e := InitEcho()
		userJSON := `{"name":"testing","username":"testing","email":"testing","phone":"012345","password":"testing","address":"malang"}`
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, UpdateUserController(c)) {
			assert.Equal(t, http.StatusCreated, rec.Code)
		}
	})
	t.Run("Test case 2, uncompleted email", func(t *testing.T) {
		test := "testing"
		config.DB.Where("email = ?", test).Delete(&user)
		e := InitEcho()
		userJSON := `{"name":"testing","username":"testing","email":"","phone":"","password":"","address":"malang"}`
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, UpdateUserController(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})
	t.Run("Test case 2.1, uncompleted name", func(t *testing.T) {
		test := "testing"
		config.DB.Where("email = ?", test).Delete(&user)
		e := InitEcho()
		userJSON := `{"name":"","username":"testing","email":"","phone":"","password":"","address":"malang"}`
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, UpdateUserController(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})

	t.Run("Test case 2.2, uncompleted phone", func(t *testing.T) {
		test := "testing"
		config.DB.Where("email = ?", test).Delete(&user)
		e := InitEcho()
		userJSON := `{"name":"testing","username":"testing","email":"testing","phone":"","password":"","address":"malang"}`
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, UpdateUserController(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})

	t.Run("Test case 2.3, uncompleted password", func(t *testing.T) {
		test := "testing"
		config.DB.Where("email = ?", test).Delete(&user)
		e := InitEcho()
		userJSON := `{"name":"testing","username":"testing","email":"testing","phone":"0","password":"","address":"malang"}`
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, UpdateUserController(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})

	t.Run("Test case 2.4, uncompleted address", func(t *testing.T) {
		test := "testing"
		config.DB.Where("email = ?", test).Delete(&user)
		e := InitEcho()
		userJSON := `{"name":"testing","username":"testing","email":"testing","phone":"0","password":"asas","address":""}`
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, UpdateUserController(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})

	t.Run("Test case 2.5, uncompleted username", func(t *testing.T) {
		test := "testing"
		config.DB.Where("email = ?", test).Delete(&user)
		e := InitEcho()
		userJSON := `{"name":"testing","username":"","email":"testing","phone":"0","password":"asas","address":""}`
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, UpdateUserController(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})

	t.Run("Test case 3, registered email", func(t *testing.T) {
		e := InitEcho()
		userJSON := `{"name":"testing","username":"testing","email":"azka@alterra.com","phone":"012121212","password":"testing","address":"malang"}`
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, UpdateUserController(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})

	t.Run("Test case 4, registered phone", func(t *testing.T) {
		test := "testing"
		config.DB.Where("email = ?", test).Delete(&user)
		e := InitEcho()
		userJSON := `{"name":"testing","username":"testing","email":"testing","phone":"081334304990","password":"testing","address":"malang"}`
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, UpdateUserController(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})

	t.Run("Test case 5, registered username", func(t *testing.T) {
		test := "testing"
		config.DB.Where("email = ?", test).Delete(&user)
		e := InitEcho()
		userJSON := `{"name":"testing","username":"testing","email":"testing","phone":"081334304990","password":"testing","address":"malang"}`
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, UpdateUserController(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})
}
