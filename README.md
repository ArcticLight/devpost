# DevPost

DevPost is a git-aware HTTP server suitable for developers. While it should probably not be used for production, it does have several features that should be useful to web developers.


# Features
DevPost is still in beta right now, and most of its features are unimplemented. It currently:
1. Serves HTML pages from its working directory
2. Serves bad-looking service pages about DevPost's status, and on 404.
3. Has a bad-looking service directory starting at /devpost. If you navigate to /devpost/stop, the server stops.

Missing from DevPost are the following planned features:
1. Ability to integrate with Git and Git-flow via the /devpost service page.
2. Ability to rename the /devpost service page in case you need, for whatever reason, the path /devpost to be available.
3. Ability to clone Git repositories automatically.
4. Ability to periodically poll a remote git repository for pushes, and pull when necessary, keeping itself up to date.
5. Better looking interface.
