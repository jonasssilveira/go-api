package payment

import (
	"gorm.io/gorm"
)

type PaymentData struct {
	gorm.Model
	Id       string             `json:"id"gorm:"type:uuid;primaryKey;id"`
	Conta             int     `json:"conta"`
	Agencia           int     `json:"agencia"`
	CodigoVerificador int     `json:"codigo_verificador"`
	Valor             float64 `json:"valor"`
}
