package main

import (
	"database/sql"
	"fmt"
	"time"

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

func main() {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("hello world!")
}
