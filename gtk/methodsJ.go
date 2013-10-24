// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gtk

import (
	T "github.com/tHinqa/outside-gtk2/types"
	// . "github.com/tHinqa/outside/types"
)

type Justification Enum

const (
	JUSTIFY_LEFT Justification = iota
	JUSTIFY_RIGHT
	JUSTIFY_CENTER
	JUSTIFY_FILL
)

var JustificationGetType func() T.GType
