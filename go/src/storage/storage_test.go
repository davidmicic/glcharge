package storage

import (
	// "glcharge/go/src/container"

	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestChangeGroupMaxCurrentById(t *testing.T) {
	db, mock, err := sqlmock.New()
	d := DalDB{Db: db}

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.
		ExpectPrepare(`UPDATE public.group SET MaxCurrent`).
		ExpectExec().
		WithArgs(500.0, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// // now we execute our method
	d.ChangeGroupMaxCurrentById(1, 500.0)

	// // we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestChangeConnectorStatusById(t *testing.T) {
	db, mock, err := sqlmock.New()
	d := DalDB{Db: db}

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.
		ExpectPrepare(`PDATE public.chargePointConnector SET "Status"`).
		ExpectExec().
		WithArgs("Available", 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// // now we execute our method
	d.ChangeConnectorStatusById(1, "Available")

	// // we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestChangeChargePointPriorityById(t *testing.T) {
	db, mock, err := sqlmock.New()
	d := DalDB{Db: db}

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.
		ExpectPrepare(`UPDATE public.chargepointstatus SET "Priority"`).
		ExpectExec().
		WithArgs(5, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// // now we execute our method
	d.ChangeChargePointPriorityById(1, 5)

	// // we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
