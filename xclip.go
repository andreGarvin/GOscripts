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

				EOFerror := writeToFile(clipboard_path, string( stdin ))
				if EOFerror != nil {
					fmt.Println( EOFerror )
				}
			}
		case "o":

			data_stream, err := readFile( clipboard_path )
			if err != nil {
				fmt.Println( err )
			} else {
				fmt.Println( data_stream )
			}
		case "a":

			stdin, err := ioutil.ReadAll( os.Stdin )
			if err != nil {
				fmt.Println( err )
			} else {

				data_stream, err := readFile( clipboard_path )
				if err != nil {
					fmt.Println( err )
				} else {

					new_data_stream := data_stream + string( stdin )
					EOFerror := writeToFile(clipboard_path, new_data_stream)
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


func readFile( file_name string ) ( string, error ) {

	byte_slice, err := ioutil.ReadFile( file_name )
	if err != nil {
		return "", err
	}
	return string( byte_slice ), nil
}

func writeToFile( file_name, data string ) error {
	
	EOFerror := ioutil.WriteFile( file_name, []byte( data ), 0644 )
	if EOFerror != nil {
		return EOFerror
	}
	return nil
}