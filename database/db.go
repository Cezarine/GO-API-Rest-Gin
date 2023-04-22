package database

import (
	"log"

	"github.com/cezarine/API-GO-GIN/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaBanco() {
	dsn := "host=DESKTOP-IKTKK3V user=sistema password=om315 dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Fatal("Erro ao conectar com o banco" + err.Error())
	}
	DB.AutoMigrate(&models.Aluno{})
}
