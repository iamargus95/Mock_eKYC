package v1service

import (
	u "iamargus95/eKYC-service-gin/apiHelpers"
	"iamargus95/eKYC-service-gin/models"
	res "iamargus95/eKYC-service-gin/resources/api/v1"
)

type ClientService struct {
	Client models.Clients
}

func (cs *ClientService) ClientList() map[string]interface{} {
	gotClient := cs.Client

	clientData := res.ClientsResponse{
		ID:    gotClient.ID,
		Name:  "XYZ",
		Email: "test@email.com",
	}
	response := u.Message(400, "This is a list of Clients.")
	response["data"] = clientData
	return response
}
