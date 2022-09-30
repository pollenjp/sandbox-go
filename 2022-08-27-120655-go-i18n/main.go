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
		msg  map[string]any
	}

	templates := []i18nTemplate{
		{
			"en",
			map[string]any{
				"MSG002": "Hello, world! (2)",
				"MSG003": "Hello, world! (3)",
				"MSG004": map[string]any{
					"description": "desc",
					"one":         "{{ .Hello }} {{ .PluralCount }} one",
					"other":       "{{ .Hello }} {{ .PluralCount }} other",
				},
			},
		},
		{
			"ja",
			map[string]any{
				"MSG002": "JP: Hello, world! (2)",
				"MSG003": "JP: Hello, world! (3)",
				"MSG004": map[string]any{
					"description": "JP: desc",
					"zero":        "JP: {{ .Hello }} zero",
					"one":         "JP: {{ .Hello }} one",
					"other":       "JP: {{ .Hello }} other",
				},
			},
		},
	}
	parse := func(t i18nTemplate) {
		if msg, err := json.Marshal(t.msg); err == nil {
			fmt.Println(string(msg))
			bundle.MustParseMessageFileBytes(msg, t.lang+".json")
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
	fmt.Printf("%v\n", bundle)
	fmt.Printf("%s\n", language.English)
	fmt.Printf("%s\n", language.Japanese)
	{
		localizer := i18n.NewLocalizer(bundle, language.Japanese.String())
		fmt.Println(localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "MSG002"}))
		fmt.Println(localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "MSG003"}))
		tmp := i18n.LocalizeConfig{
			MessageID: "MSG004",
			TemplateData: map[string]interface{}{
				"Hello": "HELLO",
			},
			PluralCount: 1,
		}
		fmt.Printf("%v\n", tmp)
		fmt.Printf("%v\n", tmp.PluralCount)
		fmt.Println(localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: "MSG004",
			TemplateData: map[string]interface{}{
				"Hello": "HELLO",
			},
			PluralCount: 0, // 日本語は対象言語ではない
		}))
	}
	{
		localizer := i18n.NewLocalizer(bundle, language.English.String())
		cnt := 1
		fmt.Println(localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: "MSG004",
			TemplateData: map[string]interface{}{
				"Hello":       "HELLO",
				"PluralCount": fmt.Sprint(cnt),
			},
			PluralCount: cnt,
		}))
	}
}
