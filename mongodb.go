package main

import "gopkg.in/mgo.v2"

type db struct {
	servers string
	name string
	session *mgo.Session
	database *mgo.Database
}

func (db *db) Connect() error {
	session, err := mgo.Dial(db.servers)
	if err != nil {
		return err
	}
	db.session = session

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	db.database = db.session.DB(db.name)
	return nil
}

func (db *db) Close() {
	db.session.Close()
}


func (db *db) GetCollection(name string) *mgo.Collection {
	return db.database.C(name)
}

func NewDatabase(servers string, name string) (*db, error) {
	db := &db{servers: servers, name: name}
	err := db.Connect()
	if err != nil {
		return nil, err
	}
	return db, nil
}

