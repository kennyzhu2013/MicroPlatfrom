package models

import "strconv"

func (m *Preferences) String() string {
	return strconv.Itoa(m.User)
}
