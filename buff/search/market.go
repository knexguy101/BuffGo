package search

import (
	structs "github.com/knexguy101/BuffGo/models/client"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

//Market searches the market for items given a query string
func Market(search string, client *structs.HttpClient) (suggestions []Suggestion, err error) {

	var (
		res *http.Response
		tempRes *marketSearch
		resData []byte
	)

	res, err = client.Get(fmt.Sprintf("https://buff.163.com/api/market/search/suggest?text=%s&game=csgo", url.QueryEscape(search)))
	if err != nil {
		return
	}
	defer res.Body.Close()

	resData, err = io.ReadAll(res.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(resData, &tempRes)
	if err != nil {
		return
	}

	suggestions = tempRes.Data.Suggestions

	return
}
