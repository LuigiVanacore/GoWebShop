package main

import (
	"encoding/json"

	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type Product struct {
	ID       int    `json:id`
	Name     string `json:name`
	Price    int    `json:price`
	Quantity int    `json:quantity`
}

var Products []Product

func getProducts(w http.ResponseWriter, r* http.Request) {
	log.Println("getProducts")

	json.NewEncoder(w).Encode(Products)
}

func getProduct(w http.ResponseWriter, r* http.Request) {
	log.Println("getProduct")

	params := mux.Vars(r)

	id, _ := strconv.Atoi(params["id"])

	for _, product := range Products {
		if product.ID == id {
			json.NewEncoder(w).Encode(&product)
		}
	}
}

func addProduct(w http.ResponseWriter, r* http.Request) {
	log.Println("addProduct")

	var product Product

	json.NewDecoder(r.Body).Decode(&product)

	Products = append(Products,product)
}

func updateProduct(w http.ResponseWriter, r* http.Request)  {
	log.Println("updateProduct")

	var product Product

	json.NewDecoder(r.Body).Decode(&product)

	for i, temp := range Products {
		if product.ID == temp.ID {
			Products[i] = product
		}
	}
}

func removeProduct(w http.ResponseWriter, r* http.Request) {
log.Println("removeProduct")

	params := mux.Vars(r)

	id, _ := strconv.Atoi(params["id"])

	for i, product := range Products {
		if product.ID == id {
			Products = append(Products[:i], Products[i+1:]... )
		}
	}

}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/products", getProducts).Methods("GET")
	router.HandleFunc("/products/{id}", getProduct).Methods("GET")
	router.HandleFunc("/products", addProduct).Methods("POST")
	router.HandleFunc("/products", updateProduct).Methods("PUT")
	router.HandleFunc("/products", removeProduct).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
