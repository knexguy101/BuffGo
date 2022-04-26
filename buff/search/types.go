package search

type marketSearch struct {
	Code string `json:"code"`
	Data struct {
		Suggestions []Suggestion `json:"suggestions"`
	} `json:"data"`
	Msg interface{} `json:"msg"`
}

type Suggestion struct {
	GoodsIds string `json:"goods_ids"`
	Option   string `json:"option"`
}

type itemSearch struct {
	Code string `json:"code"`
	Data struct {
		FopStr     string `json:"fop_str"`
		Items []SearchItem `json:"items"`
		PageNum            int `json:"page_num"`
		PageSize           int `json:"page_size"`
		PreviewScreenshots struct {
			BgImg   string `json:"bg_img"`
			Selling string `json:"selling"`
		} `json:"preview_screenshots"`
		ShowPayMethodIcon bool   `json:"show_pay_method_icon"`
		SortBy            string `json:"sort_by"`
		SrcURLBackground  string `json:"src_url_background"`
		TotalCount        int    `json:"total_count"`
		TotalPage         int    `json:"total_page"`
	} `json:"data"`
	Msg interface{} `json:"msg"`
}

type SearchItem struct {
	AllowBargain bool `json:"allow_bargain"`
	Appid        int  `json:"appid"`
	AssetInfo    struct {
		ActionLink          string `json:"action_link"`
		Appid               int    `json:"appid"`
		Assetid             string `json:"assetid"`
		Classid             string `json:"classid"`
		Contextid           int    `json:"contextid"`
		GoodsID             int    `json:"goods_id"`
		HasTradableCooldown bool   `json:"has_tradable_cooldown"`
		Info                struct {
			Fraudwarnings     interface{}   `json:"fraudwarnings"`
			IconURL           string        `json:"icon_url"`
			InspectEnSize     string        `json:"inspect_en_size"`
			InspectEnURL      string        `json:"inspect_en_url"`
			InspectMobileSize string        `json:"inspect_mobile_size"`
			InspectMobileURL  string        `json:"inspect_mobile_url"`
			InspectSize       string        `json:"inspect_size"`
			InspectStartAt    string        `json:"inspect_start_at"`
			InspectState      int           `json:"inspect_state"`
			InspectTrnSize    string        `json:"inspect_trn_size"`
			InspectTrnURL     string        `json:"inspect_trn_url"`
			InspectURL        string        `json:"inspect_url"`
			InspectVersion    int           `json:"inspect_version"`
			InspectedAt       string        `json:"inspected_at"`
			OriginalIconURL   string        `json:"original_icon_url"`
			Paintindex        int           `json:"paintindex"`
			Paintseed         int           `json:"paintseed"`
			Stickers          []Sticker `json:"stickers"`
			TournamentTags    []interface{} `json:"tournament_tags"`
		} `json:"info"`
		Instanceid           string      `json:"instanceid"`
		Paintwear            string      `json:"paintwear"`
		TradableCooldownText string      `json:"tradable_cooldown_text"`
		TradableUnfrozenTime interface{} `json:"tradable_unfrozen_time"`
	} `json:"asset_info"`
	BackgroundImageURL    string      `json:"background_image_url"`
	Bookmarked            bool        `json:"bookmarked"`
	CanBargain            bool        `json:"can_bargain"`
	CanUseInspectTrnURL   bool        `json:"can_use_inspect_trn_url"`
	CannotBargainReason   string      `json:"cannot_bargain_reason"`
	CouponInfos           interface{} `json:"coupon_infos"`
	CreatedAt             int         `json:"created_at"`
	Description           string      `json:"description"`
	Featured              int         `json:"featured"`
	Fee                   string      `json:"fee"`
	Game                  string      `json:"game"`
	GoodsID               int         `json:"goods_id"`
	ID                    string      `json:"id"`
	ImgSrc                string      `json:"img_src"`
	Income                string      `json:"income"`
	LowestBargainPrice    string      `json:"lowest_bargain_price"`
	Mode                  int         `json:"mode"`
	Price                 string      `json:"price"`
	RecentAverageDuration float64     `json:"recent_average_duration"`
	RecentDeliverRate     float64     `json:"recent_deliver_rate"`
	State                 int         `json:"state"`
	SupportedPayMethods   []int       `json:"supported_pay_methods"`
	TradableCooldown      interface{} `json:"tradable_cooldown"`
	UpdatedAt             int         `json:"updated_at"`
	UserID                string      `json:"user_id"`
}

type Sticker struct {
	ImgUrl string `json:"img_url"`
	Name string `json:"name"`
	Slot int `json:"slot"`
	StickerId int `json:"sticker_id"`
	Wear int `json:"wear"`
}