// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gtk

import (
	T "github.com/tHinqa/outside-gtk2/types"
	// . "github.com/tHinqa/outside/types"
)

type KeySnoopFunc func(grab_widget *Widget,
	event *T.GdkEventKey, funcData T.Gpointer) int
