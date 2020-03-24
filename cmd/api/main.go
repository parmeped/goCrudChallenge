package main

import (
	"flag"

	"github.com/goCrudChallenge/pkg/api"

	"github.com/goCrudChallenge/pkg/utl/config"
)

func main() {

	cfgPath := flag.String("p", "./conf.local.yaml", "Path to config file")
	flag.Parse()

	cfg, err := config.Load(*cfgPath)
	checkErr(err)

	checkErr(api.Start(cfg))
}

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
