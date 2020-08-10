// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY ROBOTS.

// WARNING: This file has automatically been generated on Mon, 10 Aug 2020 18:02:28 CST.
// Code generated by https://git.io/c-for-go. DO NOT EDIT.

package raygui

/*
#include "../lib/raygui/src/raygui.h"
#include "../lib/raygui/src/ricons.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"
import (
	"runtime"
	"unsafe"
)

// GuiEnable function as declared in src/raygui.h:425
func GuiEnable() {
	C.GuiEnable()
}

// GuiDisable function as declared in src/raygui.h:426
func GuiDisable() {
	C.GuiDisable()
}

// GuiLock function as declared in src/raygui.h:427
func GuiLock() {
	C.GuiLock()
}

// GuiUnlock function as declared in src/raygui.h:428
func GuiUnlock() {
	C.GuiUnlock()
}

// GuiFade function as declared in src/raygui.h:429
func GuiFade(alpha float32) {
	calpha, _ := (C.float)(alpha), cgoAllocsUnknown
	C.GuiFade(calpha)
}

// GuiSetState function as declared in src/raygui.h:430
func GuiSetState(state int32) {
	cstate, _ := (C.int)(state), cgoAllocsUnknown
	C.GuiSetState(cstate)
}

// GuiGetState function as declared in src/raygui.h:431
func GuiGetState() int32 {
	__ret := C.GuiGetState()
	__v := (int32)(__ret)
	return __v
}

// GuiSetFont function as declared in src/raygui.h:434
func GuiSetFont(font Font) {
	cfont, _ := font.PassValue()
	C.GuiSetFont(cfont)
}

// GuiGetFont function as declared in src/raygui.h:435
func GuiGetFont() Font {
	__ret := C.GuiGetFont()
	__v := *NewFontRef(unsafe.Pointer(&__ret))
	return __v
}

// GuiSetStyle function as declared in src/raygui.h:438
func GuiSetStyle(control int32, property int32, value int32) {
	ccontrol, _ := (C.int)(control), cgoAllocsUnknown
	cproperty, _ := (C.int)(property), cgoAllocsUnknown
	cvalue, _ := (C.int)(value), cgoAllocsUnknown
	C.GuiSetStyle(ccontrol, cproperty, cvalue)
}

// GuiGetStyle function as declared in src/raygui.h:439
func GuiGetStyle(control int32, property int32) int32 {
	ccontrol, _ := (C.int)(control), cgoAllocsUnknown
	cproperty, _ := (C.int)(property), cgoAllocsUnknown
	__ret := C.GuiGetStyle(ccontrol, cproperty)
	__v := (int32)(__ret)
	return __v
}

// GuiEnableTooltip function as declared in src/raygui.h:442
func GuiEnableTooltip() {
	C.GuiEnableTooltip()
}

// GuiDisableTooltip function as declared in src/raygui.h:443
func GuiDisableTooltip() {
	C.GuiDisableTooltip()
}

// GuiSetTooltip function as declared in src/raygui.h:444
func GuiSetTooltip(tooltip string) {
	tooltip = safeString(tooltip)
	ctooltip, _ := unpackPCharString(tooltip)
	C.GuiSetTooltip(ctooltip)
	runtime.KeepAlive(tooltip)
}

// GuiClearTooltip function as declared in src/raygui.h:445
func GuiClearTooltip() {
	C.GuiClearTooltip()
}

// GuiWindowBox function as declared in src/raygui.h:448
func GuiWindowBox(bounds Rectangle, title string) bool {
	cbounds, _ := bounds.PassValue()
	title = safeString(title)
	ctitle, _ := unpackPCharString(title)
	__ret := C.GuiWindowBox(cbounds, ctitle)
	runtime.KeepAlive(title)
	__v := (bool)(__ret)
	return __v
}

// GuiGroupBox function as declared in src/raygui.h:449
func GuiGroupBox(bounds Rectangle, text string) {
	cbounds, _ := bounds.PassValue()
	text = safeString(text)
	ctext, _ := unpackPCharString(text)
	C.GuiGroupBox(cbounds, ctext)
	runtime.KeepAlive(text)
}

// GuiLine function as declared in src/raygui.h:450
func GuiLine(bounds Rectangle, text string) {
	cbounds, _ := bounds.PassValue()
	text = safeString(text)
	ctext, _ := unpackPCharString(text)
	C.GuiLine(cbounds, ctext)
	runtime.KeepAlive(text)
}

// GuiPanel function as declared in src/raygui.h:451
func GuiPanel(bounds Rectangle) {
	cbounds, _ := bounds.PassValue()
	C.GuiPanel(cbounds)
}

// GuiScrollPanel function as declared in src/raygui.h:452
func GuiScrollPanel(bounds Rectangle, content Rectangle, scroll *Vector2) Rectangle {
	cbounds, _ := bounds.PassValue()
	ccontent, _ := content.PassValue()
	cscroll, _ := scroll.PassRef()
	__ret := C.GuiScrollPanel(cbounds, ccontent, cscroll)
	__v := *NewRectangleRef(unsafe.Pointer(&__ret))
	return __v
}

// GuiLabel function as declared in src/raygui.h:455
func GuiLabel(bounds Rectangle, text string) {
	cbounds, _ := bounds.PassValue()
	text = safeString(text)
	ctext, _ := unpackPCharString(text)
	C.GuiLabel(cbounds, ctext)
	runtime.KeepAlive(text)
}

// GuiButton function as declared in src/raygui.h:456
func GuiButton(bounds Rectangle, text string) bool {
	cbounds, _ := bounds.PassValue()
	text = safeString(text)
	ctext, _ := unpackPCharString(text)
	__ret := C.GuiButton(cbounds, ctext)
	runtime.KeepAlive(text)
	__v := (bool)(__ret)
	return __v
}

// GuiLabelButton function as declared in src/raygui.h:457
func GuiLabelButton(bounds Rectangle, text string) bool {
	cbounds, _ := bounds.PassValue()
	text = safeString(text)
	ctext, _ := unpackPCharString(text)
	__ret := C.GuiLabelButton(cbounds, ctext)
	runtime.KeepAlive(text)
	__v := (bool)(__ret)
	return __v
}

// GuiImageButton function as declared in src/raygui.h:458
func GuiImageButton(bounds Rectangle, text string, texture Texture2D) bool {
	cbounds, _ := bounds.PassValue()
	text = safeString(text)
	ctext, _ := unpackPCharString(text)
	ctexture, _ := texture.PassValue()
	__ret := C.GuiImageButton(cbounds, ctext, ctexture)
	runtime.KeepAlive(text)
	__v := (bool)(__ret)
	return __v
}

// GuiImageButtonEx function as declared in src/raygui.h:459
func GuiImageButtonEx(bounds Rectangle, text string, texture Texture2D, texSource Rectangle) bool {
	cbounds, _ := bounds.PassValue()
	text = safeString(text)
	ctext, _ := unpackPCharString(text)
	ctexture, _ := texture.PassValue()
	ctexSource, _ := texSource.PassValue()
	__ret := C.GuiImageButtonEx(cbounds, ctext, ctexture, ctexSource)
	runtime.KeepAlive(text)
	__v := (bool)(__ret)
	return __v
}

// GuiToggle function as declared in src/raygui.h:460
func GuiToggle(bounds Rectangle, text string, active bool) bool {
	cbounds, _ := bounds.PassValue()
	text = safeString(text)
	ctext, _ := unpackPCharString(text)
	cactive, _ := (C._Bool)(active), cgoAllocsUnknown
	__ret := C.GuiToggle(cbounds, ctext, cactive)
	runtime.KeepAlive(text)
	__v := (bool)(__ret)
	return __v
}

// GuiToggleGroup function as declared in src/raygui.h:461
func GuiToggleGroup(bounds Rectangle, text string, active int32) int32 {
	cbounds, _ := bounds.PassValue()
	text = safeString(text)
	ctext, _ := unpackPCharString(text)
	cactive, _ := (C.int)(active), cgoAllocsUnknown
	__ret := C.GuiToggleGroup(cbounds, ctext, cactive)
	runtime.KeepAlive(text)
	__v := (int32)(__ret)
	return __v
}

// GuiCheckBox function as declared in src/raygui.h:462
func GuiCheckBox(bounds Rectangle, text string, checked bool) bool {
	cbounds, _ := bounds.PassValue()
	text = safeString(text)
	ctext, _ := unpackPCharString(text)
	cchecked, _ := (C._Bool)(checked), cgoAllocsUnknown
	__ret := C.GuiCheckBox(cbounds, ctext, cchecked)
	runtime.KeepAlive(text)
	__v := (bool)(__ret)
	return __v
}

// GuiComboBox function as declared in src/raygui.h:463
func GuiComboBox(bounds Rectangle, text string, active int32) int32 {
	cbounds, _ := bounds.PassValue()
	text = safeString(text)
	ctext, _ := unpackPCharString(text)
	cactive, _ := (C.int)(active), cgoAllocsUnknown
	__ret := C.GuiComboBox(cbounds, ctext, cactive)
	runtime.KeepAlive(text)
	__v := (int32)(__ret)
	return __v
}

// GuiDropdownBox function as declared in src/raygui.h:464
func GuiDropdownBox(bounds Rectangle, text string, active *int32, editMode bool) bool {
	cbounds, _ := bounds.PassValue()
	text = safeString(text)
	ctext, _ := unpackPCharString(text)
	cactive, _ := (*C.int)(unsafe.Pointer(active)), cgoAllocsUnknown
	ceditMode, _ := (C._Bool)(editMode), cgoAllocsUnknown
	__ret := C.GuiDropdownBox(cbounds, ctext, cactive, ceditMode)
	runtime.KeepAlive(text)
	__v := (bool)(__ret)
	return __v
}

// GuiSpinner function as declared in src/raygui.h:465
func GuiSpinner(bounds Rectangle, text string, value *int32, minValue int32, maxValue int32, editMode bool) bool {
	cbounds, _ := bounds.PassValue()
	text = safeString(text)
	ctext, _ := unpackPCharString(text)
	cvalue, _ := (*C.int)(unsafe.Pointer(value)), cgoAllocsUnknown
	cminValue, _ := (C.int)(minValue), cgoAllocsUnknown
	cmaxValue, _ := (C.int)(maxValue), cgoAllocsUnknown
	ceditMode, _ := (C._Bool)(editMode), cgoAllocsUnknown
	__ret := C.GuiSpinner(cbounds, ctext, cvalue, cminValue, cmaxValue, ceditMode)
	runtime.KeepAlive(text)
	__v := (bool)(__ret)
	return __v
}

// GuiValueBox function as declared in src/raygui.h:466
func GuiValueBox(bounds Rectangle, text string, value *int32, minValue int32, maxValue int32, editMode bool) bool {
	cbounds, _ := bounds.PassValue()
	text = safeString(text)
	ctext, _ := unpackPCharString(text)
	cvalue, _ := (*C.int)(unsafe.Pointer(value)), cgoAllocsUnknown
	cminValue, _ := (C.int)(minValue), cgoAllocsUnknown
	cmaxValue, _ := (C.int)(maxValue), cgoAllocsUnknown
	ceditMode, _ := (C._Bool)(editMode), cgoAllocsUnknown
	__ret := C.GuiValueBox(cbounds, ctext, cvalue, cminValue, cmaxValue, ceditMode)
	runtime.KeepAlive(text)
	__v := (bool)(__ret)
	return __v
}

// GuiTextBox function as declared in src/raygui.h:467
func GuiTextBox(bounds Rectangle, text *byte, textSize int32, editMode bool) bool {
	cbounds, _ := bounds.PassValue()
	ctext, _ := (*C.char)(unsafe.Pointer(text)), cgoAllocsUnknown
	ctextSize, _ := (C.int)(textSize), cgoAllocsUnknown
	ceditMode, _ := (C._Bool)(editMode), cgoAllocsUnknown
	__ret := C.GuiTextBox(cbounds, ctext, ctextSize, ceditMode)
	__v := (bool)(__ret)
	return __v
}

// GuiTextBoxMulti function as declared in src/raygui.h:468
func GuiTextBoxMulti(bounds Rectangle, text *byte, textSize int32, editMode bool) bool {
	cbounds, _ := bounds.PassValue()
	ctext, _ := (*C.char)(unsafe.Pointer(text)), cgoAllocsUnknown
	ctextSize, _ := (C.int)(textSize), cgoAllocsUnknown
	ceditMode, _ := (C._Bool)(editMode), cgoAllocsUnknown
	__ret := C.GuiTextBoxMulti(cbounds, ctext, ctextSize, ceditMode)
	__v := (bool)(__ret)
	return __v
}

// GuiSlider function as declared in src/raygui.h:469
func GuiSlider(bounds Rectangle, textLeft string, textRight string, value float32, minValue float32, maxValue float32) float32 {
	cbounds, _ := bounds.PassValue()
	textLeft = safeString(textLeft)
	ctextLeft, _ := unpackPCharString(textLeft)
	textRight = safeString(textRight)
	ctextRight, _ := unpackPCharString(textRight)
	cvalue, _ := (C.float)(value), cgoAllocsUnknown
	cminValue, _ := (C.float)(minValue), cgoAllocsUnknown
	cmaxValue, _ := (C.float)(maxValue), cgoAllocsUnknown
	__ret := C.GuiSlider(cbounds, ctextLeft, ctextRight, cvalue, cminValue, cmaxValue)
	runtime.KeepAlive(textRight)
	runtime.KeepAlive(textLeft)
	__v := (float32)(__ret)
	return __v
}

// GuiSliderBar function as declared in src/raygui.h:470
func GuiSliderBar(bounds Rectangle, textLeft string, textRight string, value float32, minValue float32, maxValue float32) float32 {
	cbounds, _ := bounds.PassValue()
	textLeft = safeString(textLeft)
	ctextLeft, _ := unpackPCharString(textLeft)
	textRight = safeString(textRight)
	ctextRight, _ := unpackPCharString(textRight)
	cvalue, _ := (C.float)(value), cgoAllocsUnknown
	cminValue, _ := (C.float)(minValue), cgoAllocsUnknown
	cmaxValue, _ := (C.float)(maxValue), cgoAllocsUnknown
	__ret := C.GuiSliderBar(cbounds, ctextLeft, ctextRight, cvalue, cminValue, cmaxValue)
	runtime.KeepAlive(textRight)
	runtime.KeepAlive(textLeft)
	__v := (float32)(__ret)
	return __v
}

// GuiProgressBar function as declared in src/raygui.h:471
func GuiProgressBar(bounds Rectangle, textLeft string, textRight string, value float32, minValue float32, maxValue float32) float32 {
	cbounds, _ := bounds.PassValue()
	textLeft = safeString(textLeft)
	ctextLeft, _ := unpackPCharString(textLeft)
	textRight = safeString(textRight)
	ctextRight, _ := unpackPCharString(textRight)
	cvalue, _ := (C.float)(value), cgoAllocsUnknown
	cminValue, _ := (C.float)(minValue), cgoAllocsUnknown
	cmaxValue, _ := (C.float)(maxValue), cgoAllocsUnknown
	__ret := C.GuiProgressBar(cbounds, ctextLeft, ctextRight, cvalue, cminValue, cmaxValue)
	runtime.KeepAlive(textRight)
	runtime.KeepAlive(textLeft)
	__v := (float32)(__ret)
	return __v
}

// GuiStatusBar function as declared in src/raygui.h:472
func GuiStatusBar(bounds Rectangle, text string) {
	cbounds, _ := bounds.PassValue()
	text = safeString(text)
	ctext, _ := unpackPCharString(text)
	C.GuiStatusBar(cbounds, ctext)
	runtime.KeepAlive(text)
}

// GuiDummyRec function as declared in src/raygui.h:473
func GuiDummyRec(bounds Rectangle, text string) {
	cbounds, _ := bounds.PassValue()
	text = safeString(text)
	ctext, _ := unpackPCharString(text)
	C.GuiDummyRec(cbounds, ctext)
	runtime.KeepAlive(text)
}

// GuiScrollBar function as declared in src/raygui.h:474
func GuiScrollBar(bounds Rectangle, value int32, minValue int32, maxValue int32) int32 {
	cbounds, _ := bounds.PassValue()
	cvalue, _ := (C.int)(value), cgoAllocsUnknown
	cminValue, _ := (C.int)(minValue), cgoAllocsUnknown
	cmaxValue, _ := (C.int)(maxValue), cgoAllocsUnknown
	__ret := C.GuiScrollBar(cbounds, cvalue, cminValue, cmaxValue)
	__v := (int32)(__ret)
	return __v
}

// GuiGrid function as declared in src/raygui.h:475
func GuiGrid(bounds Rectangle, spacing float32, subdivs int32) Vector2 {
	cbounds, _ := bounds.PassValue()
	cspacing, _ := (C.float)(spacing), cgoAllocsUnknown
	csubdivs, _ := (C.int)(subdivs), cgoAllocsUnknown
	__ret := C.GuiGrid(cbounds, cspacing, csubdivs)
	__v := *NewVector2Ref(unsafe.Pointer(&__ret))
	return __v
}

// GuiListView function as declared in src/raygui.h:478
func GuiListView(bounds Rectangle, text string, scrollIndex *int32, active int32) int32 {
	cbounds, _ := bounds.PassValue()
	text = safeString(text)
	ctext, _ := unpackPCharString(text)
	cscrollIndex, _ := (*C.int)(unsafe.Pointer(scrollIndex)), cgoAllocsUnknown
	cactive, _ := (C.int)(active), cgoAllocsUnknown
	__ret := C.GuiListView(cbounds, ctext, cscrollIndex, cactive)
	runtime.KeepAlive(text)
	__v := (int32)(__ret)
	return __v
}

// GuiMessageBox function as declared in src/raygui.h:480
func GuiMessageBox(bounds Rectangle, title string, message string, buttons string) int32 {
	cbounds, _ := bounds.PassValue()
	title = safeString(title)
	ctitle, _ := unpackPCharString(title)
	message = safeString(message)
	cmessage, _ := unpackPCharString(message)
	buttons = safeString(buttons)
	cbuttons, _ := unpackPCharString(buttons)
	__ret := C.GuiMessageBox(cbounds, ctitle, cmessage, cbuttons)
	runtime.KeepAlive(buttons)
	runtime.KeepAlive(message)
	runtime.KeepAlive(title)
	__v := (int32)(__ret)
	return __v
}

// GuiTextInputBox function as declared in src/raygui.h:481
func GuiTextInputBox(bounds Rectangle, title string, message string, buttons string, text *byte) int32 {
	cbounds, _ := bounds.PassValue()
	title = safeString(title)
	ctitle, _ := unpackPCharString(title)
	message = safeString(message)
	cmessage, _ := unpackPCharString(message)
	buttons = safeString(buttons)
	cbuttons, _ := unpackPCharString(buttons)
	ctext, _ := (*C.char)(unsafe.Pointer(text)), cgoAllocsUnknown
	__ret := C.GuiTextInputBox(cbounds, ctitle, cmessage, cbuttons, ctext)
	runtime.KeepAlive(buttons)
	runtime.KeepAlive(message)
	runtime.KeepAlive(title)
	__v := (int32)(__ret)
	return __v
}

// GuiColorPicker function as declared in src/raygui.h:482
func GuiColorPicker(bounds Rectangle, color Color) Color {
	cbounds, _ := bounds.PassValue()
	ccolor, _ := color.PassValue()
	__ret := C.GuiColorPicker(cbounds, ccolor)
	__v := *NewColorRef(unsafe.Pointer(&__ret))
	return __v
}

// GuiColorPanel function as declared in src/raygui.h:483
func GuiColorPanel(bounds Rectangle, color Color) Color {
	cbounds, _ := bounds.PassValue()
	ccolor, _ := color.PassValue()
	__ret := C.GuiColorPanel(cbounds, ccolor)
	__v := *NewColorRef(unsafe.Pointer(&__ret))
	return __v
}

// GuiColorBarAlpha function as declared in src/raygui.h:484
func GuiColorBarAlpha(bounds Rectangle, alpha float32) float32 {
	cbounds, _ := bounds.PassValue()
	calpha, _ := (C.float)(alpha), cgoAllocsUnknown
	__ret := C.GuiColorBarAlpha(cbounds, calpha)
	__v := (float32)(__ret)
	return __v
}

// GuiColorBarHue function as declared in src/raygui.h:485
func GuiColorBarHue(bounds Rectangle, value float32) float32 {
	cbounds, _ := bounds.PassValue()
	cvalue, _ := (C.float)(value), cgoAllocsUnknown
	__ret := C.GuiColorBarHue(cbounds, cvalue)
	__v := (float32)(__ret)
	return __v
}

// GuiLoadStyle function as declared in src/raygui.h:488
func GuiLoadStyle(fileName string) {
	fileName = safeString(fileName)
	cfileName, _ := unpackPCharString(fileName)
	C.GuiLoadStyle(cfileName)
	runtime.KeepAlive(fileName)
}

// GuiLoadStyleDefault function as declared in src/raygui.h:489
func GuiLoadStyleDefault() {
	C.GuiLoadStyleDefault()
}

// GuiIconText function as declared in src/raygui.h:497
func GuiIconText(iconId int32, text string) string {
	ciconId, _ := (C.int)(iconId), cgoAllocsUnknown
	text = safeString(text)
	ctext, _ := unpackPCharString(text)
	__ret := C.GuiIconText(ciconId, ctext)
	runtime.KeepAlive(text)
	__v := packPCharString(__ret)
	return __v
}

// GuiColorPanelEx function as declared in src/raygui.h:500
func GuiColorPanelEx(bounds Rectangle, color Color, hue float32) Color {
	cbounds, _ := bounds.PassValue()
	ccolor, _ := color.PassValue()
	chue, _ := (C.float)(hue), cgoAllocsUnknown
	__ret := C.GuiColorPanelEx(cbounds, ccolor, chue)
	__v := *NewColorRef(unsafe.Pointer(&__ret))
	return __v
}

// GuiSliderPro function as declared in src/raygui.h:501
func GuiSliderPro(bounds Rectangle, textLeft string, textRight string, value float32, minValue float32, maxValue float32, sliderWidth int32) float32 {
	cbounds, _ := bounds.PassValue()
	textLeft = safeString(textLeft)
	ctextLeft, _ := unpackPCharString(textLeft)
	textRight = safeString(textRight)
	ctextRight, _ := unpackPCharString(textRight)
	cvalue, _ := (C.float)(value), cgoAllocsUnknown
	cminValue, _ := (C.float)(minValue), cgoAllocsUnknown
	cmaxValue, _ := (C.float)(maxValue), cgoAllocsUnknown
	csliderWidth, _ := (C.int)(sliderWidth), cgoAllocsUnknown
	__ret := C.GuiSliderPro(cbounds, ctextLeft, ctextRight, cvalue, cminValue, cmaxValue, csliderWidth)
	runtime.KeepAlive(textRight)
	runtime.KeepAlive(textLeft)
	__v := (float32)(__ret)
	return __v
}
