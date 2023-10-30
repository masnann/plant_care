package config

import (
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
)

type Config struct {
	ServerPort          int
	DBPort              int
	DBHost              string
	DBUser              string
	DBPass              string
	DBName              string
	Secret              string
	CloudinaryCloudName string
	CloudinaryAPIKey    string
	CloudinaryAPISecret string
	Folder              string
}

func InitConfig() *Config {
	var res = new(Config)
	res = loadConfig()

	if res == nil {
		logrus.Fatal("Config : cannot load configuration")
		return nil
	}
	return res

}

func loadConfig() *Config {

	var res = new(Config)
	// err := godotenv.Load(".env")

	// if err != nil {
	// 	logrus.Fatal("Config : cannot load config file", err.Error())
	// 	return nil
	// }

	if value, found := os.LookupEnv("SERVER"); found {
		port, err := strconv.Atoi(value)
		if err != nil {
			logrus.Fatal("Config : invalid server port", err.Error())
			return nil
		}
		res.ServerPort = port
	}

	if value, found := os.LookupEnv("DBPORT"); found {
		port, err := strconv.Atoi(value)
		if err != nil {
			logrus.Fatal("Config : invalid db port", err.Error())
			return nil
		}
		res.DBPort = port

	}
	if value, found := os.LookupEnv("DBHOST"); found {
		res.DBHost = value
	}
	if value, found := os.LookupEnv("DBUSER"); found {
		res.DBUser = value
	}
	if value, found := os.LookupEnv("DBPASS"); found {
		res.DBPass = value
	}
	if value, found := os.LookupEnv("DBNAME"); found {
		res.DBName = value
	}
	if val, found := os.LookupEnv("SECRET"); found {
		res.Secret = val
	}
	if value, found := os.LookupEnv("CLOUDINARY_CLOUD_NAME"); found {
		res.CloudinaryCloudName = value
	}

	if value, found := os.LookupEnv("CLOUDINARY_API_KEY"); found {
		res.CloudinaryAPIKey = value
	}

	if value, found := os.LookupEnv("CLOUDINARY_API_SECRET"); found {
		res.CloudinaryAPISecret = value
	}
	if value, found := os.LookupEnv("FOLDER"); found {
		res.Folder = value
	}
	return res
}
