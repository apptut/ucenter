package config

// 数据库配置
type mysqlDbModel struct {
	Host         string `yaml:"host"`
	Port         string `yaml:"port"`
	Database     string `yaml:"database"`
	Password     string `yaml:"password"`
	Username     string `yaml:"username"`
	Timezone     string `yaml:"timezone"`
	Timeout      int    `yaml:"timeout"`
	ReadTimeout  int    `yaml:"read_time_out"`
	WriteTimeout int    `yaml:"write_time_out"`

	// connect pool
	ConnMaxLifeTime int `yaml:"conn_max_life_time"`
	MaxIdleConns    int `yaml:"max_idle_conns"`
}
