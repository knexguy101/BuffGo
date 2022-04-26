package main

import (
	"fmt"
	"github.com/knexguy101/BuffGo/buff/login"
	"github.com/knexguy101/BuffGo/buff/search"
	structs "github.com/knexguy101/BuffGo/models/client"
	"math/rand"
)

func main() {
	client := structs.NewHttpClient()

	data, err := login.Login()
	if err != nil {
		panic(err)
	}

	login.AddLoginDataToClient(data, client)

	marketResults, err := search.Market("ak slate", 1, client)
	if err != nil {
		panic(err)
	}

	fmt.Println("Search results")
	for _, v := range marketResults.Data.Items {
		fmt.Println(v.ID)
	}

	randomItem := marketResults.Data.Items[rand.Intn(len(marketResults.Data.Items))]

	resItems, err := search.Item(search.NewSearchFilter(fmt.Sprintf("%d", randomItem.ID), search.DEFAULT, 1, true), client)
	if err != nil {
		panic(err)
	}

	fmt.Println("Item results")
	for _, v := range resItems {
		fmt.Println(v.ID, v.Price, v.AssetInfo.Paintwear)
	}
}
