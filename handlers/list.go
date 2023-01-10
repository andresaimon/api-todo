package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/andresaimon/api-todo/models"
)

// Função para retornar todos os registros:
func List(w http.ResponseWriter, r *http.Request) {
	todos, err := models.GetAll()
	if err != nil {
		log.Printf("Erro ao obter registros: %v", err)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}
