package main

import (
	"log"

	uuid "github.com/satori/go.uuid"
)

type Event struct {
	AccId string
	Type  string
}

type CreateEvent struct {
	Event
	AccName string
}

func NewCreateAccountEvent(name string) CreateEvent {
	newuuid, err := uuid.NewV4()

	if err != nil {
		log.Fatalf("Error creating uuid: %v", err)
	}
	event := new(CreateEvent)
	event.Type = "CreateEvent"
	event.AccId = newuuid.String()
	event.AccName = name
	return *event
}
