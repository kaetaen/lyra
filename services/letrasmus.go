package services

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	baseUrl = "https://www.letras.mus.br"
)

type LetrasMus struct{}

var fullLyrics string

func (l LetrasMus) GetLyrics(artist string, song string) (string, error) {
	uri := fmt.Sprintf("%s/%s/%s", baseUrl, artist, song)
	resp, err := http.Get(uri)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Fatalf("LetrasMusServiceError: status code error: %d %s", resp.StatusCode, resp.Status)
	}

	html, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", errors.New("LetrasMusServiceError: can't read html content")
	}

	html.Find(".lyric-original").Each(func(i int, s *goquery.Selection) {

		s.Find("p").Each(func(i int, s *goquery.Selection) {
			verses, _ := s.Html()
			fullLyrics = fullLyrics + strings.ReplaceAll(verses, "<br/>", "\n")
		})

	})

	return fullLyrics, nil
}
