package configs

import "github.com/spf13/viper"

var cfg *config

// definição das structs:
type config struct {
	API APIConfig
	DB  DBConfig
}

type APIConfig struct {
	Port string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Pass     string
	Database string
}

// definição de valores padrões através do pacote viper
// api.port, database.host e database.port: chaves de configuração
func init() {
	viper.SetDefault("api.port", "9000")
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", "5432")
}

// função que carrega as configurações:
// a função retorna um erro, caso não seja possivel carregar as configurações
// define-se o nome, o tipo e o local do arquivo
func Load() error {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	//validação do tipo de erro ocorrido:
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	//criação de um ponteiro da struct:
	cfg = new(config)

	//para atribuir o dado dentro da struct:
	cfg.API = APIConfig{
		Port: viper.GetString("api.port"),
	}

	cfg.DB = DBConfig{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetString("database.port"),
		User:     viper.GetString("database.user"),
		Pass:     viper.GetString("database.pass"),
		Database: viper.GetString("database.name"),
	}

	return nil
}

// função para retornar o objeto de configuração do banco:
func GetDB() DBConfig {
	return cfg.DB
}

// função para retornar a porta
func GetServerPort() string {
	return cfg.API.Port
}
