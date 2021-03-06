// Code generated by github.com/actgardner/gogen-avro/v7. DO NOT EDIT.
/*
 * SOURCE:
 *     namespace.avsc
 */
package avro

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v7/vm"
	"github.com/actgardner/gogen-avro/v7/vm/types"
)

type UnionNullBodyworksTraceTypeEnum int

const (
	UnionNullBodyworksTraceTypeEnumBodyworksTrace UnionNullBodyworksTraceTypeEnum = 1
)

type UnionNullBodyworksTrace struct {
	Null           *types.NullVal
	BodyworksTrace *BodyworksTrace
	UnionType      UnionNullBodyworksTraceTypeEnum
}

func writeUnionNullBodyworksTrace(r *UnionNullBodyworksTrace, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullBodyworksTraceTypeEnumBodyworksTrace:
		return writeBodyworksTrace(r.BodyworksTrace, w)
	}
	return fmt.Errorf("invalid value for *UnionNullBodyworksTrace")
}

func NewUnionNullBodyworksTrace() *UnionNullBodyworksTrace {
	return &UnionNullBodyworksTrace{}
}

func (_ *UnionNullBodyworksTrace) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullBodyworksTrace) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullBodyworksTrace) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullBodyworksTrace) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullBodyworksTrace) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullBodyworksTrace) SetString(v string)  { panic("Unsupported operation") }
func (r *UnionNullBodyworksTrace) SetLong(v int64) {
	r.UnionType = (UnionNullBodyworksTraceTypeEnum)(v)
}
func (r *UnionNullBodyworksTrace) Get(i int) types.Field {
	switch i {
	case 0:
		return r.Null
	case 1:
		r.BodyworksTrace = NewBodyworksTrace()
		return r.BodyworksTrace
	}
	panic("Unknown field index")
}
func (_ *UnionNullBodyworksTrace) NullField(i int)                  { panic("Unsupported operation") }
func (_ *UnionNullBodyworksTrace) SetDefault(i int)                 { panic("Unsupported operation") }
func (_ *UnionNullBodyworksTrace) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionNullBodyworksTrace) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *UnionNullBodyworksTrace) Finalize()                        {}

func (r *UnionNullBodyworksTrace) MarshalJSON() ([]byte, error) {
	if r == nil {
		return []byte("null"), nil
	}
	switch r.UnionType {
	case UnionNullBodyworksTraceTypeEnumBodyworksTrace:
		return json.Marshal(map[string]interface{}{"Trace": r.BodyworksTrace})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullBodyworksTrace")
}

func (r *UnionNullBodyworksTrace) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if value, ok := fields["Trace"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.BodyworksTrace)
	}
	return fmt.Errorf("invalid value for *UnionNullBodyworksTrace")
}
