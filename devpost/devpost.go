package main

import (
    "fmt"
    "github.com/skratchdot/open-golang/open"
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

//guessContent attempts to guess the Content-Type header most
//relevant for the given path. It does this by checking the file
//extension, if it has one, and then setting the appropriate header
//on the http.ResponseWriter provided.
func guessContent(w http.ResponseWriter, path string) {
//TODO: Rewrite this so it's better.
    if(len(path) <= 4) {
        return
    }
    
    last3 := path[len(path)-3:]
    
    switch {
        case last3 == "css":
            w.Header().Set("Content-Type", "text/css")
        case last3 == ".js":
            w.Header().Set("Content-Type", "text/javascript")
        case last3 == "tml", last3 == "htm":
            w.Header().Set("Content-Type", "text/html")
        case last3 == "svg":
            w.Header().Set("Content-Type", "image/svg+xml")
        case last3 == "png":
            w.Header().Set("Content-Type", "image/png")
        case last3 == "ico":
            w.Header().Set("Content-Type", "image/x-icon")
        case last3 == "son":
            w.Header().Set("Content-Type", "application/json")
        case last3 == "gif":
            w.Header().Set("Content-Type", "image/gif")
        case last3 == "jpg", last3 == "peg":
            w.Header().Set("Content-Type", "image/jpeg")
    }
}

//devpostHandler is the HTTP handler assigned to handle requests for devpost.
//It handles presenting the welcome page on first run, and further delegates
//rendering content or service pages depending on the status of the request.
//
//I could probably make better use of the go HTTP package, however I haven't
//fully investigated the API yet, and I don't know if I can enable the correct
//routing of pages with the dynamic "/devpost" handle if I don't take care of it
//myself.
func devpostHandler(w http.ResponseWriter, r *http.Request) {
    if(firstContact) {
        firstContact = false
        renderWelcomePage(w, r)
    } else {
        if(len(r.URL.Path) >= len(controlprefix)+1 && r.URL.Path[:len(controlprefix)+1] == "/"+controlprefix) {
            cmd := r.URL.Path[len(controlprefix)+1:];
            switch {
                case cmd == "/stop":
                    renderStopPage(w, r)
                    stopServer()
                default:
                    if(r.Method == "GET") {
                        renderControlPage(w, r)
                    }
            }
        } else {
            path := "./" + r.URL.Path[1:]
            if os.PathSeparator != '/' {
                path = strings.Replace(path, "/", string(os.PathSeparator), -1)
            }
            contents, err := ioutil.ReadFile(r.URL.Path[1:])
            if err != nil {
                if len(path) >= 2 && path[len(path)-1:] == string(os.PathSeparator) {
                    path = path[:len(path)-1]
                } else if len(path) < 2 {
                    path = ""
                }
                contents, err = ioutil.ReadFile(path + string(os.PathSeparator) + "index.html")
                if err != nil {
                    w.WriteHeader(404)
                    renderFileNotFoundPage(w, r, path, err)
                } else {
                    renderHTMLPage(contents)
                }
            } else {
                renderContent(w, r, path)
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
