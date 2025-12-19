package main

import "testing"

func TestAddTwoPosNum(t *testing.T) {
	n1 := 5
	n2 := 10
	n3 := -10
	n4 := 0

	r1, e1 := add_two_pos_nums(n1, n2)
	if e1 != nil || r1 != (n1+n2) {
		t.Errorf("unexpected result %d (expect pass): %d | %d\nerr: %v", r1, n1, n2, e1)
	}

	r2, e2 := add_two_pos_nums(n1, n3)
	if e2 == nil || r2 != 0 {
		t.Errorf("unexpected result %d (expect fail): %d | %d\nerr: %v", r2, n1, n3, e2)
	}

	r3, e3 := add_two_pos_nums(n1, n4)
	if e3 == nil || r3 != 0 {
		t.Errorf("unexpected result %d (expect fail): %d | %d\nerr: %v", r3, n1, n3, e3)
	}
}
