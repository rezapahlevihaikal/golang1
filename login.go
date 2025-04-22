package main

import (
	"database/sql"
	"fmt"
)

func Login(username, password string) bool {
	var dbPassword string
	query := "SELECT password FROM users WHERE username = ?"

	err := DB.QueryRow(query, username).Scan(&dbPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("User tidak ditemukan")
			return false
		}
		fmt.Println("error query message: ", err)
		return false
	}

	if password == dbPassword {
		fmt.Println("Login berhasil")
		return true
	}

	fmt.Println("password salah.")
	return false
}
