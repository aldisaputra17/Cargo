package voyage

import (
	"errors"
	"time"

	"github.com/aldisaputra17/cargo/location"
)

type Number string

type Voyage struct {
	Number   Number
	Schedule Schedule
}

func New(n Number, s Schedule) *Voyage {
	return &Voyage{Number: n, Schedule: s}
}

type Schedule struct {
	CarrierMovements []CarrierMovement
}

type CarrierMovement struct {
	DepartureLocation location.UNLcode
	ArrivalLocation   location.UNLcode
	DepartureTime     time.Time
	ArrivalTime       time.Time
}

var ErrUnknown = errors.New("unknown voyage")

type Repository interface {
	Find(Number) (*Voyage, error)
}
