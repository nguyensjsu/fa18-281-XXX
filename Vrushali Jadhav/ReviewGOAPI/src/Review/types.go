package main

import "gopkg.in/mgo.v2/bson"

type review struct {
	ReviewID     bson.ObjectId `bson:"_id" json:"ReviewID"`
	ProductName  string        `bson:"ProductName" json:"ProductName"`
        Name         string        `bson:"Name" json:"Name"`
	Reviews      string        `bson:"Reviews" json:"Reviews"`	
}
