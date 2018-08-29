package main

import (
	"flag"
	"learnGo/libray_source_file/demo1"
)

var (
	name string
)

func init() {
	flag.StringVar(&name, "name", "every one", "enter name to say hello")
}

func main() {
	flag.Parse()
	demo.Hello(name)
}
