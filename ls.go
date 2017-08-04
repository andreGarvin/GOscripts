package main

import (
	// mt lib
	"./lib"


	"io/ioutil"
	"strings"
	"fmt"
	"os"
)

func main() {

	args := lib.ArguementParser( os.Args )

	// declares slice to hold the list of items in directory
	var dir []string;

	if len( args.Payload ) != 0 {
		for _, d := range args.Payload {

			// reads the given 'directory_name'
			fs, err := ioutil.ReadDir( d )
			if err != nil {
				fmt.Println( err )
			}

			// iterates over the slice returned from ioutil.ReadDir()
			for _, f := range fs {
				// gets the file name
				f := f.Name()

				// appends the file name to the 'dir' slice
				dir = append(dir, f)
			}

			fmt.Println( strings.Join( dir, "  ") )
			dir = []string {};
		}
	} else {
		fs, err := ioutil.ReadDir(".")
		if err != nil {
			fmt.Println(err)
		}

		for _, f := range fs {
			f := f.Name()

			dir = append(dir, f)
		}

		fmt.Println( strings.Join(dir, " ") )
		dir = []string {};
	}
}
