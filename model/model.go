package model

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Session is an interface to access to the Session struct.
type Session interface {
	DB(name string) DataLayer
	Close()
}

// MongoSession is currently a Mongo session.
type MongoSession struct {
	*mgo.Session
}

// DB shadows *mgo.DB to returns a DataLayer interface instead of *mgo.Database.
func (s MongoSession) DB(name string) DataLayer {
	return &MongoDatabase{Database: s.Session.DB(name)}
}

// DataLayer is an interface to access to the database struct.
type DataLayer interface {
	C(name string) Collection
}

// MongoCollection wraps a mgo.Collection to embed methods in models.
type MongoCollection struct {
	*mgo.Collection
}

// MongoQuery wraps a mgo.Query to embed methods in models.
type MongoQuery struct {
	*mgo.Query
}

// Collection is an interface to access to the collection struct.
type Collection interface {
	Find(query interface{}) Query
	Count() (n int, err error)
	Insert(docs ...interface{}) error
	Remove(selector interface{}) error
	Update(selector interface{}, update interface{}) error
	GetMyDocuments() ([]interface{}, error)
}

// MongoDatabase wraps a mgo.Database to embed methods in models.
type MongoDatabase struct {
	*mgo.Database
}

// C shadows *mgo.DB to returns a DataLayer interface instead of *mgo.Database.
func (d MongoDatabase) C(name string) Collection {
	return &MongoCollection{Collection: d.Database.C(name)}
}

// GetMyDocuments returns all scores.
func (c *MongoCollection) GetMyDocuments() ([]interface{}, error) {
	var documents []interface{}
	err := c.Find(bson.M{}).All(&documents)
	if err != nil {
		return nil, err
	}
	return documents, nil
}

// Find shadows *mgo.Collection to returns a Query interface instead of *mgo.Query.
func (c MongoCollection) Find(query interface{}) Query {
	return MongoQuery{Query: c.Collection.Find(query)}
}

// Query is an interface to access to the database struct
type Query interface {
	All(result interface{}) error
	One(result interface{}) (err error)
	Distinct(key string, result interface{}) error
}
