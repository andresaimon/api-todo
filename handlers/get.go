package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/andresaimon/api-todo/models"
	"github.com/go-chi/chi/v5"
)

// Função para filtrar por status (true ou false):
func Get(w http.ResponseWriter, r *http.Request) {
	done, err := strconv.ParseBool(chi.URLParam(r, "done"))
	if err != nil {
		log.Printf("Erro ao fazer parse do id: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	todo, err := models.Get(done)
	if err != nil {
		log.Printf("Erro ao atualizar registro: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}
