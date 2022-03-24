package tool

import (
	"bufio"
	"encoding/json"
	"os"
)

type Config struct {
	Server   Server   `json:"server"`
	Sms      Sms      `json:"sms"`
	Database Database `json:"database"`
}

type Server struct {
	ServerName         string `json:"server_name"`
	ServerMode         string `json:"server_mode"`
	ServerHost         string `json:"server_host"`
	ServerPort         int    `json:"server_port"`
	ServerReadTimeout  int    `json:"server_read_timeout"`
	ServerWriteTimeout int    `json:"server_write_timeout"`
}

type Sms struct {
	SignName     string `json:"sign_name"`
	TemplateCode string `json:"template_code"`
	AppKey       string `json:"app_key"`
	AppSectet    string `json:"app_secret"`
	RegionId     string `json:"region_id"`
}

type Database struct {
	Driver   string `json:"driver"`
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Name     string `json:"name"`
	Charset  string `json:"charset"`
	ShowSql  bool   `json:"show_sql"`
}

var _cfg *Config = nil

func GetConfig() *Config {
	return _cfg
}

func ParseConfig() {
	file, err := os.Open("./conf/config.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	decoder := json.NewDecoder(reader)
	if err = decoder.Decode(&_cfg); err != nil {
		panic(err)
	}

}
