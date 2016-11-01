package main

import (
	"fmt"
	"github.com/mohoromitch/cps633-lab2/table"
	"github.com/mohoromitch/cps633-lab2/users"
)

const USER_PERMISSIONS_FILE = "./.user_and_permissions.db"

func main() {
	fmt.Println("Populating tables...")
	user1 := table.NewTableFromFile("User1.txt")
	user2 := table.NewTableFromFile("User2.txt")
	user3 := table.NewTableFromFile("User3.txt")
	user4 := table.NewTableFromFile("User4.txt")
	user5 := table.NewTableFromFile("User5.txt")
	fmt.Print("User 1 deviations ")
	fmt.Println(user1.RetrieveDeviations())
	fmt.Print("User 1 threshhold: ")
	fmt.Println(user1.CalculateEERThrValue())
	fmt.Print("User 2 deviations ")
	fmt.Println(user2.RetrieveDeviations())
	fmt.Print("User 2 threshhold: ")
	fmt.Println(user2.CalculateEERThrValue())
	fmt.Print("User 3 deviations ")
	fmt.Println(user3.RetrieveDeviations())
	fmt.Print("User 3 threshhold: ")
	fmt.Println(user3.CalculateEERThrValue())
	fmt.Print("User 4 deviations ")
	fmt.Println(user4.RetrieveDeviations())
	fmt.Print("User 4 threshhold: ")
	fmt.Println(user4.CalculateEERThrValue())
	fmt.Print("User 5 deviations ")
	fmt.Println(user5.RetrieveDeviations())
	fmt.Print("User 5 threshhold: ")
	fmt.Println(user5.CalculateEERThrValue())
	fmt.Println("...done populating tables")
	fmt.Println("Reading in file database...")
	lm := users.NewLoginManager(USER_PERMISSIONS_FILE)
	fmt.Println("...done reading in file database")
	lm.Login()
}
