package serdes

import (
	"text/template"

	"google.golang.org/protobuf/proto"
)

type DataReader[T proto.Message] interface {
	Read() (T, error)
}
type DataWriter[T proto.Message] interface {
	Write(T, *template.Template) error
}

type DataAccessor[T proto.Message] interface {
	DataReader[T]
	DataWriter[T]
}

type Serdes[T proto.Message] struct {
	DataAccessor DataAccessor[T]
}

func NewSerdes[T proto.Message](accessor DataAccessor[T]) *Serdes[T] {
	return &Serdes[T]{
		DataAccessor: accessor,
	}
}

func (s *Serdes[T]) Read() (T, error) {
	return s.DataAccessor.Read()
}

func (s *Serdes[T]) Write(data T, tmpl *template.Template) error {
	return s.DataAccessor.Write(data, tmpl)
}
