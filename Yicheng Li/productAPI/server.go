package main

import (
	"fmt"
	"log"
	"net/http"
	// "encoding/json"
	"github.com/codegangsta/negroni"
	// "github.com/streadway/amqp"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	// "github.com/satori/go.uuid"
	"gopkg.in/mgo.v2"
    // "gopkg.in/mgo.v2/bson"
)

// MongoDB Config
var mongodb_server = "mongodb://user_name:password@ip_add"  
var mongodb_database = "TeamProject"
var mongodb_collection = "products"


// NewServer configures and returns a Server.
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

// API Routes
func initRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/ping", pingHandler(formatter)).Methods("GET")
	//mx.HandleFunc("/product", burgerCreateProductHandler(formatter)).Methods("POST")
	mx.HandleFunc("/products", burgerProductsDetailsHandler(formatter)).Methods("GET")
	//mx.HandleFunc("/product/{id}", burgerProductDetailsHandler(formatter)).Method("GET")
	//mx.HandleFunc("/product/{id}", burgerDeleteProductHandler(formatter)).Methods("DELETE")
}

// Helper Functions
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

// API Ping Handler
func pingHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct{ Test string }{"API version 1.0 alive!"})
	}
}

// API Create Products Handler
// func burgerCreateProductHandler(formatter *render.Render) http.HandlerFunc {
// 	return func(w http.ResponseWriter, req *http.Request) {
// 				// Open MongoDB Session
// 				session, err := mgo.Dial(mongodb_server)
// 		        if err != nil {
// 		                panic(err)
// 		        }
// 		        defer session.Close()
// 		        session.SetMode(mgo.Monotonic, true)
// 		        c := session.DB(mongodb_database).C(mongodb_collection)
//
// 						var newProduct product
// 						if err := json.NewDecoder(req.Body).Decode(&newProduct); err != nil {
// 								formatter.JSON(w, http.StatusBadRequest, "Invalid request payload")
// 								return
// 						}
//
// 						if err := c.Insert(&newProduct); err != nil {
// 							formatter.JSON(w, http.StatusInternalServerError, err.Error())
// 							return
// 						}
//
// 						if err != nil {
// 							log.Fatal(err)
// 						}
//
// 						formatter.JSON(w, http.StatusOK, result)
// 	}
// }

func burgerProductsDetailsHandler(formatter *render.Render) http.HandlerFunc  {
	return func(w http.ResponseWriter, req *http.Request) {
		session, err := mgo.Dial(mongodb_server)
		         if err != nil {
		                panic(err)
		        }
		        defer session.Close()
		        session.SetMode(mgo.Monotonic, true)
		        c := session.DB(mongodb_database).C(mongodb_collection)

						var results []product
						if err := c.Find(nil).All(&results); err != nil {
							formatter.JSON(w, http.StatusInternalServerError, err.Error())
							return
						}
						if err != nil {
						                log.Fatal(err)
						}
						fmt.Println("Number of Products", len(results))

						formatter.JSON(w, http.StatusOK, results)
	}
}
//
// API Gumball Machine Handler
// func gumballHandler(formatter *render.Render) http.HandlerFunc {
// 	return func(w http.ResponseWriter, req *http.Request) {
// 		session, err := mgo.Dial(mongodb_server)
//         if err != nil {
//                 panic(err)
//         }
//         defer session.Close()
//         session.SetMode(mgo.Monotonic, true)
//         c := session.DB(mongodb_database).C(mongodb_collection)
//         var result bson.M
//         err = c.Find(bson.M{"SerialNumber" : "1234998871109"}).One(&result)
//         if err != nil {
//                 log.Fatal(err)
//         }
//         fmt.Println("Gumball Machine:", result )
// 		formatter.JSON(w, http.StatusOK, result)
// 	}
// }
//
// // API Update Gumball Inventory
// func gumballUpdateHandler(formatter *render.Render) http.HandlerFunc {
// 	return func(w http.ResponseWriter, req *http.Request) {
//     	var m gumballMachine
//     	_ = json.NewDecoder(req.Body).Decode(&m)
//     	fmt.Println("Update Gumball Inventory To: ", m.CountGumballs)
// 		session, err := mgo.Dial(mongodb_server)
//         if err != nil {
//                 panic(err)
//         }
//         defer session.Close()
//         session.SetMode(mgo.Monotonic, true)
//         c := session.DB(mongodb_database).C(mongodb_collection)
//         query := bson.M{"SerialNumber" : "1234998871109"}
//         change := bson.M{"$set": bson.M{ "CountGumballs" : m.CountGumballs}}
//         err = c.Update(query, change)
//         if err != nil {
//                 log.Fatal(err)
//         }
//        	var result bson.M
//         err = c.Find(bson.M{"SerialNumber" : "1234998871109"}).One(&result)
//         if err != nil {
//                 log.Fatal(err)
//         }
//         fmt.Println("Gumball Machine:", result )
// 		formatter.JSON(w, http.StatusOK, result)
// 	}
// }
//
// // API Create New Gumball Order
// func gumballNewOrderHandler(formatter *render.Render) http.HandlerFunc {
// 	return func(w http.ResponseWriter, req *http.Request) {
// 		uuid,_ := uuid.NewV4()
//     	var ord = order {
// 					Id: uuid.String(),
// 					OrderStatus: "Order Placed",
// 		}
// 		if orders == nil {
// 			orders = make(map[string]order)
// 		}
// 		orders[uuid.String()] = ord
// 		queue_send(uuid.String())
// 		fmt.Println( "Orders: ", orders )
// 		formatter.JSON(w, http.StatusOK, ord)
// 	}
// }
//
// // API Get Order Status
// func gumballOrderStatusHandler(formatter *render.Render) http.HandlerFunc {
// 	return func(w http.ResponseWriter, req *http.Request) {
// 		params := mux.Vars(req)
// 		var uuid string = params["id"]
// 		fmt.Println( "Order ID: ", uuid )
// 		if uuid == ""  {
// 			fmt.Println( "Orders:", orders )
// 			var orders_array [] order
// 			for key, value := range orders {
//     			fmt.Println("Key:", key, "Value:", value)
//     			orders_array = append(orders_array, value)
// 			}
// 			formatter.JSON(w, http.StatusOK, orders_array)
// 		} else {
// 			var ord = orders[uuid]
// 			fmt.Println( "Order: ", ord )
// 			formatter.JSON(w, http.StatusOK, ord)
// 		}
// 	}
// }



/*

  	-- Gumball MongoDB Collection (Create Document) --

    db.gumball.insert(
	    {
	      Id: 1,
	      CountGumballs: NumberInt(202),
	      ModelNumber: 'M102988',
	      SerialNumber: '1234998871109'
	    }
	) ;

    -- Gumball MongoDB Collection - Find Gumball Document --

    db.gumball.find( { Id: 1 } ) ;

    {
        "_id" : ObjectId("54741c01fa0bd1f1cdf71312"),
        "Id" : 1,
        "CountGumballs" : 202,
        "ModelNumber" : "M102988",
        "SerialNumber" : "1234998871109"
    }

    -- Gumball MongoDB Collection - Update Gumball Document --

    db.gumball.update(
        { Dd: 1 },
        { $set : { CountGumballs : NumberInt(10) } },
        { multi : false }
    )

    -- Gumball Delete Documents

    db.gumball.remove({})

 */
