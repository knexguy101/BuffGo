package purchase

import (
	structs "github.com/knexguy101/BuffGo/models/client"
	"encoding/json"
	"io"
	"net/http"
)

//Buy attempts to buy the given item given the args provided
func Buy(goodsId, price int, sellOrderId string, method PayMethod, client *structs.HttpClient) (purchase *PurchaseResponse, err error) {

	var (
		res *http.Response
		resData []byte
	)

	resData, err = json.Marshal(purchasePayload{
		AllowTradableCooldown: 0,
		CdkeyID: "",
		Game: "csgo",
		GoodsID: goodsId,
		PayMethod: method,
		Price: price,
		SellOrderID: sellOrderId,
		Token: "",
	})
	if err != nil {
		return
	}

	res, err = client.Post("https://buff.163.com/api/market/goods/buy", "application/json", string(resData))
	if err != nil {
		return
	}
	defer res.Body.Close()

	resData, err = io.ReadAll(res.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(resData, &purchase)
	return
}
