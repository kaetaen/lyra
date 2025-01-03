package services

type LyricService interface {
	GetLyrics(artist string, song string) (string, error)
}
