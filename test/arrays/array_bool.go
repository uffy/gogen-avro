// Code generated by github.com/actgardner/gogen-avro. DO NOT EDIT.
/*
 * SOURCE:
 *     arrays.avsc
 */
package avro

import (
	"io"

	"github.com/actgardner/gogen-avro/vm/types"
	"github.com/actgardner/gogen-avro/vm"
)

func writeArrayBool(r []bool, w io.Writer) error {
	err := vm.WriteLong(int64(len(r)),w)
	if err != nil || len(r) == 0 {
		return err
	}
	for _, e := range r {
		err = vm.WriteBool(e, w)
		if err != nil {
			return err
		}
	}
	return vm.WriteLong(0,w)
}

type ArrayBool []bool

func (_ *ArrayBool) SetBoolean(v bool) { panic("Unsupported operation") }
func (_ *ArrayBool) SetInt(v int32) { panic("Unsupported operation") }
func (_ *ArrayBool) SetLong(v int64) { panic("Unsupported operation") }
func (_ *ArrayBool) SetFloat(v float32) { panic("Unsupported operation") }
func (_ *ArrayBool) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *ArrayBool) SetBytes(v []byte) { panic("Unsupported operation") }
func (_ *ArrayBool) SetString(v string) { panic("Unsupported operation") }
func (_ *ArrayBool) SetUnionElem(v int64) { panic("Unsupported operation") }
func (_ *ArrayBool) Get(i int) types.Field { panic("Unsupported operation") }
func (_ *ArrayBool) Clear(i int) { panic("Unsupported operation") }
func (_ *ArrayBool) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *ArrayBool) ClearMap(key string) { panic("Unsupported operation") }
func (_ *ArrayBool) Finalize() { }
func (_ *ArrayBool) SetDefault(i int) { panic("Unsupported operation") }

func (r *ArrayBool) AppendArray() types.Field {
	var v bool
	
	*r = append(*r, v)
	
	return (*types.Boolean)(&(*r)[len(*r)-1])
	
}

