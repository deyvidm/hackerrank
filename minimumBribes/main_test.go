package main

import "testing"

func Test_minimumBribes(t *testing.T) {
	type args struct {
		q []int32
	}
	tests := []struct {
		name     string
		args     args
		expected int32
	}{
		{name: "no input", expected: 0},
		{name: "happy path :: 1 2 3", args: args{q: []int32{1, 2, 3}}, expected: 0},
		{name: "single bribe :: 1 3 2", args: args{q: []int32{1, 3, 2}}, expected: 1},
		{name: "two bribes :: 1 3 4 2", args: args{q: []int32{1, 3, 4, 2}}, expected: 2},
		{name: "off by one :: ", args: args{q: []int32{1, 2, 5, 3, 7, 8, 6, 4}}, expected: 7},
		//to chaootic! only 2 bribes per person allowed
		{name: "too many bribes :: 4 1 2 3", args: args{q: []int32{4, 1, 2, 3}}, expected: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minimumBribesTestable(tt.args.q)
			if tt.expected != got {
				t.Errorf("test |%s|\nproducted '%d'\nexpected '%d'\n", tt.name, got, tt.expected)
			}
		})
	}
}
