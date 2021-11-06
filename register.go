package main


type WalletType interface {
	// 获取预签名信息
	GetSignInfo(address string) (json string, err error)
	// 余额查询
	GetBalance(address string) (balance string, err error)
	// 广播交易
	BroadcastScript(transaction_script string) (transaction_id_chain string, err error)
	// 交易状态查询
	GetTxStatus(transaction_id_chain string) (resp *GetTxStatusResponse, err error)
	// 账户交易查询
	GetTxJSON(address string, timefrom_utc uint64, timetill_utc uint64, limit_n int32) (json string, err error)
}


type GetTxStatusResponse struct {
	AmountString       string
	AddressFrom        string
	AddressTo          string
	TransactionIdChain string
	CreatetimeUtc      int64
	UpdatetimeUtc      int64
	IsSuccess          bool
	IsFailed           bool
	NumBlocksConfirmed int32
}

type GetTxJSONRequest struct {
	CoinId      int32
	Address     string
	TimefromUtc int64
	TimetillUtc int64
	LimitN      int32
}