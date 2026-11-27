package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

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
