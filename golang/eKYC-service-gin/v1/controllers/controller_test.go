package v1controller_test

//Does not work
import (
	"bytes"
	"iamargus95/eKYC-service-gin/v1/routes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var unauthRequestTests = []struct {
	url          string
	method       string
	bodyData     string
	expectedCode int
	responseData string
}{
	{
		"/signup",
		"POST",
		`{"user":{"name": "wangzitian0","email": "testing@one2n.in","plan": "basic"}}`,
		http.StatusOK,
		`{"accessKey": "10-char-JWT-Token","secretKey": "20-char-JWT-Token",}`,
	},
}

func TestSignup(t *testing.T) {

	asserts := assert.New(t)
	r := gin.New() //Also write maintest
	routes.SignupClient(r.Group("/api/v1"))

	for _, testdata := range unauthRequestTests {
		requestData := testdata.bodyData
		req, err := http.NewRequest(testdata.method, testdata.url, bytes.NewBufferString(requestData))
		asserts.NoError(err)
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		asserts.Equal(testdata.expectedCode, w.Code)
	}
}
