package dataify

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"strings"
	"sync"
	"testing"
)

const (
	testStandardTokenEnv = "DATAIFY_API_TOKEN"
	testLegacyTokenEnv   = "DATAIFY_TOKEN"
	testLegacyAPIKeyEnv  = "DATAIFY_API_KEY"
)

type authorizationCaptureTransport struct {
	authorization string
}

func silenceLegacyTokenWarnings(t *testing.T) {
	t.Helper()
	resetLegacyTokenWarningState(t, io.Discard)
}

func resetLegacyTokenWarningState(t *testing.T, writer io.Writer) {
	t.Helper()
	previousTokenWarningOnce := legacyTokenEnvWarningOnce
	previousAPIKeyWarningOnce := legacyAPIKeyEnvWarningOnce
	previousWriter := legacyTokenWarningWriter
	legacyTokenEnvWarningOnce = &sync.Once{}
	legacyAPIKeyEnvWarningOnce = &sync.Once{}
	legacyTokenWarningWriter = writer
	t.Cleanup(func() {
		legacyTokenEnvWarningOnce = previousTokenWarningOnce
		legacyAPIKeyEnvWarningOnce = previousAPIKeyWarningOnce
		legacyTokenWarningWriter = previousWriter
	})
}

func (t *authorizationCaptureTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	t.authorization = req.Header.Get("Authorization")
	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(strings.NewReader(`{"ok":true}`)),
		Header:     make(http.Header),
		Request:    req,
	}, nil
}

func TestClientResolvesCompatibleEnvironmentTokens(t *testing.T) {
	silenceLegacyTokenWarnings(t)

	tests := []struct {
		name     string
		explicit string
		env      map[string]string
		want     string
	}{
		{
			name:     "uses explicit token before environment values",
			explicit: "explicit-token",
			env: map[string]string{
				testStandardTokenEnv: "standard-token",
				testLegacyTokenEnv:   "legacy-token",
				testLegacyAPIKeyEnv:  "legacy-api-key",
			},
			want: "explicit-token",
		},
		{
			name: "uses standard token before legacy values",
			env: map[string]string{
				testStandardTokenEnv: "standard-token",
				testLegacyTokenEnv:   "legacy-token",
				testLegacyAPIKeyEnv:  "legacy-api-key",
			},
			want: "standard-token",
		},
		{
			name: "uses legacy token when standard token is absent",
			env: map[string]string{
				testLegacyTokenEnv:  "legacy-token",
				testLegacyAPIKeyEnv: "legacy-api-key",
			},
			want: "legacy-token",
		},
		{
			name: "uses legacy api key when no token variable is set",
			env: map[string]string{
				testLegacyAPIKeyEnv: "legacy-api-key",
			},
			want: "legacy-api-key",
		},
		{
			name: "ignores whitespace only values",
			env: map[string]string{
				testStandardTokenEnv: "  ",
				testLegacyTokenEnv:   "\t",
				testLegacyAPIKeyEnv:  "legacy-api-key",
			},
			want: "legacy-api-key",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv(testStandardTokenEnv, "")
			t.Setenv(testLegacyTokenEnv, "")
			t.Setenv(testLegacyAPIKeyEnv, "")
			for key, value := range tt.env {
				t.Setenv(key, value)
			}

			transport := &authorizationCaptureTransport{}
			client := NewClient(tt.explicit, WithHTTPClient(&http.Client{Transport: transport}))
			if _, err := client.Google.Search(context.Background(), GoogleSearchRequest{Q: "token compatibility", JSON: "1"}); err != nil {
				t.Fatalf("Google.Search() error = %v", err)
			}
			if want := "Bearer " + tt.want; transport.authorization != want {
				t.Fatalf("Authorization = %q, want %q", transport.authorization, want)
			}
		})
	}
}

func TestClientMissingTokenErrorMentionsSupportedEnvironmentNames(t *testing.T) {
	silenceLegacyTokenWarnings(t)

	t.Setenv(testStandardTokenEnv, "")
	t.Setenv(testLegacyTokenEnv, "")
	t.Setenv(testLegacyAPIKeyEnv, "")

	client := NewClient("")
	_, err := client.Google.Search(context.Background(), GoogleSearchRequest{Q: "token compatibility", JSON: "1"})
	if !errors.Is(err, ErrMissingToken) {
		t.Fatalf("Google.Search() error = %v, want ErrMissingToken", err)
	}
	for _, name := range []string{testStandardTokenEnv, testLegacyTokenEnv, testLegacyAPIKeyEnv} {
		if !strings.Contains(err.Error(), name) {
			t.Fatalf("missing-token error %q does not mention %s", err, name)
		}
	}
}

func TestClientWarnsOnceForEachLegacyTokenEnvironment(t *testing.T) {
	var warnings bytes.Buffer
	resetLegacyTokenWarningState(t, &warnings)
	t.Setenv(testStandardTokenEnv, "standard-token")
	t.Setenv(testLegacyTokenEnv, "legacy-token")
	t.Setenv(testLegacyAPIKeyEnv, "legacy-api-key")

	NewClient("")
	NewClient("legacy-token")
	NewClient("legacy-token")
	t.Setenv(testStandardTokenEnv, "")
	NewClient("")
	NewClient("")
	t.Setenv(testLegacyTokenEnv, "")
	NewClient("")
	NewClient("")

	output := warnings.String()
	for _, warning := range []string{
		"dataify: DATAIFY_TOKEN 已兼容读取，建议迁移至 DATAIFY_API_TOKEN。",
		"dataify: DATAIFY_API_KEY 已兼容读取，建议迁移至 DATAIFY_API_TOKEN。",
	} {
		if count := strings.Count(output, warning); count != 1 {
			t.Fatalf("warning %q occurred %d times in %q, want once", warning, count, output)
		}
	}
	for _, token := range []string{"standard-token", "legacy-token", "legacy-api-key"} {
		if strings.Contains(output, token) {
			t.Fatalf("warning output leaked token value %q: %q", token, output)
		}
	}
}
