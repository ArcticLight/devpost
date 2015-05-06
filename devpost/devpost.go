package main

import (
    "fmt"
    "github.com/skratchdot/open-golang/open"
    "io/ioutil"
    "net/http"
    "os"
    "strings"
   // "log"
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



//renderHTMLPage renders the content []byte as a served HTML document.
//It does NOT validate whether the content is actually an HTML document, but it
//WILL set the Content-Type header before writing the response, and MAY perform
//devpost-specific post-processing before rendering it.
//
//At the moment, this function simply writes the content-type header and
//prints the content as-is. It is separated into its own helper function however
//so that in the future, devpost may include dynamic page rewriting features such as
//Livereload support.
func renderHTMLPage(w http.ResponseWriter, r *http.Request, content []byte) {
    w.Header().Set("Content-Type", "text/html")
    w.Write(content)
}

//stopServer stops the running devpost server.
//
//In the future, it may perform cleanup or other operations involved in shutting
//down DevPost.
func stopServer() {
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
            cmd := r.URL.RawQuery
            switch {
                case cmd == "stop":
                    if(r.Method == "GET") {
                        renderStopPage(w, r)
                        closereq<-true
                    }
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
                    renderHTMLPage(w, r, contents)
                }
            } else {
                renderContent(w, r, path, contents)
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
    stopServer()
}
