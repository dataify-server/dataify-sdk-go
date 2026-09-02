package dataify

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"
)

const (
	defaultSERPBaseURL        = "https://scraperapi.dataify.com"
	defaultScraperBaseURL     = "https://scraperapi.dataify.com"
	defaultWebUnlockerBaseURL = "https://webunlocker.dataify.com"
	standardTokenEnv          = "DATAIFY_API_TOKEN"
	legacyTokenEnv            = "DATAIFY_TOKEN"
	legacyAPIKeyEnv           = "DATAIFY_API_KEY"
)

var (
	ErrMissingToken                      = errors.New("dataify: missing API token; set DATAIFY_API_TOKEN (legacy DATAIFY_TOKEN and DATAIFY_API_KEY are also supported)")
	legacyTokenEnvWarningOnce            = &sync.Once{}
	legacyAPIKeyEnvWarningOnce           = &sync.Once{}
	legacyTokenWarningWriter   io.Writer = os.Stderr
)

type APIError struct {
	StatusCode int
	Body       []byte
}

func (e *APIError) Error() string {
	return fmt.Sprintf("dataify: api returned HTTP %d: %s", e.StatusCode, truncate(e.Body))
}

type Client struct {
	token              string
	httpClient         *http.Client
	serpBaseURL        string
	scraperBaseURL     string
	webUnlockerBaseURL string
	Scraper            *ScraperService
	Airbnb             *AirbnbService
	Amazon             *AmazonService
	Bing               *BingService
	Booking            *BookingService
	Crunchbase         *CrunchbaseService
	DuckDuckGo         *DuckDuckGoService
	EBay               *EBayService
	Facebook           *FacebookService
	GitHub             *GitHubService
	Glassdoor          *GlassdoorService
	Google             *GoogleService
	Indeed             *IndeedService
	Instagram          *InstagramService
	LinkedIn           *LinkedInService
	Reddit             *RedditService
	TikTok             *TikTokService
	Twitter            *TwitterService
	Walmart            *WalmartService
	WebUnlocker        *WebUnlockerService
	Yandex             *YandexService
	YouTube            *YouTubeService
	Zillow             *ZillowService
}

type Option func(*Client)

func WithHTTPClient(httpClient *http.Client) Option {
	return func(c *Client) {
		if httpClient != nil {
			c.httpClient = httpClient
		}
	}
}
func WithSERPBaseURL(baseURL string) Option {
	return func(c *Client) { c.serpBaseURL = strings.TrimRight(strings.TrimSpace(baseURL), "/") }
}
func WithScraperBaseURL(baseURL string) Option {
	return func(c *Client) { c.scraperBaseURL = strings.TrimRight(strings.TrimSpace(baseURL), "/") }
}
func WithWebUnlockerBaseURL(baseURL string) Option {
	return func(c *Client) { c.webUnlockerBaseURL = strings.TrimRight(strings.TrimSpace(baseURL), "/") }
}
func WithTimeout(timeout time.Duration) Option {
	return func(c *Client) {
		if timeout > 0 {
			c.httpClient.Timeout = timeout
		}
	}
}

func NewClient(token string, opts ...Option) *Client {
	resolvedToken := strings.TrimSpace(token)
	if resolvedToken == "" {
		resolvedToken = TokenFromEnv()
	} else {
		warnIfLegacyEnvironmentToken(resolvedToken)
	}

	c := &Client{
		token:              resolvedToken,
		httpClient:         &http.Client{Timeout: 120 * time.Second},
		serpBaseURL:        defaultSERPBaseURL,
		scraperBaseURL:     defaultScraperBaseURL,
		webUnlockerBaseURL: defaultWebUnlockerBaseURL,
	}
	for _, opt := range opts {
		if opt != nil {
			opt(c)
		}
	}
	c.Scraper = &ScraperService{client: c}
	c.Airbnb = &AirbnbService{client: c}
	c.Amazon = &AmazonService{client: c}
	c.Bing = &BingService{client: c}
	c.Booking = &BookingService{client: c}
	c.Crunchbase = &CrunchbaseService{client: c}
	c.DuckDuckGo = &DuckDuckGoService{client: c}
	c.EBay = &EBayService{client: c}
	c.Facebook = &FacebookService{client: c}
	c.GitHub = &GitHubService{client: c}
	c.Glassdoor = &GlassdoorService{client: c}
	c.Google = &GoogleService{client: c}
	c.Indeed = &IndeedService{client: c}
	c.Instagram = &InstagramService{client: c}
	c.LinkedIn = &LinkedInService{client: c}
	c.Reddit = &RedditService{client: c}
	c.TikTok = &TikTokService{client: c}
	c.Twitter = &TwitterService{client: c}
	c.Walmart = &WalmartService{client: c}
	c.WebUnlocker = &WebUnlockerService{client: c}
	c.Yandex = &YandexService{client: c}
	c.YouTube = &YouTubeService{client: c}
	c.Zillow = &ZillowService{client: c}
	return c
}

// TokenFromEnv returns the first configured Dataify credential from the standard
// environment variable and its legacy compatibility aliases.
func TokenFromEnv() string {
	if token := environmentToken(standardTokenEnv); token != "" {
		return token
	}
	if token := environmentToken(legacyTokenEnv); token != "" {
		warnLegacyTokenEnvironment(legacyTokenEnv)
		return token
	}
	if token := environmentToken(legacyAPIKeyEnv); token != "" {
		warnLegacyTokenEnvironment(legacyAPIKeyEnv)
		return token
	}
	return ""
}

func environmentToken(name string) string {
	return strings.TrimSpace(os.Getenv(name))
}

func warnIfLegacyEnvironmentToken(token string) {
	if token == environmentToken(legacyTokenEnv) && token != "" {
		warnLegacyTokenEnvironment(legacyTokenEnv)
		return
	}
	if token == environmentToken(legacyAPIKeyEnv) && token != "" {
		warnLegacyTokenEnvironment(legacyAPIKeyEnv)
	}
}

func warnLegacyTokenEnvironment(name string) {
	var once *sync.Once
	switch name {
	case legacyTokenEnv:
		once = legacyTokenEnvWarningOnce
	case legacyAPIKeyEnv:
		once = legacyAPIKeyEnvWarningOnce
	default:
		return
	}
	once.Do(func() {
		fmt.Fprintf(legacyTokenWarningWriter, "dataify: %s 已兼容读取，建议迁移至 DATAIFY_API_TOKEN。\n", name)
	})
}

func (c *Client) doForm(ctx context.Context, baseURL, path string, values map[string]string) (*RawResponse, error) {
	if c.token == "" {
		return nil, ErrMissingToken
	}
	form := url.Values{}
	for k, v := range values {
		if strings.TrimSpace(v) != "" {
			form.Set(k, v)
		}
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, strings.TrimRight(baseURL, "/")+path, strings.NewReader(form.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", bearer(c.token))
	return c.do(req)
}

func (c *Client) doJSON(ctx context.Context, baseURL, path string, payload map[string]string) (*RawResponse, error) {
	if c.token == "" {
		return nil, ErrMissingToken
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, strings.TrimRight(baseURL, "/")+path, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", bearer(c.token))
	return c.do(req)
}

func (c *Client) do(req *http.Request) (*RawResponse, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, &APIError{StatusCode: resp.StatusCode, Body: body}
	}
	return NewRawResponse(body)
}

func bearer(token string) string {
	token = strings.TrimSpace(token)
	if token == "" || strings.HasPrefix(strings.ToLower(token), "bearer ") {
		return token
	}
	return "Bearer " + token
}

func addStringParam(values map[string]string, key string, value string) {
	if strings.TrimSpace(value) != "" {
		values[key] = strings.TrimSpace(value)
	}
}
func addIntParam(values map[string]string, key string, value int) {
	if value != 0 {
		values[key] = fmt.Sprint(value)
	}
}
func defaultParam(values map[string]string, key string, value string) {
	if strings.TrimSpace(values[key]) == "" && value != "" {
		values[key] = value
	}
}
func requireParam(values map[string]string, key string) error {
	if strings.TrimSpace(values[key]) == "" {
		return fmt.Errorf("dataify: missing required parameter %s", key)
	}
	return nil
}
func truncate(body []byte) string {
	s := strings.TrimSpace(string(body))
	if len(s) > 500 {
		return s[:500] + "..."
	}
	return s
}
