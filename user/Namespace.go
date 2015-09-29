package user

import "strings"

type Namespace string

func (n Namespace) Owns(no Namespace) bool {
	return strings.HasPrefix(string(no), string(n))
}

func (n Namespace) Equal(n2 Namespace) bool {
	return string(n) == string(n2)
}
