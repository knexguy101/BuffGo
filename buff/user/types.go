package user

type Info struct {
	Code string `json:"code"`
	Data struct {
		AllowAutoRemark              bool    `json:"allow_auto_remark"`
		AllowBargainRejectedPush     bool    `json:"allow_bargain_rejected_push"`
		AllowBuyerBargain            bool    `json:"allow_buyer_bargain"`
		AllowCommentPush             bool    `json:"allow_comment_push"`
		AllowCsgoTradeUp             bool    `json:"allow_csgo_trade_up"`
		AllowCsgoTradeUpCommunity    bool    `json:"allow_csgo_trade_up_community"`
		AllowDeliverPush             bool    `json:"allow_deliver_push"`
		AllowEpay                    bool    `json:"allow_epay"`
		AllowFeedbackNewEntry        bool    `json:"allow_feedback_new_entry"`
		AllowMailNotification        bool    `json:"allow_mail_notification"`
		AllowNicknameSign            bool    `json:"allow_nickname_sign"`
		AllowPreviewAudit            bool    `json:"allow_preview_audit"`
		AllowPreviewRecommend        bool    `json:"allow_preview_recommend"`
		AllowPriceChangeNotify       bool    `json:"allow_price_change_notify"`
		AllowPubgRecycle             bool    `json:"allow_pubg_recycle"`
		AllowPurchasePremium         bool    `json:"allow_purchase_premium"`
		AllowSensitiveContentControl bool    `json:"allow_sensitive_content_control"`
		AllowShopDisplay             bool    `json:"allow_shop_display"`
		AllowSmsNotification         bool    `json:"allow_sms_notification"`
		AllowSocialComment           bool    `json:"allow_social_comment"`
		AllowUpPush                  bool    `json:"allow_up_push"`
		AllowWechatTradeMessage      bool    `json:"allow_wechat_trade_message"`
		Avatar                       string  `json:"avatar"`
		BuffPriceCurrency            string  `json:"buff_price_currency"`
		BuffPriceCurrencyDesc        string  `json:"buff_price_currency_desc"`
		BuffPriceCurrencyRateBaseCny float64 `json:"buff_price_currency_rate_base_cny"`
		BuffPriceCurrencyRateBaseUsd float64 `json:"buff_price_currency_rate_base_usd"`
		BuffPriceCurrencySymbol      string  `json:"buff_price_currency_symbol"`
		BuyerExamState               int     `json:"buyer_exam_state"`
		CanUnbindSteam               bool    `json:"can_unbind_steam"`
		ForceBuyerSendOffer          bool    `json:"force_buyer_send_offer"`
		ID                           string  `json:"id"`
		InventoryPrice               string  `json:"inventory_price"`
		IsForeigner                  bool    `json:"is_foreigner"`
		IsNeedSteamVerify            bool    `json:"is_need_steam_verify"`
		IsNew                        bool    `json:"is_new"`
		IsPremium                    bool    `json:"is_premium"`
		Language                     string  `json:"language"`
		LoginFrom                    int     `json:"login_from"`
		Mobile                       string  `json:"mobile"`
		Nickname                     string  `json:"nickname"`
		NicknameRemaining            int     `json:"nickname_remaining"`
		SellerExamState              int     `json:"seller_exam_state"`
		ShowLeaderboard              struct {
			Csgo  bool `json:"csgo"`
			Dota2 bool `json:"dota2"`
		} `json:"show_leaderboard"`
		ShowSteamAssetBuyPrice struct {
			Csgo  bool `json:"csgo"`
			Dota2 bool `json:"dota2"`
		} `json:"show_steam_asset_buy_price"`
		ShowSteamAssetRemark struct {
			Csgo  bool `json:"csgo"`
			Dota2 bool `json:"dota2"`
		} `json:"show_steam_asset_remark"`
		SteamAPIKeyState       int    `json:"steam_api_key_state"`
		SteamPriceCurrency     string `json:"steam_price_currency"`
		SteamPriceCurrencyDesc string `json:"steam_price_currency_desc"`
		Steamid                string `json:"steamid"`
		TradeURL               string `json:"trade_url"`
		TradeURLState          int    `json:"trade_url_state"`
	} `json:"data"`
	Msg interface{} `json:"msg"`
}
