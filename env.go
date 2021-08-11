package main

import (
	"os"
	"strings"
)

var env map[string]string

func GetEnvirons() map[string]string {
	if env == nil {
		RefreshEnvirons()
	}
	return env
}

func RefreshEnvirons() {
	env = make(map[string]string)

	for _, row := range os.Environ() {
		key, value := ParseEnviron(row)
		env[key] = value
	}
}

func ParseEnviron(row string) (string, string) {
	sp := strings.Split(row, "=")
	return sp[0], strings.Join(sp[1:], "=")
}
