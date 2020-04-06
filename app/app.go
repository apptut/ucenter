package app

import (
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"gopkg.in/redis.v5"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"strings"
	"sync"
	"time"
	"ucenter/app/svc/config"
)

var Config config.Model
var Locker sync.RWMutex

var dbServer *gorm.DB
var redisServer *redis.Client
var dbOnce, redisOnce sync.Once

// 是否是开发环境
func IsDev() bool {
	return Config.Env == config.Dev
}

// 生产环境
func IsProd() bool {
	return Config.Env == config.Prod
}

// 预上线环境
func IsStage() bool {
	return Config.Env == config.Stage
}

// 测试环境
func IsTest() bool {
	return Config.Env == config.Test
}

// 配置文件初始化
func configInit(prjHome string) error {
	configFile := strings.TrimRight(prjHome, "/") + "/config/config.yaml"
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(data, &Config)
	if err != nil {
		return err
	}

	Config.HomeDir = strings.TrimRight(prjHome, "/")

	return nil
}

func Db() *gorm.DB {
	dbOnce.Do(initDb)
	return dbServer
}

func Redis() *redis.Client {
	redisOnce.Do(initRedis)
	return redisServer
}

func Init(prjHome string) error {
	if err := configInit(prjHome); err != nil {
		return err
	}

	return nil
}

func initDb() {
	addr := Config.Mysql.Host + ":3306"
	if len(Config.Mysql.Port) != 0 {
		addr = Config.Mysql.Host + ":" + Config.Mysql.Port
	}

	loc, _ := time.LoadLocation(Config.Mysql.Timezone)

	dbConfig := &mysql.Config{
		Net:                  "tcp",
		Addr:                 addr,
		User:                 Config.Mysql.Username,
		Passwd:               Config.Mysql.Password,
		DBName:               Config.Mysql.Database,
		ParseTime:            true,
		AllowNativePasswords: true,
		Timeout:              time.Millisecond * time.Duration(Config.Mysql.Timeout),
		ReadTimeout:          time.Millisecond * time.Duration(Config.Mysql.ReadTimeout),
		WriteTimeout:         time.Millisecond * time.Duration(Config.Mysql.WriteTimeout),
		Loc:                  loc,
	}

	db, err := gorm.Open("mysql", dbConfig.FormatDSN())
	if err != nil {
		panic(err)
	}

	// 开发环境是否打印sql log
	//if IsDev() {
	//	db.LogMode(true)
	//}

	// 连接数设置
	db.DB().SetConnMaxLifetime(time.Millisecond * time.Duration(Config.Mysql.ConnMaxLifeTime))
	db.DB().SetMaxIdleConns(Config.Mysql.MaxIdleConns)

	dbServer = db
}

func initRedis() {
	redisServer = redis.NewClient(&redis.Options{
		Addr:        Config.Redis.Addr,
		PoolSize:    Config.Redis.PoolSize,
		Password:    Config.Redis.Password,
		IdleTimeout: time.Millisecond * time.Duration(Config.Redis.IdleTimeout),
		MaxRetries:  Config.Redis.Retries,
	})
}

func Destruct() {
	if dbServer != nil {
		_ = dbServer.Close()
	}

	if redisServer != nil {
		_ = redisServer.Close()
	}
}
