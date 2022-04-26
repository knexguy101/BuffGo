package search

import (
	"encoding/json"
	"fmt"
	structs "github.com/knexguy101/BuffGo/models/client"
	"io"
	"net/http"
	"net/url"
)

//Market searches the market for items given a query string
func Market(search string, pageNum int, client *structs.HttpClient) (marketSearch *MarketSearch, err error) {

	var (
		res *http.Response
		resData []byte
	)

	res, err = client.Get(client.TS(fmt.Sprintf("https://buff.163.com/api/market/goods?game=csgo&page_num=%d&search=%s&use_suggestion=0&trigger=search_input&_=", pageNum, url.QueryEscape(search))))
	if err != nil {
		return
	}
	defer res.Body.Close()

	resData, err = io.ReadAll(res.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(resData, &marketSearch)
	if err != nil {
		return
	}

	return
}
