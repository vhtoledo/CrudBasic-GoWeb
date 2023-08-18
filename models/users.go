package models

// Modelo usuario
type User struct {
	Id			int
	Username	string
	Password	string
	Email		string
}

const UserSchema string = `CREATE TABLE users (
	id INT(6) UNSINGNED AUTO_INCREMENT PRIMARY KEY,
	username VARCHAR(30) NOT NULL,
	password VARCHAR(50) NOT NULL,
	email VARCHAR(50) NOT NULL,
	create_data TIMESTAMP DEFAULT CURRENT_TIMESTAMP)`