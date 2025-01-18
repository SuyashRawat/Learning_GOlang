package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	googleOauthConfig *oauth2.Config    //stores config for oauth2
	oauthStateString  = "pseudo-random" //prevents CSRF attacks during oauth flow, randomized for each session
)

func init() {
	//load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	//setup googleOauthConfig
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/callback",
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"}, //only user's email is accessible
		Endpoint:     google.Endpoint,                                            //predefined endpoint
	}
}

func main() {
	http.HandleFunc("/", handleMain)                   //main page
	http.HandleFunc("/login", handleGoogleLogin)       //initiates google oauth login
	http.HandleFunc("/callback", handleGoogleCallback) //handles google's response after authentication
	fmt.Println(http.ListenAndServe(":8080", nil))     //starts server at 8080
}

func handleMain(w http.ResponseWriter, r *http.Request) { //starts the google login process by visiting /login
	var htmlIndex = `<html><body><a href="/login">Google Log In</a></body></html>`
	fmt.Fprint(w, htmlIndex)
}

func handleGoogleLogin(w http.ResponseWriter, r *http.Request) { //generates URL for google's oauth consent page using AuthCodeURL
	url := googleOauthConfig.AuthCodeURL(oauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect) //redirects user to google's consent page
}

func handleGoogleCallback(w http.ResponseWriter, r *http.Request) {
	content, err := getUserInfo(r.FormValue("state"), r.FormValue("code")) //state->validates response matches the request, code->authcode used to fetch access token
	if err != nil {
		fmt.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	fmt.Fprintf(w, "Content: %s\n", content) //displays user details
}

func getUserInfo(state string, code string) ([]byte, error) {
	if state != oauthStateString { //validates the state to ensure it matches oauthStateString
		return nil, fmt.Errorf("invalid oauth state")
	}

	token, err := googleOauthConfig.Exchange(context.Background(), code) //exchanges the code for an access token
	if err != nil {
		return nil, fmt.Errorf("code exchange failed: %s", err.Error())
	}

	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken) //fetch user details from google's api
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}

	defer response.Body.Close()
	contents, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed reading response body: %s", err.Error())
	}

	return contents, nil
}
