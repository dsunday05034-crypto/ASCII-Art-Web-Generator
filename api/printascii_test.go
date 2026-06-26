package api

import (
	"strings"
	"testing"
)

func TestPrintAscii(t *testing.T) {
	// Table-driven test cases covering typical inputs and edge cases
	tests := []struct {
		name       string
		banner     string
		input      string
		color      string
		subMatch   string
		wantSubstr string
		wantErr    bool
	}{
		{
			name:       "Valid Standard Generation",
			banner:     "standard",
			input:      "HELL0",
			color:      "",
			subMatch:   "",
			wantSubstr: "|______|", // Standard structural horizontal bar matching your banner file
			wantErr:    false,
		},
		{
			name:       "Color Substring Highlighting",
			banner:     "standard",
			input:      "Go",
			color:      "#ff0000",
			subMatch:   "G",
			wantSubstr: "<span style=\"color:#ff0000;\">", // Verifies RGBA HTML wrapping logic
			wantErr:    false,
		},
		{
			name:       "Missing Banner File Error Handling",
			banner:     "nonexistent_style",
			input:      "Test",
			color:      "",
			subMatch:   "",
			wantSubstr: "",
			wantErr:    true, // Expects an error from embed.FS
		},
		{
			name:       "Empty Input Handled Cleanly",
			banner:     "standard",
			input:      "",
			color:      "",
			subMatch:   "",
			wantSubstr: "",
			wantErr:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PrintAscii(tt.banner, tt.input, tt.color, tt.subMatch)

			// Validate error presence expectations
			if (err != nil) != tt.wantErr {
				t.Fatalf("PrintAscii() error = %v, wantErr %v", err, tt.wantErr)
			}

			// Validate structural output substrings if no error occurred
			if !tt.wantErr && tt.wantSubstr != "" && !strings.Contains(got, tt.wantSubstr) {
				t.Errorf("PrintAscii() output missing expected substring.\nExpected to contain: %q\nGot entire payload.", tt.wantSubstr)
			}
		})
	}
}
