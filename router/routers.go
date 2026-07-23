package router

import (
	"github.com/gorilla/mux"
	"github.com/luminous479/Stocks-API/middleware"
)


func Routers() {		
  
	router := mux.NewRouter()

	router.HandleFunc("/api/stocks/{id}", middleware.GetStocks).Methods("GET")	
	router.HandleFunc("/api/stocks", middleware.GetAllStocks).Methods("GET")
	router.HandleFunc("/api/stocks", middleware.CreateStock).Methods("POST")
	router.HandleFunc("/api/stocks/{id}", middleware.UpdateStock).Methods("PUT")
	router.HandleFunc("/api/stocks/{id}", middleware.DeleteStock).Methods("DELETE")

}