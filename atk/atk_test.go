package atk

import (
	. "github.com/tHinqa/outside-gtk2/gobject"
	"testing"
)

func TestInit(t *testing.T) {
	GTypeInit()
	//TODO(t): Both ""?
	t.Log(GetToolkitName())
	t.Log(GetToolkitVersion())
}
