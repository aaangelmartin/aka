package i18n

import "testing"

// TestCatalogParity guards the bilingual rule: every key must exist in
// both language maps. If this test fails, add the missing translation
// instead of removing the key from the other map.
func TestCatalogParity(t *testing.T) {
	for k := range en {
		if _, ok := es[k]; !ok {
			t.Errorf("key %q is in EN but missing in ES", k)
		}
	}
	for k := range es {
		if _, ok := en[k]; !ok {
			t.Errorf("key %q is in ES but missing in EN", k)
		}
	}
}

func TestDetect(t *testing.T) {
	cases := []struct {
		name      string
		cfg, flag string
		env       string
		wantLang  Lang
	}{
		{"flag-en", "", "en", "es_ES.UTF-8", EN},
		{"flag-es", "", "es", "en_US.UTF-8", ES},
		{"cfg-en", "en", "", "es_ES.UTF-8", EN},
		{"cfg-es", "es", "", "en_US.UTF-8", ES},
		{"auto-lang-es", "auto", "", "es_ES.UTF-8", ES},
		{"auto-lang-en", "auto", "", "en_US.UTF-8", EN},
		{"empty-all", "", "", "", EN},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Setenv("LANG", tc.env)
			got := Detect(tc.cfg, tc.flag)
			if got != tc.wantLang {
				t.Fatalf("Detect(%q,%q) env=%q = %s; want %s", tc.cfg, tc.flag, tc.env, got, tc.wantLang)
			}
		})
	}
}

func TestMissingKeyFallback(t *testing.T) {
	Set(EN)
	if got := T("nope.nope"); got == "" || got == "nope.nope" {
		// Empty or bare key would silently show nothing / the key; we
		// want a noisy placeholder.
		t.Fatalf("missing key returned %q; want a visible placeholder", got)
	}
}
