package checkToken

import (
	"errors"
)
type Command struct {}

func (c Command) Execute(req Request) (error) {
	return errors.New("Pene")
}

type Request struct {
	Token string
}
