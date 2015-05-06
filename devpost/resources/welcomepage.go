package resources

/* This package contains static resources for DevPost.
 * It is probably horribly organized, and might have been made
 * better by using HTML templates. Go even has features
 * for this. But I'm lazy and/or not smart, so this hasn't taken
 * place yet.
 *
 * The reason these are hard-coded snippets and NOT real resource files
 * is that DevPost should NOT rely on a specific installation directory
 * or supporting files to work. Part of DevPost's design philosophy is
 * that if you can't relocate the binary onto a flash drive and run it
 * from there, it shouldn't be in DevPost. It might make the binary
 * a bit fatter than usual, but it will ensure it Always Works.
 */
 
var dopstyle = `
    <style>
      h1 {
        width: 100%%;
        text-align: center;
      }
      .code {
        display: inline-block;
        font-family: "monospace";
        background-color: #CCC;
        border-radius: 5px;
        padding: 1px 3px 1px 3px;
        margin: 0 2px 0 2px;
      }
      .centered {
        width: 100%%;
        text-align: center;
      }
    </style>`

func Welcomepage(wdir string) string {
    return`<!DOCTYPE HTML>
<html>
  <head>
` + dopstyle +`
  </head>
<body>
  <h1>DevPost is starting up!</h1>
    <div class="centered"><p><em>(You'll only see this page once.)</em></p>
    <p>Refresh this page and you'll see whatever is located at <span class="code">`+ wdir +`</span>.</p>
    <p>To control DevPost further, navigate to <span class="code">/devpost</span> <em>(yes, in your browser! :)&nbsp;)</em></div>
</body>
</html>`
}

func Controlpage(wdir string, controlprefix *string) string {
    return `<!DOCTYPE HTML>
<html>
  <head>
`+ dopstyle +`
  </head>
<body>
  <h1>DevPost Control Page</h1>
  <p>Your current control URL: <span class="code">/` + *controlprefix + `</span></p>
  <p>Click <a href="/` + *controlprefix + `/stop">this link</a> to KILL devpost.</p>
</body>
</html>
`
}

func FileNotFound(path string, err error) string {
    return `<!DOCTYPE HTML>
<html>
<head>
`+dopstyle+`
</head>
<body>
    <h1>404 Not Found</h1>
    <p>Couldn't find file: <span class="code">`+path+`</span></p>
    <p>The exact error was:</p>
    <p class="code">`+err.Error()+`</p>
</body>
</html>
`
}

func Stoppage() string {
    return "DevPost has stopped."
}
