package options

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
)

type Options struct {
	Filename string
}

func GetOptions() Options {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	optionFileName := usr.HomeDir + "/.news"

	options := Options{
		Filename: optionFileName,
	}

	if _, err := os.Stat(optionFileName); err == nil {
		//TODO: IDK what to do with that...
	} else if os.IsNotExist(err) {
		_, err := os.Create(optionFileName)
		if err != nil {
			panic(err)
		}
		fmt.Println("Option file created:", optionFileName)
	} else {
		panic("Schrodinger File: file may or may not exist. See err for details.")
	}

	return options
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func (o Options) RemoveLine(start, n int) (err error) {
	if start < 1 {
		return errors.New("bad argument. Bad RSS number")
	}
	if n < 0 {
		return errors.New("I like trains")
	}
	var f *os.File
	if f, err = os.OpenFile(o.Filename, os.O_RDWR, 0); err != nil {
		return
	}
	defer func() {
		if cErr := f.Close(); err == nil {
			err = cErr
		}
	}()
	var b []byte
	if b, err = ioutil.ReadAll(f); err != nil {
		return
	}
	cut, ok := skip(b, start-1)
	if !ok {
		return fmt.Errorf("too big number: %d", start)
	}
	if n == 0 {
		return nil
	}
	tail, ok := skip(cut, n)
	if !ok {
		return fmt.Errorf("less than %d lines after line %d", n, start)
	}
	t := int64(len(b) - len(cut))
	if err = f.Truncate(t); err != nil {
		return
	}
	if len(tail) > 0 {
		_, err = f.WriteAt(tail, t)
	}
	return
}

func skip(b []byte, n int) ([]byte, bool) {
	for ; n > 0; n-- {
		if len(b) == 0 {
			return nil, false
		}
		x := bytes.IndexByte(b, '\n')
		if x < 0 {
			x = len(b)
		} else {
			x++
		}
		b = b[x:]
	}
	return b, true
}
