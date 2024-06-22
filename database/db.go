package database

import (
	"log"
	"os"
	"github.com/guilhermeonrails/api-go-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"fmt"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	host := os.Getenv("HOST")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DBNAME")
	port := os.Getenv("DBPORT")

	log.Printf("Conectando ao banco de dados com as seguintes configurações:\nHost: %s\nUser: %s\nDBName: %s\nPort: %s\n", host, user, dbname, port)

	stringDeConexao := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", host, user, password, dbname, port)
	DB, err := gorm.Open(postgres.Open(stringDeConexao), &gorm.Config{})
	if err != nil {
		log.Fatalf("Erro ao conectar com banco de dados: %v", err)
	}

	DB.AutoMigrate(&models.Aluno{})
}
