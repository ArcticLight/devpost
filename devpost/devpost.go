package main

import (
    "fmt"
    "github.com/skratchdot/open-golang/open"
    "io/ioutil"
    "net/http"
    "os"
    "strings"
    "os/exec"
    "log"
)

type dpstatus struct {
    Ok bool
    Giterror bool
    Gitpath string
}

var workingdir string
var closereq chan(bool)
var firstContact = true
var controlprefix = "devpost"
var gitcommand = "git"
var gitusercommand = ""

func init() {
    var err error
    workingdir, err = os.Getwd()
    if err != nil {
        panic(err)
    }
    
    closereq = make(chan(bool), 1)
}

func execChecks() *dpstatus {
    var ret = dpstatus{Ok: false, Giterror: true, Gitpath: ""}
    var err error
    if gitusercommand == "" {
        ret.Gitpath, err = exec.LookPath(gitcommand)
    } else {
        ret.Gitpath, err = exec.LookPath(gitusercommand)
    }
    
    if err != nil {
        log.Println(err)
    } else {
        ret.Ok = true
        ret.Giterror = false
    }
    
    return &ret
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
        renderWelcomePage(w, r, execChecks())
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
            path := workingdir + "/" + r.URL.Path[1:]
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
