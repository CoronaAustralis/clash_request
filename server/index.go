package server

import (
	"clash_request/cmd/flags"
	"clash_request/config"
	"clash_request/utils"
	"io"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

var Config *config.Config

func init() {
	if flags.Config == "" {
		configPath := "./config.json"
		if !utils.Exists(configPath) {
			log.Infoln("config file not exists, creating default config file")
			_, err := utils.CreateNestedFile(configPath)
			if err != nil {
				log.Fatalf("failed to create config file: %+v", err)
			}
			Config = config.DefaultConfig()
			if !utils.WriteJsonToFile(configPath, Config) {
				log.Fatalf("failed to create default config file")
			}
		} else {
			configBytes, err := os.ReadFile(configPath)
			if err != nil {
				log.Fatalf("reading config file error: %+v", err)
			}
			Config = &config.Config{}
			err = utils.Json.Unmarshal(configBytes, Config)
			if err != nil {
				log.Fatalf("load config error: %+v", err)
			}
			// update config.json struct
			confBody, err := utils.Json.MarshalIndent(Config, "", "  ")
			if err != nil {
				log.Fatalf("marshal config error: %+v", err)
			}
			err = os.WriteFile(configPath, confBody, 0777)
			if err != nil {
				log.Fatalf("update config struct error: %+v", err)
			}
		}
	} else {
		configPath := flags.Config
		configBytes, err := os.ReadFile(configPath)
		if err != nil {
			log.Fatalf("reading config file error: %+v", err)
		}
		Config = &config.Config{}
		err = utils.Json.Unmarshal(configBytes, Config)
		if err != nil {
			log.Fatalf("load config error: %+v", err)
		}
	}

	if Config.RequestUrl == "" || Config.RequestToken == "" {
		log.Panicln("no url or token")
	}

	SetLog()
}

func SetLog() {
	formatter := log.TextFormatter{
		ForceColors:               true,
		EnvironmentOverrideColors: true,
		DisableColors:             true,
		TimestampFormat:           "2006-01-02 15:04:05",
		FullTimestamp:             true,
	}
	log.SetFormatter(&formatter)
	if !utils.Exists(Config.LogPath) {
		dir := filepath.Dir(Config.LogPath)
		err := os.MkdirAll(dir, 0777)
		if err != nil {
			log.Panicln(err)
		}
	}
	w2, err := os.OpenFile(Config.LogPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		log.Panicln(err)
	}
	log.SetOutput(io.MultiWriter(os.Stdout, w2))
}

func ExecuteRequest() {
	for i := 0; i < 5; i++ {
		if GetSubscribePath() {
			break
		}
	}
	for i := 0; i < 5; i++ {
		if GetSubscribeFile() {
			break
		}
	}
}

func Server() {
	ExecuteRequest()
	SetInterval(ExecuteRequest)
	ginServer()
}
