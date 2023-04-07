package configs

import (
	"github.com/spf13/viper"
)

var cfg *conf

type conf struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBHost        string `mapstructure:"DB_HOST"`
	DBPort        string `mapstructure:"DB_PORT"`
	DBUser        string `mapstructure:"DB_USER"`
	DBPassword    string `mapstructure:"DB_PASSWORD"`
	DBName        string `mapstructure:"DB_NAME"`
	WebServerPort string `mapstructure:"WEB_SERVER_PORT"`
	GrpcPort      string `mapstructure:"GRPC_PORT"`
	GraphQLPort   string `mapstructure:"GRAPHQL_PORT"`
}

func LoadConfig(filePath string) (*conf, error) {
	//Nome da configuração
	viper.SetConfigName("app_config")
	//Tipo de artuiqvo
	viper.SetConfigType("env")
	//Caminho do arquivo
	viper.AddConfigPath(filePath)
	//Nome do arquivoo
	viper.SetConfigFile(".env")
	//Configuração automatica de envs
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	if err = viper.Unmarshal(&cfg); err != nil {
		panic(err)
	}
	return cfg, nil
}
