package main

import (
	"fmt"
	"math"
	"os"
	"strings"
	"unicode/utf8"
)

const lgo = `
  /$$$$$$                                /$$        /$$$$$$  /$$                             /$$    
 /$$__  $$                              | $$       /$$__  $$| $$                            | $$    
| $$  \__/ /$$$$$$$   /$$$$$$   /$$$$$$ | $$   /$$| $$  \__/| $$$$$$$   /$$$$$$   /$$$$$$  /$$$$$$  
|  $$$$$$ | $$__  $$ /$$__  $$ /$$__  $$| $$  /$$/| $$      | $$__  $$ /$$__  $$ |____  $$|_  $$_/  
 \____  $$| $$  \ $$| $$$$$$$$| $$$$$$$$| $$$$$$/ | $$      | $$  \ $$| $$$$$$$$  /$$$$$$$  | $$    
 /$$  \ $$| $$  | $$| $$_____/| $$_____/| $$_  $$ | $$    $$| $$  | $$| $$_____/ /$$__  $$  | $$ /$$
|  $$$$$$/| $$  | $$|  $$$$$$$|  $$$$$$$| $$ \  $$|  $$$$$$/| $$  | $$|  $$$$$$$|  $$$$$$$  |  $$$$/
 \______/ |__/  |__/ \_______/ \_______/|__/  \__/ \______/ |__/  |__/ \_______/ \_______/   \___/  `

var lt = map[rune][4]string{
	'S': {" ___ ", "/ __|", "\\__ \\", "|___/"},
	'n': {"     ", " _ _  ", "| ' \\ ", "|_||_|"},
	'e': {"     ", " ___ ", "/ -_)", "\\___|"},
	'k': {" _   ", "| |__", "| / /", "|_\\_\\"},
	'C': {" ___ ", "/ __|", "| (__", " \\___|"},
	'h': {" _    ", "| |_  ", "| ' \\ ", "|_||_|"},
	'a': {"      ", " __ _ ", "/ _` |", "\\__,_|"},
	't': {" _   ", "| |_ ", "|  _|", " \\__|"},
}

var (
	a1   []string
	a2   []string
	a3   []string
	w1   int
	w2   int
	w3   int
	cur  []string
	curw int
)

func wb(ss string) []string {
	rr := make([]string, 4)
	for _, cc := range ss {
		ll := lt[cc]
		mw := 0
		for _, xx := range ll {
			mw = max(mw, utf8.RuneCountInString(xx))
		}
		for ii, xx := range ll {
			rr[ii] += xx + strings.Repeat(" ", mw-utf8.RuneCountInString(xx))
		}
	}
	return rr
}

func ia() {
	a1 = strings.Split(strings.TrimPrefix(lgo, "\n"), "\n")
	for _, ll := range a1 {
		w1 = max(w1, utf8.RuneCountInString(ll))
	}
	xx, yy := wb("Sneek"), wb("Cheat")
	xw, yw := utf8.RuneCountInString(xx[0]), utf8.RuneCountInString(yy[0])
	for ii := range xx {
		a2 = append(a2, xx[ii]+" "+yy[ii])
	}
	w2 = xw + 1 + yw
	a3 = append(append([]string{}, xx...), yy...)
	w3 = max(xw, yw)
}

func ft() {
	sw = max(cw()-1, 10)
	switch {
	case sw >= w1:
		cur, curw = a1, w1
	case sw >= w2:
		cur, curw = a2, w2
	case sw >= w3:
		cur, curw = a3, w3
	default:
		cur, curw = []string{"SneekCheat"}, 10
	}
}

func fm(tm float64) {
	var bb strings.Builder
	bb.WriteString("\x1b[H")
	pd := strings.Repeat(" ", max((sw-curw)/2, 0))
	for ii, ll := range cur {
		sn := math.Cos(tm + float64(ii)*0.09)
		aa := math.Abs(sn)
		lv := 250 + int(aa*5)
		if sn < 0 {
			lv = 244 + int(aa*4)
		}
		fmt.Fprintf(&bb, "\x1b[38;5;%dm%s", lv, pd)
		rr := []rune(ll)
		for xx := 0; xx < curw; xx++ {
			ch := ' '
			if aa > .06 {
				if qi := int(math.Floor((float64(xx)-float64(curw)/2)/sn + float64(curw)/2)); qi >= 0 && qi < len(rr) {
					ch = rr[qi]
				}
			}
			bb.WriteRune(ch)
		}
		bb.WriteString("\x1b[K\n")
	}
	os.Stdout.WriteString(bb.String())
}
