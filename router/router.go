package router

import (
	"net/http"
	"os"

	"github.com/halissontorres/go-expert-client-server-api/config"
	"github.com/halissontorres/go-expert-client-server-api/handler"
)

func init() {
	logger := config.GetLogger("router")

	if err := config.Init(); err != nil {
		logger.Fatal("Erro ao inicializar configurações: %v", err)
	}

	porta := os.Getenv("COTACAO_SERVER_PORT")
	if porta == "" {
		porta = config.SERVER_PORT
	}
	logger.Info("Servidor escutando na porta %s", porta)

	http.HandleFunc("/cotacao", handler.Cotacao)

	logger.Fatal("Error ao iniciar servidor %s", http.ListenAndServe(":"+porta, nil))
}
