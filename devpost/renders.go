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
  
  h4 {
    display: inline;
    margin: 0;
  }
  
  span {
    padding: 3px 4px 3px 4px;
    border-radius: 5px;
  }
  
  .status {
    position: relative;
    left: 26px;
    display: inline-block;
    border: 1px solid black;
    border-left: none;
    background-color: #CCC;
    border-radius: 0 5px 5px 0;
    height: 1.1em;
    font-size: 1em;
    padding: 3px 4px 3px 4px;
    margin-left: 0;
  }
  
  .error {
    width: 50%;
    min-width: 300px;
    border: 1px solid darkred;
    background-color: #FFAA99;
    border-radius: 5px;
    margin: 1em auto 0 auto;
  }
  
  .error .code {
    background-color: #EEE;
    
  }
  
  .status.ok {
    background-color: #AAFF99;
    color: darkgreen;
    border-color: darkgreen;
  }
  
  .status.bad {
    background-color: #FFAA99;
    color: darkred;
    border-color: darkred;
  }
  
  .status::before {
    content: "Status:";
    width: 50px;
    position: absolute;
    left: -58px;
    top: -1px;
    border: 1px solid black;
    border-right: none;
    background-color: #DDD;
    color: black;
    height: 1.1em;
    font-size: 1em;
    border-radius: 5px 0 0 5px;
    padding: 3px 4px 3px 4px;
    margin-right: 0;
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
    padding-bottom: 1px;
  }
  
  .code {
    display: inline-block;
    font-family: "Courier New", Courier, monospace;
    background-color: #CCC;
    border-radius: 5px;
    padding-bottom: 1px;
  }
</style>`)

var welcomeTemplate, _ = template.New("WelcomePage").Parse(`
<!DOCTYPE HTML>
<html>
  <head>
    <title>DevPost - Start</title>
    {{.Style}}
  </head>
  <body>
    <h1>DevPost started up!</h1>
    <div class="center">
      <p><em>(You will see this page only once)</em></p>
      {{if .Status.Ok}}<span class="status ok">Everything is good!</span>{{else}}<span class="status bad">Something went wrong</span>
        {{if .Status.Giterror}}<div class="error">
          <h3>Git error:</h3>
          <p>Unable to detect <span class="code">git</span></p>
          <p>You will be unable to use Git functionality until you manually fix this <a href="{{.Prefix}}">in the settings</a>.</p>
        </div>{{end}}
      {{end}}
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
func renderWelcomePage(w http.ResponseWriter, r *http.Request, status dpstatus) {
    welcomeTemplate.ExecuteTemplate(w, "WelcomePage",
    struct {
      Style template.HTML
      Status dpstatus
      Wd, Prefix string
    } { servicestyle, status, workingdir, "/"+controlprefix })
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
