package main

const hello = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return hello + name
}