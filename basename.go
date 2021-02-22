package main

import "strings"

func basename(path string) string {
	// cut path by latest '/'
	for i := len(path) - 1; i >= 0; i-- {
		if path[i] == '/' {
			path = path[i+1:]
			break
		}
	}
	// remove suffix
	for i := len(path) - 1; i >= 0; i-- {
		if path[i] == '.' {
			path = path[:i]
			break
		}
	}
	return path
}

func main() {
	println(basename("/path/test/test1.go"))
	println(basename("/test2.go"))
	println(basename("/test3"))
	println(strings.Contains("test, test", "es"))
	println(strings.Count("test, test", "es"))
	println(strings.Fields("test, test")[0])
}
