package main

import "gopkg.in/mgo.v2/bson"

type Product struct {
	ProductID          bson.ObjectId 
	Name        string  
	Quantity       int      
	Price 		float64        
}
type cart struct {
	CartID	string
    TotalPrice float64
	
	products 	[]Product
}