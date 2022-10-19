package main

import (
	"net/http"

	"github.com/KevinDarsono/sekolahkopi-google.git/controllers"
)

func main() {
	http.HandleFunc("/google/login", controllers.GoogleLogin)
	//http.HandleFunc("/google/callback", controllers.GoogleCallback)

	http.ListenAndServe(":8080", nil)
}
