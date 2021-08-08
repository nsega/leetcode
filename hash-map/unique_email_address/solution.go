package main

import (
	"fmt"
	"strings"
)

func numUniqueEmails(emails []string) int {

	m := make(map[string]int)
	for _, email := range emails {
		s := strings.Split(email, "@")
		local := strings.Split(s[0], "+")
		l := strings.ReplaceAll(local[0], ".", "")
		email = fmt.Sprintf("%s@%s", l, s[1])
		m[email]++
	}
	return len(m)
}
