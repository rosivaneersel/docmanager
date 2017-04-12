package main

import (
	"errors"
	"gopkg.in/mgo.v2/bson"
)

type Group struct {
	ID bson.ObjectId `bson:"_id"`
	Name string
	Email string
	Users []bson.ObjectId
	DocumentTypes []DocumentType
	Documents []Document
	Batches []Batch
}

func (g Group) OK() error{
	if len(g.Name) < 5 {
		return errors.New("Name is too short")
	}
	return nil
}