package orm

import (
	"database/sql"
	"douyin/config"
	"fmt"
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

var (
	dbClient    *sql.DB
	Gorm        *gorm.DB
	err         error
	RedisClient *redis.Client
)

var c = config.Config

func Start() {
	//mysql连接池
	Gorm, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", c.Mysql.User, c.Mysql.Password, c.Mysql.Host, c.Mysql.Port, c.Mysql.Database))
	if err != nil {
		log.Println("connect mysql err:" + err.Error())
		return
	}
	dbClient = Gorm.DB()
	Gorm.LogMode(c.Mysql.Logdebug)
	dbClient.SetMaxOpenConns(c.Mysql.MaxActive) //用于设置最大打开的连接数，默认值为0表示不限制。
	dbClient.SetMaxIdleConns(c.Mysql.MaxIdle)   //最大空闲数
	dbClient.Ping()

	//goredis
	RedisClient = redis.NewClient(&redis.Options{
		DB:           c.Redis.Database,
		Password:     c.Redis.Password,
		MinIdleConns: c.Redis.MaxIdle,
		PoolSize:     c.Redis.MaxActive,
		Addr:         c.Redis.Host + ":" + c.Redis.Port,
	})
}
