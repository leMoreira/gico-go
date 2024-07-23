package basedados

import (
	"log"

	"github.com/leandro/gico-go/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	// stringDeConexao := "host=localhost user=postgres password=postgres dbname=gicoobj port=5432 sslmode=disable"
	stringDeConexao := "host=10.100.34.38 user=webadmin password=GNHxlm98862 dbname=gicoobj port=5432 sslmode=disable"

	DB, err = gorm.Open(postgres.Open(stringDeConexao))

	if err != nil {
		log.Panic("Erro ao conectar ao banco de dados")
	}

	DB.AutoMigrate(&models.Equipe{},
		&models.Prova{},
		&models.Usuario{},
		&models.Pontos{},
		&models.TotalPontos{})

}
