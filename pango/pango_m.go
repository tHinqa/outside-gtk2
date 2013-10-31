// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package pango

import (
	L "github.com/tHinqa/outside-gtk2/glib"
	O "github.com/tHinqa/outside-gtk2/gobject"
	// T "github.com/tHinqa/outside-gtk2/types"
)

type Map struct{}

var (
	MapGetEngine  func(m *Map, script Script) *Engine
	MapGetEngines func(m *Map, script Script, exactEngines, fallbackEngines **L.SList)
)

func (m *Map) GetEngine(script Script) *Engine { return MapGetEngine(m, script) }
func (m *Map) GetEngines(script Script, exactEngines, fallbackEngines **L.SList) {
	MapGetEngines(m, script, exactEngines, fallbackEngines)
}

type Matrix struct {
	Xx float64
	Xy float64
	Yx float64
	Yy float64
	X0 float64
	Y0 float64
}

var (
	MatrixGetType func() O.Type

	MatrixCopy                    func(m *Matrix) *Matrix
	MatrixFree                    func(m *Matrix)
	MatrixTranslate               func(m *Matrix, tx, ty float64)
	MatrixScale                   func(m *Matrix, scaleX, scaleY float64)
	MatrixRotate                  func(m *Matrix, degrees float64)
	MatrixConcat                  func(m, newMatrix *Matrix)
	MatrixTransformPoint          func(m *Matrix, x, y *float64)
	MatrixTransformDistance       func(m *Matrix, dx, dy *float64)
	MatrixTransformRectangle      func(m *Matrix, rect *Rectangle)
	MatrixTransformPixelRectangle func(m *Matrix, rect *Rectangle)
	MatrixGetFontScaleFactor      func(m *Matrix) float64
)

func (m *Matrix) Copy() *Matrix                      { return MatrixCopy(m) }
func (m *Matrix) Free()                              { MatrixFree(m) }
func (m *Matrix) Translate(tx, ty float64)           { MatrixTranslate(m, tx, ty) }
func (m *Matrix) Scale(scaleX, scaleY float64)       { MatrixScale(m, scaleX, scaleY) }
func (m *Matrix) Rotate(degrees float64)             { MatrixRotate(m, degrees) }
func (m *Matrix) Concat(newMatrix *Matrix)           { MatrixConcat(m, newMatrix) }
func (m *Matrix) TransformPoint(x, y *float64)       { MatrixTransformPoint(m, x, y) }
func (m *Matrix) TransformDistance(dx, dy *float64)  { MatrixTransformDistance(m, dx, dy) }
func (m *Matrix) TransformRectangle(rect *Rectangle) { MatrixTransformRectangle(m, rect) }
func (m *Matrix) TransformPixelRectangle(rect *Rectangle) {
	MatrixTransformPixelRectangle(m, rect)
}
func (m *Matrix) GetFontScaleFactor() float64 { return MatrixGetFontScaleFactor(m) }

var ModuleRegister func(module *IncludedModule)
