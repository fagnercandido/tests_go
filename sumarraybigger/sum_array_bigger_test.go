package main

import "testing"

func TestSumBigger(t *testing.T) {

	var alice []int32
	alice = append(alice, 1)
	alice = append(alice, 2)
	alice = append(alice, 3)

	var bob []int32
	bob = append(bob, 3)
	bob = append(bob, 4)
	bob = append(bob, 1)

	test1 := SumBigger(alice, bob)

	if len(test1) != 2 {
		t.Errorf("The result was be 2 %d", test1)
	}

	if test1[0] != 1 || test1[1] != 2 {
		t.Errorf("The incorrect result %d", test1)
	}

}
