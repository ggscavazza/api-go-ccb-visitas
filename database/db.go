package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	stringDeConexao := "cc182652_admin:vila@2023@alpina@tcp(177.85.99.55:3306)/cc182652_app_ccbvisitas?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(stringDeConexao))

	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}
}
