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

	marketResults, err := search.Market("ak slate", client)
	if err != nil {
		panic(err)
	}

	fmt.Println("Search results")
	for _, v := range marketResults {
		fmt.Println(v.GoodsIds)
	}

	randomItem := marketResults[rand.Intn(len(marketResults))]

	resItems, err := search.Item(search.NewSearchFilter(randomItem.GoodsIds, search.DEFAULT, 1, true), client)
	if err != nil {
		panic(err)
	}

	fmt.Println("Item results")
	for _, v := range resItems {
		fmt.Println(v.ID, v.Price, v.AssetInfo.Paintwear)
	}
}
