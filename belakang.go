package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"
)

func bj() int {
	bb, ee := os.ReadFile(pf)
	if ee != nil {
		return 0
	}
	pd, ee := strconv.Atoi(strings.TrimSpace(string(bb)))
	if ee != nil || syscall.Kill(pd, 0) != nil {
		return 0
	}
	return pd
}

func lw(ff string, aa ...any) {
	fl, ee := os.OpenFile(lp, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if ee != nil {
		return
	}
	defer fl.Close()
	fmt.Fprintf(fl, time.Now().Format("15:04:05")+" "+ff+"\n", aa...)
}

func bs() string {
	if bj() != 0 {
		return "udah jalan di belakang"
	}
	ex, ee := os.Executable()
	if ee != nil {
		return "gagal: " + ee.Error()
	}
	os.MkdirAll(filepath.Dir(lp), 0700)
	lo, ee := os.OpenFile(lp, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0600)
	if ee != nil {
		return "gagal: " + ee.Error()
	}
	defer lo.Close()
	cc := exec.Command(ex, "jalan")
	cc.Stdout, cc.Stderr = lo, lo
	cc.SysProcAttr = &syscall.SysProcAttr{Setsid: true}
	if ee := cc.Start(); ee != nil {
		return "gagal: " + ee.Error()
	}
	cc.Process.Release()
	time.Sleep(600 * time.Millisecond)
	if bj() == 0 {
		return "ga bisa jalan, cek log di " + lp
	}
	return "jalan di belakang, wake lock nyala"
}

func bk() string {
	pd := bj()
	if pd == 0 {
		return "ga ada yang jalan"
	}
	syscall.Kill(pd, syscall.SIGTERM)
	time.Sleep(500 * time.Millisecond)
	return "udah dimatiin"
}

func dm(cf *Cf) {
	if cf.Worker == "" {
		lw("url worker belum diisi, isi dulu di pengaturan")
		return
	}
	os.MkdirAll(filepath.Dir(pf), 0700)
	os.WriteFile(pf, []byte(strconv.Itoa(os.Getpid())), 0600)
	exec.Command("termux-wake-lock").Run()
	lw("jalan di belakang")
	dd := make(chan struct{})
	sg := make(chan os.Signal, 1)
	signal.Notify(sg, syscall.SIGTERM, os.Interrupt)
	go func() { <-sg; close(dd) }()
	var wg sync.WaitGroup
	wg.Add(2)
	go func() { defer wg.Done(); lc(cf, dd, lw) }()
	go func() { defer wg.Done(); ln(cf, dd, lw) }()
	wg.Wait()
	exec.Command("termux-wake-unlock").Run()
	os.Remove(pf)
	lw("udah berhenti")
}
