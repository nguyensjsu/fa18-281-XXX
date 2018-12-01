package main
import (	
	"log"
	"net/http"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/rs/xid"
	"github.com/unrolled/render"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var mongodb_server = "mongodb://admin:admin@10.0.2.249"
var mongodb_database = "cmpe281"
var mongodb_collection = "payment"

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
	mx.HandleFunc("/Pay/{paymentid}", GetPaymentHandler(formatter)).Methods("GET")
	mx.HandleFunc("/Pay", NewPaymentHandler(formatter)).Methods("POST")
}

//Ping Handler
func pingHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct{ Test string }{"API version 1.0 alive!"})
	}
}

// Get PaymentID
func GetPaymentHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		vars := mux.Vars(req)
    	paymentid:=vars["paymentid"]
    	session, err := mgo.Dial(mongodb_server)
		if err != nil {
			formatter.JSON(w, http.StatusInternalServerError, err.Error())
			return
		}

		defer session.Close()
    	session.SetMode(mgo.PrimaryPreferred, true)
		c := session.DB(mongodb_database).C(mongodb_collection)
    	var result bson.M
		err = c.Find(bson.M{"PaymentID": paymentid}).One(&result)
    	if result == nil {
			var notPaid payment
			formatter.JSON(w, http.StatusOK, notPaid)
		} else {
			formatter.JSON(w, http.StatusOK, result)
		}
	}
}

// Make Payment
func NewPaymentHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		session, err := mgo.Dial(mongodb_server)

		if err != nil {
			formatter.JSON(w, http.StatusInternalServerError, err.Error())
			return
		}

		defer session.Close()
		session.SetMode(mgo.PrimaryPreferred, true)
		c := session.DB(mongodb_database).C(mongodb_collection)
		var newPayment payment
		
		newpaymentid:=xid.New()
		newPayment.PaymentID=newpaymentid.String()
	    if err := c.Insert(&newPayment); err != nil {
			formatter.JSON(w, http.StatusInternalServerError, err.Error())
			return
		}
		if err := c.Insert(&newPayment); err != nil {
			formatter.JSON(w, http.StatusInternalServerError, err.Error())
			return
		}
		var result payment
		if err = c.Find(bson.M{"PaymentID": newPayment.PaymentID}).One(&result); err != nil {
			formatter.JSON(w, http.StatusInternalServerError, err.Error())
			return
		}
		if err != nil {
			log.Fatal(err)
		}
		formatter.JSON(w, http.StatusOK, result)
	}
}