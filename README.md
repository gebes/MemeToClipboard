
<h1 align="center">
Meme to clipboard
</h1> 

CLI tool to copy a random meme

 

# How to use
1. Build the tool with `go build`
2. Add the folder location to your path. 

or

1. Run `go install meme-to-clipboard`

![Screenshot](https://i.imgur.com/JgEXrWu.png)

## Change source
If you want different memes then change the Endpoint in the main.go file
```go
const (
	Endpoint = "https://meme-api.herokuapp.com/gimme/dankmemes"
)
```