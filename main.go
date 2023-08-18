package main

import (
	"fmt"
	"gomysql/db"
	"gomysql/models"
)

func main() {
	//Verificar conexion
	db.Connect()
	//db.Ping()
	//fmt.Println(db.ExistsTable("users"))

	//Crear tabla
	//db.CreateTable(models.UserSchema, "users")

	// Insertar un registro
	//user := models.CreateUser("victor", "victor123", "victor@gmail.com")
	//fmt.Println(user)

	// Obtener todos los registros
	//users := models.ListUsers()
	//fmt.Println(users)

	//Obtener un los registros
	users := models.GetUser(1)
	fmt.Println(users)

	// Eliminar tabla
	//db.TruncateTable("users")

	//cerrar conexion
	db.Close()
}