package users

import "strings"

type Permissions struct {
	r, w, x bool
}

func NewPermissions(r, w, x bool) Permissions {
	return Permissions{r, w, x}
}

func ParsePermissions(serialized string) Permissions {
	p := NewPermissions(false, false, false)
	p.Parse(serialized)
	return p
}

func (s *Permissions) Parse(str string) {
	s.r = strings.Contains(str, "r")
	s.w = strings.Contains(str, "w")
	s.x = strings.Contains(str, "x")
}

func (s *Permissions) SetRead(enabled bool) {
	s.r = enabled
}
func (s *Permissions) SetWrite(enabled bool) {
	s.w = enabled
}
func (s *Permissions) SetExecute(enabled bool) {
	s.x = enabled
}

func (s *Permissions) GetRead() bool {
	return s.r
}
func (s *Permissions) GetWrite() bool {
	return s.w
}
func (s *Permissions) GetExecute() bool {
	return s.x
}

func (s *Permissions) Str() string {
	str := ""
	if s.r {
		str += "r"
	} else {
		str += " "
	}
	if s.w {
		str += "w"
	} else {
		str += " "
	}
	if s.x {
		str += "x"
	} else {
		str += " "
	}
	return str
}
