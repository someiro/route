# Route

A small static site generator written in Go. Route generates HTML from Markdown.

## Installation

With Go installed, run this from the repository root:

```sh
$ go install .
$ export PATH="$(go env GOPATH)/bin:$PATH"
```

Put the Go install directory first in your `PATH` so your shell finds Route's `route` command. On macOS, this avoids a name conflict with the system `/sbin/route` command. Add the `export` line to your shell configuration (for example, `~/.zshrc`) to keep it across sessions.

## Usage

```sh
route new "Hello, Route!"
# Edit md/Hello, Route!.md
route gen "Hello, Route!"
```

Generated HTML is written to `gen/`. Customize the template in `template/article.html` and the styles in `styles/main.css`.

## License

MIT License
