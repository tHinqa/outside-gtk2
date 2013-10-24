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
	TypeCheckValue            func(value *Value) T.Gboolean
	TypeCheckValueHolds       func(value *Value, t Type) T.Gboolean
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

	typeAddClassPrivate          func(classType Type, privateSize T.Gsize)
	typeAddInterfaceDynamic      func(instanceType Type, interfaceType Type, plugin *TypePlugin)
	typeAddInterfaceStatic       func(instanceType Type, interfaceType Type, info *T.GInterfaceInfo)
	typeCheckIsValueType         func(t Type) T.Gboolean
	typeCheckValue               func(value *Value) T.Gboolean
	typeCheckValueHolds          func(value *Value, t Type) T.Gboolean
	typeChildren                 func(t Type, nChildren *uint) *Type
	typeClassPeek                func(t Type) T.Gpointer
	typeClassPeekStatic          func(t Type) T.Gpointer
	typeClassRef                 func(t Type) T.Gpointer
	typeDefaultInterfacePeek     func(gType Type) T.Gpointer
	typeDefaultInterfaceRef      func(gType Type) T.Gpointer
	typeDepth                    func(t Type) uint
	typeFundamental              func(typeId Type) Type
	typeFundamentalNext          func() Type
	typeGetPlugin                func(t Type) *TypePlugin
	typeGetQdata                 func(t Type, quark T.GQuark) T.Gpointer
	typeInterfaceAddPrerequisite func(interfaceType Type, prerequisiteType Type)
	typeInterfaceGetPlugin       func(instanceType Type, interfaceType Type) *TypePlugin
	typeInterfacePrerequisites   func(interfaceType Type, nPrerequisites *uint) *Type
	typeInterfaces               func(t Type, nInterfaces *uint) *Type
	typeIsA                      func(t, isAType Type) T.Gboolean
	typeName                     func(t Type) string
	typeNextBase                 func(leafType, rootType Type) Type
	typeParent                   func(t Type) Type
	typeQname                    func(t Type) T.GQuark
	typeQuery                    func(t Type, query *T.GTypeQuery)
	typeRegisterDynamic          func(parentType Type, typeName string, plugin *TypePlugin, flags TypeFlags) Type
	typeRegisterFundamental      func(typeId Type, typeName string, info *TypeInfo, finfo *T.GTypeFundamentalInfo, flags TypeFlags) Type
	typeRegisterStatic           func(parentType Type, typeName string, info *TypeInfo, flags TypeFlags) Type
	typeRegisterStaticSimple     func(parentType Type, typeName string, classSize uint, classInit T.GClassInitFunc, instanceSize uint, instanceInit T.GInstanceInitFunc, flags TypeFlags) Type
	typeSetQdata                 func(t Type, quark T.GQuark, data T.Gpointer)
	typeTestFlags                func(t Type, flags uint) T.Gboolean
	typeValueTablePeek           func(t Type) *T.GTypeValueTable
)

func (t Type) AddClassPrivate(privateSize T.Gsize) { typeAddClassPrivate(t, privateSize) }
func (t Type) AddInterfaceDynamic(interfaceType Type, plugin *TypePlugin) {
	typeAddInterfaceDynamic(t, interfaceType, plugin)
}
func (t Type) AddInterfaceStatic(interfaceType Type, info *T.GInterfaceInfo) {
	typeAddInterfaceStatic(t, interfaceType, info)
}
func (t Type) CheckIsValueType() T.Gboolean       { return typeCheckIsValueType(t) }
func (t Type) Children(nChildren *uint) *Type     { return typeChildren(t, nChildren) }
func (t Type) ClassPeek() T.Gpointer              { return typeClassPeek(t) }
func (t Type) ClassPeekStatic() T.Gpointer        { return typeClassPeekStatic(t) }
func (t Type) ClassRef() T.Gpointer               { return typeClassRef(t) }
func (t Type) DefaultInterfacePeek() T.Gpointer   { return typeDefaultInterfacePeek(t) }
func (t Type) DefaultInterfaceRef() T.Gpointer    { return typeDefaultInterfaceRef(t) }
func (t Type) Depth() uint                        { return typeDepth(t) }
func (t Type) Fundamental() Type                  { return typeFundamental(t) }
func (t Type) GetPlugin() *TypePlugin             { return typeGetPlugin(t) }
func (t Type) GetQdata(quark T.GQuark) T.Gpointer { return typeGetQdata(t, quark) }
func (t Type) InterfaceAddPrerequisite(prerequisiteType Type) {
	typeInterfaceAddPrerequisite(t, prerequisiteType)
}
func (t Type) InterfaceGetPlugin(interfaceType Type) *TypePlugin {
	return typeInterfaceGetPlugin(t, interfaceType)
}
func (t Type) InterfacePrerequisites(nPrerequisites *uint) *Type {
	return typeInterfacePrerequisites(t, nPrerequisites)
}
func (t Type) Interfaces(nInterfaces *uint) *Type { return typeInterfaces(t, nInterfaces) }
func (t Type) IsA(isAType Type) T.Gboolean        { return typeIsA(t, isAType) }
func (t Type) Name() string                       { return typeName(t) }
func (t Type) NextBase(rootType Type) Type        { return typeNextBase(t, rootType) }
func (t Type) Parent() Type                       { return typeParent(t) }
func (t Type) Qname() T.GQuark                    { return typeQname(t) }
func (t Type) Query(query *T.GTypeQuery)          { typeQuery(t, query) }
func (t Type) RegisterDynamic(typeName string, plugin *TypePlugin, flags TypeFlags) Type {
	return typeRegisterDynamic(t, typeName, plugin, flags)
}
func (t Type) RegisterFundamental(typeName string, info *TypeInfo, finfo *T.GTypeFundamentalInfo, flags TypeFlags) Type {
	return typeRegisterFundamental(t, typeName, info, finfo, flags)
}
func (t Type) RegisterStatic(typeName string, info *TypeInfo, flags TypeFlags) Type {
	return typeRegisterStatic(t, typeName, info, flags)
}
func (t Type) RegisterStaticSimple(typeName string, classSize uint, classInit T.GClassInitFunc, instanceSize uint, instanceInit T.GInstanceInitFunc, flags TypeFlags) Type {
	return typeRegisterStaticSimple(t, typeName, classSize, classInit, instanceSize, instanceInit, flags)
}
func (t Type) SetQdata(quark T.GQuark, data T.Gpointer) { typeSetQdata(t, quark, data) }
func (t Type) TestFlags(flags uint) T.Gboolean          { return typeTestFlags(t, flags) }
func (t Type) ValueTablePeek() *T.GTypeValueTable       { return typeValueTablePeek(t) }

type TypeClass struct {
	Type Type
}

var (
	typeCheckClassCast  func(gClass *TypeClass, isAType Type) *TypeClass
	typeCheckClassIsA   func(gClass *TypeClass, isAType Type) T.Gboolean
	typeClassGetPrivate func(klass *TypeClass, privateType Type) T.Gpointer
	typeNameFromClass   func(gClass *TypeClass) string
)

func (t *TypeClass) CheckCast(isAType Type) *TypeClass { return typeCheckClassCast(t, isAType) }
func (t *TypeClass) GetPrivate(privateType Type) T.Gpointer {
	return typeClassGetPrivate(t, privateType)
}
func (t *TypeClass) IsA(isAType Type) T.Gboolean { return typeCheckClassIsA(t, isAType) }
func (t *TypeClass) Name() string                { return typeNameFromClass(t) }

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

	typeCheckInstance      func(i *TypeInstance) T.Gboolean
	typeCheckInstanceCast  func(i *TypeInstance, ifaceType Type) *TypeInstance
	typeCheckInstanceIsA   func(i *TypeInstance, ifaceType Type) T.Gboolean
	typeFreeInstance       func(i *TypeInstance)
	typeInstanceGetPrivate func(i *TypeInstance, privateType Type) T.Gpointer
	typeNameFromInstance   func(i *TypeInstance) string
)

func (i *TypeInstance) CheckInstance() T.Gboolean { return typeCheckInstance(i) }
func (i *TypeInstance) CheckCast(ifaceType Type) *TypeInstance {
	return typeCheckInstanceCast(i, ifaceType)
}
func (i *TypeInstance) IsA(ifaceType Type) T.Gboolean {
	return typeCheckInstanceIsA(i, ifaceType)
}
func (i *TypeInstance) Free() { typeFreeInstance(i) }
func (i *TypeInstance) GetPrivate(privateType Type) T.Gpointer {
	return typeInstanceGetPrivate(i, privateType)
}
func (i *TypeInstance) NameFromInstance() string { return typeNameFromInstance(i) }

type TypeModule struct {
	Parent         Object
	UseCount       uint
	TypeInfos      *T.GSList
	InterfaceInfos *T.GSList
	Name           *T.Gchar
}

var (
	TypeModuleGetType func() Type

	typeModuleAddInterface  func(m *TypeModule, instanceType, interfaceType Type, interfaceInfo *T.GInterfaceInfo)
	typeModuleRegisterEnum  func(m *TypeModule, name string, constStaticValues *EnumValue) Type
	typeModuleRegisterFlags func(m *TypeModule, name string, constStaticValues *FlagsValue) Type
	typeModuleRegisterType  func(m *TypeModule, parentType Type, typeName string, typeInfo *TypeInfo, flags TypeFlags) Type
	typeModuleSetName       func(m *TypeModule, name string)
	typeModuleUnuse         func(m *TypeModule)
	typeModuleUse           func(m *TypeModule) T.Gboolean
)

func (m *TypeModule) AddInterface(instanceType, interfaceType Type, interfaceInfo *T.GInterfaceInfo) {
	typeModuleAddInterface(m, instanceType, interfaceType, interfaceInfo)
}
func (m *TypeModule) RegisterEnum(name string, constStaticValues *EnumValue) Type {
	return typeModuleRegisterEnum(m, name, constStaticValues)
}
func (m *TypeModule) RegisterFlags(name string, constStaticValues *FlagsValue) Type {
	return typeModuleRegisterFlags(m, name, constStaticValues)
}
func (m *TypeModule) RegisterType(parentType Type, typeName string, typeInfo *TypeInfo, flags TypeFlags) Type {
	return typeModuleRegisterType(m, parentType, typeName, typeInfo, flags)
}
func (m *TypeModule) SetName(name string) { typeModuleSetName(m, name) }
func (m *TypeModule) Unuse()              { typeModuleUnuse(m) }
func (m *TypeModule) Use() T.Gboolean     { return typeModuleUse(m) }

type TypePlugin struct{}

var (
	TypePluginGetType func() Type

	typePluginCompleteInterfaceInfo func(p *TypePlugin, instanceType, interfaceType Type, info *T.GInterfaceInfo)
	typePluginCompleteTypeInfo      func(p *TypePlugin, gType Type, info *TypeInfo, valueTable *T.GTypeValueTable)
	typePluginUnuse                 func(p *TypePlugin)
	typePluginUse                   func(p *TypePlugin)
)

func (p *TypePlugin) CompleteInterfaceInfo(instanceType, interfaceType Type, info *T.GInterfaceInfo) {
	typePluginCompleteInterfaceInfo(p, instanceType, interfaceType, info)
}
func (p *TypePlugin) CompleteTypeInfo(gType Type, info *TypeInfo, valueTable *T.GTypeValueTable) {
	typePluginCompleteTypeInfo(p, gType, info, valueTable)
}
func (p *TypePlugin) Unuse() { typePluginUnuse(p) }
func (p *TypePlugin) Use()   { typePluginUse(p) }
