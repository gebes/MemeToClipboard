# MemeToClipboard
 Golang CLI tool to fetch a direct meme URL from the web

## Usage
Build the tool with `go build` and add the folder location to your path. Afterwards you can use the application in your CLI.

![Screenshot](https://i.imgur.com/pZlV0hL.png)

## Change meme API
If you want different memes then change the Endpoint in the main.go file
```go
const (
	Endpoint = "https://meme-api.herokuapp.com/gimme/dankmemes"
)
```