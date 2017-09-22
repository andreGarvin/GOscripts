package main

import (
	// my golib
	"./golib"

	"io/ioutil"
	"strings"
	"fmt"
	"os"
)

func main() {

		// parses the cli argurments passed to the program
		args := golib.ArgumentParser( os.Args )

		// declares slice array to hold the list of items in directory
		var dir []string;

		if len( args.Payload ) != 0 {
				for _, d := range args.Payload {

						// reads fs with the given 'directory_name'
						fs, err := ioutil.ReadDir( d )
						if err != nil {
								fmt.Println( err )
						} else {

								// iterates over the slice array returned from ioutil.ReadDir()
								for _, f := range fs {
										// gets the file name
										f := f.Name()

										// appends the file name to the 'dir' slice
										dir = append(dir, f)
								}

								fmt.Println( strings.Join( dir, "  ") )
						}

				}
		} else {

				fs, err := ioutil.ReadDir(".")
				if err != nil {
						fmt.Println(err)
				} else {
						for _, f := range fs {
								f := f.Name()
								dir = append(dir, f)
						}

						fmt.Println( strings.Join(dir, " ") )
				}

		}
}
