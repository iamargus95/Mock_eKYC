package routes_test

import (
	"iamargus95/eKYC-service-gin/v1/routes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var testWelcome = struct {
	url          string
	expectedCode int
	responseData []byte
}{
	url:          "/",
	expectedCode: http.StatusOK,
	responseData: []byte(`"message": "Welcome to the Draft API."`),
}

func TestWelcome(t *testing.T) {

	gin.SetMode(gin.TestMode)
	asserts := assert.New(t)
	r := gin.New()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest(http.MethodGet, testWelcome.url, nil)

	routes.Welcome(c)

	r.ServeHTTP(w, c.Request)
	asserts.Equal(testWelcome.expectedCode, w.Code)
}

var testNotFound = struct {
	url          string
	expectedCode int
	responseData []byte
}{
	url:          "/noexistent",
	expectedCode: http.StatusNotFound,
	responseData: []byte(`"message": "Route not found."`),
}

func TestNotFound(t *testing.T) {

	gin.SetMode(gin.TestMode)
	asserts := assert.New(t)
	r := gin.New()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest(http.MethodGet, testNotFound.url, nil)

	routes.NotFound(c)

	r.ServeHTTP(w, c.Request)
	asserts.Equal(testNotFound.expectedCode, w.Code)
}
