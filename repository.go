package main

import (
	"errors"
	"log"
)

type RepositoryInMemory struct {
	Events []*Event
}

var repositoryInMemory *RepositoryInMemory
var ErrNotFound = errors.New("not found")

func init() {
	repositoryInMemory = &RepositoryInMemory{}
}

func SaveEvent(event *Event) {
	log.Println("save event:" + event.ID)
	log.Println(len(repositoryInMemory.Events))
	repositoryInMemory.Events = append(repositoryInMemory.Events, event)
	log.Println(len(repositoryInMemory.Events))
}

func LoadEvent(eventID string) (*Event, error) {
	log.Println("load event")

	for _, evt := range repositoryInMemory.Events {
		if evt.ID == eventID {
			return evt, nil
		}
	}

	return &Event{}, ErrNotFound
}

func LoadEventIds() []string {
	res := []string{}

	for _, evt := range repositoryInMemory.Events {
		log.Println(evt.ID)
		res = append(res, evt.ID)
	}

	return res
}
