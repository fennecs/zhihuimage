package main

import (
	"github.com/urfave/cli"
	"os"
	"fmt"
	"math"

	"github.com/jinxZz/zhihuimage/service"
	"github.com/jinxZz/zhihuimage/util"
)

func main() {
	//实例化cli
	app := cli.NewApp()
	app.EnableBashCompletion = true
	//Name可以设定应用的名字
	app.Name = "hello"
	// Version可以设定应用的版本号
	app.Version = "1.0.0"
	// Commands用于创建命令
	app.Flags = []cli.Flag{
		cli.StringFlag{Name: "dir, d", Usage: "Director path to save image.Make sure you have the writing authority."},
		cli.Int64Flag{Name: "id, i", Usage: "Get question ID it from zhihu."},
		cli.IntFlag{Name: "size, s", Usage: "Question count per page.Max 5.", Value: 5},
		cli.IntFlag{Name: "limit, l", Usage: "Max page count.", Value:math.MaxInt32},
	}
	app.Action = func(c *cli.Context) error {
		rootDir := util.Trim(c.String("dir"))
		questionId := c.Int64("id")
		size := c.Int("size")
		limit := c.Int("limit")

		if rootDir == "" || questionId == 0 {
			fmt.Println("Must specify both rootDir and questionId.Use -h to get help.")
		}

		service.GetWonderfulImages(questionId, rootDir, size, limit)
		return nil
	}
	app.Run(os.Args)
}
