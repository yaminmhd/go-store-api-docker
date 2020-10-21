package main

import (
	"github.com/urfave/cli"
	"github.com/yaminmhd/go-hardware-store/appcontext"
	"github.com/yaminmhd/go-hardware-store/config"
	"github.com/yaminmhd/go-hardware-store/log"
	"github.com/yaminmhd/go-hardware-store/server"
	"os"
)

func main(){
	config.Load()
	log.SetupLogger()
	appcontext.Initiate()
	appcontext.CreateTables()
	startApp()
}


func startApp() {
	app := cli.NewApp()
	app.Name = "Hardware store in GO"
	app.Version = "0.0.1"
	app.Commands = []cli.Command{
		{
			Name:        "start",
			Description: "Start HTTP api server",
			Action: func(c *cli.Context) error {
				server.StartServer()
				return nil
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
