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

		var file string;

		line := []int {};
		args := lib.ArguementParser( os.Args )

		if len( args.Payload ) != 0 {

				for _, e := range args.Payload {
						e = string( e )

						// coverting the given line number/s to int
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

														var inBound bool = false;
														var c int;

														data, err := readFile( file )
														if err != nil {
																fmt.Println( err )
														} else {

																slice_data_stream := strings.Split(, "\n")
																slice_len := len( slice_data_stream )

																/*
																	Head: The top of the file.
																	Tail: The bottom of the file.
																*/
																head, tail := 0, line[0]

																/*
																	case t:
																		"assgining the 'head' the given line number
																		and the 'tail' the len of the sliced file data
																		you will get the only that porion of the file
																		as if you are getting it from bottom to top"

																	case b:
																			"assgin the 'head' the first given line number
																			and the 'tail' the second given lien number"
																*/
																switch flag := args.Flags[0]; flag {
																case "t":
																		head = line[0]
																		tail = slice_len
																case "b":
																		head = line[0]
																		tail = line[1]
																}

																/*
																	check if the length of the second given line
																	number, if there is any, is not greater then
																	the length of 'tail' == 'slice_len'
																*/
																if tail > slice_len {
																		tail = slice_len
																}

																for c = head; c <= tail - 1; c++ {
																		// checking if the the for loop printed out lines from the file data
																		inBound = true;
																		fmt.Println( string( slice_data_stream[ c ] ) )
																}

																if !inBound {
																		fmt.Println("error in seek.go: You have either reached out of bound of file length or this program retruned a empty string.")
																}
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

// return back the file data of a given file
func readFile( file_name string ) ( string, error ) {
		byte_stream, err := ioutil.ReadFile( file_name )
		if err != nil {
				return "", err
		} else {
				return string( byte_stream ), nil
		}
}

// fully converts the number of int64 or in32 to int
func convertToInt( input string ) int {
		/*
			looping through all numbers to find a mautch
			of the given string number.

			*unlikely it will not find the number and return anything
				becasue it is a infite loop.
		*/

		for i := 0; i < ( i + 1 ); i++ {
				str := string( i )

				if str == input {
						return i
				}
		}

		return 0;
}
