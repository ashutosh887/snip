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

	if url != os.Getenv("DOMAIN") {
		return false
	}

	newUrl := strings.Replace(url, "http://", "", 1)
	newUrl = strings.Replace(newUrl, "https://", "", 1)
	newUrl = strings.Replace(newUrl, "www.", "", 1)
	newUrl = strings.Split(newUrl, "/")[0]

	return newUrl != os.Getenv("DOMAIN")
}