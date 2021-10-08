package v1controller_test

import (
	"bytes"
	authtoken "iamargus95/eKYC-service-gin/jwt"
	v1controller "iamargus95/eKYC-service-gin/v1/controllers"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var unauthRequestTests = []struct {
	url          string
	bodyData     []byte
	expectedCode int
}{
	{
		url:          "/api/v1/signup",
		bodyData:     []byte(`{"name": "testClient","email": "testing@test.in","plan": "basic"}`),
		expectedCode: http.StatusOK,
	},
}

//tests for the endpoint /api/v1/signup
func TestSignup(t *testing.T) {

	asserts := assert.New(t)
	gin.SetMode(gin.TestMode)
	r := gin.New()

	for _, testdata := range unauthRequestTests {

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		requestData := testdata.bodyData

		c.Request, _ = http.NewRequest(http.MethodPost, testdata.url, bytes.NewBuffer(requestData))
		c.Request.Header.Set("Content-Type", "application/json")

		v1controller.Signup(c)

		r.ServeHTTP(w, c.Request)

		asserts.Equal(testdata.expectedCode, w.Code)
	}
}

func newfileUploadRequest(uri string, params map[string]string, paramName, path string) *http.Request {

	file, err := os.Open(path)
	if err != nil {
		return nil
	}

	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		return nil
	}

	fi, err := file.Stat()
	if err != nil {
		return nil
	}

	file.Close()
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile(paramName, fi.Name())
	if err != nil {
		return nil
	}

	part.Write(fileContents)
	for key, val := range params {
		_ = writer.WriteField(key, val)
	}

	err = writer.Close()
	if err != nil {
		return nil
	}

	request, _ := http.NewRequest("POST", uri, body)
	request.Header.Set("Content-Type", writer.FormDataContentType())
	return request
}

//Tests for the endpoint /api/v1/image
func TestImageUpload(t *testing.T) {

	token := authtoken.JWTService().GenerateToken("testClient")
	filepath, _ := os.Getwd()
	filepath += "/test.jpeg"
	asserts := assert.New(t)
	gin.SetMode(gin.TestMode)
	r := gin.New()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	extraparams := map[string]string{
		"type": "face",
	}
	c.Request = newfileUploadRequest("/api/v1/image", extraparams, "file", filepath) //invalid token gen token on the fly
	c.Request.Header.Set("Authorization", "Bearer "+token)

	v1controller.Image(c)

	r.ServeHTTP(w, c.Request)
	asserts.Equal(200, w.Code)

}
