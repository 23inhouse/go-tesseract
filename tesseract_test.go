package tesseract

import (
	"testing"
)

var languageTests = []struct {
	lang     string
	expected string
}{
	{"eng", "Animals are multicellular, eukaryotic organisms of the kingdom Animalia (also called Metazoa). All animals are motile, meaning they can move spontaneously and independently, at some point in their lives. Their body plan eventually becomes fixed as they develop, although some undergo a process of metamorphosis later on in their lives. All animals are heterotrophs: they must ingest other organisms or their products for sustenance."},
	{"nld", "De dieren (wetenschappelijke naam: Animalia), vormen een rijk in de supergroep Unikonta, behorende tot het domein van de eukaryoten. Het dierenrijk is in diverse ondergroepen verdeeld die weer onderverdeeld zijn in stammen. De wetenschap die zich met de studie van het dierenrijk bezighoudt is zoölogie."},
}

func TestProcessWithLanguages(t *testing.T) {
	for _, tt := range languageTests {
		actual := Process("images/"+tt.lang+".png", tt.lang)
		if actual != tt.expected {
			t.Errorf("\nexpected:\t%s\nactual:\t\t%s\n", tt.expected, actual)
		}
	}
}
