package config

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"gopkg.in/natefinch/lumberjack.v2"
)

var conf Config

func InitConfig(config string) Config {

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	log.SetOutput(&lumberjack.Logger{
		Filename:   "logs/app.log",
		MaxSize:    2, // megabytes
		MaxBackups: 30,
		MaxAge:     40, // days
		Compress:   true,
	})
	log.Println("-------- * ------- Starting Logging -------- * -------")

	data, err := ioutil.ReadFile(config)
	if err != nil {
		log.Fatalln("✗ config.init ioutil.ReadFile error: ", err.Error())
	}
	conf = Config{}
	if err := json.Unmarshal(data, &conf); err != nil {
		log.Fatalln("✗ config.init json.Unmarshal error: ", err.Error())
	}

	return conf
}

type Config struct {
	Name     string    `json:"name"`
	Database Database  `json:"database"`
	App      _App      `json:"app"`
	Services _Services `json:"services"`
}

// Database ...
type Database struct {
	Host     string `json:"host"`
	Port     uint16 `json:"port"`
	User     string `json:"user"`
	Password string `json:"pass"`
	Name     string `json:"name"`
}

type _App struct {
	Name        string `json:"name"`
	Port        string `json:"port"`
	BaseAddress string `json:"baseAddress"`
	Version     string `json:"version"`
}
type _Services struct {
	FileStorage  _Path  `json:"fileStorage"`
	ImageStorage _Path  `json:"imageStorage"`
	CrmAPIUrl    string `json:"crmApiUrl"`
	CrmToken     string `json:"crm_token"`
	Sms          _Sms   `json:"sms"`
}

type _Path struct {
	Path string `json:"path"`
}

type _Sms struct {
	Address string `json:"address"`
	Token   string `json:"token"`
}
