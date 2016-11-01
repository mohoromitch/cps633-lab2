package main

import (
	"fmt"
	"github.com/mohoromitch/cps633-lab2/users"
	"strings"
)

const (
	USER_PERMISSIONS_FILE = "./.user_and_permissions.db"
)

func main() {
	fmt.Println("Hello World")
	loop()
}

func loop() {
	db, err := users.NewDatabase(USER_PERMISSIONS_FILE)
	CheckErr(err)

	for {
		var username string
		fmt.Print("Enter user name (exit to quit): ")
		fmt.Scanf("%s\n", &username)

		if equals(username, "exit") {
			break
		}

		user, exists := db.FindUser(username)
		if exists {
			fmt.Println(user.Str())
		} else {
			fmt.Printf("No user \"%s\" exists on record...\n", username)
			continue
		}
	}
}

func equals(one, two string) bool {
	return strings.Compare(one, two) == 0
}

func CheckErr(err error) {
	if err == nil {
		panic(err)
	}
}
