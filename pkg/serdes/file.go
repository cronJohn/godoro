package serdes

import (
	"os"
	"path/filepath"
	"reflect"
	"text/template"

	"github.com/bufbuild/protoyaml-go"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"

	"github.com/cronJohn/godoro/util"
)

// An unmarshaller function is any function that take a specific []byte data format
// and unmarshals it into a proto.Message
type UnmarshalFunc func([]byte, proto.Message) error

var unmarshalFuncs = map[string]UnmarshalFunc{
	".json": protojson.Unmarshal,
	".yaml": protoyaml.Unmarshal,
}

type FileHandler[T proto.Message] struct {
	path          string
	unmarshalFunc UnmarshalFunc
}

func NewFileHandler[T proto.Message](path string) DataAccessor[T] {
	return FileHandler[T]{
		path:          path,
		unmarshalFunc: unmarshalFuncs[filepath.Ext(path)],
	}
}

func (fh FileHandler[T]) Read() (T, error) {
	var msg T
	msg = reflect.New(reflect.TypeOf(msg).Elem()).Interface().(T)
	bytes, err := util.GetFileData(fh.path)
	if err != nil {
		log.Error().Msgf("Error getting data: %s", err.Error())
		return msg, err
	}

	if err := fh.unmarshalFunc(bytes, msg); err != nil {
		log.Error().Msgf("Error unmarshalling data: %s", err.Error())
		return msg, err
	}

	return msg, nil
}

func (fh FileHandler[T]) Write(data T, tmpl *template.Template) error {
	file, err := os.Create(fh.path)
	if err != nil {
		log.Error().Msgf("Error opening file: %s", err.Error())
		return err
	}
	defer file.Close()

	return tmpl.Execute(file, data)
}
