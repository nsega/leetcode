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

func main() {
	type args struct {
		emails []string
	}
	type test struct {
		name string
		arg  args
		exp  int
	}
	tests := []test{
		{
			name: "Example 1:",
			arg: args{
				emails: []string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"},
			},
			exp: 2,
		},
		{
			name: "Example 2:",
			arg: args{
				emails: []string{"a@leetcode.com", "b@leetcode.com", "c@leetcode.com"},
			},
			exp: 3,
		},
	}
	for _, tt := range tests {
		fmt.Println("Case:", tt.name)
		fmt.Println(" Actual:", numUniqueEmails(tt.arg.emails))
		fmt.Println(" Exp:", tt.exp)
	}
}
