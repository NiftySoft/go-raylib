// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY ROBOTS.

// WARNING: This file has automatically been generated on Tue, 23 Mar 2021 11:35:17 CST.
// Code generated by https://git.io/c-for-go. DO NOT EDIT.

package raylib

/*
#include "../lib/raylib/src/raylib.h"
#include "../lib/raylib/src/raymath.h"
#include "../lib/raylib/src/rlgl.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"
import "unsafe"

// Vector2 as declared in src/raylib.h:180
type gVector2 struct {
	gX             float32
	gY             float32
	ref29ca61a5    *C.Vector2
	allocs29ca61a5 interface{}
}
type Vector2 struct {
	X float32
	Y float32
}

// Vector3 as declared in src/raylib.h:187
type gVector3 struct {
	gX             float32
	gY             float32
	gZ             float32
	ref5ecd5133    *C.Vector3
	allocs5ecd5133 interface{}
}
type Vector3 struct {
	X float32
	Y float32
	Z float32
}

// Vector4 as declared in src/raylib.h:195
type gVector4 struct {
	gX             float32
	gY             float32
	gZ             float32
	gW             float32
	refc0a9c490    *C.Vector4
	allocsc0a9c490 interface{}
}
type Vector4 struct {
	X float32
	Y float32
	Z float32
	W float32
}

// Quaternion as declared in src/raylib.h:198
type gQuaternion struct {
	gX             float32
	gY             float32
	gZ             float32
	gW             float32
	refc0a9c490    *C.Vector4
	allocsc0a9c490 interface{}
}
type Quaternion struct {
	X float32
	Y float32
	Z float32
	W float32
}

// Matrix as declared in src/raylib.h:206
type gMatrix struct {
	gM0            float32
	gM4            float32
	gM8            float32
	gM12           float32
	gM1            float32
	gM5            float32
	gM9            float32
	gM13           float32
	gM2            float32
	gM6            float32
	gM10           float32
	gM14           float32
	gM3            float32
	gM7            float32
	gM11           float32
	gM15           float32
	refff9f44f9    *C.Matrix
	allocsff9f44f9 interface{}
}
type Matrix struct {
	M0  float32
	M4  float32
	M8  float32
	M12 float32
	M1  float32
	M5  float32
	M9  float32
	M13 float32
	M2  float32
	M6  float32
	M10 float32
	M14 float32
	M3  float32
	M7  float32
	M11 float32
	M15 float32
}

// Color as declared in src/raylib.h:214
type gColor struct {
	gR             byte
	gG             byte
	gB             byte
	gA             byte
	refa79767ed    *C.Color
	allocsa79767ed interface{}
}
type Color struct {
	R byte
	G byte
	B byte
	A byte
}

// Rectangle as declared in src/raylib.h:222
type gRectangle struct {
	gX             float32
	gY             float32
	gWidth         float32
	gHeight        float32
	refcee8783a    *C.Rectangle
	allocscee8783a interface{}
}
type Rectangle struct {
	X      float32
	Y      float32
	Width  float32
	Height float32
}

// Image as declared in src/raylib.h:232
type gImage struct {
	gData         unsafe.Pointer
	gWidth        int32
	gHeight       int32
	gMipmaps      int32
	gFormat       int32
	ref4fc2b5b    *C.Image
	allocs4fc2b5b interface{}
}
type Image struct {
	Data    unsafe.Pointer
	Width   int32
	Height  int32
	Mipmaps int32
	Format  int32
}

// Texture as declared in src/raylib.h:242
type gTexture struct {
	gId            uint32
	gWidth         int32
	gHeight        int32
	gMipmaps       int32
	gFormat        int32
	ref4ddb34ee    *C.Texture
	allocs4ddb34ee interface{}
}
type Texture struct {
	Id      uint32
	Width   int32
	Height  int32
	Mipmaps int32
	Format  int32
}

// Texture2D as declared in src/raylib.h:245
type gTexture2D struct {
	gId            uint32
	gWidth         int32
	gHeight        int32
	gMipmaps       int32
	gFormat        int32
	ref4ddb34ee    *C.Texture
	allocs4ddb34ee interface{}
}
type Texture2D struct {
	Id      uint32
	Width   int32
	Height  int32
	Mipmaps int32
	Format  int32
}

// TextureCubemap as declared in src/raylib.h:248
type gTextureCubemap struct {
	gId            uint32
	gWidth         int32
	gHeight        int32
	gMipmaps       int32
	gFormat        int32
	ref4ddb34ee    *C.Texture
	allocs4ddb34ee interface{}
}
type TextureCubemap struct {
	Id      uint32
	Width   int32
	Height  int32
	Mipmaps int32
	Format  int32
}

// RenderTexture as declared in src/raylib.h:255
type gRenderTexture struct {
	gId            uint32
	gTexture       gTexture
	gDepth         gTexture
	ref24cb7f02    *C.RenderTexture
	allocs24cb7f02 interface{}
}
type RenderTexture struct {
	Id      uint32
	Texture Texture
	Depth   Texture
}

// RenderTexture2D as declared in src/raylib.h:258
type gRenderTexture2D struct {
	gId            uint32
	gTexture       gTexture
	gDepth         gTexture
	ref24cb7f02    *C.RenderTexture
	allocs24cb7f02 interface{}
}
type RenderTexture2D struct {
	Id      uint32
	Texture Texture
	Depth   Texture
}

// NPatchInfo as declared in src/raylib.h:268
type gNPatchInfo struct {
	gSource        gRectangle
	gLeft          int32
	gTop           int32
	gRight         int32
	gBottom        int32
	gLayout        int32
	ref78104a03    *C.NPatchInfo
	allocs78104a03 interface{}
}
type NPatchInfo struct {
	Source Rectangle
	Left   int32
	Top    int32
	Right  int32
	Bottom int32
	Layout int32
}

// CharInfo as declared in src/raylib.h:277
type gCharInfo struct {
	gValue         int32
	gOffsetX       int32
	gOffsetY       int32
	gAdvanceX      int32
	gImage         gImage
	ref702c36c0    *C.CharInfo
	allocs702c36c0 interface{}
}
type CharInfo struct {
	Value    int32
	OffsetX  int32
	OffsetY  int32
	AdvanceX int32
	Image    Image
}

// Font as declared in src/raylib.h:287
type gFont struct {
	gBaseSize      int32
	gCharsCount    int32
	gCharsPadding  int32
	gTexture       gTexture2D
	gRecs          []gRectangle
	gChars         []gCharInfo
	ref70a6a7ec    *C.Font
	allocs70a6a7ec interface{}
}
type Font struct {
	BaseSize     int32
	CharsCount   int32
	CharsPadding int32
	Texture      Texture
	Recs         *Rectangle
	Chars        *CharInfo
}

// Camera3D as declared in src/raylib.h:298
type gCamera3D struct {
	gPosition     gVector3
	gTarget       gVector3
	gUp           gVector3
	gFovy         float32
	gProjection   int32
	ref7b09036    *C.Camera3D
	allocs7b09036 interface{}
}
type Camera3D struct {
	Position   Vector3
	Target     Vector3
	Up         Vector3
	Fovy       float32
	Projection int32
}

// Camera as declared in src/raylib.h:300
type gCamera struct {
	gPosition     gVector3
	gTarget       gVector3
	gUp           gVector3
	gFovy         float32
	gProjection   int32
	ref7b09036    *C.Camera3D
	allocs7b09036 interface{}
}
type Camera struct {
	Position   Vector3
	Target     Vector3
	Up         Vector3
	Fovy       float32
	Projection int32
}

// Camera2D as declared in src/raylib.h:308
type gCamera2D struct {
	gOffset        gVector2
	gTarget        gVector2
	gRotation      float32
	gZoom          float32
	ref1eaba177    *C.Camera2D
	allocs1eaba177 interface{}
}
type Camera2D struct {
	Offset   Vector2
	Target   Vector2
	Rotation float32
	Zoom     float32
}

// Mesh as declared in src/raylib.h:334
type gMesh struct {
	gVertexCount   int32
	gTriangleCount int32
	gVertices      []float32
	gTexcoords     []float32
	gTexcoords2    []float32
	gNormals       []float32
	gTangents      []float32
	gColors        []byte
	gIndices       []uint16
	gAnimVertices  []float32
	gAnimNormals   []float32
	gBoneIds       []int32
	gBoneWeights   []float32
	gVaoId         uint32
	gVboId         []uint32
	ref415d9568    *C.Mesh
	allocs415d9568 interface{}
}
type Mesh struct {
	VertexCount   int32
	TriangleCount int32
	Vertices      *float32
	Texcoords     *float32
	Texcoords2    *float32
	Normals       *float32
	Tangents      *float32
	Colors        *byte
	Indices       *uint16
	AnimVertices  *float32
	AnimNormals   *float32
	BoneIds       *int32
	BoneWeights   *float32
	VaoId         uint32
	VboId         *uint32
}

// Shader as declared in src/raylib.h:340
type gShader struct {
	gId            uint32
	gLocs          []int32
	reff85f9b1e    *C.Shader
	allocsf85f9b1e interface{}
}
type Shader struct {
	Id   uint32
	Locs *int32
}

// MaterialMap as declared in src/raylib.h:347
type gMaterialMap struct {
	gTexture       gTexture2D
	gColor         gColor
	gValue         float32
	refa8350ad3    *C.MaterialMap
	allocsa8350ad3 interface{}
}
type MaterialMap struct {
	Texture Texture
	Color   Color
	Value   float32
}

// Material as declared in src/raylib.h:354
type gMaterial struct {
	gShader        gShader
	gMaps          []gMaterialMap
	gParams        [4]float32
	ref85c817c3    *C.Material
	allocs85c817c3 interface{}
}
type Material struct {
	Shader Shader
	Maps   *MaterialMap
	Params [4]float32
}

// Transform as declared in src/raylib.h:361
type gTransform struct {
	gTranslation   gVector3
	gRotation      gQuaternion
	gScale         gVector3
	reff543030e    *C.Transform
	allocsf543030e interface{}
}
type Transform struct {
	Translation Vector3
	Rotation    Vector4
	Scale       Vector3
}

// BoneInfo as declared in src/raylib.h:367
type gBoneInfo struct {
	gName          [32]byte
	gParent        int32
	ref5ab7f197    *C.BoneInfo
	allocs5ab7f197 interface{}
}
type BoneInfo struct {
	Name   [32]byte
	Parent int32
}

// Model as declared in src/raylib.h:383
type gModel struct {
	gTransform     gMatrix
	gMeshCount     int32
	gMaterialCount int32
	gMeshes        []gMesh
	gMaterials     []gMaterial
	gMeshMaterial  []int32
	gBoneCount     int32
	gBones         []gBoneInfo
	gBindPose      []gTransform
	ref16545ddd    *C.Model
	allocs16545ddd interface{}
}
type Model struct {
	Transform     Matrix
	MeshCount     int32
	MaterialCount int32
	Meshes        *Mesh
	Materials     *Material
	MeshMaterial  *int32
	BoneCount     int32
	Bones         *BoneInfo
	BindPose      *Transform
}

// ModelAnimation as declared in src/raylib.h:391
type gModelAnimation struct {
	gBoneCount     int32
	gFrameCount    int32
	gBones         []gBoneInfo
	gFramePoses    [][]gTransform
	ref26dd6a24    *C.ModelAnimation
	allocs26dd6a24 interface{}
}
type ModelAnimation struct {
	BoneCount  int32
	FrameCount int32
	Bones      *BoneInfo
	FramePoses **Transform
}

// Ray as declared in src/raylib.h:397
type gRay struct {
	gPosition      gVector3
	gDirection     gVector3
	refc546b0b2    *C.Ray
	allocsc546b0b2 interface{}
}
type Ray struct {
	Position  Vector3
	Direction Vector3
}

// RayHitInfo as declared in src/raylib.h:405
type gRayHitInfo struct {
	gHit           bool
	gDistance      float32
	gPosition      gVector3
	gNormal        gVector3
	refb8de43a9    *C.RayHitInfo
	allocsb8de43a9 interface{}
}
type RayHitInfo struct {
	Hit      bool
	Distance float32
	Position Vector3
	Normal   Vector3
}

// BoundingBox as declared in src/raylib.h:411
type gBoundingBox struct {
	gMin           gVector3
	gMax           gVector3
	refa54e9d16    *C.BoundingBox
	allocsa54e9d16 interface{}
}
type BoundingBox struct {
	Min Vector3
	Max Vector3
}

// Wave as declared in src/raylib.h:420
type gWave struct {
	gSampleCount   uint32
	gSampleRate    uint32
	gSampleSize    uint32
	gChannels      uint32
	gData          unsafe.Pointer
	ref7a3602b7    *C.Wave
	allocs7a3602b7 interface{}
}
type Wave struct {
	SampleCount uint32
	SampleRate  uint32
	SampleSize  uint32
	Channels    uint32
	Data        unsafe.Pointer
}

// AudioStream as declared in src/raylib.h:434
type gAudioStream struct {
	gSampleRate    uint32
	gSampleSize    uint32
	gChannels      uint32
	gBuffer        *C.rAudioBuffer
	ref997374a2    *C.AudioStream
	allocs997374a2 interface{}
}
type AudioStream struct {
	SampleRate uint32
	SampleSize uint32
	Channels   uint32
	Buffer     *C.rAudioBuffer
}

// Sound as declared in src/raylib.h:440
type gSound struct {
	gStream        gAudioStream
	gSampleCount   uint32
	ref394fec80    *C.Sound
	allocs394fec80 interface{}
}
type Sound struct {
	Stream      AudioStream
	SampleCount uint32
}

// Music as declared in src/raylib.h:451
type gMusic struct {
	gStream       gAudioStream
	gSampleCount  uint32
	gLooping      bool
	gCtxType      int32
	gCtxData      unsafe.Pointer
	refc930d4e    *C.Music
	allocsc930d4e interface{}
}
type Music struct {
	Stream      AudioStream
	SampleCount uint32
	Looping     bool
	CtxType     int32
	CtxData     unsafe.Pointer
}

// VrDeviceInfo as declared in src/raylib.h:465
type gVrDeviceInfo struct {
	gHResolution            int32
	gVResolution            int32
	gHScreenSize            float32
	gVScreenSize            float32
	gVScreenCenter          float32
	gEyeToScreenDistance    float32
	gLensSeparationDistance float32
	gInterpupillaryDistance float32
	gLensDistortionValues   [4]float32
	gChromaAbCorrection     [4]float32
	ref6e24e41d             *C.VrDeviceInfo
	allocs6e24e41d          interface{}
}
type VrDeviceInfo struct {
	HResolution            int32
	VResolution            int32
	HScreenSize            float32
	VScreenSize            float32
	VScreenCenter          float32
	EyeToScreenDistance    float32
	LensSeparationDistance float32
	InterpupillaryDistance float32
	LensDistortionValues   [4]float32
	ChromaAbCorrection     [4]float32
}

// VrStereoConfig as declared in src/raylib.h:475
type gVrStereoConfig struct {
	gLeftLensCenter    [2]float32
	gRightLensCenter   [2]float32
	gLeftScreenCenter  [2]float32
	gRightScreenCenter [2]float32
	gScale             [2]float32
	gScaleIn           [2]float32
	refbf4ea4f3        *C.VrStereoConfig
	allocsbf4ea4f3     interface{}
}
type VrStereoConfig struct {
	LeftLensCenter    [2]float32
	RightLensCenter   [2]float32
	LeftScreenCenter  [2]float32
	RightScreenCenter [2]float32
	Scale             [2]float32
	ScaleIn           [2]float32
}

// MultiText as declared in src/raylib.h:480
type gMultiText struct {
	gText          []string
	refdf1ec495    *C.MultiText
	allocsdf1ec495 interface{}
}
type MultiText struct {
	Text *string
}
