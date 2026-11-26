package main

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"
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

const (
	w0  = "\x1b[38;5;255m"
	g0  = "\x1b[38;5;244m"
	a0  = "\x1b[38;5;114m"
	pnm = "chatgpt, deepseek, zai, kimi, openrouter, claude, gemini"
)

type Cf struct {
	Worker, Key, Provider, Model, APIKey, System, Out string
	URL, Method, Body, Path                          string
}

type Sp struct {
	URL     string            `json:"url"`
	Method  string            `json:"method"`
	Headers map[string]string `json:"headers"`
	Body    string            `json:"body"`
	Path    string            `json:"path"`
}

var pv = map[string][2]string{
	"chatgpt":    {"https://api.openai.com/v1/chat/completions", "openai"},
	"deepseek":   {"https://api.deepseek.com/chat/completions", "openai"},
	"zai":        {"https://api.z.ai/api/paas/v4/chat/completions", "openai"},
	"kimi":       {"https://api.moonshot.ai/v1/chat/completions", "openai"},
	"openrouter": {"https://openrouter.ai/api/v1/chat/completions", "openai"},
	"claude":     {"https://api.anthropic.com/v1/messages", "anthropic"},
	"gemini":     {"https://generativelanguage.googleapis.com/v1beta/models/%s:generateContent", "gemini"},
}

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
	hmm, _ = os.UserHomeDir()
	pp     = filepath.Join(hmm, ".config", "sneekcheat", "config.json")
	qq     = filepath.Join(hmm, ".cache", "sneekcheat", "reply")
	kk     = make(chan byte, 4096)
	ww     = make(chan os.Signal, 1)
	a1     []string
	a2     []string
	a3     []string
	w1     int
	w2     int
	w3     int
	cur    []string
	curw   int
	sw     = 40
	spn    = true
)

func g1(pw string) cipher.AEAD {
	hh := sha256.Sum256([]byte(pw))
	bb, _ := aes.NewCipher(hh[:])
	gg, _ := cipher.NewGCM(bb)
	return gg
}

func en(pw string, vv any) string {
	jj, _ := json.Marshal(vv)
	gg := g1(pw)
	nn := make([]byte, gg.NonceSize())
	rand.Read(nn)
	return base64.StdEncoding.EncodeToString(gg.Seal(nn, nn, jj, nil))
}

func de(pw, ss string, vv any) error {
	rr, ee := base64.StdEncoding.DecodeString(strings.TrimSpace(ss))
	gg := g1(pw)
	if ee != nil || len(rr) <= gg.NonceSize() {
		return errors.New("worker ngasih balasan yang tidak valid")
	}
	jj, ee := gg.Open(nil, rr[:gg.NonceSize()], rr[gg.NonceSize():], nil)
	if ee != nil {
		return errors.New("key kamu ga cocok sama punya worker")
	}
	return json.Unmarshal(jj, vv)
}

func bd(cf Cf, pt string) Sp {
	jj := func(vv any) string { bb, _ := json.Marshal(vv); return string(bb) }
	hd := map[string]string{"Content-Type": "application/json"}
	pr, ok := pv[cf.Provider]
	if !ok {
		es := func(ss string) string { ss = jj(ss); return ss[1 : len(ss)-1] }
		bo := strings.NewReplacer("{prompt}", es(pt), "{system}", es(cf.System), "{model}", es(cf.Model), "{key}", es(cf.APIKey)).Replace(cf.Body)
		uu := strings.NewReplacer("{prompt}", url.QueryEscape(pt), "{system}", url.QueryEscape(cf.System), "{model}", url.QueryEscape(cf.Model), "{key}", url.QueryEscape(cf.APIKey)).Replace(cf.URL)
		if cf.APIKey != "" {
			hd["Authorization"] = "Bearer " + cf.APIKey
		}
		mt := strings.ToUpper(cf.Method)
		if mt != "GET" {
			mt = "POST"
		}
		return Sp{uu, mt, hd, bo, cf.Path}
	}
	switch pr[1] {
	case "anthropic":
		hd["x-api-key"], hd["anthropic-version"] = cf.APIKey, "2023-06-01"
		bo := map[string]any{"model": cf.Model, "max_tokens": 4096, "messages": []any{map[string]string{"role": "user", "content": pt}}}
		if cf.System != "" {
			bo["system"] = cf.System
		}
		return Sp{pr[0], "POST", hd, jj(bo), "content.0.text"}
	case "gemini":
		hd["x-goog-api-key"] = cf.APIKey
		bo := map[string]any{"contents": []any{map[string]any{"parts": []any{map[string]string{"text": pt}}}}}
		if cf.System != "" {
			bo["systemInstruction"] = map[string]any{"parts": []any{map[string]string{"text": cf.System}}}
		}
		return Sp{fmt.Sprintf(pr[0], cf.Model), "POST", hd, jj(bo), "candidates.0.content.parts.0.text"}
	}
	hd["Authorization"] = "Bearer " + cf.APIKey
	ms := []any{}
	if cf.System != "" {
		ms = append(ms, map[string]string{"role": "system", "content": cf.System})
	}
	ms = append(ms, map[string]string{"role": "user", "content": pt})
	return Sp{pr[0], "POST", hd, jj(map[string]any{"model": cf.Model, "messages": ms}), "choices.0.message.content"}
}

func aq(cf Cf, pt string) (string, error) {
	rs, ee := (&http.Client{Timeout: 90 * time.Second}).Post(cf.Worker, "text/plain", strings.NewReader(en(cf.Key, bd(cf, pt))))
	if ee != nil {
		return "", ee
	}
	defer rs.Body.Close()
	bb, _ := io.ReadAll(rs.Body)
	if rs.StatusCode != 200 {
		return "", fmt.Errorf("worker bales %d. kayaknya key kamu ga cocok sama KEY di worker", rs.StatusCode)
	}
	var oo struct {
		OK   bool   `json:"ok"`
		Text string `json:"text"`
	}
	if ee := de(cf.Key, string(bb), &oo); ee != nil {
		return "", ee
	}
	if !oo.OK {
		return "", errors.New(oo.Text)
	}
	return oo.Text, nil
}

func ry(cf *Cf) bool {
	_, ip := pv[cf.Provider]
	mm := ""
	switch {
	case cf.Worker == "":
		mm = "yo, url worker belum diisi"
	case ip && cf.Model == "":
		mm = "yo, nama model belum diisi"
	case ip && cf.APIKey == "":
		mm = "yo, api key belum diisi"
	}
	if mm == "" {
		return true
	}
	fmt.Println("\n" + mm + ". isi dulu di pengaturan (3).\n\n" + g0 + "pencet tombol apa aja" + w0)
	<-kk
	return false
}

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

func mn(cf *Cf) {
	tm := 0.0
	for {
		ft()
		fmt.Print("\x1b[2J")
		fm(tm)
		fmt.Printf("\x1b[%d;1H%s", len(cur)+2, w0)
		fmt.Printf("  %s  clipboard\n  %s  notifikasi\n  %s  pengaturan\n  %s  coba dulu\n  %s  cek alat\n  %s  keluar\n\n", kb('1'), kb('2'), kb('3'), kb('4'), kb('5'), kb('q'))
		md := cf.Model
		if md == "" {
			md = "model belum diisi"
		}
		fmt.Println(g0 + "  " + ct(cf.Provider+"  "+md+"  "+cf.Out, sw-2))
		fmt.Println("  spasi buat jeda muter" + w0)
		tk := time.NewTicker(50 * time.Millisecond)
		var kv byte
	wait:
		for {
			select {
			case <-tk.C:
				if spn {
					tm += .09
				}
				fm(tm)
			case <-ww:
				break wait
			case kv = <-kk:
				if kv == ' ' {
					spn = !spn
					continue
				}
				break wait
			}
		}
		tk.Stop()
		switch kv {
		case '1':
			rc(cf)
		case '2':
			rn(cf)
		case '3':
			st(cf)
		case '4':
			cl("coba dulu")
			if !ry(cf) {
				continue
			}
			qs := pq("mau nanya apa", "halo", false)
			fmt.Println(g0 + "lagi nanya..." + w0)
			oo, ee := aq(*cf, qs)
			if ee != nil {
				oo = "gagal: " + ee.Error()
			}
			fmt.Println("\n" + oo + "\n\n" + g0 + "pencet tombol apa aja" + w0)
			<-kk
		case '5':
			dc(cf)
		case 'q':
			return
		}
	}
}

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

func lg(ff string, aa ...any) {
	fmt.Printf(g0+time.Now().Format("15:04:05")+w0+" "+ff+"\n", aa...)
}

func rc(cf *Cf) {
	cl("clipboard, pencet apa aja buat berhenti")
	if !ry(cf) {
		return
	}
	fmt.Println(g0 + "lagi monitor clipboard kamu. coba copy sesuatu." + w0)
	ll := cg()
	tk := time.NewTicker(600 * time.Millisecond)
	defer tk.Stop()
	for {
		select {
		case <-kk:
			return
		case <-tk.C:
			nw := cg()
			if nw == "" || nw == ll {
				continue
			}
			ll = nw
			lg("kekirim %d karakter, tunggu bentar", len(nw))
			oo, ee := aq(*cf, nw)
			if ee != nil {
				lg("gagal: %v", ee)
				continue
			}
			ll = strings.TrimSpace(oo)
			dl(cf, oo, false)
			lg("dapet %d karakter balik", len(oo))
		}
	}
}

func rn(cf *Cf) {
	cl("notifikasi, pencet apa aja buat berhenti")
	if !ry(cf) {
		return
	}
	os.MkdirAll(filepath.Dir(qq), 0700)
	os.Remove(qq)
	nf("SneekCheat", "tap Ask terus ketik")
	fmt.Println(g0 + "nunggu dari notifikasi kamu..." + w0)
	tk := time.NewTicker(500 * time.Millisecond)
	defer tk.Stop()
	defer exec.Command("termux-notification-remove", "sneek").Run()
	for {
		select {
		case <-kk:
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
			lg("kamu bilang: %s", qs)
			nf("SneekCheat", "mikir...")
			oo, ee := aq(*cf, qs)
			if ee != nil {
				lg("gagal: %v", ee)
				nf("SneekCheat gagal", ee.Error())
				continue
			}
			dl(cf, oo, true)
			lg("ngirim %d karakter balik", len(oo))
		}
	}
}

func dc(cf *Cf) {
	cl("cek alat")
	bad := false
	for _, bn := range []string{"termux-clipboard-get", "termux-clipboard-set", "termux-notification"} {
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
	fmt.Println("\n" + g0 + "pencet tombol apa aja" + w0)
	<-kk
}

func ld() (Cf, bool) {
	cf := Cf{Key: "loser123", Provider: "deepseek", Out: "clipboard", Method: "POST"}
	bb, ee := os.ReadFile(pp)
	if ee == nil {
		json.Unmarshal(bb, &cf)
	}
	return cf, ee == nil
}

func sv(cf Cf) {
	os.MkdirAll(filepath.Dir(pp), 0700)
	bb, _ := json.MarshalIndent(cf, "", "  ")
	os.WriteFile(pp, bb, 0600)
}

func main() {
	cf, ex := ld()
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
