package users

import "strings"
type Permissions struct {
	r, w, x bool
}

func ParsePermissions(serialized string) Permissions {
	return Permissions{strings.Contains(serialized, "r"), strings.Contains(serialized, "w"), strings.Contains(serialized, "x")}
}

func NewPermissions(r, w, x bool) Permissions {
	return Permissions{r, w, x}
}

func (s *Permissions) SetRead(enabled bool) {
	s.r = enabled;
}
func (s *Permissions) SetWrite(enabled bool) {
	s.w = enabled;
}
func (s *Permissions) SetExecute(enabled bool) {
	s.x = enabled;
}

func (s *Permissions) GetRead() bool {
	return s.r;
}
func (s *Permissions) GetWrite() bool {
	return s.w;
}
func (s *Permissions) GetExecute() bool {
	return s.x;
}

func (s *Permissions) Str() string {
	str := "";
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
