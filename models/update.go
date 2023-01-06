package models

import "github.com/silastgoes/API-To-Do-List/db"

// atualizar tarefas:
func Update(id int64, todo Todo) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()
	//$1, $2, $3 e $4 referem-se à ordem definida nos parâmetros em seguida
	res, err := conn.Exec(`UPDATE todos SET title=$2, description=$3, done=$4 WHERE id=$1`, id, todo.Title, todo.Description, todo.Done)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
