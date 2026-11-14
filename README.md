<img  width="250" alt="route" src="https://github.com/user-attachments/assets/a415c6b7-ed59-4ee0-a220-3971d51ce280" />
<br><br>

A small static site generator written in Go. Route generates HTML from Markdown.

## Installation

With Go installed, run this from the repository root:

```sh
$ go install .
```

## Usage

Run these commands from the root of the cloned repository. To use a different working directory, first copy the `template/` and `styles/` directories from the repository into it, then run the commands below from that directory.

```sh
$ groute init
$ groute new "Hello, Route!"
# Edit md/Hello, Route!.md
$ groute gen "Hello, Route!"
```

Generated HTML is written to `gen/`. Customize the template in `template/index.html`, `template/article.html` and the styles in `styles/main.css`.

## License

MIT License
