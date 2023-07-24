package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Dbhost     string `mapstructure: "postgres_host"`
	Dbusername string `mapstructure: "postgres_user"`
	Dbpassword string `mapstructure: "postgres_password"`
	Dbname     string `mapstructure: "postgres_db"`
	Dbport     string `mapstructure: "postgres_port"`
}

func LoadConfig(path string) (config Config, err error) {

	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()
	viper.BindEnv()

	err = viper.ReadInConfig()
	// fmt.Print(viper.AllSettings()) //map[postgres_db:test postgres_host:localhost postgres_password:postgres postgres_port:5432 postgres_user:postgres]
	if err != nil {
		fmt.Print("config file is not found")
		return
	}
	// viper.BindEnv("Dbhost", "postgres_host")
	// viper.BindEnv("Dbusername", "postgres_user")
	// viper.BindEnv("Dbpassword", "postgres_password")
	// viper.BindEnv("Dbname", "postgres_db")
	// viper.BindEnv("Dbport", "postgres_port")
	// viper.WriteConfigAs("abctest")

	// The current viper implementation does not allow to viper.Unmarshal environment variables into structs without binding them first (using viper.BindEnv). This behavior results in an "empty" struct (i.e no values filled) if no configuration file is provided with matching keys.

	err = viper.Unmarshal(&config) //masalah anjing config is {    }, tapi di viper.Get bisa

	// fmt.Print(err) //KNP return nil????

	return
}
