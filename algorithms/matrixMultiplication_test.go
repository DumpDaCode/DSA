package algorithms

import (
	"reflect"
	"testing"
)

func TestMatrix_NaiveMultiply(t *testing.T) {
	type args struct {
		B Matrix
	}
	tests := []struct {
		name string
		A    Matrix
		args args
		want Matrix
	}{
		// TODO: Add test cases.
		{
			name: "2X2 with positve Values Only",
			A: Matrix{
				{1, 2},
				{2, 1},
			},
			args: args{
				B: Matrix{
					{1, 0},
					{2, 1},
				},
			},
			want: Matrix{
				{5, 2},
				{4, 1},
			},
		},
		{
			name: "2X2 With negative values",
			A: Matrix{
				{1, 2},
				{2, -1},
			},
			args: args{
				B: Matrix{
					{-1, 0},
					{2, 1},
				},
			},
			want: Matrix{
				{3, 2},
				{-4, -1},
			},
		},
		{
			name: "4X4 Matrices",
			A: Matrix{
				{5, 2, 6, 1},
				{0, 6, 2, 0},
				{3, 8, 1, 4},
				{1, 8, 5, 6},
			},
			args: args{
				B: Matrix{
					{7, 5, 8, 0},
					{1, 8, 2, 6},
					{9, 4, 3, 8},
					{5, 3, 7, 9},
				},
			},
			want: Matrix{
				{96, 68, 69, 69},
				{24, 56, 18, 52},
				{58, 95, 71, 92},
				{90, 107, 81, 142},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.A.NaiveMultiply(tt.args.B); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Matrix.NaiveMultiply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix_Add(t *testing.T) {
	type args struct {
		B Matrix
	}
	tests := []struct {
		name string
		A    Matrix
		args args
		want Matrix
	}{
		// TODO: Add test cases.
		{
			name: "Test#1",
			A: Matrix{
				{1, 2},
				{2, 1},
			},
			args: args{
				B: Matrix{
					{1, 0},
					{2, 1},
				},
			},
			want: Matrix{
				{2, 2},
				{4, 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.A.Add(tt.args.B); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Matrix.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix_RecursiveMultiply(t *testing.T) {
	type args struct {
		B Matrix
	}
	tests := []struct {
		name string
		A    Matrix
		args args
		want Matrix
	}{
		// TODO: Add test cases.
		{
			name: "2X2 with positve Values Only",
			A: Matrix{
				{1, 2},
				{2, 1},
			},
			args: args{
				B: Matrix{
					{1, 0},
					{2, 1},
				},
			},
			want: Matrix{
				{5, 2},
				{4, 1},
			},
		},
		{
			name: "2X2 With negative values",
			A: Matrix{
				{1, 2},
				{2, -1},
			},
			args: args{
				B: Matrix{
					{-1, 0},
					{2, 1},
				},
			},
			want: Matrix{
				{3, 2},
				{-4, -1},
			},
		},
		{
			name: "4X4 Matrices",
			A: Matrix{
				{5, 2, 6, 1},
				{0, 6, 2, 0},
				{3, 8, 1, 4},
				{1, 8, 5, 6},
			},
			args: args{
				B: Matrix{
					{7, 5, 8, 0},
					{1, 8, 2, 6},
					{9, 4, 3, 8},
					{5, 3, 7, 9},
				},
			},
			want: Matrix{
				{96, 68, 69, 69},
				{24, 56, 18, 52},
				{58, 95, 71, 92},
				{90, 107, 81, 142},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.A.RecursiveMultiply(tt.args.B); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Matrix.RecursiveMultiply() = %v, want %v", got, tt.want)
			}
		})
	}
}
