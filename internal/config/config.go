package config

import (
	"flag"
	"os"

	"github.com/PritOriginal/zero-agency-test/pkg/logger"
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env    logger.Environment `yaml:"env" env:"ENV" env-default:"local"`
	OpenAI OpenAIConfig       `yaml:"open_ai" `
}

type OpenAIConfig struct {
	Model  string `yaml:"model" env:"OPENAI_MODEL"`
	URL    string `yaml:"url" env:"OPENAI_URL"`
	ApiKey string `yaml:"api_key" env:"OPENAI_API_KEY"`
}

func MustLoad() *Config {
	configPath := fetchConfigPath()
	if configPath == "" {
		panic("config path is empty")
	}

	return MustLoadPath(configPath)
}

func MustLoadPath(configPath string) *Config {
	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic("cannot read config: " + err.Error())
	}

	return &cfg
}

func fetchConfigPath() string {
	var res string

	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}
