package parser

import (
	"os"
	"testing"

	"github.com/darkelle24/go-epitech/game"
)

type getPathTest struct {
	args          []string
	expectedError bool
	expected      string
}

var getPathTests = []getPathTest{
	{[]string{"go", "./.gitignore"}, false, "./.gitignore"},
	{[]string{"go"}, true, ""},
	{[]string{"go", "./.gitignore", "args2"}, true, ""},
}

func TestGetPath(t *testing.T) {

	for _, test := range getPathTests {
		os.Args = test.args
		if output, err := getPath(); err != nil && !test.expectedError {
			t.Errorf("getPath returns an error when it shouldn't")
		} else if err == nil && test.expectedError {
			t.Errorf("getPath does not return an error when it should")
		} else if output != test.expected {
			t.Errorf("Output \"%s\" not equal to expected \"%s\"", output, test.expected)
		}
	}
}

type readFileTest struct {
	createdFile   bool
	path          string
	text          string
	expectedError bool
}

var readFileTests = []readFileTest{
	{true, "./test.txt", "dab", false},
	{true, "./test.txt", "dab\nqsdqs\n", false},
	{true, "./test.txt", "dab\r\nqsdqs\r\n", false},
	{false, "", "", true},
}

func subTestReadFile(t *testing.T, test readFileTest) {
	if test.createdFile {
		f, err := os.Create(test.path)

		if err != nil {
			return
		}

		defer os.Remove(test.path)
		defer f.Close()

		_, err = f.WriteString(test.text)
		if err != nil {
			return
		}
	}

	if output, err := readFile(test.path); err != nil && !test.expectedError {
		t.Errorf("readFile returns an error when it shouldn't")
	} else if err == nil && test.expectedError {
		t.Errorf("readFile does not return an error when it should")
	} else if err != nil && test.expectedError {
		return
	} else if output != test.text {
		t.Errorf("Output \"%s\" not equal to expected \"%s\"", output, test.text)
	}
}

func TestReadFile(t *testing.T) {

	for _, test := range readFileTests {
		subTestReadFile(t, test)
	}
}

type checkNumberTest struct {
	input         string
	expectedError bool
	expected      int
}

var checkNumberTests = []checkNumberTest{
	{"lol", true, 0},
	{"-1.2", true, 0},
	{"1.2", true, 0},

	{"-1", false, -1},
	{"1", false, 1},
}

func TestCheckNumber(t *testing.T) {

	for _, test := range checkNumberTests {
		if output, err := checkNumber(test.input); err != nil && !test.expectedError {
			t.Errorf("checkNumber returns an error when it shouldn't")
		} else if err == nil && test.expectedError {
			t.Errorf("checkNumber does not return an error when it should")
		} else if err != nil && test.expectedError {
			return
		} else if output != test.expected {
			t.Errorf("Output \"%d\" not equal to expected \"%d\"", output, test.expected)
		}
	}
}

type firstLineParseTest struct {
	input         string
	expectedError bool
	expectedW     int
	expectedH     int
	expectedT     int
}

var firstLineParseTests = []firstLineParseTest{
	{"sqd qsd qsqsd", true, 0, 0, 0},
	{"-1 qsd qsqsd", true, 0, 0, 0},
	{"-1 -1 qsqsd", true, 0, 0, 0},
	{"-1 qsqsd", true, 0, 0, 0},
	{"qsqsd", true, 0, 0, 0},
	{"qsqsd qsd qsdqs qsd sqdqs", true, 0, 0, 0},
	{"-1 -1 -1d", true, 0, 0, 0},
	{"-1 -1d -1", true, 0, 0, 0},
	{"-1d -1 -1", true, 0, 0, 0},
	{"-1 -1  -1", true, 0, 0, 0},
	{"-1  -1 -1", true, 0, 0, 0},
	{" -1 -1 -1", true, 0, 0, 0},
	{"-1 -1 -1 ", true, 0, 0, 0},
	{"1 -1 1", true, 0, 0, 0},
	{"1 1 -1", true, 0, 0, 0},
	{"-1 1 1", true, 0, 0, 0},
	{"1 1 1", true, 0, 0, 0},
	{"1 1 20000000000", true, 0, 0, 0},

	{"1 1 1", false, 1, 1, 1},
	{"10 1 1", false, 10, 1, 1},
	{"1 10 1", false, 1, 10, 1},
}

func TestFirstLineParse(t *testing.T) {

	for _, test := range firstLineParseTests {
		if W, H, T, err := firstLineParse(test.input); err != nil && !test.expectedError {
			t.Errorf("firstLineParse returns an error when it shouldn't")
		} else if err == nil && test.expectedError {
			t.Errorf("firstLineParse does not return an error when it should")
		} else if err != nil && test.expectedError {
			return
		} else if W != test.expectedW {
			t.Errorf("Output Width \"%d\" not equal to expected \"%d\"", W, test.expectedW)
		} else if H != test.expectedH {
			t.Errorf("Output Height \"%d\" not equal to expected \"%d\"", H, test.expectedH)
		} else if T != test.expectedT {
			t.Errorf("Output Turn \"%d\" not equal to expected \"%d\"", T, test.expectedT)
		}
	}
}

type orderParserTest struct {
	input         []string
	expectedError bool
}

var orderParserTests = []orderParserTest{
	{[]string{"5 5 1"}, true},
	{[]string{"5 5 1", "qsd 1 0 green"}, true},
	{[]string{"5 5 2", "qsd 1 0 green", "qsdqsqsdd 1 2"}, true},
	{[]string{"5 5 2", "qsd 1 0 green", "qsdqsqsdd 1 3", ""}, true},
	{[]string{"5 5 2", "qsd 1 0 green", "qsdqqsdqsfsd 1 2 100 2"}, true},
	{[]string{"5 5 2", "qsd 1 0 green", "qsdqsqsdd 1 4", "qsdqqsdqsfsd 1 2 100 2", ""}, true},

	{[]string{"5 5 2", "qsd 1 0 green", "qsdqsqsdd 1 4", "qsdqqsdqsfsd 1 2 100 2"}, false},
	{[]string{"5 5 2", "qsd 1 0 green", "qsd 1 1 green", "qsdqsqsdd 4 4", "qsdqsqsdd 0 0", "qsdqqsdqsfsd 1 2 100 2"}, false},
}

func TestOrderParser(t *testing.T) {

	for _, test := range orderParserTests {
		var gameEnv game.Game

		mapError := gameEnv.CreateMap(10, 10)
		if mapError == nil {
			if err := orderParser(test.input, &gameEnv); err == nil && !test.expectedError {
				t.Errorf("firstLineParse returns an error when it shouldn't")
			} else if err == nil && test.expectedError {
				t.Errorf("firstLineParse does not return an error when it should")
			}
		} else {
			t.Errorf("Cannot test firstLineParse because there was an error in the map initialization")
		}
	}
}
