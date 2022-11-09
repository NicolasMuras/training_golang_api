package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/NicolasMuras/training_golang_api/songs"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ConfigDB() *pgxpool.Pool {
	conn, err := pgxpool.New(context.Background(), "postgresql://postgres:postgres@postgresdb:5432/postgres")

	if err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	return conn
}

func ValidateDBUserData(conn *pgxpool.Pool, username string, password string) string {
	sqlStatement := `
	SELECT * FROM accounts 
	WHERE username = $1 AND password = $2`

	var id int
	var user string
	var pass string

	err := conn.QueryRow(context.Background(), sqlStatement, username, password).Scan(&id, &user, &pass)
	if err != nil {
		// handle this error better than this
		fmt.Fprint(os.Stderr, "\n[ERROR] Read: ", err)
		return ""
	}

	return user
}

func RetrieveSongsFromDB(conn *pgxpool.Pool, filter string) []songs.Song {
	sqlStatement := `
	SELECT * FROM songs 
	WHERE LOWER(name) = LOWER($1) OR LOWER(artist) = LOWER($1) OR LOWER(album) = LOWER($1)`
	rows, err := conn.Query(context.Background(), sqlStatement, filter)
	if err != nil {
		// handle this error better than this
		fmt.Fprint(os.Stderr, "\n[ERROR] read: ", err)
	}

	song_array := []songs.Song{}

	for rows.Next() {
		var song songs.Song
		err := rows.Scan(
			&song.Id,
			&song.Name,
			&song.Artist,
			&song.Duration,
			&song.Album,
			&song.Artwork,
			&song.Price,
			&song.Origin,
		)
		if err != nil {
			// handle this error
			panic(err)
		}

		song_array = append(song_array, song)
	}
	// get any error encountered during iteration

	defer rows.Close()

	return song_array
}

func writeSongToDB(conn *pgxpool.Pool, id float64, name string, artist string, duration float64, album string, artwork string, price float64, origin string) {
	sqlStatement := `
	INSERT INTO songs (id, name, artist, duration, album, artwork, price, origin)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8);`

	rows := conn.QueryRow(
		context.Background(),
		sqlStatement,
		id,
		name,
		artist,
		duration,
		album,
		artwork,
		price,
		origin,
	).Scan()

	if rows.Error() != "no rows in result set" && rows != nil && !strings.Contains(rows.Error(), "duplicate key value violates unique constraint") {
		panic(rows)
	}
}

func WriteRowsToDB(conn *pgxpool.Pool, body []byte) {
	var result map[string]any
	json.Unmarshal([]byte(body), &result)
	for _, song := range result["results"].([]interface{}) {
		id := song.(map[string]interface{})["trackId"].(float64)
		name := song.(map[string]interface{})["trackName"].(string)
		artist := song.(map[string]interface{})["artistName"].(string)
		duration := song.(map[string]interface{})["trackTimeMillis"].(float64)
		album := song.(map[string]interface{})["collectionName"].(string)
		artwork := song.(map[string]interface{})["artworkUrl30"].(string)
		price := song.(map[string]interface{})["trackPrice"].(float64)
		origin := "apple"
		writeSongToDB(conn, id, name, artist, duration, album, artwork, price, origin)
	}
}
