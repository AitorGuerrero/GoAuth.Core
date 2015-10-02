package user

import (
	"code.google.com/p/go-uuid/uuid"
	"errors"
)

type Id uuid.UUID

func (u User) Id() Id {
	return u.id
}

func (i Id) Equal (i2 Id) bool {
	return string(i) == string(i2)
}

func (i Id) IsEmpty () bool {
	return "" == string(i)
}

func ParseId (s string) (i Id, err error) {
	i = Id(uuid.Parse(s));
	if nil == i {
		err = errors.New("Invalid Id")
	}

	return
}

func (i Id) Serialize () string {
	return uuid.UUID(i).String()
}
