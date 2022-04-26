package search

import (
	structs "buffGo/models/client"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Item(filter *ItemFilter, client *structs.HttpClient) (items []SearchItem, err error) {
	var (
		res *http.Response
		tempRes *itemSearch
		resData []byte
	)

	res, err = client.Get(fmt.Sprintf("https://buff.163.com/api/market/goods/sell_order?game=csgo&%s", filter.ToQuery()))
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

	items = tempRes.Data.Items

	return
}
