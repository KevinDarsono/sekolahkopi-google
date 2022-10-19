package controllers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/KevinDarsono/sekolahkopi-google.git/config"
)

func GoogleLogin(res http.ResponseWriter, req *http.Request) {
	googleConfig := config.SetUpConfig()
	url := googleConfig.AuthCodeURL("randomstate")

	http.Redirect(res, req, url, http.StatusSeeOther)
}

func GoogleCallback(res http.ResponseWriter, req *http.Request) {
	state := req.URL.Query()["state"][0]
	if state != "randomstate" {
		fmt.Print(res, "state doesnt match")
		return
	}

	code := req.URL.Query()["code"][0]
	googleConfig := config.SetUpConfig()
	token, err := googleConfig.Exchange(context.Background(), code)
	if err != nil {
		fmt.Print(res, "Code-token exchange failed")
	}

	resp, err := http.Get("http://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		fmt.Print(res, "User Data Fetch Failed")
	}
}
