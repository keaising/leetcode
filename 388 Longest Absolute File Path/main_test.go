package main

import (
	"testing"
)

func Test_helper(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 int
		want2 string
	}{
		{
			name: "1",
			args: args{
				str: "\t\t\tfile.ext",
			},
			want:  true,
			want1: 3,
			want2: "file.ext",
		},
		{
			name: "2",
			args: args{
				str: "\t\t\tfile",
			},
			want:  false,
			want1: 3,
			want2: "file",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := helper(tt.args.str)
			if got != tt.want {
				t.Errorf("helper() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("helper() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("helper() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func Test_lengthLongestPath(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				input: "dir\n\tsub1\n\t\tsub2\n\t\t\tfile.ext",
			},
			want: 22,
		},
		{
			name: "2",
			args: args{
				input: "file1.txt\nfile2.txt\nlongfile.txt",
			},
			want: 12,
		},
		{
			name: "3",
			args: args{
				input: "a",
			},
			want: 0,
		},
		{
			name: "3",
			args: args{
				input: "dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext",
			},
			want: 32,
		},
		{
			name: "4",
			args: args{
				input: "a\n\tb\n\t\tc.txt\n\taaaa.txt",
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthLongestPath(tt.args.input); got != tt.want {
				t.Errorf("lengthLongestPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
