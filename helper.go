package main

import (
	"bufio"
	"os"
	"text/template"
)

func ReadStdinToStop() (b []byte, err error) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if b != nil {
			b = append(b, scanner.Bytes()...)
		} else {
			b = scanner.Bytes()
		}
		b = append(b, 0x0a)
		err = scanner.Err()
		if err != nil {
			return
		}
	}
	err = scanner.Err()
	return
}

func TemplateParseFile(t *template.Template, filename string) error {
	if filename != "" && filename != "-" {
		_, err := t.ParseFiles(filename)
		return err
	}
	b, err := ReadStdinToStop()
	if err != nil {
		return err
	}
	s := string(b)

	_, err = t.Parse(s)
	if err != nil {
		return err
	}

	return nil
}
