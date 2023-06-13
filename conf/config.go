package conf

import (
	"fmt"
	"log"

	"gopkg.in/ini.v1"
)

var (
	AppMode  string
	HttpPort string

	MyHost string
	MyPort string
	MyUser string
	MyPw   string
	MyName string

	RedisHost string
	RedisPw   string
	RedisPort string
	RedisDb   string
)

func init() {
	file, err := ini.Load("conf/config.ini")
	if err != nil {
		log.Println("配置文件读取错误，请检查文件路径:", err)
		panic(err)
	}

	LoadServer(file)
	LoadMysqlData(file)
	LoadRedis(file)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("service").Key("AppMode").String()
	HttpPort = file.Section("service").Key("HttpPort").String()
}

func LoadMysqlData(file *ini.File) {
	MyUser = file.Section("mysql").Key("MyUser").String()
	MyPw = file.Section("mysql").Key("MyPw").String()
	MyHost = file.Section("mysql").Key("MyHost").String()
	MyPort = file.Section("mysql").Key("MyPort").String()
	MyName = file.Section("mysql").Key("MyName").String()
}

func LoadRedis(file *ini.File) {
	RedisHost = file.Section("redis").Key("RedisHost").String()
	RedisPort = file.Section("redis").Key("RedisPort").String()
	RedisPw = file.Section("redis").Key("RedisPw").String()
	RedisDb = file.Section("redis").Key("RedisDb").String()

	fmt.Println(RedisHost, RedisPort, RedisPw, RedisDb)
}
