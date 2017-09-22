package main

/**
  This is rpgram is a simple server for serving files
  and web applications for, developement purposes.
  Passing a list for very useful arguements such as
  which port to listen on, to run in dev mode, or
  to serve a file or directory.
*/

import (
    "path/filepath"
    "io/ioutil"
    "net/http"
    "strings"
    "flag"
    "fmt"
    "os"

    "github.com/fatih/color"
    "github.com/andreGarvin/allPaths"
)

// global varibles
var (
    help string = "./test-files/serve-help"
    dirPaths []string
    include string
    port string
    file string
    mode string
    dir string
)

type err interface {
    Error() string
}

func main() {

    // flags given through cli
    flag.StringVar(&port, "port", ":8080", "Give a port number `:<port number>` for the server to run, default is :8080.")
    flag.StringVar(&file, "f", "./serve-help", "Give a existing file name to serve, default file is 'serve-help'.")
    flag.StringVar(&dir, "d", "", "Give a existing working directory to serve the index.html file or first file in directory, no default.")
    flag.StringVar(&mode, "mode", "", "Triggers the server to etheir in developement mode or as a regular serve, default is regular serve")
    flag.StringVar(&include, "i", "", "This adds file to include as a dir item in the dir selection.")
    // parses the flag command line agruemnts
    flag.Parse()

    // check if flag dir was called to excute allPaths.All() function
    if dir != "" {
        fs, err := ioutil.ReadDir( dir )
        if err != nil {
            color.Red( err.Error() )
        } else {
            paths, err := allPaths.All( dir, fs )
            if err != nil {
                fmt.Println(err)
            } else {
                dirPaths = paths
            }
        }
    }

    if include != "" {
       dirPaths = append(dirPaths, include)

       if file != "" {
          dirPaths = append(dirPaths, file)
       }
    }

    // routes for webserver
    if mode == "dev" {
        http.HandleFunc("/", serveDev)
    } else {
        http.HandleFunc("/", serve)
    }

    // display the webserver is running
    if dir != "" || len( dirpaths ) != 0 {
        if mode == "dev" {
            color.Blue("Running DEV Server at http://localhost%s\n", port)
        } else {
            color.Green("Serveing '"+ filepath.Base( dir ) +"' running at http://localhost%s\n", port)
        }
    } else {
        color.Green("Serveing '"+ filepath.Base( file ) +"' running at http://localhost%s\n", port)
    }

    // waits till all other parts are done and runs webserver
    defer http.ListenAndServe(port, nil)
}


// iterates over a slice array of type string and check if the item exist returns the index of the item
func includes( arr []string, item string ) int {

    for i, arr_item := range  arr {

        if filepath.Base( item ) == filepath.Base( string( arr_item ) ) {
            return i
        }
    }
    return -1
}

// returns the conents of a given existing file name
func readFile( file_name string ) ( string, error ) {

    byte_stream, err := ioutil.ReadFile( file_name )
    if err != nil {
        return "", err
    }
    return string( byte_stream ), nil
}


// serves handles the route for the webserver
func serve(w http.ResponseWriter, r *http.Request) {
    var url string = r.URL.String()

    // checks if the the url path is not "/favicon.ico"
    if url != "/favicon.ico" {

        if len( dirPaths ) != 0 {

            /*
                if the url path is '/' and it is a folder check if there
                is there a inde.html file in that folder given
            */
            if url == "/" {

                index := includes( dirPaths, filepath.Clean( filepath.Join(dir, "index.html" ) ) )
		            if index != -1 {

		                /*
                        get the item out of the 'dirPaths'
                        with he items index from the slice
                    */
                    file = dirPaths[ index ]

                    data_stream, err := readFile(file)
                    if err != nil {
                        color.Red( err.Error() )
                        fmt.Fprintf(w, err.Error())
                    }

                    color.Yellow("Serve: [GET] %s\n", file)
                    fmt.Fprintf(w, data_stream)
                }
            }

            for _, path := range dirPaths {
                slice_path := strings.Split( string( path ), ".." )
                pathItem := filepath.ToSlash( slice_path[ len( slice_path ) - 1 ] )

                // formats the path if path == 'lib/help' to '/lib/help'
                if string( pathItem[0] ) != "/" {
                    pathItem = "/" + pathItem
                }

                if pathItem == url {

                    data_stream, err := readFile( pathItem[1:] )
                    if err != nil {
                        color.Red( err.Error() )
                        fmt.Fprintf(w, err.Error())
                    } else {
                        color.Yellow("Serve: [GET] %s\n", pathItem)
                        fmt.Fprintf(w, data_stream)
                    }
                }
            }
        } else {

            data_stream, err := readFile( file )
            if err != nil {
                color.Red( err.Error() )
                fmt.Fprintf(w, err.Error())
            }
            color.Yellow("Serve: [GET] %s\n", file)
            fmt.Fprintf(w, data_stream)
        }
    }
}

func serveDev(w http.ResponseWriter, r *http.Request) {
      var url string = r.URL.String()

      if url != "/favicon.ico" {
          color.Yellow("Serve: [GET] %s\n", url)

          if url == "/" {
              index := includes( dirPaths, filepath.Clean( filepath.Join(dir, "index.html" ) ) )
              if index != -1 {
                  file = dirPaths[ index ]

                  data_stream, err := readFile(file)
                  if err != nil {
                      color.Red( err.Error() )
                      fmt.Fprintf(w, err.Error())
                  }
                  fmt.Fprintf(w, data_stream)
                  return
              }
          } else {

              index := includes(dirPaths, url)
              if index != -1 {
                  file = dirPaths[ index ]

                  data_stream, err := readFile(file)
                  if err != nil {
                      color.Red( err.Error() )
                      fmt.Fprintf(w, err.Error())
                  }
                  fmt.Fprintf(w, data_stream)
                  return
              } else {

                  index := includes( dirPaths, filepath.Clean( filepath.Join(dir, "index.html" ) ) )
                  if index != -1 {
                      file = dirPaths[ index ]

                      data_stream, err := readFile(file)
                      if err != nil {
                          color.Red( err.Error() )
                          fmt.Fprintf(w, err.Error())
                      }
                      fmt.Fprintf(w, data_stream)
                      return
                  }
              }
          }
      }
}

// // recursively gets all files paths in a given directory
// func recursiveTreeDive( root string, dir []os.FileInfo ) {
//
//     for _, f := range dir {
//         // gets the file name, concatenate them with the given root,
//         // then cleans the path slash.
//         f := filepath.Clean( filepath.Join( root, f.Name() ) )
//
//         // appends the file name to the 'dirPaths' array
//         dirPaths = append(dirPaths, f )
//
//         stat, err := os.Stat( f )
//         if err != nil {
//             color.Red( err.Error() )
//         } else {
//
//             // checks weather the file is a directory and not a git folder
//             if stat.IsDir() && filepath.Base( f ) != ".git" {
//                 newRoot := f
//
//                 fs, err := ioutil.ReadDir( newRoot )
//                 if err != nil {
//                     color.Red( err.Error() )
//                 } else {
//                     // calls the function agian passing in the 'newRoot'
//                     // and ist directory contents
//                     recursiveTreeDive( newRoot, fs )
//                 }
//             }
//         }
//     }
// }
