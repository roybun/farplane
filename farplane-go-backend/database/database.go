package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "rbunting"
	password = "pgsql7331"
	dbname   = "farplane"
)

var callCount int = 0

type IFarplaneDB interface {
	AddBudget(string) bool
}

type FarplaneDB struct {
	theDatabase *sql.DB
}

func (fdb FarplaneDB) AddBudget(budgerName string) bool {
	//psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	//	"password=%s dbname=%s sslmode=disable",
	//	host, port, user, password, dbname)
	//db, err := fdb.theDatabase.Open("postgres", psqlInfo)
	//if err != nil {
	//	panic(err)
	//}
	//defer db.Close()

	err := fdb.theDatabase.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	fmt.Println("LeBudgetAdder")
	sqlStatement := `INSERT INTO farplane.budget (name) VALUES ($1)`
	id := 0
	_, err = fdb.theDatabase.Exec(sqlStatement, budgerName)
	if err != nil {
		panic(err)
	}
	fmt.Println("New record ID is:", id)
	//var (
	//	leName string
	//)
	//rows, err := fdb.theDatabase.Query("SELECT name FROM farplane.budget")
	//for rows.Next() {
	//	err = rows.Scan(&leName)
	//	fmt.Println("\nNext Value: ", leName)
	//}
	return true
}

func AddBudget(budgetName string) bool {
	fmt.Println("Adding budget")
	return true
}

func GetCallCount() int {
	callCount++
	return callCount
}
