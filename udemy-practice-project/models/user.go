package models

import (
	"fmt"

	"example.com/rest-api/database"
	"example.com/rest-api/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

var users = []User{}

func (u User) Save() error {
	query := `INSERT INTO users (email, password) VALUES (?, ?)`
	stmt,err := database.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()

	hasedPassword, err :=  utils.HashPassword(u.Password)

	if err != nil {
		return err
	}
	

	result , err := stmt.Exec(u.Email, hasedPassword)
	
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	u.ID = id
	fmt.Printf("%v", u.ID)
	return err
}

func GetUsers() ([]User, error) {
	query := `SELECT * FROM users`
	rows, err := database.DB.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	users = []User{}
	for rows.Next() {
		var user User
		
		if err := rows.Scan(&user.ID, &user.Email, &user.Password); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, err
}