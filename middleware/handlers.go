package middleware

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/luminous479/Stocks-API/models"
)

type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

func CreateConnection() *sql.DB {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := sql.Open("mysql", os.Getenv("DB_USER")+":"+os.Getenv("DB_PASSWORD")+"@tcp("+os.Getenv("DB_HOST")+":"+os.Getenv("DB_PORT")+")/"+os.Getenv("DB_NAME"))
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func CreateStock(w http.ResponseWriter, r *http.Request) {

	var stock models.Stock

	err := json.NewDecoder(r.Body).Decode(&stock)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db := CreateConnection()
	defer db.Close()

	result, err := db.Exec("INSERT INTO stocks (stockid, name, price, company) VALUES (?, ?, ?, ?)", stock.StockID, stock.Name, stock.Price, stock.Company)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := response{
		ID:      id,
		Message: "Stock created successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func GetAllStocks() {

}

func GetStocks(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	stockID := params["id"]

	db := CreateConnection()
	defer db.Close()

	var stock models.Stock
	err := db.QueryRow("SELECT stockid, name, price, company FROM stocks WHERE stockid = ?", stockID).Scan(&stock.StockID, &stock.Name, &stock.Price, &stock.Company)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Stock not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stock)

}

func UpdateStock() {

}

func DeleteStock() {

}
