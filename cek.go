package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
)

func dc(cf *Cf) {
	cl("cek alat")
	bad := false
	for _, bn := range []string{"termux-clipboard-get", "termux-clipboard-set", "termux-notification", "termux-wake-lock"} {
		_, ee := exec.LookPath(bn)
		fmt.Printf("  %s  %s\n", map[bool]string{true: "ada ", false: "gada"}[ee == nil], bn)
		bad = bad || ee != nil
	}
	if bad && pq("mau diinstalin termux-api? y atau n", "y", false) == "y" {
		xc := exec.Command("pkg", "install", "-y", "termux-api")
		xc.Stdout, xc.Stderr = os.Stdout, os.Stderr
		xc.Run()
	}
	nn, ff := context.WithTimeout(context.Background(), 3*time.Second)
	defer ff()
	if ee := exec.CommandContext(nn, "termux-clipboard-get").Run(); ee != nil {
		fmt.Println("\n  gada  app Termux:API ga nyahut")
		fmt.Println("  ambil dari tempat kamu dapet Termux, terus buka sekali")
		fmt.Println("  https://f-droid.org/packages/com.termux.api/")
		fmt.Println("  https://github.com/termux/termux-api/releases")
	} else {
		fmt.Println("\n  ada   app Termux:API nyahut")
	}
	if cf.Worker == "" {
		fmt.Println("  gada  url worker belum diisi")
	} else if rs, ee := http.Post(cf.Worker, "text/plain", strings.NewReader("x")); ee != nil || rs.StatusCode != 401 {
		fmt.Println("  gada  worker itu ga pas dengan worker SneekCheat")
	} else {
		rs.Body.Close()
		fmt.Println("  ada   worker on")
	}
	fmt.Println("\n  catatan: Android 10 ke atas ngeblok baca clipboard pas Termux ga kebuka.")
	fmt.Println("  jalur notifikasi tetep jalan di belakang, clipboard belum tentu.")
	fmt.Println("\n" + g0 + "pencet tombol apa aja" + w0)
	<-kk
}
