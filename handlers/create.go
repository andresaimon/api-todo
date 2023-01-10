package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/andresaimon/api-todo/models"
)

// Rota para criar uma nova entrada no banco de dados:
// função do tipo HTTP
func Create(w http.ResponseWriter, r *http.Request) {
	// parse da request que está sendo recebida:
	var todo models.Todo

	// payload de uma request em JSON
	err := json.NewDecoder(r.Body).Decode(&todo)

	// validando o tipo de erro:
	// Error do HTTP recede 3 parâmetros (response writer; a mensagem de erro; e o status code)
	if err != nil {
		log.Printf("Erro ao fazer decode do json: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// dados do payload recebido:
	id, err := models.Insert(todo)

	// variável de resposta:
	var resp map[string]any

	if err != nil {
		resp = map[string]any{
			"Error":   true,
			"Message": fmt.Sprintf("Ocorreu um erro ao tentar inserir: %v", err),
		}
	} else {
		resp = map[string]any{
			"Error":   false,
			"Message": fmt.Sprintf("Tarefa inserida com sucesso! ID: %d", id),
		}
	}

	// realização de um encode do map de resposta:
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
