package transaction

type CreateTransactionRequest struct {
	AccountId     int    `json:"account_id"`
	OperationType int    `json:"operation_type"`
	Amount        string `json:"amount"`
}
