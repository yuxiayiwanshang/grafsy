package main

import (
	"flag"
	"fmt"
	grafsy "github.com/leoleovich/grafsy/grafsy"
	"os"
	"sync"
)

func main() {
	var configFile string
	flag.StringVar(&configFile, "c", "/etc/grafsy/grafsy.toml", "Path to config file.")
	flag.Parse()

	var conf grafsy.Config
	err := conf.LoadConfig(configFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	lc, err := conf.GenerateLocalConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	mon := &grafsy.Monitoring{
		Conf: &conf,
		Lc:   lc,
	}

	cli := grafsy.Client{
		&conf,
		lc,
		mon,
	}

	srv := grafsy.Server{
		&conf,
		lc,
		mon,
	}

	var wg sync.WaitGroup
	go srv.Run()
	go cli.Run()
	go mon.Run()

	wg.Add(1)
	wg.Wait()
}
