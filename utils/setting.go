package utils

import (
	"fmt"
	"os"

	"gopkg.in/ini.v1"
)

var (
	AppMode  string
	HttpPort string
	JwtKey   string

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string

	AccessKey string
	SecretKey string
	Bucket    string
	SeverName string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	LoadDatabase(file)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":3000")
	JwtKey = file.Section("server").Key("JwtKey").MustString("qwertyuiop312")
}

func LoadDatabase(file *ini.File) {
	Db = file.Section("database").Key("Db").MustString("ginblog")
	DbHost = file.Section("database").Key("DbHost").MustString("127.0.0.1")
	DbPort = file.Section("database").Key("DbPort").MustString("3303")
	DbUser = file.Section("database").Key("DbUser").MustString("ginblog")
	DbPassWord = file.Section("database").Key("DbPassWord").MustString("ginblog@123.com")
	DbName = file.Section("database").Key("DbName").MustString("ginblog")
}

func LoadUploadFile(file *ini.File) {
	AccessKey = file.Section("oss").Key("AccessKey").MustString("")
	SecretKey = file.Section("oss").Key("SecretKey").MustString("")
	Bucket = file.Section("oss").Key("Bucket").MustString("ginblog")
	SeverName = file.Section("oss").Key("SeverName").MustString("")
}
