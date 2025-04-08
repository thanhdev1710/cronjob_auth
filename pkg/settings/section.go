package settings

type Config struct {
	Port    string `mapstructure:"PORT"`
	Env     string `mapstructure:"ENV"`
	ApiKey  string `mapstructure:"API_KEY"`
	Postgre Postgre
}

type Postgre struct {
	Host            string `mapstructure:"HOST_DB"`
	Port            string `mapstructure:"PORT_DB"`
	Username        string `mapstructure:"USERNAME_DB"`
	Password        string `mapstructure:"PASSWORD_DB"`
	Database        string `mapstructure:"DATABASE_DB"`
	SslMode         string `mapstructure:"SSL_MODE_DB"`
	ConnMaxIdleTime int    `mapstructure:"CONN_MAX_IDLE_TIME_DB"`
	ConnMaxLifeTime int    `mapstructure:"CONN_MAX_LIFE_TIME"`
	ConnMaxOpen     int    `mapstructure:"CONN_MAX_OPEN"`
}
