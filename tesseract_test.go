package tesseract

import (
	"testing"
)

func TestProcess(t *testing.T) {
	actual, _ := Process("test/english.jpg", "eng")

	expected := "The file format of the source file. For images created by the library itself (via a factory function, or by running a method on an existing image), this attribute is set to None."

	if actual != expected {
		t.Errorf("\nexpected:\t\t%s\nactual:\t%s\n", expected, actual)
	}
}

func TestProcessWithPolishLanguage(t *testing.T) {
	actual, _ := Process("test/polish.png", "pol")

	expected := "Kto był pierwszym królem polski?\nKto jest Twoim tatusiem?"

	if actual != expected {
		t.Errorf("\nexpected:\t\t%s\nactual:\t%s\n", expected, actual)
	}
}
