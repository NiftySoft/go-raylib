// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY ROBOTS.

// WARNING: This file has automatically been generated on Mon, 03 Aug 2020 11:53:50 CST.
// Code generated by https://git.io/c-for-go. DO NOT EDIT.

package physac

/*
#include "../lib/raylib/src/physac.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"

// Vector2 as declared in src/physac.h:130
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

// Matrix2x2 as declared in src/physac.h:142
type Matrix2x2 struct {
	M00            float32
	M01            float32
	M10            float32
	M11            float32
	refb92f06e6    *C.Matrix2x2
	allocsb92f06e6 interface{}
}
type matrix2x2 struct {
	M00 float32
	M01 float32
	M10 float32
	M11 float32
}

// PolygonData as declared in src/physac.h:148
type PolygonData struct {
	VertexCount    uint32
	Positions      [24]Vector2
	Normals        [24]Vector2
	ref87a94eb7    *C.PolygonData
	allocs87a94eb7 interface{}
}
type polygonData struct {
	VertexCount uint32
	Positions   [24]C.Vector2
	Normals     [24]C.Vector2
}

// PhysicsShape as declared in src/physac.h:157
type PhysicsShape struct {
	Type           PhysicsShapeType
	Body           []PhysicsBodyData
	Radius         float32
	Transform      Matrix2x2
	VertexData     PolygonData
	ref4c540f1a    *C.PhysicsShape
	allocs4c540f1a interface{}
}
type physicsShape struct {
	Type       PhysicsShapeType
	Body       *C.PhysicsBodyData
	Radius     float32
	Transform  C.Matrix2x2
	VertexData C.PolygonData
}

// PhysicsBodyData as declared in src/physac.h:179
type PhysicsBodyData struct {
	Id              uint32
	Enabled         bool
	Position        Vector2
	Velocity        Vector2
	Force           Vector2
	AngularVelocity float32
	Torque          float32
	Orient          float32
	Inertia         float32
	InverseInertia  float32
	Mass            float32
	InverseMass     float32
	StaticFriction  float32
	DynamicFriction float32
	Restitution     float32
	UseGravity      bool
	IsGrounded      bool
	FreezeOrient    bool
	Shape           PhysicsShape
	refd780e53      *C.PhysicsBodyData
	allocsd780e53   interface{}
}
type physicsBodyData struct {
	Id              uint32
	Enabled         bool
	Position        C.Vector2
	Velocity        C.Vector2
	Force           C.Vector2
	AngularVelocity float32
	Torque          float32
	Orient          float32
	Inertia         float32
	InverseInertia  float32
	Mass            float32
	InverseMass     float32
	StaticFriction  float32
	DynamicFriction float32
	Restitution     float32
	UseGravity      bool
	IsGrounded      bool
	FreezeOrient    bool
	Shape           C.PhysicsShape
}

// PhysicsManifoldData as declared in src/physac.h:194
type PhysicsManifoldData struct {
	Id              uint32
	BodyA           []PhysicsBodyData
	BodyB           []PhysicsBodyData
	Penetration     float32
	Normal          Vector2
	Contacts        [2]Vector2
	ContactsCount   uint32
	Restitution     float32
	DynamicFriction float32
	StaticFriction  float32
	ref10b92967     *C.PhysicsManifoldData
	allocs10b92967  interface{}
}
type physicsManifoldData struct {
	Id              uint32
	BodyA           *C.PhysicsBodyData
	BodyB           *C.PhysicsBodyData
	Penetration     float32
	Normal          C.Vector2
	Contacts        [2]C.Vector2
	ContactsCount   uint32
	Restitution     float32
	DynamicFriction float32
	StaticFriction  float32
}
