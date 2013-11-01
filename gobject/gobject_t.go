// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package gobject

import (
	// O "github.com/tHinqa/outside-gtk2/gobject"
	T "github.com/tHinqa/outside-gtk2/types"
	// . "github.com/tHinqa/outside/types"
)

var (
	TypeFromName func(name string) Type

	TypeAddClassCacheFunc     func(cacheData T.Gpointer, cacheFunc T.GTypeClassCacheFunc)
	TypeAddInterfaceCheck     func(checkData T.Gpointer, checkFunc T.GTypeInterfaceCheckFunc)
	TypeCheckValue            func(value *Value) bool
	TypeCheckValueHolds       func(value *Value, t Type) bool
	TypeClassAddPrivate       func(gClass T.Gpointer, privateSize T.Gsize)
	TypeClassPeekParent       func(gClass T.Gpointer) T.Gpointer
	TypeClassUnref            func(gClass T.Gpointer)
	TypeClassUnrefUncached    func(gClass T.Gpointer)
	TypeDefaultInterfaceUnref func(gIface T.Gpointer)
	TypeFundamentalNext       func() Type
	TypeInitWithDebugFlags    func(debugFlags TypeDebugFlags)
	TypeInterfacePeek         func(instanceClass T.Gpointer, ifaceType Type) T.Gpointer
	TypeInterfacePeekParent   func(gIface T.Gpointer) T.Gpointer
	TypeRemoveClassCacheFunc  func(cacheData T.Gpointer, cacheFunc T.GTypeClassCacheFunc)
	TypeRemoveInterfaceCheck  func(checkData T.Gpointer, checkFunc T.GTypeInterfaceCheckFunc)

	TypeAddClassPrivate          func(classType Type, privateSize T.Gsize)
	TypeAddInterfaceDynamic      func(instanceType Type, interfaceType Type, plugin *TypePlugin)
	TypeAddInterfaceStatic       func(instanceType Type, interfaceType Type, info *T.GInterfaceInfo)
	TypeCheckIsValueType         func(t Type) bool
	TypeChildren                 func(t Type, nChildren *uint) *Type
	TypeClassPeek                func(t Type) T.Gpointer
	TypeClassPeekStatic          func(t Type) T.Gpointer
	TypeClassRef                 func(t Type) T.Gpointer
	TypeDefaultInterfacePeek     func(gType Type) T.Gpointer
	TypeDefaultInterfaceRef      func(gType Type) T.Gpointer
	TypeDepth                    func(t Type) uint
	TypeFundamental              func(typeId Type) Type
	TypeGetPlugin                func(t Type) *TypePlugin
	TypeGetQdata                 func(t Type, quark T.Quark) T.Gpointer
	TypeInterfaceAddPrerequisite func(interfaceType Type, prerequisiteType Type)
	TypeInterfaceGetPlugin       func(instanceType Type, interfaceType Type) *TypePlugin
	TypeInterfacePrerequisites   func(interfaceType Type, nPrerequisites *uint) *Type
	TypeInterfaces               func(t Type, nInterfaces *uint) *Type
	TypeIsA                      func(t, isAType Type) bool
	TypeName                     func(t Type) string
	TypeNextBase                 func(leafType, rootType Type) Type
	TypeParent                   func(t Type) Type
	TypeQname                    func(t Type) T.Quark
	TypeQuery                    func(t Type, query *T.GTypeQuery)
	TypeRegisterDynamic          func(parentType Type, typeName string, plugin *TypePlugin, flags TypeFlags) Type
	TypeRegisterFundamental      func(typeId Type, typeName string, info *TypeInfo, finfo *T.GTypeFundamentalInfo, flags TypeFlags) Type
	TypeRegisterStatic           func(parentType Type, typeName string, info *TypeInfo, flags TypeFlags) Type
	TypeRegisterStaticSimple     func(parentType Type, typeName string, classSize uint, classInit T.GClassInitFunc, instanceSize uint, instanceInit T.GInstanceInitFunc, flags TypeFlags) Type
	TypeSetQdata                 func(t Type, quark T.Quark, data T.Gpointer)
	TypeTestFlags                func(t Type, flags uint) bool
	TypeValueTablePeek           func(t Type) *T.GTypeValueTable
)

func (t Type) AddClassPrivate(privateSize T.Gsize) { TypeAddClassPrivate(t, privateSize) }
func (t Type) AddInterfaceDynamic(interfaceType Type, plugin *TypePlugin) {
	TypeAddInterfaceDynamic(t, interfaceType, plugin)
}
func (t Type) AddInterfaceStatic(interfaceType Type, info *T.GInterfaceInfo) {
	TypeAddInterfaceStatic(t, interfaceType, info)
}
func (t Type) CheckIsValueType() bool            { return TypeCheckIsValueType(t) }
func (t Type) Children(nChildren *uint) *Type    { return TypeChildren(t, nChildren) }
func (t Type) ClassPeek() T.Gpointer             { return TypeClassPeek(t) }
func (t Type) ClassPeekStatic() T.Gpointer       { return TypeClassPeekStatic(t) }
func (t Type) ClassRef() T.Gpointer              { return TypeClassRef(t) }
func (t Type) DefaultInterfacePeek() T.Gpointer  { return TypeDefaultInterfacePeek(t) }
func (t Type) DefaultInterfaceRef() T.Gpointer   { return TypeDefaultInterfaceRef(t) }
func (t Type) Depth() uint                       { return TypeDepth(t) }
func (t Type) Fundamental() Type                 { return TypeFundamental(t) }
func (t Type) GetPlugin() *TypePlugin            { return TypeGetPlugin(t) }
func (t Type) GetQdata(quark T.Quark) T.Gpointer { return TypeGetQdata(t, quark) }
func (t Type) InterfaceAddPrerequisite(prerequisiteType Type) {
	TypeInterfaceAddPrerequisite(t, prerequisiteType)
}
func (t Type) InterfaceGetPlugin(interfaceType Type) *TypePlugin {
	return TypeInterfaceGetPlugin(t, interfaceType)
}
func (t Type) InterfacePrerequisites(nPrerequisites *uint) *Type {
	return TypeInterfacePrerequisites(t, nPrerequisites)
}
func (t Type) Interfaces(nInterfaces *uint) *Type { return TypeInterfaces(t, nInterfaces) }
func (t Type) IsA(isAType Type) bool              { return TypeIsA(t, isAType) }
func (t Type) Name() string                       { return TypeName(t) }
func (t Type) NextBase(rootType Type) Type        { return TypeNextBase(t, rootType) }
func (t Type) Parent() Type                       { return TypeParent(t) }
func (t Type) Qname() T.Quark                     { return TypeQname(t) }
func (t Type) Query(query *T.GTypeQuery)          { TypeQuery(t, query) }
func (t Type) RegisterDynamic(typeName string, plugin *TypePlugin, flags TypeFlags) Type {
	return TypeRegisterDynamic(t, typeName, plugin, flags)
}
func (t Type) RegisterFundamental(typeName string, info *TypeInfo, finfo *T.GTypeFundamentalInfo, flags TypeFlags) Type {
	return TypeRegisterFundamental(t, typeName, info, finfo, flags)
}
func (t Type) RegisterStatic(typeName string, info *TypeInfo, flags TypeFlags) Type {
	return TypeRegisterStatic(t, typeName, info, flags)
}
func (t Type) RegisterStaticSimple(typeName string, classSize uint, classInit T.GClassInitFunc, instanceSize uint, instanceInit T.GInstanceInitFunc, flags TypeFlags) Type {
	return TypeRegisterStaticSimple(t, typeName, classSize, classInit, instanceSize, instanceInit, flags)
}
func (t Type) SetQdata(quark T.Quark, data T.Gpointer) { TypeSetQdata(t, quark, data) }
func (t Type) TestFlags(flags uint) bool               { return TypeTestFlags(t, flags) }
func (t Type) ValueTablePeek() *T.GTypeValueTable      { return TypeValueTablePeek(t) }

type TypeClass struct {
	Type Type
}

var (
	TypeCheckClassCast  func(class *TypeClass, isAType Type) *TypeClass
	TypeCheckClassIsA   func(class *TypeClass, isAType Type) bool
	TypeClassGetPrivate func(class *TypeClass, privateType Type) T.Gpointer
	TypeNameFromClass   func(class *TypeClass) string
)

func (t *TypeClass) CheckCast(isAType Type) *TypeClass { return TypeCheckClassCast(t, isAType) }
func (t *TypeClass) GetPrivate(privateType Type) T.Gpointer {
	return TypeClassGetPrivate(t, privateType)
}
func (t *TypeClass) IsA(isAType Type) bool { return TypeCheckClassIsA(t, isAType) }
func (t *TypeClass) Name() string          { return TypeNameFromClass(t) }

type TypeDebugFlags Enum

const (
	TYPE_DEBUG_OBJECTS TypeDebugFlags = 1 << iota
	TYPE_DEBUG_SIGNALS
	TYPE_DEBUG_NONE TypeDebugFlags = 0
	TYPE_DEBUG_MASK TypeDebugFlags = 0x03
)

type TypeFlags Enum

const (
	TYPE_FLAG_ABSTRACT TypeFlags = 1 << (4 + iota)
	TYPE_FLAG_VALUE_ABSTRACT
)

type TypeInfo struct {
	ClassSize     uint16
	BaseInit      T.GBaseInitFunc
	BaseFinalize  T.GBaseFinalizeFunc
	ClassInit     T.GClassInitFunc
	ClassFinalize T.GClassFinalizeFunc
	ClassData     T.Gconstpointer
	InstanceSize  uint16
	NPreallocs    uint16
	InstanceInit  T.GInstanceInitFunc
	ValueTable    *T.GTypeValueTable
}

type TypeInstance struct {
	Class *TypeClass
}

var (
	TypeCreateInstance func(t Type) *TypeInstance

	TypeCheckInstance      func(i *TypeInstance) bool
	TypeCheckInstanceCast  func(i *TypeInstance, ifaceType Type) *TypeInstance
	TypeCheckInstanceIsA   func(i *TypeInstance, ifaceType Type) bool
	TypeFreeInstance       func(i *TypeInstance)
	TypeInstanceGetPrivate func(i *TypeInstance, privateType Type) T.Gpointer
	TypeNameFromInstance   func(i *TypeInstance) string
)

func (i *TypeInstance) CheckInstance() bool { return TypeCheckInstance(i) }
func (i *TypeInstance) CheckCast(ifaceType Type) *TypeInstance {
	return TypeCheckInstanceCast(i, ifaceType)
}
func (i *TypeInstance) IsA(ifaceType Type) bool {
	return TypeCheckInstanceIsA(i, ifaceType)
}
func (i *TypeInstance) Free() { TypeFreeInstance(i) }
func (i *TypeInstance) GetPrivate(privateType Type) T.Gpointer {
	return TypeInstanceGetPrivate(i, privateType)
}
func (i *TypeInstance) NameFromInstance() string { return TypeNameFromInstance(i) }

type TypeModule struct {
	Parent         Object
	UseCount       uint
	TypeInfos      *SList
	InterfaceInfos *SList
	Name           *T.Gchar
}

var (
	TypeModuleGetType func() Type

	TypeModuleAddInterface  func(m *TypeModule, instanceType, interfaceType Type, interfaceInfo *T.GInterfaceInfo)
	TypeModuleRegisterEnum  func(m *TypeModule, name string, constStaticValues *EnumValue) Type
	TypeModuleRegisterFlags func(m *TypeModule, name string, constStaticValues *FlagsValue) Type
	TypeModuleRegisterType  func(m *TypeModule, parentType Type, typeName string, typeInfo *TypeInfo, flags TypeFlags) Type
	TypeModuleSetName       func(m *TypeModule, name string)
	TypeModuleUnuse         func(m *TypeModule)
	TypeModuleUse           func(m *TypeModule) bool
)

func (m *TypeModule) AddInterface(instanceType, interfaceType Type, interfaceInfo *T.GInterfaceInfo) {
	TypeModuleAddInterface(m, instanceType, interfaceType, interfaceInfo)
}
func (m *TypeModule) RegisterEnum(name string, constStaticValues *EnumValue) Type {
	return TypeModuleRegisterEnum(m, name, constStaticValues)
}
func (m *TypeModule) RegisterFlags(name string, constStaticValues *FlagsValue) Type {
	return TypeModuleRegisterFlags(m, name, constStaticValues)
}
func (m *TypeModule) RegisterType(parentType Type, typeName string, typeInfo *TypeInfo, flags TypeFlags) Type {
	return TypeModuleRegisterType(m, parentType, typeName, typeInfo, flags)
}
func (m *TypeModule) SetName(name string) { TypeModuleSetName(m, name) }
func (m *TypeModule) Unuse()              { TypeModuleUnuse(m) }
func (m *TypeModule) Use() bool           { return TypeModuleUse(m) }

type TypePlugin struct{}

var (
	TypePluginGetType func() Type

	TypePluginCompleteInterfaceInfo func(p *TypePlugin, instanceType, interfaceType Type, info *T.GInterfaceInfo)
	TypePluginCompleteTypeInfo      func(p *TypePlugin, gType Type, info *TypeInfo, valueTable *T.GTypeValueTable)
	TypePluginUnuse                 func(p *TypePlugin)
	TypePluginUse                   func(p *TypePlugin)
)

func (p *TypePlugin) CompleteInterfaceInfo(instanceType, interfaceType Type, info *T.GInterfaceInfo) {
	TypePluginCompleteInterfaceInfo(p, instanceType, interfaceType, info)
}
func (p *TypePlugin) CompleteTypeInfo(gType Type, info *TypeInfo, valueTable *T.GTypeValueTable) {
	TypePluginCompleteTypeInfo(p, gType, info, valueTable)
}
func (p *TypePlugin) Unuse() { TypePluginUnuse(p) }
func (p *TypePlugin) Use()   { TypePluginUse(p) }
