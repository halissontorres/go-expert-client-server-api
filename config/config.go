package config

import (
	"fmt"
	"time"

	"github.com/halissontorres/go-expert-client-server-api/config/db/sqlite"
	"github.com/halissontorres/go-expert-client-server-api/config/log"
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *log.Logger
)

const (
	COTACAO_API_URL     = "https://economia.awesomeapi.com.br/json/last/"
	COTACAO_API_TIMEOUT = 200 * time.Millisecond
	COTACAO_DB_TIMEOUT  = 10 * time.Millisecond
	SERVER_PORT         = "8080"
	MOEDAS              = "USD-BRL"
)

func Init() error {
	var err error

	db, err = sqlite.InitializeSQLite()

	if err != nil {
		return fmt.Errorf("Erro ao inicializar sqlite: %v", err)
	}

	return nil
}

func GetSQLite() *gorm.DB {
	return db
}

func GetLogger(p string) *log.Logger {
	logger = log.NewLogger(p)
	return logger
}
