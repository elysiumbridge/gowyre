package main

// Tx represent a blockchain transaction object
type Tx struct {
	ID            string  `json:"id"`
	NetworkTxID   string  `json:"networkTxId"`
	CreatedAt     uint64  `json:"createdAt"`
	Confirmations uint64  `json:"confirmations"`
	TimeObserved  uint64  `json:"timeObserved"`
	BlockTime     uint64  `json:"blockTime"`
	Blockhash     string  `json:"blockhash"`
	Amount        float64 `json:"amount"`
	Direction     string  `json:"direction"`
	NetworkFee    float64 `json:"networkFee"`
	Address       string  `json:"address"`
	SourceAddress *string `json:"sourceAddress"`
	Currency      string  `json:"currency"`
	TwinTxId      *string `json:"twinTxId"`
}
