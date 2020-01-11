package main

import (
	"encoding/json"
	"gorilla/mux"
	"github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"strconv"
)

// Detail represents the food's info
type Detail struct {
	ID       int
	Name     string
	Price    string
	MakeTime string
}

var food = []Detail{

	Detail{
		ID:       1,
		Name:     "Pizza",
		Price:    "$4.99",
		MakeTime: "10mins",
	},
	Detail{
		ID:       2,
		Name:     "Cheese Burger",
		Price:    "$4.53",
		MakeTime: "5mins",
	},
	Detail{
		ID:       3,
		Name:     "Sandwich",
		Price:    "$3.99",
		MakeTime: "3mins",
	},
	Detail{
		ID:       4,
		Name:     "Burger",
		Price:    "$5.99",
		MakeTime: "4mins",
	},
	Detail{
		ID:       5,
		Name:     "Ice Cream",
		Price:    "$2.50",
		MakeTime: "nil",
	},
}

// GetByID returns the details of food by and Id
func GetByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var tmp Detail

	str, err := strconv.Atoi(id)
	if err != nil {
		w.Write([]byte("ID value in URL Request is Invalid"))
	}

	for ix := range food {
		ref := &food[ix]
		if str == ref.ID {
			tmp = *ref
			break
		}
	}
	json, err := json.MarshalIndent(tmp, "", "	")
	if err != nil {
		log.Fatal(err)
	}
	w.Write(json)
}

// GetByName returns the details of the food by searching name
func GetByName(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	var tmp Detail

	for ix := range food {
		ref := &food[ix]
		if name == ref.Name {
			tmp = *ref
			break
		}
	}

	json, err := json.MarshalIndent(tmp, "", "	")
	if err != nil {
		log.Fatal(err)
	}
	w.Write(json)
}

// AddNew add new food to the array
// This method assumes a form or cURL will be used to send the data
func AddNew(w http.ResponseWriter, r *http.Request) {
	id := newID()
	name := r.FormValue("name")
	price := r.FormValue("price")
	time := r.FormValue("time")

	err := process(&price, &time)
	if err != nil {
		log.Fatal(err)
		w.Write([]byte("Invalid Data Supplied: " + err.Error()))
	}

	var tmp *Detail
	tmp.ID = id
	tmp.Name = name
	tmp.Price = price
	tmp.MakeTime = time

	food = append(food, *tmp) //add the supplied data to the slice

	json, err := json.MarshalIndent(food, "", "	")
	if err != nil {
		log.Fatal(err)
	}
	w.Write(json)
}
