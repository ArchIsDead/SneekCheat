package main

import (
	"fmt"
)

func st(cf *Cf) {
	type fl struct {
		kv byte
		nm string
		hp string
		pt *string
		sc bool
	}
	for {
		_, ip := pv[cf.Provider]
		fs := []fl{
			{'1', "worker", "url worker (yang di cloudflare)", &cf.Worker, false},
			{'2', "key", "key enkripsi (samain sama KEY di worker)", &cf.Key, true},
			{'3', "provider", "provider (" + pnm + ") atau nama lain buat api sendiri", &cf.Provider, false},
			{'4', "model", "nama model, ketik persis", &cf.Model, false},
			{'5', "api key", "api key", &cf.APIKey, true},
			{'6', "system", "system prompt (buat AI nya)", &cf.System, false},
			{'7', "output", "jawabannya ke mana, isi: clipboard, notification, atau both", &cf.Out, false},
		}
		if !ip {
			fs = append(fs,
				fl{'8', "url", "url api kamu, bisa pake {prompt} {system} {model} {key}", &cf.URL, false},
				fl{'9', "method", "GET atau POST", &cf.Method, false},
				fl{'a', "body", "template body, placeholdernya sama", &cf.Body, false},
				fl{'b', "path", "posisi jawaban di balasan, contoh data.0.text", &cf.Path, false})
		}
		cl("pengaturan")
		for _, xx := range fs {
			vv := *xx.pt
			if xx.sc && vv != "" {
				vv = "****"
			}
			fmt.Printf("  %s  %s\n", kb(xx.kv), ct(xx.nm+": "+vv, sw-5))
		}
		if cf.Key == "loser123" {
			fmt.Println("\n  awas, key masih loser123. ganti di sini dan di KEY worker")
		}
		fmt.Println("\n  " + kb('q') + "  balik")
		kv := <-kk
		if kv == 'q' {
			return
		}
		for _, xx := range fs {
			if xx.kv != kv {
				continue
			}
			od := *xx.pt
			*xx.pt = pq(xx.hp, od, xx.sc)
			if xx.pt == &cf.Provider && od != cf.Provider {
				cf.Model = pq("nama model buat "+cf.Provider+", ketik sendiri", "", false)
			}
			sv(*cf)
		}
	}
}
