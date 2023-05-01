package user

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var httpCode = 200
var getAll = "some-get-all-data"
var getByID = "some-get-by-id-data"
var create = "some-create-data"
var update = "some-update-data"
var delete = "some-delete-data"

type ctrlMocked struct{}

func MockController() Controller {
	return &ctrlMocked{}
}

func (*ctrlMocked) GetAll(c *gin.Context)  { c.JSON(httpCode, getAll) }
func (*ctrlMocked) GetByID(c *gin.Context) { c.JSON(httpCode, getByID) }
func (*ctrlMocked) Create(c *gin.Context)  { c.JSON(httpCode, create) }
func (*ctrlMocked) Update(c *gin.Context)  { c.JSON(httpCode, update) }
func (*ctrlMocked) Delete(c *gin.Context)  { c.JSON(httpCode, delete) }

func GetRouterMock() *gin.Engine {
	router := gin.Default()
	mockController := MockController()
	NewRouterWithDeps(RouterDependencies{
		controller: mockController,
	}).SetupRouter(router)

	return router
}

func TestGetAll(t *testing.T) {
	router := GetRouterMock()
	expected, _ := json.Marshal(getAll)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/user", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, httpCode, w.Code)
	assert.Equal(t, string(expected), w.Body.String())
}

func TestGetByID(t *testing.T) {
	router := GetRouterMock()
	expected, _ := json.Marshal(getByID)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/user/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, httpCode, w.Code)
	assert.Equal(t, string(expected), w.Body.String())
}

func TestCreate(t *testing.T) {
	router := GetRouterMock()
	expected, _ := json.Marshal(create)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/user", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, httpCode, w.Code)
	assert.Equal(t, string(expected), w.Body.String())
}

func TestUpdate(t *testing.T) {
	router := GetRouterMock()
	expected, _ := json.Marshal(update)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/user/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, httpCode, w.Code)
	assert.Equal(t, string(expected), w.Body.String())
}

func TestDelete(t *testing.T) {
	router := GetRouterMock()
	expected, _ := json.Marshal(delete)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/user/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, httpCode, w.Code)
	assert.Equal(t, string(expected), w.Body.String())
}
