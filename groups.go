package main

import (
	"errors"
	"gopkg.in/mgo.v2/bson"
)

type Group struct {
	ID bson.ObjectId `bson:"_id"`
	Name string `bson:"name"`
	Email string `bson:"email"`
	DocumentTypes []DocumentType `bson:"document_types"`
	Documents []Document `bson:"documents"`
	Batches []Batch `bson:"batches"`
}

func (g Group) OK() error{
	if len(g.Name) < 5 {
		return errors.New("Name is too short")
	}
	return nil
}

func (g *Group) FindDocumentType(code string) (int, bool) {
	for i := 0; i < len(g.DocumentTypes); i++ {
		if g.DocumentTypes[i].Code == code {
			return i, true
		}
	}
	return -1, false
}

func (g *Group) CreateOrUpdateDocumentType(code string, dt DocumentType) {
	i, found := g.FindDocumentType(code)
	if found {
		g.DocumentTypes[i] = dt
	} else {
		g.DocumentTypes = append(g.DocumentTypes, dt)
	}
}

type groups struct {
	db *db
}

func (g *groups) GetByID(id string) (*Group, error) {
	c := g.db.GetCollection("groups")
	group := &Group{}
	err := c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&group)
	if err != nil {
		return nil, err
	}
	return group, nil
}

func (g *groups) GetByIDs(ids []bson.ObjectId) ([]Group, error) {
	c := g.db.GetCollection("groups")
	var groups []Group

	err := c.Find(bson.M{"_id":bson.M{"$in" : ids}}).All(&groups)
	if err != nil {
		return nil, err
	}
	return groups, nil
}

func (g *groups) Create(group *Group) error {
	err := group.OK()
	if err != nil {
		return err
	}

	if !group.ID.Valid() {
		group.ID = bson.NewObjectId()
	}

	c := g.db.GetCollection("groups")
	err = c.Insert(group)
	if err != nil {
		return err
	}
	return nil
}

func (g *groups) Update(group *Group) error {
	err := group.OK()
	if err != nil {
		return err
	}

	c:= g.db.GetCollection("groups")
	err = c.Update(bson.M{"_id": group.ID}, group)
	if err != nil {
		return err
	}
	return nil
}

func NewGroupController(db *db) *groups {
	return &groups{db: db}
}
