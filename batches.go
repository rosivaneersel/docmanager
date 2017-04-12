package main

import (
	"errors"
	"time"
	"gopkg.in/mgo.v2/bson"
)

type Batch struct {
	Name string
	Emails string
	ExecutionDate time.Time
	DocumentTypes []bson.ObjectId
}

func (b Batch) OK() error {
	if b.ExecutionDate.IsZero() {
		return errors.New("Execution date can't be empty")
	}
	return nil
}