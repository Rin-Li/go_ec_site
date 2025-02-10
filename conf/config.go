package conf

import (
	"fmt"
	"gin-mall-tmp/dao"
	"log"
	"strings"

	"gopkg.in/ini.v1"
)

var (
	AppModel string
	HttpPort string

	Db          string
	DbHost      string
	DbPort      string
	DbUser      string
	DbPassword  string
	DbName      string
	RedisDb     string
	RedisAddr   string
	RedisPw     string
	RedisDbName string
	ValidEmail  string
	SmtpHost    string
	SmtpEmail   string
	SmtpPass    string
	Host        string
	ProductPath string
	AvaterPath  string
)

func Init() {
	//本地读取环境变量
	file, err := ini.Load("./conf/config.ini")
	if err != nil {
		log.Fatal("Error loading config file:", err)
	}

	LoadServer(file)
	LoadMySql(file)
	LoadRedis(file)
	LoadEmail(file)
	LoadPhotoPath(file)

	// mysql 连接字符串拼接
	pathRead := strings.Join([]string{
		DbUser, ":", DbPassword, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8mb4&parseTime=true",
	}, "")
	pathWrite := strings.Join([]string{
		DbUser, ":", DbPassword, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8mb4&parseTime=true",
	}, "")

	// 打印数据库连接字符串
	fmt.Println("Database connection string (Read):", pathRead)
	fmt.Println("Database connection string (Write):", pathWrite)

	dao.Database(pathRead, pathWrite)
}

func LoadServer(file *ini.File) {
	AppModel = file.Section("service").Key("AppModel").String() //在service 的section下拿到Appmode的值，然后String()化
	HttpPort = file.Section("service").Key("HttpPort").String()
}

func LoadMySql(file *ini.File) {
	Db = file.Section("mysql").Key("Db").String() //在service 的section下拿到Appmode的值，然后String()化
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbPort = file.Section("mysql").Key("DbPort").String()
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbPassword = file.Section("mysql").Key("DbPassword").String()
	DbName = file.Section("mysql").Key("DbName").String()
}

func LoadEmail(file *ini.File) {
	ValidEmail = file.Section("email").Key("ValidEmail").String() //在service 的section下拿到Appmode的值，然后String()化
	SmtpHost = file.Section("email").Key("SmtpHost").String()
	SmtpEmail = file.Section("email").Key("SmtpEmail").String() //在service 的section下拿到Appmode的值，然后String()化
	SmtpPass = file.Section("email").Key("SmtpPass").String()

}

func LoadRedis(file *ini.File) {
	RedisDb = file.Section("redis").Key("RedisDb").String() //在service 的section下拿到Appmode的值，然后String()化
	RedisAddr = file.Section("redis").Key("RedisAddr").String()
	RedisPw = file.Section("redis").Key("RedisPw").String()
	RedisDbName = file.Section("redis").Key("RedisDbName").String()
}

func LoadPhotoPath(file *ini.File) {
	Host = file.Section("path").Key("Host").String() //在service 的section下拿到Appmode的值，然后String()化
	ProductPath = file.Section("path").Key("ProductPath").String()
	AvaterPath = file.Section("path").Key("AvaterPath").String()
}
