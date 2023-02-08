package providers

// 使用插件：https://github.com/go-redis/redis

import (
	"context"
	"ginvel/common/helper"
	"ginvel/internal/reader/framework_toml"
	"github.com/go-redis/redis/v8"
	"log"
)

type InitRedis struct{}

var RedisDb *redis.Client

func (initRedis *InitRedis) InitRedis1() {
	log.Println("尝试连接GoRedis1...")

	var rdbConfig map[string]string = framework_toml.GetRedisConfig1()
	var address = rdbConfig["Host"] + ":" + rdbConfig["Port"]

	defer func() {
		if r := recover(); r != nil {
			log.Println("GoRedis1初始化出现问题，已经跳过。。。")
		}
	}()

	RedisDb = redis.NewClient(&redis.Options{ // 连接服务
		Addr:     address,                                  // string
		Password: rdbConfig["Password"],                    // string
		DB:       int(helper.StringToInt(rdbConfig["DB"])), // int
	})
	RedisPong, RedisErr := RedisDb.Ping(context.Background()).Result() // 心跳
	if RedisErr != nil {
		log.Println("Redis1服务未运行。。。", RedisPong, RedisErr)
		log.Println("Redis常用命令：\n" +
			" 启动：src/redis-server \n" +
			" 进入命令行：src/redis-cli \n" +
			" 关闭安全模式：CONFIG SET protected-mode no \n" +
			" 重置密码：config set requirepass [密码]\n")
		//os.Exit(200)
	} else {
		log.Println("GoRedis1已连接 >>> ")
	}

	return
}

var RedisDb2 *redis.Client

func (initRedis *InitRedis) InitRedis2() {
	log.Println("尝试连接GoRedis2...")

	var rdbConfig map[string]string = framework_toml.GetRedisConfig2()
	var address = rdbConfig["Host"] + ":" + rdbConfig["Port"]

	defer func() {
		if r := recover(); r != nil {
			log.Println("GoRedis2初始化出现问题，已经跳过。。。")
		}
	}()

	RedisDb = redis.NewClient(&redis.Options{ // 连接服务
		Addr:     address,                                  // string
		Password: rdbConfig["Password"],                    // string
		DB:       int(helper.StringToInt(rdbConfig["DB"])), // int
	})
	RedisPong, RedisErr := RedisDb.Ping(context.Background()).Result() // 心跳
	if RedisErr != nil {
		log.Println("Redis2服务未运行。。。", RedisPong, RedisErr)
		log.Println("Redis常用命令：\n" +
			" 启动：src/redis-server \n" +
			" 进入命令行：src/redis-cli \n" +
			" 关闭安全模式：CONFIG SET protected-mode no \n" +
			" 重置密码：config set requirepass [密码]\n")
		//os.Exit(200)
	} else {
		log.Println("GoRedis2已连接 >>> ")
	}

	return
}
