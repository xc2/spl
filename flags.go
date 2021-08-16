package main

import (
	"os"
	"strings"
)

func ParseExpression(row string) (string, string) {
	sp := strings.SplitN(row, "=", 2)
	return sp[0], sp[1]
}

type mapVar map[string]string

func (m mapVar) Set(s string) error {
	key, value := ParseExpression(s)
	m[key] = value
	return nil
}
func (m mapVar) String() string {
	return ""
}

type fileVar struct {
	file        *os.File
	shouldClose bool
}

func (f *fileVar) Set(s string) error {
	err := f.Close()
	if err != nil {
		return err
	}

	if s == "-" {
		f.file = os.Stdout
		f.shouldClose = false
	} else {
		fi, err := os.OpenFile(s, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			return err
		}
		f.file = fi
		f.shouldClose = true
	}

	return nil
}
func (f *fileVar) String() string {
	return ""
}
func (f *fileVar) Close() error {

	if f.file != nil && f.shouldClose {
		err := f.file.Close()
		if err != nil {
			return err
		}
	}
	return nil
}
