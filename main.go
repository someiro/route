package main

import (
	"bytes"
	"fmt"
	"github.com/yuin/goldmark"
	"html/template"
	"os"
	"time"
)

const lang = "ja"
const siteName = "Route"

func genHTML(title string, mdBuf bytes.Buffer) bytes.Buffer {
	now := time.Now()
	data := struct {
		Lang      string
		SiteName  string
		Title     string
		Date      string
		Body      template.HTML
		Copyright string
	}{
		Lang:      lang,
		SiteName:  siteName,
		Title:     title,
		Date:      now.Format("2006 · 01 · 02"),
		Body:      template.HTML(mdBuf.String()),
		Copyright: now.Format("2006"),
	}

	tmpl, err := template.ParseFiles("template/article.html")
	if err != nil {
		fmt.Fprintln(os.Stderr, "> Template Error 💥\nLOG:", err)
		os.Exit(1)
	}

	var outBuf bytes.Buffer

	if err := tmpl.Execute(&outBuf, data); err != nil {
		fmt.Fprintln(os.Stderr, "> Template Error 💥\nLOG:", err)
		os.Exit(1)
	}

	return outBuf
}

func main() {
	args := os.Args
	argsCount := len(args)

	if argsCount == 1 {
		fmt.Println(`

    ____              __     
   / __ \____  __  __/ /____ 
  / /_/ / __ \/ / / / __/ _ \
 / _, _/ /_/ / /_/ / /_/  __/
/_/ |_|\____/\__,_/\__/\___/ 



|Create Markdown| $ route new <FileName>

|Generate HTML| $ route gen <FileName>



Copyright (c) 2026 Route / someiro All right reserved.


`)
		os.Exit(0)
	}

	switch args[1] {
	case "new":
		if argsCount == 3 {
			fileName := args[2]
			f, err := os.OpenFile("md/"+fileName+".md", os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0644)
			if err != nil {
				fmt.Println("Create Error 💥", err, "\n")
				os.Exit(1)
			}
			f.Close()

			fmt.Printf("> Create Markdown ✨\nROUTE: md/%s.md\n\n", fileName)
			os.Exit(0)
		} else {
			fmt.Println("> Create Error 💥\nLOG: No FileName Specified\n")
			os.Exit(1)
		}
	case "gen":
		if argsCount == 3 {
			fileName := args[2]
			source, err := os.ReadFile("md/" + fileName + ".md")
			if err != nil {
				fmt.Println("> Generate Error 💥\nLOG:", err, "\n")
				os.Exit(1)
			}

			var buf bytes.Buffer
			if err := goldmark.Convert(source, &buf); err != nil {
				fmt.Println("> Convert Error 💥\nLOG:", err)
				os.Exit(1)
			}

			outBuf := genHTML(fileName, buf)

			if err := os.WriteFile("gen/"+fileName+".html", outBuf.Bytes(), 0644); err != nil {
				fmt.Println("> Write Error 💥\nLOG:", err)
				os.Exit(1)
			}

			fmt.Printf("> Generate HTML ✨\nROUTE: gen/%s.html\n\n", fileName)
			os.Exit(0)
		} else {
			fmt.Println("> Generate Error 💥\nLOG: No FileName Specified\n")
			os.Exit(1)
		}

	default:
		fmt.Println("> Command Error 💥\nLOG: Not Found Command\n")
		os.Exit(1)
	}
}
