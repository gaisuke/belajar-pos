package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type Products struct {
	ID        int
	Name      string
	Price     int
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type OrderItems struct {
	OrderId   int
	ProductId int
	Qty       int
}

type Orders struct {
	ID         int
	TotalPrice int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "asdasd"
	dbname   = "belajar-pos"
)

var db *sql.DB

func main() {
	db = initDB()
	defer db.Close()

	router := mux.NewRouter()
	router.HandleFunc("/products", GetAllProducts).Methods("GET")
	// router.HandleFunc("/products/{id}", GetProductById).Methods("GET")
	router.HandleFunc("/products", CreateProduct).Methods("POST")
	// router.HandleFunc("/products/{id}", UpdateProduct).Methods("PUT")
	// router.HandleFunc("/products/{id}", DeleteProduct).Methods("DELETE")
	// router.HandleFunc("/orders", GetAllOrders).Methods("GET")
	// router.HandleFunc("/orders/{id}", GetOrderById).Methods("GET")
	// router.HandleFunc("/orders", CreateOrder).Methods("POST")

	fmt.Println("server running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))

}

func initDB() *sql.DB {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT * FROM products")
	if err != nil {
		ResponseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer rows.Close()
	var products []Products
	for rows.Next() {
		var product Products
		err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Status, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			ResponseWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		products = append(products, product)
	}
	ResponseWithJSON(w, http.StatusOK, products)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product Products
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	product.ID = 0
	_, err := db.Exec("INSERT INTO products (name, price, status) VALUES ($1, $2, $3)", product.Name, product.Price, product.Status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res := json.NewEncoder(w).Encode(&product)
	ResponseWithJSON(w, http.StatusOK, res)
}

func ResponseWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func ResponseWithError(w http.ResponseWriter, code int, message string) {
	ResponseWithJSON(w, code, map[string]string{"error": message})
}
