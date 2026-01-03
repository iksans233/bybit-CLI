package main

import (
	"bybit/config"
	"bybit/rootcmd"
	"os"
)

func main() {
	config.Init()

	err := root.Cmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
