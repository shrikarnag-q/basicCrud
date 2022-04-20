package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var product Product
	json.NewDecoder(r.Body).Decode(&product)
	DB.Create(&product)
	json.NewEncoder(w).Encode(product)

}

func UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var product Product
	DB.First(&product, params["id"])
	json.NewDecoder(r.Body).Decode(&product)
	DB.Save(&product)
	json.NewEncoder(w).Encode(product)

}

func GetAllProductsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var products []Product
	DB.Find(&products)
	json.NewEncoder(w).Encode(products)

}

func GetProductHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var product Product
	DB.First(&product, params["id"])
	json.NewEncoder(w).Encode(product)

}

func DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var product Product
	DB.Delete(&product, params["id"])
	json.NewEncoder(w).Encode("The product is Deleted Successfully!")

}
