package middlewares_test

import (
	authtoken "iamargus95/eKYC-service-gin/jwt"
	"iamargus95/eKYC-service-gin/middlewares"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAuthenticateHeader(t *testing.T) {

	asserts := assert.New(t)
	token := authtoken.JWTService().GenerateToken("testClient")

	gin.SetMode(gin.TestMode)
	r := gin.New()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request, _ = http.NewRequest(http.MethodPost, "/api/v1/image", nil)
	c.Request.Header.Set("Authorization", "Bearer "+token)
	middlewares.EnsureLoggedIn(authtoken.JWTService())

	r.ServeHTTP(w, c.Request)

	asserts.Equal(200, w.Code)
}
