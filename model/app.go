package model

import "gopkg.in/mgo.v2/bson"

// App for database
type App struct {
	ID         bson.ObjectId `json:"id" bson:"_id"`
	Name       string        `json:"name" bson:"name"`
	Token      string        `json:"token" bson:"token"`
	Maintainer string        `json:"maintainer" bson:"maintainer"`
}
