package transaction

import (
	"fmt"

	"github.com/ArturLima/pismo/internal/store/pgstore"
)

var operationTypeNames = map[int]string{
	1: "normal_purchase",
	2: "purchase_in_installments",
	3: "withdrawal",
	4: "credit_voucher",
}

type TransactionResponse struct {
	AccountId     int    `json:"account_id"`
	OperationType string `json:"operation_type"`
	Amount        string `json:"amount"`
}

func ToTransacationResponse(req pgstore.Transaction) TransactionResponse {

	floatVal, _ := req.Amount.Float64Value()
	amountStr := fmt.Sprintf("%.2f", floatVal.Float64)

	return TransactionResponse{
		AccountId:     int(req.AccountID),
		OperationType: operationTypeNames[int(req.OperationID)],
		Amount:        amountStr,
	}
}
