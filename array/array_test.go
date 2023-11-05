package array

import (
	"reflect"
	"testing"
)

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "first",
			args: args{
				nums: []int{0, 0, 1, 2, 3, 3, 4, 4},
			},
			want: 5,
		},
		{
			name: "second",
			args: args{
				nums: []int{0, 0, 1, 1, 2, 3, 4},
			},
			want: 5,
		},
		{
			name: "third",
			args: args{
				nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.args.nums); got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxProfit(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "first",
			args: args{
				prices: []int{7, 1, 5, 3, 6, 4},
			},
			want: 7,
		},
		{
			name: "second",
			args: args{
				prices: []int{1, 2, 3, 4, 5},
			},
			want: 4,
		},
		{
			name: "third",
			args: args{
				prices: []int{7, 6, 4, 3, 1},
			},
			want: 0,
		},
		{
			name: "forth",
			args: args{
				prices: []int{6, 1, 3, 2, 4, 7},
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rotate(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name   string
		args   args
		expect []int
	}{
		{
			name: "first",
			args: args{
				nums: []int{1, 2, 3, 4, 5},
				k:    3,
			},
			expect: []int{3, 4, 5, 1, 2},
		},
		{
			name: "second",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 6, 7},
				k:    3,
			},
			expect: []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			name: "third",
			args: args{
				nums: []int{-1, -100, 3, 99},
				k:    2,
			},
			expect: []int{3, 99, -1, -100},
		},
		{
			name: "forth",
			args: args{
				nums: []int{-1},
				k:    2,
			},
			expect: []int{-1},
		},
		{
			name: "fifth",
			args: args{
				nums: []int{1, 2},
				k:    3,
			},
			expect: []int{2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate(tt.args.nums, tt.args.k)
			if !reflect.DeepEqual(tt.args.nums, tt.expect) {
				t.Errorf("rotate() = %v, want %v", tt.args.nums, tt.expect)
				return
			}
		})

	}
}

func Test_containsDuplicate(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "first",
			args: args{
				nums: []int{1, 2, 3, 1},
			},
			want: true,
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			if got := containsDuplicate(tt.args.nums); got != tt.want {
				t.Errorf("containsDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
