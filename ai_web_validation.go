package dataify

import (
	"encoding/json"
	"fmt"
	"math"
	"net/url"
	"sort"
	"strings"
	"unicode/utf8"
)

const (
	aiWebMaxQueryLength    = 256
	aiWebMaxSelectorItems  = 50
	aiWebMaxSelectorBytes  = 2048
	aiWebMaxResultsMinimum = 1
	aiWebMaxResultsMaximum = 100
	aiWebMapLimitMinimum   = 1
	aiWebMapLimitMaximum   = 100000
	aiWebMaxAgeMinimum     = 0
	aiWebMaxAgeMaximum     = 172800000
	aiWebTimeoutMinimum    = 1000
	aiWebTimeoutMaximum    = 120000
	aiWebWaitForMinimum    = 0
	aiWebWaitForMaximum    = 30000
)

var (
	aiWebTimeRangeValues = map[string]struct{}{
		"hour": {}, "day": {}, "week": {}, "month": {}, "year": {},
	}
	aiWebFormatValues = map[string]struct{}{
		AIWebFormatMarkdown: {}, AIWebFormatRawHTML: {}, AIWebFormatCleanHTML: {}, AIWebFormatLinks: {},
	}
	aiWebDomainFilterValues = map[string]struct{}{
		AIWebDomainFilterAll: {}, AIWebDomainFilterDomain: {},
	}
	aiWebSitemapValues = map[string]struct{}{
		AIWebSitemapSkip: {}, AIWebSitemapInclude: {}, AIWebSitemapOnly: {},
	}
	aiWebForbiddenHeaders = map[string]struct{}{
		"host": {}, "connection": {}, "keep-alive": {}, "proxy-authenticate": {},
		"proxy-authorization": {}, "te": {}, "trailer": {}, "transfer-encoding": {}, "upgrade": {},
	}
)

func prepareAIWebRequest(operation AIWebOperation, req any) (any, error) {
	if _, ok := aiWebPath(operation); !ok {
		return nil, fmt.Errorf("%w: %q", ErrUnsupportedAIWebOperation, operation)
	}
	switch value := req.(type) {
	case AIWebSearchRequest:
		return prepareAIWebSearch(operation, value)
	case *AIWebSearchRequest:
		if value == nil {
			return nil, newValidationError("request", "must not be nil")
		}
		return prepareAIWebSearch(operation, *value)
	case AIWebMapRequest:
		return prepareAIWebMap(operation, value)
	case *AIWebMapRequest:
		if value == nil {
			return nil, newValidationError("request", "must not be nil")
		}
		return prepareAIWebMap(operation, *value)
	case AIWebExtractRequest:
		return prepareAIWebExtract(operation, value)
	case *AIWebExtractRequest:
		if value == nil {
			return nil, newValidationError("request", "must not be nil")
		}
		return prepareAIWebExtract(operation, *value)
	case map[string]any:
		return prepareAIWebPayload(operation, value)
	default:
		return nil, newValidationError("request", "must be AIWebSearchRequest, AIWebMapRequest, AIWebExtractRequest, or map[string]any")
	}
}

func prepareAIWebSearch(operation AIWebOperation, r AIWebSearchRequest) (AIWebSearchRequest, error) {
	if operation != AIWebOperationSearch {
		return r, newValidationError("operation", "must be %q for AIWebSearchRequest", AIWebOperationSearch)
	}
	uniqueID, err := normalizeUniqueID(r.UniqueID)
	if err != nil {
		return r, err
	}
	r.UniqueID = uniqueID
	r.Query = strings.TrimSpace(r.Query)
	if err := validateQuery(r.Query); err != nil {
		return r, err
	}
	r.Source = strings.TrimSpace(r.Source)
	if expected, known := aiWebSearchSourceForUniqueID(r.UniqueID); known {
		switch {
		case r.Source == "":
			r.Source = expected
		case r.Source != expected:
			return r, newValidationError("source", "must be %q for unique_id %q", expected, r.UniqueID)
		}
	}
	r.Country = normalizeCountry(r.Country)
	if err := validateCountry(r.Country); err != nil {
		return r, err
	}
	r.Language = strings.TrimSpace(r.Language)
	r.TimeRange = strings.TrimSpace(r.TimeRange)
	if err := validateEnum("time_range", r.TimeRange, aiWebTimeRangeValues); err != nil {
		return r, err
	}
	switch r.UniqueID {
	case AIWebUniqueIDShopping:
		if r.TimeRange != "" {
			return r, newValidationError("time_range", "is not supported for shopping search")
		}
	case AIWebUniqueIDScholar:
		if r.TimeRange != "" {
			return r, newValidationError("time_range", "is not supported for scholar search")
		}
		if r.Country != "" {
			return r, newValidationError("country", "is not supported for scholar search")
		}
	}
	if r.MaxResults != nil {
		if err := validateIntRange("max_results", *r.MaxResults, aiWebMaxResultsMinimum, aiWebMaxResultsMaximum); err != nil {
			return r, err
		}
	}
	if r.ExtractOptions != nil {
		if r.ExtractContent == nil || !*r.ExtractContent {
			return r, newValidationError("extract_options", "requires extract_content=true")
		}
		options, err := prepareAIWebExtractOptions(*r.ExtractOptions)
		if err != nil {
			return r, err
		}
		r.ExtractOptions = &options
	}
	return r, nil
}

func prepareAIWebMap(operation AIWebOperation, r AIWebMapRequest) (AIWebMapRequest, error) {
	if operation != AIWebOperationMap {
		return r, newValidationError("operation", "must be %q for AIWebMapRequest", AIWebOperationMap)
	}
	uniqueID, err := normalizeUniqueID(r.UniqueID)
	if err != nil {
		return r, err
	}
	r.UniqueID = uniqueID
	r.URL = strings.TrimSpace(r.URL)
	if err := validateURL("url", r.URL); err != nil {
		return r, err
	}
	r.Country = normalizeCountry(r.Country)
	if err := validateCountry(r.Country); err != nil {
		return r, err
	}
	r.DomainFilter = strings.TrimSpace(r.DomainFilter)
	if err := validateEnum("domain_filter", r.DomainFilter, aiWebDomainFilterValues); err != nil {
		return r, err
	}
	r.Sitemap = strings.TrimSpace(r.Sitemap)
	if err := validateEnum("sitemap", r.Sitemap, aiWebSitemapValues); err != nil {
		return r, err
	}
	if r.Limit != nil {
		if err := validateIntRange("limit", *r.Limit, aiWebMapLimitMinimum, aiWebMapLimitMaximum); err != nil {
			return r, err
		}
	}
	return r, nil
}

func prepareAIWebExtract(operation AIWebOperation, r AIWebExtractRequest) (AIWebExtractRequest, error) {
	if operation != AIWebOperationExtract {
		return r, newValidationError("operation", "must be %q for AIWebExtractRequest", AIWebOperationExtract)
	}
	uniqueID, err := normalizeUniqueID(r.UniqueID)
	if err != nil {
		return r, err
	}
	r.UniqueID = uniqueID
	r.URL = strings.TrimSpace(r.URL)
	if err := validateURL("url", r.URL); err != nil {
		return r, err
	}
	r.Country = normalizeCountry(r.Country)
	if err := validateCountry(r.Country); err != nil {
		return r, err
	}
	formats, err := normalizeFormats(r.Formats)
	if err != nil {
		return r, err
	}
	r.Formats = formats
	includeTags, err := normalizeSelectorList("include_tags", r.IncludeTags)
	if err != nil {
		return r, err
	}
	r.IncludeTags = includeTags
	excludeTags, err := normalizeSelectorList("exclude_tags", r.ExcludeTags)
	if err != nil {
		return r, err
	}
	r.ExcludeTags = excludeTags
	headers, err := normalizeHeaders(r.Headers)
	if err != nil {
		return r, err
	}
	r.Headers = headers
	if err := validateOptionalIntRange("max_age_ms", r.MaxAgeMS, aiWebMaxAgeMinimum, aiWebMaxAgeMaximum); err != nil {
		return r, err
	}
	if err := validateOptionalIntRange("timeout_ms", r.TimeoutMS, aiWebTimeoutMinimum, aiWebTimeoutMaximum); err != nil {
		return r, err
	}
	if err := validateOptionalIntRange("wait_for_ms", r.WaitForMS, aiWebWaitForMinimum, aiWebWaitForMaximum); err != nil {
		return r, err
	}
	if err := validateWaitForTimeout(r.TimeoutMS, r.WaitForMS); err != nil {
		return r, err
	}
	return r, nil
}

func prepareAIWebExtractOptions(options AIWebExtractOptions) (AIWebExtractOptions, error) {
	formats, err := normalizeFormats(options.Formats)
	if err != nil {
		return options, err
	}
	options.Formats = formats
	includeTags, err := normalizeSelectorList("include_tags", options.IncludeTags)
	if err != nil {
		return options, err
	}
	options.IncludeTags = includeTags
	excludeTags, err := normalizeSelectorList("exclude_tags", options.ExcludeTags)
	if err != nil {
		return options, err
	}
	options.ExcludeTags = excludeTags
	if err := validateOptionalIntRange("max_age_ms", options.MaxAgeMS, aiWebMaxAgeMinimum, aiWebMaxAgeMaximum); err != nil {
		return options, err
	}
	if err := validateOptionalIntRange("timeout_ms", options.TimeoutMS, aiWebTimeoutMinimum, aiWebTimeoutMaximum); err != nil {
		return options, err
	}
	if err := validateOptionalIntRange("wait_for_ms", options.WaitForMS, aiWebWaitForMinimum, aiWebWaitForMaximum); err != nil {
		return options, err
	}
	if err := validateWaitForTimeout(options.TimeoutMS, options.WaitForMS); err != nil {
		return options, err
	}
	return options, nil
}

func prepareAIWebPayload(operation AIWebOperation, payload map[string]any) (map[string]any, error) {
	if payload == nil {
		return nil, newValidationError("request", "must not be nil")
	}
	out := cloneMap(payload)
	uniqueID, err := payloadString(out, "unique_id")
	if err != nil {
		return nil, err
	}
	uniqueID = strings.TrimSpace(uniqueID)
	if uniqueID == "" {
		return nil, newValidationError("unique_id", "must not be empty")
	}
	out["unique_id"] = uniqueID

	switch operation {
	case AIWebOperationSearch:
		if err := prepareAIWebSearchPayload(out); err != nil {
			return nil, err
		}
	case AIWebOperationMap:
		if err := prepareAIWebMapPayload(out); err != nil {
			return nil, err
		}
	case AIWebOperationExtract:
		if err := prepareAIWebExtractPayload(out); err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("%w: %q", ErrUnsupportedAIWebOperation, operation)
	}
	return out, nil
}

func prepareAIWebSearchPayload(payload map[string]any) error {
	queryValue, err := payloadString(payload, "query")
	if err != nil {
		return err
	}
	query := strings.TrimSpace(queryValue)
	if err := validateQuery(query); err != nil {
		return err
	}
	payload["query"] = query

	uniqueID, _ := payload["unique_id"].(string)
	source, hasSource, err := optionalPayloadString(payload, "source")
	if err != nil {
		return err
	}
	if expected, known := aiWebSearchSourceForUniqueID(uniqueID); known {
		switch {
		case !hasSource || source == "":
			payload["source"] = expected
		case source != expected:
			return newValidationError("source", "must be %q for unique_id %q", expected, uniqueID)
		}
	} else if hasSource {
		payload["source"] = strings.TrimSpace(source)
	}

	if err := normalizeOptionalPayloadInt(payload, "max_results", aiWebMaxResultsMinimum, aiWebMaxResultsMaximum); err != nil {
		return err
	}
	country, hasCountry, err := optionalPayloadString(payload, "country")
	if err != nil {
		return err
	}
	if hasCountry {
		country = normalizeCountry(country)
		if err := validateCountry(country); err != nil {
			return err
		}
		payload["country"] = country
	}
	timeRange, hasTimeRange, err := optionalPayloadString(payload, "time_range")
	if err != nil {
		return err
	}
	if hasTimeRange {
		timeRange = strings.TrimSpace(timeRange)
		if err := validateEnum("time_range", timeRange, aiWebTimeRangeValues); err != nil {
			return err
		}
		payload["time_range"] = timeRange
	}
	switch uniqueID {
	case AIWebUniqueIDShopping:
		if timeRange != "" {
			return newValidationError("time_range", "is not supported for shopping search")
		}
	case AIWebUniqueIDScholar:
		if timeRange != "" {
			return newValidationError("time_range", "is not supported for scholar search")
		}
		if country != "" {
			return newValidationError("country", "is not supported for scholar search")
		}
	}

	extractContent := false
	if value, exists := payload["extract_content"]; exists && value != nil {
		flagged, err := payloadBool(value, "extract_content")
		if err != nil {
			return err
		}
		extractContent = flagged
		payload["extract_content"] = flagged
	}
	if value, exists := payload["extract_options"]; exists && value != nil {
		options, ok := value.(map[string]any)
		if !ok {
			return newValidationError("extract_options", "must be an object")
		}
		if !extractContent {
			return newValidationError("extract_options", "requires extract_content=true")
		}
		normalized, err := prepareAIWebExtractOptionsPayload(options)
		if err != nil {
			return err
		}
		payload["extract_options"] = normalized
	}
	return nil
}

func prepareAIWebMapPayload(payload map[string]any) error {
	urlValue, err := payloadString(payload, "url")
	if err != nil {
		return err
	}
	urlValue = strings.TrimSpace(urlValue)
	if err := validateURL("url", urlValue); err != nil {
		return err
	}
	payload["url"] = urlValue

	if country, exists, err := optionalPayloadString(payload, "country"); err != nil {
		return err
	} else if exists {
		country = normalizeCountry(country)
		if err := validateCountry(country); err != nil {
			return err
		}
		payload["country"] = country
	}
	if domainFilter, exists, err := optionalPayloadString(payload, "domain_filter"); err != nil {
		return err
	} else if exists {
		domainFilter = strings.TrimSpace(domainFilter)
		if err := validateEnum("domain_filter", domainFilter, aiWebDomainFilterValues); err != nil {
			return err
		}
		payload["domain_filter"] = domainFilter
	}
	if sitemap, exists, err := optionalPayloadString(payload, "sitemap"); err != nil {
		return err
	} else if exists {
		sitemap = strings.TrimSpace(sitemap)
		if err := validateEnum("sitemap", sitemap, aiWebSitemapValues); err != nil {
			return err
		}
		payload["sitemap"] = sitemap
	}
	return normalizeOptionalPayloadInt(payload, "limit", aiWebMapLimitMinimum, aiWebMapLimitMaximum)
}

func prepareAIWebExtractPayload(payload map[string]any) error {
	urlValue, err := payloadString(payload, "url")
	if err != nil {
		return err
	}
	urlValue = strings.TrimSpace(urlValue)
	if err := validateURL("url", urlValue); err != nil {
		return err
	}
	payload["url"] = urlValue

	if country, exists, err := optionalPayloadString(payload, "country"); err != nil {
		return err
	} else if exists {
		country = normalizeCountry(country)
		if err := validateCountry(country); err != nil {
			return err
		}
		payload["country"] = country
	}
	if formats, exists, err := optionalPayloadStringSlice(payload, "formats"); err != nil {
		return err
	} else if exists {
		normalized, err := normalizeFormats(formats)
		if err != nil {
			return err
		}
		payload["formats"] = normalized
	}
	for _, key := range []string{"include_tags", "exclude_tags"} {
		if items, exists, err := optionalPayloadStringSlice(payload, key); err != nil {
			return err
		} else if exists {
			normalized, err := normalizeSelectorList(key, items)
			if err != nil {
				return err
			}
			payload[key] = normalized
		}
	}
	if headers, exists, err := optionalPayloadHeaders(payload, "headers"); err != nil {
		return err
	} else if exists {
		normalized, err := normalizeHeaders(headers)
		if err != nil {
			return err
		}
		payload["headers"] = normalized
	}
	if err := normalizeOptionalPayloadBool(payload, "only_main_content"); err != nil {
		return err
	}
	if err := normalizeOptionalPayloadBool(payload, "advanced_mode"); err != nil {
		return err
	}
	if err := normalizeOptionalPayloadInt(payload, "max_age_ms", aiWebMaxAgeMinimum, aiWebMaxAgeMaximum); err != nil {
		return err
	}
	if err := normalizeOptionalPayloadInt(payload, "timeout_ms", aiWebTimeoutMinimum, aiWebTimeoutMaximum); err != nil {
		return err
	}
	if err := normalizeOptionalPayloadInt(payload, "wait_for_ms", aiWebWaitForMinimum, aiWebWaitForMaximum); err != nil {
		return err
	}
	return validateWaitForTimeout(payloadIntPointer(payload, "timeout_ms"), payloadIntPointer(payload, "wait_for_ms"))
}

func prepareAIWebExtractOptionsPayload(options map[string]any) (map[string]any, error) {
	out := cloneMap(options)
	if formats, exists, err := optionalPayloadStringSlice(out, "formats"); err != nil {
		return nil, err
	} else if exists {
		normalized, err := normalizeFormats(formats)
		if err != nil {
			return nil, err
		}
		out["formats"] = normalized
	}
	for _, key := range []string{"include_tags", "exclude_tags"} {
		if items, exists, err := optionalPayloadStringSlice(out, key); err != nil {
			return nil, err
		} else if exists {
			normalized, err := normalizeSelectorList(key, items)
			if err != nil {
				return nil, err
			}
			out[key] = normalized
		}
	}
	if err := normalizeOptionalPayloadBool(out, "only_main_content"); err != nil {
		return nil, err
	}
	if err := normalizeOptionalPayloadBool(out, "advanced_mode"); err != nil {
		return nil, err
	}
	if err := normalizeOptionalPayloadInt(out, "max_age_ms", aiWebMaxAgeMinimum, aiWebMaxAgeMaximum); err != nil {
		return nil, err
	}
	if err := normalizeOptionalPayloadInt(out, "timeout_ms", aiWebTimeoutMinimum, aiWebTimeoutMaximum); err != nil {
		return nil, err
	}
	if err := normalizeOptionalPayloadInt(out, "wait_for_ms", aiWebWaitForMinimum, aiWebWaitForMaximum); err != nil {
		return nil, err
	}
	if err := validateWaitForTimeout(payloadIntPointer(out, "timeout_ms"), payloadIntPointer(out, "wait_for_ms")); err != nil {
		return nil, err
	}
	return out, nil
}

func normalizeUniqueID(raw string) (string, error) {
	uniqueID := strings.TrimSpace(raw)
	if uniqueID == "" {
		return "", newValidationError("unique_id", "must not be empty")
	}
	return uniqueID, nil
}

func normalizeCountry(value string) string {
	return strings.ToUpper(strings.TrimSpace(value))
}

func validateQuery(query string) error {
	if query == "" {
		return newValidationError("query", "must not be empty")
	}
	if utf8.RuneCountInString(query) > aiWebMaxQueryLength {
		return newValidationError("query", "must be at most %d characters", aiWebMaxQueryLength)
	}
	return nil
}

func validateCountry(value string) error {
	if value == "" {
		return nil
	}
	if len(value) != 2 || !isASCIIAlpha(value) {
		return newValidationError("country", "must be a two-letter country code")
	}
	return nil
}

func isASCIIAlpha(value string) bool {
	for i := 0; i < len(value); i++ {
		character := value[i]
		if (character < 'A' || character > 'Z') && (character < 'a' || character > 'z') {
			return false
		}
	}
	return true
}

func validateEnum(name string, value string, allowed map[string]struct{}) error {
	if value == "" {
		return nil
	}
	if _, ok := allowed[value]; !ok {
		return newValidationError(name, "must be one of %s", allowedValues(allowed))
	}
	return nil
}

func allowedValues(allowed map[string]struct{}) string {
	values := make([]string, 0, len(allowed))
	for value := range allowed {
		values = append(values, value)
	}
	sort.Strings(values)
	return strings.Join(values, ", ")
}

func validateIntRange(name string, value int, minimum int, maximum int) error {
	if value < minimum || value > maximum {
		return newValidationError(name, "must be between %d and %d", minimum, maximum)
	}
	return nil
}

func validateOptionalIntRange(name string, value *int, minimum int, maximum int) error {
	if value == nil {
		return nil
	}
	return validateIntRange(name, *value, minimum, maximum)
}

func validateWaitForTimeout(timeout *int, waitFor *int) error {
	if timeout == nil || waitFor == nil {
		return nil
	}
	if *waitFor > *timeout/2 {
		return newValidationError("wait_for_ms", "must not exceed half of timeout_ms")
	}
	return nil
}

func validateURL(name string, value string) error {
	if value == "" {
		return newValidationError(name, "must not be empty")
	}
	parsed, err := url.Parse(value)
	if err != nil {
		return newValidationError(name, "must be a valid URL")
	}
	if parsed.Scheme != "http" && parsed.Scheme != "https" {
		return newValidationError(name, "must use http:// or https://")
	}
	if parsed.Host == "" {
		return newValidationError(name, "must include a host")
	}
	if parsed.User != nil {
		return newValidationError(name, "must not include user information")
	}
	return nil
}

func normalizeFormats(formats []string) ([]string, error) {
	if len(formats) == 0 {
		return formats, nil
	}
	seen := make(map[string]struct{}, len(formats))
	normalized := make([]string, 0, len(formats))
	for _, raw := range formats {
		format := strings.TrimSpace(raw)
		if _, ok := aiWebFormatValues[format]; !ok {
			return nil, newValidationError("formats", "unsupported format %q; allowed values: %s", format, allowedValues(aiWebFormatValues))
		}
		if _, exists := seen[format]; exists {
			return nil, newValidationError("formats", "must not contain duplicates")
		}
		seen[format] = struct{}{}
		normalized = append(normalized, format)
	}
	return normalized, nil
}

func normalizeSelectorList(name string, values []string) ([]string, error) {
	if len(values) > aiWebMaxSelectorItems {
		return nil, newValidationError(name, "must contain at most %d items", aiWebMaxSelectorItems)
	}
	normalized := make([]string, 0, len(values))
	for _, raw := range values {
		item := strings.TrimSpace(raw)
		if item == "" {
			return nil, newValidationError(name, "must not contain empty values")
		}
		if len([]byte(item)) > aiWebMaxSelectorBytes {
			return nil, newValidationError(name, "values must be at most %d bytes", aiWebMaxSelectorBytes)
		}
		normalized = append(normalized, item)
	}
	return normalized, nil
}

func normalizeHeaders(headers map[string]string) (map[string]string, error) {
	if headers == nil {
		return nil, nil
	}
	normalized := make(map[string]string, len(headers))
	for name, value := range headers {
		name = strings.TrimSpace(name)
		value = strings.TrimSpace(value)
		if name == "" || value == "" {
			return nil, newValidationError("headers", "names and values must not be empty")
		}
		if strings.ContainsAny(name, "\r\n") || strings.ContainsAny(value, "\r\n") {
			return nil, newValidationError("headers", "names and values must not contain newlines")
		}
		if _, forbidden := aiWebForbiddenHeaders[strings.ToLower(name)]; forbidden {
			return nil, newValidationError("headers", "must not contain Host or hop-by-hop headers")
		}
		normalized[name] = value
	}
	return normalized, nil
}

func cloneMap(source map[string]any) map[string]any {
	result := make(map[string]any, len(source))
	for key, value := range source {
		result[key] = value
	}
	return result
}

func payloadString(payload map[string]any, key string) (string, error) {
	value, exists := payload[key]
	if !exists || value == nil {
		return "", newValidationError(key, "must not be empty")
	}
	text, ok := value.(string)
	if !ok {
		return "", newValidationError(key, "must be a string")
	}
	return text, nil
}

func optionalPayloadString(payload map[string]any, key string) (string, bool, error) {
	value, exists := payload[key]
	if !exists || value == nil {
		return "", false, nil
	}
	text, ok := value.(string)
	if !ok {
		return "", false, newValidationError(key, "must be a string")
	}
	return text, true, nil
}

func payloadBool(value any, key string) (bool, error) {
	flagged, ok := value.(bool)
	if !ok {
		return false, newValidationError(key, "must be a boolean")
	}
	return flagged, nil
}

func payloadInt(value any, key string) (int, error) {
	switch number := value.(type) {
	case int:
		return number, nil
	case int8:
		return int(number), nil
	case int16:
		return int(number), nil
	case int32:
		return int(number), nil
	case int64:
		if int64(int(number)) != number {
			return 0, newValidationError(key, "is out of range")
		}
		return int(number), nil
	case uint:
		if uint64(number) > uint64(maxIntValue()) {
			return 0, newValidationError(key, "is out of range")
		}
		return int(number), nil
	case uint8:
		return int(number), nil
	case uint16:
		return int(number), nil
	case uint32:
		return int(number), nil
	case uint64:
		if number > uint64(maxIntValue()) {
			return 0, newValidationError(key, "is out of range")
		}
		return int(number), nil
	case float32:
		return intFromFloat(float64(number), key)
	case float64:
		return intFromFloat(number, key)
	case json.Number:
		parsed, err := number.Int64()
		if err != nil {
			return 0, newValidationError(key, "must be an integer")
		}
		if int64(int(parsed)) != parsed {
			return 0, newValidationError(key, "is out of range")
		}
		return int(parsed), nil
	default:
		return 0, newValidationError(key, "must be an integer")
	}
}

func intFromFloat(value float64, key string) (int, error) {
	if math.Trunc(value) != value {
		return 0, newValidationError(key, "must be an integer")
	}
	converted := int(value)
	if float64(converted) != value {
		return 0, newValidationError(key, "is out of range")
	}
	return converted, nil
}

func maxIntValue() int {
	return int(^uint(0) >> 1)
}

func normalizeOptionalPayloadInt(payload map[string]any, key string, minimum int, maximum int) error {
	value, exists := payload[key]
	if !exists || value == nil {
		return nil
	}
	number, err := payloadInt(value, key)
	if err != nil {
		return err
	}
	if err := validateIntRange(key, number, minimum, maximum); err != nil {
		return err
	}
	payload[key] = number
	return nil
}

func normalizeOptionalPayloadBool(payload map[string]any, key string) error {
	value, exists := payload[key]
	if !exists || value == nil {
		return nil
	}
	flagged, err := payloadBool(value, key)
	if err != nil {
		return err
	}
	payload[key] = flagged
	return nil
}

func payloadIntPointer(payload map[string]any, key string) *int {
	value, exists := payload[key]
	if !exists {
		return nil
	}
	number, ok := value.(int)
	if !ok {
		return nil
	}
	return &number
}

func payloadStringSlice(value any, key string) ([]string, error) {
	switch items := value.(type) {
	case []string:
		return items, nil
	case []any:
		result := make([]string, 0, len(items))
		for index, item := range items {
			text, ok := item.(string)
			if !ok {
				return nil, newValidationError(key, "item %d must be a string", index)
			}
			result = append(result, text)
		}
		return result, nil
	default:
		return nil, newValidationError(key, "must be an array of strings")
	}
}

func optionalPayloadStringSlice(payload map[string]any, key string) ([]string, bool, error) {
	value, exists := payload[key]
	if !exists || value == nil {
		return nil, false, nil
	}
	items, err := payloadStringSlice(value, key)
	return items, true, err
}

func optionalPayloadHeaders(payload map[string]any, key string) (map[string]string, bool, error) {
	value, exists := payload[key]
	if !exists || value == nil {
		return nil, false, nil
	}
	switch headers := value.(type) {
	case map[string]string:
		return headers, true, nil
	case map[string]any:
		result := make(map[string]string, len(headers))
		for name, item := range headers {
			text, ok := item.(string)
			if !ok {
				return nil, false, newValidationError(key+"."+name, "must be a string")
			}
			result[name] = text
		}
		return result, true, nil
	default:
		return nil, false, newValidationError(key, "must be an object")
	}
}

func aiWebSearchSourceForUniqueID(uniqueID string) (string, bool) {
	switch uniqueID {
	case AIWebUniqueIDWeb:
		return AIWebSourceWeb, true
	case AIWebUniqueIDShopping:
		return AIWebSourceShopping, true
	case AIWebUniqueIDNews:
		return AIWebSourceNews, true
	case AIWebUniqueIDScholar:
		return AIWebSourceScholar, true
	default:
		return "", false
	}
}
