package main

import (
	"io/ioutil"
	"fmt"
	"os"
)


func main() {

	files_slices := os.Args[1:]

	if len( files_slices ) != 0 {

		for _, i := range files_slices {
			i = string( i )

			stat, err := os.Stat( i )
			if err != nil {
				fmt.Printf("error in cat.go: %s file does not exist;\nprogram error: %s.\n", i, err)
			} else {

				if stat.IsDir() {
						fmt.Printf("error in cat.go: %s is a directory.\n", i)
				} else {

					data_stream, err := readFile( i )
					if err != nil {
						fmt.Println( err )
					} else {
						fmt.Println( data_stream )
					}
				}
			}
		} 
	} else {
		fmt.Println("error in cat.go: Must provide input.")
	}
}


func readFile( file_name string ) ( string, error ) {

	byte_slice, err := ioutil.ReadFile( file_name )
	if err != nil {
		return "", err
	}
	return string( byte_slice ), nil
}