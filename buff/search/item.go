package search

import (
	structs "github.com/knexguy101/BuffGo/models/client"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

//Item searches a specific item with a specified ItemFilter
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
