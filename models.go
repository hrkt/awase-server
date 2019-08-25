package main

import (
	"github.com/google/uuid"
)

type Event struct {
	ID             string
	Label          string
	CandidateDates []CandidateDate
	Entries        []EventEntry
}

func NewEvent(label string) *Event {
	event := &Event{}
	event.ID = uuid.New().String()
	event.Label = label
	event.CandidateDates = []CandidateDate{}
	return event
}

type EventEntry struct {
	ID              string
	ParticipantName string
}

type CandidateDate struct {
	ID       string
	Yyyymmdd string
}

type CandidatePreference struct {
	CandidateDateId string
}

type PreferenceType int

const (
	Ok PreferenceType = iota
	MaybeOk
	Ng
)

func (pt PreferenceType) String() string {
	switch pt {
	case Ok:
		return "Ok"
	case MaybeOk:
		return "MaybeOk"
	case Ng:
		return "Ng"
	default:
		return "Unknown"
	}
}
