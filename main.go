package main

import (
	"flag"
	"fmt"
	"os"
	"text/template"
)

func main() {
	f := flag.NewFlagSet("smptpl", flag.ExitOnError)
	f.Usage = func() {
		_, err := fmt.Fprintf(f.Output(), "Usage:\n  %s [options] <template file>\n\nOptions:\n", f.Name())
		if err != nil {
			panic(err)
		}
		f.PrintDefaults()
	}

	vars := make(map[string]string)
	outfile := fileVar{os.Stdout, false}

	defer outfile.Close()

	f.Var((mapVar)(vars), "var", "")

	f.Var(outfile, "outfile", "")

	err := f.Parse(os.Args[1:])
	if err != nil {
		panic(err)
	}

	tmpl, err := template.ParseFiles(f.Arg(0))

	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(outfile.file, GetEnvirons())
	if err != nil {
		panic(err)
	}

}
