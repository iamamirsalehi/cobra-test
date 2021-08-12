package main

import(
	"github.com/iamamirsalehi/note-cobra/data"
	"github.com/iamamirsalehi/note-cobra/cmd"
)

func main() {
	data.OpenDatabase()
	cmd.Execute()
}
