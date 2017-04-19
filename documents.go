package main

import (
	"time"
	"errors"
	"gopkg.in/mgo.v2/bson"
)

type DocumentType struct {
	Code string `bson:"code"`
	Name string `bson:"name"`
}

// Todo: Code uniqueness

func (d DocumentType) OK() error{
	if len(d.Name) < 5 {
		return errors.New("Name is too short")
	}
	return nil
}

type Document struct {
	ID bson.ObjectId `bson:"_id"`
	Group Group `bson:"group"`
	UserID bson.ObjectId `bson:"user_id"`
	User User `bson:"-"`
	DocumentType DocumentType `bson:"documenttype_id"`
	Name string `bson:"name"`
	DocumentDate time.Time `bson:"document_date"`
	File string `bson:"file"`
}

func (d Document) OK() error {
	if len(d.Name) < 5 {
		return errors.New("Name is too short")
	}
	return nil
}