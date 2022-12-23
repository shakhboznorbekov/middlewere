package config

type Config struct {
	HTTPPort string

	PostgresHost           string
	PostgresUser           string
	PostgresDatabase       string
	PostgresPassword       string
	PostgresPort           string
	PostgresMaxConnections int32

	AuthSecretKey string
	SuperAdmin    string
	Client        string
}

func Load() Config {

	var cfg Config

	cfg.HTTPPort = ":4000"

	cfg.PostgresHost = "localhost"
	cfg.PostgresUser = "admin"
	cfg.PostgresDatabase = "dvdrental"
	cfg.PostgresPassword = "2711"
	cfg.PostgresPort = "5432"
	cfg.PostgresMaxConnections = 20

	cfg.AuthSecretKey = "9K+WgNTglA44Hg=="
	cfg.SuperAdmin = "Super"
	cfg.Client = "Client"

	return cfg
}
