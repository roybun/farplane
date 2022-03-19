package main

import (
	"farplane/database"
	"farplane/farplanewebapi"
	"fmt"
)

func main() {
	fmt.Println("HALLOWWWWW!!!!!! ^.^")
	farplanewebapi.StartServer(database.FarplaneDB{})
}
