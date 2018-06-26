package test

import (
	"bytes"
	"testing"

	// internal
	"github.com/sniperkit/snk.golang.json/pkg/json/v1"
)

func Benchmark_encode_string_with_SetEscapeHTML(b *testing.B) {
	type V struct {
		S string
		B bool
		I int
	}
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		buf := &bytes.Buffer{}
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(true)
		if err := enc.Encode(V{S: "s", B: true, I: 233}); err != nil {
			b.Fatal(err)
		}
	}
}