package config

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func SetUpConfig() *oauth2.Config {
	conf := &oauth2.Config{
		ClientID:     "933044355937-f7qdlfpvg3u3b40sroih6dkoqjv77o6n.apps.googleusercontent.com",
		ClientSecret: "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",
		RedirectURL:  "http://localhost:8080/callback",
		Scopes: []string{
			"http://www.googleapis.com/auth/userinfo.email",
			"http://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}
	return conf
}
