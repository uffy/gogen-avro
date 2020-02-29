// Code generated by github.com/actgardner/gogen-avro. DO NOT EDIT.
/*
 * SOURCE:
 *     union.avsc
 */
package avro

import (
	"io"
	
	"github.com/actgardner/gogen-avro/vm/types"
	"github.com/actgardner/gogen-avro/vm"
	"github.com/actgardner/gogen-avro/compiler"
)

  
type NestedUnionRecord struct { 
	
	
		IntField int32
	

}

func DeserializeNestedUnionRecord(r io.Reader) (t NestedUnionRecord, err error) {
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err == nil {
		err = vm.Eval(r, deser, &t)
	}
	return
}

func DeserializeNestedUnionRecordFromSchema(r io.Reader, schema string) (t NestedUnionRecord, err error) {
	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err == nil {
		err = vm.Eval(r, deser, &t)
	}
	return
}

func writeNestedUnionRecord(r NestedUnionRecord, w io.Writer) error {
	var err error
	
	err = vm.WriteInt( r.IntField, w)
	if err != nil {
		return err			
	}
	
	return err
}

func (r NestedUnionRecord) Serialize(w io.Writer) error {
	return writeNestedUnionRecord(r, w)
}

func (r NestedUnionRecord) Schema() string {
	return "{\"fields\":[{\"name\":\"IntField\",\"type\":\"int\"}],\"name\":\"NestedUnionRecord\",\"type\":\"record\"}"
}

func (r NestedUnionRecord) SchemaName() string {
	return "NestedUnionRecord"
}

func (_ *NestedUnionRecord) SetBoolean(v bool) { panic("Unsupported operation") }
func (_ *NestedUnionRecord) SetInt(v int32) { panic("Unsupported operation") }
func (_ *NestedUnionRecord) SetLong(v int64) { panic("Unsupported operation") }
func (_ *NestedUnionRecord) SetFloat(v float32) { panic("Unsupported operation") }
func (_ *NestedUnionRecord) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *NestedUnionRecord) SetBytes(v []byte) { panic("Unsupported operation") }
func (_ *NestedUnionRecord) SetString(v string) { panic("Unsupported operation") }
func (_ *NestedUnionRecord) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *NestedUnionRecord) Get(i int) types.Field {
	switch (i) {
	
	case 0:
		
		
			return (*types.Int)(&r.IntField)
		
	
	default:
		panic("Unknown field index")
	}
}

func (r *NestedUnionRecord) SetDefault(i int) {
	switch (i) { 
	default:
		panic("Unknown field index")
	}
}

func (r *NestedUnionRecord) Clear(i int) {
	switch (i) { 
	default:
		panic("Non-optional field index")
	}
}

func (_ *NestedUnionRecord) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *NestedUnionRecord) ClearMap(key string) { panic("Unsupported operation") }
func (_ *NestedUnionRecord) AppendArray() types.Field { panic("Unsupported operation") }
func (_ *NestedUnionRecord) Finalize() { }
