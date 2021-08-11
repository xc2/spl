package main

import (
	"os"
)

type mapVar map[string]string

func (m mapVar) Set(s string) error {
	key, value := ParseEnviron(s)
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

func (f fileVar) Set(s string) error {
	fi, err := os.OpenFile(s, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	err = f.Close()
	if err != nil {
		fi.Close()
		return err
	}
	f.file = fi
	f.shouldClose = true

	return nil
}
func (f fileVar) String() string {
	return ""
}
func (f fileVar) Close() error {

	if f.file != nil && f.shouldClose {
		err := f.file.Close()
		if err != nil {
			return err
		}
	}
	return nil
}
func (f fileVar) IsBoolFlag() bool {
	return false
}
