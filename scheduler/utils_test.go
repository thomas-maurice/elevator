package scheduler

import "testing"

func TestComputePathCost(t *testing.T) {
	cases := []struct {
		start int64
		path  []int64
		cost  int64
		end   int64
	}{
		{0, []int64{1}, 2, 1},
		{1, []int64{1, 2, 3}, 5, 3},
		{10, []int64{9, 5, 0}, 13, 0},
	}
	for _, c := range cases {
		cost, end := ComputePathCost(c.start, c.path)
		if cost != c.cost {
			t.Errorf("Invalid cost, got %d, wanted %d", cost, c.cost)
		}
		if end != c.end {
			t.Errorf("Invalid end, got %d, wanted %d", end, c.end)
		}
	}
}

func TestOptimizePath(t *testing.T) {
	cases := []struct {
		in  []int64
		out []int64
	}{
		{[]int64{1, 1}, []int64{1}},
		{[]int64{1, 1, 2, 2, 3}, []int64{1, 2, 3}},
		{[]int64{1, 3, 3}, []int64{1, 3}},
		{[]int64{1, 2, 1, 2}, []int64{1, 2, 1, 2}},
	}
	for _, c := range cases {
		path := OptimizePath(c.in)

		if c.out == nil && path == nil {
			continue
		}

		if c.out == nil || path == nil {
			t.Errorf("Invalid end, got %d, wanted %d", path, c.out)
		}

		if len(c.out) != len(path) {
			t.Errorf("Invalid end, got %d, wanted %d", path, c.out)
		}

		for i := range c.out {
			if c.out[i] != path[i] {
				t.Errorf("Invalid end, got %d, wanted %d", path, c.out)
			}
		}
	}
}

func TestInsertASAP(t *testing.T) {
	cases := []struct {
		in  []int64
		req PickUpRequest
		out []int64
	}{
		{[]int64{1, 3}, PickUpRequest{Source: 2, Destination: 3}, []int64{1, 2, 3}},
		{[]int64{1, 5}, PickUpRequest{Source: 2, Destination: 3}, []int64{1, 2, 3, 5}},
		{[]int64{1, 2}, PickUpRequest{Source: 5, Destination: 3}, []int64{1, 2, 5, 3}},
		{[]int64{10, 3}, PickUpRequest{Source: 7, Destination: 5}, []int64{10, 7, 5, 3}},
		{[]int64{1, 5}, PickUpRequest{Source: 10, Destination: 8}, []int64{1, 5, 10, 8}},
	}
	for _, c := range cases {
		path := InsertASAP(c.in, c.req)

		if c.out == nil && path == nil {
			continue
		}

		if c.out == nil || path == nil {
			t.Errorf("Invalid end, got %d, wanted %d", path, c.out)
		}

		if len(c.out) != len(path) {
			t.Errorf("Invalid end, got %d, wanted %d", path, c.out)
		}

		for i := range c.out {
			if c.out[i] != path[i] {
				t.Errorf("Invalid end, got %d, wanted %d", path, c.out)
			}
		}
	}
}
