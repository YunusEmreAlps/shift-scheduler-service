package config

import (
	"os"
	"path/filepath"

	"shift-scheduler-service/pkg/logger"

	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/viper"
)

type Config struct {
	App     App     `mapstructure:"app"`
	Auth    Auth    `mapstructure:"auth"`
	DB      DB      `mapstructure:"db"`
	Cache   Cache   `mapstructure:"cache"`
	Broker  Broker  `mapstructure:"broker"`
	Cookie  Cookie  `mapstructure:"cookie"`
	Session Session `mapstructure:"session"`
	Metric  Metric  `mapstructure:"metric"`
	Logger  Logger  `mapstructure:"logger"`
	Jaeger  Jaeger  `mapstructure:"jaeger"`
	Cdn     Cdn     `mapstructure:"cdn"`
}

type App struct {
	Mode    string `mapstructure:"mode"`
	Port    string `mapstructure:"port"`
	Version string `mapstructure:"version"`
	Name    string `mapstructure:"name"`
}

type Auth struct {
	JwtPub string `mapstructure:"jwt_pub"`
}

type DB struct {
	Url string `mapstructure:"url"`
}

type Cache struct {
	Url string `mapstructure:"url"`
}

type Broker struct {
	Url           string `mapstructure:"url"`
	ConsumerGroup string `mapstructure:"consumer_group"`
	Topic         string `mapstructure:"topic"`
}

type Cookie struct {
	Name     string `mapstructure:"name"`
	MaxAge   int    `mapstructure:"max_age"`
	Secure   bool   `mapstructure:"secure"`
	HTTPOnly bool   `mapstructure:"http_only"`
}

type Session struct {
	Name   string `mapstructure:"name"`
	Prefix string `mapstructure:"prefix"`
	Expire int    `mapstructure:"expire"`
}

type Metric struct {
	Url     string `mapstructure:"url"`
	Service string `mapstructure:"service"`
}

type Logger struct {
	Development       bool   `mapstructure:"development"`
	DisableCaller     bool   `mapstructure:"disable_caller"`
	DisableStacktrace bool   `mapstructure:"disable_stacktrace"`
	Encoding          string `mapstructure:"encoding"`
	Level             string `mapstructure:"level"`
}

type Jaeger struct {
	Host        string `mapstructure:"host"`
	ServiceName string `mapstructure:"service_name"`
	LogSpans    bool   `mapstructure:"log_spans"`
}

type Cdn struct {
	Endpoint  string `mapstructure:"endpoint"`
	Region    string `mapstructure:"region"`
	Bucket    string `mapstructure:"bucket"`
	AccessKey string `mapstructure:"access_key"`
	SecretKey string `mapstructure:"secret_key"`
}

/*type config struct {
	App struct {
		Mode    string `mapstructure:"mode"`
		Port    string `mapstructure:"port"`
		Version string `mapstructure:"version"`
		Name    string `mapstructure:"name"`
	} `mapstructure:"app"`

	Auth struct {
		JwtPub string `mapstructure:"jwt_pub"`
	} `mapstructure:"auth"`

	DB struct {
		Url string `mapstructure:"url"`
	} `mapstructure:"db"`

	Cache struct {
		Url string `mapstructure:"url"`
	} `mapstructure:"cache"`

	Broker struct {
		Url           string `mapstructure:"url"`
		ConsumerGroup string `mapstructure:"consumer_group"`
		Topic         string `mapstructure:"topic"`
	} `mapstructure:"broker"`

	Metrics struct {
		Url     string `mapstructure:"url"`
		Service string `mapstructure:"service"`
	} `mapstructure:"metrics"`

	Logger struct {
		Development       bool   `mapstructure:"development"`
		DisableCaller     bool   `mapstructure:"disable_caller"`
		DisableStacktrace bool   `mapstructure:"disable_stacktrace"`
		Encoding          string `mapstructure:"encoding"`
		Level             string `mapstructure:"level"`
	} `mapstructure:"logger"`

	Cdn struct {
		Endpoint  string `mapstructure:"endpoint"`
		Region    string `mapstructure:"region"`
		Bucket    string `mapstructure:"bucket"`
		AccessKey string `mapstructure:"access_key"`
		SecretKey string `mapstructure:"secret_key"`
	} `mapstructure:"cdn"`
}*/

var C Config

func ReadConfig(processCwdir string) {
	Config := &C
	viper.SetConfigName(".env")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(filepath.Join(processCwdir, "config"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		// fmt.Println("Cannot read config file:", err)
		logger.CLogger.Info("Cannot read config file:", err)
	}

	if err := viper.Unmarshal(&Config); err != nil {
		logger.CLogger.Info("Cannot unmarshal config file:", err)
		os.Exit(1)
	}

	spew.Dump(C)
}
