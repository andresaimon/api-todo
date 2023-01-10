package models

import "github.com/andresaimon/api-todo/db"

// Busca de todas as tarefas cadastrada no banco de dados:
// retorna um slice de tarefas
func GetAll() (todos []Todo, err error) {
	// abrindo conexão com o banco de dados:
	conn, err := db.OpenConnection()

	// validando o tipo de erro:
	if err != nil {
		return
	}

	// fechando a conexão com o banco de dados:
	defer conn.Close()

	// SQL a ser executado:
	rows, err := conn.Query(`SELECT * FROM todos`)
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
