package models

import "strconv"

func (m *Preferences) String() string {
	return strconv.Itoa(m.User)
}

//base operations of this model define here..
//func1 of complex select