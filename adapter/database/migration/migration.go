package migration

import (
	"e-commerce/domain/entity/produto"
	"gorm.io/gorm"
)

func RunMigration(db *gorm.DB) {
	db.AutoMigrate(
		produto.Product{},
		produto.Categoria{},
	)
}
