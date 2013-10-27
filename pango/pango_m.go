// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package pango

import (
	O "github.com/tHinqa/outside-gtk2/gobject"
	T "github.com/tHinqa/outside-gtk2/types"
)

type Map struct{}

var (
	mapGetEngine  func(m *Map, script Script) *Engine
	mapGetEngines func(m *Map, script Script, exactEngines, fallbackEngines **T.GSList)
)

func (m *Map) GetEngine(script Script) *Engine { return mapGetEngine(m, script) }
func (m *Map) GetEngines(script Script, exactEngines, fallbackEngines **T.GSList) {
	mapGetEngines(m, script, exactEngines, fallbackEngines)
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

	matrixCopy                    func(m *Matrix) *Matrix
	matrixFree                    func(m *Matrix)
	matrixTranslate               func(m *Matrix, tx, ty float64)
	matrixScale                   func(m *Matrix, scaleX, scaleY float64)
	matrixRotate                  func(m *Matrix, degrees float64)
	matrixConcat                  func(m, newMatrix *Matrix)
	matrixTransformPoint          func(m *Matrix, x, y *float64)
	matrixTransformDistance       func(m *Matrix, dx, dy *float64)
	matrixTransformRectangle      func(m *Matrix, rect *Rectangle)
	matrixTransformPixelRectangle func(m *Matrix, rect *Rectangle)
	matrixGetFontScaleFactor      func(m *Matrix) float64
)

func (m *Matrix) Copy() *Matrix                      { return matrixCopy(m) }
func (m *Matrix) Free()                              { matrixFree(m) }
func (m *Matrix) Translate(tx, ty float64)           { matrixTranslate(m, tx, ty) }
func (m *Matrix) Scale(scaleX, scaleY float64)       { matrixScale(m, scaleX, scaleY) }
func (m *Matrix) Rotate(degrees float64)             { matrixRotate(m, degrees) }
func (m *Matrix) Concat(newMatrix *Matrix)           { matrixConcat(m, newMatrix) }
func (m *Matrix) TransformPoint(x, y *float64)       { matrixTransformPoint(m, x, y) }
func (m *Matrix) TransformDistance(dx, dy *float64)  { matrixTransformDistance(m, dx, dy) }
func (m *Matrix) TransformRectangle(rect *Rectangle) { matrixTransformRectangle(m, rect) }
func (m *Matrix) TransformPixelRectangle(rect *Rectangle) {
	matrixTransformPixelRectangle(m, rect)
}
func (m *Matrix) GetFontScaleFactor() float64 { return matrixGetFontScaleFactor(m) }

var ModuleRegister func(module *IncludedModule)
