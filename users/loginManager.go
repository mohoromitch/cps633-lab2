package users

import (
	"fmt"
)

type loginManager struct {
	db *database
}

func NewLoginManager(filename string) *loginManager {
	db, err := NewDatabase(filename)
	if err != nil {
		panic(err)
	}
	return &loginManager{db: db}
}

func (lm *loginManager) Login() int {
	var username, file, permission string
	username = lm.readUsername()
	file = lm.readFile()
	permission = lm.readPermission()
	fmt.Println("Requesting file " + file + " with permission " + permission + " for user " + username)
	return 0
}

func (lm *loginManager) readUsername() (username string) {
	fmt.Println("Please enter in your username")
	fmt.Scanf("%s\n", &username)
	return username
}

func (lm *loginManager) readFile() (file string) {
	fmt.Println("Please enter in the file you want to access")
	fmt.Scanf("%s\n", &file)
	return file
}

func (lm *loginManager) readPermission() (permission string) {
	fmt.Println("Please enter in the permission you would like to request")
	fmt.Scanf("%s\n", &permission)
	return permission
}
