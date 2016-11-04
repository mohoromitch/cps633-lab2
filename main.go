package main

import (
	"fmt"
	"github.com/mohoromitch/cps633-lab2/table"
	"github.com/mohoromitch/cps633-lab2/users"
)

const USER_PERMISSIONS_FILE = "./.user_and_permissions.db"

func main() {
	fmt.Println("Populating tables...")
	userTables := make([]*table.Table, 5)
	userTables[0] = table.NewTableFromFile("User1.txt")
	userTables[1] = table.NewTableFromFile("User2.txt")
	userTables[2] = table.NewTableFromFile("User3.txt")
	userTables[3] = table.NewTableFromFile("User4.txt")
	userTables[4] = table.NewTableFromFile("User5.txt")
	fmt.Println("...done populating tables")
	for index, u := range userTables {
		fmt.Printf("User %d deviations: \n", index+1)
		fmt.Println(u.RetrieveDeviations())
		fmt.Printf("User %d threshold: \n", index+1)
		fmt.Println(u.CalculateEERThrValue())
	}
	fmt.Println("Reading in file database...")
	lm := users.NewLoginManager(USER_PERMISSIONS_FILE)
	fmt.Println("...done reading in file database")
	lm.Login()
}
