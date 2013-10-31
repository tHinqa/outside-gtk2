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

	valueGetCurrentValue     func(v *Value, value *O.Value)
	valueGetMaximumValue     func(v *Value, value *O.Value)
	valueGetMinimumValue     func(v *Value, value *O.Value)
	valueSetCurrentValue     func(v *Value, value *O.Value) bool
	valueGetMinimumIncrement func(v *Value, value *O.Value)
)

func (v *Value) GetCurrentValue(value *O.Value)      { valueGetCurrentValue(v, value) }
func (v *Value) GetMaximumValue(value *O.Value)      { valueGetMaximumValue(v, value) }
func (v *Value) GetMinimumValue(value *O.Value)      { valueGetMinimumValue(v, value) }
func (v *Value) SetCurrentValue(value *O.Value) bool { return valueSetCurrentValue(v, value) }
func (v *Value) GetMinimumIncrement(value *O.Value)  { valueGetMinimumIncrement(v, value) }
