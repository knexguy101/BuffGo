package purchase

const (
	ALIPAY PayMethod = 10
	BUFFBALANCE PayMethod = 3
	BUFFBALANCEBANKCARD PayMethod = 1
	WECHATPAY PayMethod = 6
)

type PayMethod int
