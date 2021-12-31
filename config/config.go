package config

import "time"

type DBConfig struct {
	DBDriver         string `mapstructure:"db_driver"`
	PostgresUser     string `mapstructure:"postgres_user"`
	PostgresPassword string `mapstructure:"psotgres_password"`
	PostgresDB       string `mapstructure:"postgres_db"`
	PostgresHost     string `mapstructure:"postgres_host"`
	PostgresPort     string `mapstructure:"postgres_port"`
}

type RouterConfig struct {
	Address      string        `mapstructure:"address"`
	ReadTimeout  time.Duration `mapstructure:"read_timeout"`
	WriteTimeout time.Duration `mapstructure:"write_timeout"`
}

type GeneralConfig struct {
	RouterConfig RouterConfig `mapstructure:"router_config"`
	DBConfig     DBConfig     `mapstructure:"db_config"`
}
