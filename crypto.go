package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"
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
