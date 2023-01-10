package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/andresaimon/api-todo/models"
	"github.com/go-chi/chi/v5"
)

// Os valores recebidos pela URL são strings
// Atoi: converte uma string para um inteiro
func Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	// validando o tipo de erro:
	if err != nil {
		log.Printf("Erro ao fazer parse do ID: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// parse da request que está sendo recebida:
	var todo models.Todo

	// payload de uma request em JSON
	err = json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		log.Printf("Erro ao fazer decode do json: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// atualizando os dados no banco de dados:
	rows, err := models.Update(int64(id), todo)
	if err != nil {
		log.Printf("Erro ao atualizar registro: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// mensagem de erro, caso sejam atualizados mais de um registro
	if rows > 1 {
		log.Printf("Error: foram atualizados %d registros", rows)
	}

	resp := map[string]any{
		"Error":   false,
		"Message": "Dados atualizados com sucesso!",
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
