package config

import (
	"encoding/json"
	"github.com/spf13/viper"
	"log"
	"regexp"
	"strings"
)

var config Config = Config{
	Debug: &DebugConfig{
		Enable:  false,
		Verbose: false,
	},
	Event: &EventConfig{},
}

func init() {
	log.Println("init config...")

	instance := viper.New()

	// only for dev
	instance.AddConfigPath("/etc/twitter-oracle-backend/")
	instance.AddConfigPath(".")

	instance.SetConfigType("yaml")
	instance.SetConfigName("config.yaml")

	instance.SetEnvPrefix("tob")
	instance.AutomaticEnv()
	instance.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := instance.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			panic(err)
		}
	}

	err = instance.Unmarshal(&config)
	if err != nil {
		panic(err)
	}

	priKeyRegex := regexp.MustCompile("[0-9a-hA-H]{64}")
	configJson, err := json.Marshal(config)
	if err != nil {
		return
	}

	log.Printf("dump config: %s", priKeyRegex.ReplaceAll(configJson, []byte("[PRIVATE KEY]")))
}

func Get() *Config {
	return &config
}
