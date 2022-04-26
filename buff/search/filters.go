package search

import (
	"fmt"
	"strings"
	"time"
)

const (
	DEFAULT SortBy = "default"
	NEWEST SortBy = "created.desc"
	PRICEASC SortBy = "price.asc"
	PRICEDESC SortBy = "price.desc"
	PAINTWEARASC SortBy = "paintwear.asc"
	PAINTWEARDESC SortBy = "paintwear.desc"
	HOTDESC SortBy = "heat.desc"
)

type SortBy string

type ItemFilter struct {
	GoodsId string
	PageNum int
	SortBy SortBy
	AllowTradeCooldown bool
}

//NewSearchFilter Creates a new search filter
func NewSearchFilter(goodsId string, sortBy SortBy, pageNum int, allowTradeCooldown bool) *ItemFilter {
	return &ItemFilter{
		GoodsId: goodsId,
		SortBy: sortBy,
		PageNum: pageNum,
		AllowTradeCooldown: allowTradeCooldown,
	}
}

//ToQuery pushes the ItemFilter into a query string
func (f *ItemFilter) ToQuery() string {
	trade := 1
	if !f.AllowTradeCooldown {
		trade = 0
	}
	var items = []string {
		fmt.Sprintf("goods_id=%s", f.GoodsId),
		fmt.Sprintf("page_num=%d", f.PageNum),
		fmt.Sprintf("sort_by=%s", f.SortBy),
		fmt.Sprintf("allow_tradable_cooldown=%d", trade),
		fmt.Sprintf("_=%d", time.Now().UnixMilli()),
	}
	return strings.Join(items, "&")
}
