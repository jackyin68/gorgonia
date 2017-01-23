package main

import (
	"reflect"
	"strings"
	"text/template"
)

var parameterizedKinds = [...]reflect.Kind{
	reflect.Array,
	reflect.Chan,
	reflect.Func,
	reflect.Interface,
	reflect.Map,
	reflect.Ptr,
	reflect.Slice,
	reflect.Struct,
}

var rangeable = [...]reflect.Kind{
	reflect.Int,
	reflect.Int8,
	reflect.Int16,
	reflect.Int32,
	reflect.Int64,
	reflect.Uint,
	reflect.Uint8,
	reflect.Uint16,
	reflect.Uint32,
	reflect.Uint64,
	reflect.Float32,
	reflect.Float64,
	reflect.Complex64,
	reflect.Complex128,
}

var specialized = [...]reflect.Kind{
	reflect.Bool,
	reflect.Int,
	reflect.Int8,
	reflect.Int16,
	reflect.Int32,
	reflect.Int64,
	reflect.Uint,
	reflect.Uint8,
	reflect.Uint16,
	reflect.Uint32,
	reflect.Uint64,
	reflect.Float32,
	reflect.Float64,
	reflect.Complex64,
	reflect.Complex128,
	reflect.String,
}

var funcs = template.FuncMap{
	"lower":           strings.ToLower,
	"title":           strings.Title,
	"hasPrefix":       strings.HasPrefix,
	"isParameterized": isParameterized,
	"isRangeable":     isRangeable,
	"isSpecialized":   isSpecialized,

	"short": short,
	"clean": clean,
	"strip": strip,

	"reflectKind": reflectKind,
	"asType":      asType,
}

func isParameterized(a reflect.Kind) bool {
	for _, v := range parameterizedKinds {
		if v == a {
			return true
		}
	}
	return false
}

func isRangeable(a reflect.Kind) bool {
	for _, v := range rangeable {
		if v == a {
			return true
		}
	}
	return false
}

func isSpecialized(a reflect.Kind) bool {
	for _, v := range specialized {
		if v == a {
			return true
		}
	}
	return false
}

var shortNames = map[reflect.Kind]string{
	reflect.Invalid:       "Invalid",
	reflect.Bool:          "B",
	reflect.Int:           "I",
	reflect.Int8:          "I8",
	reflect.Int16:         "I16",
	reflect.Int32:         "I32",
	reflect.Int64:         "I64",
	reflect.Uint:          "U",
	reflect.Uint8:         "U8",
	reflect.Uint16:        "U16",
	reflect.Uint32:        "U32",
	reflect.Uint64:        "U64",
	reflect.Uintptr:       "Uintptr",
	reflect.Float32:       "F32",
	reflect.Float64:       "F64",
	reflect.Complex64:     "C64",
	reflect.Complex128:    "C128",
	reflect.Array:         "Array",
	reflect.Chan:          "Chan",
	reflect.Func:          "Func",
	reflect.Interface:     "Interface",
	reflect.Map:           "Map",
	reflect.Ptr:           "Ptr",
	reflect.Slice:         "Slice",
	reflect.String:        "Str",
	reflect.Struct:        "Struct",
	reflect.UnsafePointer: "UnsafePointer",
}

func short(a reflect.Kind) string {
	return shortNames[a]
}

func clean(a string) string {
	if a == "unsafe.pointer" {
		return "unsafe.Pointer"
	}
	return a
}

func strip(a string) string {
	return strings.Replace(a, ".", "", -1)
}

func reflectKind(a reflect.Kind) string {
	return strip(strings.Title(a.String()))
}

func asType(a reflect.Kind) string {
	return clean(a.String())
}
