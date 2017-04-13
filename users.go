package main

import (
	"gopkg.in/mgo.v2/bson"
	"golang.org/x/crypto/bcrypt"
)

type users struct {
	db *db
}

func (u *users) GetUserByAuthentication(email string, password string) (*User, error) {
	user := &User{}
	c := u.db.GetCollection("users")
	err := c.Find(bson.M{"email": email}).One(user)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *users) GetUserByID(id string) (*User, error) {
	user := &User{}
	c := u.db.GetCollection("users")
	err := c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *users) Create(user *User) error {
	err := user.OK()
	if err != nil {
		return err
	}

	if !user.ID.Valid() {
		user.ID = bson.NewObjectId()
	}

	c := u.db.GetCollection("users")
	err = c.Insert(user)
	if err != nil {
		return err
	}
	return nil
}

func (u *users) Update(user *User) error  {
	err := user.OK()
	if err != nil {
		return err
	}

	c := u.db.GetCollection("users")
	err = c.Update(bson.M{"_id": user.ID}, user)
	if err != nil {
		return err
	}
	return nil
}

func NewUserController(db *db) *users {
	return &users{db: db}
}

type User struct {
	ID bson.ObjectId `bson:"_id"`
	Username string
	Email string
	Password string
	Groups []bson.ObjectId
}

func (u User) OK() error {
	return nil
}

func (u *User) SetPassword(password string) error {
	p := []byte(password)

	h, err := bcrypt.GenerateFromPassword(p, bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(h)
	return nil
}
