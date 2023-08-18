package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//creamos conexion a db
const url = "root:123456@tpc(localhost:3306)/goweb_db"

var db *sql.DB

func Connect(){
	conection, err := sql.Open("mysql", url)
	if err != nil {
		panic(err)
	}
	fmt.Println("Conexion Exitosa")
	db = conection
}

func Close(){
	db.Close()
}