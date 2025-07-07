package transaction

import "time"

type Passbook struct {
	AccountNo    int
	Transactions []*Transaction
}

type Transaction struct {
	TxID      int
	Type      string
	Amount    float32
	Timestamp time.Time
}
