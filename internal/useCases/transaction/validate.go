package transaction

import "errors"

const (
	OperationNormalPurchase         = 1
	OperationPurchaseInInstallments = 2
	OperationWithdrawal             = 3
	OperationCreditVoucher          = 4
)

func ValidateTransaction(operationType int, amount float64) error {

	switch operationType {
	case OperationNormalPurchase, OperationWithdrawal, OperationPurchaseInInstallments:
		if amount >= 0 {
			return errors.New("amount must be negative for this operation type")
		}
	case OperationCreditVoucher:
		if amount <= 0 {
			return errors.New("amount must be positive for credit voucher operation")
		}
	default:
		return errors.New("invalid operation type")
	}

	return nil
}
