package payment

type PaymentType struct {
	Boleto      bool `json:"boleto"`
	Credit      bool `json:"credit"`
	Pix         bool `json:"pix"`
	CreditData  `json:"creditData"`
	PaymentData `json:"paymentData"`
}
