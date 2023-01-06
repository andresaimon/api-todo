package models

import "github.com/silastgoes/API-To-Do-List/db"

// deletar tarefas:
func Delete(id int64) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()
	//$1, $2, $3 e $4 referem-se à ordem definida nos parâmetros em seguida
	res, err := conn.Exec(`DELETE FROM todos WHERE id=$1`, id)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
