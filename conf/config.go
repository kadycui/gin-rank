package conf

import (
	"fmt"
	"log"

	"gopkg.in/ini.v1"
)

var (
	AppMode  string
	HttpPort string

	MySQL_Host string
	MySQL_Port string
	MySQL_User string
	MySQL_Password   string
	MySQL_Db string

	Redis_Host string
	Redis_Password   string
	Redis_Port string
	Redis_Db   string
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
	MySQL_User = file.Section("mysql").Key("user").String()
	MySQL_Password = file.Section("mysql").Key("password").String()
	MySQL_Host = file.Section("mysql").Key("host").String()
	MySQL_Port = file.Section("mysql").Key("port").String()
	MySQL_Db = file.Section("mysql").Key("dbname").String()

	fmt.Println(MySQL_User, MySQL_Password, MySQL_Host, MySQL_Port, MySQL_Db)
}

func LoadRedis(file *ini.File) {
	Redis_Host = file.Section("redis").Key("host").String()
	Redis_Port = file.Section("redis").Key("port").String()
	Redis_Password = file.Section("redis").Key("password").String()
	Redis_Db = file.Section("redis").Key("db").String()

	fmt.Println(Redis_Host, Redis_Port, Redis_Password, Redis_Db)
}
