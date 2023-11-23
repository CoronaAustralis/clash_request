package config

import "clash_request/cmd/flags"

type Config struct {
	RequestUrl       string `json:"requestUrl"`
	RequestToken     string `json:"requestToken"`
	RequestInterval  string `json:"requestInterval"`
	SubsrciptionPath string `json:"subsrciptionPath"`
	Port             int    `json:"port"`
	LogPath          string `json:"logPath"`
}

func DefaultConfig() *Config {
	return &Config{
		RequestUrl:       "http://127.0.0.1:7777/get/",
		RequestToken:     "",
		RequestInterval:  "12h",
		SubsrciptionPath: "./data/sub/sub.yaml",
		Port:             flags.Port,
		LogPath:          "./data/log/server.log",
	}
}
