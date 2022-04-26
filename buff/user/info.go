package user

import (
	structs "buffGo/models/client"
	"encoding/json"
	"io"
	"net/http"
)

func GetUserInfo(client *structs.HttpClient) (info *Info, err error) {

	var (
		res *http.Response
		resData []byte
	)

	res, err = client.Get(client.TS("https://buff.163.com/account/api/user/info?_="))
	if err != nil {
		return
	}
	defer res.Body.Close()

	resData, err = io.ReadAll(res.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(resData, &info)
	return
}
