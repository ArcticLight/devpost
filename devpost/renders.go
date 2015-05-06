package main

import (
    "net/http"
    "html/template"
)

var servicestyle = template.HTML(`<style>
  html {
    font-family: "Trebuchet MS", Helvetica, sans-serif;
  }

  body {
    background-color: #EAEAEA;
  }
  
  h1 {
    width: 100%;
    text-align: center;
  }
  
  div.center {
    margin: 0;
    width: 100%;
    text-align: center;
  }
  
  .path {
    display: inline-block;
    font-family: "Courier New", Courier, monospace;
    background-color: #CCC;
    padding: 3px 4px 1px 4px;
    border-radius: 5px;
  }
  
  .code {
    display: inline-block;
    font-family: "Courier New", Courier, monospace;
    background-color: #CCC;
    padding: 3px 4px 1px 4px;
    border-radius: 5px;
  }
</style>`)

var welcomeTemplate, _ = template.New("WelcomePage").Parse(`<!DOCTYPE HTML>
<html>
  <head>
    <title>DevPost - Start</title>
    {{.Style}}
  </head>
  <body>
    <h1>DevPost is running!</h1>
    <div class="center">
      <p><em>(You will see this page only once)</em></p>
      <p>Refreshing this page will serve from <span class="path">{{.Wd}}</span></p>
      <p>To change DevPost settings, navigate to <span class="path">{{.Prefix}}</span>, or click <a href="{{.Prefix}}">this link</a></p>
    </div>
  </body>
</html>`)

var stopTemplate, _ = template.New("StopTemplate").Parse(`<!DOCTYPE HTML>
<html>
  <head>
    <title>DevPost - Stopping</title>
    {{.}}
  </head>
  <body>
    <h1>DevPost has stopped.</h1>
    <div class="center">
      <p><em>You will no longer get pages from me.</em></p>
    </div>
  </body>
</html>`)

var fofTemplate, _ = template.New("fofTemplate").Parse(`<!DOCTYPE HTML>
<html>
  <head>
    <title>DevPost - 404</title>
    {{.Style}}
  </head>
  <body>
    <h1>404 - Not Found</h1>
    <p>Unable to serve the file <span class="path">{{.Path}}</span></p>
    <p>The exact error was this: <span class="code">{{.Err}}</span></p>
  </body>
</html>`)

//renderWelcomePage renders the HTML document first seen when DevPost starts up.
func renderWelcomePage(w http.ResponseWriter, r *http.Request) {
    welcomeTemplate.Execute(w,
    struct {
      Style template.HTML
      Wd, Prefix string
    } { servicestyle, workingdir, "/"+controlprefix })
}

func renderStopPage(w http.ResponseWriter, r *http.Request) {
    stopTemplate.Execute(w, servicestyle)
}

func renderControlPage(w http.ResponseWriter, r *http.Request) {
    //stub
}

func renderFileNotFoundPage(w http.ResponseWriter, r *http.Request, path string, err error) {
    fofTemplate.Execute(w,
    struct {
      Style template.HTML
      Path, Err string
    } {servicestyle, path, err.Error()})
}

//guessContent attempts to guess the Content-Type header most
//relevant for the given path. It does this by checking the file
//extension, if it has one, and then setting the appropriate header
//on the http.ResponseWriter provided.
func guessContent(w http.ResponseWriter, path string) {
//TODO: Rewrite this so it's better. Preferably by using Real Extensions
//and not just the last 3 characters of the string.
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

//renderContent renders generic content to a browser. It will try to match a content-type
//header from the known set and set the header for the browser, before writing the byte-buffer
//to the HTTP connection.
func renderContent(w http.ResponseWriter, r *http.Request, path string, content []byte) {
  guessContent(w, path)
  w.Write(content)
}
