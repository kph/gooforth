package gooforth

import (
	"strconv"
)

type TypeBool struct {
	V bool
}

type TypeString struct {
	V string
}

type TypeInt struct {
	V int
}

type TypeInt8 struct {
	V int8
}

type TypeInt16 struct {
	V int16
}

type TypeInt32 struct {
	V int32
}

type TypeInt64 struct {
	V int64
}

type TypeUint struct {
	V uint
}

type TypeUint8 struct {
	V uint8
}

type TypeUint16 struct {
	V uint16
}

type TypeUint32 struct {
	V uint32
}

type TypeUint64 struct {
	V uint64
}

type TypeByte struct {
	V byte
}

type TypeRune struct {
	V rune
}

type TypeFloat32 struct {
	V float32
}

type TypeFloat64 struct {
	V float64
}

type TypeComplex64 struct {
	V complex64
}

type TypeComplex128 struct {
	V complex128
}

func (w *TypeBool) Execute(t Task) (n Word) {
	t.Push(w)
	return nil
}

func (w *TypeBool) String() (s string) {
	if w.V {
		return "true"
	}
	return "false"
}

func (w *TypeString) Execute(t Task) (n Word) {
	t.Push(w)
	return nil
}

func (w *TypeString) String() (s string) {
	return w.V
}

func (w *TypeInt) Execute(t Task) (n Word) {
	t.Push(w)
	return nil
}

func (w *TypeInt) String() (s string) {
	return strconv.Itoa(w.V)
}

func (w *TypeInt8) Execute(t Task) (n Word) {
	t.Push(w)
	return nil
}

func (w *TypeInt8) String() (s string) {
	return strconv.FormatInt(int64(w.V), 10)
}

func (w *TypeInt16) Execute(t Task) (n Word) {
	t.Push(w)
	return nil
}

func (w *TypeInt16) String() (s string) {
	return strconv.FormatInt(int64(w.V), 10)
}

func (w *TypeInt32) Execute(t Task) (n Word) {
	t.Push(w)
	return nil
}

func (w *TypeInt32) String() (s string) {
	return strconv.FormatInt(int64(w.V), 10)
}

func (w *TypeInt64) Execute(t Task) (n Word) {
	t.Push(w)
	return nil
}

func (w *TypeInt64) String() (s string) {
	return strconv.FormatInt(w.V, 10)
}

func (w *TypeUint) Execute(t Task) (n Word) {
	t.Push(w)
	return nil
}

func (w *TypeUint) String() (s string) {
	return strconv.FormatUint(uint64(w.V), 10)
}

func (w *TypeUint8) Execute(t Task) (n Word) {
	t.Push(w)
	return nil
}

func (w *TypeUint8) String() (s string) {
	return strconv.FormatUint(uint64(w.V), 10)
}

func (w *TypeUint16) Execute(t Task) (n Word) {
	t.Push(w)
	return nil
}

func (w *TypeUint16) String() (s string) {
	return strconv.FormatUint(uint64(w.V), 10)
}

func (w *TypeUint32) Execute(t Task) (n Word) {
	t.Push(w)
	return nil
}

func (w *TypeUint32) String() (s string) {
	return strconv.FormatUint(uint64(w.V), 10)
}

func (w *TypeUint64) Execute(t Task) (n Word) {
	t.Push(w)
	return nil
}

func (w *TypeUint64) String() (s string) {
	return strconv.FormatUint(w.V, 10)
}

func (w *TypeByte) Execute(t Task) (n Word) {
	t.Push(w)
	return nil
}

func (w *TypeByte) String() (s string) {
	return string(w.V)
}

func (w *TypeRune) Execute(t Task) (n Word) {
	t.Push(w)
	return nil
}

func (w *TypeRune) String() (s string) {
	return string(w.V)
}

func (w *TypeFloat32) Execute(t Task) (n Word) {
	t.Push(w)
	return nil
}

func (w *TypeFloat32) String() (s string) {
	return strconv.FormatFloat(float64(w.V), 'G', -1, 32)
}

func (w *TypeFloat64) Execute(t Task) (n Word) {
	t.Push(w)
	return nil
}

func (w *TypeFloat64) String() (s string) {
	return strconv.FormatFloat(w.V, 'G', -1, 64)
}

func (w *TypeComplex64) Execute(t Task) (n Word) {
	t.Push(w)
	return nil
}

func (w *TypeComplex64) String() (s string) {
	return strconv.FormatComplex(complex128(w.V), 'G', -1, 64)
}

func (w *TypeComplex128) Execute(t Task) (n Word) {
	t.Push(w)
	return nil
}

func (w *TypeComplex128) String() (s string) {
	return strconv.FormatComplex(w.V, 'G', -1, 128)
}

var TypeIntOne = TypeInt{V: 1}
