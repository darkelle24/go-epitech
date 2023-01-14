package parser

import (
	"testing"
)

type parserPackageTest struct {
	input          string
	expectedError  bool
	expectedName   string
	expectedX      int
	expectedY      int
	expectedWeight int
}

var parserPackageTests = []parserPackageTest{
	{" qsd 0 0 green", true, "", 0, 0, 0},
	{"qsd  0 0 green", true, "", 0, 0, 0},
	{"qsd 0  0 green", true, "", 0, 0, 0},
	{"qsd 0 0  green", true, "", 0, 0, 0},
	{"qsd 0 0 grqsqsdeen", true, "", 0, 0, 0},
	{"qsd 0 0 green qds", true, "", 0, 0, 0},
	{"qsd 0 0", true, "", 0, 0, 0},
	{"qsd 0 -1 green", true, "", 0, 0, 0},
	{"qsd -1 0 green", true, "", 0, 0, 0},
	{"qsd 0 qsdqd green", true, "", 0, 0, 0},
	{"qsd qsdqqds 0 green", true, "", 0, 0, 0},
	{"q sd 0 0 green", true, "", 0, 0, 0},

	{"qsd 0 0 green", false, "qsd", 0, 0, 200},
	{"qsd 1 0 green", false, "qsd", 1, 0, 200},
	{"qsd 1 1 green", false, "qsd", 1, 1, 200},
	{"qsd 1 1 blue", false, "qsd", 1, 1, 500},
	{"qsd 1 1 yellow", false, "qsd", 1, 1, 100},
	{"qsd 1 1 YELLOW", false, "qsd", 1, 1, 100},
	{"qsd 1 1 Yellow", false, "qsd", 1, 1, 100},
	{"qsd 1 1 YeLlOw", false, "qsd", 1, 1, 100},
}

func TestParserPackage(t *testing.T) {
	for _, test := range parserPackageTests {
		switch name, x, y, weight, err := parserPackage(test.input); {
		case err != nil && !test.expectedError:
			t.Errorf("parserPackage returns an error when it shouldn't")
		case err == nil && test.expectedError:
			t.Errorf("parserPackage does not return an error when it should")
		case err != nil && test.expectedError:
			return
		case name != test.expectedName:
			t.Errorf("Output Name \"%s\" not equal to expected \"%s\"", name, test.expectedName)
		case x != test.expectedX:
			t.Errorf("Output X \"%d\" not equal to expected \"%d\"", x, test.expectedX)
		case y != test.expectedY:
			t.Errorf("Output Y \"%d\" not equal to expected \"%d\"", y, test.expectedY)
		case weight != test.expectedWeight:
			t.Errorf("Output Weight \"%d\" not equal to expected \"%d\"", weight, test.expectedWeight)
		}
	}
}
