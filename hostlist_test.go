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
	{"n01p[001-005]", []string{"n01p001", "n01p002", "n01p003", "n01p004", "n01p005"}},
	{"ap[1-3]z", []string{"ap1z", "ap2z", "ap3z"}},
	{"dgx[01-03]", []string{"dgx01", "dgx02", "dgx03"}},
	{"adev[6,13,15]", []string{"adev6", "adev13", "adev15"}},
	{"lx[62-64,128]", []string{"lx62", "lx63", "lx64", "lx128"}},
}

var testPairsCPU = []testPair{
	{"36-45,84-85", []string{"36", "37", "38", "39", "40", "41", "42", "43", "44", "45", "84", "85"}},
}

func TestExpandNodeList(t *testing.T) {
	for _, pair := range testPairsNode {
		resultHostlist := ExpandNodeList(pair.input)
		t.Log("input:", pair.input)
		t.Log("output:", resultHostlist)
		if !reflect.DeepEqual(resultHostlist, pair.output) {
			t.Errorf("%#v not equal to %#v", resultHostlist, pair.output)
		}
	}
}

func TestExpandCPUList(t *testing.T) {
	for _, pair := range testPairsCPU {
		resultHostlist := ExpandCPUList(pair.input)
		t.Log(pair.input)
		t.Log(resultHostlist)
		if !reflect.DeepEqual(resultHostlist, pair.output) {
			t.Errorf("%#v not equal to %#v", resultHostlist, pair.output)
		}
	}
}
