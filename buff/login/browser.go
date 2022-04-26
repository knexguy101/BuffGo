package login

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
)

//Login opens a browser and collects cookies needed to automate the site.
func Login() (*LoginData, error) {
	u := launcher.New().Headless(false).MustLaunch()
	b := rod.New().ControlURL(u).MustConnect()
	defer b.Close()

	p := b.MustPage("https://buff.163.com/")
	defer p.Close()

	var cookies []*proto.NetworkCookie
	sessionFound := false
	for !sessionFound {
		cookies = p.MustCookies()
		for _, v := range cookies {
			if v.Name == "session" {
				sessionFound = true
				break
			}
		}
	}

	return &LoginData{
		Cookies: cookies,
	}, nil
}
