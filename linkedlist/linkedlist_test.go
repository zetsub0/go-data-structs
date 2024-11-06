package linkedlist

import (
	"testing"
)

func TestLinkedList_Append(t *testing.T) {
	type args[T comparable] struct {
		v T
	}
	tests := []struct {
		name string
		l    LinkedList[int]
		args args[int]
		want int
	}{
		{
			name: "Append to empty list",
			l:    LinkedList[int]{},
			args: args[int]{v: 1},
			want: 1,
		},
		{
			name: "Append to non-empty list",
			l: LinkedList[int]{
				head: &node[int]{
					value: 1,
				},
				length: 1,
			},
			args: args[int]{v: 2},
			want: 2,
		},
		{
			name: "Append multiple values",
			l:    LinkedList[int]{},
			args: args[int]{v: 1},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.l.Append(tt.args.v)
			if tt.l.length != tt.want {
				t.Errorf("LinkedList.Append() = %v, want %v", tt.l.length, tt.want)
			}
		})
	}
}

func TestLinkedList_Prepend(t *testing.T) {
	type args[T comparable] struct {
		v T
	}
	tests := []struct {
		name string
		l    LinkedList[int]
		args args[int]
		want int
	}{
		{
			name: "Append to empty list",
			l:    LinkedList[int]{},
			args: args[int]{v: 1},
			want: 1,
		},
		{
			name: "Append to non-empty list",
			l: LinkedList[int]{
				head: &node[int]{
					value: 1,
				},
				length: 1,
			},
			args: args[int]{v: 2},
			want: 2,
		},
		{
			name: "Append multiple values",
			l:    LinkedList[int]{},
			args: args[int]{v: 1},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.l.Prepend(tt.args.v)
			if tt.l.length != tt.want {
				t.Errorf("LinkedList.Append() = %v, want %v", tt.l.length, tt.want)
			}
		})
	}
}
