package cargo

import (
	"errors"

	"github.com/aldisaputra17/cargo/location"
	"github.com/aldisaputra17/cargo/voyage"
)

type HandlingActivity struct {
	Type         HandlingEventType
	Location     location.UNLcode
	VoyageNumber voyage.Number
}

type HandlingEvent struct {
	TrackingID TrackingID
	Activity   HandlingActivity
}

type HandlingEventType int

const (
	NotHandled HandlingEventType = iota
	Load
	Unload
	Receive
	Claim
	Customs
)

func (t HandlingEventType) String() string {
	switch t {
	case NotHandled:
		return "Not Handled"
	case Load:
		return "Load"
	case Unload:
		return "Unload"
	case Receive:
		return "Receive"
	case Claim:
		return "Claim"
	case Customs:
		return "Customs"
	}

	return ""
}

type HandlingHistory struct {
	HandlingEvents []HandlingEvent
}

func (h HandlingHistory) MostRecentlyCompletedEvent() (HandlingEvent, error) {
	if len(h.HandlingEvents) == 0 {
		return HandlingEvent{}, errors.New("delivery history is empty")
	}

	return h.HandlingEvents[len(h.HandlingEvents)-1], nil
}

type HandlingEventRepository interface {
	Store(e HandlingEvent)
	QueryHandlingHistory(TrackingID) HandlingHistory
}

type HandlingEventFactory struct {
	CargoRepository    Repository
	VoyageRepository   voyage.Repository
	LocationRepository location.Repository
}
