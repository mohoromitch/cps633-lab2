package main

import (
	"fmt"
	"github.com/mohoromitch/cps633-lab2/table"
)

func main() {
	fmt.Println("Populating tables...")
	user1 := table.NewTableFromFile("User1.txt")
	user2 := table.NewTableFromFile("User2.txt")
	user3 := table.NewTableFromFile("User3.txt")
	user4 := table.NewTableFromFile("User4.txt")
	user5 := table.NewTableFromFile("User5.txt")
	fmt.Println(user1.RetrieveDeviations())
	fmt.Println(user2.RetrieveDeviations())
	fmt.Println(user3.RetrieveDeviations())
	fmt.Println(user4.RetrieveDeviations())
	fmt.Println(user5.RetrieveDeviations())
}
