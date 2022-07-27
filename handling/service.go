package handling

import "errors"

var ErrInvalidArgument = errors.New("invalid argument")

type EventHandler interface {
	CargoWasHandled()
}
