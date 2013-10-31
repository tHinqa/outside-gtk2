// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package atk

import (
	O "github.com/tHinqa/outside-gtk2/gobject"
	// T "github.com/tHinqa/outside-gtk2/types"
)

type Value struct{}

var (
	ValueGetType func() O.Type

	ValueGetCurrentValue     func(v *Value, value *O.Value)
	ValueGetMaximumValue     func(v *Value, value *O.Value)
	ValueGetMinimumValue     func(v *Value, value *O.Value)
	ValueSetCurrentValue     func(v *Value, value *O.Value) bool
	ValueGetMinimumIncrement func(v *Value, value *O.Value)
)

func (v *Value) GetCurrentValue(value *O.Value)      { ValueGetCurrentValue(v, value) }
func (v *Value) GetMaximumValue(value *O.Value)      { ValueGetMaximumValue(v, value) }
func (v *Value) GetMinimumValue(value *O.Value)      { ValueGetMinimumValue(v, value) }
func (v *Value) SetCurrentValue(value *O.Value) bool { return ValueSetCurrentValue(v, value) }
func (v *Value) GetMinimumIncrement(value *O.Value)  { ValueGetMinimumIncrement(v, value) }
