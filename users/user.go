package users

import (
	"strings"
)

type User struct {
	name        string
	permissions map[string]Permissions
}

func NewUser(name string) User {
	return User{name, make(map[string]Permissions)}
}

func ParseUserFromString(serializedUser string) User {
	fields := strings.Fields(serializedUser)
	perm := make(map[string]Permissions)
	for i := 1; i < len(fields); i += 2 {
		perm[fields[i]] = ParsePermissions(fields[i+1])
	}
	return User{fields[0], perm}
}

func (s *User) SetFilePermissions(filename string, p Permissions) {
	s.permissions[filename] = p
}

func (s *User) GetFilePermissions(filename string) Permissions {
	p, exists := s.permissions[filename]
	if !exists {
		p = NewPermissions(true, true, false)
	}
	return p
}

func (s *User) GetName() string {
	return s.name
}

func (s *User) Str() string {
	str := "Current permissions for " + s.name + ":\n"
	for filename, permission := range s.permissions {
		str += "\t" + filename + ": "  + permission.Str() + "\n"
	}
	return str
}

func (s *User) Serialize() string {
	str := s.name + " ";
	for filename, permission := range(s.permissions) {
		str += filename + " " + permission.Str() + " ";
	}
	return str
}
