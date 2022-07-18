package entities

import uuid "github.com/satori/go.uuid"

type Student struct {
	Id     uuid.UUID
	Name   string
	Age    int
	Gender bool
}
