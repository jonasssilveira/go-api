package payment

type PaymentData struct {
	Conta             int     `json:"conta"`
	Agencia           int     `json:"agencia"`
	CodigoVerificador int     `json:"codigo_verificador"`
	Valor             float64 `json:"valor"`
}
