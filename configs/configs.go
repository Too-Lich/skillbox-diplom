package configs

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

type Config struct {
	SMS       string `json:"SMS"`
	VoiceCall string `json:"voice_call"`
	Billing   string `json:"billing"`
	MMS       string `json:"MMS"`
	Support   string `json:"support"`
	Incident  string `json:"incident"`
	Email     string `json:"email"`
}

func GetConfig() Config {
	dir, err := filepath.Abs("")
	if err != nil {
		log.Print("Не могу найти файл настроек: ", err)
	}
	pathConfigFile := filepath.Join(dir, "configs", "config.json")
	log.Print(pathConfigFile)

	confByte, err := os.ReadFile(pathConfigFile)
	if err != nil {
		log.Fatal("Не могу получить настройки:", err)
	}

	var config Config
	if err = json.Unmarshal(confByte, &config); err != nil {
		log.Fatalln("Не могу получить настройки:", err)
	}

	return config
}
