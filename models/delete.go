package models

import "github.com/silastgoes/API-To-Do-List/db"

// Deletar tarefas:
func Delete(id int64) (int64, error) {
	// abrindo a conexão:
	conn, err := db.OpenConnection()

	// validando o erro:
	if err != nil {
		return 0, err
	}

	// fechando a conexão:
	defer conn.Close()

	/// query SQL para realizar o Delete:
	res, err := conn.Exec(`DELETE FROM todos WHERE id=$1`, id)
	if err != nil {
		return 0, err
	}

	// retorna o número de linhas que foram deletadas:
	return res.RowsAffected()
}
