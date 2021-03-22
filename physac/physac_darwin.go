// +build darwin

package physac

/*
#cgo darwin LDFLAGS: -framework OpenGL -framework Cocoa -framework IOKit -framework CoreVideo -framework CoreFoundation -L${SRCDIR}/../lib/raylib/a/darwin -lraylib
#cgo darwin CFLAGS: -I../lib/raylib/src -DPHYSAC_IMPLEMENTATION -DPHYSAC_STATIC -DPHYSAC_DEFINE_VECTOR2_TYPE
*/
import "C"
