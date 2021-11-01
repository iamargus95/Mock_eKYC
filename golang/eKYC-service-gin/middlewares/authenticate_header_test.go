package middlewares

import (
	"fmt"
	authtoken "iamargus95/eKYC-service-gin/jwt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {

	os.Exit(m.Run())
}

const (
	AuthKeyHeader  = "Authorization"
	AuthTypeBearer = "Bearer"
)

func addAuthorization(t *testing.T, request *http.Request, tokenMaker authtoken.JWTInterface,
	authorizationType string, username string) {

	token := tokenMaker.GenerateToken(username)

	authorizationHeader := fmt.Sprintf("%s %s", AuthTypeBearer, token)
	request.Header.Set(AuthKeyHeader, authorizationHeader)
}

func TestAuthMiddleware(t *testing.T) {
	testCases := []struct {
		name          string
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker authtoken.JWTInterface)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker authtoken.JWTInterface) {
				addAuthorization(t, request, tokenMaker, AuthTypeBearer, "test")
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {

			gin.SetMode(gin.TestMode)
			r := gin.New()
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			authPath := "/api/v1/image"
			c.Request, _ = http.NewRequest(http.MethodPost, authPath, nil)
			r.POST(
				authPath,
				EnsureLoggedIn(authtoken.JWTService()),
			)

			tc.setupAuth(t, c.Request, authtoken.JWTService())
			r.ServeHTTP(w, c.Request)
			tc.checkResponse(t, w)
		})
	}
}
