package config

import (
	"log"
	"os"
	"strconv"

	"github.com/spf13/viper"
)

var JWT_SECRRET = ""

type AppConfig struct {
	DBUsername     string
	DBPassword     string
	DBHost         string
	DBPort         int
	DBName         string
	jwtKey         string
	CLOUD_NAME     string
	KEY_API        string
	KEY_API_SECRET string
}

func InitConfig() *AppConfig {
	return ReadENV()
}

func ReadENV() *AppConfig {
	app := AppConfig{}
	isRead := true

	if val, found := os.LookupEnv("JWTSECRET"); found {
		app.jwtKey = val
		isRead = false
	}
	if val, found := os.LookupEnv("DBUSER"); found {
		app.DBUsername = val
		isRead = false
	}
	if val, found := os.LookupEnv("DBPASS"); found {
		app.DBPassword = val
		isRead = false
	}
	if val, found := os.LookupEnv("DBHOST"); found {
		app.DBHost = val
		isRead = false
	}
	if val, found := os.LookupEnv("DBPORT"); found {
		conv, _ := strconv.Atoi(val)
		app.DBPort = conv
		isRead = false
	}
	if val, found := os.LookupEnv("DBNAME"); found {
		app.DBName = val
		isRead = false
	}
	if val, found := os.LookupEnv("CLOUDINARY_CLOUD_NAME"); found {
		app.CLOUD_NAME = val
		isRead = false
	}
	if val, found := os.LookupEnv("CLOUDINARY_KEY_API"); found {
		app.KEY_API = val
		isRead = false
	}
	if val, found := os.LookupEnv("CLOUDINARY_KEY_API_SECRET"); found {
		app.KEY_API_SECRET = val
		isRead = false
	}

	if isRead {
		viper.AddConfigPath(".")
		viper.SetConfigName("local")
		// viper.SetConfigName("server")
		viper.SetConfigType("env")

		err := viper.ReadInConfig()
		if err != nil {
			log.Println("error read config: ", err.Error())
			return nil
		}
		app.jwtKey = viper.Get("JWTSECRET").(string)
		app.DBUsername = viper.Get("DBUSER").(string)
		app.DBPassword = viper.Get("DBPASS").(string)
		app.DBHost = viper.Get("DBHOST").(string)
		app.DBPort, _ = strconv.Atoi(viper.Get("DBPORT").(string))
		app.DBName = viper.Get("DBNAME").(string)
	}
	JWT_SECRRET = app.jwtKey
	return &app
}
