package main

import (

	"github.com/sirupsen/logrus"
	"transfer-banking/models"
)

func main() {
	models.DB.TruncateAll()
	william := &models.Account{
		Name: "William",
		CPF: "12513623194",
		Ballance: 10000.00,
	}

	guilherme := &models.Account{
		Name: "Guilherme",
		CPF: "12513623195",
		Ballance: 10000.00,
	}

	err := models.DB.Create(william)
	if err != nil {
		logrus.Fatalf("Error creating user %v with err %v", william, err)
	}

	err = models.DB.Create(guilherme)
	if err != nil {
		logrus.Fatalf("Error creating user %v with err %v", guilherme, err)
	}

}
