package main

import (
	"io/ioutil"
	"os"

	"github.com/fatih/color"
)

type err interface {
		Error() string
}

func main() {
		files_slices := os.Args[1:]

		if len( files_slices ) != 0 {

				for _, i := range files_slices {
						i = string( i )

						stat, err := os.Stat( i )
						if err != nil {
								color.Red("*error in cat.go: %s file does not exist;\nprogram error: %s.\n", i, err)
						} else {

							if stat.IsDir() {
									color.Red("*error in cat.go: %s is a directory.\n", i)
							} else {

									data_stream, err := readFile( i )
									if err != nil {
											color.Red("*%s", err.Error())
									} else {
											color.White( data_stream )
									}
							}
						}
				}
		} else {
				color.Red("*error in cat.go: Must provide input.")
		}
}

// returns a file data and error (EOFerror) if one occurs
func readFile( file_name string ) ( string, error ) {

		byte_slice, err := ioutil.ReadFile( file_name )
		if err != nil {
			return "", err
		}
		return string( byte_slice ), nil
}
