// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package glib

import (
	T "github.com/tHinqa/outside-gtk2/types"
	. "github.com/tHinqa/outside/types"
)

type Rand struct{}

var (
	RandNew              func() *Rand
	RandNewWithSeed      func(seed T.GUint32) *Rand
	RandNewWithSeedArray func(seed *T.GUint32, seedLength uint) *Rand

	RandomDouble      func() float64
	RandomDoubleRange func(begin, end float64) float64
	RandomInt         func() T.GUint32
	RandomIntRange    func(begin, end T.GInt32) T.GInt32
	RandomSetSeed     func(seed T.GUint32)

	RandFree         func(r *Rand)
	RandCopy         func(r *Rand) *Rand
	RandSetSeed      func(r *Rand, seed T.GUint32)
	RandSetSeedArray func(r *Rand, seed *T.GUint32, seedLength uint)
	RandInt          func(r *Rand) T.GUint32
	RandIntRange     func(r *Rand, begin, end T.GInt32) T.GInt32
	RandDouble       func(r *Rand) float64
	RandDoubleRange  func(r *Rand, begin, end float64) float64
)

func (r *Rand) Free()                                         { RandFree(r) }
func (r *Rand) Copy() *Rand                                   { return RandCopy(r) }
func (r *Rand) SetSeed(seed T.GUint32)                        { RandSetSeed(r, seed) }
func (r *Rand) SetSeedArray(seed *T.GUint32, seedLength uint) { RandSetSeedArray(r, seed, seedLength) }
func (r *Rand) Int() T.GUint32                                { return RandInt(r) }
func (r *Rand) IntRange(begin, end T.GInt32) T.GInt32         { return RandIntRange(r, begin, end) }
func (r *Rand) Double() float64                               { return RandDouble(r) }
func (r *Rand) DoubleRange(begin, end float64) float64        { return RandDoubleRange(r, begin, end) }

type Regex struct{}

var (
	RegexNew func(pattern string, compileOptions RegexCompileFlags, matchOptions RegexMatchFlags, e **Error) *Regex

	RegexCheckReplacement func(replacement string, hasReferences *bool, e **Error) bool
	RegexEscapeString     func(str string, length int) string
	RegexMatchSimple      func(pattern, str string, compileOptions RegexCompileFlags, matchOptions RegexMatchFlags) bool
	RegexSplitSimple      func(pattern string, str string, compileOptions RegexCompileFlags, matchOptions RegexMatchFlags) []string

	RegexGetCaptureCount func(r *Regex) int
	RegexGetCompileFlags func(r *Regex) RegexCompileFlags
	RegexGetMatchFlags   func(r *Regex) RegexMatchFlags
	RegexGetMaxBackref   func(r *Regex) int
	RegexGetPattern      func(r *Regex) string
	RegexGetStringNumber func(r *Regex, name string) int
	RegexMatch           func(r *Regex, str string, matchOptions RegexMatchFlags, matchInfo **T.GMatchInfo) bool
	RegexMatchAll        func(r *Regex, str string, matchOptions RegexMatchFlags, matchInfo **T.GMatchInfo) bool
	RegexMatchAllFull    func(r *Regex, str string, stringLen T.Gssize, startPosition int, matchOptions RegexMatchFlags, matchInfo **T.GMatchInfo, e **Error) bool
	RegexMatchFull       func(r *Regex, str string, stringLen T.Gssize, startPosition int, matchOptions RegexMatchFlags, matchInfo **T.GMatchInfo, e **Error) bool
	RegexRef             func(r *Regex) *Regex
	RegexReplace         func(r *Regex, str string, stringLen T.Gssize, startPosition int, replacement string, matchOptions RegexMatchFlags, e **Error) string
	RegexReplaceEval     func(r *Regex, str string, stringLen T.Gssize, startPosition int, matchOptions RegexMatchFlags, eval RegexEvalCallback, userData T.Gpointer, e **Error) string
	RegexReplaceLiteral  func(r *Regex, str string, stringLen T.Gssize, startPosition int, replacement string, matchOptions RegexMatchFlags, e **Error) string
	RegexSplit           func(r *Regex, str string, matchOptions RegexMatchFlags) []string
	RegexSplitFull       func(r *Regex, str string, stringLen T.Gssize, startPosition int, matchOptions RegexMatchFlags, maxTokens int, e **Error) []string
	RegexUnref           func(r *Regex)
)

func (r *Regex) GetCaptureCount() int               { return RegexGetCaptureCount(r) }
func (r *Regex) GetCompileFlags() RegexCompileFlags { return RegexGetCompileFlags(r) }
func (r *Regex) GetMatchFlags() RegexMatchFlags     { return RegexGetMatchFlags(r) }
func (r *Regex) GetMaxBackref() int                 { return RegexGetMaxBackref(r) }
func (r *Regex) GetPattern() string                 { return RegexGetPattern(r) }
func (r *Regex) GetStringNumber(name string) int    { return RegexGetStringNumber(r, name) }
func (r *Regex) Match(str string, matchOptions RegexMatchFlags, matchInfo **T.GMatchInfo) bool {
	return RegexMatch(r, str, matchOptions, matchInfo)
}
func (r *Regex) MatchAll(str string, matchOptions RegexMatchFlags, matchInfo **T.GMatchInfo) bool {
	return RegexMatchAll(r, str, matchOptions, matchInfo)
}
func (r *Regex) MatchAllFull(str string, stringLen T.Gssize, startPosition int, matchOptions RegexMatchFlags, matchInfo **T.GMatchInfo, e **Error) bool {
	return RegexMatchAllFull(r, str, stringLen, startPosition, matchOptions, matchInfo, e)
}
func (r *Regex) MatchFull(str string, stringLen T.Gssize, startPosition int, matchOptions RegexMatchFlags, matchInfo **T.GMatchInfo, e **Error) bool {
	return RegexMatchFull(r, str, stringLen, startPosition, matchOptions, matchInfo, e)
}
func (r *Regex) Ref() *Regex { return RegexRef(r) }
func (r *Regex) Replace(str string, stringLen T.Gssize, startPosition int, replacement string, matchOptions RegexMatchFlags, e **Error) string {
	return RegexReplace(r, str, stringLen, startPosition, replacement, matchOptions, e)
}
func (r *Regex) ReplaceEval(str string, stringLen T.Gssize, startPosition int, matchOptions RegexMatchFlags, eval RegexEvalCallback, userData T.Gpointer, e **Error) string {
	return RegexReplaceEval(r, str, stringLen, startPosition, matchOptions, eval, userData, e)
}
func (r *Regex) ReplaceLiteral(str string, stringLen T.Gssize, startPosition int, replacement string, matchOptions RegexMatchFlags, e **Error) string {
	return RegexReplaceLiteral(r, str, stringLen, startPosition, replacement, matchOptions, e)
}
func (r *Regex) Split(str string, matchOptions RegexMatchFlags) []string {
	return RegexSplit(r, str, matchOptions)
}
func (r *Regex) SplitFull(str string, stringLen T.Gssize, startPosition int, matchOptions RegexMatchFlags, maxTokens int, e **Error) []string {
	return RegexSplitFull(r, str, stringLen, startPosition, matchOptions, maxTokens, e)
}
func (r *Regex) Unref() { RegexUnref(r) }

type RegexCompileFlags Enum

const (
	REGEX_CASELESS RegexCompileFlags = 1 << iota
	REGEX_MULTILINE
	REGEX_DOTALL
	REGEX_EXTENDED
	REGEX_ANCHORED
	REGEX_DOLLAR_ENDONLY
	_
	_
	_
	REGEX_UNGREEDY
	_
	REGEX_RAW
	REGEX_NO_AUTO_CAPTURE
	REGEX_OPTIMIZE
	_
	_
	_
	_
	_

	REGEX_DUPNAMES
	REGEX_NEWLINE_CR
	REGEX_NEWLINE_LF
	REGEX_NEWLINE_CRLF = REGEX_NEWLINE_CR | REGEX_NEWLINE_LF
)

type RegexEvalCallback func(matchInfo *T.GMatchInfo,
	result *String, userData T.Gpointer) bool

type RegexMatchFlags Enum

const (
	REGEX_MATCH_ANCHORED RegexMatchFlags = 1 << (iota + 4)
	_
	_
	REGEX_MATCH_NOTBOL
	REGEX_MATCH_NOTEOL
	_
	REGEX_MATCH_NOTEMPTY
	_
	_
	_
	_
	REGEX_MATCH_PARTIAL
	_
	_
	_
	_
	REGEX_MATCH_NEWLINE_CR
	REGEX_MATCH_NEWLINE_LF
	REGEX_MATCH_NEWLINE_ANY
	REGEX_MATCH_NEWLINE_CRLF = REGEX_MATCH_NEWLINE_CR | REGEX_MATCH_NEWLINE_LF
)

type Relation struct{}

var (
	RelationNew func(fields int) *Relation

	RelationCount   func(r *Relation, key T.Gconstpointer, field int) int
	RelationDelete  func(r *Relation, key T.Gconstpointer, field int) int
	RelationDestroy func(r *Relation)
	RelationExists  func(r *Relation, v ...VArg) bool
	RelationIndex   func(r *Relation, field int, hashFunc HashFunc, keyEqualFunc T.GEqualFunc)
	RelationInsert  func(r *Relation, v ...VArg)
	RelationPrint   func(r *Relation)
	RelationSelect  func(r *Relation, key T.Gconstpointer, field int) *Tuples
)

func (r *Relation) Count(key T.Gconstpointer, field int) int  { return RelationCount(r, key, field) }
func (r *Relation) Delete(key T.Gconstpointer, field int) int { return RelationDelete(r, key, field) }
func (r *Relation) Destroy()                                  { RelationDestroy(r) }
func (r *Relation) Exists(v ...VArg) bool                     { return RelationExists(r, v) }
func (r *Relation) Index(field int, hashFunc HashFunc, keyEqualFunc T.GEqualFunc) {
	RelationIndex(r, field, hashFunc, keyEqualFunc)
}
func (r *Relation) Insert(v ...VArg) { RelationInsert(r, v) }
func (r *Relation) Print()           { RelationPrint(r) }
func (r *Relation) Select(key T.Gconstpointer, field int) *Tuples {
	return RelationSelect(r, key, field)
}
