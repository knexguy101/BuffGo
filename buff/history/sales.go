package history

import (
	structs "buffGo/models/client"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

//Sales gives back the recent sales of an item
func Sales(goodsId string, client *structs.HttpClient) (history *SaleHistory, err error) {

	var (
		res *http.Response
		resData []byte
	)

	res, err = client.Get(client.TS(fmt.Sprintf("https://buff.163.com/api/market/goods/bill_order?game=csgo&goods_id=%s&_=", goodsId)))
	if err != nil {
		return
	}
	defer res.Body.Close()

	resData, err = io.ReadAll(res.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(resData, &history)

	return
}


