package main

import "log"

func init() {
	log.SetPrefix("TRACE: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

func main() {
	log.Println("hello")
	log.Fatalln("fatal message")
	log.Panicln("panic message")
}
