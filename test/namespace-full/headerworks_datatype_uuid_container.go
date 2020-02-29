// Code generated by github.com/actgardner/gogen-avro. DO NOT EDIT.
/*
 * SOURCE:
 *     namespace.avsc
 */
package avro

import (
	"io"

	"github.com/actgardner/gogen-avro/container"
	"github.com/actgardner/gogen-avro/vm"
	"github.com/actgardner/gogen-avro/compiler"
)

func NewHeaderworksDatatypeUUIDWriter(writer io.Writer, codec container.Codec, recordsPerBlock int64) (*container.Writer, error) {
	t := HeaderworksDatatypeUUID{}
	return container.NewWriter(writer, codec, recordsPerBlock, t.Schema())
}

// container reader
type HeaderworksDatatypeUUIDReader struct {
	r io.Reader
	p *vm.Program
}

func NewHeaderworksDatatypeUUIDReader(r io.Reader) (*HeaderworksDatatypeUUIDReader, error){
	containerReader, err := container.NewReader(r)
	if err != nil {
		return nil, err
	}

	t := HeaderworksDatatypeUUID{}
	deser, err := compiler.CompileSchemaBytes([]byte(containerReader.AvroContainerSchema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	return &HeaderworksDatatypeUUIDReader {
		r: containerReader,
		p: deser,
	}, nil
}

func (r HeaderworksDatatypeUUIDReader) Read() (t HeaderworksDatatypeUUID, err error) {
	err = vm.Eval(r.r, r.p, &t)
	return
}
