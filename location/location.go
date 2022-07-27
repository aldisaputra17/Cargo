package location

import "errors"

type UNLcode string

type Location struct {
	UNLcode UNLcode
	Name    string
}

var ErrUnknown = errors.New("unknown location")

type Repository interface {
	Find(locode UNLcode) (*Location, error)
	FindAll() []*Location
}
