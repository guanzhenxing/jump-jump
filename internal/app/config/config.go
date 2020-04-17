package config

import (
	"github.com/go-redis/redis"
	"github.com/jwma/reborn"
	"time"
)

var config *reborn.Reborn

func getDefaultConfig() *reborn.Config {
	d := reborn.NewConfig()
	d.SetValue("landingHosts", []string{"http://127.0.0.1:8081/"})
	d.SetValue("idMinimumLength", 2)
	d.SetValue("idLength", 6)
	d.SetValue("idMaximumLength", 10)

	return d
}

func GetConfig() *reborn.Reborn {
	return config
}

func SetupConfig(rdb *redis.Client) error {
	var err error
	config, err = reborn.NewWithDefaults(rdb, "j2config", getDefaultConfig())
	if err != nil {
		return err
	}
	config.SetAutoReloadDuration(time.Second * 30)
	config.StartAutoReload()

	return nil
}
