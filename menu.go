package main

import (
	"fmt"
	"time"
)

var (
	spn = true
	ps  string
)

func mn(cf *Cf) {
	tm := 0.0
	for {
		ft()
		fmt.Print("\x1b[2J")
		fm(tm)
		fmt.Printf("\x1b[%d;1H%s", len(cur)+2, w0)
		bt := "mati"
		if bj() != 0 {
			bt = "nyala"
		}
		fmt.Printf("  %s  clipboard\n  %s  notifikasi\n  %s  pengaturan\n  %s  coba dulu\n  %s  cek alat\n  %s  di belakang [%s]\n  %s  keluar\n\n", kb('1'), kb('2'), kb('3'), kb('4'), kb('5'), kb('6'), bt, kb('q'))
		md := cf.Model
		if md == "" {
			md = "model belum diisi"
		}
		fmt.Println(g0 + "  " + ct(cf.Provider+"  "+md+"  "+cf.Out, sw-2))
		fmt.Println("  spasi buat jeda muter" + w0)
		if ps != "" {
			fmt.Println(g0 + "  " + ct(ps, sw-2) + w0)
			ps = ""
		}
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
		case '6':
			switch {
			case bj() != 0:
				ps = bk()
			case cf.Worker == "":
				ps = "isi url worker dulu di pengaturan (3)"
			default:
				ps = bs()
			}
		case 'q':
			return
		}
	}
}
