package main

import (
	"flag"
	"fmt"
	"github.com/Masterminds/sprig"
	"os"
	"text/template"
)

func main() {
	f := flag.NewFlagSet("spl", flag.ExitOnError)
	f.Usage = func() {
		Usage(f)
	}

	vars := make(map[string]string)
	outfile := fileVar{os.Stdout, false}

	defer outfile.Close()

	f.Var((mapVar)(vars), "v", "")
	f.Var((mapVar)(vars), "var", "")

	f.Var(&outfile, "o", "")
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

func Usage(f *flag.FlagSet) {
	name := f.Name()
	_, err := fmt.Fprintf(f.Output(), `NAME
  %[1]s - Read, parse as golang template and write.

SYNOPSIS
  %[1]s [options] [path/to/template]

DESCRIPTION
  This utility simply reads text, parses it as golang template and writes it to somewhere.
  github.com/masterminds/sprig is used to enhance capbilities of template syntax. 

  [path/to/template]
       Path to template file. Read from stdin if ignored or set to -.

  -v <key>=<value>, -var <key>=<value>
       Set template vars. Use multiple times to set multiple vars.

  -o <path>, -outfile <path>
       Path to result file. Write to stdout if ignored or set to -.

  -h, -help
       Show this message.

EXAMPLES
  Read from stdin and write to stdout 

      $ echo '{{ .hello }}' | %[1]s -var 'hello=world'

  Read environment variables by "env" function

      $ echo 'Current directory is {{ env "PWD" }}' | %[1]s

SEE ALSO
  - [Template Syntax](https://pkg.go.dev/text/template)
  - [Additional Functions](https://masterminds.github.io/sprig/)

`, name)
	if err != nil {
		panic(err)
	}
}
