package main

import (
	"io/ioutil"
	"strings"
	"flag"
	"fmt"
	"os"
)

func main() {
	
	file := flag.String("f", "", "Given file name to read.")
	flag.Parse()

	if *file != "" {

		stat, err := os.Stat( *file )
		if err != nil {
			fmt.Println("error in nl.go: ", err)
		} else {

			if !stat.IsDir() {
				
				data_stream, err := readFile( *file )
				if err != nil {
					fmt.Println("error in nl.go: ", err)
				} else {
					slice_data_stream := strings.Split(data_stream, "\n")
					
					for ln := 0; ln < len( slice_data_stream ); ln++ {
						l := slice_data_stream[ln]

						if l != "" {
							fmt.Printf("     %d  %s\n", (ln + 1), l )
						} else {
							fmt.Println( l )
						}
					}
				}
			} else {

				fmt.Println("error in nl.go: Is not a file but a directory, this program only excepts file to be given.")
			}
		}
	} else {
		fmt.Println("error in nl.go: You must pass in a file in execute the program.")
	}
}

func readFile( file_name string ) ( string, error ) {

	byte_stream, err := ioutil.ReadFile( file_name )
	if err != nil {
		return "", err
	}
	return string( byte_stream ), nil
}