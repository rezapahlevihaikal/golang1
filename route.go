package main

import (
	"fmt"
	"net/http"
)

func Route() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "login.html")
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			username := r.FormValue("username")
			password := r.FormValue("password")
			if Login(username, password) {
				http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
			} else {
				fmt.Println(w, "Login gagal")
			}
		}
	})

	http.HandleFunc("/dashboard", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "dashboard.html")
	})
}
