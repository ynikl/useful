package internal

import (
	"ft/internal/repository/ft"

	"github.com/spf13/viper"
)

func Init(path string, name string) error {
	viper.SetConfigName(name)
	viper.AddConfigPath(path)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	ft.Init(viper.GetString("mysql.dsn"))
	return nil
}
