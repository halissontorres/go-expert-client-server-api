package handler

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/halissontorres/go-expert-client-server-api/config"
	"github.com/halissontorres/go-expert-client-server-api/model"
)

func Cotacao(w http.ResponseWriter, r *http.Request) {
	logger := config.GetLogger("handler.Cotacao")
	timeOut, err := time.ParseDuration(os.Getenv("COTACAO_API_TIMEOUT"))
	if err != nil {
		timeOut = config.COTACAO_API_TIMEOUT
	}

	ctx, cancel := context.WithTimeout(r.Context(), timeOut)
	ctx, cancel = context.WithTimeout(ctx, timeOut)
	defer cancel()

	url := config.COTACAO_API_URL + config.MOEDAS
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusRequestTimeout)
		w.Write([]byte(`{"error":"Tempo de resposta excedido"}`))
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		logger.Error("status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Error("Erro ao ler resposta: %v", err)
	}

	defer resp.Body.Close()

	var cotacao model.Cotacao
	err = json.Unmarshal(body, &cotacao)
	if err != nil {
		logger.Error("Erro ao desserializar resposta: %v", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cotacao)
}
