package gobject

import "testing"

func TestInit(t *testing.T) {
	var n uint
	GTypeInit()
	o := ObjectGetType()
	t.Log(ClosureGetType().Name())
	t.Log(ValueGetType().Name())
	t.Log(ValueArrayGetType().Name())
	t.Log(DateGetType().Name())
	t.Log(StrvGetType().Name())
	t.Log(GstringGetType().Name())
	t.Log(HashTableGetType().Name())
	t.Log(ArrayGetType().Name())
	t.Log(ByteArrayGetType().Name())
	t.Log(PtrArrayGetType().Name())
	t.Log(VariantTypeGetGtype().Name())
	t.Log(RegexGetType().Name())
	t.Log(ErrorGetType().Name())
	t.Log(DateTimeGetType().Name())
	t.Log(VariantGetGtype().Name())
	t.Log(BindingFlagsGetType().Qname())
	t.Log(o.Qname())
	t.Log(o.Name())
	t.Log(n, o.Children(&n).Name())
	i := InitiallyUnownedGetType()
	t.Log(i.Qname())
	t.Log(n, o.Children(&n).Name())
	BindingGetType()
	o.Children(&n)
	t.Log(n)
	t.Log(i.Name())
	t.Log(i.Parent().Children(&n).Name())
}
