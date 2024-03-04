package main

import "testing"

func TestAdd(t *testing.T) {

	expectedResult := 5
	x := 2
	y := 2

	result := add(x, y)
	if result != expectedResult {
		t.Logf("%d %s %d", result, "не равен", expectedResult)
		t.Fail()
	}
}
func TestAddV2(t *testing.T) {
	for _, v := range []struct {
		x   int
		y   int
		exp int
	}{
		{
			x:   1,
			y:   2,
			exp: 3,
		},
		{
			x:   2,
			y:   2,
			exp: 4,
		},
		// {
		// 	x:   5,
		// 	y:   2,
		// 	exp: 8,
		// },
	} {
		res := add(v.x, v.y)
		if res != v.exp {
			t.Logf("%v не равно %v\n", res, v.exp)
			t.Fail()
		}
	}
}
