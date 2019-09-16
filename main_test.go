package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	f, err := ioutil.TempFile("", "main-test-")
	if err != nil {
		t.Fatalf("faield to create temp file: %s\n", err)
	}
	defer f.Close()

	os.Stdout = f

	err = Run([]string{"main"})
	if err != nil {
		t.Fatalf("%s\n", err)
	}

	_, err = f.Seek(0, os.SEEK_SET)
	if err != nil {
		t.Fatalf("seeking output file: %s\n", err)
	}

	bs, err := ioutil.ReadAll(f)
	lines := bytes.Split(bs, []byte("\n"))

	a := strings.Fields(string(lines[0]))
	if len(a) < 5 && a[0] != "A" || a[2] != "C" || a[4] != "foo" {
		t.Fatalf("the first line of main output is incorrect: want=%s got=%s\n", "A vX.X.X: C vX.X.X: foo", a)
	}

	b := strings.Fields(string(lines[1]))
	if len(b) < 5 && b[0] != "B" || b[2] != "D" || b[4] != "foo" {
		t.Fatalf("the first line of main output is incorrect: want=%s got=%s\n", "B vX.X.X: D vX.X.X: foo", b)
	}

}
