package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type User struct {
	Id       int
	Username string
	Password string
	Token    string
}

var users = []User{
	{Id: 1, Username: "c137@onecause.com", Password: "#th@nH@rm#y#r!$100%D0p#", Token: ""}}

func main() {
	r := mux.NewRouter()

	r.Handle("/login", LoginHandler).Methods("POST")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4200"},
		AllowCredentials: true,
	})

	handler := c.Handler(r)

	http.ListenAndServe(":8080", handler)
}

var LoginHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Error occured during login. Error message: "+err.Error(), http.StatusUnauthorized)
	}

	var userData map[string]string
	json.Unmarshal(body, &userData)

	w.Header().Set("Content-Type", "application/json")

	hours, minutes, _ := time.Now().Clock()

	for _, u := range users {
		if userData["username"] == u.Username && userData["password"] == u.Password && userData["token"] == fmt.Sprintf("%02d%02d", hours, minutes) {
			json.NewEncoder(w).Encode(u)
		} else {
			http.Error(w, "Login unauthorized", http.StatusUnauthorized)
		}
	}
})
