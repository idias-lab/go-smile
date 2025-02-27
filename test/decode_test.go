package test

import (
	"encoding/json"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/zencoder/go-smile/smile"
	"github.com/zencoder/go-smile/test/testdata"
)

func TestDecode(t *testing.T) {
	filenames, err := testdata.TestFilenames()
	require.NoError(t, err)

	for _, f := range filenames {
		f := f
		t.Run(filepath.Base(f), func(t *testing.T) {
			jsonFile := testdata.LoadTestFile(t, f+".json")
			smileFile := testdata.LoadTestFile(t, f+".smile")

			actualJSON, err := smile.DecodeToJSON(smileFile)
			require.NoError(t, err, "Error while decoding %q", f)

			require.JSONEq(t, string(jsonFile), actualJSON, "Decoding %q didn't produce the expected result", f)

		})
	}
}

func TestDecodeTestWithSmileJs(t *testing.T) {
	filenames, err := testdata.SmileJsTestFile("basic")
	require.NoError(t, err)

	for _, f := range filenames {
		f := f

		t.Run(filepath.Base(f), func(t *testing.T) {
			jsonFile := testdata.LoadTestFile(t, f+".json")
			smileFile := testdata.LoadTestFile(t, f+".smile")

			actualJSON, err := smile.DecodeToObject(smileFile)
			require.NoError(t, err, "Error while decoding %q", f)
			var expectedObj interface{}
			err = json.Unmarshal(jsonFile, &expectedObj)
			assert.NoError(t, err)

			marshalledExpectedJson, err := json.Marshal(expectedObj.(map[string]interface{})["value"])
			assert.NoError(t, err)
			marshalledActualJson, err := json.Marshal(actualJSON)
			assert.NoError(t, err)
			require.JSONEq(t, string(marshalledExpectedJson), string(marshalledActualJson), "Decoding %q didn't produce the expected result", f)

		})
	}
}
