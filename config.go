package main

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Cf struct {
	Worker, Key, Provider, Model, APIKey, System, Out string
	URL, Method, Body, Path                          string
}

var (
	hmm, _ = os.UserHomeDir()
	pp     = filepath.Join(hmm, ".config", "sneekcheat", "config.json")
	qq     = filepath.Join(hmm, ".cache", "sneekcheat", "reply")
	lp     = filepath.Join(hmm, ".cache", "sneekcheat", "log")
	pf     = filepath.Join(hmm, ".cache", "sneekcheat", "pid")
)

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
