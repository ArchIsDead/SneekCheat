package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cf, ex := ld()
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "jalan":
			dm(&cf)
		case "mati":
			fmt.Println(bk())
		default:
			fmt.Println("pake: sneekcheat, sneekcheat jalan, sneekcheat mati")
		}
		return
	}
	ia()
	sy("-icanon", "-echo", "min", "1")
	fmt.Print("\x1b[?25l\x1b[48;5;0m")
	rs := func() { sy("sane"); fmt.Print("\x1b[0m\x1b[?25h\x1b[2J\x1b[H") }
	defer rs()
	sg := make(chan os.Signal, 1)
	signal.Notify(sg, os.Interrupt)
	signal.Notify(ww, syscall.SIGWINCH)
	go func() { <-sg; rs(); os.Exit(0) }()
	go rk()
	if !ex {
		dc(&cf)
		st(&cf)
	}
	mn(&cf)
}
