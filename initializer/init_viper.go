package initializer

import (
	"auto-manager/global"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func InitViper() *viper.Viper {
	v := viper.New()
	v.SetConfigFile("config.yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf(" *** FATAL ERROR *** \n%s", err))
	}
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println(" *** config file changed ***")
		if err := v.Unmarshal(&global.AM_CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(&global.AM_CONFIG); err != nil {
		fmt.Println(err)
	}
	return v
}
