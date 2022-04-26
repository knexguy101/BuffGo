package structs

import (
	"net/http"
	"strings"
)

func (c *HttpClient) Get(u string) (*http.Response, error) {
	req, _ := http.NewRequest("GET", u, nil)
	req.Header = map[string][]string {
		"Accept": {"*/*"},
		"Accept-Language": {"en-US,en;q=0.9"},
		"User-Agent": {c.UserAgent},
		"X-Requested-With": {"XMLHttpRequest"},
	}
	return c.Send(req, true, true)
}

func (c *HttpClient) Post(u, contentType, body string) (*http.Response, error) {
	req, _ := http.NewRequest("POST", u, strings.NewReader(body))
	req.Header = map[string][]string {
		"Accept": {"*/*"},
		"Accept-Language": {"en-US,en;q=0.9"},
		"User-Agent": {c.UserAgent},
		"Content-Type": {contentType},
		"X-Requested-With": {"XMLHttpRequest"},
	}
	return c.Send(req, true, true)
}