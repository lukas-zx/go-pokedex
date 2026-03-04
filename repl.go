package main

import "strings"

func cleanInput(text string) []string {
	res := []string{}
	words := strings.FieldsSeq(text)
	for word := range words {
		res = append(res, strings.ToLower(strings.TrimSpace(word)))
	}
	return res
}
