package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

// proccess checks the values in the URL for integer equivalence
// Pointers used to save memory
func process(price *string, time *string) error {
	_, err := strconv.ParseFloat(*price, 64)
	if err != nil {
		return err
	}
	_, err = strconv.ParseFloat(*time, 64)
	if err != nil {
		return err
	}
	return nil
}

func query(crit string, struc *Detail, variable string) error {

	qry := "SELECT * From Food WHERE " + crit + " = ?"
	row := db.QueryRow(qry, variable)

	return row.Scan(&struc.ID, &struc.Name, &struc.Price, &struc.MakeTime)
}

func update(crit string, var1 string, var2 string) error {
	updqry := "UPDATE Food SET " + crit + " = ? WHERE ID = ?"
	_, err := db.Exec(updqry, var1, var2)
	if err != nil {
		return err
	}
	return nil
}

func render(tmp *Detail, w *http.ResponseWriter) {
	ref := *tmp
	json, err := json.MarshalIndent(&ref, "", "	")
	if err != nil {
		log.Fatal(err)
	}
	res := *w
	res.Write(json)
}
