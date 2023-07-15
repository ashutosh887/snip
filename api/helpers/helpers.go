package helpers

import (
	"os"
	"strings"
)

// enforcing http
func EnforceHTTP(url string) string {
	if url[:4] != "http" {
		return "http://" + url
	}

	return url
}

// removes all commonly found prefixes
// checks if the remaining string is the DOMAIN itself 🙂
func RemoveDomainError(url string) bool {
	if url == os.Getenv("DOMAIN") {
		return false
	}
	newURL := strings.Replace(url, "http://", "", 1)
	newURL = strings.Replace(newURL, "https://", "", 1)
	newURL = strings.Replace(newURL, "www.", "", 1)
	newURL = strings.Split(newURL, "/")[0]

	// return (newURL == os.Getenv("DOMAIN"))
	if newURL == os.Getenv("DOMAIN") {
		return false
	} else {
		return true
	}
}