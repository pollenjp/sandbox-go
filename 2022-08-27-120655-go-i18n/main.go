package main

import (
	"encoding/json"
	"fmt"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var (
	bundle *i18n.Bundle
)

func init() {
	initI18n()
}

func initI18n() {
	bundle = i18n.NewBundle(language.Japanese)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	type i18nTemplate struct {
		lang string
		msg map[string]any
	}

	templates := []i18nTemplate{
		{
			"en",
			map[string]any{
				"MSG002": "Hello, world! (2)",
				"MSG003": "Hello, world! (3)",
			},
		},
		{
			"ja",
			map[string]any{
				"MSG002": "Hello, world! (2)",
				"MSG003": "Hello, world! (3)",
			},
		},
	}
	parse := func(t i18nTemplate) {
		if msg, err := json.Marshal(t.msg); err == nil {
			bundle.MustParseMessageFileBytes(msg, t.lang + ".json")
		}
	}

	for _, t := range templates {
		parse(t)
	}

	assertI18nTemplateMissing := func(ts []i18nTemplate) {
		if len(ts) == 0 {
			return
		}
		t1 := ts[0]
		langs := map[string]bool{}
		langs[t1.lang] = true
		for i, t2 := range ts {
			if i == 0 {
				continue
			}

			// lang must be unique
			if _, ok := langs[t2.lang]; ok {
				panic(fmt.Sprintf("duplicate lang: %s", t2.lang))
			}
			langs[t2.lang] = true

			// key check
			if len(t1.msg) != len(t2.msg) {
				panic(fmt.Sprintf("%s and %s have different number of messages", t1.lang, t2.lang))
			}
			for k := range t1.msg {
				if _, ok := t2.msg[k]; !ok {
					panic(fmt.Sprintf("%s is missing %s", t2.lang, k))
				}
			}
		}
	}
	assertI18nTemplateMissing(templates)

}

func main() {
	{
		localizer := i18n.NewLocalizer(bundle, "ja-JP")
		fmt.Println(localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "MSG002"}))
		fmt.Println(localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "MSG003"}))
	}
}
