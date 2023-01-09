package db

import (
	"database/sql"
	"fmt"

	// Definição do drive do Postgres:
	_ "github.com/lib/pq"
	"github.com/silastgoes/API-To-Do-List/configs"
)

// Função para abrir conexão com o banco de dados:
func OpenConnection() (*sql.DB, error) {
	conf := configs.GetDB()

	// sc = string connection
	sc := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.Host, conf.Port, conf.User, conf.Pass, conf.Database)

	// abrindo conexão com o banco de dados:
	conn, err := sql.Open("postgres", sc)

	// validando o tipo de erro:
	if err != nil {
		panic(err)
	}

	// ping no banco de dados para verificar se a conexão foi estabelecida:
	err = conn.Ping()

	return conn, err
}
