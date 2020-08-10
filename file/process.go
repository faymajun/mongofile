package file

import (
	"fmt"
	"github.com/urfave/cli"
	"time"
)

func Action() func(*cli.Context) error {
	return func(ctx *cli.Context) error {
		dict := ctx.String("dictionary")
		fmt.Println(dict)
		defaultConfig.scanFrequency = time.Duration(ctx.Int("scanFrequency")) * time.Second
		return nil
	}
}
