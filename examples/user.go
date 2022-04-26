package main

import (
	"fmt"
	"github.com/knexguy101/BuffGo/buff/login"
	"github.com/knexguy101/BuffGo/buff/user"
	structs "github.com/knexguy101/BuffGo/models/client"
)

func main() {
	client := structs.NewHttpClient()

	data, err := login.Login()
	if err != nil {
		panic(err)
	}

	login.AddLoginDataToClient(data, client)

	info, err := user.GetUserInfo(client)
	if err != nil {
		panic(err)
	}

	fmt.Println(info.Data.ID, info.Data.Steamid)
}
