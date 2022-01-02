package payment

import (
	"time"
)
type Pedido struct {
	Carrinho    *Carrinho `json:"carrinho"`
	Deadline    time.Time `json:"deadline"`
	PaymentType `json:"paymentType"`

}
