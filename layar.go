package main

import (
	"fmt"
)

func rc(cf *Cf) {
	cl("clipboard, pencet apa aja buat berhenti")
	if !ry(cf) {
		return
	}
	fmt.Println(g0 + "lagi monitor clipboard kamu. coba copy sesuatu." + w0)
	lc(cf, kd(), lg)
}

func rn(cf *Cf) {
	cl("notifikasi, pencet apa aja buat berhenti")
	if !ry(cf) {
		return
	}
	fmt.Println(g0 + "nunggu dari notifikasi kamu..." + w0)
	ln(cf, kd(), lg)
}
