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
	file = lm.readFile(username)
	permission = lm.readPermission()
	if lm.accessGranted(username, file, permission) {
		fmt.Println("Access granted!")
	} else {
		fmt.Println("Access denied!")
	}
	return 0
}

func (lm *loginManager) readUsername() (username string) {
	fmt.Println("Please enter in a username")
	fmt.Scanf("%s\n", &username)
	if _, exists := lm.db.users[username]; !exists {
		fmt.Println("User does not exist.")
		return lm.readUsername()
	}
	return username
}

func (lm *loginManager) readFile(username string) (file string) {
	fmt.Println("Please enter in the file you want to access")
	fmt.Scanf("%s\n", &file)
	if _, exists := lm.db.users[username].permissions[file]; !exists { //LOL
		fmt.Println("File does not exist.")
		return lm.readFile(username)
	}
	return file
}

func (lm *loginManager) readPermission() (permission string) {
	fmt.Println("Please enter in the permission you would like to request")
	fmt.Scanf("%s\n", &permission)
	if !(permission == "r" || permission == "w" || permission == "x") {
		fmt.Print("Invalid permission")
		return lm.readPermission()
	}
	return permission
}

func (lm *loginManager) accessGranted(username, file, permission string) bool {
	p := lm.db.users[username].permissions[file]
	if permission == "r" {
		return p.r
	} else if permission == "w" {
		return p.w
	} else { //permission == 'x'
		return p.x
	}
}
