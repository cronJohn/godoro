package parser

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/durationpb"

	"github.com/cronJohn/godoro/pkg/dapter/proto"
)

const (
	CSV = iota
	JSON
)

type FileParser struct {
	Parser
	Format byte
	Path   string
}

func NewFileParser(format byte, path string) *FileParser {
	return &FileParser{
		Format: format,
		Path:   path,
	}
}

func (f *FileParser) GetSessions() (*proto.PomoSessions, error) {
	var buf *proto.PomoSessions
	var err error

	switch f.Format {
	case CSV:
		buf, err = f.ReadCSV()
	case JSON:
		buf, err = f.ReadJSON()
	}

	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (f *FileParser) ReadCSV() (*proto.PomoSessions, error) {
	file, err := os.Open(f.Path)
	if err != nil {
		log.Error().Msgf("Error opening file: %s", err.Error())
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1

	readerData, err := reader.ReadAll()
	if err != nil {
		log.Error().Msgf("Error reading file: %s", err.Error())
		return nil, err
	}

	sessions := make([]*proto.PomoSession, len(readerData))

	for i, r := range readerData {
		if len(r) < 3 {
			log.Error().Msgf("Error parsing row: %s", r)
			return nil, err
		}
		id, err := strconv.Atoi(r[0])
		if err != nil {
			log.Error().Msgf("Error parsing id: %s", err.Error())
			return nil, err
		}
		workDuration, err := time.ParseDuration(r[1])
		if err != nil {
			log.Error().Msgf("Error parsing work duration: %s", err.Error())
			return nil, err
		}
		breakDuration, err := time.ParseDuration(r[2])
		if err != nil {
			log.Error().Msgf("Error parsing break duration: %s", err.Error())
			return nil, err
		}
		sessions[i] = &proto.PomoSession{
			Id:            int32(id),
			WorkDuration:  durationpb.New(workDuration),
			BreakDuration: durationpb.New(breakDuration),
			Tags:          r[3:],
		}
	}

	return &proto.PomoSessions{Sessions: sessions}, nil
}

func (f *FileParser) ReadJSON() (*proto.PomoSessions, error) {
	file, err := os.Open(f.Path)
	if err != nil {
		log.Error().Msgf("Error opening file: %s", err.Error())
		return nil, err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		log.Error().Msgf("Error reading file: %s", err.Error())
		return nil, err
	}

	var sessions proto.PomoSessions
	if err := protojson.Unmarshal(bytes, &sessions); err != nil {
		log.Error().Msgf("Error unmarshalling JSON: %s", err.Error())
		return nil, err
	}

	return &sessions, nil
}
