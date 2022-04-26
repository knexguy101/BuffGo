package purchase

import (
	structs "github.com/knexguy101/BuffGo/models/client"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

//Status gets the status of an orderId
func Status(orderId string, client *structs.HttpClient) (purchase *PurchaseStatus, err error) {

	var (
		res *http.Response
		resData []byte
	)

	res, err = client.Get(client.TS(fmt.Sprintf("https://buff.163.com/api/market/bill_order/batch/info?bill_orders=%s&_=", orderId)))
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