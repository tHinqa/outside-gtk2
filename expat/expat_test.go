package expat

import "testing"

func TestInit(t *testing.T) {
	t.Log("Expat version", XML_ExpatVersion())
	//t.Logf(XML_GetFeatureList().name)
}
