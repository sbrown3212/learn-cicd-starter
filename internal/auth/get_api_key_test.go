package auth

import (
	"fmt"
	"net/http"
	"strings"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		key       string
		value     string
		expect    string
		expectErr string
	}{
		{
			expectErr: "no authorization header",
		},
		{
			key:       "Authorization",
			expectErr: "no authorization header",
		},
		{
			key:       "Authorization",
			value:     "-",
			expectErr: "malformed authorization header",
		},
		{
			key:       "Authorization",
			value:     "Bearer abc123",
			expectErr: "malformed authorization header",
		},
		{
			key:       "Authorization",
			value:     "ApiKey abc123",
			expect:    "abc123",
			expectErr: "not expecting an error",
		},
	}

	for i, tc := range tests {
		t.Run(fmt.Sprintf("TestGetAPIKey Case #%v", i), func(t *testing.T) {
			header := http.Header{}
			header.Add(tc.key, tc.value)

			output, err := GetAPIKey(header)
			if err != nil {
				if strings.Contains(err.Error(), tc.expectErr) {
					return
				}
				t.Errorf("Unexpected: TestGetAPIKey: %v\n", err)
				return
			}

			if output != tc.expect {
				t.Errorf("Unexpected: TestGetAPIKey: %s\n", output)
				return
			}
		})
	}
}
