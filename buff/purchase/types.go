package purchase

type purchasePayload struct {
	Game                  string `json:"game"`
	GoodsID               int    `json:"goods_id"`
	SellOrderID           string `json:"sell_order_id"`
	Price                 int    `json:"price"`
	PayMethod             PayMethod    `json:"pay_method"`
	AllowTradableCooldown int    `json:"allow_tradable_cooldown"`
	Token                 string `json:"token"`
	CdkeyID               string `json:"cdkey_id"`
}

type PurchaseResponse struct {
	Code string `json:"code"`
	Data struct {
		Appid     int `json:"appid"`
		AssetInfo struct {
			ActionLink          string `json:"action_link"`
			Appid               int    `json:"appid"`
			Assetid             string `json:"assetid"`
			Classid             string `json:"classid"`
			Contextid           int    `json:"contextid"`
			GoodsID             int    `json:"goods_id"`
			HasTradableCooldown bool   `json:"has_tradable_cooldown"`
			Info                struct {
				Fraudwarnings     string        `json:"fraudwarnings"`
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
				Paintindex        int           `json:"paintindex"`
				Paintseed         int           `json:"paintseed"`
				Stickers          []interface{} `json:"stickers"`
				TournamentTags    []interface{} `json:"tournament_tags"`
			} `json:"info"`
			Instanceid           string `json:"instanceid"`
			Paintwear            string `json:"paintwear"`
			TradableCooldownText string `json:"tradable_cooldown_text"`
			TradableUnfrozenTime int    `json:"tradable_unfrozen_time"`
		} `json:"asset_info"`
		BundleInfo struct {
		} `json:"bundle_info"`
		BuyerCancelTimeout       interface{} `json:"buyer_cancel_timeout"`
		BuyerCookieInvalid       bool        `json:"buyer_cookie_invalid"`
		BuyerID                  string      `json:"buyer_id"`
		BuyerPayTime             int         `json:"buyer_pay_time"`
		BuyerSendOfferTimeout    int         `json:"buyer_send_offer_timeout"`
		CanReplaceAsset          bool        `json:"can_replace_asset"`
		CouponInfo               interface{} `json:"coupon_info"`
		CouponInfos              interface{} `json:"coupon_infos"`
		CreatedAt                int         `json:"created_at"`
		DeliverExpireTimeout     int         `json:"deliver_expire_timeout"`
		ErrorText                interface{} `json:"error_text"`
		FailConfirm              interface{} `json:"fail_confirm"`
		Fee                      string      `json:"fee"`
		Game                     string      `json:"game"`
		GoodsID                  int         `json:"goods_id"`
		HasBargain               bool        `json:"has_bargain"`
		HasSentOffer             bool        `json:"has_sent_offer"`
		ID                       string      `json:"id"`
		Income                   string      `json:"income"`
		IsSellerAskedToSendOffer bool        `json:"is_seller_asked_to_send_offer"`
		Mode                     int         `json:"mode"`
		OriginalPrice            interface{} `json:"original_price"`
		PayExpireTimeout         int         `json:"pay_expire_timeout"`
		PayMethod                int         `json:"pay_method"`
		PayMethodText            string      `json:"pay_method_text"`
		Price                    string      `json:"price"`
		PriceWithPayFee          string      `json:"price_with_pay_fee"`
		Progress                 int         `json:"progress"`
		ReceiveExpireTimeout     int         `json:"receive_expire_timeout"`
		SellOrderID              interface{} `json:"sell_order_id"`
		SellerCanCancel          bool        `json:"seller_can_cancel"`
		SellerCookieInvalid      bool        `json:"seller_cookie_invalid"`
		SellerID                 string      `json:"seller_id"`
		State                    string      `json:"state"`
		StateText                string      `json:"state_text"`
		TradeOfferTraceURL       interface{} `json:"trade_offer_trace_url"`
		TradeOfferURL            interface{} `json:"trade_offer_url"`
		Tradeofferid             interface{} `json:"tradeofferid"`
		TransactTime             interface{} `json:"transact_time"`
		Type                     int         `json:"type"`
		UpdatedAt                int         `json:"updated_at"`
	} `json:"data"`
	Msg interface{} `json:"msg"`
}

type PurchaseStatus struct {
	Code string `json:"code"`
	Data struct {
		Items []struct {
			Appid     int `json:"appid"`
			AssetInfo struct {
				ActionLink          string `json:"action_link"`
				Appid               int    `json:"appid"`
				Assetid             string `json:"assetid"`
				Classid             string `json:"classid"`
				Contextid           int    `json:"contextid"`
				GoodsID             int    `json:"goods_id"`
				HasTradableCooldown bool   `json:"has_tradable_cooldown"`
				Info                struct {
					Fraudwarnings     string        `json:"fraudwarnings"`
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
					Stickers          []interface{} `json:"stickers"`
					TournamentTags    []interface{} `json:"tournament_tags"`
				} `json:"info"`
				Instanceid           string `json:"instanceid"`
				Paintwear            string `json:"paintwear"`
				TradableCooldownText string `json:"tradable_cooldown_text"`
				TradableUnfrozenTime int    `json:"tradable_unfrozen_time"`
			} `json:"asset_info"`
			BundleInfo struct {
			} `json:"bundle_info"`
			BuyerCancelTimeout       interface{} `json:"buyer_cancel_timeout"`
			BuyerCookieInvalid       bool        `json:"buyer_cookie_invalid"`
			BuyerID                  string      `json:"buyer_id"`
			BuyerPayTime             int         `json:"buyer_pay_time"`
			BuyerSendOfferTimeout    int         `json:"buyer_send_offer_timeout"`
			CanReplaceAsset          bool        `json:"can_replace_asset"`
			CouponInfo               interface{} `json:"coupon_info"`
			CouponInfos              interface{} `json:"coupon_infos"`
			CreatedAt                int         `json:"created_at"`
			DeliverExpireTimeout     int         `json:"deliver_expire_timeout"`
			ErrorText                interface{} `json:"error_text"`
			FailConfirm              interface{} `json:"fail_confirm"`
			Fee                      string      `json:"fee"`
			Game                     string      `json:"game"`
			GoodsID                  int         `json:"goods_id"`
			HasBargain               bool        `json:"has_bargain"`
			HasSentOffer             bool        `json:"has_sent_offer"`
			ID                       string      `json:"id"`
			Income                   string      `json:"income"`
			IsSellerAskedToSendOffer bool        `json:"is_seller_asked_to_send_offer"`
			Mode                     int         `json:"mode"`
			OriginalPrice            interface{} `json:"original_price"`
			PayExpireTimeout         int         `json:"pay_expire_timeout"`
			PayMethod                int         `json:"pay_method"`
			PayMethodText            string      `json:"pay_method_text"`
			Price                    string      `json:"price"`
			PriceWithPayFee          string      `json:"price_with_pay_fee"`
			Progress                 int         `json:"progress"`
			ReceiveExpireTimeout     int         `json:"receive_expire_timeout"`
			SellOrderID              string      `json:"sell_order_id"`
			SellerCanCancel          bool        `json:"seller_can_cancel"`
			SellerCookieInvalid      bool        `json:"seller_cookie_invalid"`
			SellerID                 string      `json:"seller_id"`
			State                    string      `json:"state"`
			StateText                string      `json:"state_text"`
			TradeOfferTraceURL       interface{} `json:"trade_offer_trace_url"`
			TradeOfferURL            interface{} `json:"trade_offer_url"`
			Tradeofferid             interface{} `json:"tradeofferid"`
			TransactTime             interface{} `json:"transact_time"`
			Type                     int         `json:"type"`
			UpdatedAt                int         `json:"updated_at"`
		} `json:"items"`
	} `json:"data"`
	Msg interface{} `json:"msg"`
}