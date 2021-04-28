// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY ROBOTS.

// WARNING: This file has automatically been generated on Wed, 28 Apr 2021 17:24:23 CST.
// Code generated by https://git.io/c-for-go. DO NOT EDIT.

package physac

/*
#include "../lib/raylib/src/physac.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"

const (
	// PHYSAC_STANDALONE as defined in go-raylib/<predefine>:24
	PHYSAC_STANDALONE = 1
	// PHYSAC_MAX_BODIES as defined in src/physac.h:102
	PHYSAC_MAX_BODIES = 64
	// PHYSAC_MAX_MANIFOLDS as defined in src/physac.h:103
	PHYSAC_MAX_MANIFOLDS = 4096
	// PHYSAC_MAX_VERTICES as defined in src/physac.h:104
	PHYSAC_MAX_VERTICES = 24
	// PHYSAC_CIRCLE_VERTICES as defined in src/physac.h:105
	PHYSAC_CIRCLE_VERTICES = 24
	// PHYSAC_COLLISION_ITERATIONS as defined in src/physac.h:107
	PHYSAC_COLLISION_ITERATIONS = 100
	// PHYSAC_PENETRATION_ALLOWANCE as defined in src/physac.h:108
	PHYSAC_PENETRATION_ALLOWANCE = 0.05
	// PHYSAC_PENETRATION_CORRECTION as defined in src/physac.h:109
	PHYSAC_PENETRATION_CORRECTION = 0.4
	// PHYSAC_PI as defined in src/physac.h:111
	PHYSAC_PI = 3.14159265358979323846
	// PHYSAC_DEG2RAD as defined in src/physac.h:112
	PHYSAC_DEG2RAD = (PHYSAC_PI / 180.0)
	// PHYSAC_FLT_MAX as defined in src/physac.h:285
	PHYSAC_FLT_MAX = 3.402823466e+38
	// PHYSAC_EPSILON as defined in src/physac.h:286
	PHYSAC_EPSILON = 0.000001
	// PHYSAC_K as defined in src/physac.h:287
	PHYSAC_K = 1.0 / 3.0
)

// PhysicsShapeType as declared in src/physac.h:133
type PhysicsShapeType int32

// PhysicsShapeType enumeration from src/physac.h:133
const (
	PHYSICS_CIRCLE  PhysicsShapeType = iota
	PHYSICS_POLYGON PhysicsShapeType = 1
)

const ()
