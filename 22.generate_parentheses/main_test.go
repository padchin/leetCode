package main

import (
	"reflect"
	"testing"
)

func Test_generateParenthesisBits(t *testing.T) {
	type args struct {
		size int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "1",
			args: args{
				size: 1,
			},
			want: []string{"()"},
		},
		{
			name: "2",
			args: args{
				size: 2,
			},
			want: []string{"(())", "()()"},
		},
		{
			name: "3",
			args: args{
				size: 3,
			},
			want: []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
		{
			name: "4",
			args: args{
				size: 4,
			},
			want: []string{"(((())))", "((()()))", "((())())", "((()))()", "(()(()))", "(()()())", "(()())()", "(())(())", "(())()()", "()((()))", "()(()())", "()(())()", "()()(())", "()()()()"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateParenthesisBits(tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateParenthesis() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generateParenthesisRecursion(t *testing.T) {
	type args struct {
		size int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "1",
			args: args{
				size: 1,
			},
			want: []string{"()"},
		},
		{
			name: "2",
			args: args{
				size: 2,
			},
			want: []string{"(())", "()()"},
		},
		{
			name: "3",
			args: args{
				size: 3,
			},
			want: []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
		{
			name: "4",
			args: args{
				size: 4,
			},
			want: []string{"(((())))", "((()()))", "((())())", "((()))()", "(()(()))", "(()()())", "(()())()", "(())(())", "(())()()", "()((()))", "()(()())", "()(())()", "()()(())", "()()()()"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateParenthesisRecursion(tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateParenthesisRecursion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_generateParenthesisRecursion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		generateParenthesisRecursion(4)
	}
}

func Benchmark_generateParenthesisBits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		generateParenthesisBits(4)
	}
}
