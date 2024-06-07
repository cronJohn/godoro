package serdes

import (
	"testing"

	"google.golang.org/protobuf/proto"

	mp "github.com/cronJohn/godoro/pkg/serdes/proto/compiled"
)

func TestJSONParser(t *testing.T) {
	handler := NewFileHandler[*mp.TestMessage](".test-output/test.json")
	err := handler.Write(&ProtoTestMessage, JSONTestTemplate)
	if err != nil {
		t.Error(err)
	}

	readData, err := handler.Read()
	if err != nil {
		t.Error(err)
	}

	if !proto.Equal(readData, &ProtoTestMessage) {
		t.Errorf("Expected %v, got %v", &ProtoTestMessage, readData)
	}
}

func TestYAMLParser(t *testing.T) {
	handler := NewFileHandler[*mp.TestMessage](".test-output/test.yaml")
	err := handler.Write(&ProtoTestMessage, YAMLTestTemplate)
	if err != nil {
		t.Error(err)
	}

	readData, err := handler.Read()
	if err != nil {
		t.Error(err)
	}

	if !proto.Equal(readData, &ProtoTestMessage) {
		t.Errorf("Expected %v, got %v", &ProtoTestMessage, readData)
	}
}
