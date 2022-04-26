package history

import (
	structs "buffGo/models/client"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

//Prices gives back the price history of an item
func Prices(goodsId, currency string, client *structs.HttpClient) (history *PriceHistory, err error) {

	var (
		res *http.Response
		resData []byte
	)

	res, err = client.Get(client.TS(fmt.Sprintf("https://buff.163.com/api/market/goods/price_history/buff?game=csgo&goods_id=%s&currency=%s&_=", goodsId, currency)))
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

