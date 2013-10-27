// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package cairo

type TextCluster struct {
	NumBytes  int
	NumGlyphs int
}

var (
	TextClusterAllocate func(numClusters int) *TextCluster

	textClusterFree func(t *TextCluster)
)

func (t *TextCluster) Free() { textClusterFree(t) }

type TextClusterFlags Enum

const (
	TEXT_CLUSTER_FLAG_BACKWARD TextClusterFlags = 1 << iota
)

type TextExtents struct {
	XBearing float64
	YBearing float64
	Width    float64
	Height   float64
	XAdvance float64
	YAdvance float64
}
