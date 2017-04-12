package main

import (
	"time"
	"errors"
	"gopkg.in/mgo.v2/bson"
)

type DocumentType struct {
	Name string
}

func (d DocumentType) OK() error{
	if len(d.Name) < 5 {
		return errors.New("Name is too short")
	}
	return nil
}

type Document struct {
	ID bson.ObjectId `bson:"_id"`
	Group Group
	UserID bson.ObjectId `bson:"_id"`
	User User `bson:"-"`
	DocumentType DocumentType
	Name string
	DocumentDate time.Time
	File string
}

func (d Document) OK() error {
	if len(d.Name) < 5 {
		return errors.New("Name is too short")
	}
	return nil
}