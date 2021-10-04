package main

import (
	"reflect"
	"testing"
)

func Test_fullJustify(t *testing.T) {
	type args struct {
		words    []string
		maxWidth int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "1",
			args: args{
				words:    []string{"This", "is", "an", "example", "of", "text", "justification."},
				maxWidth: 16,
			},
			want: []string{
				"This    is    an",
				"example  of text",
				"justification.  ",
			},
		},
		{
			name: "2",
			args: args{
				words:    []string{"What", "must", "be", "acknowledgment", "shall", "be"},
				maxWidth: 16,
			},
			want: []string{
				"What   must   be",
				"acknowledgment  ",
				"shall be        ",
			},
		},
		{
			name: "3",
			args: args{
				words:    []string{"a"},
				maxWidth: 2,
			},
			want: []string{"a "},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fullJustify(tt.args.words, tt.args.maxWidth); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fullJustify() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_expand(t *testing.T) {
	type args struct {
		words    []string
		index    []int
		maxWidth int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				words:    []string{"This", "is", "an", "example", "of", "text", "justification."},
				index:    []int{0, 2},
				maxWidth: 16,
			},
			want: "This    is    an",
		},
		{
			name: "2",
			args: args{
				words:    []string{"This", "is", "an", "example", "of", "text", "justification."},
				index:    []int{3, 5},
				maxWidth: 16,
			},
			want: "example  of text",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := expand(tt.args.words, tt.args.index, tt.args.maxWidth); got != tt.want {
				t.Errorf("expand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_splitToLines(t *testing.T) {
	type args struct {
		words    []string
		maxWidth int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "1",
			args: args{
				words:    []string{"This", "is", "an", "example", "of", "text", "justification."},
				maxWidth: 16,
			},
			want: [][]int{
				{0, 2},
				{3, 5},
				{6, 6},
			},
		},
		{
			name: "2",
			args: args{
				words:    []string{"What", "must", "be", "acknowledgment", "shall", "be"},
				maxWidth: 16,
			},
			want: [][]int{
				{0, 2},
				{3, 3},
				{4, 5},
			},
		},
		{
			name: "3",
			args: args{
				words:    []string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"},
				maxWidth: 20,
			},
			want: [][]int{
				{0, 3},
				{4, 5},
				{6, 9},
				{10, 13},
				{14, 16},
				{17, 17},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitToLines(tt.args.words, tt.args.maxWidth); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitToLines() = %v, want %v", got, tt.want)
			}
		})
	}
}
