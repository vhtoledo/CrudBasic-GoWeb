package main

import (
	"fmt"
	"gomysql/db"
	"gomysql/models"
)

func main() {
	//Verificar conexion
	db.Connect()
	db.Ping()
	fmt.Println(db.ExistsTable("users"))

	//Clear tabla
	db.CreateTable(models.UserSchema, "users")
}