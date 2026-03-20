package main

import (
	"github.com/halissontorres/go-expert-client-server-api/config"
	_ "github.com/halissontorres/go-expert-client-server-api/router"
)

func main() {
	config.GetLogger("main").Info("INICIANDO SERVIÇO DE COTAÇÃO")
}
