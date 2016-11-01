package users

import (
	"os"
	"bufio"
)

var users map[string]User

func Load(filename string) error {
	users = make(map[string]User)

	f, err := os.Open(filename)
	defer f.Close()

	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		user := ParseUserFromString(scanner.Text());
		AddUser(user)
	}

	return nil
}

func AddUser(user User) {
	users[user.GetName()] = user
}

func FindUser(username string) (*User, bool) {
	user, exists := users[username]
	if exists {
		return &user, true
	} else {
		return nil, false
	}
}
