package conf

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Database struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Pass     string `yaml:"password"`
	Database string `yaml:"database"`
	SSLMode  string `yaml:"sslmode"`
	Timezone string `yaml:"timezone"`
}

type Log struct {
	MaxSize int `yaml:"max_size"`
}

type GinConfig struct {
	GinMode string `yaml:"gin_mode"`
	SSLMode string `yaml:"sslMode"`
	GinHost string `yaml:"host"`
}

type GitHubAuth struct {
	ClientID     string `yaml:"client_id"`
	ClientSecret string `yaml:"client_secret"`
}

type RedisConf struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

type Config struct {
	Database   Database   `yaml:"database"`
	Log        Log        `yaml:"log"`
	GinConf    GinConfig  `yaml:"gin"`
	GitHubAuth GitHubAuth `yaml:"github_auth"`
	RedisConf  RedisConf  `yaml:"redis"`
}

var CFG *Config

func InitConfig() {
	if cfgFile, err := os.ReadFile("conf/config.yaml"); err != nil {
		panic(err)
	} else {
		if err = yaml.Unmarshal(cfgFile, &CFG); err != nil {
			panic(err)
		}
	}
}
