package glib

import (
	// T "github.com/tHinqa/outside-gtk2/types"
	// m "math"
	"testing"
)

func TestIni(t *testing.T) {
	// t.Log(GetUserName())
	// t.Log(GetRealName())
	// t.Log(GetApplicationName())
	// t.Log(GetPrgname())
	// t.Log(FindProgramInPath(GetPrgname()))
	// t.Log(GetHostName())
	// t.Log(GetHomeDir())
	// for i, s := range Listenv() {
	// 	t.Logf("%d: %s=%s", i, s, Getenv(s))
	// }
	// t.Log(GetSystemConfigDirs())
	// t.Log(GetSystemDataDirs())
	// t.Log(GetLanguageNames())
	// t.Log(GetLocaleVariants("en_ZA"))
	// t.Log(AsciiDtostr(
	// 	&(new([G_ASCII_DTOSTR_BUF_SIZE]T.Gchar)[0]),
	// 	G_ASCII_DTOSTR_BUF_SIZE,
	// 	m.Pi))
	// t.Log(Strtod(AsciiFormatd(
	// 	&(new([G_ASCII_DTOSTR_BUF_SIZE]T.Gchar)[0]),
	// 	G_ASCII_DTOSTR_BUF_SIZE,
	// 	"%g",
	// 	m.Pi), nil))
	Strchug("   aaa   ")
	if Strchomp("   aaa   ") != "   aaa" ||
		Strreverse("hello world") != Strreverse("hello world") ||
		Strdelimit("'a'", "'", '`') == Strdelimit("'a'", "'", '"') {
		t.Error("In place string return failed")
	}
	// for i := 0; i < 5; i++ {
	// 	t.Log(i, Strerror(i))
	// }
	// for i := 0; i < 5; i++ {
	// 	t.Log(i, Strsignal(i))
	// }
	// t.Log(Strconcat("1", "2", "3", "4", "5", "6", "7"))
	// t.Log(Strjoinv("+", []string{"1", "2", "3", "4", "5", "6", "7"}))
	// t.Log(Strsplit(string(Strjoin("+", "1", "2", "3", "4", "5", "6", "7")), "+", 4))
	// t.Log(StrsplitSet("1+2-3+4-5", "+-", -1))
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Listenv()
	}
}

func BenchmarkRand(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandomInt()
	}
}

func TestRand(t *testing.T) {
	RandomSetSeed(0)
	t.Log(RandomInt())
}
