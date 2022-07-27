package cargo

import (
	"time"

	"github.com/aldisaputra17/cargo/location"
	"github.com/aldisaputra17/cargo/voyage"
)

type Leg struct {
	VoyageNumber   voyage.Number    `json:"voyage_number"`
	LoadLocation   location.UNLcode `json:"from"`
	UnloadLocation location.UNLcode `json:"to"`
	LoadTime       time.Time        `json:"load_time"`
	UnloadTime     time.Time        `json:"unload_time"`
}

func NewLeg(voyageNumber voyage.Number, loadLocation, unloadLocation location.UNLcode, loadTime, unloadTime time.Time) Leg {
	return Leg{
		VoyageNumber:   voyageNumber,
		LoadLocation:   loadLocation,
		UnloadLocation: unloadLocation,
		LoadTime:       loadTime,
		UnloadTime:     unloadTime,
	}
}

type Itinerary struct {
	Legs []Leg `json:"legs"`
}

func (i Itinerary) InitialDepartureLocation() location.UNLcode {
	if i.IsEmpty() {
		return location.UNLcode("")
	}
	return i.Legs[0].LoadLocation
}

func (i Itinerary) FinalArrivalLocation() location.UNLcode {
	if i.IsEmpty() {
		return location.UNLcode("")
	}
	return i.Legs[len(i.Legs)-1].UnloadLocation
}

func (i Itinerary) FinalArrivalTime() time.Time {
	return i.Legs[len(i.Legs)-1].UnloadTime
}

// IsEmpty checks if the itinerary contains at least one leg.
func (i Itinerary) IsEmpty() bool {
	return i.Legs == nil || len(i.Legs) == 0
}

func (i Itinerary) IsExpected(event HandlingEvent) bool {
	if i.IsEmpty() {
		return true
	}

	switch event.Activity.Type {
	case Receive:
		return i.InitialDepartureLocation() == event.Activity.Location
	case Load:
		for _, l := range i.Legs {
			if l.LoadLocation == event.Activity.Location && l.VoyageNumber == event.Activity.VoyageNumber {
				return true
			}
		}
		return false
	case Unload:
		for _, l := range i.Legs {
			if l.UnloadLocation == event.Activity.Location && l.VoyageNumber == event.Activity.VoyageNumber {
				return true
			}
		}
		return false
	case Claim:
		return i.FinalArrivalLocation() == event.Activity.Location
	}

	return true
}
