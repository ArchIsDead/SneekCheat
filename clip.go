package main

import (
	"context"
	"os/exec"
	"strings"
	"time"
)

func cg() string {
	xx, ff := context.WithTimeout(context.Background(), 3*time.Second)
	defer ff()
	bb, _ := exec.CommandContext(xx, "termux-clipboard-get").Output()
	return strings.TrimSpace(string(bb))
}

func cs(ss string) {
	cc := exec.Command("termux-clipboard-set")
	cc.Stdin = strings.NewReader(ss)
	cc.Run()
}
