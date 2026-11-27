package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

type Lg func(string, ...any)

func kd() <-chan struct{} {
	dd := make(chan struct{})
	go func() { <-kk; close(dd) }()
	return dd
}

func lg(ff string, aa ...any) {
	fmt.Printf(g0+time.Now().Format("15:04:05")+w0+" "+ff+"\n", aa...)
}

func lc(cf *Cf, dd <-chan struct{}, lo Lg) {
	ll := cg()
	tk := time.NewTicker(600 * time.Millisecond)
	defer tk.Stop()
	ks := 0
	for {
		select {
		case <-dd:
			return
		case <-tk.C:
			nw := cg()
			if nw == "" {
				ks++
				if ks == 100 {
					lo("clipboard kosong terus, kayaknya Android ngeblok baca clipboard pas Termux ga kebuka")
				}
				continue
			}
			ks = 0
			if nw == ll {
				continue
			}
			ll = nw
			lo("kekirim %d karakter, tunggu bentar", len(nw))
			oo, ee := aq(*cf, nw)
			if ee != nil {
				lo("gagal: %v", ee)
				continue
			}
			ll = strings.TrimSpace(oo)
			dl(cf, oo, false)
			lo("dapet %d karakter balik", len(oo))
		}
	}
}

func ln(cf *Cf, dd <-chan struct{}, lo Lg) {
	os.MkdirAll(filepath.Dir(qq), 0700)
	os.Remove(qq)
	nf("SneekCheat", "tap Ask terus ketik")
	tk := time.NewTicker(500 * time.Millisecond)
	defer tk.Stop()
	defer exec.Command("termux-notification-remove", "sneek").Run()
	for {
		select {
		case <-dd:
			return
		case <-tk.C:
			rb, ee := os.ReadFile(qq)
			if ee != nil {
				continue
			}
			os.Remove(qq)
			qs := strings.TrimSpace(string(rb))
			if qs == "" {
				continue
			}
			lo("kamu bilang: %s", qs)
			nf("SneekCheat", "mikir...")
			oo, ee := aq(*cf, qs)
			if ee != nil {
				lo("gagal: %v", ee)
				nf("SneekCheat gagal", ee.Error())
				continue
			}
			dl(cf, oo, true)
			lo("ngirim %d karakter balik", len(oo))
		}
	}
}
