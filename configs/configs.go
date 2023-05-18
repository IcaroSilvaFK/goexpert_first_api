package configs

import (
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)

var cfg *conf

const (
	configFile      = "configs"
	configExtension = "toml"
)

type conf struct {
	DBDriver      string `mapstructure:"db_driver"`
	DBHost        string `mapstructure:"db_host"`
	DBPort        string `mapstructure:"db_port"`
	DBUser        string `mapstructure:"db_user"`
	DBPassword    string `mapstructure:"db_pass"`
	DBName        string `mapstructure:"db_name"`
	WebServerPort string `mapstructure:"port"`
	JWTSecret     string `mapstructure:"jwt_secret"`
	JWTExpiresIn  int    `mapstructure:"jwt_expires_in"`
	TokenAuth     *jwtauth.JWTAuth
}

func LoadConfig(path string) (*conf, error) {

	viper.SetConfigName(configFile)
	viper.SetConfigType(configExtension)
	viper.AddConfigPath(path)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Println("Config file changed:", e.Name)
	})

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	cfg.TokenAuth = jwtauth.New("HS256", []byte(cfg.JWTSecret), nil)

	return cfg, nil
}
