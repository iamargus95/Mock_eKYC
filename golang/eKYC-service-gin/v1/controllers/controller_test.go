package v1controller_test

//Does not work
import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var unauthRequestTests = []struct {
	url          string
	bodyData     []byte
	expectedCode int
	responseData []byte
}{
	{
		url:          "/api/v1/signup",
		bodyData:     []byte(`{"name": "wangzitian0","email": "testing@one2n.in","plan": "basic"}`),
		expectedCode: http.StatusOK,
		responseData: []byte(`{"accessKey": "10-char-JWT-Token","secretKey": "20-char-JWT-Token",}`),
	},
}

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

		r.ServeHTTP(w, c.Request)

		//c.Writer.Status() is the workaround. Used instead of w.Code
		asserts.Equal(testdata.expectedCode, c.Writer.Status())
	}
}
