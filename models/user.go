package models

import (
	"errors"

	"example.com/rest-api/db"
	"example.com/rest-api/utils"
)

type User struct {
	ID       int64
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type UserExample struct {
	Email    string `json:"email" example:"someone@something.com"`
	Password string `json:"password" example:"somewhere"`
}

type UserCreated struct {
	Message string `json:"message" example:"User created successfully"`
}
type UserCreatedUnsuccessful struct {
	Message string `json:"message" example:"Could not save user."`
}

type UserLogin struct {
	Message string `json:"message" example:"Login successful!"`
	Token string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3QyQGV4YW1wbGUuY29tIiwiZXhwIjoxNzA3MjUzMTU4LCJ1c2VySWQiOjJ9.IuP3ati5IfjYSUnifkM4Ri9htCtwWPBepddSF6MTNUI"`
}
type UserLoginUnsuccessful struct {
	Message string `json:"message" example:"Could not authenticate user."`
}

func (u User) Save() error {
	query := "INSERT INTO users(email, password) VALUES (?, ?)"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashedPassword)

	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()

	u.ID = userId
	return err
}

func (u *User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err := row.Scan(&u.ID, &retrievedPassword)

	if err != nil {
		return errors.New("credentials invalid")
	}

	passwordIsValid := utils.CheckPasswordHash(u.Password, retrievedPassword)

	if !passwordIsValid {
		return errors.New("credentials invalid")
	}

	return nil
}
