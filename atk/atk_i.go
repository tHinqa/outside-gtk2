// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package atk

import (
	O "github.com/tHinqa/outside-gtk2/gobject"
	// T "github.com/tHinqa/outside-gtk2/types"
)

type Image struct{}

var (
	ImageGetType func() O.Type

	ImageGetImageDescription func(i *Image) string
	ImageGetImageLocale      func(i *Image) string
	ImageGetImagePosition    func(i *Image, x, y *int, coordType CoordType)
	ImageGetImageSize        func(i *Image, width, height *int)
	ImageSetImageDescription func(i *Image, description string) bool
)

func (i *Image) GetImageDescription() string { return ImageGetImageDescription(i) }
func (i *Image) GetImageLocale() string      { return ImageGetImageLocale(i) }
func (i *Image) GetImagePosition(x, y *int, coordType CoordType) {
	ImageGetImagePosition(i, x, y, coordType)
}
func (i *Image) GetImageSize(width, height *int) { ImageGetImageSize(i, width, height) }
func (i *Image) SetImageDescription(description string) bool {
	return ImageSetImageDescription(i, description)
}

type Implementor struct{}

var (
	ImplementorGetType func() O.Type

	ImplementorRefAccessible func(i *Implementor) *Object
)

func (i *Implementor) RefAccessible() *Object { return ImplementorRefAccessible(i) }
