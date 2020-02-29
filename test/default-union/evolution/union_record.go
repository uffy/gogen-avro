// Code generated by github.com/actgardner/gogen-avro. DO NOT EDIT.
/*
 * SOURCE:
 *     evolution.avsc
 */
package avro

import (
	"io"
	
	"github.com/actgardner/gogen-avro/vm/types"
	"github.com/actgardner/gogen-avro/vm"
	"github.com/actgardner/gogen-avro/compiler"
)

  
type UnionRecord struct { 
	
	
		Id string
	

	
	
		UnionNull *UnionNullString
	

	
	
		UnionString UnionStringInt
	

	
	
		UnionRecord UnionUnionRecString
	

}

func DeserializeUnionRecord(r io.Reader) (t UnionRecord, err error) {
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err == nil {
		err = vm.Eval(r, deser, &t)
	}
	return
}

func DeserializeUnionRecordFromSchema(r io.Reader, schema string) (t UnionRecord, err error) {
	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err == nil {
		err = vm.Eval(r, deser, &t)
	}
	return
}

func writeUnionRecord(r UnionRecord, w io.Writer) error {
	var err error
	
	err = vm.WriteString( r.Id, w)
	if err != nil {
		return err			
	}
	
	err = writeUnionNullString( r.UnionNull, w)
	if err != nil {
		return err			
	}
	
	err = writeUnionStringInt( r.UnionString, w)
	if err != nil {
		return err			
	}
	
	err = writeUnionUnionRecString( r.UnionRecord, w)
	if err != nil {
		return err			
	}
	
	return err
}

func (r UnionRecord) Serialize(w io.Writer) error {
	return writeUnionRecord(r, w)
}

func (r UnionRecord) Schema() string {
	return "{\"fields\":[{\"default\":\"test_id\",\"name\":\"id\",\"type\":\"string\"},{\"default\":null,\"name\":\"unionNull\",\"type\":[\"null\",\"string\"]},{\"default\":\"hello\",\"name\":\"unionString\",\"type\":[\"string\",\"int\"]},{\"default\":{\"a\":1},\"name\":\"unionRecord\",\"type\":[{\"fields\":[{\"name\":\"a\",\"type\":\"int\"}],\"name\":\"unionRec\",\"type\":\"record\"},\"string\"]}],\"name\":\"UnionRecord\",\"type\":\"record\"}"
}

func (r UnionRecord) SchemaName() string {
	return "UnionRecord"
}

func (_ *UnionRecord) SetBoolean(v bool) { panic("Unsupported operation") }
func (_ *UnionRecord) SetInt(v int32) { panic("Unsupported operation") }
func (_ *UnionRecord) SetLong(v int64) { panic("Unsupported operation") }
func (_ *UnionRecord) SetFloat(v float32) { panic("Unsupported operation") }
func (_ *UnionRecord) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionRecord) SetBytes(v []byte) { panic("Unsupported operation") }
func (_ *UnionRecord) SetString(v string) { panic("Unsupported operation") }
func (_ *UnionRecord) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *UnionRecord) Get(i int) types.Field {
	switch (i) {
	
	case 0:
		
		
			return (*types.String)(&r.Id)
		
	
	case 1:
		
			r.UnionNull = &UnionNullString{}

		
		
			return r.UnionNull
		
	
	case 2:
		
			r.UnionString = UnionStringInt{}

		
		
			return &r.UnionString
		
	
	case 3:
		
			r.UnionRecord = UnionUnionRecString{}

		
		
			return &r.UnionRecord
		
	
	default:
		panic("Unknown field index")
	}
}

func (r *UnionRecord) SetDefault(i int) {
	switch (i) { 
	case 0:
		r.Id = "test_id"
		
	case 1:
		r.UnionNull = nil
		
	case 2:
		r.UnionString = UnionStringInt{}
r.UnionString.String = "hello"
		
	case 3:
		r.UnionRecord = UnionUnionRecString{}
r.UnionRecord.UnionRec = UnionRec{}
r.UnionRecord.UnionRec.A = 1

		
	default:
		panic("Unknown field index")
	}
}

func (r *UnionRecord) Clear(i int) {
	switch (i) { 
	case 1:
		r.UnionNull = nil
	default:
		panic("Non-optional field index")
	}
}

func (_ *UnionRecord) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionRecord) ClearMap(key string) { panic("Unsupported operation") }
func (_ *UnionRecord) AppendArray() types.Field { panic("Unsupported operation") }
func (_ *UnionRecord) Finalize() { }
