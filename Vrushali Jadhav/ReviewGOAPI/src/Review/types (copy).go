package main

import "gopkg.in/mgo.v2/bson"

type review struct {
	ReviewID           bson.ObjectId `bson:"_id" json:"ReviewID"`
	productIDString    string        `bson:"productIDString" json:"productIDString"`
	ReviewString       string        `bson:"ReviewString" json:"ReviewString"`	
}
