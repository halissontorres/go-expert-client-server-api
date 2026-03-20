//go:build client

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type Cotacao struct {
	USDBRL struct {
		Bid string `json:"bid"`
	} `json:"USDBRL"`
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		log.Fatalf("Erro ao criar requisição: %v", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Timeout ou erro ao obter cotação: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Erro ao ler resposta: %v", err)
	}

	var cotacao Cotacao
	if err := json.Unmarshal(body, &cotacao); err != nil {
		log.Fatalf("Erro ao desserializar cotação: %v", err)
	}

	content := fmt.Sprintf("Dólar: %s", cotacao.USDBRL.Bid)

	if err := os.WriteFile("cotacao.txt", []byte(content), 0644); err != nil {
		log.Fatalf("Erro ao salvar cotação em arquivo: %v", err)
	}

	fmt.Println(content)
}
