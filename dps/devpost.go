package main

import (
    "fmt"
    "github.com/skratchdot/open-golang/open"
    rez "github.com/arcticlight/devpost/dps/resources"
    "io/ioutil"
    "net/http"
    "os"
    "strings"
)

var workingdir string
var closereq chan(bool)
var firstContact = true
var controlprefix = "devpost"

func init() {
    var err error
    workingdir, err = os.Getwd()
    if err != nil {
        panic(err)
    }
    
    closereq = make(chan(bool), 1)
}

func devpostHandler(w http.ResponseWriter, r *http.Request) {
    if(firstContact) {
        firstContact = false
        fmt.Fprintf(w, rez.Welcomepage(workingdir))
    } else {
        if(len(r.URL.Path) >= len(controlprefix)+1 && r.URL.Path[:len(controlprefix)+1] == "/"+controlprefix) {
            cmd := r.URL.Path[len(controlprefix)+1:];
            switch {
                case cmd == "/stop":
                    fmt.Fprintf(w, rez.Stoppage())
                    closereq<-true
                default:
                    if(r.Method == "GET") {
                        fmt.Fprintf(w, rez.Controlpage(workingdir, &controlprefix))
                    }
            }
        } else {
            path := r.URL.Path[1:]
            if(os.PathSeparator != '/') {
                path = strings.Replace(path, "/", string(os.PathSeparator), -1)
            }
            contents, err := ioutil.ReadFile(r.URL.Path[1:])
            if(err != nil) {
                w.WriteHeader(404)
                fmt.Fprintf(w, rez.FileNotFound(path, err))
            } else {
                w.Write(contents)
            }
        }
    }
}

func main() {
    fmt.Printf("Starting up in %s...\n", workingdir)
    http.HandleFunc("/", devpostHandler)
    go http.ListenAndServe(":8080", nil)
    fmt.Println("Launching your browser at DevPost!")
    open.Run("http://localhost:8080/")
    <- closereq
}
