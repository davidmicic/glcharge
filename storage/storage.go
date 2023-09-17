package storage

import (
	"database/sql"
	"fmt"
	"glcharge/models"
	"log"
)

type IDal interface {
	InitDB(conn_str string)
	GetGroups() ([]models.Group, error)
	GetChargePointStatus() ([]models.ChargePointStatus, error)
	GetChargePointConnector() ([]models.ChargePointConnector, error)
	ChangeGroupMaxCurrentById(id int, maxCurrent float64)
	ChangeConnectorStatusById(id int, status string)
	ChangeChargePointPriorityById(id int, priority int)
}

type DalDB struct {
	Db *sql.DB
}

// ChangeChargePointPriorityById implements IDal.
func (d *DalDB) ChangeChargePointPriorityById(id int, priority int) {
	fmt.Print("Called ChangeChargePointPriorityById")
	stmt, err := d.Db.Prepare(`UPDATE public.chargepointstatus SET "Priority" = $1 WHERE ChargePointId = $2`)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	_, err = stmt.Exec(priority, id)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}

// ChangeConnectorStatusById implements IDal.
func (d *DalDB) ChangeConnectorStatusById(id int, status string) {
	fmt.Print("Called ChangeConnectorStatusById")
	stmt, err := d.Db.Prepare(`UPDATE public.chargePointConnector SET "Status" = $1 WHERE chargepointconnectorid = $2`)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	_, err = stmt.Exec(status, id)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}

// ChangeGroupMaxCurrentById implements IDal.
func (d *DalDB) ChangeGroupMaxCurrentById(id int, maxCurrent float64) {
	fmt.Print("Called ChangeGroupMaxCurrentById")
	stmt, err := d.Db.Prepare(`UPDATE public.group SET MaxCurrent = $1 WHERE Id = $2`)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	_, err = stmt.Exec(maxCurrent, id)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}

func (d *DalDB) InitDB(conn_str string) {
	db, err := sql.Open("postgres", conn_str)

	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	d.Db = db
}

func (d *DalDB) GetGroups() ([]models.Group, error) {
	fmt.Print("Called GetGroups")
	rows, err := d.Db.Query("select * from public.group")
	var groups []models.Group

	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var obj models.Group
		if err := rows.Scan(&obj.Id, &obj.MaxCurrent); err != nil {
			log.Fatal(err)
		}

		groups = append(groups, obj)
	}

	return groups, nil
}

func (d *DalDB) GetChargePointStatus() ([]models.ChargePointStatus, error) {
	fmt.Print("Called GetChargePointStatus")
	rows, err := d.Db.Query("select * from public.chargepointstatus")
	var chargePoints []models.ChargePointStatus

	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var obj models.ChargePointStatus
		if err := rows.Scan(&obj.ChargePointId, &obj.Priority, &obj.GroupId); err != nil {
			log.Fatal(err)
		}

		chargePoints = append(chargePoints, obj)
	}

	return chargePoints, nil
}

func (d *DalDB) GetChargePointConnector() ([]models.ChargePointConnector, error) {
	fmt.Print("Called GetChargePointConnector")
	rows, err := d.Db.Query("select * from public.chargepointconnector")
	var chargePointConnector []models.ChargePointConnector

	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var obj models.ChargePointConnector
		if err := rows.Scan(&obj.Id, &obj.Status, &obj.ChargePointId); err != nil {
			log.Fatal(err)
		}

		chargePointConnector = append(chargePointConnector, obj)
	}

	return chargePointConnector, nil
}
