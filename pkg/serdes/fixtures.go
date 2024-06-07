package serdes

import (
	"text/template"

	mp "github.com/cronJohn/godoro/pkg/serdes/proto/compiled"
)

var ProtoTestMessage = mp.TestMessage{
	Name:  "test",
	Value: 1,
}

var JSONTestTemplate = template.Must(template.New("JSONTmpl").Parse(`{
  "name": "{{ .Name }}",
  "value": {{ .Value }}
}
`))

var YAMLTestTemplate = template.Must(template.New("YAMLTmpl").Parse(`name: {{ .Name }}
value: {{ .Value }}
`))
