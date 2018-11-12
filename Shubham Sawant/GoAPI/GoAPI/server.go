package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/rs/xid"
	"github.com/unrolled/render"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// MongoDB Config
var mongodb_server = "mongodb://admin:cmpe281@34.221.158.202,54.185.27.49,18.237.78.8,54.186.122.30,34.216.148.220"

//var mongodb_server1 string
//var mongodb_server2 string
//var redis_server string

//="mongodb://54.67.13.87:27017,54.67.106.101:27017,13.57.39.192:27017,54.153.26.217://27017,52.53.154.42:27017"

var mongodb_database = "TeamProject"
var mongodb_collection = "user"

// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {
	formatter := render.New(render.Options{
		IndentJSON: true,
	})

	//mongodb_server = os.Getenv("MONGO1")
	//mongodb_server1 = os.Getenv("MONGO2")
	//mongodb_server2 = os.Getenv("MONGO3")
	//mongodb_database = os.Getenv("MONGO_DB")
	//mongodb_collection = os.Getenv("MONGO_COLLECTION")
	//redis_server = os.Getenv("REDIS")

	n := negroni.Classic()
	mx := mux.NewRouter()
	initRoutes(mx, formatter)
	n.UseHandler(mx)
	return n
}

// API Routes
func initRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/ping", pingHandler(formatter)).Methods("GET")
	mx.HandleFunc("/users/{email}", userHandler(formatter)).Methods("GET")
	mx.HandleFunc("/users", userSignUpHandler(formatter)).Methods("POST")
	mx.HandleFunc("/users/{email}", deleteUserHandler(formatter)).Methods("DELETE")
	mx.HandleFunc("/users", updateUserHandler(formatter)).Methods("PUT")

}

// API Ping Handler
func pingHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct{ Test string }{"API version 1.0 alive!"})
	}
}

// API  Handler --------------- Get the user info (GET) ------------------
func userHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		vars := mux.Vars(req)
		email := vars["email"]

		session, err := mgo.Dial(mongodb_server)
		if err != nil {
			panic(err)
		}

		defer session.Close()
		session.SetMode(mgo.PrimaryPreferred, true)
		c := session.DB(mongodb_database).C(mongodb_collection)

		fmt.Println(email)
		var result bson.M
		err = c.Find(bson.M{"Email": email}).One(&result)

		fmt.Println("Result :", result)
		if result == nil {
			var noUser user
			noUser.Message = "User does not exist"
			formatter.JSON(w, http.StatusOK, noUser)
		} else {
			formatter.JSON(w, http.StatusOK, result)
		}
	}
}

// API  Handler --------------- Register a new user (POST) ------------------
func userSignUpHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		session, err := mgo.Dial(mongodb_server)
		defer session.Close()
		session.SetMode(mgo.PrimaryPreferred, true)
		c := session.DB(mongodb_database).C(mongodb_collection)

		var result bson.M
		var newUser user
		_ = json.NewDecoder(req.Body).Decode(&newUser)

		uuidForNewUser := xid.New()

		newUser.UserID = uuidForNewUser.String()

		fmt.Println("Register new user ")

		if err != nil {
			panic(err)
		}
		fmt.Println("Connected to Database")

		query := bson.M{"UserID": newUser.UserID, "Name": newUser.Name, "Email": newUser.Email, "Password": newUser.Password}
		err = c.Insert(query)
		if err != nil {
			log.Fatal(err)
		}

		err = c.Find(bson.M{"UserID": newUser.UserID}).One(&result)
		if err != nil {
			log.Fatal(err)
		}
		formatter.JSON(w, http.StatusOK, result)
	}
}

// API  Handler --------------- Delete a user (DELETE) ------------------
func deleteUserHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		vars := mux.Vars(req)
		email := vars["email"]

		session, err := mgo.Dial(mongodb_server)

		if err != nil {
			panic(err)
		}

		defer session.Close()
		session.SetMode(mgo.PrimaryPreferred, true)
		c := session.DB(mongodb_database).C(mongodb_collection)

		c.Remove(bson.M{"Email": email})
		var deletedUser user
		deletedUser.Email = email
		deletedUser.Message = "User has been deleted"
		formatter.JSON(w, http.StatusOK, deletedUser)
	}
}

//TODO
func updateUserHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		session, err := mgo.Dial(mongodb_server)

		if err != nil {
			panic(err)
		}

		defer session.Close()
		session.SetMode(mgo.PrimaryPreferred, true)
		c := session.DB(mongodb_database).C(mongodb_collection)

		var newUser user
		_ = json.NewDecoder(req.Body).Decode(&newUser)

		var result bson.M
		err = c.Find(bson.M{"Email": newUser.Email}).One(&result)

		err = c.Update(bson.M{"Email": newUser.Email}, newUser)

		var updatedUser user
		updatedUser.Email = newUser.Email
		updatedUser.Message = "User details have been successfully updated."
		formatter.JSON(w, http.StatusOK, updatedUser)

	}
}
