package app

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
)

type Customer struct {
	Name string `json:"full_name xml:"name"`
	City string `json:"city xml:"city"`
	ZipCode string `json:"zip_code xml:"zipcode"`
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer {
		{ Name: "Carlos", City: "São Paulo", ZipCode: "00000-000" },
		{ Name: "Neto", City: "São Paulo", ZipCode: "00000-000" },
	}

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header.Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header.Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}

}