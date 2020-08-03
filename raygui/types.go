// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY ROBOTS.

// WARNING: This file has automatically been generated on Mon, 03 Aug 2020 11:53:51 CST.
// Code generated by https://git.io/c-for-go. DO NOT EDIT.

package raygui

/*
#include "../lib/raygui/src/raygui.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"

// GuiStyleProp as declared in src/raygui.h:256
type GuiStyleProp struct {
	ControlId      uint16
	PropertyId     uint16
	PropertyValue  int32
	ref8760bcee    *C.GuiStyleProp
	allocs8760bcee interface{}
}
type guiStyleProp struct {
	ControlId     uint16
	PropertyId    uint16
	PropertyValue int32
}

// Vector2 as declared in src/raylib.h:176
type Vector2 struct {
	X              float32
	Y              float32
	ref29ca61a5    *C.Vector2
	allocs29ca61a5 interface{}
}
type vector2 struct {
	X float32
	Y float32
}

// Vector3 as declared in src/raylib.h:183
type Vector3 struct {
	X              float32
	Y              float32
	Z              float32
	ref5ecd5133    *C.Vector3
	allocs5ecd5133 interface{}
}
type vector3 struct {
	X float32
	Y float32
	Z float32
}

// Color as declared in src/raylib.h:210
type Color struct {
	R              byte
	G              byte
	B              byte
	A              byte
	refa79767ed    *C.Color
	allocsa79767ed interface{}
}
type color struct {
	R byte
	G byte
	B byte
	A byte
}

// Rectangle as declared in src/raylib.h:218
type Rectangle struct {
	X              float32
	Y              float32
	Width          float32
	Height         float32
	refcee8783a    *C.Rectangle
	allocscee8783a interface{}
}
type rectangle struct {
	X      float32
	Y      float32
	Width  float32
	Height float32
}

// Texture2D as declared in src/raylib.h:238
type Texture2D struct {
	Id             uint32
	Width          int32
	Height         int32
	Mipmaps        int32
	Format         int32
	ref3c51a40b    *C.Texture2D
	allocs3c51a40b interface{}
}
type texture2D struct {
	Id      uint32
	Width   int32
	Height  int32
	Mipmaps int32
	Format  int32
}

// Font as declared in src/raylib.h:283
type Font struct {
	BaseSize       int32
	CharsCount     int32
	Texture        Texture2D
	Recs           []Rectangle
	ref70a6a7ec    *C.Font
	allocs70a6a7ec interface{}
}
type font struct {
	BaseSize   int32
	CharsCount int32
	Texture    C.Texture2D
	Recs       *C.Rectangle
}
