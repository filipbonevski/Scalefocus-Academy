package main

import (
	"io"
	"strings"
	"testing"
)

//TASK2

func TestRead(t *testing.T) {
	buf := new(strings.Builder)
	io.Copy(buf, NewReverseStringReader("avion"))

	ans := buf.String()

	if ans != "noiva" {
		t.Errorf("expected noiva, got %s", ans)
	}
}
