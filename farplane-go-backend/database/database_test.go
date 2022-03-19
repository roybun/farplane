package database

import (
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestAddBudget(t *testing.T) {
	db, mock, err := sqlmock.New()
	defer db.Close()
	if err != nil {
		t.Fatalf("Dhet on")
	}
	mock.ExpectExec("INSERT INTO farplane.budget").WithArgs("BUDGERINO").WillReturnResult(sqlmock.NewResult(1, 1))

	testStruct := FarplaneDB{
		theDatabase: db}
	testStruct.AddBudget("BUDGERINO")
	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Fatalf("DID NOT MEET EXPECTATIONS BRUV")
	}
	fmt.Println("meh")
}
