package payment

import (
	"gorm.io/gorm"
)

type PaymentType struct {
	gorm.Model
	Id       string             `json:"id"gorm:"type:uuid;primaryKey;id"`
	Boleto      bool `json:"boleto"`
	Credit      bool `json:"credit"`
	Pix         bool `json:"pix"`
	CreditData  `json:"creditData"`
	PaymentData `json:"paymentData"`
}
