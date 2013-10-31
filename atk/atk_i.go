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

	imageGetImageDescription func(i *Image) string
	imageGetImageLocale      func(i *Image) string
	imageGetImagePosition    func(i *Image, x, y *int, coordType CoordType)
	imageGetImageSize        func(i *Image, width, height *int)
	imageSetImageDescription func(i *Image, description string) bool
)

func (i *Image) GetImageDescription() string { return imageGetImageDescription(i) }
func (i *Image) GetImageLocale() string      { return imageGetImageLocale(i) }
func (i *Image) GetImagePosition(x, y *int, coordType CoordType) {
	imageGetImagePosition(i, x, y, coordType)
}
func (i *Image) GetImageSize(width, height *int) { imageGetImageSize(i, width, height) }
func (i *Image) SetImageDescription(description string) bool {
	return imageSetImageDescription(i, description)
}

type Implementor struct{}

var (
	ImplementorGetType func() O.Type

	implementorRefAccessible func(i *Implementor) *Object
)

func (i *Implementor) RefAccessible() *Object { return implementorRefAccessible(i) }
