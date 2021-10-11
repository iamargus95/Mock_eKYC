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
	bodyData     []byte
	expectedCode int
}{
	{
		// OK
		bodyData:     []byte(`{"name": "testClient","email": "testing@test.in","plan": "basic"}`),
		expectedCode: http.StatusOK,
	},
	{
		// Invalid URL
		bodyData:     []byte(`{"name": "testClient1","email": "testing1@test.in","plan": "basic"}`),
		expectedCode: http.StatusNotFound,
	},
	{
		// Invalid plan
		bodyData:     []byte(`{"name": "testClient2","email": "testing2@test.in","plan": "Basic"}`),
		expectedCode: http.StatusForbidden,
	},
	{
		// Invalid email
		bodyData:     []byte(`{"name": "testClient3","email": "testing3@in","plan": "basic"}`),
		expectedCode: http.StatusForbidden,
	},
	{
		// Email not unique
		bodyData:     []byte(`{"name": "testClient4","email": "testing@test.in","plan": "basic"}`),
		expectedCode: 400,
	},
	{
		// Name not unique
		bodyData:     []byte(`{"name": "testClient","email": "testing5@test.in","plan": "basic"}`),
		expectedCode: 400,
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

		c.Request, _ = http.NewRequest(http.MethodPost, "api/v1/signup", bytes.NewBuffer(requestData))
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

	var testUpload = []struct {
		filepath     string
		imageType    string
		expectedCode int
	}{
		{
			filepath:     "/test.jpeg",
			imageType:    "face",
			expectedCode: 200,
		},
		{
			filepath:     "/test.jpeg",
			imageType:    "id_card",
			expectedCode: 200,
		},
		{
			filepath:     "/test.jpeg",
			imageType:    "idcard",
			expectedCode: 400,
		},
		{
			filepath:     "/test.jpeg",
			imageType:    "undefined",
			expectedCode: 400,
		},
	}

	for _, test := range testUpload {

		token := authtoken.JWTService().GenerateToken("testClient")
		filepath, _ := os.Getwd()
		filepath += test.filepath
		asserts := assert.New(t)
		gin.SetMode(gin.TestMode)
		r := gin.New()

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		extraparams := map[string]string{
			"type": test.imageType,
		}

		c.Request = newfileUploadRequest("/api/v1/image", extraparams, "file", filepath)
		c.Request.Header.Set("Authorization", "Bearer "+token)

		v1controller.Image(c)

		r.ServeHTTP(w, c.Request)
		asserts.Equal(test.expectedCode, w.Code)
	}
}
