package scraping

import (
	"strings"

	"github.com/NicolasMuras/training_golang_api/utils"

	"github.com/jackc/pgx/v5/pgxpool"
)

func buildITunesUrl(url string, filter string) string {
	filter = strings.ReplaceAll(filter, " ", "+")
	return "https://itunes.apple.com/search?term=" + filter + "&entity=musicTrack&limit=10"
}

func SearchSongs(filter string, conn *pgxpool.Pool) []byte {
	url := buildITunesUrl("https://itunes.apple.com/search?", filter)
	body := utils.ApiRequest(url, "")
	return body
}
