package ds

import (
	"reflect"
	"testing"
)

func TestNewMaxPQ(t *testing.T) {
	type args struct {
		maxN int
	}
	tests := []struct {
		name string
		args args
		want *MaxPQ
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMaxPQ(tt.args.maxN); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMaxPQ() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxPQ_IsEmpty(t *testing.T) {
	type fields struct {
		pq   []string
		size int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &MaxPQ{
				pq:   tt.fields.pq,
				size: tt.fields.size,
			}
			if got := p.IsEmpty(); got != tt.want {
				t.Errorf("MaxPQ.IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxPQ_GetSize(t *testing.T) {
	type fields struct {
		pq   []string
		size int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &MaxPQ{
				pq:   tt.fields.pq,
				size: tt.fields.size,
			}
			if got := p.GetSize(); got != tt.want {
				t.Errorf("MaxPQ.GetSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxPQ_GetPQ(t *testing.T) {
	type fields struct {
		pq   []string
		size int
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &MaxPQ{
				pq:   tt.fields.pq,
				size: tt.fields.size,
			}
			if got := p.GetPQ(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MaxPQ.GetPQ() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxPQ_Insert(t *testing.T) {
	type fields struct {
		pq   []string
		size int
	}
	type args struct {
		v string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &MaxPQ{
				pq:   tt.fields.pq,
				size: tt.fields.size,
			}
			p.Insert(tt.args.v)
		})
	}
}

func TestMaxPQ_DeleteMax(t *testing.T) {
	type fields struct {
		pq   []string
		size int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &MaxPQ{
				pq:   tt.fields.pq,
				size: tt.fields.size,
			}
			if got := p.DeleteMax(); got != tt.want {
				t.Errorf("MaxPQ.DeleteMax() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxPQ_less(t *testing.T) {
	type fields struct {
		pq   []string
		size int
	}
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &MaxPQ{
				pq:   tt.fields.pq,
				size: tt.fields.size,
			}
			if got := p.less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("MaxPQ.less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxPQ_exchange(t *testing.T) {
	type fields struct {
		pq   []string
		size int
	}
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &MaxPQ{
				pq:   tt.fields.pq,
				size: tt.fields.size,
			}
			p.exchange(tt.args.i, tt.args.j)
		})
	}
}

func TestMaxPQ_swim(t *testing.T) {
	type fields struct {
		pq   []string
		size int
	}
	type args struct {
		k int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &MaxPQ{
				pq:   tt.fields.pq,
				size: tt.fields.size,
			}
			p.swim(tt.args.k)
		})
	}
}

func TestMaxPQ_sink(t *testing.T) {
	type fields struct {
		pq   []string
		size int
	}
	type args struct {
		k int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &MaxPQ{
				pq:   tt.fields.pq,
				size: tt.fields.size,
			}
			p.sink(tt.args.k)
		})
	}
}
