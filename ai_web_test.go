package dataify

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func aiWebTestClient(t *testing.T, transport http.RoundTripper, opts ...Option) *Client {
	t.Helper()
	options := []Option{WithHTTPClient(&http.Client{Transport: transport})}
	options = append(options, opts...)
	return NewClient("test-token", options...)
}

type aiWebCaptureTransport struct {
	request    *http.Request
	body       []byte
	statusCode int
	response   string
	headers    http.Header
}

func (t *aiWebCaptureTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	t.request = req
	if req.Body != nil {
		body, _ := io.ReadAll(req.Body)
		t.body = body
	}
	status := t.statusCode
	if status == 0 {
		status = http.StatusOK
	}
	headers := t.headers
	if headers == nil {
		headers = make(http.Header)
	}
	body := t.response
	if body == "" {
		body = `[]`
	}
	return &http.Response{
		StatusCode: status,
		Header:     headers,
		Body:       io.NopCloser(strings.NewReader(body)),
		Request:    req,
	}, nil
}

func TestAIWebSearchBuildsDocumentedRequestAndReadsRequestID(t *testing.T) {
	transport := &aiWebCaptureTransport{
		headers:  http.Header{"X-Request-Id": []string{" req-123 "}},
		response: `[{"title":"OpenAI","url":"https://example.com/openai"}]`,
	}
	client := aiWebTestClient(t, transport, WithAIWebBaseURL("https://ai-web.example/"))

	result, err := client.AIWeb.Search(context.Background(), AIWebSearchRequest{
		UniqueID:       AIWebUniqueIDNews,
		Query:          " OpenAI ",
		MaxResults:     IntPtr(5),
		Source:         AIWebSourceNews,
		Country:        "US",
		Language:       "en",
		TimeRange:      "week",
		ExtractContent: BoolPtr(true),
		ExtractOptions: &AIWebExtractOptions{
			Formats:         []string{AIWebFormatMarkdown, AIWebFormatLinks},
			OnlyMainContent: BoolPtr(false),
			MaxAgeMS:        IntPtr(0),
			AdvancedMode:    BoolPtr(false),
		},
	})
	if err != nil {
		t.Fatalf("Search() error = %v", err)
	}
	if transport.request.Method != http.MethodPost {
		t.Fatalf("method = %s, want POST", transport.request.Method)
	}
	if transport.request.URL.String() != "https://ai-web.example/search" {
		t.Fatalf("url = %q, want https://ai-web.example/search", transport.request.URL.String())
	}
	if got := transport.request.Header.Get("Authorization"); got != "Bearer test-token" {
		t.Fatalf("Authorization = %q", got)
	}
	if got := transport.request.Header.Get("Content-Type"); got != "application/json" {
		t.Fatalf("Content-Type = %q", got)
	}
	if result.RequestID != "req-123" {
		t.Fatalf("RequestID = %q, want req-123", result.RequestID)
	}
	var payload map[string]any
	if err := json.Unmarshal(transport.body, &payload); err != nil {
		t.Fatalf("payload is not JSON: %v (%s)", err, transport.body)
	}
	if payload["unique_id"] != AIWebUniqueIDNews || payload["query"] != "OpenAI" {
		t.Fatalf("payload = %#v", payload)
	}
	if payload["max_results"] != float64(5) {
		t.Fatalf("max_results = %#v, want 5", payload["max_results"])
	}
	if payload["extract_content"] != true {
		t.Fatalf("extract_content = %#v, want explicit true", payload["extract_content"])
	}
	options, ok := payload["extract_options"].(map[string]any)
	if !ok {
		t.Fatalf("extract_options = %#v", payload["extract_options"])
	}
	if options["only_main_content"] != false {
		t.Fatalf("only_main_content = %#v, want explicit false", options["only_main_content"])
	}
	if options["advanced_mode"] != false {
		t.Fatalf("advanced_mode = %#v, want explicit false", options["advanced_mode"])
	}
	if options["max_age_ms"] != float64(0) {
		t.Fatalf("max_age_ms = %#v, want explicit 0", options["max_age_ms"])
	}
	items, err := result.OutputsList()
	if err != nil {
		t.Fatalf("OutputsList() error = %v", err)
	}
	if len(items) != 1 || !strings.Contains(string(items[0]), "OpenAI") {
		t.Fatalf("outputs = %s", result.Outputs)
	}
}

func TestAIWebMapAndExtractUseTheirRoutes(t *testing.T) {
	tests := []struct {
		name    string
		path    string
		request any
		run     func(*AIWebService) (*AIWebResult, error)
	}{
		{
			name: "map",
			path: "/map",
			request: AIWebMapRequest{
				UniqueID:     AIWebUniqueIDMap,
				URL:          "https://example.com",
				Country:      "US",
				DomainFilter: AIWebDomainFilterAll,
				Limit:        IntPtr(10),
				Sitemap:      AIWebSitemapInclude,
			},
			run: func(s *AIWebService) (*AIWebResult, error) {
				return s.Map(context.Background(), AIWebMapRequest{
					UniqueID:     AIWebUniqueIDMap,
					URL:          "https://example.com",
					Country:      "US",
					DomainFilter: AIWebDomainFilterAll,
					Limit:        IntPtr(10),
					Sitemap:      AIWebSitemapInclude,
				})
			},
		},
		{
			name: "extract",
			path: "/extract",
			request: AIWebExtractRequest{
				UniqueID: AIWebUniqueIDExtract,
				URL:      "https://example.com/article",
			},
			run: func(s *AIWebService) (*AIWebResult, error) {
				return s.Extract(context.Background(), AIWebExtractRequest{
					UniqueID:        AIWebUniqueIDExtract,
					URL:             "https://example.com/article",
					Formats:         []string{AIWebFormatMarkdown},
					OnlyMainContent: BoolPtr(false),
					Headers:         map[string]string{"User-Agent": "dataify-sdk-go"},
					MaxAgeMS:        IntPtr(0),
					TimeoutMS:       IntPtr(60000),
					AdvancedMode:    BoolPtr(false),
				})
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transport := &aiWebCaptureTransport{response: `{"ok":true}`}
			client := aiWebTestClient(t, transport, WithAIWebBaseURL("https://ai-web.example"))

			result, err := tt.run(client.AIWeb)
			if err != nil {
				t.Fatalf("request error = %v", err)
			}
			if transport.request.URL.Path != tt.path {
				t.Fatalf("path = %q, want %q", transport.request.URL.Path, tt.path)
			}
			if string(result.Outputs) != `[{"ok":true}]` {
				t.Fatalf("outputs = %s, want wrapped object", result.Outputs)
			}
		})
	}
}

func TestAIWebExtractSerializesHeadersAndOmittedScalars(t *testing.T) {
	transport := &aiWebCaptureTransport{response: `[]`}
	client := aiWebTestClient(t, transport, WithAIWebBaseURL("https://ai-web.example"))

	_, err := client.AIWeb.Extract(context.Background(), AIWebExtractRequest{
		UniqueID:        AIWebUniqueIDExtract,
		URL:             "https://example.com/article",
		Formats:         []string{AIWebFormatMarkdown},
		OnlyMainContent: BoolPtr(false),
		Headers:         map[string]string{"User-Agent": "dataify-sdk-go"},
		MaxAgeMS:        IntPtr(0),
	})
	if err != nil {
		t.Fatalf("Extract() error = %v", err)
	}
	var payload map[string]any
	if err := json.Unmarshal(transport.body, &payload); err != nil {
		t.Fatalf("payload is not JSON: %v", err)
	}
	if payload["only_main_content"] != false {
		t.Fatalf("only_main_content = %#v, want explicit false", payload["only_main_content"])
	}
	if payload["max_age_ms"] != float64(0) {
		t.Fatalf("max_age_ms = %#v, want explicit 0", payload["max_age_ms"])
	}
	if _, exists := payload["advanced_mode"]; exists {
		t.Fatalf("advanced_mode should be omitted when nil: %#v", payload)
	}
	headers, ok := payload["headers"].(map[string]any)
	if !ok || headers["User-Agent"] != "dataify-sdk-go" {
		t.Fatalf("headers = %#v", payload["headers"])
	}
}

func TestAIWebExecuteNormalizesOutputs(t *testing.T) {
	tests := []struct {
		name     string
		response string
		want     string
	}{
		{name: "array", response: `[{"a":1}]`, want: `[{"a":1}]`},
		{name: "object", response: `{"a":1}`, want: `[{"a":1}]`},
		{name: "null", response: `null`, want: `[]`},
		{name: "empty", response: ``, want: `[]`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transport := &aiWebCaptureTransport{response: tt.response}
			client := aiWebTestClient(t, transport, WithAIWebBaseURL("https://ai-web.example"))

			result, err := client.AIWeb.Extract(context.Background(), AIWebExtractRequest{
				UniqueID: AIWebUniqueIDExtract,
				URL:      "https://example.com",
			})
			if err != nil {
				t.Fatalf("Extract() error = %v", err)
			}
			if string(result.Outputs) != tt.want {
				t.Fatalf("outputs = %s, want %s", result.Outputs, tt.want)
			}
		})
	}
}

func TestAIWebExecuteRawPreservesUpstreamShape(t *testing.T) {
	transport := &aiWebCaptureTransport{
		headers:  http.Header{"X-Request-Id": []string{"raw-1"}},
		response: `{"markdown":"# title"}`,
	}
	client := aiWebTestClient(t, transport, WithAIWebBaseURL("https://ai-web.example"))

	result, err := client.AIWeb.ExecuteRaw(context.Background(), AIWebOperationExtract, AIWebExtractRequest{
		UniqueID: AIWebUniqueIDExtract,
		URL:      "https://example.com",
	})
	if err != nil {
		t.Fatalf("ExecuteRaw() error = %v", err)
	}
	if string(result.Outputs) != `{"markdown":"# title"}` {
		t.Fatalf("outputs = %s, want raw object", result.Outputs)
	}
	if result.RequestID != "raw-1" {
		t.Fatalf("RequestID = %q, want raw-1", result.RequestID)
	}
}

func TestAIWebDownloadResultPostsTaskID(t *testing.T) {
	transport := &aiWebCaptureTransport{response: `{"markdown":"# downloaded"}`}
	client := aiWebTestClient(t, transport, WithAIWebBaseURL("https://ai-web.example"))

	result, err := client.AIWeb.DownloadResult(context.Background(), " req-123 ")
	if err != nil {
		t.Fatalf("DownloadResult() error = %v", err)
	}
	if transport.request.URL.Path != "/download" {
		t.Fatalf("path = %q, want /download", transport.request.URL.Path)
	}
	var payload map[string]string
	if err := json.Unmarshal(transport.body, &payload); err != nil {
		t.Fatalf("payload is not JSON: %v", err)
	}
	if payload["task_id"] != "req-123" {
		t.Fatalf("task_id = %q, want req-123", payload["task_id"])
	}
	if result.RequestID != "req-123" {
		t.Fatalf("RequestID = %q, want req-123", result.RequestID)
	}
	if string(result.Outputs) != `{"markdown":"# downloaded"}` {
		t.Fatalf("Outputs = %s", result.Outputs)
	}
	var decoded map[string]any
	if err := result.Decode(&decoded); err != nil {
		t.Fatalf("Decode() error = %v", err)
	}
	if decoded["markdown"] != "# downloaded" {
		t.Fatalf("downloaded = %#v", decoded)
	}
}

func TestAIWebRejectsMissingConfigurationAndInputsBeforeRequest(t *testing.T) {
	t.Run("missing base URL", func(t *testing.T) {
		t.Setenv(aiWebBaseURLEnv, "")
		transport := &aiWebCaptureTransport{}
		client := aiWebTestClient(t, transport, WithAIWebBaseURL(""))

		_, err := client.AIWeb.Search(context.Background(), AIWebSearchRequest{UniqueID: AIWebUniqueIDNews, Query: "OpenAI"})
		if !errors.Is(err, ErrMissingAIWebBaseURL) {
			t.Fatalf("Search() error = %v, want ErrMissingAIWebBaseURL", err)
		}
		if transport.request != nil {
			t.Fatal("Search() sent a request without a base URL")
		}
	})

	t.Run("missing token", func(t *testing.T) {
		t.Setenv(testStandardTokenEnv, "")
		t.Setenv(testLegacyTokenEnv, "")
		t.Setenv(testLegacyAPIKeyEnv, "")
		transport := &aiWebCaptureTransport{}
		client := NewClient("", WithHTTPClient(&http.Client{Transport: transport}), WithAIWebBaseURL("https://ai-web.example"))

		_, err := client.AIWeb.Search(context.Background(), AIWebSearchRequest{UniqueID: AIWebUniqueIDNews, Query: "OpenAI"})
		if !errors.Is(err, ErrMissingToken) {
			t.Fatalf("Search() error = %v, want ErrMissingToken", err)
		}
		if transport.request != nil {
			t.Fatal("Search() sent a request without a token")
		}
	})

	t.Run("missing query", func(t *testing.T) {
		transport := &aiWebCaptureTransport{}
		client := aiWebTestClient(t, transport, WithAIWebBaseURL("https://ai-web.example"))

		_, err := client.AIWeb.Search(context.Background(), AIWebSearchRequest{UniqueID: AIWebUniqueIDNews})
		if err == nil || !strings.Contains(err.Error(), "query") {
			t.Fatalf("Search() error = %v, want missing query", err)
		}
		if transport.request != nil {
			t.Fatal("Search() sent a request without query")
		}
	})

	t.Run("missing unique id", func(t *testing.T) {
		transport := &aiWebCaptureTransport{}
		client := aiWebTestClient(t, transport, WithAIWebBaseURL("https://ai-web.example"))

		_, err := client.AIWeb.Map(context.Background(), AIWebMapRequest{URL: "https://example.com"})
		if err == nil || !strings.Contains(err.Error(), "unique_id") {
			t.Fatalf("Map() error = %v, want missing unique_id", err)
		}
		if transport.request != nil {
			t.Fatal("Map() sent a request without unique_id")
		}
	})

	t.Run("unsupported operation", func(t *testing.T) {
		transport := &aiWebCaptureTransport{}
		client := aiWebTestClient(t, transport, WithAIWebBaseURL("https://ai-web.example"))

		_, err := client.AIWeb.Execute(context.Background(), AIWebOperation("crawl"), AIWebExtractRequest{UniqueID: AIWebUniqueIDExtract, URL: "https://example.com"})
		if err == nil || !strings.Contains(err.Error(), "unsupported") {
			t.Fatalf("Execute() error = %v, want unsupported operation", err)
		}
		if transport.request != nil {
			t.Fatal("Execute() sent a request for an unsupported operation")
		}
	})
}

func TestAIWebPropagatesHTTPFailure(t *testing.T) {
	transport := &aiWebCaptureTransport{
		statusCode: http.StatusBadRequest,
		response:   `{"code":"INVALID_TOKEN","message":"invalid token"}`,
	}
	client := aiWebTestClient(t, transport, WithAIWebBaseURL("https://ai-web.example"))

	_, err := client.AIWeb.Search(context.Background(), AIWebSearchRequest{UniqueID: AIWebUniqueIDNews, Query: "OpenAI"})
	var apiErr *APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("Search() error = %v, want APIError", err)
	}
	if apiErr.StatusCode != http.StatusBadRequest {
		t.Fatalf("status = %d, want 400", apiErr.StatusCode)
	}
	if !strings.Contains(string(apiErr.Body), "INVALID_TOKEN") {
		t.Fatalf("body = %s", apiErr.Body)
	}
}

func TestAIWebExecuteRawSupportsFlexiblePayloadAndMissingBaseURL(t *testing.T) {
	t.Run("flexible payload", func(t *testing.T) {
		transport := &aiWebCaptureTransport{response: `[{"ok":true}]`}
		client := aiWebTestClient(t, transport, WithAIWebBaseURL("https://ai-web.example"))

		result, err := client.AIWeb.ExecuteRaw(context.Background(), AIWebOperationSearch, map[string]any{
			"unique_id":   "custom_search",
			"query":       "OpenAI",
			"max_results": 3,
		})
		if err != nil {
			t.Fatalf("ExecuteRaw() error = %v", err)
		}
		if string(result.Outputs) != `[{"ok":true}]` {
			t.Fatalf("outputs = %s", result.Outputs)
		}
		var payload map[string]any
		if err := json.Unmarshal(transport.body, &payload); err != nil {
			t.Fatalf("payload is not JSON: %v", err)
		}
		if payload["unique_id"] != "custom_search" || payload["max_results"] != float64(3) {
			t.Fatalf("payload = %#v", payload)
		}
	})

	t.Run("missing base URL is checked before request shape", func(t *testing.T) {
		t.Setenv(aiWebBaseURLEnv, "")
		transport := &aiWebCaptureTransport{}
		client := aiWebTestClient(t, transport, WithAIWebBaseURL(""))

		_, err := client.AIWeb.ExecuteRaw(context.Background(), AIWebOperationSearch, map[string]any{})
		if !errors.Is(err, ErrMissingAIWebBaseURL) {
			t.Fatalf("ExecuteRaw() error = %v, want ErrMissingAIWebBaseURL", err)
		}
		if transport.request != nil {
			t.Fatal("ExecuteRaw() sent a request without a base URL")
		}
	})
}
func TestAIWebBaseURLResolution(t *testing.T) {
	t.Run("explicit option wins over environment", func(t *testing.T) {
		t.Setenv(aiWebBaseURLEnv, "https://from-env.example")
		client := NewClient("token", WithAIWebBaseURL(" https://explicit.example/ "))
		if client.aiWebBaseURL != "https://explicit.example" {
			t.Fatalf("aiWebBaseURL = %q, want https://explicit.example", client.aiWebBaseURL)
		}
	})

	t.Run("environment supplies the base URL", func(t *testing.T) {
		t.Setenv(aiWebBaseURLEnv, " https://from-env.example/ ")
		client := NewClient("token")
		if client.aiWebBaseURL != "https://from-env.example" {
			t.Fatalf("aiWebBaseURL = %q, want https://from-env.example", client.aiWebBaseURL)
		}
	})

	t.Run("empty environment uses production default", func(t *testing.T) {
		t.Setenv(aiWebBaseURLEnv, "   ")
		client := NewClient("token")
		if client.aiWebBaseURL != defaultAIWebBaseURL {
			t.Fatalf("aiWebBaseURL = %q, want %q", client.aiWebBaseURL, defaultAIWebBaseURL)
		}
	})
}

func TestAIWebDefaultBaseURLUsedForRequests(t *testing.T) {
	t.Setenv(aiWebBaseURLEnv, "")
	tests := []struct {
		name string
		path string
		run  func(*AIWebService) error
	}{
		{
			name: "search",
			path: "/search",
			run: func(service *AIWebService) error {
				_, err := service.Search(context.Background(), AIWebSearchRequest{UniqueID: AIWebUniqueIDNews, Query: "OpenAI"})
				return err
			},
		},
		{
			name: "map",
			path: "/map",
			run: func(service *AIWebService) error {
				_, err := service.Map(context.Background(), AIWebMapRequest{UniqueID: AIWebUniqueIDMap, URL: "https://example.com"})
				return err
			},
		},
		{
			name: "extract",
			path: "/extract",
			run: func(service *AIWebService) error {
				_, err := service.Extract(context.Background(), AIWebExtractRequest{UniqueID: AIWebUniqueIDExtract, URL: "https://example.com"})
				return err
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transport := &aiWebCaptureTransport{response: `[]`}
			client := NewClient("token", WithHTTPClient(&http.Client{Transport: transport}))
			if err := tt.run(client.AIWeb); err != nil {
				t.Fatalf("request error = %v", err)
			}
			if transport.request == nil {
				t.Fatal("request was not sent")
			}
			if got := transport.request.URL.Scheme; got != "https" {
				t.Fatalf("scheme = %q, want https", got)
			}
			if got := transport.request.URL.Host; got != "aiwebapi.dataify.com" {
				t.Fatalf("host = %q, want aiwebapi.dataify.com", got)
			}
			if got := transport.request.URL.Path; got != tt.path {
				t.Fatalf("path = %q, want %q", got, tt.path)
			}
		})
	}
}
func TestNormalizeAIWebOutputs(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{name: "empty", input: "", want: "[]"},
		{name: "null", input: "null", want: "[]"},
		{name: "array", input: `[1,2]`, want: `[1,2]`},
		{name: "object", input: `{"a":1}`, want: `[{"a":1}]`},
		{name: "number", input: `42`, want: `[42]`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := string(NormalizeAIWebOutputs(json.RawMessage(tt.input))); got != tt.want {
				t.Fatalf("NormalizeAIWebOutputs(%q) = %s, want %s", tt.input, got, tt.want)
			}
		})
	}
}

func TestAIWebUsesHTTPServerEndToEnd(t *testing.T) {
	var gotPath string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotPath = r.URL.Path
		w.Header().Set("X-Request-Id", "server-1")
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`[{"title":"result"}]`))
	}))
	defer server.Close()

	client := NewClient("server-token", WithAIWebBaseURL(server.URL))
	result, err := client.AIWeb.Search(context.Background(), AIWebSearchRequest{
		UniqueID: AIWebUniqueIDNews,
		Query:    "OpenAI",
	})
	if err != nil {
		t.Fatalf("Search() error = %v", err)
	}
	if gotPath != "/search" {
		t.Fatalf("path = %q, want /search", gotPath)
	}
	if result.RequestID != "server-1" {
		t.Fatalf("RequestID = %q, want server-1", result.RequestID)
	}
	if !bytes.Equal(result.Outputs, []byte(`[{"title":"result"}]`)) {
		t.Fatalf("outputs = %s", result.Outputs)
	}
}

type aiWebRoundTripError struct {
	err error
}

func (t aiWebRoundTripError) RoundTrip(*http.Request) (*http.Response, error) { return nil, t.err }

func assertAIWebValidationError(t *testing.T, err error, wantParameter string) {
	t.Helper()
	if err == nil {
		t.Fatal("error = nil, want validation error")
	}
	var validation *ValidationError
	if !errors.As(err, &validation) {
		t.Fatalf("error = %T (%v), want *ValidationError", err, err)
	}
	if validation.Parameter != wantParameter {
		t.Fatalf("validation parameter = %q, want %q (%v)", validation.Parameter, wantParameter, err)
	}
	if !IsDataifyError(err) {
		t.Fatalf("IsDataifyError(%v) = false", err)
	}
}

func TestAIWebWebSearchMode(t *testing.T) {
	transport := &aiWebCaptureTransport{response: `[{"title":"OpenAI"}]`}
	client := aiWebTestClient(t, transport, WithAIWebBaseURL("https://ai-web.example"))

	result, err := client.AIWeb.Search(context.Background(), AIWebSearchRequest{
		UniqueID:   AIWebUniqueIDWeb,
		Query:      "OpenAI",
		MaxResults: IntPtr(2),
		Country:    "us",
		TimeRange:  "week",
	})
	if err != nil {
		t.Fatalf("Search() error = %v", err)
	}
	var payload map[string]any
	if err := json.Unmarshal(transport.body, &payload); err != nil {
		t.Fatalf("payload is not JSON: %v", err)
	}
	if payload["unique_id"] != AIWebUniqueIDWeb {
		t.Fatalf("unique_id = %#v, want %q", payload["unique_id"], AIWebUniqueIDWeb)
	}
	if payload["source"] != AIWebSourceWeb {
		t.Fatalf("source = %#v, want %q", payload["source"], AIWebSourceWeb)
	}
	if payload["country"] != "US" {
		t.Fatalf("country = %#v, want US", payload["country"])
	}
	if payload["time_range"] != "week" {
		t.Fatalf("time_range = %#v, want week", payload["time_range"])
	}
	items, err := result.OutputsList()
	if err != nil {
		t.Fatalf("OutputsList() error = %v", err)
	}
	if len(items) != 1 || !strings.Contains(string(items[0]), "OpenAI") {
		t.Fatalf("outputs = %s", result.Outputs)
	}
}
func TestAIWebValidatesSearchModeAndParameters(t *testing.T) {
	longQuery := strings.Repeat("a", aiWebMaxQueryLength+1)
	tests := []struct {
		name string
		req  AIWebSearchRequest
		want string
	}{
		{
			name: "shopping rejects time_range",
			req:  AIWebSearchRequest{UniqueID: AIWebUniqueIDShopping, Query: "mouse", Source: AIWebSourceShopping, TimeRange: "week"},
			want: "time_range",
		},
		{
			name: "scholar rejects time_range",
			req:  AIWebSearchRequest{UniqueID: AIWebUniqueIDScholar, Query: "llm", Source: AIWebSourceScholar, TimeRange: "week"},
			want: "time_range",
		},
		{
			name: "scholar rejects country",
			req:  AIWebSearchRequest{UniqueID: AIWebUniqueIDScholar, Query: "llm", Source: AIWebSourceScholar, Country: "US"},
			want: "country",
		},
		{
			name: "source must match unique_id",
			req:  AIWebSearchRequest{UniqueID: AIWebUniqueIDShopping, Query: "mouse", Source: AIWebSourceNews},
			want: "source",
		},
		{
			name: "web source must match unique_id",
			req:  AIWebSearchRequest{UniqueID: AIWebUniqueIDWeb, Query: "OpenAI", Source: AIWebSourceNews},
			want: "source",
		},
		{
			name: "max_results lower bound",
			req:  AIWebSearchRequest{UniqueID: AIWebUniqueIDNews, Query: "OpenAI", MaxResults: IntPtr(0)},
			want: "max_results",
		},
		{
			name: "max_results upper bound",
			req:  AIWebSearchRequest{UniqueID: AIWebUniqueIDNews, Query: "OpenAI", MaxResults: IntPtr(101)},
			want: "max_results",
		},
		{
			name: "time_range enum",
			req:  AIWebSearchRequest{UniqueID: AIWebUniqueIDNews, Query: "OpenAI", TimeRange: "decade"},
			want: "time_range",
		},
		{
			name: "query length",
			req:  AIWebSearchRequest{UniqueID: AIWebUniqueIDNews, Query: longQuery},
			want: "query",
		},
		{
			name: "extract_options requires extract_content",
			req:  AIWebSearchRequest{UniqueID: AIWebUniqueIDNews, Query: "OpenAI", ExtractOptions: &AIWebExtractOptions{Formats: []string{AIWebFormatMarkdown}}},
			want: "extract_options",
		},
		{
			name: "extract format enum",
			req: AIWebSearchRequest{
				UniqueID:       AIWebUniqueIDNews,
				Query:          "OpenAI",
				ExtractContent: BoolPtr(true),
				ExtractOptions: &AIWebExtractOptions{Formats: []string{"bogus"}},
			},
			want: "formats",
		},
		{
			name: "extract format duplicates",
			req: AIWebSearchRequest{
				UniqueID:       AIWebUniqueIDNews,
				Query:          "OpenAI",
				ExtractContent: BoolPtr(true),
				ExtractOptions: &AIWebExtractOptions{Formats: []string{AIWebFormatMarkdown, AIWebFormatMarkdown}},
			},
			want: "formats",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transport := &aiWebCaptureTransport{}
			client := aiWebTestClient(t, transport, WithAIWebBaseURL("https://ai-web.example"))
			_, err := client.AIWeb.Search(context.Background(), tt.req)
			assertAIWebValidationError(t, err, tt.want)
			if transport.request != nil {
				t.Fatal("Search() sent a request for an invalid payload")
			}
		})
	}
}

func TestAIWebNormalizesSourceQueryAndCountry(t *testing.T) {
	transport := &aiWebCaptureTransport{response: `[]`}
	client := aiWebTestClient(t, transport, WithAIWebBaseURL("https://ai-web.example"))

	_, err := client.AIWeb.Search(context.Background(), AIWebSearchRequest{
		UniqueID: AIWebUniqueIDShopping,
		Query:    "  wireless mouse  ",
		Country:  "us",
	})
	if err != nil {
		t.Fatalf("Search() error = %v", err)
	}
	var payload map[string]any
	if err := json.Unmarshal(transport.body, &payload); err != nil {
		t.Fatalf("payload is not JSON: %v", err)
	}
	if payload["query"] != "wireless mouse" {
		t.Fatalf("query = %#v, want trimmed value", payload["query"])
	}
	if payload["source"] != AIWebSourceShopping {
		t.Fatalf("source = %#v, want %q", payload["source"], AIWebSourceShopping)
	}
	if payload["country"] != "US" {
		t.Fatalf("country = %#v, want US", payload["country"])
	}
}

func TestAIWebValidatesMapAndExtractParameters(t *testing.T) {
	selectors := make([]string, aiWebMaxSelectorItems+1)
	for i := range selectors {
		selectors[i] = "div"
	}
	tests := []struct {
		name string
		run  func(*AIWebService) error
		want string
	}{
		{
			name: "map limit lower bound",
			run: func(service *AIWebService) error {
				_, err := service.Map(context.Background(), AIWebMapRequest{UniqueID: AIWebUniqueIDMap, URL: "https://example.com", Limit: IntPtr(0)})
				return err
			},
			want: "limit",
		},
		{
			name: "map limit upper bound",
			run: func(service *AIWebService) error {
				_, err := service.Map(context.Background(), AIWebMapRequest{UniqueID: AIWebUniqueIDMap, URL: "https://example.com", Limit: IntPtr(100001)})
				return err
			},
			want: "limit",
		},
		{
			name: "map domain_filter enum",
			run: func(service *AIWebService) error {
				_, err := service.Map(context.Background(), AIWebMapRequest{UniqueID: AIWebUniqueIDMap, URL: "https://example.com", DomainFilter: "bogus"})
				return err
			},
			want: "domain_filter",
		},
		{
			name: "map sitemap enum",
			run: func(service *AIWebService) error {
				_, err := service.Map(context.Background(), AIWebMapRequest{UniqueID: AIWebUniqueIDMap, URL: "https://example.com", Sitemap: "bogus"})
				return err
			},
			want: "sitemap",
		},
		{
			name: "map country format",
			run: func(service *AIWebService) error {
				_, err := service.Map(context.Background(), AIWebMapRequest{UniqueID: AIWebUniqueIDMap, URL: "https://example.com", Country: "usa"})
				return err
			},
			want: "country",
		},
		{
			name: "extract URL format",
			run: func(service *AIWebService) error {
				_, err := service.Extract(context.Background(), AIWebExtractRequest{UniqueID: AIWebUniqueIDExtract, URL: "not-a-url"})
				return err
			},
			want: "url",
		},
		{
			name: "extract format enum",
			run: func(service *AIWebService) error {
				_, err := service.Extract(context.Background(), AIWebExtractRequest{UniqueID: AIWebUniqueIDExtract, URL: "https://example.com", Formats: []string{"bogus"}})
				return err
			},
			want: "formats",
		},
		{
			name: "extract wait_for_ms constraint",
			run: func(service *AIWebService) error {
				_, err := service.Extract(context.Background(), AIWebExtractRequest{
					UniqueID:  AIWebUniqueIDExtract,
					URL:       "https://example.com",
					TimeoutMS: IntPtr(10000),
					WaitForMS: IntPtr(6000),
				})
				return err
			},
			want: "wait_for_ms",
		},
		{
			name: "extract selector count",
			run: func(service *AIWebService) error {
				_, err := service.Extract(context.Background(), AIWebExtractRequest{
					UniqueID:    AIWebUniqueIDExtract,
					URL:         "https://example.com",
					IncludeTags: selectors,
				})
				return err
			},
			want: "include_tags",
		},
		{
			name: "extract forbidden header",
			run: func(service *AIWebService) error {
				_, err := service.Extract(context.Background(), AIWebExtractRequest{
					UniqueID: AIWebUniqueIDExtract,
					URL:      "https://example.com",
					Headers:  map[string]string{"Host": "example.com"},
				})
				return err
			},
			want: "headers",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transport := &aiWebCaptureTransport{}
			client := aiWebTestClient(t, transport, WithAIWebBaseURL("https://ai-web.example"))
			err := tt.run(client.AIWeb)
			assertAIWebValidationError(t, err, tt.want)
			if transport.request != nil {
				t.Fatal("request was sent for an invalid payload")
			}
		})
	}
}

func TestAIWebValidationPayloadNormalization(t *testing.T) {
	transport := &aiWebCaptureTransport{response: `[]`}
	client := aiWebTestClient(t, transport, WithAIWebBaseURL("https://ai-web.example"))

	_, err := client.AIWeb.Execute(context.Background(), AIWebOperationSearch, map[string]any{
		"unique_id":       AIWebUniqueIDNews,
		"query":           "  OpenAI  ",
		"max_results":     3,
		"country":         "us",
		"extract_content": true,
		"extract_options": map[string]any{
			"formats": []any{"markdown"},
		},
	})
	if err != nil {
		t.Fatalf("Execute() error = %v", err)
	}
	var payload map[string]any
	if err := json.Unmarshal(transport.body, &payload); err != nil {
		t.Fatalf("payload is not JSON: %v", err)
	}
	if payload["query"] != "OpenAI" || payload["country"] != "US" {
		t.Fatalf("payload normalization = %#v", payload)
	}
	if payload["source"] != AIWebSourceNews {
		t.Fatalf("source = %#v, want %q", payload["source"], AIWebSourceNews)
	}
	if payload["max_results"] != float64(3) {
		t.Fatalf("max_results = %#v, want 3", payload["max_results"])
	}
	options, ok := payload["extract_options"].(map[string]any)
	if !ok {
		t.Fatalf("extract_options = %#v", payload["extract_options"])
	}
	formats, ok := options["formats"].([]any)
	if !ok || len(formats) != 1 || formats[0] != AIWebFormatMarkdown {
		t.Fatalf("formats = %#v", options["formats"])
	}
}

func TestAIWebErrorTypes(t *testing.T) {
	validSearch := AIWebSearchRequest{UniqueID: AIWebUniqueIDNews, Query: "OpenAI"}

	t.Run("unsupported operation", func(t *testing.T) {
		transport := &aiWebCaptureTransport{}
		client := aiWebTestClient(t, transport, WithAIWebBaseURL("https://ai-web.example"))
		_, err := client.AIWeb.Execute(context.Background(), AIWebOperation("crawl"), AIWebExtractRequest{
			UniqueID: AIWebUniqueIDExtract,
			URL:      "https://example.com",
		})
		if !errors.Is(err, ErrUnsupportedAIWebOperation) {
			t.Fatalf("error = %v, want ErrUnsupportedAIWebOperation", err)
		}
		if !IsDataifyError(err) {
			t.Fatalf("IsDataifyError(%v) = false", err)
		}
		if transport.request != nil {
			t.Fatal("request was sent for an unsupported operation")
		}
	})

	t.Run("validation error", func(t *testing.T) {
		transport := &aiWebCaptureTransport{}
		client := aiWebTestClient(t, transport, WithAIWebBaseURL("https://ai-web.example"))
		_, err := client.AIWeb.Search(context.Background(), AIWebSearchRequest{UniqueID: AIWebUniqueIDNews})
		assertAIWebValidationError(t, err, "query")
	})

	t.Run("timeout", func(t *testing.T) {
		transport := aiWebRoundTripError{err: &url.Error{Op: "Post", URL: "https://ai-web.example/search", Err: context.DeadlineExceeded}}
		client := aiWebTestClient(t, transport, WithAIWebBaseURL("https://ai-web.example"))
		_, err := client.AIWeb.Search(context.Background(), validSearch)
		if !errors.Is(err, ErrTimeout) {
			t.Fatalf("error = %v, want ErrTimeout", err)
		}
		if !errors.Is(err, context.DeadlineExceeded) {
			t.Fatalf("error = %v, want wrapped context deadline", err)
		}
		var network *NetworkError
		if !errors.As(err, &network) || !network.Timeout() {
			t.Fatalf("error = %T (%v), want timeout NetworkError", err, err)
		}
		if !IsDataifyError(err) {
			t.Fatalf("IsDataifyError(%v) = false", err)
		}
	})

	t.Run("canceled", func(t *testing.T) {
		transport := aiWebRoundTripError{err: &url.Error{Op: "Post", URL: "https://ai-web.example/search", Err: context.Canceled}}
		client := aiWebTestClient(t, transport, WithAIWebBaseURL("https://ai-web.example"))
		_, err := client.AIWeb.Search(context.Background(), validSearch)
		if !errors.Is(err, ErrCanceled) {
			t.Fatalf("error = %v, want ErrCanceled", err)
		}
		if !errors.Is(err, context.Canceled) {
			t.Fatalf("error = %v, want wrapped context canceled", err)
		}
	})

	t.Run("API error", func(t *testing.T) {
		transport := &aiWebCaptureTransport{statusCode: http.StatusUnauthorized, response: `{"code":"UNAUTHORIZED"}`}
		client := aiWebTestClient(t, transport, WithAIWebBaseURL("https://ai-web.example"))
		_, err := client.AIWeb.Search(context.Background(), validSearch)
		var apiErr *APIError
		if !errors.As(err, &apiErr) || apiErr.StatusCode != http.StatusUnauthorized {
			t.Fatalf("error = %T (%v), want APIError 401", err, err)
		}
		if !IsDataifyError(err) {
			t.Fatalf("IsDataifyError(%v) = false", err)
		}
	})

	t.Run("missing token", func(t *testing.T) {
		t.Setenv(testStandardTokenEnv, "")
		t.Setenv(testLegacyTokenEnv, "")
		t.Setenv(testLegacyAPIKeyEnv, "")
		transport := &aiWebCaptureTransport{}
		client := NewClient("", WithHTTPClient(&http.Client{Transport: transport}), WithAIWebBaseURL("https://ai-web.example"))
		_, err := client.AIWeb.Search(context.Background(), validSearch)
		if !errors.Is(err, ErrMissingToken) {
			t.Fatalf("error = %v, want ErrMissingToken", err)
		}
		if !IsDataifyError(err) {
			t.Fatalf("IsDataifyError(%v) = false", err)
		}
	})
}
