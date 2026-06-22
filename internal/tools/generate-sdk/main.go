package main

import (
	"bytes"
	"fmt"
	"go/format"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

type Param struct {
	Name     string
	GoName   string
	Type     string
	Required bool
	Default  string
	Min      string
	Max      string
}

type Tool struct {
	Name          string
	Group         string
	GroupType     string
	Method        string
	RequestType   string
	ServiceKind   string
	SourceFile    string
	Engine        string
	SpiderName    string
	SpiderID      string
	DynamicSpider bool
	Params        []Param
}

var initialisms = map[string]string{
	"api": "API", "id": "ID", "url": "URL", "json": "JSON", "html": "HTML", "png": "PNG", "http": "HTTP",
	"serp": "SERP", "mcp": "MCP", "ai": "AI", "cid": "CID", "uid": "UID", "db": "DB", "seo": "SEO",
}

var brandWords = map[string]string{
	"youtube": "YouTube", "tiktok": "TikTok", "github": "GitHub", "linkedin": "LinkedIn", "ebay": "EBay", "duckduckgo": "DuckDuckGo",
}

var serpTools = map[string]string{
	"google_search": "google", "google_ai_mode": "google_ai_mode", "google_news": "google_news", "google_images": "google_images",
	"google_maps": "google_maps", "google_flights": "google_flights", "google_jobs": "google_jobs", "google_local": "google_local",
	"google_videos": "google_videos", "google_shopping": "google_shopping", "google_trends": "google_trends", "google_play": "google_play",
	"google_scholar": "google_scholar", "google_finance": "google_finance", "google_hotels": "google_hotels", "google_patents": "google_patents",
	"google_lens": "google_lens",
	"bing_search": "bing", "bing_images": "bing_images", "bing_maps": "bing_maps", "bing_news": "bing_news", "bing_shopping": "bing_shopping", "bing_videos": "bing_videos",
	"duckduckgo_search": "duckduckgo", "yandex_search": "yandex",
}

func main() {
	sourceRoot := getenv("DATAIFY_MCP_SOURCE", "")
	if sourceRoot == "" {
		must(fmt.Errorf("DATAIFY_MCP_SOURCE is required"))
	}
	outRoot := getenv("DATAIFY_SDK_OUT", ".")

	tools, err := loadTools(filepath.Join(sourceRoot, "internal", "tools"))
	must(err)
	sort.Slice(tools, func(i, j int) bool { return tools[i].Name < tools[j].Name })

	must(os.MkdirAll(filepath.Join(outRoot, "docs"), 0755))
	writeText(filepath.Join(outRoot, "go.mod"), goMod())
	writeGo(filepath.Join(outRoot, "client.go"), clientGo(tools))
	writeGo(filepath.Join(outRoot, "response.go"), responseGo())
	writeGo(filepath.Join(outRoot, "tools_generated.go"), toolsGo(tools))
	writeText(filepath.Join(outRoot, "README.md"), readme(tools))
	writeText(filepath.Join(outRoot, "docs", "tools.yaml"), toolsYAML(tools))
	writeText(filepath.Join(outRoot, "docs", "tools.md"), toolsMarkdown(tools))

	fmt.Printf("generated %d tools in %s\n", len(tools), outRoot)
}

func getenv(key, fallback string) string {
	if v := strings.TrimSpace(os.Getenv(key)); v != "" {
		return v
	}
	return fallback
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func loadTools(root string) ([]Tool, error) {
	var tools []Tool
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || filepath.Base(path) != "tool.go" {
			return nil
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		text := string(data)
		name := firstSubmatch(regexp.MustCompile(`mcp\.NewTool\("([^"]+)"`), text)
		if name == "" {
			return nil
		}
		if strings.HasPrefix(name, "query_") {
			return nil
		}
		constants := parseConstants(text)
		params := parseParams(text, constants)
		tool := Tool{
			Name:       name,
			SourceFile: strings.ReplaceAll(path, `\\`, `/`),
			Params:     params,
		}
		classify(&tool, text, constants)
		applyToolCorrections(&tool)
		tools = append(tools, tool)
		return nil
	})
	return tools, err
}

func classify(t *Tool, text string, constants map[string]string) {
	if t.Name == "request_web_unlocker" {
		t.ServiceKind = "web_unlocker"
		t.Group = "WebUnlocker"
		t.Method = "Request"
	} else if engine, ok := serpTools[t.Name]; ok {
		t.ServiceKind = "serp"
		t.Engine = engine
		t.Group = groupName(t.Name)
		t.Method = methodName(t.Name, t.Group)
	} else {
		t.ServiceKind = "scraper_builder"
		t.Group = groupName(t.Name)
		t.Method = methodName(t.Name, t.Group)
		spiderNameRef := firstSubmatch(regexp.MustCompile(`SpiderName:\s*([A-Za-z0-9_]+)`), text)
		t.SpiderName = constants[spiderNameRef]
		spiderIDRef := firstSubmatch(regexp.MustCompile(`SpiderID:\s*([A-Za-z0-9_]+)`), text)
		if spiderIDRef == "spiderID" || spiderIDRef == "strings" || spiderIDRef == "" {
			t.DynamicSpider = true
			t.SpiderID = defaultForParam(t.Params, "spider_id")
		} else {
			t.SpiderID = constants[spiderIDRef]
		}
		if t.SpiderName == "" {
			t.SpiderName = inferredSpiderName(t.Name)
		}
	}
	t.GroupType = t.Group + "Service"
	t.RequestType = camel(t.Name) + "Request"
}

func applyToolCorrections(t *Tool) {
	switch t.Name {
	case "glassdoor_company":
		setParamDefault(t.Params, "url", "https://www.glassdoor.co.uk/Overview/Working-at-Apple-EI_IE1138.11,16.htm")
	case "glassdoor_joblistings":
		setParamDefault(t.Params, "url", "https://www.glassdoor.com/Job/new-york-data-analyst-jobs-SRCH_IL.0,8_IC1132348_KO9,21.htm")
	case "youtube_video_post":
		setParamDefault(t.Params, "url", "https://www.youtube.com/@stephcurry/videos")
		setParamDefault(t.Params, "num_of_posts", "5")
	}
}

func setParamDefault(params []Param, name string, value string) {
	for i := range params {
		if params[i].Name == name {
			params[i].Default = value
			return
		}
	}
}

func inferredSpiderName(tool string) string {
	switch strings.Split(tool, "_")[0] {
	case "amazon":
		return "amazon.com"
	case "youtube":
		return "youtube.com"
	case "tiktok":
		return "tiktok.com"
	case "facebook":
		return "facebook.com"
	case "instagram":
		return "instagram.com"
	case "reddit":
		return "reddit.com"
	case "walmart":
		return "walmart.com"
	case "zillow":
		return "zillow.com"
	case "google":
		return "google.com"
	case "github":
		return "github.com"
	case "linkedin":
		return "linkedin.com"
	case "glassdoor":
		return "glassdoor.com"
	case "indeed":
		return "indeed.com"
	case "crunchbase":
		return "crunchbase.com"
	case "airbnb":
		return "airbnb.com"
	case "booking":
		return "booking.com"
	case "ebay":
		return "ebay.com"
	}
	return strings.Split(tool, "_")[0]
}

func defaultForParam(params []Param, name string) string {
	for _, p := range params {
		if p.Name == name {
			return p.Default
		}
	}
	return ""
}

func parseConstants(text string) map[string]string {
	out := map[string]string{}
	re := regexp.MustCompile("(?m)^\\s*([A-Za-z_][A-Za-z0-9_]*)\\s*=\\s*(\"(?:[^\"\\\\]|\\\\.)*\"|`[^`]*`)")
	for _, m := range re.FindAllStringSubmatch(text, -1) {
		v, err := strconv.Unquote(m[2])
		if err == nil {
			out[m[1]] = v
		}
	}
	return out
}

func parseParams(text string, constants map[string]string) []Param {
	lines := strings.Split(text, "\n")
	var params []Param
	seen := map[string]bool{}
	reStart := regexp.MustCompile(`mcp\.With(String|Integer|Number|Boolean)\("([^"]+)"`)
	for i := 0; i < len(lines); i++ {
		m := reStart.FindStringSubmatch(lines[i])
		if m == nil {
			continue
		}
		typ := strings.ToLower(m[1])
		name := m[2]
		block := lines[i]
		balance := strings.Count(lines[i], "(") - strings.Count(lines[i], ")")
		for balance > 0 && i+1 < len(lines) {
			i++
			block += "\n" + lines[i]
			balance += strings.Count(lines[i], "(") - strings.Count(lines[i], ")")
		}
		if seen[name] {
			continue
		}
		seen[name] = true
		p := Param{Name: name, GoName: fieldName(name), Type: typ}
		p.Required = strings.Contains(block, "mcp.Required()")
		p.Default = parseDefault(block, constants)
		p.Min = firstSubmatch(regexp.MustCompile(`mcp\.Min\(([^\)]+)\)`), block)
		p.Max = firstSubmatch(regexp.MustCompile(`mcp\.Max\(([^\)]+)\)`), block)
		params = append(params, p)
	}
	return params
}

func parseDefault(block string, constants map[string]string) string {
	if m := regexp.MustCompile(`mcp\.DefaultString\(([^\)]+)\)`).FindStringSubmatch(block); m != nil {
		raw := strings.TrimSpace(m[1])
		if strings.HasPrefix(raw, "\"") {
			if v, err := strconv.Unquote(raw); err == nil {
				return v
			}
		}
		return constants[raw]
	}
	if m := regexp.MustCompile(`mcp\.DefaultNumber\(([^\)]+)\)`).FindStringSubmatch(block); m != nil {
		return strings.TrimSpace(m[1])
	}
	return ""
}

func firstSubmatch(re *regexp.Regexp, s string) string {
	m := re.FindStringSubmatch(s)
	if m == nil {
		return ""
	}
	return strings.TrimSpace(m[1])
}

func groupName(tool string) string {
	if tool == "request_web_unlocker" {
		return "WebUnlocker"
	}
	return camel(strings.Split(tool, "_")[0])
}

func methodName(tool, group string) string {
	prefix := strings.ToLower(group)
	for k, v := range brandWords {
		if v == group {
			prefix = k
		}
	}
	rest := strings.TrimPrefix(tool, prefix+"_")
	if rest == tool {
		rest = tool
	}
	return camel(rest)
}

func fieldName(name string) string {
	return camel(name)
}

func camel(s string) string {
	parts := splitWords(s)
	if len(parts) == 0 {
		return "Value"
	}
	for i, p := range parts {
		lower := strings.ToLower(p)
		if v, ok := brandWords[lower]; ok {
			parts[i] = v
			continue
		}
		if v, ok := initialisms[lower]; ok {
			parts[i] = v
			continue
		}
		parts[i] = strings.ToUpper(lower[:1]) + lower[1:]
	}
	out := strings.Join(parts, "")
	if out == "" {
		return "Value"
	}
	if out == "Type" {
		return "Type"
	}
	if r := rune(out[0]); unicode.IsDigit(r) {
		return "Value" + out
	}
	return out
}

func splitWords(s string) []string {
	f := func(r rune) bool { return !(unicode.IsLetter(r) || unicode.IsDigit(r)) }
	raw := strings.FieldsFunc(s, f)
	var out []string
	for _, p := range raw {
		if p != "" {
			out = append(out, p)
		}
	}
	return out
}

func goMod() string { return "module github.com/dataify-server/dataify-sdk-go\n\ngo 1.20\n" }

func clientGo(tools []Tool) string {
	groups := uniqueGroups(tools)
	var b strings.Builder
	b.WriteString(`package dataify

import (
    "bytes"
    "context"
    "encoding/json"
    "errors"
    "fmt"
    "io"
    "net/http"
    "net/url"
    "strings"
    "time"
)

const (
    defaultSERPBaseURL = "https://scraperapi.dataify.com"
    defaultScraperBaseURL = "https://scraperapi.dataify.com"
    defaultWebUnlockerBaseURL = "https://webunlocker.dataify.com"
)

var ErrMissingToken = errors.New("dataify: missing API token")

type APIError struct {
    StatusCode int
    Body []byte
}

func (e *APIError) Error() string {
    return fmt.Sprintf("dataify: api returned HTTP %d: %s", e.StatusCode, truncate(e.Body))
}

type Client struct {
    token string
    httpClient *http.Client
    serpBaseURL string
    scraperBaseURL string
    webUnlockerBaseURL string
`)
	for _, g := range groups {
		fmt.Fprintf(&b, "    %s *%sService\n", g, g)
	}
	b.WriteString(`}

type Option func(*Client)

func WithHTTPClient(httpClient *http.Client) Option { return func(c *Client) { if httpClient != nil { c.httpClient = httpClient } } }
func WithSERPBaseURL(baseURL string) Option { return func(c *Client) { c.serpBaseURL = strings.TrimRight(strings.TrimSpace(baseURL), "/") } }
func WithScraperBaseURL(baseURL string) Option { return func(c *Client) { c.scraperBaseURL = strings.TrimRight(strings.TrimSpace(baseURL), "/") } }
func WithWebUnlockerBaseURL(baseURL string) Option { return func(c *Client) { c.webUnlockerBaseURL = strings.TrimRight(strings.TrimSpace(baseURL), "/") } }
func WithTimeout(timeout time.Duration) Option { return func(c *Client) { if timeout > 0 { c.httpClient.Timeout = timeout } } }

func NewClient(token string, opts ...Option) *Client {
    c := &Client{
        token: strings.TrimSpace(token),
        httpClient: &http.Client{Timeout: 120 * time.Second},
        serpBaseURL: defaultSERPBaseURL,
        scraperBaseURL: defaultScraperBaseURL,
        webUnlockerBaseURL: defaultWebUnlockerBaseURL,
    }
    for _, opt := range opts { if opt != nil { opt(c) } }
`)
	for _, g := range groups {
		fmt.Fprintf(&b, "    c.%s = &%sService{client: c}\n", g, g)
	}
	b.WriteString(`    return c
}

func (c *Client) doForm(ctx context.Context, baseURL, path string, values map[string]string) (*RawResponse, error) {
    if c.token == "" { return nil, ErrMissingToken }
    form := url.Values{}
    for k, v := range values {
        if strings.TrimSpace(v) != "" { form.Set(k, v) }
    }
    req, err := http.NewRequestWithContext(ctx, http.MethodPost, strings.TrimRight(baseURL, "/")+path, strings.NewReader(form.Encode()))
    if err != nil { return nil, err }
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
    req.Header.Set("Authorization", bearer(c.token))
    return c.do(req)
}

func (c *Client) doJSON(ctx context.Context, baseURL, path string, payload map[string]string) (*RawResponse, error) {
    if c.token == "" { return nil, ErrMissingToken }
    body, err := json.Marshal(payload)
    if err != nil { return nil, err }
    req, err := http.NewRequestWithContext(ctx, http.MethodPost, strings.TrimRight(baseURL, "/")+path, bytes.NewReader(body))
    if err != nil { return nil, err }
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Authorization", bearer(c.token))
    return c.do(req)
}

func (c *Client) do(req *http.Request) (*RawResponse, error) {
    resp, err := c.httpClient.Do(req)
    if err != nil { return nil, err }
    defer resp.Body.Close()
    body, err := io.ReadAll(resp.Body)
    if err != nil { return nil, err }
    if resp.StatusCode < 200 || resp.StatusCode >= 300 { return nil, &APIError{StatusCode: resp.StatusCode, Body: body} }
    return NewRawResponse(body)
}

func bearer(token string) string {
    token = strings.TrimSpace(token)
    if token == "" || strings.HasPrefix(strings.ToLower(token), "bearer ") { return token }
    return "Bearer " + token
}

func addStringParam(values map[string]string, key string, value string) { if strings.TrimSpace(value) != "" { values[key] = strings.TrimSpace(value) } }
func addIntParam(values map[string]string, key string, value int) { if value != 0 { values[key] = fmt.Sprint(value) } }
func defaultParam(values map[string]string, key string, value string) { if strings.TrimSpace(values[key]) == "" && value != "" { values[key] = value } }
func requireParam(values map[string]string, key string) error { if strings.TrimSpace(values[key]) == "" { return fmt.Errorf("dataify: missing required parameter %s", key) }; return nil }
func truncate(body []byte) string { s := strings.TrimSpace(string(body)); if len(s) > 500 { return s[:500] + "..." }; return s }
`)
	return b.String()
}

func responseGo() string {
	return `package dataify

import (
    "encoding/json"
    "strings"
)

type RawResponse struct { raw json.RawMessage }

func NewRawResponse(body []byte) (*RawResponse, error) {
    raw, err := normalizeJSON(body)
    if err != nil { return nil, err }
    return &RawResponse{raw: raw}, nil
}

func (r *RawResponse) Bytes() []byte { if r == nil { return nil }; return append([]byte(nil), r.raw...) }
func (r *RawResponse) String() string { if r == nil { return "" }; return string(r.raw) }
func (r *RawResponse) Decode(v any) error { if r == nil { return nil }; return json.Unmarshal(r.raw, v) }
func (r *RawResponse) MarshalJSON() ([]byte, error) { if r == nil { return []byte("null"), nil }; return r.raw.MarshalJSON() }

func normalizeJSON(body []byte) (json.RawMessage, error) {
    var value any
    if err := json.Unmarshal(body, &value); err != nil { return nil, err }
    if text, ok := value.(string); ok {
        trimmed := strings.TrimSpace(text)
        if trimmed != "" && json.Valid([]byte(trimmed)) { return normalizeJSON([]byte(trimmed)) }
    }
    normalized, err := json.Marshal(value)
    if err != nil { return nil, err }
    return json.RawMessage(normalized), nil
}
`
}

func toolsGo(tools []Tool) string {
	var b strings.Builder
	b.WriteString("package dataify\n\nimport (\n    \"context\"\n    \"encoding/json\"\n)\n\n")
	for _, g := range uniqueGroups(tools) {
		fmt.Fprintf(&b, "type %sService struct { client *Client }\n\n", g)
	}
	for _, t := range tools {
		writeRequestStruct(&b, t)
		writeMethod(&b, t)
	}
	writeSpecs(&b, tools)
	return b.String()
}

func writeRequestStruct(b *strings.Builder, t Tool) {
	fmt.Fprintf(b, "type %s struct {\n", t.RequestType)
	for _, p := range t.Params {
		typ := "string"
		if p.Type == "integer" {
			typ = "int"
		}
		fmt.Fprintf(b, "    %s %s `json:%s`\n", p.GoName, typ, strconv.Quote(p.Name+",omitempty"))
	}
	b.WriteString("}\n\n")
	fmt.Fprintf(b, "func (r %s) params() map[string]string {\n    values := map[string]string{}\n", t.RequestType)
	for _, p := range t.Params {
		if p.Type == "integer" {
			fmt.Fprintf(b, "    addIntParam(values, %s, r.%s)\n", strconv.Quote(p.Name), p.GoName)
		} else {
			fmt.Fprintf(b, "    addStringParam(values, %s, r.%s)\n", strconv.Quote(p.Name), p.GoName)
		}
	}
	b.WriteString("    return values\n}\n\n")
}

func writeMethod(b *strings.Builder, t Tool) {
	fmt.Fprintf(b, "func (s *%s) %s(ctx context.Context, req %s) (*RawResponse, error) {\n", t.GroupType, t.Method, t.RequestType)
	b.WriteString("    values := req.params()\n")
	for _, p := range t.Params {
		if p.Default == "" {
			continue
		}
		shouldDefault := t.ServiceKind != "scraper_builder" || !t.DynamicSpider || p.Required || p.Name == "spider_id" || p.Name == "file_name"
		if shouldDefault {
			fmt.Fprintf(b, "    defaultParam(values, %s, %s)\n", strconv.Quote(p.Name), strconv.Quote(p.Default))
		}
	}
	if t.DynamicSpider {
		defaults := defaultSpiderParamDefaults(t)
		if len(defaults) > 0 && t.SpiderID != "" {
			fmt.Fprintf(b, "    if values[\"spider_id\"] == %s {\n", strconv.Quote(t.SpiderID))
			for _, p := range defaults {
				fmt.Fprintf(b, "        defaultParam(values, %s, %s)\n", strconv.Quote(p.Name), strconv.Quote(p.Default))
			}
			b.WriteString("    }\n")
		}
	}
	for _, p := range t.Params {
		if p.Required {
			fmt.Fprintf(b, "    if err := requireParam(values, %s); err != nil { return nil, err }\n", strconv.Quote(p.Name))
		}
	}
	for _, name := range extraRequiredParams(t.Name) {
		fmt.Fprintf(b, "    if err := requireParam(values, %s); err != nil { return nil, err }\n", strconv.Quote(name))
	}
	switch t.ServiceKind {
	case "web_unlocker":
		b.WriteString("    values[\"isjson\"] = \"1\"\n")
		b.WriteString("    return s.client.doJSON(ctx, s.client.webUnlockerBaseURL, \"/request\", values)\n")
	case "serp":
		fmt.Fprintf(b, "    values[\"engine\"] = %s\n", strconv.Quote(t.Engine))
		b.WriteString("    return s.client.doForm(ctx, s.client.serpBaseURL, \"/request\", values)\n")
	case "scraper_builder":
		if t.DynamicSpider {
			b.WriteString("    spiderID := values[\"spider_id\"]\n")
			if t.SpiderID != "" {
				fmt.Fprintf(b, "    if spiderID == \"\" { spiderID = %s }\n", strconv.Quote(t.SpiderID))
			}
		} else {
			fmt.Fprintf(b, "    spiderID := %s\n", strconv.Quote(t.SpiderID))
		}
		fmt.Fprintf(b, "    spiderName := %s\n", strconv.Quote(t.SpiderName))
		b.WriteString("    fileName := values[\"file_name\"]\n    if fileName == \"\" { fileName = \"{{TasksID}}\" }\n")
		b.WriteString("    delete(values, \"spider_id\")\n    delete(values, \"file_name\")\n")
		b.WriteString("    spiderParameters, err := json.Marshal([]map[string]string{values})\n    if err != nil { return nil, err }\n")
		b.WriteString("    payload := map[string]string{\n        \"spider_name\": spiderName,\n        \"spider_id\": spiderID,\n        \"spider_parameters\": string(spiderParameters),\n        \"spider_errors\": \"true\",\n        \"file_name\": fileName,\n    }\n")
		b.WriteString("    return s.client.doForm(ctx, s.client.scraperBaseURL, \"/builder?platform=1\", payload)\n")
	}
	b.WriteString("}\n\n")
}

func defaultSpiderParamDefaults(t Tool) []Param {
	names := map[string]bool{}
	switch t.Name {
	case "glassdoor_company":
		names["url"] = true
	case "glassdoor_joblistings":
		names["url"] = true
	case "youtube_video_post":
		names["url"] = true
		names["order_by"] = true
		names["start_index"] = true
		names["num_of_posts"] = true
	}
	if len(names) == 0 {
		return nil
	}
	var out []Param
	for _, p := range t.Params {
		if names[p.Name] && p.Default != "" {
			out = append(out, p)
		}
	}
	return out
}

func extraRequiredParams(toolName string) []string {
	switch toolName {
	case "google_flights":
		return []string{"departure_id", "arrival_id", "outbound_date"}
	case "google_hotels":
		return []string{"check_in_date", "check_out_date"}
	}
	return nil
}

func writeSpecs(b *strings.Builder, tools []Tool) {
	b.WriteString("type ToolParamSpec struct { Name string; Type string; Required bool; Default string; Min string; Max string }\n")
	b.WriteString("type ToolSpec struct { Name string; Group string; Method string; ServiceKind string; Engine string; SpiderName string; SpiderID string; Params []ToolParamSpec }\n")
	b.WriteString("var ToolSpecs = []ToolSpec{\n")
	for _, t := range tools {
		fmt.Fprintf(b, "    {Name:%s, Group:%s, Method:%s, ServiceKind:%s, Engine:%s, SpiderName:%s, SpiderID:%s, Params: []ToolParamSpec{\n", strconv.Quote(t.Name), strconv.Quote(t.Group), strconv.Quote(t.Method), strconv.Quote(t.ServiceKind), strconv.Quote(t.Engine), strconv.Quote(t.SpiderName), strconv.Quote(t.SpiderID))
		for _, p := range t.Params {
			fmt.Fprintf(b, "        {Name:%s, Type:%s, Required:%v, Default:%s, Min:%s, Max:%s},\n", strconv.Quote(p.Name), strconv.Quote(p.Type), p.Required, strconv.Quote(p.Default), strconv.Quote(p.Min), strconv.Quote(p.Max))
		}
		b.WriteString("    }},\n")
	}
	b.WriteString("}\n")
}

func uniqueGroups(tools []Tool) []string {
	m := map[string]bool{}
	for _, t := range tools {
		m[t.Group] = true
	}
	var groups []string
	for g := range m {
		groups = append(groups, g)
	}
	sort.Strings(groups)
	return groups
}

func readme(tools []Tool) string {
	return fmt.Sprintf(`# Dataify Go SDK

Generated Go SDK for Dataify tools.

This SDK exposes Go methods for Dataify API Key based runtime tools. SERP, Web Unlocker, and Scraper Builder tools call public Dataify HTTP APIs.

## Install

`+"```bash"+`
go get github.com/dataify-server/dataify-sdk-go
`+"```"+`

## Usage

`+"```go"+`
client := dataify.NewClient("YOUR_DATAIFY_API_TOKEN")
resp, err := client.Google.Search(ctx, dataify.GoogleSearchRequest{Q: "coffee", JSON: "1"})
`+"```"+`

## Tool Count

%d API Key runtime tools are documented in docs/tools.md and docs/tools.yaml.
`, len(tools))
}

func toolsYAML(tools []Tool) string {
	var b strings.Builder
	b.WriteString("tools:\n")
	for _, t := range tools {
		fmt.Fprintf(&b, "  - name: %s\n", yamlQuote(t.Name))
		fmt.Fprintf(&b, "    group: %s\n    method: %s\n    request_type: %s\n    service_kind: %s\n", yamlQuote(t.Group), yamlQuote(t.Method), yamlQuote(t.RequestType), yamlQuote(t.ServiceKind))
		if t.Engine != "" {
			fmt.Fprintf(&b, "    engine: %s\n", yamlQuote(t.Engine))
		}
		if t.SpiderName != "" {
			fmt.Fprintf(&b, "    spider_name: %s\n", yamlQuote(t.SpiderName))
		}
		if t.SpiderID != "" {
			fmt.Fprintf(&b, "    spider_id: %s\n", yamlQuote(t.SpiderID))
		}
		b.WriteString("    params:\n")
		for _, p := range t.Params {
			fmt.Fprintf(&b, "      - name: %s\n        go_name: %s\n        type: %s\n        required: %v\n", yamlQuote(p.Name), yamlQuote(p.GoName), yamlQuote(p.Type), p.Required)
			if p.Default != "" {
				fmt.Fprintf(&b, "        default: %s\n", yamlQuote(p.Default))
			}
			if p.Min != "" {
				fmt.Fprintf(&b, "        min: %s\n", yamlQuote(p.Min))
			}
			if p.Max != "" {
				fmt.Fprintf(&b, "        max: %s\n", yamlQuote(p.Max))
			}
		}
	}
	return b.String()
}

func toolsMarkdown(tools []Tool) string {
	var b strings.Builder
	b.WriteString("# Dataify Tool Manifest\n\n")
	b.WriteString("This document is generated from the MCP tool source. Files are UTF-8.\n\n")
	for _, t := range tools {
		fmt.Fprintf(&b, "## `%s`\n\n", t.Name)
		fmt.Fprintf(&b, "- SDK: `%s.%s(ctx, dataify.%s{...})`\n", t.Group, t.Method, t.RequestType)
		fmt.Fprintf(&b, "- Service kind: `%s`\n", t.ServiceKind)
		if t.Engine != "" {
			fmt.Fprintf(&b, "- Engine: `%s`\n", t.Engine)
		}
		if t.SpiderName != "" {
			fmt.Fprintf(&b, "- Spider name: `%s`\n", t.SpiderName)
		}
		if t.SpiderID != "" {
			fmt.Fprintf(&b, "- Default spider ID: `%s`\n", t.SpiderID)
		}
		b.WriteString("\n| Parameter | Go Field | Type | Required | Default | Min | Max |\n")
		b.WriteString("| --- | --- | --- | --- | --- | --- | --- |\n")
		for _, p := range t.Params {
			fmt.Fprintf(&b, "| `%s` | `%s` | `%s` | `%v` | `%s` | `%s` | `%s` |\n", escapeMD(p.Name), p.GoName, p.Type, p.Required, escapeMD(p.Default), escapeMD(p.Min), escapeMD(p.Max))
		}
		b.WriteString("\n")
	}
	return b.String()
}

func yamlQuote(s string) string { return strconv.Quote(s) }
func escapeMD(s string) string  { return strings.ReplaceAll(s, "|", "\\|") }

func writeText(path, content string) {
	must(os.MkdirAll(filepath.Dir(path), 0755))
	must(os.WriteFile(path, []byte(content), 0644))
}

func writeGo(path, content string) {
	formatted, err := format.Source([]byte(content))
	if err != nil {
		var buf bytes.Buffer
		fmt.Fprintf(&buf, "// gofmt failed: %v\n", err)
		buf.WriteString(content)
		writeText(path, buf.String())
		return
	}
	writeText(path, string(formatted))
}
