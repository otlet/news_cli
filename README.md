# News CLI

Read news in CLI like a boss!
That your boss will see the work in the console, not in the browser :)

## Dependency

* `github.com/mmcdole/gofeed`
* `github.com/olekukonko/tablewriter`
* `github.com/urfave/cli`

## Installation

`go install news_cli.go`

## Build

```bash
go build news_cli.go # Build for your system!
env GOOS=linux GOARCH=amd64 go build -v news_cli.go # Build for Linux
env GOOS=darwin GOARCH=amd64 go build -v news_cli.go # Build for Mac
env GOOS=windows GOARCH=amd64 go build -v news_cli.go # Build for Windows
```

## Usage

```bash
news_cli command [arguments...]
news_cli add <link> # Add new RSS channel
news_cli del <id> # Remove url from list
news_cli list # RSS List
news_cli news # Show news list
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## Versioning

We use [SemVer](http://semver.org/) for versioning.
For the versions available, see the [tags on this repository](https://github.com/otlet/News_Cli/tags). 

## Authors

* **Paweł Otlewski** - *Initial work* - [otlet](https://github.com/otlet)

See also the list of [contributors](https://github.com/otlet/news_cli/contributors) who participated in this project.

## License
[MIT](https://choosealicense.com/licenses/mit/)