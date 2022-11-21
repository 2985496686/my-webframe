package gee

import (
	"fmt"
	"strings"
	"testing"
)

func ParseRouter(router string) []string {
	split := strings.Split(router, "/")
	parts := make([]string, 0)
	for _, part := range split {
		if part != "" {
			parts = append(parts, part)
			if part[0] == '*' {
				return parts
			}
		}
	}
	return parts
}

func TestTire(t *testing.T) {
	root := &node{}
	router1 := "/user/get/:id"
	router2 := "/user/add"
	router3 := "/user/*"
	root.insert(router1, ParseRouter(router1), -1)
	root.insert(router2, ParseRouter(router2), -1)
	root.insert(router3, ParseRouter(router3), -1)
	fmt.Println(root.search(ParseRouter(router1), -1).pattern)
	fmt.Println(root.search(ParseRouter("/user/name"), -1).pattern)
	//root.insert("")
}
