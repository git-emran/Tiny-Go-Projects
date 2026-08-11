package walk

import "testing"

func TestFilterOut(t *testing.T) {
	testCases := []struct {
		name     string
		file     string
		ext      string
		minSize  int64
		expected bool
	}{
		{"FilterNoExtension", "testdata/dir.log", "", 0, false},
		{"FilterNoExtensionMatch", "testdata/dir.log", "", 0, false},
	}
}
