package config

// redis 配置
type redisConfig struct {
	Addr        string `yaml:"addr"`
	Password    string `yaml:"password"`
	PoolSize    int    `yaml:"pool_size"`
	IdleTimeout int    `yaml:"idle_timeout"`
	Retries     int    `yaml:"retries"`
}
