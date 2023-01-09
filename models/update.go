package models

import "github.com/silastgoes/API-To-Do-List/db"

// Atualizando tarefas:
// função para receber como parâmetro o ID e os itens cadastrados
func Update(id int64, todo Todo) (int64, error) {
	// abrindo a conexão:
	conn, err := db.OpenConnection()

	// validando o erro:
	if err != nil {
		return 0, err
	}

	// fechando a conexão:
	defer conn.Close()

	// query SQL para realizar o Update:
	// $1, $2, $3 e $4 referem-se à ordem definida nos parâmetros em seguida
	res, err := conn.Exec(`UPDATE todos SET title=$2, description=$3, done=$4 WHERE id=$1`, id, todo.Title, todo.Description, todo.Done)
	if err != nil {
		return 0, err
	}

	// retorna o número de linhas que foram afetadas com a execução do Update:
	return res.RowsAffected()
}
