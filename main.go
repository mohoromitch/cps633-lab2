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
	fmt.Println(user1.CalculateEERThrValue())
	fmt.Println(user2.CalculateEERThrValue())
	fmt.Println(user3.CalculateEERThrValue())
	fmt.Println(user4.CalculateEERThrValue())
	fmt.Println(user5.CalculateEERThrValue())
	lm := users.NewLoginManager(USER_PERMISSIONS_FILE)
	lm.Login()
}
