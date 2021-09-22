package v1

import (
	u "iamargus95/eKYC-service-gin/apiHelpers"
	"iamargus95/eKYC-service-gin/conn"
	m "iamargus95/eKYC-service-gin/models"
	v1s "iamargus95/eKYC-service-gin/services/api/v1"

	"github.com/gin-gonic/gin"
)

func ClientsList(c *gin.Context) {

	var clients v1s.ClientService

	db := conn.GetDB()
	err := db.Debug().Model(&m.Client{}).Limit(100).Find(&clients).Error
	if err != nil {
		u.Respond(c.Writer, u.Message(400, "api call failed."))
		return
	}

	resp := clients.ClientList()
	//return response using api helper
	u.Respond(c.Writer, resp)
}
