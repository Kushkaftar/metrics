package urlKit

import "net/url"

func (s *Scheme) GetUrl(path string) (string, error) {
	u, err := url.Parse(s.Url)
	if err != nil {
		return "", err
	}

	u.Path = path

	return u.String(), nil
}
