package main

import (
	"fmt"
	"time-tracker/internal/config"
)

func main() {
	cfg, err := config.LoadCfg()
	if err != nil {
		panic("time-tracker: failed to read config: " + err.Error())
	}

	fmt.Print(cfg)

	// TODO:логирование

	// TODO:запуск сервера

	// TODO: безопасное окончание программы
}
