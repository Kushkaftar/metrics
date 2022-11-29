package urlKit

import (
	"net/url"
	"strings"
)

func (s *Scheme) AddQueryParams(path string, query map[string]string) (string, error) {

	u, err := url.Parse(s.Url)
	if err != nil {
		return "", err
	}

	u.Path = path

	q := u.Query()

	if len(query) != 0 {

		for key, value := range query {
			if strings.Trim(key, " ") != "" {
				q.Add(key, value)
			}

		}
	}

	u.RawQuery = q.Encode()

	return u.String(), nil
}
