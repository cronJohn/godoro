package serdes

import (
	"testing"

	"google.golang.org/protobuf/proto"

	mp "github.com/cronJohn/godoro/pkg/serdes/proto/compiled"
)

func TestSerdesFileYAML(t *testing.T) {
	serdes := NewSerdes(NewFileHandler[*mp.TestMessage](".test-output/serdes.yaml"))

	err := serdes.Write(&ProtoTestMessage, YAMLTestTemplate)
	if err != nil {
		t.Error(err)
	}

	readData, err := serdes.Read()
	if err != nil {
		t.Error(err)
	}

	if !proto.Equal(readData, &ProtoTestMessage) {
		t.Errorf("Expected %v, got %v", &ProtoTestMessage, readData)
	}
}
