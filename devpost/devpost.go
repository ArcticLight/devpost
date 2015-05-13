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

func init() {
    var err error
    status.workingdir, err = os.Getwd()
    if err != nil {
        panic(err)
    }
}

func execChecks() *dpstatus {
    status.ok = false
    status.giterror = nil
    status.gitpath = ""
    
    var err error
    
    if status.usersets.gitusercommand == "" {
        status.gitpath, err = exec.LookPath("git")
    } else {
        status.gitpath, err = exec.LookPath(status.usersets.gitusercommand)
    }
    
    if err != nil {
        log.Println(err)
        status.giterror = err
    } else {
        status.ok = true
        status.giterror = nil
    }
    
    return status
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
    if(status.firstContact) {
        status.firstContact = false
        renderWelcomePage(w, r, execChecks())
    } else {
        if(len(r.URL.Path) >= len(status.usersets.controlprefix)+1 && r.URL.Path[:len(status.usersets.controlprefix)+1] == "/"+status.usersets.controlprefix) {
            cmd := r.URL.RawQuery
            switch {
                case cmd == "stop":
                    if(r.Method == "GET") {
                        renderStopPage(w, r)
                        status.closereq<-true
                    }
                default:
                    if(r.Method == "GET") {
                        renderControlPage(w, r)
                    }
            }
        } else {
            path := status.workingdir + "/" + r.URL.Path[1:]
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
    fmt.Printf("Starting up in %s...\n", status.workingdir)
    http.HandleFunc("/", devpostHandler)
    go http.ListenAndServe(":8080", nil)
    fmt.Println("Launching your browser at DevPost!")
    open.Run("http://localhost:8080/")
    <- status.closereq
    stopServer()
}
