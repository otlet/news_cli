package flags

import "github.com/urfave/cli"

type RegisterFlags struct {
}

func (f RegisterFlags) addNewsFlag() cli.StringFlag {
	var add string
	return cli.StringFlag{
		Name:        "add, a",
		Value:       "",
		Usage:       "Add RSS url",
		Destination: &add,
	}
}

func (f RegisterFlags) delNewsFlag() cli.StringFlag {
	var del string
	return cli.StringFlag{
		Name:        "del, delete, d",
		Value:       "",
		Usage:       "Delete RSS url",
		Destination: &del,
	}
}

func (f RegisterFlags) listRssNewsFlag() cli.BoolFlag {
	var list bool
	return cli.BoolFlag{
		Name:        "list, l",
		Usage:       "List RSS urls",
		Destination: &list,
	}
}

func (f RegisterFlags) GetRegisterFlags() []cli.Flag {
	return []cli.Flag{
		//f.addNewsFlag(),
		//f.delNewsFlag(),
		//f.listRssNewsFlag(),
	}
}
