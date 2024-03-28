package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	initDb()
	http.HandleFunc("/resolve", resolve)
	http.HandleFunc("/register", register)
	log.Print("running server")
	log.Print(http.ListenAndServe("0.0.0.0:8080", nil))
}
func resolve(w http.ResponseWriter, req *http.Request) {
	name := req.URL.Query().Get("name")
	ip := getIp(name)
	fmt.Fprintf(w,ip)
}
func register(w http.ResponseWriter, req *http.Request) {
	name := req.URL.Query().Get("name")
	ip := req.URL.Query().Get("ip")
	insertIp(ip,name)
	log.Println("registered")
}
func initDb() {
	os.Remove("sqlite-database.db") 
	log.Println("Creating sqlite-database.db...")
	file, err := os.Create("sqlite-database.db")
	if err != nil {
		log.Fatal(err.Error())
	}
	file.Close()
	log.Println("sqlite-database.db created")
	sqliteDatabase, _ := sql.Open("sqlite3", "./sqlite-database.db")
	defer sqliteDatabase.Close()
	createTable(sqliteDatabase)
}
func createTable(db *sql.DB) {
	createAliasTableSQL := `CREATE TABLE alias (
		"idAlias" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"name" TEXT,
		"ip" TEXT
	  );`
	log.Println("Create alias table...")
	statement, err := db.Prepare(createAliasTableSQL) 
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec() 
	log.Println("alias table created")
}

func insertIp(ip string, name string) {
	db, _ := sql.Open("sqlite3", "./sqlite-database.db") 
	defer db.Close() 
	insertStudentSQL := `INSERT OR REPLACE INTO alias(ip, name) VALUES (?, ?)`
	statement, err := db.Prepare(insertStudentSQL) 
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(ip, name)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func getIp(name string) string{
	db, _ := sql.Open("sqlite3", "./sqlite-database.db") 
	defer db.Close()
	row, err := db.Query("SELECT ip FROM alias WHERE name='" + name + "'")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	var ip string
	for row.Next() { 
		row.Scan(&ip)
	}
	return ip
}
