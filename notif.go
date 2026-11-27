package main

import (
	"os/exec"
)

func nf(tt, ss string) {
	exec.Command("termux-notification", "--id", "sneek", "--title", tt, "--content", ss,
		"--priority", "high", "--alert-once",
		"--button1", "Ask", "--button1-action", `printf '%s' "$REPLY" > '`+qq+`'`).Run()
}

func dl(cf *Cf, oo string, nn bool) {
	if cf.Out != "notification" {
		cs(oo)
	}
	if cf.Out != "clipboard" {
		nf("SneekCheat", oo)
	} else if nn {
		nf("SneekCheat", "udah kecopy, tinggal paste")
	}
}
