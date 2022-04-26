package structs

import (
	"fmt"
	"net/http"
	"strings"
)

type HttpClient struct {
	Tr      *http.Transport
	Cookies map[string]*http.Cookie
	UserAgent string
}

func NewHttpClient() *HttpClient {
	return &HttpClient{
		Tr: &http.Transport{
			DisableKeepAlives: true,
			ForceAttemptHTTP2: true,
		},
		Cookies: make(map[string]*http.Cookie),
		UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.127 Safari/537.36",
	}
}

func (c *HttpClient) ClearCookies() {
	c.Cookies = make(map[string]*http.Cookie)
}

func (c *HttpClient) AddCookie(name, value, domain, path string) {
	c.Cookies[name] = &http.Cookie{
		Name:  name,
		Value: value,
	}
}

func (c *HttpClient) RemoveCookie(name string) {
	delete(c.Cookies, name)
}

func (c *HttpClient) GetCookie(name string) (string, bool) {
	val, ok := c.Cookies[name]
	if !ok {
		return "", ok
	}
	return val.Value, ok
}

func (c *HttpClient) Send(request *http.Request, useCookies, followRedirect bool) (*http.Response, error) {

	customCookie := true
	_, cookie := request.Header["cookie"]
	if useCookies && !cookie {
		request.Header.Add("cookie", c.GetCookieString())
		customCookie = false
	} else if !useCookies {
		request.Header.Add("cookie", "")
		customCookie = false
	}
	reqHeader := request.Header
	res, err := c.Tr.RoundTrip(request)
	if err != nil {
		return nil, err
	}
	if useCookies {
		c.SaveCookies(res)
	}

	storedLocation := ""
	if followRedirect {
		for res.StatusCode >= 300 && res.StatusCode < 400 {
			location := res.Header.Get("Location")
			if location == "" {
				break
			} else if !strings.HasPrefix(location, "http") {
				location = "https://" + request.Host + location
			}
			request, _ = http.NewRequest("GET", location, nil)

			request.Close = true

			location = storedLocation
			request.Header = reqHeader

			if !customCookie {
				request.Header["cookie"] = []string{c.GetCookieString()}
			}
			res, err = c.Tr.RoundTrip(request)
			if err != nil {
				return nil, err
			}
			if useCookies {
				c.SaveCookies(res)
			}

		}
	}
	return res, nil
}

func (c *HttpClient) SaveCookies(res *http.Response) {
	for _, v := range res.Cookies() {
		c.Cookies[v.Name] = v
	}
}

func (c *HttpClient) GetCookieString() string {
	var str []string
	for k, v := range c.Cookies {
		str = append(str, fmt.Sprintf("%s=%s", k, v.Value))
	}
	return strings.Join(str, "; ")
}