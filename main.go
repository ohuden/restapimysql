package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB

//Student structure
type Student struct {
	id    int
	name  string
	score int
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/students", studentsIndex).Methods("GET")
	router.HandleFunc("/students/{id}", getByID).Methods("GET")
	router.HandleFunc("/students/{id}", deleteByID).Methods("DELETE")
	router.HandleFunc("/students/{id}/{name}/{score}", updateByID).Methods("PUT")
	router.HandleFunc("/students/{name}/{score}", addStudent).Methods("POST")
	log.Fatal(http.ListenAndServe(":5555", router))
}
func updateByID(w http.ResponseWriter, r *http.Request) {

}
func deleteByID(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "uhqsapzt_root:ujCeofPog5@tcp(https://coolstickers.pro:3306)/uhqsapzt_test")
	if err = db.Ping(); err != nil {
		panic(err)
	}
	defer db.Close()

	fmt.Println("connected to database")
	params := mux.Vars(r)
	i, _ := strconv.Atoi(params["id"])

	result, err := db.Exec("DELETE FROM students WHERE id=?", i)
	_, err = result.LastInsertId()

	fmt.Fprintf(w, "deletedID = %d", i)
}

func getByID(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "uhqsapzt_root:ujCeofPog5@tcp(https://coolstickers.pro:3306)/uhqsapzt_test")
	if err = db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("connected to your database")
	defer db.Close()
	params := mux.Vars(r)
	i := params["id"]
	rows, err := db.Query("SELECT * FROM students WHERE id=?", i)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	defer rows.Close()

	for rows.Next() {
		bk := Student{}
		err := rows.Scan(&bk.id, &bk.name, &bk.score)
		fmt.Fprintf(w, "%s, %s, %s\n", strconv.Itoa(bk.id), bk.name, strconv.Itoa(bk.score))
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
	}
	if err = rows.Err(); err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

}

func addStudent(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "uhqsapzt_root:ujCeofPog5@tcp(https://coolstickers.pro:3306)/uhqsapzt_test")
	if err = db.Ping(); err != nil {
		panic(err)
	}
	defer db.Close()
	fmt.Println("connected to your database")
	params := mux.Vars(r)

	n := params["name"]
	s, _ := strconv.Atoi(params["score"])
	//fmt.Println(n, s)

	result, err := db.Exec("INSERT INTO students (`name`, `score`) VALUES (?, ?)", n, s)
	lastID, err := result.LastInsertId()

	fmt.Fprintf(w, "LastInsertId: %d", lastID)

}

func studentsIndex(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "uhqsapzt_root:ujCeofPog5@tcp(https://coolstickers.pro:3306)/uhqsapzt_test")
	if err = db.Ping(); err != nil {
		panic(err)
	}
	defer db.Close()
	fmt.Println("connected to your database")

	rows, err := db.Query("SELECT * FROM students")
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	defer rows.Close()

	bks := make([]Student, 0)
	for rows.Next() {
		bk := Student{}
		err := rows.Scan(&bk.id, &bk.name, &bk.score)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
		bks = append(bks, bk)
	}
	if err = rows.Err(); err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	for _, bk := range bks {
		fmt.Fprintf(w, "%s, %s, %s\n", strconv.Itoa(bk.id), bk.name, strconv.Itoa(bk.score))
	}
}
