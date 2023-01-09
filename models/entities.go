package models

// struct para carregar os dados para dentro do banco de dados
// e enviar dados para a resposta da API:
type Todo struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}
