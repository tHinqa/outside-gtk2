// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package cairo

import (
	// T "github.com/tHinqa/outside-gtk2/types"
)

type Matrix struct {
	Xx, Yx, Xy, Yy, X0, Y0 float64
}

var (
	MatrixInit func(matrix *Matrix, xx, yx, xy, yy, x0, y0 float64)

	MatrixInitIdentity      func(m *Matrix)
	MatrixInitRotate        func(m *Matrix, radians float64)
	MatrixInitScale         func(m *Matrix, sx, sy float64)
	MatrixInitTranslate     func(m *Matrix, tx, ty float64)
	MatrixInvert            func(m *Matrix) Status
	MatrixMultiply          func(m, a, b *Matrix)
	MatrixRotate            func(m *Matrix, radians float64)
	MatrixScale             func(m *Matrix, sx, sy float64)
	MatrixTransformDistance func(m *Matrix, dx, dy *float64)
	MatrixTransformPoint    func(m *Matrix, x, y *float64)
	MatrixTranslate         func(m *Matrix, tx, ty float64)
)

func (m *Matrix) InitIdentity()                     { MatrixInitIdentity(m) }
func (m *Matrix) InitRotate(radians float64)        { MatrixInitRotate(m, radians) }
func (m *Matrix) InitScale(sx, sy float64)          { MatrixInitScale(m, sx, sy) }
func (m *Matrix) InitTranslate(tx, ty float64)      { MatrixInitTranslate(m, tx, ty) }
func (m *Matrix) Invert() Status                    { return MatrixInvert(m) }
func (m *Matrix) Multiply(a, b *Matrix)             { MatrixMultiply(m, a, b) }
func (m *Matrix) Rotate(radians float64)            { MatrixRotate(m, radians) }
func (m *Matrix) Scale(sx, sy float64)              { MatrixScale(m, sx, sy) }
func (m *Matrix) TransformDistance(dx, dy *float64) { MatrixTransformDistance(m, dx, dy) }
func (m *Matrix) TransformPoint(x, y *float64)      { MatrixTransformPoint(m, x, y) }
func (m *Matrix) Translate(tx, ty float64)          { MatrixTranslate(m, tx, ty) }
