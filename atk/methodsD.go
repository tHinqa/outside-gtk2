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

	documentGetAttributes     func(d *Document) *AttributeSet
	documentGetAttributeValue func(d *Document, attributeName string) string
	documentGetDocument       func(d *Document) T.Gpointer
	documentGetDocumentType   func(d *Document) string
	documentGetLocale         func(d *Document) string
	documentSetAttributeValue func(d *Document, attributeName, attributeValue string) T.Gboolean
)

func (d *Document) GetAttributes() *AttributeSet { return documentGetAttributes(d) }
func (d *Document) GetAttributeValue(attributeName string) string {
	return documentGetAttributeValue(d, attributeName)
}
func (d *Document) GetDocument() T.Gpointer { return documentGetDocument(d) }
func (d *Document) GetDocumentType() string { return documentGetDocumentType(d) }
func (d *Document) GetLocale() string       { return documentGetLocale(d) }
func (d *Document) SetAttributeValue(attributeName, attributeValue string) T.Gboolean {
	return documentSetAttributeValue(d, attributeName, attributeValue)
}
