package run

import (
	"fmt"
	"strings"
	"testing"
)

func TestMnemonic_Tidy(t *testing.T) {
	mnemonics := []string{
		"farm employ cup erosion half birth become love excite private swallow grit",
		"farm employ cup erosion half    birth become love excite private swallow grit",
		"farm employ cup erosion half birth become love excite private swallow grit    ",
		"    farm employ cup erosion half birth become love excite private swallow grit",
		"    farm employ cup     erosion half birth become love   excite private swallow grit    ",
	}

	expected := []string{"farm", "employ", "cup", "erosion", "half", "birth", "become", "love", "excite", "private", "swallow", "grit"}

	// test without inserting a newline in mnemonic
	for _, mnemonic := range mnemonics {
		fields := strings.Fields(mnemonic)
		for i, field := range fields {
			if field != expected[i] {
				t.Fatal("expected", expected[i], ", got", field, ", for mnemonic", mnemonic)
			}
		}
	}

	// test by inserting a newline in mnemonic
	for _, mnemonic := range mnemonics {
		mnemonic = fmt.Sprintf("%s\n", mnemonic)
		fields := strings.Fields(mnemonic)
		for i, field := range fields {
			if field != expected[i] {
				t.Fatal("expected", expected[i], ", got", field, ", for mnemonic", mnemonic)
			}
		}
	}
}
