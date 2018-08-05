// +build linux,!386
package conntrack

import (
	"reflect"
	"testing"

	"golang.org/x/net/bpf"
)

func TestConstructFilter(t *testing.T) {
	tests := []struct {
		name     string
		table    CtTable
		filters  []ConnAttr
		rawInstr []bpf.RawInstruction
		err      error
	}{
		// Modified example from libnetfilter_conntrack/utils/conntrack_filter.c
		{name: "conntrack_filter.c", table: Ct, filters: []ConnAttr{
			{Type: AttrOrigL4Proto, Data: []byte{0x11}},                                                      // TCP
			{Type: AttrOrigL4Proto, Data: []byte{0x06}},                                                      // UDP
			{Type: AttrTCPState, Data: []byte{0x3}},                                                          // TCP_CONNTRACK_ESTABLISHED
			{Type: AttrOrigIPv4Src, Data: []byte{0x7F, 0x0, 0x0, 0x1}, Mask: []byte{0xff, 0xff, 0xff, 0xff}}, // SrcIP: 127.0.0.1
		}, rawInstr: []bpf.RawInstruction{
			{Op: 1, Jt: 0, Jf: 0, K: 0x00000004},
			{Op: 80, Jt: 0, Jf: 0, K: 0x00000001},
			{Op: 21, Jt: 1, Jf: 0, K: 0x00000001},
			{Op: 6, Jt: 0, Jf: 0, K: 0xffffffff},
			{Op: 0, Jt: 0, Jf: 0, K: 0x0000000e},
			{Op: 1, Jt: 0, Jf: 0, K: 0x00000001},
			{Op: 48, Jt: 0, Jf: 0, K: 0xfffff00c},
			{Op: 21, Jt: 13, Jf: 0, K: 0x00000000},
			{Op: 4, Jt: 0, Jf: 0, K: 0x00000004},
			{Op: 1, Jt: 0, Jf: 0, K: 0x00000002},
			{Op: 48, Jt: 0, Jf: 0, K: 0xfffff00c},
			{Op: 21, Jt: 9, Jf: 0, K: 0x00000000},
			{Op: 4, Jt: 0, Jf: 0, K: 0x00000004},
			{Op: 1, Jt: 0, Jf: 0, K: 0x00000001},
			{Op: 48, Jt: 0, Jf: 0, K: 0xfffff00c},
			{Op: 21, Jt: 5, Jf: 0, K: 0x00000000},
			{Op: 7, Jt: 0, Jf: 0, K: 0x00000000},
			{Op: 80, Jt: 0, Jf: 0, K: 0x00000004},
			{Op: 21, Jt: 2, Jf: 0, K: 0x00000011},
			{Op: 21, Jt: 1, Jf: 0, K: 0x00000006},
			{Op: 6, Jt: 0, Jf: 0, K: 0x00000000},
			{Op: 0, Jt: 0, Jf: 0, K: 0x0000000e},
			{Op: 1, Jt: 0, Jf: 0, K: 0x00000004},
			{Op: 48, Jt: 0, Jf: 0, K: 0xfffff00c},
			{Op: 21, Jt: 12, Jf: 0, K: 0x00000000},
			{Op: 4, Jt: 0, Jf: 0, K: 0x00000004},
			{Op: 1, Jt: 0, Jf: 0, K: 0x00000001},
			{Op: 48, Jt: 0, Jf: 0, K: 0xfffff00c},
			{Op: 21, Jt: 8, Jf: 0, K: 0x00000000},
			{Op: 4, Jt: 0, Jf: 0, K: 0x00000004},
			{Op: 1, Jt: 0, Jf: 0, K: 0x00000001},
			{Op: 48, Jt: 0, Jf: 0, K: 0xfffff00c},
			{Op: 21, Jt: 4, Jf: 0, K: 0x00000000},
			{Op: 7, Jt: 0, Jf: 0, K: 0x00000000},
			{Op: 80, Jt: 0, Jf: 0, K: 0x00000004},
			{Op: 21, Jt: 1, Jf: 0, K: 0x00000003},
			{Op: 6, Jt: 0, Jf: 0, K: 0x00000000},
			{Op: 0, Jt: 0, Jf: 0, K: 0x0000000e},
			{Op: 1, Jt: 0, Jf: 0, K: 0x00000001},
			{Op: 48, Jt: 0, Jf: 0, K: 0xfffff00c},
			{Op: 21, Jt: 15, Jf: 0, K: 0x00000000},
			{Op: 4, Jt: 0, Jf: 0, K: 0x00000004},
			{Op: 1, Jt: 0, Jf: 0, K: 0x00000001},
			{Op: 48, Jt: 0, Jf: 0, K: 0xfffff00c},
			{Op: 21, Jt: 11, Jf: 0, K: 0x00000000},
			{Op: 4, Jt: 0, Jf: 0, K: 0x00000004},
			{Op: 1, Jt: 0, Jf: 0, K: 0x00000001},
			{Op: 48, Jt: 0, Jf: 0, K: 0xfffff00c},
			{Op: 21, Jt: 7, Jf: 0, K: 0x00000000},
			{Op: 7, Jt: 0, Jf: 0, K: 0x00000000},
			{Op: 64, Jt: 0, Jf: 0, K: 0x00000004},
			{Op: 7, Jt: 0, Jf: 0, K: 0x00000000},
			{Op: 84, Jt: 0, Jf: 0, K: 0xffffffff},
			{Op: 21, Jt: 2, Jf: 0, K: 0x7f000001},
			{Op: 135, Jt: 0, Jf: 0, K: 0x00000000},
			{Op: 6, Jt: 0, Jf: 0, K: 0x00000000},
			{Op: 6, Jt: 0, Jf: 0, K: 0xffffffff},
		}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			rawInstr, err := constructFilter(tc.table, tc.filters)

			if err != tc.err {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(rawInstr, tc.rawInstr) {
				t.Fatalf("unexpected replies:\n- want: %#v\n-  got: %#v", tc.rawInstr, rawInstr)
			}

		})
	}
}
