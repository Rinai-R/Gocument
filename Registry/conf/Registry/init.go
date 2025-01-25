package conf

import (
	"github.com/Rinai-R/Gocument/Logger"
	"github.com/Rinai-R/Gocument/Registry/Nacos"
	"github.com/spf13/viper"
)

type registry struct {
	Nacos.Nacos
}

var Registry *registry

func init() {
	viper.SetConfigType("yaml")
	viper.SetConfigName("registry")
	viper.SetConfigFile("./Registry/conf/Registry/registry.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		Logger.Logger.Panic("Viper: " + err.Error())
	}
	err = viper.Unmarshal(&Registry)
	if err != nil {
		Logger.Logger.Panic("Viper: " + err.Error())
	}

	Logger.Logger.Debug("Viper: Registry OK")

}
