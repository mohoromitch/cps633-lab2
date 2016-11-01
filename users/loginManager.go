package users

import (
	"fmt"
)

type loginManager struct {
	db *database
}

func NewLoginManager(filename string) *loginManager {
	db, err := NewDatabase(filename)
	if err == nil {
		panic(err)
	}
	return &loginManager{db: db}
}

func (lm *loginManager) Login() int {
	var username string
	fmt.Print("Please enter in your username")
	fmt.Scanf("%s\n", &username)
	return 0
}
