package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/NicolasMuras/training_golang_api/auth"
	"github.com/NicolasMuras/training_golang_api/scraping"
	"github.com/NicolasMuras/training_golang_api/utils"
)

func handleSearch(writer http.ResponseWriter, request *http.Request) {
	_, err := auth.VerifyJWT(writer, request)

	if err != nil {
		log.Fatalln("[ERROR] Extract claims error:", err)
	} else {
		writer.Header().Set("Token", "%v")
		type_ := "application/json"
		writer.Header().Set("Content-Type", type_)

		conn := utils.ConfigDB()
		filter := request.URL.Query().Get("filter")
		body := scraping.SearchSongs(filter, conn)
		utils.WriteRowsToDB(conn, body)
		song_array := utils.RetrieveSongsFromDB(conn, filter)

		err = json.NewEncoder(writer).Encode(song_array)
		if err != nil {
			return
		}
	}
}

// Call to db to validate user credentials. If user/pass is correct: retrieve a valid token for the user.
func handleLogin(writer http.ResponseWriter, request *http.Request) {
	username := request.URL.Query().Get("username")
	password := request.URL.Query().Get("password")
	conn := utils.ConfigDB()
	validation := utils.ValidateDBUserData(conn, username, password)

	if validation != "" {
		token := auth.AuthPage()
		json.NewEncoder(writer).Encode(token)
	} else {
		json.NewEncoder(writer).Encode("[ERROR] Invalid username or password.")
	}
}

func main() {
	// http://127.0.0.1:3000/login?username=nico&password=123456
	http.HandleFunc("/login", handleLogin)
	// http://127.0.0.1:3000/search?filter=maiden
	http.HandleFunc("/search", handleSearch)

	fmt.Println("[+] Server started on port:", 8080)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
