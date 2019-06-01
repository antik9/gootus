package shortener

import (
	"math/rand"
)

const (
	symbols    = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	linkLength = 6
	prefix     = "antik.com/"
)

type Shortener interface {
	Shorten(url string) string
	Resolve(url string) string
}

type LinkShortener struct {
	linksMap        map[string]string
	linksReverseMap map[string]string
}

func NewLinkShortener() LinkShortener {
	return LinkShortener{map[string]string{}, map[string]string{}}
}

func randomString(length int) string {
	bytes := make([]byte, length)
	for i := 0; i < length; i++ {
		bytes[i] = symbols[rand.Intn(len(symbols))]
	}
	return string(bytes)
}

func (shortener LinkShortener) Shorten(url string) string {
	if shortenedLink, ok := shortener.linksMap[url]; ok {
		return prefix + shortenedLink
	} else {
		for {
			shortenedLink := randomString(linkLength)
			if _, ok := shortener.linksReverseMap[shortenedLink]; !ok {
				shortener.linksMap[url] = shortenedLink
				shortener.linksReverseMap[shortenedLink] = url
				return prefix + shortenedLink
			}
		}
	}
}

func (shortener LinkShortener) Resolve(url string) string {
	if len(url) != len(prefix)+linkLength {
		return ""
	} else {
		hash := url[len(prefix):]
		if fullUrl, ok := shortener.linksReverseMap[hash]; ok {
			return fullUrl
		} else {
			return ""
		}
	}
}
