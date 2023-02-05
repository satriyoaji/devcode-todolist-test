package repositories

import (
	"database/sql"
	"fmt"
	"github.com/satriyoaji/todolist-challenge-go/database"
	"github.com/satriyoaji/todolist-challenge-go/helpers"
)

type User struct {
	Id    int    `json:"id"`
	Email string `json:"email" validate:"required,email,unique"`
}

func CheckLogin(email, password string) (bool, error) {
	var obj User
	var pwd string

	//v := validator.New()
	//credentialStruct := User{
	//	Email: email,
	//}
	//// validation input
	//errorLogin := v.Struct(credentialStruct)
	//if errorLogin != nil {
	//	return false, errorLogin
	//}

	con := database.GetConnection()

	sqlStatement := "SELECT * FROM users WHERE email = ?"
	err := con.QueryRow(sqlStatement, email).Scan(
		&obj.Id, &obj.Email, &pwd,
	)

	if err == sql.ErrNoRows {
		fmt.Println("Email not found")
		return false, err
	}
	if err != nil {
		fmt.Println("Query error")
		return false, err
	}

	match, err := helpers.CheckPasswordHash(password, pwd)
	if !match {
		fmt.Println("Hash and password doesn't match.")
		return false, err
	}

	return true, nil
}
