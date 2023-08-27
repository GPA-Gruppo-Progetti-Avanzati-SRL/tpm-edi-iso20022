package iso20022_model

import (
	"github.com/stretchr/testify/require"
	"io/fs"
	"io/ioutil"
	"os"
	"testing"
)

const Example_Pain_013_001_07 = "example-document-pain_013_001_07.xml"

func TestDocumentPain_013_001_07(t *testing.T) {
	d := Pain_013_001_07_Document{}

	b, err := d.ToXML()
	require.NoError(t, err)

	err = ioutil.WriteFile(Example_Pain_013_001_07, b, fs.ModePerm)
	require.NoError(t, err)
	defer os.Remove(Example_Pain_013_001_07)

}
