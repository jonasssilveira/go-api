package produto

import (
	"time"
)

type Categoria struct {
	ID        string    `json:"id"gorm:"type:uuid;primaryKey;id"`
	Name      string    `json:"name"`
	Descricao string    `json:"descricao"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
