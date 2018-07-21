package config

import (
	"github.com/damnever/cc"
)

func GetConfig() *cc.Config {
	var c, _ = cc.NewConfigFromFile("./config/app.yml")

	return c
}
