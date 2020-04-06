package config

const (
	Prod  = "prod"
	Stage = "stage"
	Test  = "test"
	Dev   = "dev"
)

type Model struct {
	Env     string       `yaml:"env"`
	Listen  string       `yaml:"listen"`
	HomeDir string       `yaml:"home_dir"`
	Mysql   mysqlDbModel `yaml:"mysql_db"`
	Redis   redisConfig  `yaml:"redis"`
	Http    httpConfig
	Https   httpsConfig `yaml:"https"`
}

type httpConfig struct {
	Listen string `yaml:"listen"`
}

type httpsConfig struct {
	Listen   string `yaml:"listen"`
	CertFile string `yaml:"cert_file"`
	CertKey  string `yaml:"cert_key"`
}
