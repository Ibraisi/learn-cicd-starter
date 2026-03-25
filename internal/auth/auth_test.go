package auth

import (
	"net/http"
	"testing"
)

func Test_GetAPIkey(t *testing.T) {
	cases := []struct {
		name     string
		input    http.Header
		errorMsh string
	}{
		{
			name:     "empty api key",
			input:    createHeader("", ""),
			errorMsh: "no authorization header included",
		},
		{
			name:     "wrong format",
			input:    createHeader("Authorization", "wrongformattedd"),
			errorMsh: "malformed authorization header",
		},
		{
			name:     "valid format",
			input:    createHeader("Authorization", "ApiKey not important"),
			errorMsh: "",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			_, err := GetAPIKey(c.input)
			if c.errorMsh == "" && err != nil {
				t.Errorf("expected no error, got: %v", err)
			}
			if c.errorMsh != "" && (err == nil || err.Error() != c.errorMsh) {
				t.Errorf("expected error %q, got: %v", c.errorMsh, err)
			}
		})
	}
}

func createHeader(key, value string) http.Header {
	header := http.Header{}
	header.Add(key, value)
	return header
}
