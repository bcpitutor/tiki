package appconfig

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	ViperConf       *viper.Viper
	SelectedProfile string
	JsonOutput      bool
	Debug           bool
}

var AppConfig *Config

func InitConfig(tikidir string) error {
	v := viper.New()
	v.SetConfigName("tiki")
	v.SetConfigType("ini")
	//viper.AddConfigPath(tikidir) // TODO:
	v.AddConfigPath(".")

	err := v.ReadInConfig()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return err
	}

	AppConfig = &Config{
		ViperConf: v,
	}

	return nil
}
