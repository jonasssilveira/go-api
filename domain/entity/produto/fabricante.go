package produto

import (
	"gorm.io/gorm"
)

type Fabricante struct {
	gorm.Model
	Id       string             `json:"id"gorm:"type:uuid;primaryKey;id"`
	Name      string `json:"name"`
	Descricao string `json:"descricao"`
}
