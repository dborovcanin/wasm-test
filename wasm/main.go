package main

import (
	"bytes"
	"embed"
	"strconv"
	"syscall/js"
	"text/template"
)

//go:embed templates/*
var pageTemplate embed.FS

var counter = 1

var tmplts map[string]string = map[string]string{}

var tmpltFiles map[string]string = map[string]string{
	"test": "templates/template.html",
}

// Simulated request handler
func handleRequest(this js.Value, args []js.Value) any {
	counter++
	request := args[0].String() // Get request input
	type V struct {
		Value string
	}
	tmpl, err := template.New("test").Parse(tmplts["test"])
	if err != nil {
		panic(err)
	}
	val := V{Value: request + ": " + strconv.Itoa(counter)}
	var buff bytes.Buffer

	err = tmpl.Execute(&buff, val)
	if err != nil {
		panic(err)
	}
	return js.ValueOf(buff.String())
}

func main() {
	// Add templates
	for k, v := range tmpltFiles {
		bytes, _ := pageTemplate.ReadFile(v)
		tmplts[k] = string(bytes)

	}
	// Expose the function to JavaScript
	js.Global().Set("wasmHandleRequest", js.FuncOf(handleRequest))

	// Keep the program running
	select {}
}
