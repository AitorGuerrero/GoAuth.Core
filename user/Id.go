package user

import (
	"code.google.com/p/go-uuid/uuid"
)

type Id uuid.UUID

func (i Id) Equal (i2 Id) bool {
	return uuid.Equal(uuid.UUID(i), uuid.UUID(i2))
}

func (i Id) IsEmpty () bool {
	return i.Serialize() == ""
}

func ParseId (s string) (i Id) {
	return Id(uuid.Parse(s))
}

func (i Id) Serialize () string {
	return uuid.UUID(i).String()
}
