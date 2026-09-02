package dataify

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"testing"
)

type scraperDownloadTransport struct {
	request    *http.Request
	statusCode int
	body       []byte
}

func (t *scraperDownloadTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	t.request = req
	return &http.Response{
		StatusCode: t.statusCode,
		Header:     make(http.Header),
		Body:       io.NopCloser(bytes.NewReader(t.body)),
		Request:    req,
	}, nil
}

func TestScraperDownloadTaskResultRequestsJSONWithoutBearerHeader(t *testing.T) {
	transport := &scraperDownloadTransport{
		statusCode: http.StatusOK,
		body:       []byte(`{"items":[{"id":"result-1"}]}`),
	}
	client := NewClient("api key? &", WithHTTPClient(&http.Client{Transport: transport}))

	result, err := client.Scraper.DownloadTaskResult(context.Background(), " task id? & ")
	if err != nil {
		t.Fatalf("DownloadTaskResult() error = %v", err)
	}
	var decoded map[string]any
	if err := result.Decode(&decoded); err != nil {
		t.Fatalf("Decode() error = %v", err)
	}
	if transport.request.Method != http.MethodGet {
		t.Fatalf("method = %s, want %s", transport.request.Method, http.MethodGet)
	}
	if transport.request.URL.Path != "/download" {
		t.Fatalf("path = %q, want /download", transport.request.URL.Path)
	}
	query := transport.request.URL.Query()
	if got := query.Get("api_key"); got != "api key? &" {
		t.Fatalf("api_key = %q, want %q", got, "api key? &")
	}
	if got := query.Get("task_id"); got != "task id? &" {
		t.Fatalf("task_id = %q, want %q", got, "task id? &")
	}
	if got := query.Get("type"); got != string(ScraperDownloadJSON) {
		t.Fatalf("type = %q, want %q", got, ScraperDownloadJSON)
	}
	if got := transport.request.Header.Get("Authorization"); got != "" {
		t.Fatalf("Authorization = %q, want no header", got)
	}
}

func TestScraperDownloadTaskFilePreservesSupportedFileStreams(t *testing.T) {
	tests := []struct {
		name       string
		format     ScraperDownloadType
		body       []byte
		wantFormat string
	}{
		{name: "json", format: ScraperDownloadJSON, body: []byte(`{"ok":true}`), wantFormat: "json"},
		{name: "csv", format: ScraperDownloadCSV, body: []byte("id,name\n1,coffee\n"), wantFormat: "csv"},
		{name: "xlsx", format: ScraperDownloadXLSX, body: []byte{0x50, 0x4b, 0x03, 0x04}, wantFormat: "xlsx"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transport := &scraperDownloadTransport{statusCode: http.StatusOK, body: tt.body}
			client := NewClient("api-key", WithHTTPClient(&http.Client{Transport: transport}))

			response, err := client.Scraper.DownloadTaskFile(context.Background(), "task-123", tt.format)
			if err != nil {
				t.Fatalf("DownloadTaskFile() error = %v", err)
			}
			defer response.Body.Close()
			body, err := io.ReadAll(response.Body)
			if err != nil {
				t.Fatalf("ReadAll() error = %v", err)
			}
			if !bytes.Equal(body, tt.body) {
				t.Fatalf("body = %v, want %v", body, tt.body)
			}
			if got := transport.request.URL.Query().Get("type"); got != tt.wantFormat {
				t.Fatalf("type = %q, want %q", got, tt.wantFormat)
			}
		})
	}
}

func TestScraperDownloadTaskFileRejectsUnsupportedFormatBeforeRequest(t *testing.T) {
	transport := &scraperDownloadTransport{statusCode: http.StatusOK}
	client := NewClient("api-key", WithHTTPClient(&http.Client{Transport: transport}))

	_, err := client.Scraper.DownloadTaskFile(context.Background(), "task-123", ScraperDownloadType("xml"))
	if !errors.Is(err, ErrUnsupportedDownloadType) {
		t.Fatalf("DownloadTaskFile() error = %v, want ErrUnsupportedDownloadType", err)
	}
	if transport.request != nil {
		t.Fatal("DownloadTaskFile() sent a request for an unsupported format")
	}
}

func TestScraperDownloadPropagatesHTTPFailuresAndMissingLocalToken(t *testing.T) {
	t.Run("http failure", func(t *testing.T) {
		transport := &scraperDownloadTransport{
			statusCode: http.StatusBadRequest,
			body:       []byte(`{"data":"Task_id is error!","code":400}`),
		}
		client := NewClient("api-key", WithHTTPClient(&http.Client{Transport: transport}))

		_, err := client.Scraper.DownloadTaskResult(context.Background(), "unknown-task")
		var apiErr *APIError
		if !errors.As(err, &apiErr) || apiErr.StatusCode != http.StatusBadRequest {
			t.Fatalf("DownloadTaskResult() error = %v, want HTTP 400 APIError", err)
		}
	})

	t.Run("missing local token", func(t *testing.T) {
		t.Setenv(testStandardTokenEnv, "")
		t.Setenv(testLegacyTokenEnv, "")
		t.Setenv(testLegacyAPIKeyEnv, "")
		client := NewClient("")

		_, err := client.Scraper.DownloadTaskFile(context.Background(), "task-123", ScraperDownloadCSV)
		if !errors.Is(err, ErrMissingToken) {
			t.Fatalf("DownloadTaskFile() error = %v, want ErrMissingToken", err)
		}
	})
}
