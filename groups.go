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

type groups struct {
	db *db
}

func (u *groups) GetByID(id string) (*Group, error) {
	c := u.db.GetCollection("groups")
	group := &Group{}
	err := c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&group)
	if err != nil {
		return nil, err
	}
	return group, nil
}

func (u *groups) GetByUserID(id string) ([]Group, error) {
	c := u.db.GetCollection("groups")
	var groups []Group

	err := c.Find(bson.M{"users":bson.M{"$elemMatch" : bson.M{"$eq": bson.ObjectIdHex(id)}}}).All(&groups)
	if err != nil {
		return nil, err
	}
	return groups, nil
}

func (u *groups) Create(group *Group) error {
	err := group.OK()
	if err != nil {
		return err
	}

	if !group.ID.Valid() {
		group.ID = bson.NewObjectId()
	}

	c := u.db.GetCollection("groups")
	err = c.Insert(group)
	if err != nil {
		return err
	}
	return nil
}

func NewGroupController(db *db) *groups {
	return &groups{db: db}
}
