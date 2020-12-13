package main

import (
	"container/ring"
	"fmt"
	"testing"
)

func TestSingleton(t *testing.T) {
	var r *ring.Ring
	r = ring.New(1)
	r.Value = 0

	l := r.Len()
	if l != 1 {
		t.Errorf("len == %d, want 1", l)
	}

	v := r.Value.(int)
	if v != 0 {
		t.Errorf("r[0].Value == %d, want 0", v)
	}
}

func TestAddAtEndOfSingleton(t *testing.T) {
	var r *ring.Ring
	r = ring.New(1)
	r.Value = 0

	start := r
	
	r = addMarble(1, r)

	l := r.Len()
	if l != 2 {
		t.Errorf("len == %d, want 2", l)
	}

	v := start.Value.(int)
	if v != 0 {
		t.Errorf("start[0].Value == %d, want 0", v)
	}

	start = start.Move(1)
	v = start.Value.(int)
	if v != 1 {
		t.Errorf("start[1].Value == %d, want 1", v)
	}
}

func TestAddAtEndOfDuo(t *testing.T) {
	var r *ring.Ring
	r = ring.New(1)
	r.Value = 0

	start := r
	
	r = addMarble(1, r)
	r = addMarble(2, r)

	c := start
	
	l := c.Len()
	if l != 3 {
		t.Errorf("len == %d, want 3", l)
	}

	v := c.Value.(int)
	if v != 0 {
		t.Errorf("r[0].Value == %d, want 0", v)
	}

	c = c.Move(1)
	v = c.Value.(int)
	if v != 2 {
		t.Errorf("r[1].Value == %d, want 2", v)
	}

	c = c.Move(1)
	v = c.Value.(int)
	if v != 1 {
		t.Errorf("r[1].Value == %d, want 1", v)
	}

	start.Do(func(p interface{}) {
		fmt.Print(p.(int), " ")
	})
	fmt.Println("")
}

func TestAddInMiddleOfTriple(t *testing.T) {
	var r *ring.Ring
	r = ring.New(1)
	r.Value = 0

	start := r
	
	r = addMarble(1, r)
	r = addMarble(2, r)
	r = addMarble(3, r)
	
	c := start
	
	l := c.Len()
	if l != 4 {
		t.Errorf("len == %d, want 4", l)
	}

	v := c.Value.(int)
	if v != 0 {
		t.Errorf("r[0].Value == %d, want 0", v)
	}

	c = c.Move(1)
	v = c.Value.(int)
	if v != 2 {
		t.Errorf("r[1].Value == %d, want 2", v)
	}

	c = c.Move(1)
	v = c.Value.(int)
	if v != 1 {
		t.Errorf("r[1].Value == %d, want 1", v)
	}

	c = c.Move(1)
	v = c.Value.(int)
	if v != 3 {
		t.Errorf("r[1].Value == %d, want 3", v)
	}

	start.Do(func(p interface{}) {
		fmt.Print(p.(int), " ")
	})
	fmt.Println("")
}
