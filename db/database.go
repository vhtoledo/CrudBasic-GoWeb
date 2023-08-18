package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//creamos conexion a db
const url = "root:123456@tpc(localhost:3306)/goweb_db"
// Guardar la Conexion
var db *sql.DB

// Realiza la conexion
func Connect(){
	conection, err := sql.Open("mysql", url)
	if err != nil {
		panic(err)
	}
	fmt.Println("Conexion Exitosa")
	db = conection
}

// Cierra la conexion
func Close(){
	db.Close()
}

// Verificar conexion
func Ping(){
	if err := db.Ping(); err != nil {
		panic(err)
	}
}

// crea tabla usuarios
func CreateTable(schema string) {
	db.Exec(schema)
}