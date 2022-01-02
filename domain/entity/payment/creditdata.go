package payment

type CreditData struct {
	CardNumber         string `json:"card_number"`
	CardCVV            string `json:"card_cvv"`
	CarditName         string `json:"cardit_name"`
	CarditCpf          int    `json:"cardit_cpf"`
	CardValidationDate string `json:"card_validation_date"`
}
