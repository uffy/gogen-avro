// Code generated by github.com/actgardner/gogen-avro. DO NOT EDIT.
/*
 * SOURCE:
 *     union.avsc
 */
package avro

import (
	"io"
	"fmt"

	"github.com/actgardner/gogen-avro/vm"
	"github.com/actgardner/gogen-avro/vm/types"
)


type UnionIntLongFloatDoubleStringBoolNullTypeEnum int

const (
	UnionIntLongFloatDoubleStringBoolNullTypeEnumInt UnionIntLongFloatDoubleStringBoolNullTypeEnum = 0

	UnionIntLongFloatDoubleStringBoolNullTypeEnumLong UnionIntLongFloatDoubleStringBoolNullTypeEnum = 1

	UnionIntLongFloatDoubleStringBoolNullTypeEnumFloat UnionIntLongFloatDoubleStringBoolNullTypeEnum = 2

	UnionIntLongFloatDoubleStringBoolNullTypeEnumDouble UnionIntLongFloatDoubleStringBoolNullTypeEnum = 3

	UnionIntLongFloatDoubleStringBoolNullTypeEnumString UnionIntLongFloatDoubleStringBoolNullTypeEnum = 4

	UnionIntLongFloatDoubleStringBoolNullTypeEnumBool UnionIntLongFloatDoubleStringBoolNullTypeEnum = 5
)

type UnionIntLongFloatDoubleStringBoolNull struct { 
	Int int32

	Long int64

	Float float32

	Double float64

	String string

	Bool bool

	UnionType UnionIntLongFloatDoubleStringBoolNullTypeEnum
}

func writeUnionIntLongFloatDoubleStringBoolNull(r *UnionIntLongFloatDoubleStringBoolNull, w io.Writer) error { 
	if r == nil {
		return vm.WriteLong(int64(6), w)
	} 
	if err := vm.WriteLong(int64(r.UnionType), w); err != nil {
		return err
	}
	switch r.UnionType{ 
	case UnionIntLongFloatDoubleStringBoolNullTypeEnumInt:
		return vm.WriteInt(r.Int, w)
	case UnionIntLongFloatDoubleStringBoolNullTypeEnumLong:
		return vm.WriteLong(r.Long, w)
	case UnionIntLongFloatDoubleStringBoolNullTypeEnumFloat:
		return vm.WriteFloat(r.Float, w)
	case UnionIntLongFloatDoubleStringBoolNullTypeEnumDouble:
		return vm.WriteDouble(r.Double, w)
	case UnionIntLongFloatDoubleStringBoolNullTypeEnumString:
		return vm.WriteString(r.String, w)
	case UnionIntLongFloatDoubleStringBoolNullTypeEnumBool:
		return vm.WriteBool(r.Bool, w)
	}
	return fmt.Errorf("invalid value for *UnionIntLongFloatDoubleStringBoolNull")
}

func (_ *UnionIntLongFloatDoubleStringBoolNull) SetBoolean(v bool) { panic("Unsupported operation") }
func (_ *UnionIntLongFloatDoubleStringBoolNull) SetInt(v int32) { panic("Unsupported operation") }
func (_ *UnionIntLongFloatDoubleStringBoolNull) SetFloat(v float32) { panic("Unsupported operation") }
func (_ *UnionIntLongFloatDoubleStringBoolNull) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionIntLongFloatDoubleStringBoolNull) SetBytes(v []byte) { panic("Unsupported operation") }
func (_ *UnionIntLongFloatDoubleStringBoolNull) SetString(v string) { panic("Unsupported operation") }

func (r *UnionIntLongFloatDoubleStringBoolNull) SetLong(v int64) { 
	r.UnionType = (UnionIntLongFloatDoubleStringBoolNullTypeEnum)(v)
}

func (r *UnionIntLongFloatDoubleStringBoolNull) Get(i int) types.Field {
	switch (i) { 
	case 0:
		
		
		return (*types.Int)(&r.Int)
		
	
	case 1:
		
		
		return (*types.Long)(&r.Long)
		
	
	case 2:
		
		
		return (*types.Float)(&r.Float)
		
	
	case 3:
		
		
		return (*types.Double)(&r.Double)
		
	
	case 4:
		
		
		return (*types.String)(&r.String)
		
	
	case 5:
		
		
		return (*types.Boolean)(&r.Bool)
		
	
	}
	panic("Unknown field index")
}

func (r *UnionIntLongFloatDoubleStringBoolNull) Clear(i int) { panic("Unsupported operation") }
func (_ *UnionIntLongFloatDoubleStringBoolNull) SetDefault(i int) { panic("Unsupported operation") }
func (_ *UnionIntLongFloatDoubleStringBoolNull) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionIntLongFloatDoubleStringBoolNull) ClearMap(key string) { panic("Unsupported operation") }
func (_ *UnionIntLongFloatDoubleStringBoolNull) AppendArray() types.Field { panic("Unsupported operation") }
func (_ *UnionIntLongFloatDoubleStringBoolNull) Finalize()  { }
