package parser

import (
	"testing"
)

type parserTruckTest struct {
	input          string
	expectedError  bool
	expectedName   string
	expectedX      int
	expectedY      int
	expectedWeight int
	expectedTurn   int
}

var parserTruckTests = []parserTruckTest{
	{" qsd 0 0 100 10", true, "", 0, 0, 0, 0},
	{"qsd  0 0 100 10", true, "", 0, 0, 0, 0},
	{"qsd 0  0 100 10", true, "", 0, 0, 0, 0},
	{"qsd 0 0  100 10", true, "", 0, 0, 0, 0},
	{"qsd 0 0 100  10", true, "", 0, 0, 0, 0},
	{"qsd 0 0 100 10 ", true, "", 0, 0, 0, 0},

	{"qsd 0 0 10 10", true, "", 0, 0, 0, 0},
	{"qsd 0 0 100 -1", true, "", 0, 0, 0, 0},
	{"qsd 0 -1 100 10", true, "", 0, 0, 0, 0},
	{"qsd -1 0 100 10", true, "", 0, 0, 0, 0},

	{"q sd 0 0 100 10", true, "", 0, 0, 0, 0},
	{"qsd 0 0 100 10 dqs", true, "", 0, 0, 0, 0},

	{"qsd qsd 0 100 10", true, "", 0, 0, 0, 0},
	{"qsd 0 qsdqs 100 10", true, "", 0, 0, 0, 0},
	{"qsd 0 0 qsdqs 10", true, "", 0, 0, 0, 0},
	{"qsd 0 0 100 qsdqs", true, "", 0, 0, 0, 0},

	{"qsd 0 0 100 10", false, "qsd", 0, 0, 100, 10},
	{"qsd 1 0 100 10", false, "qsd", 1, 0, 100, 10},
	{"qsd 1 1 100 10", false, "qsd", 1, 1, 100, 10},
	{"qsd 1 1 101 10", false, "qsd", 1, 1, 101, 10},
}

func TestParserTruck(t *testing.T) {
	for _, test := range parserTruckTests {
		switch name, x, y, weight, turn, err := parserTruck(test.input); {
		case err != nil && !test.expectedError:
			t.Errorf("parserTruck returns an error when it shouldn't")
		case err == nil && test.expectedError:
			t.Errorf("parserTruck does not return an error when it should")
		case err != nil && test.expectedError:
			return
		case name != test.expectedName:
			t.Errorf("Output Name \"%s\" not equal to expected \"%s\"", name, test.expectedName)
		case x != test.expectedX:
			t.Errorf("Output X \"%d\" not equal to expected \"%d\"", x, test.expectedX)
		case y != test.expectedY:
			t.Errorf("Output Y \"%d\" not equal to expected \"%d\"", y, test.expectedY)
		case weight != test.expectedWeight:
			t.Errorf("Output Max Weight \"%d\" not equal to expected \"%d\"", weight, test.expectedWeight)
		case turn != test.expectedTurn:
			t.Errorf("Output Turn \"%d\" not equal to expected \"%d\"", turn, test.expectedTurn)
		}
	}
}
