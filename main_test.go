package main

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func init() {
	Slow(ioutil.Discard)
	Fast(ioutil.Discard)
}

// -----
// go test -v

func TestSearch(t *testing.T) {
	slowOut := new(bytes.Buffer)
	Slow(slowOut)
	slowResult := slowOut.String()

	fastOut := new(bytes.Buffer)
	Fast(fastOut)
	fastResult := fastOut.String()

	if slowResult != fastResult {
		t.Errorf("results not match\nGot:\n%v\nExpected:\n%v", fastResult, slowResult)
	}
}


// -----
// go test -bench . -benchmem

func BenchmarkSlow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Slow(ioutil.Discard)
	}
}

func BenchmarkFast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fast(ioutil.Discard)
	}
}
