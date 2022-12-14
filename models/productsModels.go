package models

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

type Products struct {
	gorm.Model
	NomeProduto      string  `json:"nomeproduto"`
	Descricao        string  `json:"descricao"`
	Valor            float64 `json:"valor"`
	Fornecedor       string  `json:"fornecedor"`
	AvaliacaoProduto int64   `json:"avaliacaoproduto"`
}

func InitialMigration() {
	password := os.Getenv("PASSWORD")

	DB, err = gorm.Open(mysql.New(mysql.Config{
		DSN: "root:" + password + "@/products?charset=utf8mb4&parseTime=True&loc=Local",
	}), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Nao se conectou com o DB")
	}

	DB.AutoMigrate(&Products{})
}
