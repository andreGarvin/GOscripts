package main

import (
	"./lib"

	"io/ioutil"
	"strings"
	"strconv"
	"fmt"
	"os"
)

func main() {

		line := []int {};
		var file string;
		args := lib.ArguementParser( os.Args )

		if len( args.Payload ) != 0 {

				for _, e := range args.Payload {
						e = string( e )

						i, err := strconv.ParseInt(e, 10, 23)
						if err == nil {
								line = append(line, convertToInt( string( i ) ))
						} else {

								stat, err := os.Stat( e )
								if err != nil {
										fmt.Println( err )
								} else {

										if !stat.IsDir() {
												file = e;

												if len( args.Flags ) != 0 {

														slice_data_stream := strings.Split(readFile( file ), "\n")
														slice_len := len( slice_data_stream )
														head, tail := 0, line[0]
														var inBound bool = false;
														var c int;

														switch flag := args.Flags[0]; flag {
														case "t":
																head = line[0]
																tail = slice_len
														case "b":
																head = line[0]
																tail = line[1]
														}

														if ( len( line ) == 2 && line[1] > tail ) {
																tail = slice_len
														} else if tail > slice_len {
																tail = slice_len
														}

														// fmt.Println( head, tail )
														for c = head; c <= tail - 1; c++ {
																inBound = true;
																fmt.Println( string( slice_data_stream[ c ] ) )
														}
														if !inBound {
																fmt.Println("error in seek.go: You have eaither reached out of bound of file length or this program retruned a empty string.")
														}

												} else {
														fmt.Println("error in seek.go: This program must get a flag in order to excute a command.")
												}
										} else {
												fmt.Println("error in seek.go: This program only takes files as input.")
										}
								}
						}
				}
		}
}


func readFile( file_name string ) string {
		byte_stream, err := ioutil.ReadFile( file_name )
		if err != nil {
				fmt.Println( err )
		}

		return string( byte_stream )
}


func convertToInt( input string ) int {

		var i int;
		for i = 0; i < ( i + 1 ); i++ {
				str := string( i )

				if str == input {
						return i
				}
		}
		return 0;
}
