package gocv

/*
#include <stdlib.h>
#include "core.h"
*/
import "C"
import (
	"image"
	"unsafe"
)

// Mat represents an n-dimensional dense numerical single-channel
// or multi-channel array. It can be used to store real or complex-valued
// vectors and matrices, grayscale or color images, voxel volumes,
// vector fields, point clouds, tensors, and histograms.
//
// For further details, please see:
// http://docs.opencv.org/3.3.0/d3/d63/classcv_1_1Mat.html
//
type Mat struct {
	p C.Mat
}

// NewMat returns a new Mat.
func NewMat() Mat {
	return Mat{p: C.Mat_New()}
}

// NewMatWithSize returns a new Mat with a specific size.
func NewMatWithSize(rows int, cols int, mt MatType) Mat {
	return Mat{p: C.Mat_NewWithSize(C.int(rows), C.int(cols), C.int(mt))}
}

// Close the Mat object.
func (m *Mat) Close() error {
	C.Mat_Close(m.p)
	m.p = nil
	return nil
}

// Ptr returns the Mat's underlying object pointer.
func (m *Mat) Ptr() C.Mat {
	return m.p
}

// Empty determines if the Mat is empty or not.
func (m *Mat) Empty() bool {
	isEmpty := C.Mat_Empty(m.p)
	return isEmpty != 0
}

// Region returns a new Mat that points to a region of this Mat. Changes made to the
// region Mat will affect the original Mat, sinec they are pointers to the underlying
// OpenCV Mat object.
func (m *Mat) Region(rio image.Rectangle) Mat {
	cRect := C.struct_Rect{
		x:      C.int(rio.Min.X),
		y:      C.int(rio.Min.Y),
		width:  C.int(rio.Size().X),
		height: C.int(rio.Size().Y),
	}

	return Mat{p: C.Mat_Region(m.p, cRect)}
}

// Rows returns the number of rows for this Mat.
func (m *Mat) Rows() int {
	return int(C.Mat_Rows(m.p))
}

// Cols returns the number of columns for this Mat.
func (m *Mat) Cols() int {
	return int(C.Mat_Cols(m.p))
}

type MatType int

const (
	MatTypeCV8U  MatType = 0
	MatTypeCV8S          = 1
	MatTypeCV16U         = 2
	MatTypeCV16S         = 3
	MatTypeCV32S         = 4
	MatTypeCV32F         = 5
	MatTypeCV64F         = 6
)

// Scalar is a 4-element vector widely used in OpenCV to pass pixel values.
//
// For further details, please see:
// http://docs.opencv.org/3.3.0/d1/da0/classcv_1_1Scalar__.html
//
type Scalar struct {
	Val1 float64
	Val2 float64
	Val3 float64
	Val4 float64
}

// NewScalar returns a new Scalar. These are usually colors typically being in BGR order.
func NewScalar(v1 float64, v2 float64, v3 float64, v4 float64) Scalar {
	s := Scalar{Val1: v1, Val2: v2, Val3: v3, Val4: v4}
	return s
}

func toByteArray(b []byte) C.struct_ByteArray {
	return C.struct_ByteArray{
		data:   (*C.char)(unsafe.Pointer(&b[0])),
		length: C.int(len(b)),
	}
}

func toGoBytes(b C.struct_ByteArray) []byte {
	return C.GoBytes(unsafe.Pointer(b.data), b.length)
}
