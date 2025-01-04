package services

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"html"

	"github.com/PuerkitoBio/goquery"
)

const (
	baseUrl       = "https://www.letras.mus.br"
	lyricCSSClass = ".lyric-original"
)

var fullLyrics string

func GetLyricsFromLetrasMus(artist string, song string) (string, error) {
	uri := fmt.Sprintf("%s/%s/%s", baseUrl, artist, song)
	resp, err := http.Get(uri)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Fatalf("LetrasMusServiceError: status code error: %d %s", resp.StatusCode, resp.Status)
	}

	htmlDocument, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", errors.New("LetrasMusServiceError: can't read html content")
	}

	htmlDocument.Find(lyricCSSClass).Each(func(i int, s *goquery.Selection) {

		s.Find("p").Each(func(i int, s *goquery.Selection) {
			verses, _ := s.Html()
			fullLyrics = fullLyrics + strings.ReplaceAll(verses, "<br/>", "\n")
		})

	})

	fullLyrics = html.UnescapeString(fullLyrics)

	return fullLyrics, nil
}
