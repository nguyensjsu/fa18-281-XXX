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
var mongodb_server = "mongodb://admin:cmpe281@184.169.234.66,13.56.119.112,13.52.39.43"
var mongodb_database = "TeamProject"
var mongodb_collection = "cart"
func NewServer() *negroni.Negroni {
	formatter := render.New(render.Options{
		IndentJSON: true,
	})
	n := negroni.Classic()
	mx := mux.NewRouter()
	initRoutes(mx, formatter)
	n.UseHandler(mx)
	return n
}


	


func initRoutes(mx *mux.Router, formatter *render.Render) {
	
	mx.HandleFunc("/ping", pingHandler(formatter)).Methods("GET")
	mx.HandleFunc("/Cart/{cartid}", GetCartHandler(formatter)).Methods("GET")
	mx.HandleFunc("/Cart", NewCartHandler(formatter)).Methods("POST")
	//mx.HandleFunc("/Cart/{Cartid}", UpdateCart(formatter)).Methods("PUT")
	mx.HandleFunc("/Cart/{cartid}", DeleteCartHandler(formatter)).Methods("DELETE")
	
	
	//Ping Handler
}
func pingHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct{ Test string }{"API version 1.0 alive!"})
	}
}
func NewCartHandler(formatter *render.Render) http.HandlerFunc{
	return func(w http.ResponseWriter, req *http.Request){
		session, err := mgo.Dial(mongodb_server)
		defer session.Close()
		session.SetMode(mgo.PrimaryPreferred, true)
		c := session.DB(mongodb_database).C(mongodb_collection)
		var result bson.M
		var newCart cart
		_ = json.NewDecoder(req.Body).Decode(&newCart)
		newcartid:=xid.New()
		newCart.CartID=newcartid.String()
		totalPrice := float64(0)
    for _, num := range newCart.products {
	  totalPrice += float64(num.Price * float64(num.Quantity))
		}
		newCart.TotalPrice = totalPrice	
									 

		if err != nil {
			panic(err)
		}
		fmt.Println("Connected to Database....")
		query := bson.M{"CartID": newCart.CartID,"TotalPrice":newCart.TotalPrice }
		err = c.Insert(query)
		if err != nil {
			log.Fatal(err)
		}
		err = c.Find(bson.M{"CartID": newCart.CartID}).One(&result)
		if err != nil {
			log.Fatal(err)
		}
		formatter.JSON(w, http.StatusOK, result)	
		
	}
}


//TODO
//Update cart by ID
/*func UpdateCart(formatter *render.Render) http.HandlerFunc{
	return func(w http.ResponseWriter, req *http.Request){
		session, err := mgo.Dial(mongodb_server)
		defer session.Close()
		session.SetMode(mgo.PrimaryPreferred, true)
		c := session.DB(mongodb_database).C(mongodb_collection)
		var result bson.M
		var newproduct Product
		
		_ = json.NewDecoder(req.Body).Decode(&newproduct)
		vars := mux.Vars(req)
		Cartid := vars["Cartid"]
		if err != nil {
			panic(err)
		}
		fmt.Println("Connected to Database")
		err = c.Find(bson.M{"CartID": Cartid}).One(&result)
        if result == nil {
			fmt.Println("No such Cart....")
		} else {
			formatter.JSON(w, http.StatusOK, result)
		}
		totalPrice := float64(0)
		for _, num := range newproduct.Product {
			totalPrice += float64(num.Price * float64(num.Quantity))
			  }

		
		
        

	}
	
}*/
//Get Cart Details by ID
func GetCartHandler(formatter *render.Render) http.HandlerFunc{
	return func(w http.ResponseWriter, req *http.Request){
		vars := mux.Vars(req)
		cartid := vars["cartid"]
		fmt.Println(cartid)
		session, err := mgo.Dial(mongodb_server)
		if err != nil {
			panic(err)
		}
		defer session.Close()
		session.SetMode(mgo.PrimaryPreferred, true)
		c := session.DB(mongodb_database).C(mongodb_collection)
		fmt.Println(cartid)
		var result bson.M
		err = c.Find(bson.M{"CartID": cartid}).One(&result)
		fmt.Println("Result :", result)
		if result == nil {
			fmt.Println("No such cart")
		} else {
			formatter.JSON(w, http.StatusOK, result)
		}
	}
}

func DeleteCartHandler(formatter *render.Render) http.HandlerFunc{
			return func(w http.ResponseWriter, req *http.Request){
		vars := mux.Vars(req)
		cartid:= vars["cartid"]
		session, err := mgo.Dial(mongodb_server)
		if err != nil {
			panic(err)
		}
		defer session.Close()
		session.SetMode(mgo.PrimaryPreferred, true)
		c := session.DB(mongodb_database).C(mongodb_collection)
		var result bson.M
		err = c.Find(bson.M{"CartID": cartid}).One(&result)
		if result == nil {
			fmt.Println("No such cart")
		} else {
			formatter.JSON(w, http.StatusOK, result)
		}
		c.Remove(bson.M{"CartID": cartid})

			}
		}

	

