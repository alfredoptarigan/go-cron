package helper

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"reflect"
	"regexp"
)

func SetConfig(key string, defaultVal string) string {
	viper.SetConfigType("env")

	// set config path to root directory
	viper.AddConfigPath(SetRootPath() + "/.")
	viper.SetConfigName("app")

	viper.SetDefault(key, defaultVal)
	viper.WatchConfig()
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	return viper.GetString(key)
}
func isNilFixed(i interface{}) bool {
	if i == nil {
		return true
	}
	switch reflect.TypeOf(i).Kind() {
	case reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan, reflect.Slice:
		return reflect.ValueOf(i).IsNil()
	}
	return false
}

func GetConfig(key string) string {
	return SetConfig(key, "")
}

func SetRootPath() string {
	_, ok := os.LookupEnv("ENVIRONMENT")
	mainPath := "build"
	if ok {
		mainPath = "build"
	} else {
		mainPath = "crongolang"
	}

	projectName := regexp.MustCompile(`^(.*` + mainPath + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	return string(rootPath)
}
