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

	//Insertar un registro
	//user := models.CreateUser("agostina", "agostina123", "agostina@gmail.com")
	//fmt.Println(user)

	// Obtener todos los registros
	//user := models.ListUsers()
	//fmt.Println(user)

	//Obtener un los registros
	//user := models.GetUser(1)
	//fmt.Println(user)

	//Actualizar registro por id
	//user := models.GetUser(1)
	//fmt.Println(user)
	//user.Username = "Hugo"
	//user.Password = "Hugo123"
	//user.Email = "hugo@gmail.com"
	//user.Save()
	//fmt.Println(models.ListUsers())

	//Eliminar un registro
	//user := models.GetUser(3)
	//fmt.Println(user)
	//user.Delete()
	//fmt.Println(models.ListUsers())

	// Eliminar todos registro de la tabla
	db.TruncateTable("users")
	fmt.Println(models.ListUsers())

	//cerrar conexion
	db.Close()
}