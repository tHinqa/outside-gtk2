package atk

import (
	G "github.com/tHinqa/outside-gtk2/gobject"
	"testing"
)

func TestInit(t *testing.T) {
	G.GTypeInit()
	//TODO(t): Both ""?
	t.Log(GetToolkitName())
	t.Log(GetToolkitVersion())
}
