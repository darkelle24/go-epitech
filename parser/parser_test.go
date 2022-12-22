package parser

import (
	"os"
	"testing"
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
	} else if output != test.text {
		t.Errorf("Output \"%s\" not equal to expected \"%s\"", output, test.text)
	}
}

func TestReadFile(t *testing.T) {

	for _, test := range readFileTests {
		subTestReadFile(t, test)
	}
}

/* type getPathTest struct {
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
} */
