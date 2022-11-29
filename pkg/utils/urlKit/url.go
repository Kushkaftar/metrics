package urlKit

import "net/url"

type Scheme struct {
	Url string
}

func NewScheme(scheme, host string) *Scheme {

	u := url.URL{
		Scheme: scheme,
		Host:   host,
	}
	return &Scheme{
		Url: u.String(),
	}
}
