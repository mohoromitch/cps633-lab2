package users

import (
	"bufio"
	"os"
)

type database struct {
	users map[string]User
}

func NewDatabase(filename string) (*database, error) {
	db := database{}
	err := db.Load(filename)
	return &db, err
}

func (db *database) Load(filename string) error {
	users := make(map[string]User)

	f, err := os.Open(filename)
	defer f.Close()

	if err != nil {
		return err
	}

	db.users = users

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	db.users = users

	for scanner.Scan() {
		user := ParseUserFromString(scanner.Text())
		db.AddUser(user)
	}

	return nil
}

func (db *database) AddUser(user User) {
	db.users[user.GetName()] = user
}

func (db *database) FindUser(username string) (*User, bool) {
	user, exists := db.users[username]
	if exists {
		return &user, true
	} else {
		return nil, false
	}
}
