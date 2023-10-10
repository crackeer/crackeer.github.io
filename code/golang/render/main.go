package main

import (
	"fmt"
	"os"

	"github.com/unrolled/render"
)

func main() {
	// ...
	r := render.New(render.Options{
		Directory:  "/Users/liuhu016/gocode/render/templates", // Specify what path to load the templates from.
		FileSystem: nil,                                       // Specify filesystem from where files are loaded.
		Layout:     "layout/default",                          // Specify a layout template. Layouts can call {{ yield }} to render the current template or {{ partial "css" }} to render a partial from the current template.
		Delims:     render.Delims{"{[{", "}]}"},               // Sets delimiters to the specified strings.
		Extensions: []string{".html"},
	})

	err := r.HTML(os.Stdout, 200, "home", map[string]interface{}{
		"page": 756,
	})
	if err != nil {
		fmt.Println(err.Error())
	}
}
