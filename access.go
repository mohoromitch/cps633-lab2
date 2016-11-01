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
			for {
				fmt.Printf("Logged in as %s\n", user.GetName())

				fmt.Print("Enter name of file to edit permisions for (exit to quit): ")
				var filename string
				fmt.Scanf("%s", &filename)

				if equals(filename, "exit") {
					break
				}

				perm := user.GetFilePermissions(filename)
				fmt.Printf("Set permissions for file \"%s\" to: ", filename)
				var newPerms string
				fmt.Scanf("%s", &newPerms)

				perm.Parse(newPerms)
				user.SetFilePermissions(filename, perm)

			}
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
	if err != nil {
		panic(err)
	}
}
