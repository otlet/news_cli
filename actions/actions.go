package actions

import (
	"bufio"
	"fmt"
	"github.com/mmcdole/gofeed"
	"github.com/olekukonko/tablewriter"
	"github.com/otlet/news_cli/options"
	"github.com/urfave/cli"
	"os"
	"regexp"
	"strconv"
)

var (
	domainRegExp = `^(?:https?:\/\/)?(?:[^@\/\n]+@)?(?:www\.)?([^:\/\n]+)`
	settings     = options.GetOptions()
)

type Actions struct{}

func (actions Actions) AddUrl(c *cli.Context) error {
	rss := c.Args().First()
	RegExp := regexp.MustCompile(domainRegExp)
	match := RegExp.MatchString(rss)
	if match {
		f, err := os.OpenFile(settings.Filename, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		options.Check(err)

		_, err = f.WriteString(rss + "\n")
		options.Check(err)

		err = f.Close()
		options.Check(err)

		fmt.Println("Added: ", rss)
	} else {
		fmt.Println("Sorry, it's not a URL")
	}
	return nil
}

func (actions Actions) DelUrl(c *cli.Context) error {
	arg := c.Args().First()

	rss, err := strconv.ParseInt(arg, 10, 0)
	options.Check(err)

	err = settings.RemoveLine(int(rss), 1)
	options.Check(err)
	return nil
}

func (actions Actions) ListUrls(context *cli.Context) error {
	file, err := os.Open(settings.Filename)
	options.Check(err)
	defer file.Close()

	index := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		index++
		fmt.Println(index, scanner.Text())
	}

	if index == 0 {
		fmt.Println("List is empty!")
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func (actions Actions) GetNews(context *cli.Context) error {
	file, err := os.Open(settings.Filename)
	options.Check(err)
	defer file.Close()

	var urlList []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		urlList = append(urlList, scanner.Text())
	}

	fp := gofeed.NewParser()

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Title", "URL", "Published"})
	table.SetRowLine(true)

	for _, url := range urlList {
		feed, _ := fp.ParseURL(url)
		for _, item := range feed.Items {
			table.Append([]string{
				item.Title,
				item.Link,
				item.Published,
			})
		}
	}
	table.Render()

	return nil
}
