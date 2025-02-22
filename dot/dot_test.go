package dot_test

import (
	"bytes"
	"context"
	"testing"

	"github.com/cayleygraph/quad"
	"github.com/cayleygraph/quad/dot"
)

var testData = []struct {
	quads []quad.Quad
	data  string
}{
	{
		[]quad.Quad{
			{
				Subject:   quad.BNode("subject1"),
				Predicate: quad.IRI("/film/performance/character"),
				Object:    quad.String("Tomas de Torquemada"),
				Label:     nil,
			},
			{
				Subject:   quad.BNode("subject1"),
				Predicate: quad.IRI("http://an.example/predicate1"),
				Object:    quad.String("object1"),
				Label:     nil,
			},
			{
				Subject:   quad.IRI("http://example.org/bob#me"),
				Predicate: quad.IRI("http://schema.org/birthDate"),
				Object: quad.TypedString{
					Value: "1990-07-04",
					Type:  "http://www.w3.org/2001/XMLSchema#date",
				},
				Label: nil,
			},
			{
				Subject:   quad.BNode("subject1"),
				Predicate: quad.IRI("text"),
				Object:    quad.String(`some "text"`),
				Label:     nil,
			},
		},
		`digraph cayley_graph {
	"_:subject1" -> "\"Tomas de Torquemada\"" [ label = "</film/performance/character>" ];
	"_:subject1" -> "\"object1\"" [ label = "<http://an.example/predicate1>" ];
	"<http://example.org/bob#me>" -> "\"1990-07-04\"^^<http://www.w3.org/2001/XMLSchema#date>" [ label = "<http://schema.org/birthDate>" ];
	"_:subject1" -> "\"some \\\"text\\\"\"" [ label = "<text>" ];
}
`,
	},
}

func TestWriter(t *testing.T) {
	buf := bytes.NewBuffer(nil)
	ctx := context.Background()
	for _, c := range testData {
		buf.Reset()
		w := dot.NewWriter(buf)
		n, err := quad.Copy(ctx, w, quad.NewReader(c.quads))
		if err != nil {
			t.Fatalf("write failed after %d quads: %v", n, err)
		}
		if err = w.Close(); err != nil {
			t.Fatal("error on close:", err)
		}
		if c.data != buf.String() {
			t.Fatalf("wrong output:\n%s\n\nvs\n\n%s", buf.String(), c.data)
		}
	}
}
