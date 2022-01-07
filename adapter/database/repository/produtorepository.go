package repository

import (
	"e-commerce/domain/entity/produto"
 	"errors"
	"gorm.io/gorm"
	"log"
)

const (
	ERROR_DEFAULT = "ocorreu um erro "
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB)*ProductRepository{
	return &ProductRepository{
		db: db,
	}
}

func (p *ProductRepository)ListAllProducts() []produto.Product{
	var produtos []produto.Product
	err := p.db.Find(&produtos).Error
	if err != nil{
		log.Println(ERROR_DEFAULT,err)
		return produtos
	}
	return produtos
}
func (p *ProductRepository)ListAllProductsByCategory() []produto.Product{
	return nil
}
func (p *ProductRepository)GetProductById(id string) *produto.Product{
	var produto produto.Product
	err := p.db.Find(&produto, "id = ?", id).Error
	if err != nil{
		log.Println(ERROR_DEFAULT,err)
		return &produto
	}
	return &produto
}
func (p *ProductRepository)GetProductByName(name string) (*produto.Product,error){
	var produto produto.Product
	err := p.db.Find(&produto, "name = ?", name).Error
	if err != nil{
		log.Println("",err)
		return nil, errors.New(ERROR_DEFAULT+err.Error())
	}
	return &produto, nil
}
func (p *ProductRepository)GetProductByNameOrCategoryName(name string) *produto.Product{
	return nil
}
func (p *ProductRepository)SaveProduct(product *produto.Product) *produto.Product{
	err := p.db.Create(&product)
	if err!= nil{
		log.Println(ERROR_DEFAULT,err)
		return nil
	}
	return product
}
func (p *ProductRepository)UpdateProduct(product *produto.Product) *produto.Product{
	return nil
}
func (p *ProductRepository)GetProductByVenda(id string) (*produto.Product,error){
	return nil, nil
}
func (p *ProductRepository)DeleteProduct(id string) error{
	var product produto.Product

	err := db.Find(product, "id = ?", id).Error

	if err != nil {
		log.Println(ERROR_DEFAULT,err)
		return err
	}

	err = db.Delete(&product).Error

	if err != nil{
		log.Println(ERROR_DEFAULT,err)
		return err
	}
	return nil
}