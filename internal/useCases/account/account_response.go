package account

import "time"

type AccountResponse struct {
	ID             int32     `json:"id"`
	DocumentNumber string    `json:"document_number"`
	CreatedAt      time.Time `json:"created_at"`
}
