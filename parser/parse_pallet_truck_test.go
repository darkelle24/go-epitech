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
		if name, x, y, err := parserPalletTruck(test.input); err != nil && !test.expectedError {
			t.Errorf("parserPackage returns an error when it shouldn't")
		} else if err == nil && test.expectedError {
			t.Errorf("parserPackage does not return an error when it should")
		} else if err != nil && test.expectedError {
			return
		} else if name != test.expectedName {
			t.Errorf("Output Name \"%s\" not equal to expected \"%s\"", name, test.expectedName)
		} else if x != test.expectedX {
			t.Errorf("Output X \"%d\" not equal to expected \"%d\"", x, test.expectedX)
		} else if y != test.expectedY {
			t.Errorf("Output Y \"%d\" not equal to expected \"%d\"", y, test.expectedY)
		}
	}
}
