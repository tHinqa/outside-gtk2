// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package atk

import (
	O "github.com/tHinqa/outside-gtk2/gobject"
	T "github.com/tHinqa/outside-gtk2/types"
)

type Document struct{}

var (
	DocumentGetType func() O.Type

	DocumentGetAttributes     func(d *Document) *AttributeSet
	DocumentGetAttributeValue func(d *Document, attributeName string) string
	DocumentGetDocument       func(d *Document) T.Gpointer
	DocumentGetDocumentType   func(d *Document) string
	DocumentGetLocale         func(d *Document) string
	DocumentSetAttributeValue func(d *Document, attributeName, attributeValue string) bool
)

func (d *Document) GetAttributes() *AttributeSet { return DocumentGetAttributes(d) }
func (d *Document) GetAttributeValue(attributeName string) string {
	return DocumentGetAttributeValue(d, attributeName)
}
func (d *Document) GetDocument() T.Gpointer { return DocumentGetDocument(d) }
func (d *Document) GetDocumentType() string { return DocumentGetDocumentType(d) }
func (d *Document) GetLocale() string       { return DocumentGetLocale(d) }
func (d *Document) SetAttributeValue(attributeName, attributeValue string) bool {
	return DocumentSetAttributeValue(d, attributeName, attributeValue)
}
