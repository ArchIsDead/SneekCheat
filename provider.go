package main

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

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
