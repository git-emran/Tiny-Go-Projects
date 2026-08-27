# Example Markdown File 
This is an example Markdown File to test the preview tool

## Features
* Support for links [google](https:/google.com)
* Basic markdown syntax support

## `cli-markdown-preview`
**What it is:** A Minimalistic command line based Markdown preview tool. Using `bluemonday` package from golang which is an html sanitizer, `blackfriday/v2` which is a markdown processor and provides full UTF-8 support. 

### Features
     1	markdown → html conversion with blackfriday
     2	html sanitization with bluemonday
     3	custom html templates
     4	temporary html file generation
     5	automatic browser preview
     6	cross-platform support: macos, linux, and windows

### Usage
go run . -file README.md

Skip opening the browser:
`
go run . -file README.md -s
`

Use a custom HTML template:
`
go run . -file README.md -t template.html
`

The generated HTML file path is printed to stdout.

Build
`go build -o mdp .`

Then:
`./mdp -file README.md`
