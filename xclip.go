package main

import (
	// my lib
	"./lib"

	"path/filepath"
	"io/ioutil"
	"fmt"
	"os"
)

func main() {

		args := lib.ArguementParser( os.Args )
		cwd, err := os.Getwd()
		if err != nil {
			fmt.Println( err )
		}

		// path to the clipboard
		clipboard_path := filepath.Clean( filepath.Join(cwd, ".clipboard") )
		if len ( args.Flags ) != 0 {

				switch flag := args.Flags[0]; flag {
				case "i":

						stdin, err := ioutil.ReadAll( os.Stdin )
						if err != nil {
							fmt.Println( err )
						} else {

								EOFerror := ioutil.WriteFile( clipboard_path, stdin, 0644 )
								if EOFerror != nil {
									fmt.Println( err )
								}
						}
				case "o":

						byte_stream, err := ioutil.ReadFile( clipboard_path )
						if err != nil {
							fmt.Println( err )
						}
						fmt.Println( string( byte_stream ) )
				case "a":

						stdin, err := ioutil.ReadAll( os.Stdin )
						if err != nil {
								fmt.Println( err )
						} else {

								byte_stream, err := ioutil.ReadFile( clipboard_path )
								if err != nil {
										fmt.Println( err )
								} else {

										data_stream := string( byte_stream ) + string( stdin )

										EOFerror := ioutil.WriteFile( clipboard_path, []byte( data_stream ), 0644 )
										if EOFerror != nil {
											fmt.Println( err )
										}
								}
						}
				}

		} else {
				fmt.Println("error in xclips.go: Must provide either a input or oupt flag to use program, exit.")
		}
}
