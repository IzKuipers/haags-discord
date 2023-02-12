package main

import (
	"regexp"
)

type Dict = map[*regexp.Regexp]string

type DictRaw = map[string]string

func convert(input string) string {
	dict := *getDict()

	for key, value := range dict {
		input = key.ReplaceAllLiteralString(input, value)
	}

	return input
}

func getDict() *Dict {
	raw := getRawDict()
	dict := Dict{}

	for k, v := range *raw {
		dict[regexp.MustCompile(k)] = v
	}

	return &dict
}
