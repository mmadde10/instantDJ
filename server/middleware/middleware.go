package middleware

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/zmb3/spotify"
	"golang.org/x/oauth2"
)

// // DB connection string
// // const connectionString = "mongodb://localhost:27017"
// const connectionString = "Connection String"

// // Database Name
// const dbName = "test"

// // Collection name
// const collName = "todolist"

// // collection object/instance
// var collection *mongo.Collection

const redirectURI = "http://localhost:8080/api/callback"

var clientID = os.Getenv("clientID")
var secretKey = os.Getenv("secretKey")

type userLogin struct {
	ID           string
	Name         string
	AccessToken  string
	RefreshToken string
	Email        string
}

// GenerateRandomBytes returns securely generated random bytes.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

// GenerateRandomString returns a URL-safe, base64 encoded
// securely generated random string.
func GenerateRandomString(s int) (string, error) {
	b, err := GenerateRandomBytes(s)
	return base64.URLEncoding.EncodeToString(b), err
}

// TODO: Move random generator into a util package

var stateToken, err = GenerateRandomString(32)

var (
	auth  = spotify.NewAuthenticator(redirectURI, spotify.ScopeUserReadEmail)
	ch    = make(chan *spotify.Client)
	state = stateToken
	token = make(chan *oauth2.Token)
)

// AuthenticateUser Auth user
func AuthenticateUser(w http.ResponseWriter, r *http.Request) {
	auth.SetAuthInfo(clientID, secretKey)
	url := auth.AuthURL(state)
	fmt.Println("Please log in to Spotify by visiting the following page in your browser:", url)
	// wait for auth to complete
	client := <-ch
	tok := <-token

	// use the client to make calls that require authorization
	user, err := client.CurrentUser()
	if err != nil {
		log.Fatal(err)
	}

	AuthInfo := userLogin{
		ID:           user.ID,
		Name:         user.DisplayName,
		AccessToken:  tok.AccessToken,
		RefreshToken: tok.RefreshToken,
		Email:        user.Email,
	}

	fmt.Println("\n You are logged in as:", user.ID)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(AuthInfo)
}

// CompleteAuth Auth user
func CompleteAuth(w http.ResponseWriter, r *http.Request) {
	tok, err := auth.Token(state, r)
	if err != nil {
		http.Error(w, "Couldn't get token", http.StatusForbidden)
		log.Fatal(err)
	}
	if st := r.FormValue("state"); st != state {
		http.NotFound(w, r)
		log.Fatalf("State mismatch: %s != %s\n", st, state)
	}
	// use the token to get an authenticated client
	client := auth.NewClient(tok)
	fmt.Fprintf(w, "Login Completed!")
	ch <- &client
	token <- tok
}

// BEGIN TRACK MIDDLEWARE

//GetTrack Checks for token, then gets track by ID
func GetTrack(w http.ResponseWriter, r *http.Request) {
	client := &http.Client{}
	tok := r.Header["Authorization"][0]

	params := mux.Vars(r)
	id := params["id"]

	url := "https://api.spotify.com/v1/tracks/" + id

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", tok)
	resp, err := client.Do(req)

	if resp.StatusCode == 200 {

		defer resp.Body.Close()
		var result map[string]interface{}
		json.NewDecoder(resp.Body).Decode(&result)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("hee: ", result["name"])

		// TrackResult := Track{
		// 	id: result["id"],
		// }

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(result)
	}
	if resp.StatusCode == 401 {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("401 - Unauthorized"))
	}
	if resp.StatusCode == 400 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Bad Request"))
	}
}
