package testdata

import (
	"io/ioutil"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFilenames() ([]string, error) {
	pattern := filepath.Join(getTestdataDir(), "*.smile")
	files, err := filepath.Glob(pattern)
	if err != nil {
		return files, err
	}

	for i, filename := range files {
		files[i] = strings.TrimSuffix(filename, ".smile")
	}
	return files, nil
}

func LoadTestFile(t *testing.T, filepath string) []byte {
	b, err := ioutil.ReadFile(filepath)
	require.NoError(t, err, "Error reading test file %q", filepath)

	return b
}

func SmileJsTestFile(part string) ([]string, error) {
	pattern := filepath.Join(getTestdataDir(), "..", "..", "smile-js", "testdata", part, "*.smile")
	files, err := filepath.Glob(pattern)
	if err != nil {
		return files, err
	}

	for i, filename := range files {
		files[i] = strings.TrimSuffix(filename, ".smile")
	}
	return files, nil
}

func getTestdataDir() string {
	_, testdataFile, _, _ := runtime.Caller(0)
	return filepath.Dir(testdataFile)
}
