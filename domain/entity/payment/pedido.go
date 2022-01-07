package payment

import (
	"gorm.io/gorm"
	"time"
)

type Pedido struct {
	gorm.Model
	Id          string `json:"id"gorm:"type:uuid;primaryKey;id"`
	Carrinho    *Carrinho `json:"carrinho"gorm:"foreignKey:Id;references:Id"`
	Deadline    time.Time `json:"deadline"`
	PaymentType `json:"paymentType"`
}
