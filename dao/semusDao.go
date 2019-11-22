package dao

import (
	"github.com/joaopandolfi/blackwhale/remotes/mysql"
	"github.com/joaopandolfi/blackwhale/utils"
	"github.com/joaopandolfi/miia_api/models"
)

type SemusDao interface {
	GetMany(companyID int) (semus []models.Semus, err error)
	Save(semus models.Semus) (success bool, err error)
	Get(semusID int, companyID int) (semus models.Semus, sucess bool, err error)
}

type Semus struct {
}

func (cc Semus) Save(semus models.Semus) (success bool, err error) {
	_, err = mysql.Driver.ExecuteAndReturnLastId("INSERT INTO semus (iduser,idcompany,result,input) VALUES (?,?,?,?);", 1, 1, 1, 1)
	success = false

	if err == nil {
		success = true
	}
	return
}

func (cc Semus) GetMany(companyID int) (semus []models.Semus, err error) {
	var semusSql []models.SemusSQL

	err = mysql.Driver.Execute(mysql.THE_CASE, &semusSql, `SELECT * FROM semus WHERE idcompany  = ?`, companyID)
	if err != nil {
		utils.CriticalError("[SemusDao][GetMany] - Error on get query", err.Error())
		return
	}

	for _, semu := range semusSql {
		semus = append(semus, semu.Parse())
	}

	return
}

func (cc Semus) Get(semusID int, companyID int) (semus models.Semus, sucess bool, err error) {
	var semuss []models.Semus
	sucess = false

	err = mysql.Driver.Execute(mysql.THE_CASE, &semuss, `SELECT * FROM semus WHERE idsemus = ? AND idcompany  = ?`, semusID, companyID)
	if err != nil {
		utils.CriticalError("[SemusDao][Get] - Error on get query", err.Error())
		return
	}

	if len(semuss) > 0 {
		semus = semuss[0]
		sucess = true
	}

	return
}
