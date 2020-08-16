package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {

	//<<<<===================>>>>
	//sub conmmand
	//cli.AppHelpTemplate = fmt.Sprintf("good a ")
	app := cli.NewApp()
	app.EnableBashCompletion = true
	app.Commands = []*cli.Command{
		{
			Name:    "add",
			Aliases: []string{"a"},
			Usage:   "add a task to the list",
			Action: func(c *cli.Context) error {
				fmt.Println("added task: ", c.Args().First())
				return nil
			},
		},
		{
			Name:    "complete",
			Aliases: []string{"c"},
			Usage:   "complete a task on the list",
			Action: func(c *cli.Context) error {
				fmt.Println("completed task: ", c.Args().First())
				return nil
			},
		},
		{
			Name:    "template",
			Aliases: []string{"t"},
			Usage:   "options for task templates",
			Subcommands: []*cli.Command{
				{
					Name:  "add",
					Usage: "add a new template",
					Action: func(c *cli.Context) error {
						fmt.Println("new task template: ", c.Args().First())
						fmt.Println("zz task template: ", c.Args().First())
						fmt.Println("zz task template: ", c.Args().First())
						fmt.Println("zz task template: ", c.Args().First())
						fmt.Println("zz task template: ", c.Args().First())
						fmt.Println("zz task template: ", c.Args().First())
						fmt.Println("zz task template: ", c.Args().First())
						fmt.Println("zz task template: ", c.Args().First())
						fmt.Println("zz task template: ", c.Args().First())
						fmt.Println("zz task template: ", c.Args().First())
						fmt.Println("zz task template: ", c.Args().First())
						fmt.Println("zz task template: ", c.Args().First())
						return nil
					},
				},
				{
					Name:  "remove",
					Usage: "remove an existing template",
					Action: func(c *cli.Context) error {
						fmt.Println("removed task template: ", c.Args().First())
						return nil
					},
				},
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
	//<<<<===================>>>>

	//<<<<===================>>>>
	//Setting and querying flags is simple.
	//app := &cli.App{
	//	Name:  "测试一些擦参数",
	//	Usage: "测试一些参数而已兄弟们",
	//	Flags: []cli.Flag{
	//		&cli.StringFlag{
	//			Name:     "city",
	//			Value:    "上海",
	//			Aliases:  []string{"c"},
	//			Usage:    "城市中文名",
	//			Required: true,
	//		},
	//		&cli.StringFlag{
	//			Name:     "day",
	//			Aliases:  []string{"d"},
	//			Value:    "今天",
	//			Usage:    "可选: 今天, 昨天, 预测",
	//			Required: true,
	//		},
	//	},
	//	Action: func(c *cli.Context) error {
	//		city := c.String("city")
	//		day := c.String("day")
	//		fmt.Printf("city:%s,day:%s\n", city, day)
	//		return nil
	//	},
	//}
	//
	//err := app.Run(os.Args)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//<<<<===================>>>>

	//<<<<===================>>>>
	//You can lookup arguments by calling the Args function on cli.Context, e.g.:
	//app := &cli.App{
	//	Action: func(c *cli.Context) error {
	//		fmt.Printf("Hello %q", c.Args().Get(0))
	//		return nil
	//	},
	//}
	//
	//err := app.Run(os.Args)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//<<<<===================>>>>

	//<<<<===================>>>>
	//app := &cli.App{
	//	Name: "greet",
	//	Usage: "fight the loneliness!",
	//	Action: func(c *cli.Context) error {
	//		fmt.Println("Hello friend!")
	//		return nil
	//	},
	//}
	//
	//err := app.Run(os.Args)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//<<<<===================>>>>

	//<<<<===================>>>>
	//app := cli.NewApp()
	//app.Name = "weather-cli"
	//app.Usage = "天气预报小程序"
	//
	//app.Flags = &cli.Flag{
	//	cli.StringFlag{
	//		Name:  "city, c",
	//		Value: "上海",
	//		Usage: "城市中文名",
	//	},
	//	cli.StringFlag{
	//		Name:  "day, d",
	//		Value: "今天",
	//		Usage: "可选: 今天, 昨天, 预测",
	//	},
	//}
	//
	//app.Action = func(c *cli.Context) error {
	//	city := c.String("city")
	//	day := c.String("day")
	//
	//	var body, err = Request(apiUrl + city)
	//	if err != nil {
	//		fmt.Printf("err was %v", err)
	//		return nil
	//	}
	//
	//	var r Response
	//	err = json.Unmarshal([]byte(body), &r)
	//	if err != nil {
	//		fmt.Printf("\nError message: %v", err)
	//		return nil
	//	}
	//	if r.Status != 200 {
	//		fmt.Printf("获取天气API出现错误, %s", r.Message)
	//		return nil
	//	}
	//	Print(day, r)
	//	return nil
	//}
	//app.Run(os.Args)
}
