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


func GetPassBookPaginated(allTx []*Transaction, page, size int) []*Transaction {
	start := (page - 1) * size
	end := start + size

	if start >= len(allTx) {
		return []*Transaction{}
	}
	if end > len(allTx) {
		end = len(allTx)
	}
	return allTx[start:end]
}

