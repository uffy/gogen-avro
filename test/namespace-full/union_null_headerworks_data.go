// Code generated by github.com/actgardner/gogen-avro. DO NOT EDIT.
/*
 * SOURCE:
 *     namespace.avsc
 */
package avro

import (
	"io"
	"fmt"

	"github.com/actgardner/gogen-avro/vm"
	"github.com/actgardner/gogen-avro/vm/types"
)


type UnionNullHeaderworksDataTypeEnum int

const (
	UnionNullHeaderworksDataTypeEnumHeaderworksData UnionNullHeaderworksDataTypeEnum = 1
)

type UnionNullHeaderworksData struct { 
	HeaderworksData HeaderworksData

	UnionType UnionNullHeaderworksDataTypeEnum
}

func writeUnionNullHeaderworksData(r *UnionNullHeaderworksData, w io.Writer) error { 
	if r == nil {
		return vm.WriteLong(int64(0), w)
	} 
	if err := vm.WriteLong(int64(r.UnionType), w); err != nil {
		return err
	}
	switch r.UnionType{ 
	case UnionNullHeaderworksDataTypeEnumHeaderworksData:
		return writeHeaderworksData(r.HeaderworksData, w)
	}
	return fmt.Errorf("invalid value for *UnionNullHeaderworksData")
}

func (_ *UnionNullHeaderworksData) SetBoolean(v bool) { panic("Unsupported operation") }
func (_ *UnionNullHeaderworksData) SetInt(v int32) { panic("Unsupported operation") }
func (_ *UnionNullHeaderworksData) SetFloat(v float32) { panic("Unsupported operation") }
func (_ *UnionNullHeaderworksData) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullHeaderworksData) SetBytes(v []byte) { panic("Unsupported operation") }
func (_ *UnionNullHeaderworksData) SetString(v string) { panic("Unsupported operation") }

func (r *UnionNullHeaderworksData) SetLong(v int64) { 
	r.UnionType = (UnionNullHeaderworksDataTypeEnum)(v)
}

func (r *UnionNullHeaderworksData) Get(i int) types.Field {
	switch (i) { 
	case 1:
		
		r.HeaderworksData = HeaderworksData{}
		
		
		return &r.HeaderworksData
		
	
	}
	panic("Unknown field index")
}

func (r *UnionNullHeaderworksData) Clear(i int) { panic("Unsupported operation") }
func (_ *UnionNullHeaderworksData) SetDefault(i int) { panic("Unsupported operation") }
func (_ *UnionNullHeaderworksData) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionNullHeaderworksData) ClearMap(key string) { panic("Unsupported operation") }
func (_ *UnionNullHeaderworksData) AppendArray() types.Field { panic("Unsupported operation") }
func (_ *UnionNullHeaderworksData) Finalize()  { }
