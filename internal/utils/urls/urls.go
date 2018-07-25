package urls

import (
	"errors"
	"regexp"
)

const (
	URL_REGEX = "(https?)(://)([a-z0-9.:]+)([/.]*)"
)

var regex = regexp.MustCompile(URL_REGEX)

func IsValidBaseUrl(url string) bool {
	_, _, err := ExtractSchemeAndDomainWithPortFromUrl(url)
	return err == nil
}

func ExtractSchemeAndDomainWithPortFromUrl(url string) (*string, *string, error) {
	regexMatch := regex.FindStringSubmatch(url)
	if cap(regexMatch) == 0 {
		return nil, nil, errors.New("regex match failed")
	} else {
		// regexMatch[3] - Domain including port
		// regexMatch[1] - scheme http or https
		return &regexMatch[1], &regexMatch[3], nil
	}
}
