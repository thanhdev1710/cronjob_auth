package initialize

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"github.com/thanhdev1710/cronjob_auth/global"
)

func LoadConfig() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("No .env file found, loading from OS ENV")
	}

	viper.BindEnv("PORT")
	viper.BindEnv("API_KEY")
	viper.BindEnv("ENV")

	viper.BindEnv("LOG_LEVEL")
	viper.BindEnv("LOG_FILE")
	viper.BindEnv("LOG_MAXSIZE")
	viper.BindEnv("LOG_MAXBACKUPS")
	viper.BindEnv("LOG_MAXAGE")
	viper.BindEnv("LOG_COMPRESS")

	viper.BindEnv("HOST_DB")
	viper.BindEnv("PORT_DB")
	viper.BindEnv("USERNAME_DB")
	viper.BindEnv("PASSWORD_DB")
	viper.BindEnv("DATABASE_DB")
	viper.BindEnv("SSL_MODE_DB")
	viper.BindEnv("CONN_MAX_IDLE_TIME_DB")
	viper.BindEnv("CONN_MAX_OPEN")
	viper.BindEnv("CONN_MAX_LIFE_TIME")

	err = viper.Unmarshal(&global.Config)
	if err != nil {
		log.Fatalf("Cannot unmarshal config: %v", err)
	}
	err = viper.Unmarshal(&global.Config.Logger)
	if err != nil {
		log.Fatalf("Cannot unmarshal config: %v", err)
	}
	err = viper.Unmarshal(&global.Config.Postgre)
	if err != nil {
		log.Fatalf("Cannot unmarshal config: %v", err)
	}
}
