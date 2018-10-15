package config

import (
	"strings"

	"github.com/spf13/viper"

	"github.com/Eric-GreenComb/recharge-linked/bean"
)

// Server Server Config
var Server bean.ServerConfig

// Dgraph Dgraph配置
var Dgraph bean.DgraphConfig

func init() {
	readConfig()
	initConfig()
}

func readConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	viper.ReadInConfig()
}

func initConfig() {
	Server.Port = strings.Split(viper.GetString("server.port"), ",")
	Server.Mode = viper.GetString("server.mode")
	Server.IPs = strings.Split(viper.GetString("server.ip"), ",")

	Dgraph.Host = viper.GetString("dgraph.host")
}

// HasInIPs HasInIPs
func HasInIPs(ip string) bool {
	for _, _ip := range Server.IPs {
		if _ip == ip {
			return true
		}
	}
	return false
}
