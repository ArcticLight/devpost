# Deprecation Notice

DevPost development has been cancelled, since I've been made aware of plenty of better tools that serve its purpose. Anyone who is interested in the concept of a fast, easy to use HTTP development server is encouraged to go check out [Browser-sync](http://www.browsersync.io/) an excellent tool for many of the uses that DevPost attempted to meet.

## Why was DevPost cancelled?
DevPost was an interesting side-project for me to work on to try and produce "the ideal HTTP development server"; one that came up instantly, required no configuration to work, and that could easily integrate with Git-Flow while working on a website.

I started DevPost as an answer to this problem, but only because I was not aware that there *are* a few incredibly good tools out there that do this, such as the above-mentioned browser-sync. Since I'd rather not have a case of [competing standards](https://xkcd.com/927/) on our hands, I have chosen to discontinue development of DevPost in favor of recommending browser-sync.

### DevPost (original documentation)

DevPost was a git-aware HTTP server suitable for developers.

### Why should I use DevPost?

One reason: Simplicity and portability. DevPost was designed with the goal of offering a simple click-and-forget solution for getting right to what matters: your work. DevPost offers a simple, small, capable little server that can be started with one click or command. This means it can be easily integrated with your existing workflow, with no extra overhead. And since it's a self-contained binary, it's portable and doesn't have any dependencies like a Python module (which is hard to get installed on Windows) or Node.js (which is difficult to configure).

In addition, it can be configured right from your browser, without keeping the terminal open. In the future, all of DevPost's options will be configurable from the browser, so there won't be anything to learn, and no need to bring up or keep yet another terminal or command prompt open.

### Features
DevPost was only available in early alpha form before development was discontinued, and most of its features remain unimplemented. It:

1. Serves HTML pages from its working directory
2. Serves bad-looking service pages about DevPost's status, and on 404.
3. Has a bad-looking service directory starting at /devpost. If you navigate to /devpost?stop, the server stops.

### Installing DevPost

Getting DevPost if you have Go setup on your PC is very simple: simply go-get this package, with this command:

```go get github.com/arcticlight/devpost/devpost```

There may also be a binary release of DevPost available for your system. DevPost - when compiled - has no external dependencies (other than git which is optional), so you can just run the file you download and it should work.
