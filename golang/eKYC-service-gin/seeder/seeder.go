package seed

import (
	"iamargus95/eKYC-service-gin/infra"
)

// var clients = []models.Client{
// 	{
// 		Name:  "E Corp",
// 		Email: "ecorp@evil.com",
// 		Plan:  models.Plan{ClientID: 1, Plan: "Enterprise"},
// 	},
// 	{
// 		Name:  "Corp2",
// 		Email: "corp2@corp2.com",
// 		Plan:  models.Plan{ClientID: 2, Plan: "Enterprise"},
// 	},
// }

func Load() {

	m := infra.GetMigrationTool()
	m.Up()
}
