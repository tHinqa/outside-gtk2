// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package glib

import (
	// O "github.com/tHinqa/outside-gtk2/gobject"
	T "github.com/tHinqa/outside-gtk2/types"
	// . "github.com/tHinqa/outside/types"
)

type Date struct {
	Bits1, Bits2 uint
	// JulianDays : 32
	// Julian : 1
	// Dmy : 1
	// Day : 6
	// Month : 4
	// Year : 16
}

var (
	DateNew       func() *Date
	DateNewDmy    func(day DateDay, month DateMonth, year DateYear) *Date
	DateNewJulian func(julianDay T.GUint32) *Date

	DateGetDaysInMonth       func(month DateMonth, year DateYear) uint8
	DateGetMondayWeeksInYear func(year DateYear) uint8
	DateGetSundayWeeksInYear func(year DateYear) uint8
	DateIsLeapYear           func(year DateYear) bool
	DateStrftime             func(s string, slen T.Gsize, format string, date *Date) T.Gsize
	DateValidDay             func(day DateDay) bool
	DateValidDmy             func(day DateDay, month DateMonth, year DateYear) bool
	DateValidJulian          func(julianDate T.GUint32) bool
	DateValidMonth           func(month DateMonth) bool
	DateValidWeekday         func(weekday DateWeekday) bool
	DateValidYear            func(year DateYear) bool

	DateFree                 func(d *Date)
	DateValid                func(d *Date) bool
	DateGetWeekday           func(d *Date) DateWeekday
	DateGetMonth             func(d *Date) DateMonth
	DateGetYear              func(d *Date) DateYear
	DateGetDay               func(d *Date) DateDay
	DateGetJulian            func(d *Date) T.GUint32
	DateGetDayOfYear         func(d *Date) uint
	DateGetMondayWeekOfYear  func(d *Date) uint
	DateGetSundayWeekOfYear  func(d *Date) uint
	DateGetIso8601WeekOfYear func(d *Date) uint
	DateClear                func(d *Date, nDates uint)
	DateSetParse             func(d *Date, str string)
	DateSetTimeT             func(d *Date, timet T.TimeT)
	DateSetTimeVal           func(d *Date, timeval *TimeVal)
	DateSetTime              func(d *Date, time T.GTime)
	DateSetMonth             func(d *Date, month DateMonth)
	DateSetDay               func(d *Date, day DateDay)
	DateSetYear              func(d *Date, year DateYear)
	DateSetDmy               func(d *Date, day DateDay, month DateMonth, y DateYear)
	DateSetJulian            func(d *Date, julianDate T.GUint32)
	DateIsFirstOfMonth       func(d *Date) bool
	DateIsLastOfMonth        func(d *Date) bool
	DateAddDays              func(d *Date, nDays uint)
	DateSubtractDays         func(d *Date, nDays uint)
	DateAddMonths            func(d *Date, nMonths uint)
	DateSubtractMonths       func(d *Date, nMonths uint)
	DateAddYears             func(d *Date, nYears uint)
	DateSubtractYears        func(d *Date, nYears uint)
	DateDaysBetween          func(d *Date, d2 *Date) int
	DateCompare              func(d *Date, d2 *Date) int
	DateToStructTm           func(d *Date, tm *T.Tm)
	DateClamp                func(d *Date, minDate, maxDate *Date)
	DateOrder                func(d *Date, d2 *Date)
)

func (d *Date) Free()                                           { DateFree(d) }
func (d *Date) Valid() bool                                     { return DateValid(d) }
func (d *Date) GetWeekday() DateWeekday                         { return DateGetWeekday(d) }
func (d *Date) GetMonth() DateMonth                             { return DateGetMonth(d) }
func (d *Date) GetYear() DateYear                               { return DateGetYear(d) }
func (d *Date) GetDay() DateDay                                 { return DateGetDay(d) }
func (d *Date) GetJulian() T.GUint32                            { return DateGetJulian(d) }
func (d *Date) GetDayOfYear() uint                              { return DateGetDayOfYear(d) }
func (d *Date) GetMondayWeekOfYear() uint                       { return DateGetMondayWeekOfYear(d) }
func (d *Date) GetSundayWeekOfYear() uint                       { return DateGetSundayWeekOfYear(d) }
func (d *Date) GetIso8601WeekOfYear() uint                      { return DateGetIso8601WeekOfYear(d) }
func (d *Date) Clear(nDates uint)                               { DateClear(d, nDates) }
func (d *Date) SetParse(str string)                             { DateSetParse(d, str) }
func (d *Date) SetTimeT(timet T.TimeT)                          { DateSetTimeT(d, timet) }
func (d *Date) SetTimeVal(timeval *TimeVal)                     { DateSetTimeVal(d, timeval) }
func (d *Date) SetTime(time T.GTime)                            { DateSetTime(d, time) }
func (d *Date) SetMonth(month DateMonth)                        { DateSetMonth(d, month) }
func (d *Date) SetDay(day DateDay)                              { DateSetDay(d, day) }
func (d *Date) SetYear(year DateYear)                           { DateSetYear(d, year) }
func (d *Date) SetDmy(day DateDay, month DateMonth, y DateYear) { DateSetDmy(d, day, month, y) }
func (d *Date) SetJulian(julianDate T.GUint32)                  { DateSetJulian(d, julianDate) }
func (d *Date) IsFirstOfMonth() bool                            { return DateIsFirstOfMonth(d) }
func (d *Date) IsLastOfMonth() bool                             { return DateIsLastOfMonth(d) }
func (d *Date) AddDays(nDays uint)                              { DateAddDays(d, nDays) }
func (d *Date) SubtractDays(nDays uint)                         { DateSubtractDays(d, nDays) }
func (d *Date) AddMonths(nMonths uint)                          { DateAddMonths(d, nMonths) }
func (d *Date) SubtractMonths(nMonths uint)                     { DateSubtractMonths(d, nMonths) }
func (d *Date) AddYears(nYears uint)                            { DateAddYears(d, nYears) }
func (d *Date) SubtractYears(nYears uint)                       { DateSubtractYears(d, nYears) }
func (d *Date) DaysBetween(d2 *Date) int                        { return DateDaysBetween(d, d2) }
func (d *Date) Compare(d2 *Date) int                            { return DateCompare(d, d2) }
func (d *Date) ToStructTm(tm *T.Tm)                             { DateToStructTm(d, tm) }
func (d *Date) Clamp(minDate, maxDate *Date)                    { DateClamp(d, minDate, maxDate) }
func (d *Date) Order(d2 *Date)                                  { DateOrder(d, d2) }

type DateDay uint8

type DateDMY Enum

const (
	DATE_DAY DateDMY = iota
	DATE_MONTH
	DATE_YEAR
)

type DateMonth Enum

const (
	DATE_BAD_MONTH DateMonth = iota
	DATE_JANUARY
	DATE_FEBRUARY
	DATE_MARCH
	DATE_APRIL
	DATE_MAY
	DATE_JUNE
	DATE_JULY
	DATE_AUGUST
	DATE_SEPTEMBER
	DATE_OCTOBER
	DATE_NOVEMBER
	DATE_DECEMBER
)

type DateTime T.GInt32

var (
	DateTimeNew                 func(tz *TimeZone, year int, month int, day int, hour int, minute int, seconds float64) *DateTime
	DateTimeNewFromTimevalLocal func(tv *TimeVal) *DateTime
	DateTimeNewFromTimevalUtc   func(tv *TimeVal) *DateTime
	DateTimeNewFromUnixLocal    func(t int64) *DateTime
	DateTimeNewFromUnixUtc      func(t int64) *DateTime
	DateTimeNewLocal            func(year int, month int, day int, hour int, minute int, seconds float64) *DateTime
	DateTimeNewNow              func(tz *TimeZone) *DateTime
	DateTimeNewNowLocal         func() *DateTime
	DateTimeNewNowUtc           func() *DateTime
	DateTimeNewUtc              func(year int, month int, day int, hour int, minute int, seconds float64) *DateTime

	DateTimeCompare func(dt1 T.Gconstpointer, dt2 T.Gconstpointer) int
	DateTimeHash    func(datetime T.Gconstpointer) uint
	DateTimeEqual   func(dt1 T.Gconstpointer, dt2 T.Gconstpointer) bool

	DateTimeAdd                     func(d *DateTime, timespan T.GTimeSpan) *DateTime
	DateTimeAddDays                 func(d *DateTime, days int) *DateTime
	DateTimeAddFull                 func(d *DateTime, years, months, days, hours, minutes int, seconds float64) *DateTime
	DateTimeAddHours                func(d *DateTime, hours int) *DateTime
	DateTimeAddMinutes              func(d *DateTime, minutes int) *DateTime
	DateTimeAddMonths               func(d *DateTime, months int) *DateTime
	DateTimeAddSeconds              func(d *DateTime, seconds float64) *DateTime
	DateTimeAddWeeks                func(d *DateTime, weeks int) *DateTime
	DateTimeAddYears                func(d *DateTime, years int) *DateTime
	DateTimeDifference              func(end, begin *DateTime) T.GTimeSpan
	DateTimeFormat                  func(d *DateTime, format string) string
	DateTimeGetDayOfMonth           func(d *DateTime) int
	DateTimeGetDayOfWeek            func(d *DateTime) int
	DateTimeGetDayOfYear            func(d *DateTime) int
	DateTimeGetHour                 func(d *DateTime) int
	DateTimeGetMicrosecond          func(d *DateTime) int
	DateTimeGetMinute               func(d *DateTime) int
	DateTimeGetMonth                func(d *DateTime) int
	DateTimeGetSecond               func(d *DateTime) int
	DateTimeGetSeconds              func(d *DateTime) float64
	DateTimeGetTimezoneAbbreviation func(d *DateTime) string
	DateTimeGetUtcOffset            func(d *DateTime) T.GTimeSpan
	DateTimeGetWeekNumberingYear    func(d *DateTime) int
	DateTimeGetWeekOfYear           func(d *DateTime) int
	DateTimeGetYear                 func(d *DateTime) int
	DateTimeGetYmd                  func(d *DateTime, year, month, day *int)
	DateTimeIsDaylightSavings       func(d *DateTime) bool
	DateTimeRef                     func(d *DateTime) *DateTime
	DateTimeToLocal                 func(d *DateTime) *DateTime
	DateTimeToTimeval               func(d *DateTime, tv *TimeVal) bool
	DateTimeToTimezone              func(d *DateTime, tz *TimeZone) *DateTime
	DateTimeToUnix                  func(d *DateTime) int64
	DateTimeToUtc                   func(d *DateTime) *DateTime
	DateTimeUnref                   func(d *DateTime)
)

func (d *DateTime) Add(timespan T.GTimeSpan) *DateTime { return DateTimeAdd(d, timespan) }
func (d *DateTime) AddDays(days int) *DateTime         { return DateTimeAddDays(d, days) }
func (d *DateTime) AddFull(years, months, days, hours, minutes int, seconds float64) *DateTime {
	return DateTimeAddFull(d, years, months, days, hours, minutes, seconds)
}
func (d *DateTime) AddHours(hours int) *DateTime           { return DateTimeAddHours(d, hours) }
func (d *DateTime) AddMinutes(minutes int) *DateTime       { return DateTimeAddMinutes(d, minutes) }
func (d *DateTime) AddMonths(months int) *DateTime         { return DateTimeAddMonths(d, months) }
func (d *DateTime) AddSeconds(seconds float64) *DateTime   { return DateTimeAddSeconds(d, seconds) }
func (d *DateTime) AddWeeks(weeks int) *DateTime           { return DateTimeAddWeeks(d, weeks) }
func (d *DateTime) AddYears(years int) *DateTime           { return DateTimeAddYears(d, years) }
func (d *DateTime) Format(format string) string            { return DateTimeFormat(d, format) }
func (d *DateTime) GetDayOfMonth() int                     { return DateTimeGetDayOfMonth(d) }
func (d *DateTime) GetDayOfWeek() int                      { return DateTimeGetDayOfWeek(d) }
func (d *DateTime) GetDayOfYear() int                      { return DateTimeGetDayOfYear(d) }
func (d *DateTime) GetHour() int                           { return DateTimeGetHour(d) }
func (d *DateTime) GetMicrosecond() int                    { return DateTimeGetMicrosecond(d) }
func (d *DateTime) GetMinute() int                         { return DateTimeGetMinute(d) }
func (d *DateTime) GetMonth() int                          { return DateTimeGetMonth(d) }
func (d *DateTime) GetSecond() int                         { return DateTimeGetSecond(d) }
func (d *DateTime) GetSeconds() float64                    { return DateTimeGetSeconds(d) }
func (d *DateTime) GetTimezoneAbbreviation() string        { return DateTimeGetTimezoneAbbreviation(d) }
func (d *DateTime) GetUtcOffset() T.GTimeSpan              { return DateTimeGetUtcOffset(d) }
func (d *DateTime) GetWeekNumberingYear() int              { return DateTimeGetWeekNumberingYear(d) }
func (d *DateTime) GetWeekOfYear() int                     { return DateTimeGetWeekOfYear(d) }
func (d *DateTime) GetYear() int                           { return DateTimeGetYear(d) }
func (d *DateTime) GetYmd(year, month, day *int)           { DateTimeGetYmd(d, year, month, day) }
func (d *DateTime) IsDaylightSavings() bool                { return DateTimeIsDaylightSavings(d) }
func (d *DateTime) Ref() *DateTime                         { return DateTimeRef(d) }
func (d *DateTime) ToLocal() *DateTime                     { return DateTimeToLocal(d) }
func (d *DateTime) ToTimeval(tv *TimeVal) bool             { return DateTimeToTimeval(d, tv) }
func (d *DateTime) ToTimezone(tz *TimeZone) *DateTime      { return DateTimeToTimezone(d, tz) }
func (d *DateTime) ToUnix() int64                          { return DateTimeToUnix(d) }
func (d *DateTime) ToUtc() *DateTime                       { return DateTimeToUtc(d) }
func (d *DateTime) Unref()                                 { DateTimeUnref(d) }
func (e *DateTime) Difference(begin *DateTime) T.GTimeSpan { return DateTimeDifference(e, begin) }

type DateWeekday Enum

const (
	DATE_BAD_WEEKDAY DateWeekday = iota
	DATE_MONDAY
	DATE_TUESDAY
	DATE_WEDNESDAY
	DATE_THURSDAY
	DATE_FRIDAY
	DATE_SATURDAY
	DATE_SUNDAY
)

type DateYear uint16

type Dir struct{}

var (
	DirOpen func(path string, flags uint, e **T.GError) *Dir

	DirClose    func(d *Dir)
	DirReadName func(d *Dir) string
	DirRewind   func(d *Dir)
)

func (d *Dir) Close()           { DirClose(d) }
func (d *Dir) ReadName() string { return DirReadName(d) }
func (d *Dir) Rewind()          { DirRewind(d) }
