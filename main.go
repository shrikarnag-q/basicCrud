package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	InitializeDB()
	InitializeRouter()

}

func InitializeRouter() {
	r := mux.NewRouter()
	r.HandleFunc("/addproduct", CreateProductHandler).Methods("POST")
	r.HandleFunc("/updateproduct/{id}", UpdateProductHandler).Methods("PUT")
	r.HandleFunc("/deleteproduct/{id}", DeleteProductHandler).Methods("DELETE")
	r.HandleFunc("/products", GetAllProductsHandler).Methods("GET")
	r.HandleFunc("/product/{id}", GetProductHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":9000", r))
}
