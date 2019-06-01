// +build go1.8

package otoware_go

import (
	"os"

	"github.com/kkiyama117/otoware-go/cmd/otowari"
	"github.com/urfave/cli"

	"github.com/kkiyama117/otoware-go/pkg/version"
)

const APP_VER = version.VERSION

func init() {
	// setting.AppVer = APP_VER
}

func main() {
	app := cli.NewApp()
	app.Name = "Gogs"
	app.Usage = "A painless self-hosted Git service"
	app.Version = APP_VER
	app.Commands = []cli.Command{
		otowari.Otowari,
	}
	app.Run(os.Args)
}
