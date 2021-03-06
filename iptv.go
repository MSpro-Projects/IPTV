package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"

	"./factory"
	"./iptv-server"
	"./version"
)

type (
	// Config : IPTV Configuration file information
	Config struct {
		iptvcfg string
	}
)

var config Config

var iptvFlags = []cli.Flag{
	&cli.StringFlag{
		Name:  "iptvcfg",
		Usage: "config file",
	},
}

func action(c *cli.Context) error {
	if err := initialize(c); err != nil {
		fmt.Print(err)
	}
	if err := start(); err != nil {
		fmt.Print(err)
	}
	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "IPTV"
	app.Usage = "-free5gccfg common configuration file -iptvcfg iptv configuration file"
	app.Action = action
	fmt.Println(app.Name)
	app.Flags = iptvFlags
	if err := app.Run(os.Args); err != nil {
		fmt.Printf("Error args: %v", err)
	}
}

func initialize(c *cli.Context) error {
	config = Config{
		iptvcfg: c.String("iptvcfg"),
	}

	if config.iptvcfg != "" {
		factory.InitConfigFactory(config.iptvcfg)
	} else {
		DefaultIptvConfigPath := "./iptvcfg.conf"
		factory.InitConfigFactory(DefaultIptvConfigPath)
	}
	return nil
}

func start() error {
	// Run gin Server
	server := iptvserver.Server{}
	server.IptvServerIpv4Port = factory.IptvConfig.Configuration.IPTVServer.ServerAddr
	server.Channels = factory.IptvConfig.Configuration.IPTVServer.Channel
	server.CacheFolder = factory.IptvConfig.Configuration.IPTVServer.CacheFolder
	server.WebClient = factory.IptvConfig.Configuration.IPTVServer.WebClientFolder
	server.Version = version.GetVersion()
	server.Start()
	// Run AF
	return nil
}
