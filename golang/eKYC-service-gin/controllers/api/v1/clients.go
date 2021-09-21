package v1

import (
	"encoding/json"
	u "iamargus95/eKYC-service-gin/apiHelpers"
	v1s "iamargus95/eKYC-service-gin/services/api/v1"

	"github.com/gin-gonic/gin"
)

func ClientsList(c *gin.Context) {

	var clientsService v1s.ClientService

	err := json.NewDecoder(c.Request.Body).Decode(&clientsService.Client)
	if err != nil {
		u.Respond(c.Writer, u.Message(1, "Invalid Request."))
		return
	}

	//call service
	resp := clientsService.ClientList()

	//return response using api helper
	u.Respond(c.Writer, resp)
}

func CreateClient(c *gin.Context) { //Fix queries before writing this

	var clientsService v1s.ClientService

	err := json.NewDecoder()
}
