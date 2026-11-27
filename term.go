package main

import (
	"fmt"
	"os"
	"os/exec"
	"unicode/utf8"
)

var (
	kk = make(chan byte, 4096)
	ww = make(chan os.Signal, 1)
	sw = 40
)

func sy(aa ...string) []byte {
	cc := exec.Command("stty", aa...)
	cc.Stdin = os.Stdin
	bb, _ := cc.Output()
	return bb
}

func rk() {
	bb := make([]byte, 256)
	for {
		nn, ee := os.Stdin.Read(bb)
		if ee != nil {
			return
		}
		for _, xx := range bb[:nn] {
			kk <- xx
		}
	}
}

func cw() int {
	var rw, co int
	fmt.Sscan(string(sy("size")), &rw, &co)
	return co
}

func ct(ss string, nn int) string {
	rr := []rune(ss)
	if nn < 2 || len(rr) <= nn {
		return ss
	}
	return string(rr[:nn-1]) + "…"
}

func kb(kv byte) string { return a0 + string(rune(kv)) + w0 }

func cl(tt string) {
	fmt.Print("\x1b[2J\x1b[H\x1b[48;5;0m" + w0 + "\x1b[1mSneekCheat\x1b[22m " + g0 + tt + w0 + "\n\n")
}

func pq(lb, df string, sc bool) string {
	ss := df
	if sc && df != "" {
		ss = "****"
	}
	fmt.Printf("\n%s\n%s\n> ", lb, g0+"["+ct(ss, sw-6)+"]  enter biarin, - hapus"+w0)
	var bb []byte
	for kv := range kk {
		switch {
		case kv == '\r' || kv == '\n':
			fmt.Println()
			switch tx := string(bb); tx {
			case "":
				return df
			case "-":
				return ""
			default:
				return tx
			}
		case kv == 127 || kv == 8:
			if len(bb) > 0 {
				_, nn := utf8.DecodeLastRune(bb)
				bb = bb[:len(bb)-nn]
				fmt.Print("\b \b")
			}
		case kv >= 32:
			bb = append(bb, kv)
			if sc && kv&0xC0 != 0x80 {
				fmt.Print("*")
			} else if !sc {
				os.Stdout.Write([]byte{kv})
			}
		}
	}
	return df
}
