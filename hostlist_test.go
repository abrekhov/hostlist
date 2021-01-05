package hostlist

import (
	"reflect"
	"testing"
)

type testPair struct {
	input  string
	output []string
}

var testPairsNode = []testPair{
	{"n01p[001-005], ap[1-3]z", []string{"n01p001", "n01p002", "n01p003", "n01p004", "n01p005", "ap1z", "ap2z", "ap3z"}},
}

var testPairsCPU = []testPair{
	{"36-45,84-85", []string{"36", "37", "38", "39", "40", "41", "42", "43", "44", "45", "84", "85"}},
}

func TestExpandNodeList(t *testing.T) {
	for _, pair := range testPairsNode {
		resultHostlist := ExpandNodeList(pair.input)
		if !reflect.DeepEqual(resultHostlist, pair.output) {
			t.Errorf("%#v not equal to %#v", resultHostlist, pair.output)
		}
	}
}

func TestExpandCpuList(t *testing.T) {
	for _, pair := range testPairsCPU {
		resultHostlist := ExpandCpuList(pair.input)
		t.Log(pair.input)
		t.Log(resultHostlist)
		if !reflect.DeepEqual(resultHostlist, pair.output) {
			t.Errorf("%#v not equal to %#v", resultHostlist, pair.output)
		}
	}
}
