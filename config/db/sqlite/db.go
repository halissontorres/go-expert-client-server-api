package sqlite

import (
	"os"

	"github.com/halissontorres/go-expert-client-server-api/config/log"
	"github.com/halissontorres/go-expert-client-server-api/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := log.NewLogger("sqlite")
	dbPath := "./db/cotacao.db"

	_, err := os.Stat(dbPath)

	if os.IsNotExist(err) {
		logger.Info("Database não encontrada. Criando banco de dados.")

		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}
		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}
		file.Close()
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Error("Erro ao conectar ao banco de dados: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&model.UsdBrl{})
	if err != nil {
		logger.Error("sqlite erro em automigrate: %v", err)
		return nil, err
	}
	return db, nil
}
