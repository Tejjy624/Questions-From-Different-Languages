package triangle

import "testing"

func TestGetTriangleType(t *testing.T) {
	type Test struct {
		a, b, c  int
		expected triangleType
	}

	var tests = []Test{
		{1001, 5, 6, UnknownTriangle},
		// TODO add more tests for 100% test coverage
		{900, 2001, 1500, UnknownTriangle},
		{900, 1900, 3001, UnknownTriangle},
		{-1, 10, 100, UnknownTriangle},
		{10, -1, 100, UnknownTriangle},
		{10, 100, -1, UnknownTriangle},
		{1, 10, 100, InvalidTriangle},
		{3, 4, 5, RightTriangle},
		{3, 4, 3, AcuteTriangle},
		{5, 8, 10, ObtuseTriangle},
	}

	for _, test := range tests {
		actual := getTriangleType(test.a, test.b, test.c)
		if actual != test.expected {
			t.Errorf("getTriangleType(%d, %d, %d)=%v; want %v", test.a, test.b, test.c, actual, test.expected)
		}
	}
}
