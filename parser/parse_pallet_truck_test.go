package parser

import (
	"testing"
)

type parserPalletTruckTest struct {
	input         string
	expectedError bool
	expectedName  string
	expectedX     int
	expectedY     int
}

var parserPalletTruckTests = []parserPalletTruckTest{
	{" qsd 0 0", true, "", 0, 0},
	{"qsd  0 0", true, "", 0, 0},
	{"qsd 0  0", true, "", 0, 0},
	{"qsd 0 0 ", true, "", 0, 0},
	{"qsd 0 -1", true, "", 0, 0},
	{"qsd -1 0", true, "", 0, 0},
	{"qsd 0 qsdqd", true, "", 0, 0},
	{"qsd qsdqqds 0", true, "", 0, 0},
	{"q sd 0 0", true, "", 0, 0},

	{"qsd 0 0", false, "qsd", 0, 0},
	{"qsd 1 0", false, "qsd", 1, 0},
	{"qsd 1 1", false, "qsd", 1, 1},
}

func TestParserPalletTruck(t *testing.T) {
	for _, test := range parserPalletTruckTests {
		switch name, x, y, err := parserPalletTruck(test.input); {
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
		}
	}
}
