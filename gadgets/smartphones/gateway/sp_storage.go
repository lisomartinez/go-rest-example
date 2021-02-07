package gateway

import (
	"github.com/lisomartinez/go-rest-example/gadgets/smartphones/models"
	"github.com/lisomartinez/go-rest-example/internal/database"
	"github.com/lisomartinez/go-rest-example/internal/logs"
)

type SmartphoneStorageGateway interface {
	Add(cmd *models.CreateSmartphoneCMD) (*models.Smartphone, error)
}

type SmartphoneStorage struct {
	*database.MySqlClient
}

func (storage *SmartphoneStorage) Add(cmd *models.CreateSmartphoneCMD) (*models.Smartphone, error) {
	tx, err := storage.MySqlClient.Begin()

	if err != nil {
		logs.Log().Error("cannot create transaction")
	}

	res, err := tx.Exec(`insert into smartphone (name, price, country_origin, os) values (?, ?, ?, ?)`,
		cmd.Name, cmd.Price, cmd.CountryOrigin, cmd.OS)

	if err != nil {
		logs.Log().Error("cannot execute statement")
		_ = tx.Rollback()
		return nil, err
	}

	id, err := res.LastInsertId()

	if err != nil {
		logs.Log().Error("cannot fetch last id")
		_ = tx.Rollback()
		return nil, err
	}

	_ = tx.Commit()

	return &models.Smartphone{
		Id:            id,
		Name:          cmd.Name,
		Price:         cmd.Price,
		CountryOrigin: cmd.CountryOrigin,
		OS:            cmd.OS,
	}, nil
}
