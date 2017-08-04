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

			if !os.IsNotExist( err ) {
				if err != nil {

					fmt.Println(err)
				} else {

					if stat.IsDir() {

						fmt.Printf("error in cat.go: %s is a directory.\n", i)
					} else {

						file := i
						byte_slice, err := ioutil.ReadFile( file )
						if err != nil {

							fmt.Println(err)
						} else {

							fmt.Printf( string( byte_slice ) )
						}
					}
				}
			} else {

				fmt.Printf("error in cat.go: %s file does not exist.\n", i)
			}
		}
	} else {
		fmt.Println("error in cat.go: Must provide input.")
	}
}
