package cache

import (
	"fmt"
	"strconv"
	"github.com/go-redis/redis"
	"gopkg.in/ini.v1"
)

var(
	RedisClient *redis.Client 
	RedisDb string
	RedisAddr string
	RedisPw string
	RedisDbName string
)

func init() {
	file, err := ini.Load("./conf/config.ini")
	if err != nil{
		fmt.Println("redis config err", err)
	}
	LoadRedisData(file)
	Redis()
}

func LoadRedisData(file *ini.File){
	RedisDb = file.Section("redis").Key("RedisDb").String() 
	RedisAddr = file.Section("redis").Key("RedisHost").String()
	RedisPw = file.Section("redis").Key("RedisPort").String()
	RedisDbName= file.Section("redis").Key("RedisUser").String()
}

func Redis(){
	db, _ := strconv.ParseUint(RedisDbName, 10, 64)
	client := redis.NewClient(&redis.Options{
		Addr: RedisAddr,
		//password
		DB: int(db),
	})
	_, err := client.Ping().Result()
	if err != nil{
		panic(err)
	}
	RedisClient = client

}




