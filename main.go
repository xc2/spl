package main

import (
	"flag"
	"fmt"
	sprig "github.com/Masterminds/sprig"
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

	f.Var(&outfile, "outfile", "")

	err := f.Parse(os.Args[1:])

	if err != nil {
		panic(err)
	}

	infile := f.Arg(0)
	if infile == "" {
		infile = "-"
	}

	tmpl := template.New(infile)
	tmpl.Option("missingkey=zero")

	tmpl.Funcs(sprig.GenericFuncMap())

	err = TemplateParseFile(tmpl, infile)

	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(outfile.file, vars)
	if err != nil {
		panic(err)
	}

}
