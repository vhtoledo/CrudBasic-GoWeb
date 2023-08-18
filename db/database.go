package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//creamos conexion a db
const url = "root:123456@tcp(localhost:3306)/goweb_db"

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

// Verificar si una tabla existe o no
func ExistsTable(tableName string) bool {
	sql := fmt.Sprintf("SHOW TABLES LIKE '%s'", tableName)
	rows, err := db.Query(sql)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	return rows.Next()
}

// crea tabla usuarios
func CreateTable(schema, name string) {

	if !ExistsTable(name) {
		_, err := db.Exec(schema)
		if err != nil {
			fmt.Println(err)
		}
	}
}

//Eliminara Tabla
func TruncateTable(tableName string) {
	sql := fmt.Sprintf("TRUNCATE %s", tableName)
	db.Exec(sql)
}

//Polimorfismo a Exec
func Exec(query string, args ...interface{}) (sql.Result, error) {
	result, err := db.Exec(query, args...)
	if err != nil {
		fmt.Println(err)
	}
	return result, err
}

//Polimorfismo a Query
func Query(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := db.Query(query, args...)
	if err != nil {
		fmt.Println(err)
	}
	return rows, err
}