# DevPost

DevPost is a git-aware HTTP server suitable for developers. While it should probably not be used for production, it does have several features that should be useful to web developers.

### Why should I use DevPost?

One reason: Simplicity and portability. DevPost was designed with the goal of offering a simple click-and-forget solution for getting right to what matters: your work. DevPost offers a simple, small, capable little server that can be started with one click or command. This means it can be easily integrated with your existing workflow, with no extra overhead. And since it's a self-contained binary, it's portable and doesn't have any dependencies like a Python module (which is hard to get installed on Windows) or Node.js (which is difficult to configure).

In addition, it can be configured right from your browser, without keeping the terminal open. In the future, all of DevPost's options will be configurable from the browser, so there won't be anything to learn, and no need to bring up or keep yet another terminal or command prompt open.

### Features
DevPost is still in early alpha right now, and most of its features are unimplemented. It currently:

1. Serves HTML pages from its working directory
2. Serves bad-looking service pages about DevPost's status, and on 404.
3. Has a bad-looking service directory starting at /devpost. If you navigate to /devpost/stop, the server stops.

Missing from DevPost are the following planned features:

1. Ability to integrate with Git and Git-flow via the /devpost service page.
2. Ability to rename the /devpost service page in case you need, for whatever reason, the path /devpost to be available.
3. Ability to clone Git repositories automatically.
4. Ability to periodically poll a remote git repository for pushes, and pull when necessary, keeping itself up to date.
5. Better looking interface.

### Installing DevPost

Getting DevPost if you have Go setup on your PC is very simple: simply go-get this package, with this command:

```go get github.com/arcticlight/devpost/devpost```

There may also be a binary release of DevPost available for your system. DevPost - when compiled - has no external dependencies, so you can just run the file you download and it should work.
