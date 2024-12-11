package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB

	err error
)

func ConectaComBancodeDados() {
	strConexao := "host=localhost user=postgres password=admin dbname=postgres port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(strConexao))

	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}

}
