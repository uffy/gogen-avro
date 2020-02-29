// Code generated by github.com/actgardner/gogen-avro. DO NOT EDIT.
/*
 * SOURCE:
 *     namespace.avsc
 */
package avro

import (
	"io"
	
	"github.com/actgardner/gogen-avro/vm/types"
	"github.com/actgardner/gogen-avro/vm"
	"github.com/actgardner/gogen-avro/compiler"
)

// Trace  
type BodyworksTrace struct { 
	// Trace Identifier
	
		TraceId *UnionNullDatatypeUUID
	

}

func DeserializeBodyworksTrace(r io.Reader) (t BodyworksTrace, err error) {
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err == nil {
		err = vm.Eval(r, deser, &t)
	}
	return
}

func DeserializeBodyworksTraceFromSchema(r io.Reader, schema string) (t BodyworksTrace, err error) {
	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err == nil {
		err = vm.Eval(r, deser, &t)
	}
	return
}

func writeBodyworksTrace(r BodyworksTrace, w io.Writer) error {
	var err error
	
	err = writeUnionNullDatatypeUUID( r.TraceId, w)
	if err != nil {
		return err			
	}
	
	return err
}

func (r BodyworksTrace) Serialize(w io.Writer) error {
	return writeBodyworksTrace(r, w)
}

func (r BodyworksTrace) Schema() string {
	return "{\"doc\":\"Trace\",\"fields\":[{\"default\":null,\"doc\":\"Trace Identifier\",\"name\":\"traceId\",\"type\":[\"null\",{\"doc\":\"A Universally Unique Identifier, in canonical form in lowercase. Example: de305d54-75b4-431b-adb2-eb6b9e546014\",\"fields\":[{\"default\":\"\",\"name\":\"uuid\",\"type\":\"string\"}],\"name\":\"UUID\",\"namespace\":\"headerworks.datatype\",\"type\":\"record\"}]}],\"name\":\"bodyworks.Trace\",\"type\":\"record\"}"
}

func (r BodyworksTrace) SchemaName() string {
	return "bodyworks.Trace"
}

func (_ *BodyworksTrace) SetBoolean(v bool) { panic("Unsupported operation") }
func (_ *BodyworksTrace) SetInt(v int32) { panic("Unsupported operation") }
func (_ *BodyworksTrace) SetLong(v int64) { panic("Unsupported operation") }
func (_ *BodyworksTrace) SetFloat(v float32) { panic("Unsupported operation") }
func (_ *BodyworksTrace) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *BodyworksTrace) SetBytes(v []byte) { panic("Unsupported operation") }
func (_ *BodyworksTrace) SetString(v string) { panic("Unsupported operation") }
func (_ *BodyworksTrace) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *BodyworksTrace) Get(i int) types.Field {
	switch (i) {
	
	case 0:
		
			r.TraceId = &UnionNullDatatypeUUID{}

		
		
			return r.TraceId
		
	
	default:
		panic("Unknown field index")
	}
}

func (r *BodyworksTrace) SetDefault(i int) {
	switch (i) { 
	case 0:
		r.TraceId = nil
		
	default:
		panic("Unknown field index")
	}
}

func (r *BodyworksTrace) Clear(i int) {
	switch (i) { 
	case 0:
		r.TraceId = nil
	default:
		panic("Non-optional field index")
	}
}

func (_ *BodyworksTrace) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *BodyworksTrace) ClearMap(key string) { panic("Unsupported operation") }
func (_ *BodyworksTrace) AppendArray() types.Field { panic("Unsupported operation") }
func (_ *BodyworksTrace) Finalize() { }
