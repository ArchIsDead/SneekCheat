package main

import (
	"fmt"
)

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
