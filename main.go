package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Response struct {
	date string
	code string
}

func main() {
	/*router := mux.NewRouter()
	router.HandleFunc("/currency/save/{date}")
	*/
	sp_GetRates("cba", "1996-05-13")
}
func DB() *sql.DB {
	connStr := "user=postgres password=1234 dbname=Test sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	fmt.Println("Succesfully connected")
	//defer db.Close()
	return db
}

var (
	id     int
	title  string
	code   string
	value  float32
	a_date string
)

func sp_GetRates(code string, a_date string) {
	var db = DB()
	rows, err := db.Query("Select * from r_currency where a_date=$2 and code=$1 ", code, a_date)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		err := rows.Scan(&id, &title, &code, &value, &a_date)
		if err != nil {
			panic(err)
		}
		log.Println(id, title, code, value, a_date)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
