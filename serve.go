package main

import (
    "path/filepath"
    "io/ioutil"
    "net/http"
    "strings"
    "flag"
    "fmt"
    "os"
)

// global varibles
var (
    help string = "./serve-help"
    dirPaths []string
    port string
    file string
    dir string
)

func main() {

    // flags given through cli
    flag.StringVar(&port, "port", ":8080", "Give a port number for the server to run, default is 8080.")
    flag.StringVar(&file, "f", "./serve-help", "Give a existing file name to serve, default file is 'serve-help'.")
    flag.StringVar(&dir, "d", "", "Give a existing working directory to serve the index.html file or first file in directory, no default.")

    // parses the flag command line agruemnts
    flag.Parse()

    // check si f flag dir was called to excute 'recursiveTreeDive()'
    if dir != "" {
        fs, err := ioutil.ReadDir( dir )
        if err != nil {
            fmt.Println( err )
        } else {
            recursiveTreeDive( dir, fs )
        }
    }

    // routes for webserver
    http.HandleFunc("/", serveFile)

    // display the webserver is running
    if dir != "" {
        fmt.Printf("Serveing '"+ filepath.Base( dir ) +"' running at http://localhost%s\n", port)
    } else {
        fmt.Printf("Serveing '"+ filepath.Base( file ) +"' running at http://localhost%s\n", port)
    }
    // waits till all other parts are doen and runs webserver
    defer http.ListenAndServe(port, nil)
}

// returns the conents of a given existing file name
func readFile( file_name string ) string {

    byte_stream, err := ioutil.ReadFile( file_name )
    if err != nil {
        fmt.Println( err )
        return "server-error: '"+ file_name +"' does not exist."
    }

    return string( byte_stream )
}


// serves handles the route for the webserver
func serveFile(w http.ResponseWriter, r *http.Request) {
    var url string = r.URL.String()

    // checks if the the url path is not "/favicon.ico"
    if url != "/favicon.ico" {

        if len( dirPaths ) != 0 {

            for _, path := range dirPaths {
                slice_path := strings.Split( string( path ), ".." )
                pathItem := filepath.ToSlash( slice_path[ len( slice_path ) - 1 ] )

                // formats the path if path == 'lib/help' to '/lib/help'
                if string( pathItem[0] ) != "/" {
                    pathItem = "/" + pathItem
                }

                if pathItem == url {

                    fmt.Println("Serving ", pathItem)
                    fmt.Fprintf(w, readFile( string( path ) ))
                }
            }
        } else {

            fmt.Println("Serving ", file)
            fmt.Fprintf(w, readFile(file))
        }
    }
}

// recursively gets all files paths in a given directory
func recursiveTreeDive( root string, dir []os.FileInfo ) {

    for _, f := range dir {
        // gets the file name, concatenate them with the given root,
        // then cleans the path slash.
        f := filepath.Clean( filepath.Join( root, f.Name() ) )

        // appends the file name to the 'dirPaths' array
        dirPaths = append(dirPaths, f )

        stat, err := os.Stat( f )
        if err != nil {
            fmt.Println( err )
        } else {

            // checks weather the file is a directory and not a git folder
            if stat.IsDir() && filepath.Base( f ) != ".git" {
                newRoot := f

                fs, err := ioutil.ReadDir( newRoot )
                if err != nil {
                    fmt.Println( err )
                } else {
                    // calls the function agian passing in the 'newRoot'
                    // and ist directory contents
                    recursiveTreeDive( newRoot, fs )
                }
            }
        }
    }
}
