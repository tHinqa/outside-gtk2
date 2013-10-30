package glib

import (
	T "github.com/tHinqa/outside-gtk2/types"
	m "math"
	"testing"
)

func TestIni(t *testing.T) {
	t.Log(GetUserName())
	t.Log(GetRealName())
	t.Log(GetApplicationName())
	t.Log(GetPrgname())
	t.Log(FindProgramInPath(GetPrgname()))
	t.Log(GetHostName())
	t.Log(GetHomeDir())
	for i, s := range Listenv() {
		t.Logf("%d: %s=%s", i, s, Getenv(s))
	}
	t.Log(GetSystemConfigDirs())
	t.Log(GetSystemDataDirs())
	t.Log(GetLanguageNames())
	t.Log(GetLocaleVariants("en_ZA"))
	t.Log(AsciiDtostr(
		&(new([G_ASCII_DTOSTR_BUF_SIZE]T.Gchar)[0]),
		G_ASCII_DTOSTR_BUF_SIZE,
		m.Pi))
	t.Log(Strtod(AsciiFormatd(
		&(new([G_ASCII_DTOSTR_BUF_SIZE]T.Gchar)[0]),
		G_ASCII_DTOSTR_BUF_SIZE,
		"%g",
		m.Pi), nil))
	t.Log(">>>", Strchug("            hello world                 "), "<<<")
	t.Log(">>>", Strchomp("            hello world                 "), "<<<")
	t.Log(Strreverse("hello world"))
	t.Log(Strreverse("hello world"))
	t.Log(Strdelimit("'hello world'", "'", '`'))
	t.Log(Strdelimit("'hello world'", "'", '"'))
	for i := 0; i < 5; i++ {
		t.Log(i, Strerror(i))
	}
	for i := 0; i < 5; i++ {
		t.Log(i, Strsignal(i))
	}
	t.Log(Strconcat("1", "2", "3", "4", "5", "6", "7"))
	t.Log(Strsplit(string(Strjoin("+", "1", "2", "3", "4", "5", "6", "7")), "+", 4))
	t.Log(StrsplitSet("1+2-3+4-5", "+-", -1))
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Listenv()
	}
}
