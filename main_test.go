package main

import (
	"encoding/json"
	"testing"
)

func TestSeal(t *testing.T) {
	var out map[string]string
	if err := unseal("k", seal("k", map[string]string{"a": "b"}), &out); err != nil || out["a"] != "b" {
		t.Fatal("roundtrip failed", err)
	}
	if unseal("x", seal("k", 1), &out) == nil {
		t.Fatal("wrong key accepted")
	}
}

func TestCustomBody(t *testing.T) {
	s := build(Cfg{Provider: "x", Method: "post", Body: `{"q":"{prompt}"}`}, "say \"hi\"\n")
	var m map[string]string
	if json.Unmarshal([]byte(s.Body), &m) != nil || m["q"] != "say \"hi\"\n" {
		t.Fatal(s.Body)
	}
}
