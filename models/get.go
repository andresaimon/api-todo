package models

import "github.com/andresaimon/api-todo/db"

// Busca de tarefas por status (true ou false):
func Get(done bool) (todos []Todo, err error) {
	// abrindo conexão com o banco de dados:
	conn, err := db.OpenConnection()

	// validando o tipo de erro:
	if err != nil {
		return
	}

	// fechando a conexão com o banco de dados:
	defer conn.Close()

	// SQL a ser executado:
	rows, err := conn.Query(`SELECT * FROM todos WHERE done=$1`, done)
	if err != nil {
		return
	}

	// seleciona os itens adicionados ao banco de dados:
	for rows.Next() {
		var todo Todo
		err = rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done)
		if err != nil {
			continue
		}

		todos = append(todos, todo)
	}

	return
}
