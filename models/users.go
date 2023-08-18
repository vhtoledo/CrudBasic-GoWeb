package models

import "gomysql/db"

// Modelo usuario
type User struct {
	Id       int64
	Username string
	Password string
	Email    string
}

type Users []User

const UserSchema string = `CREATE TABLE users(
	id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	username VARCHAR(30) NOT NULL,
	password VARCHAR(64) NOT NULL,
	email VARCHAR(50),
	create_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP)`

//Construir usuario
func NewUser(username, password, email string) *User {
	user := &User{Username: username, Password: password, Email: email}
	return user
}

//Crear usuario e insertar
func CreateUser(username, password, email string) *User {
	user := NewUser(username, password, email)
	user.insert()
	return user
}

//Insertar Registro
func (user *User) insert() {
	sql := "INSERT users SET username=?, password=?,email=?"
	result, _ := db.Exec(sql, user.Username, user.Password, user.Email)
	user.Id, _ = result.LastInsertId()
}