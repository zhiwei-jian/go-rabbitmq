package config

import (
	compostgres "github.com/zhiwei-jian/common-go-postgres"
	"github.com/zhiwei-jian/go-rabbitmq/redis"
)

var dbContext compostgres.AppContext

var RedisConfig = &redis.RedisConfig{
	"10.199.196.93:30997",
	"zisefeizhu",
	0,
}

var PostgresConfig = &compostgres.PostgresConfig{
	"10.199.196.93",
	31656,
	"postgres",
	"postgres",
	"k8s",
}

func GetDbContext() *compostgres.AppContext {
	if dbContext.Db != nil {
		return &dbContext
	}

	dbContext, err := compostgres.ConnectDB(PostgresConfig)
	if err != "" {
		return nil
	}

	return dbContext
}
