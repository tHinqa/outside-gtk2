package glib

import "testing"

func TestIni(t *testing.T) {
	t.Log(GetUserName())
	t.Log(GetRealName())
	t.Log(GetApplicationName())
	t.Log(GetPrgname())
	t.Log(FindProgramInPath(GetPrgname()))
	t.Log(GetHostName())
	t.Log(GetHomeDir())
	for _, s := range Listenv() {
		t.Logf("%s=%s", s, Getenv(s))
	}
	t.Log(GetSystemConfigDirs())
	t.Log(GetSystemDataDirs())
	t.Log(GetLanguageNames())
	t.Log(GetLocaleVariants("en_ZA"))
}
