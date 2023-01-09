package models

import "github.com/silastgoes/API-To-Do-List/db"

// Abrindo conexão com o banco de dados e inserindo dados:
func Insert(todo Todo) (id int64, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	// fechando a conexão:
	defer conn.Close()

	// SQL a ser executado:
	sql := `INSERT INTO todos (title, description, done) VALUES ($1, $2, $3) RETURNING id`

	// transação com o banco de dados:
	err = conn.QueryRow(sql, todo.Title, todo.Description, todo.Done).Scan(&id)

	// retorna qual o ID, caso ocorra algum erro:
	return
}
