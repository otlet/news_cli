package main

import (
	"github.com/otlet/news_cli/actions"
	newsflag "github.com/otlet/news_cli/flags"
	"github.com/urfave/cli"
	"log"
	"os"
	"sort"
)

func main() {
	app := cli.NewApp()
	app.Name = "News CLI"
	app.Usage = "Read news in CLI!"
	app.Copyright = "MIT"
	app.Author = "Paweł \"Otlet\" Otlewski <otlet@otlet.pl>"
	app.Version = "1.0.0"

	flags := newsflag.RegisterFlags{}
	app.Flags = flags.GetRegisterFlags()
	cliCommandActions := actions.Actions{}

	app.Commands = []cli.Command{
		{
			Name:      "add",
			Aliases:   []string{"a"},
			Usage:     "Add RSS URL to list",
			UsageText: "news_cli add <url to rss>",
			Action:    func(c *cli.Context) error { return cliCommandActions.AddUrl(c) },
		},
		{
			Name:      "delete",
			ShortName: "del",
			Aliases:   []string{"d"},
			Usage:     "Remove RSS URL to list",
			UsageText: "news_cli delete <id>",
			Action:    func(c *cli.Context) error { return cliCommandActions.DelUrl(c) },
		},
		{
			Name:    "list",
			Aliases: []string{"l"},
			Usage:   "List URLs",
			Action:  func(c *cli.Context) error { return cliCommandActions.ListUrls(c) },
		},
		{
			Name:    "news",
			Aliases: []string{"n"},
			Usage:   "List News",
			Action:  func(c *cli.Context) error { return cliCommandActions.GetNews(c) },
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
