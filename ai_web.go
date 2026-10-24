package dataify

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
)

// AIWebUniqueIDWeb, AIWebUniqueIDShopping, AIWebUniqueIDNews,
// AIWebUniqueIDScholar, AIWebUniqueIDMap, and AIWebUniqueIDExtract are the
// fixed unique_id values used by the built-in AI Web tools. A custom gateway
// may still accept other values.
const (
	AIWebUniqueIDWeb      = "search_web"
	AIWebUniqueIDShopping = "search_shopping"
	AIWebUniqueIDNews     = "search_news"
	AIWebUniqueIDScholar  = "search_scholar"
	AIWebUniqueIDMap      = "map"
	AIWebUniqueIDExtract  = "extract"
)

// AIWebSource values for AIWebSearchRequest.Source.
const (
	AIWebSourceWeb      = "web"
	AIWebSourceNews     = "news"
	AIWebSourceShopping = "shopping"
	AIWebSourceScholar  = "scholar"
)

// AIWebDomainFilter values for AIWebMapRequest.DomainFilter.
const (
	AIWebDomainFilterAll    = "all"
	AIWebDomainFilterDomain = "domain"
)

// AIWebSitemapMode values for AIWebMapRequest.Sitemap.
const (
	AIWebSitemapSkip    = "skip"
	AIWebSitemapInclude = "include"
	AIWebSitemapOnly    = "only"
)

// AIWebFormat values accepted by Extracts formats.
const (
	AIWebFormatMarkdown  = "markdown"
	AIWebFormatRawHTML   = "raw_html"
	AIWebFormatCleanHTML = "clean_html"
	AIWebFormatLinks     = "links"
)

// AIWebOperation selects the synchronous AI Web gateway route.
type AIWebOperation string

const (
	AIWebOperationSearch  AIWebOperation = "search"
	AIWebOperationMap     AIWebOperation = "map"
	AIWebOperationExtract AIWebOperation = "extract"
)

// ErrMissingAIWebBaseURL is returned before sending a request when the AI Web
// base URL has been explicitly cleared, for example with
// WithAIWebBaseURL(""). Otherwise the SDK defaults to
// https://aiwebapi.dataify.com.
var ErrMissingAIWebBaseURL = newSDKError("dataify: missing AI Web base URL; set AI_WEB_API_BASE_URL or use WithAIWebBaseURL")

// ErrUnsupportedAIWebOperation is matched when Execute or ExecuteRaw receives
// an operation other than search, map, or extract.
var ErrUnsupportedAIWebOperation = newSDKError("dataify: unsupported AI Web operation")

// BoolPtr returns a pointer to value. It is a convenience for the optional
// boolean fields on the AI Web request types, where nil means "omit" and a
// pointer to false keeps the explicit false value on the wire.
func BoolPtr(value bool) *bool { return &value }

// IntPtr returns a pointer to value. It is a convenience for the optional
// integer fields on the AI Web request types, where nil means "omit" and a
// pointer to zero keeps the explicit zero value on the wire.
func IntPtr(value int) *int { return &value }

type AIWebExtractOptions struct {
	Formats         []string `json:"formats,omitempty"`
	OnlyMainContent *bool    `json:"only_main_content,omitempty"`
	IncludeTags     []string `json:"include_tags,omitempty"`
	ExcludeTags     []string `json:"exclude_tags,omitempty"`
	MaxAgeMS        *int     `json:"max_age_ms,omitempty"`
	TimeoutMS       *int     `json:"timeout_ms,omitempty"`
	WaitForMS       *int     `json:"wait_for_ms,omitempty"`
	AdvancedMode    *bool    `json:"advanced_mode,omitempty"`
}

type AIWebSearchRequest struct {
	UniqueID       string               `json:"unique_id"`
	Query          string               `json:"query"`
	MaxResults     *int                 `json:"max_results,omitempty"`
	Source         string               `json:"source,omitempty"`
	Country        string               `json:"country,omitempty"`
	Language       string               `json:"language,omitempty"`
	TimeRange      string               `json:"time_range,omitempty"`
	ExtractContent *bool                `json:"extract_content,omitempty"`
	ExtractOptions *AIWebExtractOptions `json:"extract_options,omitempty"`
}

type AIWebMapRequest struct {
	UniqueID     string `json:"unique_id"`
	URL          string `json:"url"`
	Country      string `json:"country,omitempty"`
	DomainFilter string `json:"domain_filter,omitempty"`
	Limit        *int   `json:"limit,omitempty"`
	Sitemap      string `json:"sitemap,omitempty"`
}

type AIWebExtractRequest struct {
	UniqueID        string            `json:"unique_id"`
	URL             string            `json:"url"`
	Formats         []string          `json:"formats,omitempty"`
	OnlyMainContent *bool             `json:"only_main_content,omitempty"`
	IncludeTags     []string          `json:"include_tags,omitempty"`
	ExcludeTags     []string          `json:"exclude_tags,omitempty"`
	Headers         map[string]string `json:"headers,omitempty"`
	MaxAgeMS        *int              `json:"max_age_ms,omitempty"`
	TimeoutMS       *int              `json:"timeout_ms,omitempty"`
	WaitForMS       *int              `json:"wait_for_ms,omitempty"`
	Country         string            `json:"country,omitempty"`
	AdvancedMode    *bool             `json:"advanced_mode,omitempty"`
}

// AIWebResult contains the request id reported by the AI Web gateway and the
// gateway outputs payload. Search, Map, and Extract return normalized outputs
// (a JSON array); ExecuteRaw and DownloadResult preserve the upstream shape.
// DownloadResult echoes the request id passed by the caller.
type AIWebResult struct {
	RequestID string          `json:"request_id"`
	Outputs   json.RawMessage `json:"outputs"`
}

func (r *AIWebResult) Bytes() []byte {
	if r == nil {
		return nil
	}
	return append([]byte(nil), r.Outputs...)
}

func (r *AIWebResult) String() string {
	if r == nil {
		return ""
	}
	return string(r.Outputs)
}

func (r *AIWebResult) Decode(v any) error {
	if r == nil || len(r.Outputs) == 0 {
		return nil
	}
	return json.Unmarshal(r.Outputs, v)
}

// OutputsList decodes normalized outputs into individual JSON values. It
// returns an empty slice for null or empty outputs.
func (r *AIWebResult) OutputsList() ([]json.RawMessage, error) {
	if r == nil || len(r.Outputs) == 0 {
		return []json.RawMessage{}, nil
	}
	var items []json.RawMessage
	if err := json.Unmarshal(r.Outputs, &items); err != nil {
		return nil, err
	}
	if items == nil {
		return []json.RawMessage{}, nil
	}
	return items, nil
}

// NormalizeAIWebOutputs converts upstream outputs into a JSON array: null and
// empty payloads become [], arrays are preserved, and a single object becomes a
// one-element array. It mirrors the normalization used by the MCP server and
// the TypeScript SDK.
func NormalizeAIWebOutputs(outputs json.RawMessage) json.RawMessage {
	trimmed := strings.TrimSpace(string(outputs))
	if trimmed == "" || trimmed == "null" {
		return json.RawMessage("[]")
	}
	if strings.HasPrefix(trimmed, "[") && json.Valid([]byte(trimmed)) {
		return json.RawMessage(trimmed)
	}
	wrapped := make(json.RawMessage, 0, len(trimmed)+2)
	wrapped = append(wrapped, '[')
	wrapped = append(wrapped, trimmed...)
	wrapped = append(wrapped, ']')
	if !json.Valid(wrapped) {
		return json.RawMessage(trimmed)
	}
	return wrapped
}

type AIWebService struct {
	client *Client
}

// Search calls POST /search and returns normalized outputs.
func (s *AIWebService) Search(ctx context.Context, req AIWebSearchRequest) (*AIWebResult, error) {
	return s.Execute(ctx, AIWebOperationSearch, req)
}

// Map calls POST /map and returns normalized outputs.
func (s *AIWebService) Map(ctx context.Context, req AIWebMapRequest) (*AIWebResult, error) {
	return s.Execute(ctx, AIWebOperationMap, req)
}

// Extract calls POST /extract and returns normalized outputs.
func (s *AIWebService) Extract(ctx context.Context, req AIWebExtractRequest) (*AIWebResult, error) {
	return s.Execute(ctx, AIWebOperationExtract, req)
}

// Execute calls the selected AI Web route and normalizes the outputs to a JSON
// array so callers can decode a single shape across search, map, and extract.
func (s *AIWebService) Execute(ctx context.Context, operation AIWebOperation, req any) (*AIWebResult, error) {
	result, err := s.ExecuteRaw(ctx, operation, req)
	if err != nil {
		return nil, err
	}
	result.Outputs = NormalizeAIWebOutputs(result.Outputs)
	return result, nil
}

// ExecuteRaw calls the selected AI Web route and preserves the upstream outputs
// shape instead of normalizing it to an array.
func (s *AIWebService) ExecuteRaw(ctx context.Context, operation AIWebOperation, req any) (*AIWebResult, error) {
	if s == nil || s.client == nil {
		return nil, newSDKError("dataify: AI Web service is not initialized")
	}
	if err := s.client.requireAIWebBaseURL(); err != nil {
		return nil, err
	}
	prepared, err := prepareAIWebRequest(operation, req)
	if err != nil {
		return nil, err
	}
	path, ok := aiWebPath(operation)
	if !ok {
		return nil, fmt.Errorf("%w: %q", ErrUnsupportedAIWebOperation, operation)
	}
	body, err := json.Marshal(prepared)
	if err != nil {
		return nil, err
	}
	header, responseBody, err := s.client.doJSONRaw(ctx, s.client.aiWebBaseURL, path, body)
	if err != nil {
		return nil, err
	}
	var normalized json.RawMessage
	if len(strings.TrimSpace(string(responseBody))) == 0 {
		normalized = json.RawMessage("null")
	} else {
		normalized, err = normalizeJSON(responseBody)
		if err != nil {
			return nil, fmt.Errorf("%w: %v", ErrInvalidResponse, err)
		}
	}
	return &AIWebResult{
		RequestID: strings.TrimSpace(header.Get("X-Request-Id")),
		Outputs:   append(json.RawMessage(nil), normalized...),
	}, nil
}

// DownloadResult downloads a retained AI Web result. Pass the request_id
// returned by an earlier call.
func (s *AIWebService) DownloadResult(ctx context.Context, requestID string) (*AIWebResult, error) {
	if s == nil || s.client == nil {
		return nil, newSDKError("dataify: AI Web service is not initialized")
	}
	if err := s.client.requireAIWebBaseURL(); err != nil {
		return nil, err
	}
	if s.client.token == "" {
		return nil, ErrMissingToken
	}
	requestID = strings.TrimSpace(requestID)
	if requestID == "" {
		return nil, newValidationError("request_id", "must not be empty")
	}
	body, err := json.Marshal(map[string]string{"task_id": requestID})
	if err != nil {
		return nil, err
	}
	_, responseBody, err := s.client.doJSONRaw(ctx, s.client.aiWebBaseURL, "/download", body)
	if err != nil {
		return nil, err
	}
	normalized, err := normalizeJSON(responseBody)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInvalidResponse, err)
	}
	return &AIWebResult{RequestID: requestID, Outputs: normalized}, nil
}

func (c *Client) requireAIWebBaseURL() error {
	if strings.TrimSpace(c.aiWebBaseURL) == "" {
		return ErrMissingAIWebBaseURL
	}
	return nil
}

func aiWebPath(operation AIWebOperation) (string, bool) {
	switch strings.ToLower(strings.TrimSpace(string(operation))) {
	case "search":
		return "/search", true
	case "map":
		return "/map", true
	case "extract":
		return "/extract", true
	default:
		return "", false
	}
}
