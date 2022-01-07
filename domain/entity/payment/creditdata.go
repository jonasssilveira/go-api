package payment

import (
	"gorm.io/gorm"
)

type CreditData struct {
	gorm.Model
	Id       string             `json:"id"gorm:"type:uuid;primaryKey;id"`
	CardNumber         string `json:"card_number"`
	CardCVV            string `json:"card_cvv"`
	CarditName         string `json:"cardit_name"`
	CarditCpf          int    `json:"cardit_cpf"`
	CardValidationDate string `json:"card_validation_date"`
}
