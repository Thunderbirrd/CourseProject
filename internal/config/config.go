package config

import "github.com/Thunderbirrd/CourseProject/pkg/helpers"

type Config struct {
	HttpPort   string
	DBHost     string
	DBPort     string
	DBName     string
	DBUsername string
	DBPassword string
	DBSSLMode  string
}

func New() (*Config, error) {
	cfg := &Config{}

	helpers.EnvToString(&cfg.HttpPort, "HTTP_PORT", "8090")
	helpers.EnvToString(&cfg.DBHost, "DB_HOST", "ec2-52-212-228-71.eu-west-1.compute.amazonaws.com")
	helpers.EnvToString(&cfg.DBPort, "DB_PORT", "5432")
	helpers.EnvToString(&cfg.DBUsername, "DB_USERNAME", "fzcamrgntritxl")
	helpers.EnvToString(&cfg.DBName, "DB_NAME", "d7r6s29qsbs7ah")
	helpers.EnvToString(&cfg.DBPassword, "DB_PASSWORD", "7945968eed261e91271a72e87735b1f85715d098b45847238aaeae55b2d9d535")
	helpers.EnvToString(&cfg.DBSSLMode, "DB_SSL_MODE", "require")

	return cfg, nil
}
