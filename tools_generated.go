package dataify

import (
	"context"
	"encoding/json"
)

type AirbnbService struct{ client *Client }

type AmazonService struct{ client *Client }

type BingService struct{ client *Client }

type BookingService struct{ client *Client }

type CrunchbaseService struct{ client *Client }

type DuckDuckGoService struct{ client *Client }

type EBayService struct{ client *Client }

type FacebookService struct{ client *Client }

type GitHubService struct{ client *Client }

type GlassdoorService struct{ client *Client }

type GoogleService struct{ client *Client }

type IndeedService struct{ client *Client }

type InstagramService struct{ client *Client }

type LinkedInService struct{ client *Client }

type RedditService struct{ client *Client }

type TikTokService struct{ client *Client }

type TwitterService struct{ client *Client }

type WalmartService struct{ client *Client }

type WebUnlockerService struct{ client *Client }

type YandexService struct{ client *Client }

type YouTubeService struct{ client *Client }

type ZillowService struct{ client *Client }

type AirbnbProductRequest struct {
	Searchurl string `json:"searchurl,omitempty"`
	Country   string `json:"country,omitempty"`
	FileName  string `json:"file_name,omitempty"`
}

func (r AirbnbProductRequest) params() map[string]string {
	values := map[string]string{}
	addStringParam(values, "searchurl", r.Searchurl)
	addStringParam(values, "country", r.Country)
	addStringParam(values, "file_name", r.FileName)
	return values
}

func (s *AirbnbService) Product(ctx context.Context, req AirbnbProductRequest) (*RawResponse, error) {
	values := req.params()
	defaultParam(values, "searchurl", "https://www.airbnb.com/s/Greece/homes?query=Greece&refinement_paths%5B%5D=%2Fhomes&place_id=ChIJY2xxEcdKWxMRHS2a3HUXOjY&flexible_trip_lengths%5B%5D=one_week&monthly_start_date=2025-03-01&monthly_length=3&monthly_end_date=2025-06-01&search_mode=regular_search&price_filter_input_type=0&channel=EXPLORE&date_picker_type=calendar&source=structured_search_input_header&search_type=filter_change&price_filter_num_nights=5&flexible_date_search_filter_type=1")
	defaultParam(values, "country", "HK")
	defaultParam(values, "file_name", "{{TasksID}}")
	if err := requireParam(values, "searchurl"); err != nil {
		return nil, err
	}
	spiderID := "airbnb_product_by-searchurl"
	spiderName := "airbnb.com"
	fileName := values["file_name"]
	if fileName == "" {
		fileName = "{{TasksID}}"
	}
	delete(values, "spider_id")
	delete(values, "file_name")
	spiderParameters, err := json.Marshal([]map[string]string{values})
	if err != nil {
		return nil, err
	}
	payload := map[string]string{
		"spider_name":       spiderName,
		"spider_id":         spiderID,
		"spider_parameters": string(spiderParameters),
		"spider_errors":     "true",
		"file_name":         fileName,
	}
	return s.client.doForm(ctx, s.client.scraperBaseURL, "/builder?platform=1", payload)
}

type AmazonCommentRequest struct {
	URL      string `json:"url,omitempty"`
	FileName string `json:"file_name,omitempty"`
}

func (r AmazonCommentRequest) params() map[string]string {
	values := map[string]string{}
	addStringParam(values, "url", r.URL)
	addStringParam(values, "file_name", r.FileName)
	return values
}

func (s *AmazonService) Comment(ctx context.Context, req AmazonCommentRequest) (*RawResponse, error) {
	values := req.params()
	defaultParam(values, "url", "https://www.amazon.com/HISDERN-Checkered-Handkerchief-Classic-Necktie/dp/B0BRXPR726")
	defaultParam(values, "file_name", "{{TasksID}}")
	if err := requireParam(values, "url"); err != nil {
		return nil, err
	}
	spiderID := "amazon_comment_by-url"
	spiderName := "amazon.com"
	fileName := values["file_name"]
	if fileName == "" {
		fileName = "{{TasksID}}"
	}
	delete(values, "spider_id")
	delete(values, "file_name")
	spiderParameters, err := json.Marshal([]map[string]string{values})
	if err != nil {
		return nil, err
	}
	payload := map[string]string{
		"spider_name":       spiderName,
		"spider_id":         spiderID,
		"spider_parameters": string(spiderParameters),
		"spider_errors":     "true",
		"file_name":         fileName,
	}
	return s.client.doForm(ctx, s.client.scraperBaseURL, "/builder?platform=1", payload)
}

type AmazonGlobalProductRequest struct {
	SpiderID     string `json:"spider_id,omitempty"`
	URL          string `json:"url,omitempty"`
	Maximum      string `json:"maximum,omitempty"`
	SortBy       string `json:"sort_by,omitempty"`
	GetSponsored string `json:"get_sponsored,omitempty"`
	Keyword      string `json:"keyword,omitempty"`
	Domain       string `json:"domain,omitempty"`
	Brands       string `json:"brands,omitempty"`
	PageTurning  string `json:"page_turning,omitempty"`
	LowestPrice  string `json:"lowest_price,omitempty"`
	HighestPrice string `json:"highest_price,omitempty"`
	FileName     string `json:"file_name,omitempty"`
}

func (r AmazonGlobalProductRequest) params() map[string]string {
	values := map[string]string{}
	addStringParam(values, "spider_id", r.SpiderID)
	addStringParam(values, "url", r.URL)
	addStringParam(values, "maximum", r.Maximum)
	addStringParam(values, "sort_by", r.SortBy)
	addStringParam(values, "get_sponsored", r.GetSponsored)
	addStringParam(values, "keyword", r.Keyword)
	addStringParam(values, "domain", r.Domain)
	addStringParam(values, "brands", r.Brands)
	addStringParam(values, "page_turning", r.PageTurning)
	addStringParam(values, "lowest_price", r.LowestPrice)
	addStringParam(values, "highest_price", r.HighestPrice)
	addStringParam(values, "file_name", r.FileName)
	return values
}

func (s *AmazonService) GlobalProduct(ctx context.Context, req AmazonGlobalProductRequest) (*RawResponse, error) {
	values := req.params()
	defaultParam(values, "spider_id", "amazon_global-product_by-url")
	defaultParam(values, "file_name", "{{TasksID}}")
	if err := requireParam(values, "spider_id"); err != nil {
		return nil, err
	}
	spiderID := values["spider_id"]
	if spiderID == "" {
		spiderID = "amazon_global-product_by-url"
	}
	spiderName := "amazon.com"
	fileName := values["file_name"]
	if fileName == "" {
		fileName = "{{TasksID}}"
	}
	delete(values, "spider_id")
	delete(values, "file_name")
	spiderParameters, err := json.Marshal([]map[string]string{values})
	if err != nil {
		return nil, err
	}
	payload := map[string]string{
		"spider_name":       spiderName,
		"spider_id":         spiderID,
		"spider_parameters": string(spiderParameters),
		"spider_errors":     "true",
		"file_name":         fileName,
	}
	return s.client.doForm(ctx, s.client.scraperBaseURL, "/builder?platform=1", payload)
}

type AmazonProductRequest struct {
	SpiderID             string `json:"spider_id,omitempty"`
	Asin                 string `json:"asin,omitempty"`
	URL                  string `json:"url,omitempty"`
	CategoryURL          string `json:"category_url,omitempty"`
	Keyword              string `json:"keyword,omitempty"`
	PageTurning          string `json:"page_turning,omitempty"`
	LowestPrice          string `json:"lowest_price,omitempty"`
	HighestPrice         string `json:"highest_price,omitempty"`
	SortBy               string `json:"sort_by,omitempty"`
	CollectSubcategories string `json:"collect_subcategories,omitempty"`
	ZipCode              string `json:"zip_code,omitempty"`
	FileName             string `json:"file_name,omitempty"`
}

func (r AmazonProductRequest) params() map[string]string {
	values := map[string]string{}
	addStringParam(values, "spider_id", r.SpiderID)
	addStringParam(values, "asin", r.Asin)
	addStringParam(values, "url", r.URL)
	addStringParam(values, "category_url", r.CategoryURL)
	addStringParam(values, "keyword", r.Keyword)
	addStringParam(values, "page_turning", r.PageTurning)
	addStringParam(values, "lowest_price", r.LowestPrice)
	addStringParam(values, "highest_price", r.HighestPrice)
	addStringParam(values, "sort_by", r.SortBy)
	addStringParam(values, "collect_subcategories", r.CollectSubcategories)
	addStringParam(values, "zip_code", r.ZipCode)
	addStringParam(values, "file_name", r.FileName)
	return values
}

func (s *AmazonService) Product(ctx context.Context, req AmazonProductRequest) (*RawResponse, error) {
	values := req.params()
	defaultParam(values, "spider_id", "amazon_product_by-asin")
	defaultParam(values, "file_name", "{{TasksID}}")
	if err := requireParam(values, "spider_id"); err != nil {
		return nil, err
	}
	spiderID := values["spider_id"]
	if spiderID == "" {
		spiderID = "amazon_product_by-asin"
	}
	spiderName := "amazon.com"
	fileName := values["file_name"]
	if fileName == "" {
		fileName = "{{TasksID}}"
	}
	delete(values, "spider_id")
	delete(values, "file_name")
	spiderParameters, err := json.Marshal([]map[string]string{values})
	if err != nil {
		return nil, err
	}
	payload := map[string]string{
		"spider_name":       spiderName,
		"spider_id":         spiderID,
		"spider_parameters": string(spiderParameters),
		"spider_errors":     "true",
		"file_name":         fileName,
	}
	return s.client.doForm(ctx, s.client.scraperBaseURL, "/builder?platform=1", payload)
}

type AmazonProductListRequest struct {
	Keyword     string `json:"keyword,omitempty"`
	Domain      string `json:"domain,omitempty"`
	PageTurning string `json:"page_turning,omitempty"`
	FileName    string `json:"file_name,omitempty"`
}

func (r AmazonProductListRequest) params() map[string]string {
	values := map[string]string{}
	addStringParam(values, "keyword", r.Keyword)
	addStringParam(values, "domain", r.Domain)
	addStringParam(values, "page_turning", r.PageTurning)
	addStringParam(values, "file_name", r.FileName)
	return values
}

func (s *AmazonService) ProductList(ctx context.Context, req AmazonProductListRequest) (*RawResponse, error) {
	values := req.params()
	defaultParam(values, "keyword", "X-box")
	defaultParam(values, "domain", "https://www.amazon.com/")
	defaultParam(values, "page_turning", "1")
	defaultParam(values, "file_name", "{{TasksID}}")
	if err := requireParam(values, "keyword"); err != nil {
		return nil, err
	}
	if err := requireParam(values, "domain"); err != nil {
		return nil, err
	}
	spiderID := "amazon_product-list_by-keywords-domain"
	spiderName := "amazon.com"
	fileName := values["file_name"]
	if fileName == "" {
		fileName = "{{TasksID}}"
	}
	delete(values, "spider_id")
	delete(values, "file_name")
	spiderParameters, err := json.Marshal([]map[string]string{values})
	if err != nil {
		return nil, err
	}
	payload := map[string]string{
		"spider_name":       spiderName,
		"spider_id":         spiderID,
		"spider_parameters": string(spiderParameters),
		"spider_errors":     "true",
		"file_name":         fileName,
	}
	return s.client.doForm(ctx, s.client.scraperBaseURL, "/builder?platform=1", payload)
}

type AmazonSellerRequest struct {
	URL      string `json:"url,omitempty"`
	FileName string `json:"file_name,omitempty"`
}

func (r AmazonSellerRequest) params() map[string]string {
	values := map[string]string{}
	addStringParam(values, "url", r.URL)
	addStringParam(values, "file_name", r.FileName)
	return values
}

func (s *AmazonService) Seller(ctx context.Context, req AmazonSellerRequest) (*RawResponse, error) {
	values := req.params()
	defaultParam(values, "url", "https://www.amazon.com/sp?ie=UTF8&seller=ADZ7LD48GVFQJ&asin=B07H56J7K1&ref_=dp_merchant_link&isAmazonFulfilled=1")
	defaultParam(values, "file_name", "{{TasksID}}")
	if err := requireParam(values, "url"); err != nil {
		return nil, err
	}
	spiderID := "amazon_seller_by-url"
	spiderName := "amazon.com"
	fileName := values["file_name"]
	if fileName == "" {
		fileName = "{{TasksID}}"
	}
	delete(values, "spider_id")
	delete(values, "file_name")
	spiderParameters, err := json.Marshal([]map[string]string{values})
	if err != nil {
		return nil, err
	}
	payload := map[string]string{
		"spider_name":       spiderName,
		"spider_id":         spiderID,
		"spider_parameters": string(spiderParameters),
		"spider_errors":     "true",
		"file_name":         fileName,
	}
	return s.client.doForm(ctx, s.client.scraperBaseURL, "/builder?platform=1", payload)
}

type BingImagesRequest struct {
	Q         string `json:"q,omitempty"`
	JSON      string `json:"json,omitempty"`
	Mkt       string `json:"mkt,omitempty"`
	Cc        string `json:"cc,omitempty"`
	First     string `json:"first,omitempty"`
	Count     string `json:"count,omitempty"`
	Imagesize string `json:"imagesize,omitempty"`
	Color2    string `json:"color2,omitempty"`
	Photo     string `json:"photo,omitempty"`
	Aspect    string `json:"aspect,omitempty"`
	Face      string `json:"face,omitempty"`
	Age       string `json:"age,omitempty"`
	License   string `json:"license,omitempty"`
	NoCache   string `json:"no_cache,omitempty"`
}

func (r BingImagesRequest) params() map[string]string {
	values := map[string]string{}
	addStringParam(values, "q", r.Q)
	addStringParam(values, "json", r.JSON)
	addStringParam(values, "mkt", r.Mkt)
	addStringParam(values, "cc", r.Cc)
	addStringParam(values, "first", r.First)
	addStringParam(values, "count", r.Count)
	addStringParam(values, "imagesize", r.Imagesize)
	addStringParam(values, "color2", r.Color2)
	addStringParam(values, "photo", r.Photo)
	addStringParam(values, "aspect", r.Aspect)
	addStringParam(values, "face", r.Face)
	addStringParam(values, "age", r.Age)
	addStringParam(values, "license", r.License)
	addStringParam(values, "no_cache", r.NoCache)
	return values
}

func (s *BingService) Images(ctx context.Context, req BingImagesRequest) (*RawResponse, error) {
	values := req.params()
	defaultParam(values, "q", "Pizza")
	defaultParam(values, "json", "1")
	defaultParam(values, "first", "0")
	defaultParam(values, "no_cache", "false")
	if err := requireParam(values, "q"); err != nil {
		return nil, err
	}
	if err := requireParam(values, "json"); err != nil {
		return nil, err
	}
	values["engine"] = "bing_images"
	return s.client.doForm(ctx, s.client.serpBaseURL, "/request", values)
}

type BingMapsRequest struct {
	Q       string `json:"q,omitempty"`
	JSON    string `json:"json,omitempty"`
	Cp      string `json:"cp,omitempty"`
	Setlang string `json:"setlang,omitempty"`
	PlaceID string `json:"place_id,omitempty"`
	First   string `json:"first,omitempty"`
	Count   string `json:"count,omitempty"`
	NoCache string `json:"no_cache,omitempty"`
}

func (r BingMapsRequest) params() map[string]string {
	values := map[string]string{}
	addStringParam(values, "q", r.Q)
	addStringParam(values, "json", r.JSON)
	addStringParam(values, "cp", r.Cp)
	addStringParam(values, "setlang", r.Setlang)
	addStringParam(values, "place_id", r.PlaceID)
	addStringParam(values, "first", r.First)
	addStringParam(values, "count", r.Count)
	addStringParam(values, "no_cache", r.NoCache)
	return values
}

func (s *BingService) Maps(ctx context.Context, req BingMapsRequest) (*RawResponse, error) {
	values := req.params()
	defaultParam(values, "json", "1")
	defaultParam(values, "first", "0")
	defaultParam(values, "no_cache", "false")
	if err := requireParam(values, "q"); err != nil {
		return nil, err
	}
	if err := requireParam(values, "json"); err != nil {
		return nil, err
	}
	values["engine"] = "bing_maps"
	return s.client.doForm(ctx, s.client.serpBaseURL, "/request", values)
}

type BingNewsRequest struct {
	Q          string `json:"q,omitempty"`
	JSON       string `json:"json,omitempty"`
	Mkt        string `json:"mkt,omitempty"`
	Cc         string `json:"cc,omitempty"`
	First      string `json:"first,omitempty"`
	Count      string `json:"count,omitempty"`
	Qft        string `json:"qft,omitempty"`
	Safesearch string `json:"safeSearch,omitempty"`
	NoCache    string `json:"no_cache,omitempty"`
}

func (r BingNewsRequest) params() map[string]string {
	values := map[string]string{}
	addStringParam(values, "q", r.Q)
	addStringParam(values, "json", r.JSON)
	addStringParam(values, "mkt", r.Mkt)
	addStringParam(values, "cc", r.Cc)
	addStringParam(values, "first", r.First)
	addStringParam(values, "count", r.Count)
	addStringParam(values, "qft", r.Qft)
	addStringParam(values, "safeSearch", r.Safesearch)
	addStringParam(values, "no_cache", r.NoCache)
	return values
}

func (s *BingService) News(ctx context.Context, req BingNewsRequest) (*RawResponse, error) {
	values := req.params()
	defaultParam(values, "q", "Pizza")
	defaultParam(values, "json", "1")
	defaultParam(values, "first", "0")
	defaultParam(values, "no_cache", "false")
	if err := requireParam(values, "q"); err != nil {
		return nil, err
	}
	if err := requireParam(values, "json"); err != nil {
		return nil, err
	}
	values["engine"] = "bing_news"
	return s.client.doForm(ctx, s.client.serpBaseURL, "/request", values)
}

type BingSearchRequest struct {
	Q          string `json:"q,omitempty"`
	JSON       string `json:"json,omitempty"`
	Location   string `json:"location,omitempty"`
	Lat        string `json:"lat,omitempty"`
	Lon        string `json:"lon,omitempty"`
	Mkt        string `json:"mkt,omitempty"`
	Cc         string `json:"cc,omitempty"`
	First      string `json:"first,omitempty"`
	Safesearch string `json:"safeSearch,omitempty"`
	Filters    string `json:"filters,omitempty"`
	NoCache    string `json:"no_cache,omitempty"`
}

func (r BingSearchRequest) params() map[string]string {
	values := map[string]string{}
	addStringParam(values, "q", r.Q)
	addStringParam(values, "json", r.JSON)
	addStringParam(values, "location", r.Location)
	addStringParam(values, "lat", r.Lat)
	addStringParam(values, "lon", r.Lon)
	addStringParam(values, "mkt", r.Mkt)
	addStringParam(values, "cc", r.Cc)
	addStringParam(values, "first", r.First)
	addStringParam(values, "safeSearch", r.Safesearch)
	addStringParam(values, "filters", r.Filters)
	addStringParam(values, "no_cache", r.NoCache)
	return values
}

func (s *BingService) Search(ctx context.Context, req BingSearchRequest) (*RawResponse, error) {
	values := req.params()
	defaultParam(values, "q", "Pizza")
	defaultParam(values, "json", "1")
	defaultParam(values, "first", "0")
	defaultParam(values, "no_cache", "false")
	if err := requireParam(values, "q"); err != nil {
		return nil, err
	}
	if err := requireParam(values, "json"); err != nil {
		return nil, err
	}
	values["engine"] = "bing"
	return s.client.doForm(ctx, s.client.serpBaseURL, "/request", values)
}

type BingShoppingRequest struct {
	Q       string `json:"q,omitempty"`
	JSON    string `json:"json,omitempty"`
	Mkt     string `json:"mkt,omitempty"`
	Cc      string `json:"cc,omitempty"`
	Efirst  string `json:"efirst,omitempty"`
	Filters string `json:"filters,omitempty"`
	NoCache string `json:"no_cache,omitempty"`
}

func (r BingShoppingRequest) params() map[string]string {
	values := map[string]string{}
	addStringParam(values, "q", r.Q)
	addStringParam(values, "json", r.JSON)
	addStringParam(values, "mkt", r.Mkt)
	addStringParam(values, "cc", r.Cc)
	addStringParam(values, "efirst", r.Efirst)
	addStringParam(values, "filters", r.Filters)
	addStringParam(values, "no_cache", r.NoCache)
	return values
}

func (s *BingService) Shopping(ctx context.Context, req BingShoppingRequest) (*RawResponse, error) {
	values := req.params()
	defaultParam(values, "q", "Pizza")
	defaultParam(values, "json", "1")
	defaultParam(values, "no_cache", "false")
	if err := requireParam(values, "q"); err != nil {
		return nil, err
	}
	if err := requireParam(values, "json"); err != nil {
		return nil, err
	}
	values["engine"] = "bing_shopping"
	return s.client.doForm(ctx, s.client.serpBaseURL, "/request", values)
}

type BingVideosRequest struct {
	Q          string `json:"q,omitempty"`
	JSON       string `json:"json,omitempty"`
	Mkt        string `json:"mkt,omitempty"`
	Cc         string `json:"cc,omitempty"`
	Setlang    string `json:"setlang,omitempty"`
	First      string `json:"first,omitempty"`
	Length     string `json:"length,omitempty"`
	Date       string `json:"date,omitempty"`
	Resolution string `json:"resolution,omitempty"`
	SourceSite string `json:"source_site,omitempty"`
	Price      string `json:"price,omitempty"`
	NoCache    string `json:"no_cache,omitempty"`
}

func (r BingVideosRequest) params() map[string]string {
	values := map[string]string{}
	addStringParam(values, "q", r.Q)
	addStringParam(values, "json", r.JSON)
	addStringParam(values, "mkt", r.Mkt)
	addStringParam(values, "cc", r.Cc)
	addStringParam(values, "setlang", r.Setlang)
	addStringParam(values, "first", r.First)
	addStringParam(values, "length", r.Length)
	addStringParam(values, "date", r.Date)
	addStringParam(values, "resolution", r.Resolution)
	addStringParam(values, "source_site", r.SourceSite)
	addStringParam(values, "price", r.Price)
	addStringParam(values, "no_cache", r.NoCache)
	return values
}

func (s *BingService) Videos(ctx context.Context, req BingVideosRequest) (*RawResponse, error) {
	values := req.params()
	defaultParam(values, "q", "Pizza")
	defaultParam(values, "json", "1")
	defaultParam(values, "first", "0")
	defaultParam(values, "no_cache", "false")
	if err := requireParam(values, "q"); err != nil {
		return nil, err
	}
	if err := requireParam(values, "json"); err != nil {
		return nil, err
	}
	values["engine"] = "bing_videos"
	return s.client.doForm(ctx, s.client.serpBaseURL, "/request", values)
}

type BookingHotellistRequest struct {
	URL      string `json:"url,omitempty"`
	FileName string `json:"file_name,omitempty"`
}

func (r BookingHotellistRequest) params() map[string]string {
	values := map[string]string{}
	addStringParam(values, "url", r.URL)
	addStringParam(values, "file_name", r.FileName)
	return values
}

func (s *BookingService) Hotellist(ctx context.Context, req BookingHotellistRequest) (*RawResponse, error) {
	values := req.params()
	defaultParam(values, "url", "https://www.booking.com/hotel/gb/westlands-of-pitlochry.en-gb.html#tab-main")
	defaultParam(values, "file_name", "{{TasksID}}")
	if err := requireParam(values, "url"); err != nil {
		return nil, err
	}
	spiderID := "booking_hotellist_by-url"
	spiderName := "booking.com"
	fileName := values["file_name"]
	if fileName == "" {
		fileName = "{{TasksID}}"
	}
	delete(values, "spider_id")
	delete(values, "file_name")
	spiderParameters, err := json.Marshal([]map[string]string{values})
	if err != nil {
		return nil, err
	}
	payload := map[string]string{
		"spider_name":       spiderName,
		"spider_id":         spiderID,
		"spider_parameters": string(spiderParameters),
		"spider_errors":     "true",
		"file_name":         fileName,
	}
	return s.client.doForm(ctx, s.client.scraperBaseURL, "/builder?platform=1", payload)
}

type CrunchbaseCompanyRequest struct {
	SpiderID string `json:"spider_id,omitempty"`
	URL      string `json:"url,omitempty"`
	Keyword  string `json:"keyword,omitempty"`
	FileName string `json:"file_name,omitempty"`
}

func (r CrunchbaseCompanyRequest) params() map[string]string {
	values := map[string]string{}
	addStringParam(values, "spider_id", r.SpiderID)
	addStringParam(values, "url", r.URL)
	addStringParam(values, "keyword", r.Keyword)
	addStringParam(values, "file_name", r.FileName)
	return values
}

func (s *CrunchbaseService) Company(ctx context.Context, req CrunchbaseCompanyRequest) (*RawResponse, error) {
	values := req.params()
	defaultParam(values, "spider_id", "crunchbase_company_by-url")
	defaultParam(values, "file_name", "{{TasksID}}")
	if err := requireParam(values, "spider_id"); err != nil {
		return nil, err
	}
	spiderID := values["spider_id"]
	if spiderID == "" {
		spiderID = "crunchbase_company_by-url"
	}
	spiderName := "crunchbase.com"
	fileName := values["file_name"]
	if fileName == "" {
		fileName = "{{TasksID}}"
	}
	delete(values, "spider_id")
	delete(values, "file_name")
	spiderParameters, err := json.Marshal([]map[string]string{values})
	if err != nil {
		return nil, err
	}
	payload := map[string]string{
		"spider_name":       spiderName,
		"spider_id":         spiderID,
		"spider_parameters": string(spiderParameters),
		"spider_errors":     "true",
		"file_name":         fileName,
	}
	return s.client.doForm(ctx, s.client.scraperBaseURL, "/builder?platform=1", payload)
}

type DuckDuckGoSearchRequest struct {
	Q            string `json:"q,omitempty"`
	JSON         string `json:"json,omitempty"`
	Kl           string `json:"kl,omitempty"`
	SearchAssist string `json:"search_assist,omitempty"`
	Safe         string `json:"safe,omitempty"`
	Df           string `json:"df,omitempty"`
	Start        string `json:"start,omitempty"`
	M            string `json:"m,omitempty"`
	NoCache      string `json:"no_cache,omitempty"`
}

func (r DuckDuckGoSearchRequest) params() map[string]string {
	values := map[string]string{}
	addStringParam(values, "q", r.Q)
	addStringParam(values, "json", r.JSON)
	addStringParam(values, "kl", r.Kl)
	addStringParam(values, "search_assist", r.SearchAssist)
	addStringParam(values, "safe", r.Safe)
	addStringParam(values, "df", r.Df)
	addStringParam(values, "start", r.Start)
	addStringParam(values, "m", r.M)
	addStringParam(values, "no_cache", r.NoCache)
	return values
}

func (s *DuckDuckGoService) Search(ctx context.Context, req DuckDuckGoSearchRequest) (*RawResponse, error) {
	values := req.params()
	defaultParam(values, "q", "Pizza")
	defaultParam(values, "json", "1")
	defaultParam(values, "search_assist", "false")
	defaultParam(values, "safe", "-1")
	defaultParam(values, "start", "0")
	defaultParam(values, "m", "10")
	defaultParam(values, "no_cache", "false")
	if err := requireParam(values, "q"); err != nil {
		return nil, err
	}
	if err := requireParam(values, "json"); err != nil {
		return nil, err
	}
	values["engine"] = "duckduckgo"
	return s.client.doForm(ctx, s.client.serpBaseURL, "/request", values)
}

type EBayInfoRequest struct {
	SpiderID string `json:"spider_id,omitempty"`
	URL      string `json:"url,omitempty"`
	Keywords string `json:"keywords,omitempty"`
	Count    string `json:"count,omitempty"`
	FileName string `json:"file_name,omitempty"`
}

func (r EBayInfoRequest) params() map[string]string {
	values := map[string]string{}
	addStringParam(values, "spider_id", r.SpiderID)
	addStringParam(values, "url", r.URL)
	addStringParam(values, "keywords", r.Keywords)
	addStringParam(values, "count", r.Count)
	addStringParam(values, "file_name", r.FileName)
	return values
}

func (s *EBayService) Info(ctx context.Context, req EBayInfoRequest) (*RawResponse, error) {
	values := req.params()
	defaultParam(values, "spider_id", "ebay_ebay_by-url")
	defaultParam(values, "file_name", "{{TasksID}}")
	if err := requireParam(values, "spider_id"); err != nil {
		return nil, err
	}
	spiderID := values["spider_id"]
	if spiderID == "" {
		spiderID = "ebay_ebay_by-url"
	}
	spiderName := "ebay.com"
	fileName := values["file_name"]
	if fileName == "" {
		fileName = "{{TasksID}}"
	}
	delete(values, "spider_id")
	delete(values, "file_name")
	spiderParameters, err := json.Marshal([]map[string]string{values})
	if err != nil {
		return nil, err
	}
	payload := map[string]string{
		"spider_name":       spiderName,
		"spider_id":         spiderID,
		"spider_parameters": string(spiderParameters),
		"spider_errors":     "true",
		"file_name":         fileName,
	}
	return s.client.doForm(ctx, s.client.scraperBaseURL, "/builder?platform=1", payload)
}

type FacebookCommentRequest struct {
	URL           string `json:"url,omitempty"`
	GetAllReplies string `json:"get_all_replies,omitempty"`
	LimitRecords  string `json:"limit_records,omitempty"`
	CommentsSort  string `json:"comments_sort,omitempty"`
	FileName      string `json:"file_name,omitempty"`
}

func (r FacebookCommentRequest) params() map[string]string {
	values := map[string]string{}
	addStringParam(values, "url", r.URL)
	addStringParam(values, "get_all_replies", r.GetAllReplies)
	addStringParam(values, "limit_records", r.LimitRecords)
	addStringParam(values, "comments_sort", r.CommentsSort)
	addStringParam(values, "file_name", r.FileName)
	return values
}

func (s *FacebookService) Comment(ctx context.Context, req FacebookCommentRequest) (*RawResponse, error) {
	values := req.params()
	defaultParam(values, "url", "https://www.facebook.com/share/p/1K6xfHFkrK/")
	defaultParam(values, "get_all_replies", "True")
	defaultParam(values, "limit_records", "10")
	defaultParam(values, "comments_sort", "所有评论")
	defaultParam(values, "file_name", "{{TasksID}}")
	if err := requireParam(values, "url"); err != nil {
		return nil, err
	}
	spiderID := "facebook_comment_by-comments-url"
	spiderName := "facebook.com"
	fileName := values["file_name"]
	if fileName == "" {
		fileName = "{{TasksID}}"
	}
	delete(values, "spider_id")
	delete(values, "file_name")
	spiderParameters, err := json.Marshal([]map[string]string{values})
	if err != nil {
		return nil, err
	}
	payload := map[string]string{
		"spider_name":       spiderName,
		"spider_id":         spiderID,
		"spider_parameters": string(spiderParameters),
		"spider_errors":     "true",
		"file_name":         fileName,
	}
	return s.client.doForm(ctx, s.client.scraperBaseURL, "/builder?platform=1", payload)
}

type FacebookEventRequest struct {
	SpiderID string `json:"spider_id,omitempty"`
	URL      string `json:"url,omitempty"`
	FileName string `json:"file_name,omitempty"`
}

func (r FacebookEventRequest) params() map[string]string {
	values := map[string]string{}
	addStringParam(values, "spider_id", r.SpiderID)
	addStringParam(values, "url", r.URL)
	addStringParam(values, "file_name", r.FileName)
	return values
}

func (s *FacebookService) Event(ctx context.Context, req FacebookEventRequest) (*RawResponse, error) {
	values := req.params()
	defaultParam(values, "spider_id", "facebook_event_by-eventlist-url")
	defaultParam(values, "url", "https://www.facebook.com/nohoclub/events")
	defaultParam(values, "file_name", "{{TasksID}}")
	if err := requireParam(values, "spider_id"); err != nil {
		return nil, err
	}
	if err := requireParam(values, "url"); err != nil {
		return nil, err
	}
	spiderID := values["spider_id"]
	if spiderID == "" {
		spiderID = "facebook_event_by-eventlist-url"
	}
	spiderName := "facebook.com"
	fileName := values["file_name"]
	if fileName == "" {
		fileName = "{{TasksID}}"
	}
	delete(values, "spider_id")
	delete(values, "file_name")
	spiderParameters, err := json.Marshal([]map[string]string{values})
	if err != nil {
		return nil, err
	}
	payload := map[string]string{
		"spider_name":       spiderName,
		"spider_id":         spiderID,
		"spider_parameters": string(spiderParameters),
		"spider_errors":     "true",
		"file_name":         fileName,
	}
	return s.client.doForm(ctx, s.client.scraperBaseURL, "/builder?platform=1", payload)
}

type FacebookPostRequest struct {
	URL      string `json:"url,omitempty"`
	FileName string `json:"file_name,omitempty"`
}

func (r FacebookPostRequest) params() map[string]string {
	values := map[string]string{}
	addStringParam(values, "url", r.URL)
	addStringParam(values, "file_name", r.FileName)
	return values
}

func (s *FacebookService) Post(ctx context.Context, req FacebookPostRequest) (*RawResponse, error) {
	values := req.params()
	defaultParam(values, "url", "https://www.facebook.com/permalink.php?story_fbid=pfbid0gNjZBhqCxSqj9xJS5aygNwqFqNEM2fYbTFKKbsvvGdEfTgFyAYWSckvkEHPqAE7gl&id=61574926580533&rdid=86oaujwNGCCdPLfj#")
	defaultParam(values, "file_name", "{{TasksID}}")
	if err := requireParam(values, "url"); err != nil {
		return nil, err
	}
	spiderID := "facebook_post_by-posts-url"
	spiderName := "facebook.com"
	fileName := values["file_name"]
	if fileName == "" {
		fileName = "{{TasksID}}"
	}
	delete(values, "spider_id")
	delete(values, "file_name")
	spiderParameters, err := json.Marshal([]map[string]string{values})
	if err != nil {
		return nil, err
	}
	payload := map[string]string{
		"spider_name":       spiderName,
		"spider_id":         spiderID,
		"spider_parameters": string(spiderParameters),
		"spider_errors":     "true",
		"file_name":         fileName,
	}
	return s.client.doForm(ctx, s.client.scraperBaseURL, "/builder?platform=1", payload)
}

type FacebookProfileRequest struct {
	URL      string `json:"url,omitempty"`
	FileName string `json:"file_name,omitempty"`
}

func (r FacebookProfileRequest) params() map[string]string {
	values := map[string]string{}
	addStringParam(values, "url", r.URL)
	addStringParam(values, "file_name", r.FileName)
	return values
}

func (s *FacebookService) Profile(ctx context.Context, req FacebookProfileRequest) (*RawResponse, error) {
	values := req.params()
	defaultParam(values, "url", "https://www.facebook.com/MayeMusk")
	defaultParam(values, "file_name", "{{TasksID}}")
	if err := requireParam(values, "url"); err != nil {
		return nil, err
	}
	spiderID := "facebook_profile_by-profiles-url"
	spiderName := "facebook.com"
	fileName := values["file_name"]
	if fileName == "" {
		fileName = "{{TasksID}}"
	}
	delete(values, "spider_id")
	delete(values, "file_name")
	spiderParameters, err := json.Marshal([]map[string]string{values})
	if err != nil {
		return nil, err
	}
	payload := map[string]string{
		"spider_name":       spiderName,
		"spider_id":         spiderID,
		"spider_parameters": string(spiderParameters),
		"spider_errors":     "true",
		"file_name":         fileName,
	}
	return s.client.doForm(ctx, s.client.scraperBaseURL, "/builder?platform=1", payload)
}

type GitHubRepositoryRequest struct {
	SpiderID    string `json:"spider_id,omitempty"`
	RepoURL     string `json:"repo_url,omitempty"`
	SearchURL   string `json:"search_url,omitempty"`
	URL         string `json:"url,omitempty"`
	PageTurning string `json:"page_turning,omitempty"`
	MaxNum      string `json:"max_num,omitempty"`
	FileName    string `json:"file_name,omitempty"`
}

func (r GitHubRepositoryRequest) params() map[string]string {
	values := map[string]string{}
	addStringParam(values, "spider_id", r.SpiderID)
	addStringParam(values, "repo_url", r.RepoURL)
	addStringParam(values, "search_url", r.SearchURL)
	addStringParam(values, "url", r.URL)
	addStringParam(values, "page_turning", r.PageTurning)
	addStringParam(values, "max_num", r.MaxNum)
	addStringParam(values, "file_name", r.FileName)
	return values
}

func (s *GitHubService) Repository(ctx context.Context, req GitHubRepositoryRequest) (*RawResponse, error) {
	values := req.params()
	defaultParam(values, "spider_id", "github_repository_by-repo-url")
	defaultParam(values, "file_name", "{{TasksID}}")
	if err := requireParam(values, "spider_id"); err != nil {
		return nil, err
	}
	spiderID := values["spider_id"]
	if spiderID == "" {
		spiderID = "github_repository_by-repo-url"
	}
	spiderName := "github.com"
	fileName := values["file_name"]
	if fileName == "" {
		fileName = "{{TasksID}}"
	}
	delete(values, "spider_id")
	delete(values, "file_name")
	spiderParameters, err := json.Marshal([]map[string]string{values})
	if err != nil {
		return nil, err
	}
	payload := map[string]string{
		"spider_name":       spiderName,
		"spider_id":         spiderID,
		"spider_parameters": string(spiderParameters),
		"spider_errors":     "true",
		"file_name":         fileName,
	}
	return s.client.doForm(ctx, s.client.scraperBaseURL, "/builder?platform=1", payload)
}

type GlassdoorCompanyRequest struct {
	SpiderID         string `json:"spider_id,omitempty"`
	URL              string `json:"url,omitempty"`
	Location         string `json:"location,omitempty"`
	CompanyName      string `json:"company_name,omitempty"`
	Industries       string `json:"industries,omitempty"`
	JobTitle         string `json:"Job title,omitempty"`
	SearchURL        string `json:"search_url,omitempty"`
	MaxSearchResults string `json:"max_search_results,omitempty"`
	FileName         string `json:"file_name,omitempty"`
}

func (r GlassdoorCompanyRequest) params() map[string]string {
	values := map[string]string{}
	addStringParam(values, "spider_id", r.SpiderID)
	addStringParam(values, "url", r.URL)
	addStringParam(values, "location", r.Location)
	addStringParam(values, "company_name", r.CompanyName)
	addStringParam(values, "industries", r.Industries)
	addStringParam(values, "Job title", r.JobTitle)
	addStringParam(values, "search_url", r.SearchURL)
	addStringParam(values, "max_search_results", r.MaxSearchResults)
	addStringParam(values, "file_name", r.FileName)
	return values
}

func (s *GlassdoorService) Company(ctx context.Context, req GlassdoorCompanyRequest) (*RawResponse, error) {
	values := req.params()
	defaultParam(values, "spider_id", "glassdoor_company_by-url")
	defaultParam(values, "file_name", "{{TasksID}}")
	if values["spider_id"] == "glassdoor_company_by-url" {
		defaultParam(values, "url", "https://www.glassdoor.co.uk/Overview/Working-at-Apple-EI_IE1138.11,16.htm")
	}
	if err := requireParam(values, "spider_id"); err != nil {
		return nil, err
	}
	spiderID := values["spider_id"]
	if spiderID == "" {
		spiderID = "glassdoor_company_by-url"
	}
	spiderName := "glassdoor.com"
	fileName := values["file_name"]
	if fileName == "" {
		fileName = "{{TasksID}}"
	}
	delete(values, "spider_id")
	delete(values, "file_name")
	spiderParameters, err := json.Marshal([]map[string]string{values})
	if err != nil {
		return nil, err
	}
	payload := map[string]string{
		"spider_name":       spiderName,
		"spider_id":         spiderID,
		"spider_parameters": string(spiderParameters),
		"spider_errors":     "true",
		"file_name":         fileName,
	}
	return s.client.doForm(ctx, s.client.scraperBaseURL, "/builder?platform=1", payload)
}

type GlassdoorJoblistingsRequest struct {
	SpiderID string `json:"spider_id,omitempty"`
	URL      string `json:"url,omitempty"`
	Keyword  string `json:"keyword,omitempty"`
	Location string `json:"location,omitempty"`
	Country  string `json:"country,omitempty"`
	FileName string `json:"file_name,omitempty"`
}

func (r GlassdoorJoblistingsRequest) params() map[string]string {
	values := map[string]string{}
	addStringParam(values, "spider_id", r.SpiderID)
	addStringParam(values, "url", r.URL)
	addStringParam(values, "keyword", r.Keyword)
	addStringParam(values, "location", r.Location)
	addStringParam(values, "country", r.Country)
	addStringParam(values, "file_name", r.FileName)
	return values
}

func (s *GlassdoorService) Joblistings(ctx context.Context, req GlassdoorJoblistingsRequest) (*RawResponse, error) {
	values := req.params()
	defaultParam(values, "spider_id", "glassdoor_joblistings_by-url")
	defaultParam(values, "file_name", "{{TasksID}}")
	if values["spider_id"] == "glassdoor_joblistings_by-url" {
		defaultParam(values, "url", "https://www.glassdoor.com/Job/new-york-data-analyst-jobs-SRCH_IL.0,8_IC1132348_KO9,21.htm")
	}
	if err := requireParam(values, "spider_id"); err != nil {
		return nil, err
	}
	spiderID := values["spider_id"]
	if spiderID == "" {
		spiderID = "glassdoor_joblistings_by-url"
	}
	spiderName := "glassdoor.com"
	fileName := values["file_name"]
	if fileName == "" {
		fileName = "{{TasksID}}"
	}
	delete(values, "spider_id")
	delete(values, "file_name")
	spiderParameters, err := json.Marshal([]map[string]string{values})
	if err != nil {
		return nil, err
	}
	payload := map[string]string{
		"spider_name":       spiderName,
		"spider_id":         spiderID,
		"spider_parameters": string(spiderParameters),
		"spider_errors":     "true",
		"file_name":         fileName,
	}
	return s.client.doForm(ctx, s.client.scraperBaseURL, "/builder?platform=1", payload)
}

type GoogleAIModeRequest struct {
	Q        string `json:"q,omitempty"`
	JSON     string `json:"json,omitempty"`
	Location string `json:"location,omitempty"`
	Uule     string `json:"uule,omitempty"`
	NoCache  string `json:"no_cache,omitempty"`
	Gl       string `json:"gl,omitempty"`
	Hl       string `json:"hl,omitempty"`
}

func (r GoogleAIModeRequest) params() map[string]string {
	values := map[string]string{}
	addStringParam(values, "q", r.Q)
	addStringParam(values, "json", r.JSON)
	addStringParam(values, "location", r.Location)
	addStringParam(values, "uule", r.Uule)
	addStringParam(values, "no_cache", r.NoCache)
	addStringParam(values, "gl", r.Gl)
	addStringParam(values, "hl", r.Hl)
	return values
}

func (s *GoogleService) AIMode(ctx context.Context, req GoogleAIModeRequest) (*RawResponse, error) {
	values := req.params()
	defaultParam(values, "q", "pizza")
	defaultParam(values, "json", "1")
	defaultParam(values, "no_cache", "false")
	if err := requireParam(values, "q"); err != nil {
		return nil, err
	}
	if err := requireParam(values, "json"); err != nil {
		return nil, err
	}
	values["engine"] = "google_ai_mode"
	return s.client.doForm(ctx, s.client.serpBaseURL, "/request", values)
}

type GoogleFinanceRequest struct {
	Q       string `json:"q,omitempty"`
	JSON    string `json:"json,omitempty"`
	Hl      string `json:"hl,omitempty"`
	Window  string `json:"window,omitempty"`
	NoCache string `json:"no_cache,omitempty"`
}

func (r GoogleFinanceRequest) params() map[string]string {
	values := map[string]string{}
	addStringParam(values, "q", r.Q)
	addStringParam(values, "json", r.JSON)
	addStringParam(values, "hl", r.Hl)
	addStringParam(values, "window", r.Window)
	addStringParam(values, "no_cache", r.NoCache)
	return values
}

func (s *GoogleService) Finance(ctx context.Context, req GoogleFinanceRequest) (*RawResponse, error) {
	values := req.params()
	defaultParam(values, "json", "1")
	defaultParam(values, "window", "1D")
	defaultParam(values, "no_cache", "false")
	if err := requireParam(values, "q"); err != nil {
		return nil, err
	}
	if err := requireParam(values, "json"); err != nil {
		return nil, err
	}
	values["engine"] = "google_finance"
	return s.client.doForm(ctx, s.client.serpBaseURL, "/request", values)
}

type GoogleFlightsRequest struct {
	DepartureID     string `json:"departure_id,omitempty"`
	ArrivalID       string `json:"arrival_id,omitempty"`
	JSON            string `json:"json,omitempty"`
	Gl              string `json:"gl,omitempty"`
	Hl              string `json:"hl,omitempty"`
	Currency        string `json:"currency,omitempty"`
	Type            string `json:"type,omitempty"`
	OutboundDate    string `json:"outbound_date,omitempty"`
	ReturnDate      string `json:"return_date,omitempty"`
	TravelClass     string `json:"travel_class,omitempty"`
	MultiCityJSON   string `json:"multi_city_json,omitempty"`
	ShowHidden      string `json:"show_hidden,omitempty"`
	ExcludeBasic    string `json:"exclude_basic,omitempty"`
	DeepSearch      string `json:"deep_search,omitempty"`
	Adults          string `json:"adults,omitempty"`
	Children        string `json:"children,omitempty"`
	InfantsInSeat   string `json:"infants_in_seat,omitempty"`
	InfantsOnLap    string `json:"infants_on_lap,omitempty"`
	SortBy          string `json:"sort_by,omitempty"`
	Stops           string `json:"stops,omitempty"`
	ExcludeAirlines string `json:"exclude_airlines,omitempty"`
	IncludeAirlines string `json:"include_airlines,omitempty"`
	Bags            string `json:"bags,omitempty"`
	MaxPrice        string `json:"max_price,omitempty"`
	OutboundTimes   string `json:"outbound_times,omitempty"`
	ReturnTimes     string `json:"return_times,omitempty"`
	Emissions       string `json:"emissions,omitempty"`
	LayoverDuration string `json:"layover_duration,omitempty"`
	ExcludeConns    string `json:"exclude_conns,omitempty"`
	MaxDuration     string `json:"max_duration,omitempty"`
	DepartureToken  string `json:"departure_token,omitempty"`
	NoCache         string `json:"no_cache,omitempty"`
}

func (r GoogleFlightsRequest) params() map[string]string {
	values := map[string]string{}
	addStringParam(values, "departure_id", r.DepartureID)
	addStringParam(values, "arrival_id", r.ArrivalID)
	addStringParam(values, "json", r.JSON)
	addStringParam(values, "gl", r.Gl)
	addStringParam(values, "hl", r.Hl)
	addStringParam(values, "currency", r.Currency)
	addStringParam(values, "type", r.Type)
	addStringParam(values, "outbound_date", r.OutboundDate)
	addStringParam(values, "return_date", r.ReturnDate)
	addStringParam(values, "travel_class", r.TravelClass)
	addStringParam(values, "multi_city_json", r.MultiCityJSON)
	addStringParam(values, "show_hidden", r.ShowHidden)
	addStringParam(values, "exclude_basic", r.ExcludeBasic)
	addStringParam(values, "deep_search", r.DeepSearch)
	addStringParam(values, "adults", r.Adults)
	addStringParam(values, "children", r.Children)
	addStringParam(values, "infants_in_seat", r.InfantsInSeat)
	addStringParam(values, "infants_on_lap", r.InfantsOnLap)
	addStringParam(values, "sort_by", r.SortBy)
	addStringParam(values, "stops", r.Stops)
	addStringParam(values, "exclude_airlines", r.ExcludeAirlines)
	addStringParam(values, "include_airlines", r.IncludeAirlines)
	addStringParam(values, "bags", r.Bags)
	addStringParam(values, "max_price", r.MaxPrice)
	addStringParam(values, "outbound_times", r.OutboundTimes)
	addStringParam(values, "return_times", r.ReturnTimes)
	addStringParam(values, "emissions", r.Emissions)
	addStringParam(values, "layover_duration", r.LayoverDuration)
	addStringParam(values, "exclude_conns", r.ExcludeConns)
	addStringParam(values, "max_duration", r.MaxDuration)
	addStringParam(values, "departure_token", r.DepartureToken)
	addStringParam(values, "no_cache", r.NoCache)
	return values
}

func (s *GoogleService) Flights(ctx context.Context, req GoogleFlightsRequest) (*RawResponse, error) {
	values := req.params()
	defaultParam(values, "json", "1")
	defaultParam(values, "currency", "USD")
	defaultParam(values, "type", "1")
	defaultParam(values, "travel_class", "1")
	defaultParam(values, "show_hidden", "false")
	defaultParam(values, "exclude_basic", "false")
	defaultParam(values, "deep_search", "false")
	defaultParam(values, "adults", "1")
	defaultParam(values, "children", "0")
	defaultParam(values, "infants_in_seat", "0")
	defaultParam(values, "infants_on_lap", "0")
	defaultParam(values, "sort_by", "1")
	defaultParam(values, "stops", "0")
	defaultParam(values, "bags", "0")
	defaultParam(values, "no_cache", "false")
	if err := requireParam(values, "json"); err != nil {
		return nil, err
	}
	if err := requireParam(values, "departure_id"); err != nil {
		return nil, err
	}
	if err := requireParam(values, "arrival_id"); err != nil {
		return nil, err
	}
	if err := requireParam(values, "outbound_date"); err != nil {
		return nil, err
	}
	values["engine"] = "google_flights"
	return s.client.doForm(ctx, s.client.serpBaseURL, "/request", values)
}

type GoogleHotelsRequest struct {
	Q                string `json:"q,omitempty"`
	JSON             string `json:"json,omitempty"`
	Hl               string `json:"hl,omitempty"`
	Gl               string `json:"gl,omitempty"`
	Currency         string `json:"currency,omitempty"`
	CheckInDate      string `json:"check_in_date,omitempty"`
	CheckOutDate     string `json:"check_out_date,omitempty"`
	Adults           string `json:"adults,omitempty"`
	Children         string `json:"children,omitempty"`
	ChildrenAges     string `json:"children_ages,omitempty"`
	SortBy           string `json:"sort_by,omitempty"`
	MinPrice         string `json:"min_price,omitempty"`
	MaxPrice         string `json:"max_price,omitempty"`
	PropertyTypes    string `json:"property_types,omitempty"`
	Amenities        string `json:"amenities,omitempty"`
	Rating           string `json:"rating,omitempty"`
	Brands           string `json:"brands,omitempty"`
	HotelClass       string `json:"hotel_class,omitempty"`
	FreeCancellation string `json:"free_cancellation,omitempty"`
	SpecialOffers    string `json:"special_offers,omitempty"`
	EcoCertified     string `json:"eco_certified,omitempty"`
	VacationRentals  string `json:"vacation_rentals,omitempty"`
	Bedrooms         string `json:"bedrooms,omitempty"`
	Bathrooms        string `json:"bathrooms,omitempty"`
	NextPageToken    string `json:"next_page_token,omitempty"`
	NoCache          string `json:"no_cache,omitempty"`
	PropertyToken    string `json:"property_token,omitempty"`
}

func (r GoogleHotelsRequest) params() map[string]string {
	values := map[string]string{}
	addStringParam(values, "q", r.Q)
	addStringParam(values, "json", r.JSON)
	addStringParam(values, "hl", r.Hl)
	addStringParam(values, "gl", r.Gl)
	addStringParam(values, "currency", r.Currency)
	addStringParam(values, "check_in_date", r.CheckInDate)
	addStringParam(values, "check_out_date", r.CheckOutDate)
	addStringParam(values, "adults", r.Adults)
	addStringParam(values, "children", r.Children)
	addStringParam(values, "children_ages", r.ChildrenAges)
	addStringParam(values, "sort_by", r.SortBy)
	addStringParam(values, "min_price", r.MinPrice)
	addStringParam(values, "max_price", r.MaxPrice)
	addStringParam(values, "property_types", r.PropertyTypes)
	addStringParam(values, "amenities", r.Amenities)
	addStringParam(values, "rating", r.Rating)
	addStringParam(values, "brands", r.Brands)
	addStringParam(values, "hotel_class", r.HotelClass)
	addStringParam(values, "free_cancellation", r.FreeCancellation)
	addStringParam(values, "special_offers", r.SpecialOffers)
	addStringParam(values, "eco_certified", r.EcoCertified)
	addStringParam(values, "vacation_rentals", r.VacationRentals)
	addStringParam(values, "bedrooms", r.Bedrooms)
	addStringParam(values, "bathrooms", r.Bathrooms)
	addStringParam(values, "next_page_token", r.NextPageToken)
	addStringParam(values, "no_cache", r.NoCache)
	addStringParam(values, "property_token", r.PropertyToken)
	return values
}

func (s *GoogleService) Hotels(ctx context.Context, req GoogleHotelsRequest) (*RawResponse, error) {
	values := req.params()
	defaultParam(values, "q", "pizza")
	defaultParam(values, "json", "1")
	defaultParam(values, "currency", "USD")
	defaultParam(values, "adults", "2")
	defaultParam(values, "children", "0")
	defaultParam(values, "bedrooms", "0")
	defaultParam(values, "bathrooms", "0")
	defaultParam(values, "no_cache", "false")
	if err := requireParam(values, "q"); err != nil {
		return nil, err
	}
	if err := requireParam(values, "json"); err != nil {
		return nil, err
	}
	if err := requireParam(values, "check_in_date"); err != nil {
		return nil, err
	}
	if err := requireParam(values, "check_out_date"); err != nil {
		return nil, err
	}
	if err := requireParam(values, "check_in_date"); err != nil {
		return nil, err
	}
	if err := requireParam(values, "check_out_date"); err != nil {
		return nil, err
	}
	values["engine"] = "google_hotels"
	return s.client.doForm(ctx, s.client.serpBaseURL, "/request", values)
}

type GoogleImagesRequest struct {
	Q            string `json:"q,omitempty"`
	JSON         string `json:"json,omitempty"`
	GoogleDomain string `json:"google_domain,omitempty"`
	Gl           string `json:"gl,omitempty"`
	Hl           string `json:"hl,omitempty"`
	Cr           string `json:"cr,omitempty"`
	Lr           string `json:"lr,omitempty"`
	Location     string `json:"location,omitempty"`
	Uule         string `json:"uule,omitempty"`
	Lat          string `json:"lat,omitempty"`
	Lon          string `json:"lon,omitempty"`
	Radius       string `json:"radius,omitempty"`
	Start        string `json:"start,omitempty"`
	Tbm          string `json:"tbm,omitempty"`
	Ludocid      string `json:"ludocid,omitempty"`
	Lsig         string `json:"lsig,omitempty"`
	Kgmid        string `json:"kgmid,omitempty"`
	Si           string `json:"si,omitempty"`
	Ibp          string `json:"ibp,omitempty"`
	Uds          string `json:"uds,omitempty"`
	Tbs          string `json:"tbs,omitempty"`
	Safe         string `json:"safe,omitempty"`
	Nfpr         string `json:"nfpr,omitempty"`
	Filter       string `json:"filter,omitempty"`
	Device       string `json:"device,omitempty"`
	RenderJs     string `json:"render_js,omitempty"`
	NoCache      string `json:"no_cache,omitempty"`
	AIOverview   string `json:"ai_overview,omitempty"`
}

func (r GoogleImagesRequest) params() map[string]string {
	values := map[string]string{}
	addStringParam(values, "q", r.Q)
	addStringParam(values, "json", r.JSON)
	addStringParam(values, "google_domain", r.GoogleDomain)
	addStringParam(values, "gl", r.Gl)
	addStringParam(values, "hl", r.Hl)
	addStringParam(values, "cr", r.Cr)
	addStringParam(values, "lr", r.Lr)
	addStringParam(values, "location", r.Location)
	addStringParam(values, "uule", r.Uule)
	addStringParam(values, "lat", r.Lat)
	addStringParam(values, "lon", r.Lon)
	addStringParam(values, "radius", r.Radius)
	addStringParam(values, "start", r.Start)
	addStringParam(values, "tbm", r.Tbm)
	addStringParam(values, "ludocid", r.Ludocid)
	addStringParam(values, "lsig", r.Lsig)
	addStringParam(values, "kgmid", r.Kgmid)
	addStringParam(values, "si", r.Si)
	addStringParam(values, "ibp", r.Ibp)
	addStringParam(values, "uds", r.Uds)
	addStringParam(values, "tbs", r.Tbs)
	addStringParam(values, "safe", r.Safe)
	addStringParam(values, "nfpr", r.Nfpr)
	addStringParam(values, "filter", r.Filter)
	addStringParam(values, "device", r.Device)
	addStringParam(values, "render_js", r.RenderJs)
	addStringParam(values, "no_cache", r.NoCache)
	addStringParam(values, "ai_overview", r.AIOverview)
	return values
}

func (s *GoogleService) Images(ctx context.Context, req GoogleImagesRequest) (*RawResponse, error) {
	values := req.params()
	defaultParam(values, "q", "pizza")
	defaultParam(values, "json", "1")
	defaultParam(values, "google_domain", "google.com")
	defaultParam(values, "start", "0")
	defaultParam(values, "tbm", "isch")
	defaultParam(values, "filter", "1")
	defaultParam(values, "device", "desktop")
	defaultParam(values, "no_cache", "false")
	if err := requireParam(values, "q"); err != nil {
		return nil, err
	}
	if err := requireParam(values, "json"); err != nil {
		return nil, err
	}
	values["engine"] = "google_images"
	return s.client.doForm(ctx, s.client.serpBaseURL, "/request", values)
}

type GoogleJobsRequest struct {
	Q             string `json:"q,omitempty"`
	JSON          string `json:"json,omitempty"`
	GoogleDomain  string `json:"google_domain,omitempty"`
	Gl            string `json:"gl,omitempty"`
	Hl            string `json:"hl,omitempty"`
	Location      string `json:"location,omitempty"`
	Uule          string `json:"uule,omitempty"`
	NextPageToken string `json:"next_page_token,omitempty"`
	Chips         string `json:"chips,omitempty"`
	Lrad          string `json:"lrad,omitempty"`
	Ltype         string `json:"ltype,omitempty"`
	Uds           string `json:"uds,omitempty"`
	NoCache       string `json:"no_cache,omitempty"`
}

func (r GoogleJobsRequest) params() map[string]string {
	values := map[string]string{}
	addStringParam(values, "q", r.Q)
	addStringParam(values, "json", r.JSON)
	addStringParam(values, "google_domain", r.GoogleDomain)
	addStringParam(values, "gl", r.Gl)
	addStringParam(values, "hl", r.Hl)
	addStringParam(values, "location", r.Location)
	addStringParam(values, "uule", r.Uule)
	addStringParam(values, "next_page_token", r.NextPageToken)
	addStringParam(values, "chips", r.Chips)
	addStringParam(values, "lrad", r.Lrad)
	addStringParam(values, "ltype", r.Ltype)
	addStringParam(values, "uds", r.Uds)
	addStringParam(values, "no_cache", r.NoCache)
	return values
}

func (s *GoogleService) Jobs(ctx context.Context, req GoogleJobsRequest) (*RawResponse, error) {
	values := req.params()
	defaultParam(values, "json", "1")
	defaultParam(values, "google_domain", "google.com")
	defaultParam(values, "no_cache", "false")
	if err := requireParam(values, "q"); err != nil {
		return nil, err
	}
	if err := requireParam(values, "json"); err != nil {
		return nil, err
	}
	values["engine"] = "google_jobs"
	return s.client.doForm(ctx, s.client.serpBaseURL, "/request", values)
}

type GoogleLensRequest struct {
	URL     string `json:"url,omitempty"`
	JSON    string `json:"json,omitempty"`
	Hl      string `json:"hl,omitempty"`
	Country string `json:"country,omitempty"`
	Type    string `json:"type,omitempty"`
	Q       string `json:"q,omitempty"`
	Safe    string `json:"safe,omitempty"`
	NoCache string `json:"no_cache,omitempty"`
}

func (r GoogleLensRequest) params() map[string]string {
	values := map[string]string{}
	addStringParam(values, "url", r.URL)
	addStringParam(values, "json", r.JSON)
	addStringParam(values, "hl", r.Hl)
	addStringParam(values, "country", r.Country)
	addStringParam(values, "type", r.Type)
	addStringParam(values, "q", r.Q)
	addStringParam(values, "safe", r.Safe)
	addStringParam(values, "no_cache", r.NoCache)
	return values
}

func (s *GoogleService) Lens(ctx context.Context, req GoogleLensRequest) (*RawResponse, error) {
	values := req.params()
	defaultParam(values, "json", "1")
	defaultParam(values, "type", "all")
	defaultParam(values, "no_cache", "false")
	if err := requireParam(values, "url"); err != nil {
		return nil, err
	}
	if err := requireParam(values, "json"); err != nil {
		return nil, err
	}
	values["engine"] = "google_lens"
	return s.client.doForm(ctx, s.client.serpBaseURL, "/request", values)
}

type GoogleLocalRequest struct {
	Q            string `json:"q,omitempty"`
	JSON         string `json:"json,omitempty"`
	GoogleDomain string `json:"google_domain,omitempty"`
	Gl           string `json:"gl,omitempty"`
	Hl           string `json:"hl,omitempty"`
	Location     string `json:"location,omitempty"`
	Uule         string `json:"uule,omitempty"`
	Start        string `json:"start,omitempty"`
	Ludocid      string `json:"ludocid,omitempty"`
	Tbs          string `json:"tbs,omitempty"`
	NoCache      string `json:"no_cache,omitempty"`
}

func (r GoogleLocalRequest) params() map[string]string {
	values := map[string]string{}
	addStringParam(values, "q", r.Q)
	addStringParam(values, "json", r.JSON)
	addStringParam(values, "google_domain", r.GoogleDomain)
	addStringParam(values, "gl", r.Gl)
	addStringParam(values, "hl", r.Hl)
	addStringParam(values, "location", r.Location)
	addStringParam(values, "uule", r.Uule)
	addStringParam(values, "start", r.Start)
	addStringParam(values, "ludocid", r.Ludocid)
	addStringParam(values, "tbs", r.Tbs)
	addStringParam(values, "no_cache", r.NoCache)
	return values
}

func (s *GoogleService) Local(ctx context.Context, req GoogleLocalRequest) (*RawResponse, error) {
	values := req.params()
	defaultParam(values, "json", "1")
	defaultParam(values, "google_domain", "google.com")
	defaultParam(values, "no_cache", "false")
	if err := requireParam(values, "q"); err != nil {
		return nil, err
	}
	if err := requireParam(values, "json"); err != nil {
		return nil, err
	}
	values["engine"] = "google_local"
	return s.client.doForm(ctx, s.client.serpBaseURL, "/request", values)
}

type GoogleMapCommentRequest struct {
	URL       string `json:"url,omitempty"`
	DaysLimit string `json:"days_limit,omitempty"`
	FileName  string `json:"file_name,omitempty"`
}

func (r GoogleMapCommentRequest) params() map[string]string {
	values := map[string]string{}
	addStringParam(values, "url", r.URL)
	addStringParam(values, "days_limit", r.DaysLimit)
	addStringParam(values, "file_name", r.FileName)
	return values
}

func (s *GoogleService) MapComment(ctx context.Context, req GoogleMapCommentRequest) (*RawResponse, error) {
	values := req.params()
	defaultParam(values, "url", "https://www.google.com/maps/place/Waterfront+Botanical+Gardens/@38.2630366,-85.7288454,15z/data=!4m8!3m7!1s0x8869731e16a7bdbd:0x2f5d238fefed7ca1!8m2!3d38.2632837!4d-85.7239738!9m1!1b1!16s%2Fg%2F11c709xzzx?hl=en&entry=ttu")
	defaultParam(values, "days_limit", "20")
	defaultParam(values, "file_name", "{{TasksID}}")
	if err := requireParam(values, "url"); err != nil {
		return nil, err
	}
	spiderID := "google_comment_by-url"
	spiderName := "google.com"
	fileName := values["file_name"]
	if fileName == "" {
		fileName = "{{TasksID}}"
	}
	delete(values, "spider_id")
	delete(values, "file_name")
	spiderParameters, err := json.Marshal([]map[string]string{values})
	if err != nil {
		return nil, err
	}
	payload := map[string]string{
		"spider_name":       spiderName,
		"spider_id":         spiderID,
		"spider_parameters": string(spiderParameters),
		"spider_errors":     "true",
		"file_name":         fileName,
	}
	return s.client.doForm(ctx, s.client.scraperBaseURL, "/builder?platform=1", payload)
}

type GoogleMapDetailsRequest struct {
	SpiderID  string `json:"spider_id,omitempty"`
	URL       string `json:"url,omitempty"`
	CID       string `json:"CID,omitempty"`
	Keyword   string `json:"keyword,omitempty"`
	Country   string `json:"country,omitempty"`
	Lat       string `json:"lat,omitempty"`
	Long      string `json:"long,omitempty"`
	ZoomLevel string `json:"zoom_level,omitempty"`
	PlaceID   string `json:"place_id,omitempty"`
	FileName  string `json:"file_name,omitempty"`
}

func (r GoogleMapDetailsRequest) params() map[string]string {
	values := map[string]string{}
	addStringParam(values, "spider_id", r.SpiderID)
	addStringParam(values, "url", r.URL)
	addStringParam(values, "CID", r.CID)
	addStringParam(values, "keyword", r.Keyword)
	addStringParam(values, "country", r.Country)
	addStringParam(values, "lat", r.Lat)
	addStringParam(values, "long", r.Long)
	addStringParam(values, "zoom_level", r.ZoomLevel)
	addStringParam(values, "place_id", r.PlaceID)
	addStringParam(values, "file_name", r.FileName)
	return values
}

func (s *GoogleService) MapDetails(ctx context.Context, req GoogleMapDetailsRequest) (*RawResponse, error) {
	values := req.params()
	defaultParam(values, "spider_id", "google_map-details_by-url")
	defaultParam(values, "file_name", "{{TasksID}}")
	if err := requireParam(values, "spider_id"); err != nil {
		return nil, err
	}
	spiderID := values["spider_id"]
	if spiderID == "" {
		spiderID = "google_map-details_by-url"
	}
	spiderName := "google.com"
	fileName := values["file_name"]
	if fileName == "" {
		fileName = "{{TasksID}}"
	}
	delete(values, "spider_id")
	delete(values, "file_name")
	spiderParameters, err := json.Marshal([]map[string]string{values})
	if err != nil {
		return nil, err
	}
	payload := map[string]string{
		"spider_name":       spiderName,
		"spider_id":         spiderID,
		"spider_parameters": string(spiderParameters),
		"spider_errors":     "true",
		"file_name":         fileName,
	}
	return s.client.doForm(ctx, s.client.scraperBaseURL, "/builder?platform=1", payload)
}

type GoogleMapsRequest struct {
	Q            string `json:"q,omitempty"`
	JSON         string `json:"json,omitempty"`
	Ll           string `json:"ll,omitempty"`
	Location     string `json:"location,omitempty"`
	Lat          string `json:"lat,omitempty"`
	Lon          string `json:"lon,omitempty"`
	Z            string `json:"z,omitempty"`
	M            string `json:"m,omitempty"`
	Nearby       string `json:"nearby,omitempty"`
	GoogleDomain string `json:"google_domain,omitempty"`
	Hl           string `json:"hl,omitempty"`
	Gl           string `json:"gl,omitempty"`
	Start        string `json:"start,omitempty"`
	Type         string `json:"type,omitempty"`
	Data         string `json:"data,omitempty"`
	PlaceID      string `json:"place_id,omitempty"`
	DataCID      string `json:"data_cid,omitempty"`
	NoCache      string `json:"no_cache,omitempty"`
}

func (r GoogleMapsRequest) params() map[string]string {
	values := map[string]string{}
	addStringParam(values, "q", r.Q)
	addStringParam(values, "json", r.JSON)
	addStringParam(values, "ll", r.Ll)
	addStringParam(values, "location", r.Location)
	addStringParam(values, "lat", r.Lat)
	addStringParam(values, "lon", r.Lon)
	addStringParam(values, "z", r.Z)
	addStringParam(values, "m", r.M)
	addStringParam(values, "nearby", r.Nearby)
	addStringParam(values, "google_domain", r.GoogleDomain)
	addStringParam(values, "hl", r.Hl)
	addStringParam(values, "gl", r.Gl)
	addStringParam(values, "start", r.Start)
	addStringParam(values, "type", r.Type)
	addStringParam(values, "data", r.Data)
	addStringParam(values, "place_id", r.PlaceID)
	addStringParam(values, "data_cid", r.DataCID)
	addStringParam(values, "no_cache", r.NoCache)
	return values
}

func (s *GoogleService) Maps(ctx context.Context, req GoogleMapsRequest) (*RawResponse, error) {
	values := req.params()
	defaultParam(values, "json", "1")
	defaultParam(values, "google_domain", "google.com")
	defaultParam(values, "start", "0")
	defaultParam(values, "no_cache", "false")
	if err := requireParam(values, "q"); err != nil {
		return nil, err
	}
	if err := requireParam(values, "json"); err != nil {
		return nil, err
	}
	values["engine"] = "google_maps"
	return s.client.doForm(ctx, s.client.serpBaseURL, "/request", values)
}

type GoogleNewsRequest struct {
	Q                string `json:"q,omitempty"`
	JSON             string `json:"json,omitempty"`
	NoCache          string `json:"no_cache,omitempty"`
	Gl               string `json:"gl,omitempty"`
	Hl               string `json:"hl,omitempty"`
	TopicToken       string `json:"topic_token,omitempty"`
	Kgmid            string `json:"kgmid,omitempty"`
	PublicationToken string `json:"publication_token,omitempty"`
	SectionToken     string `json:"section_token,omitempty"`
	StoryToken       string `json:"story_token,omitempty"`
	So               string `json:"so,omitempty"`
}

func (r GoogleNewsRequest) params() map[string]string {
	values := map[string]string{}
	addStringParam(values, "q", r.Q)
	addStringParam(values, "json", r.JSON)
	addStringParam(values, "no_cache", r.NoCache)
	addStringParam(values, "gl", r.Gl)
	addStringParam(values, "hl", r.Hl)
	addStringParam(values, "topic_token", r.TopicToken)
	addStringParam(values, "kgmid", r.Kgmid)
	addStringParam(values, "publication_token", r.PublicationToken)
	addStringParam(values, "section_token", r.SectionToken)
	addStringParam(values, "story_token", r.StoryToken)
	addStringParam(values, "so", r.So)
	return values
}

func (s *GoogleService) News(ctx context.Context, req GoogleNewsRequest) (*RawResponse, error) {
	values := req.params()
	defaultParam(values, "q", "pizza")
	defaultParam(values, "json", "1")
	defaultParam(values, "no_cache", "false")
	defaultParam(values, "so", "0")
	if err := requireParam(values, "q"); err != nil {
		return nil, err
	}
	if err := requireParam(values, "json"); err != nil {
		return nil, err
	}
	values["engine"] = "google_news"
	return s.client.doForm(ctx, s.client.serpBaseURL, "/request", values)
}

type GooglePatentsRequest struct {
	Q          string `json:"q,omitempty"`
	JSON       string `json:"json,omitempty"`
	Page       string `json:"page,omitempty"`
	Num        string `json:"num,omitempty"`
	Sort       string `json:"sort,omitempty"`
	Clustered  string `json:"clustered,omitempty"`
	Dups       string `json:"dups,omitempty"`
	Patents    string `json:"patents,omitempty"`
	Scholar    string `json:"scholar,omitempty"`
	Before     string `json:"before,omitempty"`
	After      string `json:"after,omitempty"`
	Inventor   string `json:"inventor,omitempty"`
	Assignee   string `json:"assignee,omitempty"`
	Country    string `json:"country,omitempty"`
	Language   string `json:"language,omitempty"`
	Status     string `json:"status,omitempty"`
	Type       string `json:"type,omitempty"`
	Litigation string `json:"litigation,omitempty"`
	NoCache    string `json:"no_cache,omitempty"`
}

func (r GooglePatentsRequest) params() map[string]string {
	values := map[string]string{}
	addStringParam(values, "q", r.Q)
	addStringParam(values, "json", r.JSON)
	addStringParam(values, "page", r.Page)
	addStringParam(values, "num", r.Num)
	addStringParam(values, "sort", r.Sort)
	addStringParam(values, "clustered", r.Clustered)
	addStringParam(values, "dups", r.Dups)
	addStringParam(values, "patents", r.Patents)
	addStringParam(values, "scholar", r.Scholar)
	addStringParam(values, "before", r.Before)
	addStringParam(values, "after", r.After)
	addStringParam(values, "inventor", r.Inventor)
	addStringParam(values, "assignee", r.Assignee)
	addStringParam(values, "country", r.Country)
	addStringParam(values, "language", r.Language)
	addStringParam(values, "status", r.Status)
	addStringParam(values, "type", r.Type)
	addStringParam(values, "litigation", r.Litigation)
	addStringParam(values, "no_cache", r.NoCache)
	return values
}

func (s *GoogleService) Patents(ctx context.Context, req GooglePatentsRequest) (*RawResponse, error) {
	values := req.params()
	defaultParam(values, "json", "1")
	defaultParam(values, "page", "1")
	defaultParam(values, "dups", "family")
	defaultParam(values, "patents", "true")
	defaultParam(values, "scholar", "false")
	defaultParam(values, "no_cache", "false")
	if err := requireParam(values, "q"); err != nil {
		return nil, err
	}
	if err := requireParam(values, "json"); err != nil {
		return nil, err
	}
	values["engine"] = "google_patents"
	return s.client.doForm(ctx, s.client.serpBaseURL, "/request", values)
}

type GooglePlayRequest struct {
	Q                string `json:"q,omitempty"`
	JSON             string `json:"json,omitempty"`
	Hl               string `json:"hl,omitempty"`
	Gl               string `json:"gl,omitempty"`
	AppsCategory     string `json:"apps_category,omitempty"`
	NextPageToken    string `json:"next_page_token,omitempty"`
	SectionPageToken string `json:"section_page_token,omitempty"`
	Chart            string `json:"chart,omitempty"`
	SeeMoreToken     string `json:"see_more_token,omitempty"`
	StoreDevice      string `json:"store_device,omitempty"`
	Age              string `json:"age,omitempty"`
	NoCache          string `json:"no_cache,omitempty"`
}

func (r GooglePlayRequest) params() map[string]string {
	values := map[string]string{}
	addStringParam(values, "q", r.Q)
	addStringParam(values, "json", r.JSON)
	addStringParam(values, "hl", r.Hl)
	addStringParam(values, "gl", r.Gl)
	addStringParam(values, "apps_category", r.AppsCategory)
	addStringParam(values, "next_page_token", r.NextPageToken)
	addStringParam(values, "section_page_token", r.SectionPageToken)
	addStringParam(values, "chart", r.Chart)
	addStringParam(values, "see_more_token", r.SeeMoreToken)
	addStringParam(values, "store_device", r.StoreDevice)
	addStringParam(values, "age", r.Age)
	addStringParam(values, "no_cache", r.NoCache)
	return values
}

func (s *GoogleService) Play(ctx context.Context, req GooglePlayRequest) (*RawResponse, error) {
	values := req.params()
	defaultParam(values, "json", "1")
	defaultParam(values, "gl", "us")
	defaultParam(values, "no_cache", "false")
	if err := requireParam(values, "q"); err != nil {
		return nil, err
	}
	if err := requireParam(values, "json"); err != nil {
		return nil, err
	}
	values["engine"] = "google_play"
	return s.client.doForm(ctx, s.client.serpBaseURL, "/request", values)
}

type GooglePlayStoreInformationRequest struct {
	AppURL   string `json:"app_url,omitempty"`
	FileName string `json:"file_name,omitempty"`
}

func (r GooglePlayStoreInformationRequest) params() map[string]string {
	values := map[string]string{}
	addStringParam(values, "app_url", r.AppURL)
	addStringParam(values, "file_name", r.FileName)
	return values
}

func (s *GoogleService) PlayStoreInformation(ctx context.Context, req GooglePlayStoreInformationRequest) (*RawResponse, error) {
	values := req.params()
	defaultParam(values, "app_url", "https://play.google.com/store/apps/details?id=com.linkedin.android")
	defaultParam(values, "file_name", "{{TasksID}}")
	if err := requireParam(values, "app_url"); err != nil {
		return nil, err
	}
	spiderID := "google-play-store_information_by-url"
	spiderName := "play.google.com"
	fileName := values["file_name"]
	if fileName == "" {
		fileName = "{{TasksID}}"
	}
	delete(values, "spider_id")
	delete(values, "file_name")
	spiderParameters, err := json.Marshal([]map[string]string{values})
	if err != nil {
		return nil, err
	}
	payload := map[string]string{
		"spider_name":       spiderName,
		"spider_id":         spiderID,
		"spider_parameters": string(spiderParameters),
		"spider_errors":     "true",
		"file_name":         fileName,
	}
	return s.client.doForm(ctx, s.client.scraperBaseURL, "/builder?platform=1", payload)
}

type GooglePlayStoreReviewsRequest struct {
	AppURL   string `json:"app_url,omitempty"`
	FileName string `json:"file_name,omitempty"`
}

func (r GooglePlayStoreReviewsRequest) params() map[string]string {
	values := map[string]string{}
	addStringParam(values, "app_url", r.AppURL)
	addStringParam(values, "file_name", r.FileName)
	return values
}

func (s *GoogleService) PlayStoreReviews(ctx context.Context, req GooglePlayStoreReviewsRequest) (*RawResponse, error) {
	values := req.params()
	defaultParam(values, "app_url", "https://play.google.com/store/apps/details?id=com.linkedin.android")
	defaultParam(values, "file_name", "{{TasksID}}")
	if err := requireParam(values, "app_url"); err != nil {
		return nil, err
	}
	spiderID := "google-play-store_reviews_by-url"
	spiderName := "play.google.com"
	fileName := values["file_name"]
	if fileName == "" {
		fileName = "{{TasksID}}"
	}
	delete(values, "spider_id")
	delete(values, "file_name")
	spiderParameters, err := json.Marshal([]map[string]string{values})
	if err != nil {
		return nil, err
	}
	payload := map[string]string{
		"spider_name":       spiderName,
		"spider_id":         spiderID,
		"spider_parameters": string(spiderParameters),
		"spider_errors":     "true",
		"file_name":         fileName,
	}
	return s.client.doForm(ctx, s.client.scraperBaseURL, "/builder?platform=1", payload)
}

type GoogleScholarRequest struct {
	Q       string `json:"q,omitempty"`
	JSON    string `json:"json,omitempty"`
	Hl      string `json:"hl,omitempty"`
	Lr      string `json:"lr,omitempty"`
	Start   string `json:"start,omitempty"`
	Num     string `json:"num,omitempty"`
	Cites   string `json:"cites,omitempty"`
	AsYlo   string `json:"as_ylo,omitempty"`
	AsYhi   string `json:"as_yhi,omitempty"`
	Scisbd  string `json:"scisbd,omitempty"`
	Cluster string `json:"cluster,omitempty"`
	AsSdt   string `json:"as_sdt,omitempty"`
	Safe    string `json:"safe,omitempty"`
	Filter  string `json:"filter,omitempty"`
	AsVis   string `json:"as_vis,omitempty"`
	AsRr    string `json:"as_rr,omitempty"`
	NoCache string `json:"no_cache,omitempty"`
}

func (r GoogleScholarRequest) params() map[string]string {
	values := map[string]string{}
	addStringParam(values, "q", r.Q)
	addStringParam(values, "json", r.JSON)
	addStringParam(values, "hl", r.Hl)
	addStringParam(values, "lr", r.Lr)
	addStringParam(values, "start", r.Start)
	addStringParam(values, "num", r.Num)
	addStringParam(values, "cites", r.Cites)
	addStringParam(values, "as_ylo", r.AsYlo)
	addStringParam(values, "as_yhi", r.AsYhi)
	addStringParam(values, "scisbd", r.Scisbd)
	addStringParam(values, "cluster", r.Cluster)
	addStringParam(values, "as_sdt", r.AsSdt)
	addStringParam(values, "safe", r.Safe)
	addStringParam(values, "filter", r.Filter)
	addStringParam(values, "as_vis", r.AsVis)
	addStringParam(values, "as_rr", r.AsRr)
	addStringParam(values, "no_cache", r.NoCache)
	return values
}

func (s *GoogleService) Scholar(ctx context.Context, req GoogleScholarRequest) (*RawResponse, error) {
	values := req.params()
	defaultParam(values, "json", "1")
	defaultParam(values, "start", "0")
	defaultParam(values, "num", "10")
	defaultParam(values, "scisbd", "0")
	defaultParam(values, "as_sdt", "0")
	defaultParam(values, "filter", "1")
	defaultParam(values, "as_vis", "0")
	defaultParam(values, "as_rr", "0")
	defaultParam(values, "no_cache", "false")
	if err := requireParam(values, "json"); err != nil {
		return nil, err
	}
	values["engine"] = "google_scholar"
	return s.client.doForm(ctx, s.client.serpBaseURL, "/request", values)
}

type GoogleSearchRequest struct {
	Q            string `json:"q,omitempty"`
	JSON         string `json:"json,omitempty"`
	GoogleDomain string `json:"google_domain,omitempty"`
	Gl           string `json:"gl,omitempty"`
	Hl           string `json:"hl,omitempty"`
	Cr           string `json:"cr,omitempty"`
	Lr           string `json:"lr,omitempty"`
	Location     string `json:"location,omitempty"`
	Uule         string `json:"uule,omitempty"`
	Start        string `json:"start,omitempty"`
	Tbs          string `json:"tbs,omitempty"`
	Safe         string `json:"safe,omitempty"`
	Nfpr         string `json:"nfpr,omitempty"`
	Filter       string `json:"filter,omitempty"`
	Device       string `json:"device,omitempty"`
	RenderJs     string `json:"render_js,omitempty"`
	NoCache      string `json:"no_cache,omitempty"`
	AIOverview   string `json:"ai_overview,omitempty"`
}

func (r GoogleSearchRequest) params() map[string]string {
	values := map[string]string{}
	addStringParam(values, "q", r.Q)
	addStringParam(values, "json", r.JSON)
	addStringParam(values, "google_domain", r.GoogleDomain)
	addStringParam(values, "gl", r.Gl)
	addStringParam(values, "hl", r.Hl)
	addStringParam(values, "cr", r.Cr)
	addStringParam(values, "lr", r.Lr)
	addStringParam(values, "location", r.Location)
	addStringParam(values, "uule", r.Uule)
	addStringParam(values, "start", r.Start)
	addStringParam(values, "tbs", r.Tbs)
	addStringParam(values, "safe", r.Safe)
	addStringParam(values, "nfpr", r.Nfpr)
	addStringParam(values, "filter", r.Filter)
	addStringParam(values, "device", r.Device)
	addStringParam(values, "render_js", r.RenderJs)
	addStringParam(values, "no_cache", r.NoCache)
	addStringParam(values, "ai_overview", r.AIOverview)
	return values
}

func (s *GoogleService) Search(ctx context.Context, req GoogleSearchRequest) (*RawResponse, error) {
	values := req.params()
	defaultParam(values, "q", "pizza")
	defaultParam(values, "json", "1")
	defaultParam(values, "google_domain", "google.com")
	defaultParam(values, "start", "0")
	defaultParam(values, "filter", "1")
	defaultParam(values, "device", "desktop")
	defaultParam(values, "no_cache", "false")
	if err := requireParam(values, "q"); err != nil {
		return nil, err
	}
	if err := requireParam(values, "json"); err != nil {
		return nil, err
	}
	values["engine"] = "google"
	return s.client.doForm(ctx, s.client.serpBaseURL, "/request", values)
}

type GoogleShoppingRequest struct {
	Q             string `json:"q,omitempty"`
	JSON          string `json:"json,omitempty"`
	GoogleDomain  string `json:"google_domain,omitempty"`
	Gl            string `json:"gl,omitempty"`
	Hl            string `json:"hl,omitempty"`
	Location      string `json:"location,omitempty"`
	Uule          string `json:"uule,omitempty"`
	Start         string `json:"start,omitempty"`
	Shoprs        string `json:"shoprs,omitempty"`
	MinPrice      string `json:"min_price,omitempty"`
	MaxPrice      string `json:"max_price,omitempty"`
	SortBy        string `json:"sort_by,omitempty"`
	FreeShipping  string `json:"free_shipping,omitempty"`
	OnSale        string `json:"on_sale,omitempty"`
	SmallBusiness string `json:"small_business,omitempty"`
	NoCache       string `json:"no_cache,omitempty"`
}

func (r GoogleShoppingRequest) params() map[string]string {
	values := map[string]string{}
	addStringParam(values, "q", r.Q)
	addStringParam(values, "json", r.JSON)
	addStringParam(values, "google_domain", r.GoogleDomain)
	addStringParam(values, "gl", r.Gl)
	addStringParam(values, "hl", r.Hl)
	addStringParam(values, "location", r.Location)
	addStringParam(values, "uule", r.Uule)
	addStringParam(values, "start", r.Start)
	addStringParam(values, "shoprs", r.Shoprs)
	addStringParam(values, "min_price", r.MinPrice)
	addStringParam(values, "max_price", r.MaxPrice)
	addStringParam(values, "sort_by", r.SortBy)
	addStringParam(values, "free_shipping", r.FreeShipping)
	addStringParam(values, "on_sale", r.OnSale)
	addStringParam(values, "small_business", r.SmallBusiness)
	addStringParam(values, "no_cache", r.NoCache)
	return values
}

func (s *GoogleService) Shopping(ctx context.Context, req GoogleShoppingRequest) (*RawResponse, error) {
	values := req.params()
	defaultParam(values, "json", "1")
	defaultParam(values, "google_domain", "google.com")
	defaultParam(values, "no_cache", "false")
	if err := requireParam(values, "json"); err != nil {
		return nil, err
	}
	values["engine"] = "google_shopping"
	return s.client.doForm(ctx, s.client.serpBaseURL, "/request", values)
}

type GoogleShoppingInfoRequest struct {
	Keyword  string `json:"keyword,omitempty"`
	FileName string `json:"file_name,omitempty"`
}

func (r GoogleShoppingInfoRequest) params() map[string]string {
	values := map[string]string{}
	addStringParam(values, "keyword", r.Keyword)
	addStringParam(values, "file_name", r.FileName)
	return values
}

func (s *GoogleService) ShoppingInfo(ctx context.Context, req GoogleShoppingInfoRequest) (*RawResponse, error) {
	values := req.params()
	defaultParam(values, "keyword", "iphone")
	defaultParam(values, "file_name", "{{TasksID}}")
	if err := requireParam(values, "keyword"); err != nil {
		return nil, err
	}
	spiderID := "google_shopping_by-keywords"
	spiderName := "google.com"
	fileName := values["file_name"]
	if fileName == "" {
		fileName = "{{TasksID}}"
	}
	delete(values, "spider_id")
	delete(values, "file_name")
	spiderParameters, err := json.Marshal([]map[string]string{values})
	if err != nil {
		return nil, err
	}
	payload := map[string]string{
		"spider_name":       spiderName,
		"spider_id":         spiderID,
		"spider_parameters": string(spiderParameters),
		"spider_errors":     "true",
		"file_name":         fileName,
	}
	return s.client.doForm(ctx, s.client.scraperBaseURL, "/builder?platform=1", payload)
}

type GoogleTrendsRequest struct {
	Q                      string `json:"q,omitempty"`
	JSON                   string `json:"json,omitempty"`
	Hl                     string `json:"hl,omitempty"`
	Geo                    string `json:"geo,omitempty"`
	Region                 string `json:"region,omitempty"`
	DataType               string `json:"data_type,omitempty"`
	Tz                     string `json:"tz,omitempty"`
	Cat                    string `json:"cat,omitempty"`
	Gprop                  string `json:"gprop,omitempty"`
	Date                   string `json:"date,omitempty"`
	Csv                    string `json:"csv,omitempty"`
	IncludeLowSearchVolume string `json:"include_low_search_volume,omitempty"`
	NoCache                string `json:"no_cache,omitempty"`
}

func (r GoogleTrendsRequest) params() map[string]string {
	values := map[string]string{}
	addStringParam(values, "q", r.Q)
	addStringParam(values, "json", r.JSON)
	addStringParam(values, "hl", r.Hl)
	addStringParam(values, "geo", r.Geo)
	addStringParam(values, "region", r.Region)
	addStringParam(values, "data_type", r.DataType)
	addStringParam(values, "tz", r.Tz)
	addStringParam(values, "cat", r.Cat)
	addStringParam(values, "gprop", r.Gprop)
	addStringParam(values, "date", r.Date)
	addStringParam(values, "csv", r.Csv)
	addStringParam(values, "include_low_search_volume", r.IncludeLowSearchVolume)
	addStringParam(values, "no_cache", r.NoCache)
	return values
}

func (s *GoogleService) Trends(ctx context.Context, req GoogleTrendsRequest) (*RawResponse, error) {
	values := req.params()
	defaultParam(values, "q", "pizza")
	defaultParam(values, "json", "1")
	defaultParam(values, "tz", "420")
	defaultParam(values, "cat", "0")
	defaultParam(values, "no_cache", "false")
	if err := requireParam(values, "q"); err != nil {
		return nil, err
	}
	if err := requireParam(values, "json"); err != nil {
		return nil, err
	}
	values["engine"] = "google_trends"
	return s.client.doForm(ctx, s.client.serpBaseURL, "/request", values)
}

type GoogleVideosRequest struct {
	Q            string `json:"q,omitempty"`
	JSON         string `json:"json,omitempty"`
	GoogleDomain string `json:"google_domain,omitempty"`
	Gl           string `json:"gl,omitempty"`
	Hl           string `json:"hl,omitempty"`
	Location     string `json:"location,omitempty"`
	Uule         string `json:"uule,omitempty"`
	Start        string `json:"start,omitempty"`
	Tbs          string `json:"tbs,omitempty"`
	NoCache      string `json:"no_cache,omitempty"`
	Lr           string `json:"lr,omitempty"`
	Safe         string `json:"safe,omitempty"`
	Nfpr         string `json:"nfpr,omitempty"`
	Filter       string `json:"filter,omitempty"`
}

func (r GoogleVideosRequest) params() map[string]string {
	values := map[string]string{}
	addStringParam(values, "q", r.Q)
	addStringParam(values, "json", r.JSON)
	addStringParam(values, "google_domain", r.GoogleDomain)
	addStringParam(values, "gl", r.Gl)
	addStringParam(values, "hl", r.Hl)
	addStringParam(values, "location", r.Location)
	addStringParam(values, "uule", r.Uule)
	addStringParam(values, "start", r.Start)
	addStringParam(values, "tbs", r.Tbs)
	addStringParam(values, "no_cache", r.NoCache)
	addStringParam(values, "lr", r.Lr)
	addStringParam(values, "safe", r.Safe)
	addStringParam(values, "nfpr", r.Nfpr)
	addStringParam(values, "filter", r.Filter)
	return values
}

func (s *GoogleService) Videos(ctx context.Context, req GoogleVideosRequest) (*RawResponse, error) {
	values := req.params()
	defaultParam(values, "json", "1")
	defaultParam(values, "google_domain", "google.com")
	defaultParam(values, "no_cache", "false")
	defaultParam(values, "filter", "1")
	if err := requireParam(values, "q"); err != nil {
		return nil, err
	}
	if err := requireParam(values, "json"); err != nil {
		return nil, err
	}
	values["engine"] = "google_videos"
	return s.client.doForm(ctx, s.client.serpBaseURL, "/request", values)
}

type IndeedCompaniesInfoRequest struct {
	SpiderID       string `json:"spider_id,omitempty"`
	CompanyListURL string `json:"company_list_url,omitempty"`
	Keyword        string `json:"keyword,omitempty"`
	Industry       string `json:"industry,omitempty"`
	State          string `json:"state,omitempty"`
	CompanyURL     string `json:"company_url,omitempty"`
	FileName       string `json:"file_name,omitempty"`
}

func (r IndeedCompaniesInfoRequest) params() map[string]string {
	values := map[string]string{}
	addStringParam(values, "spider_id", r.SpiderID)
	addStringParam(values, "company_list_url", r.CompanyListURL)
	addStringParam(values, "keyword", r.Keyword)
	addStringParam(values, "industry", r.Industry)
	addStringParam(values, "state", r.State)
	addStringParam(values, "company_url", r.CompanyURL)
	addStringParam(values, "file_name", r.FileName)
	return values
}

func (s *IndeedService) CompaniesInfo(ctx context.Context, req IndeedCompaniesInfoRequest) (*RawResponse, error) {
	values := req.params()
	defaultParam(values, "spider_id", "indeed_companies-info_by-company-list-url")
	defaultParam(values, "file_name", "{{TasksID}}")
	if err := requireParam(values, "spider_id"); err != nil {
		return nil, err
	}
	spiderID := values["spider_id"]
	if spiderID == "" {
		spiderID = "indeed_companies-info_by-company-list-url"
	}
	spiderName := "indeed.com"
	fileName := values["file_name"]
	if fileName == "" {
		fileName = "{{TasksID}}"
	}
	delete(values, "spider_id")
	delete(values, "file_name")
	spiderParameters, err := json.Marshal([]map[string]string{values})
	if err != nil {
		return nil, err
	}
	payload := map[string]string{
		"spider_name":       spiderName,
		"spider_id":         spiderID,
		"spider_parameters": string(spiderParameters),
		"spider_errors":     "true",
		"file_name":         fileName,
	}
	return s.client.doForm(ctx, s.client.scraperBaseURL, "/builder?platform=1", payload)
}

type IndeedJobListingsRequest struct {
	JobURL   string `json:"job_url,omitempty"`
	FileName string `json:"file_name,omitempty"`
}

func (r IndeedJobListingsRequest) params() map[string]string {
	values := map[string]string{}
	addStringParam(values, "job_url", r.JobURL)
	addStringParam(values, "file_name", r.FileName)
	return values
}

func (s *IndeedService) JobListings(ctx context.Context, req IndeedJobListingsRequest) (*RawResponse, error) {
	values := req.params()
	defaultParam(values, "job_url", "https://fr.indeed.com/viewjob?jk=55b3e5dfa0c2ff66")
	defaultParam(values, "file_name", "{{TasksID}}")
	if err := requireParam(values, "job_url"); err != nil {
		return nil, err
	}
	spiderID := "indeed_job-listings_by-job-url"
	spiderName := "indeed.com"
	fileName := values["file_name"]
	if fileName == "" {
		fileName = "{{TasksID}}"
	}
	delete(values, "spider_id")
	delete(values, "file_name")
	spiderParameters, err := json.Marshal([]map[string]string{values})
	if err != nil {
		return nil, err
	}
	payload := map[string]string{
		"spider_name":       spiderName,
		"spider_id":         spiderID,
		"spider_parameters": string(spiderParameters),
		"spider_errors":     "true",
		"file_name":         fileName,
	}
	return s.client.doForm(ctx, s.client.scraperBaseURL, "/builder?platform=1", payload)
}

type InstagramCommentRequest struct {
	Posturl  string `json:"posturl,omitempty"`
	FileName string `json:"file_name,omitempty"`
}

func (r InstagramCommentRequest) params() map[string]string {
	values := map[string]string{}
	addStringParam(values, "posturl", r.Posturl)
	addStringParam(values, "file_name", r.FileName)
	return values
}

func (s *InstagramService) Comment(ctx context.Context, req InstagramCommentRequest) (*RawResponse, error) {
	values := req.params()
	defaultParam(values, "posturl", "https://www.instagram.com/cats_of_instagram/reel/C4GLo_eLO2e/")
	defaultParam(values, "file_name", "{{TasksID}}")
	if err := requireParam(values, "posturl"); err != nil {
		return nil, err
	}
	spiderID := "ins_comment_by-posturl"
	spiderName := "instagram.com"
	fileName := values["file_name"]
	if fileName == "" {
		fileName = "{{TasksID}}"
	}
	delete(values, "spider_id")
	delete(values, "file_name")
	spiderParameters, err := json.Marshal([]map[string]string{values})
	if err != nil {
		return nil, err
	}
	payload := map[string]string{
		"spider_name":       spiderName,
		"spider_id":         spiderID,
		"spider_parameters": string(spiderParameters),
		"spider_errors":     "true",
		"file_name":         fileName,
	}
	return s.client.doForm(ctx, s.client.scraperBaseURL, "/builder?platform=1", payload)
}

type InstagramProfilesRequest struct {
	SpiderID   string `json:"spider_id,omitempty"`
	Username   string `json:"username,omitempty"`
	Profileurl string `json:"profileurl,omitempty"`
	FileName   string `json:"file_name,omitempty"`
}

func (r InstagramProfilesRequest) params() map[string]string {
	values := map[string]string{}
	addStringParam(values, "spider_id", r.SpiderID)
	addStringParam(values, "username", r.Username)
	addStringParam(values, "profileurl", r.Profileurl)
	addStringParam(values, "file_name", r.FileName)
	return values
}

func (s *InstagramService) Profiles(ctx context.Context, req InstagramProfilesRequest) (*RawResponse, error) {
	values := req.params()
	defaultParam(values, "spider_id", "ins_profiles_by-username")
	defaultParam(values, "file_name", "{{TasksID}}")
	if err := requireParam(values, "spider_id"); err != nil {
		return nil, err
	}
	spiderID := values["spider_id"]
	if spiderID == "" {
		spiderID = "ins_profiles_by-username"
	}
	spiderName := "instagram.com"
	fileName := values["file_name"]
	if fileName == "" {
		fileName = "{{TasksID}}"
	}
	delete(values, "spider_id")
	delete(values, "file_name")
	spiderParameters, err := json.Marshal([]map[string]string{values})
	if err != nil {
		return nil, err
	}
	payload := map[string]string{
		"spider_name":       spiderName,
		"spider_id":         spiderID,
		"spider_parameters": string(spiderParameters),
		"spider_errors":     "true",
		"file_name":         fileName,
	}
	return s.client.doForm(ctx, s.client.scraperBaseURL, "/builder?platform=1", payload)
}

type InstagramReelRequest struct {
	SpiderID          string `json:"spider_id,omitempty"`
	URL               string `json:"url,omitempty"`
	NumOfPosts        string `json:"num_of_posts,omitempty"`
	PostsToNotInclude string `json:"posts_to_not_include,omitempty"`
	StartDate         string `json:"start_date,omitempty"`
	EndDate           string `json:"end_date,omitempty"`
	FileName          string `json:"file_name,omitempty"`
}

func (r InstagramReelRequest) params() map[string]string {
	values := map[string]string{}
	addStringParam(values, "spider_id", r.SpiderID)
	addStringParam(values, "url", r.URL)
	addStringParam(values, "num_of_posts", r.NumOfPosts)
	addStringParam(values, "posts_to_not_include", r.PostsToNotInclude)
	addStringParam(values, "start_date", r.StartDate)
	addStringParam(values, "end_date", r.EndDate)
	addStringParam(values, "file_name", r.FileName)
	return values
}

func (s *InstagramService) Reel(ctx context.Context, req InstagramReelRequest) (*RawResponse, error) {
	values := req.params()
	defaultParam(values, "spider_id", "ins_reel_by-url")
	defaultParam(values, "url", "https://www.instagram.com/reel/C5Rdyj_q7YN/")
	defaultParam(values, "file_name", "{{TasksID}}")
	if err := requireParam(values, "spider_id"); err != nil {
		return nil, err
	}
	if err := requireParam(values, "url"); err != nil {
		return nil, err
	}
	spiderID := values["spider_id"]
	if spiderID == "" {
		spiderID = "ins_reel_by-url"
	}
	spiderName := "instagram.com"
	fileName := values["file_name"]
	if fileName == "" {
		fileName = "{{TasksID}}"
	}
	delete(values, "spider_id")
	delete(values, "file_name")
	spiderParameters, err := json.Marshal([]map[string]string{values})
	if err != nil {
		return nil, err
	}
	payload := map[string]string{
		"spider_name":       spiderName,
		"spider_id":         spiderID,
		"spider_parameters": string(spiderParameters),
		"spider_errors":     "true",
		"file_name":         fileName,
	}
	return s.client.doForm(ctx, s.client.scraperBaseURL, "/builder?platform=1", payload)
}

type LinkedInCompanyInformationRequest struct {
	URL      string `json:"url,omitempty"`
	FileName string `json:"file_name,omitempty"`
}

func (r LinkedInCompanyInformationRequest) params() map[string]string {
	values := map[string]string{}
	addStringParam(values, "url", r.URL)
	addStringParam(values, "file_name", r.FileName)
	return values
}

func (s *LinkedInService) CompanyInformation(ctx context.Context, req LinkedInCompanyInformationRequest) (*RawResponse, error) {
	values := req.params()
	defaultParam(values, "url", "https://www.linkedin.com/company/dynamo-software")
	defaultParam(values, "file_name", "{{TasksID}}")
	if err := requireParam(values, "url"); err != nil {
		return nil, err
	}
	spiderID := "linkedin_company_information_by-url"
	spiderName := "linkedin.com"
	fileName := values["file_name"]
	if fileName == "" {
		fileName = "{{TasksID}}"
	}
	delete(values, "spider_id")
	delete(values, "file_name")
	spiderParameters, err := json.Marshal([]map[string]string{values})
	if err != nil {
		return nil, err
	}
	payload := map[string]string{
		"spider_name":       spiderName,
		"spider_id":         spiderID,
		"spider_parameters": string(spiderParameters),
		"spider_errors":     "true",
		"file_name":         fileName,
	}
	return s.client.doForm(ctx, s.client.scraperBaseURL, "/builder?platform=1", payload)
}

type LinkedInJobListingsInformationRequest struct {
	SpiderID      string `json:"spider_id,omitempty"`
	JobListingURL string `json:"job_listing_url,omitempty"`
	JobURL        string `json:"job_url,omitempty"`
	Location      string `json:"location,omitempty"`
	Keyword       string `json:"keyword,omitempty"`
	FileName      string `json:"file_name,omitempty"`
}

func (r LinkedInJobListingsInformationRequest) params() map[string]string {
	values := map[string]string{}
	addStringParam(values, "spider_id", r.SpiderID)
	addStringParam(values, "job_listing_url", r.JobListingURL)
	addStringParam(values, "job_url", r.JobURL)
	addStringParam(values, "location", r.Location)
	addStringParam(values, "keyword", r.Keyword)
	addStringParam(values, "file_name", r.FileName)
	return values
}

func (s *LinkedInService) JobListingsInformation(ctx context.Context, req LinkedInJobListingsInformationRequest) (*RawResponse, error) {
	values := req.params()
	defaultParam(values, "spider_id", "linkedin_job_listings_information_by-job-listing-url")
	defaultParam(values, "file_name", "{{TasksID}}")
	if err := requireParam(values, "spider_id"); err != nil {
		return nil, err
	}
	spiderID := values["spider_id"]
	if spiderID == "" {
		spiderID = "linkedin_job_listings_information_by-job-listing-url"
	}
	spiderName := "linkedin.com"
	fileName := values["file_name"]
	if fileName == "" {
		fileName = "{{TasksID}}"
	}
	delete(values, "spider_id")
	delete(values, "file_name")
	spiderParameters, err := json.Marshal([]map[string]string{values})
	if err != nil {
		return nil, err
	}
	payload := map[string]string{
		"spider_name":       spiderName,
		"spider_id":         spiderID,
		"spider_parameters": string(spiderParameters),
		"spider_errors":     "true",
		"file_name":         fileName,
	}
	return s.client.doForm(ctx, s.client.scraperBaseURL, "/builder?platform=1", payload)
}

type RedditCommentRequest struct {
	URL          string `json:"url,omitempty"`
	DaysBack     string `json:"days_back,omitempty"`
	CommentLimit string `json:"comment_limit,omitempty"`
	FileName     string `json:"file_name,omitempty"`
}

func (r RedditCommentRequest) params() map[string]string {
	values := map[string]string{}
	addStringParam(values, "url", r.URL)
	addStringParam(values, "days_back", r.DaysBack)
	addStringParam(values, "comment_limit", r.CommentLimit)
	addStringParam(values, "file_name", r.FileName)
	return values
}

func (s *RedditService) Comment(ctx context.Context, req RedditCommentRequest) (*RawResponse, error) {
	values := req.params()
	defaultParam(values, "url", "https://www.reddit.com/r/datascience/comments/1cmnf0m/comment/l32204i/?utm_source=share&utm_medium=web3x&utm_name=web3xcss&utm_term=1&utm_content=share_button")
	defaultParam(values, "days_back", "10")
	defaultParam(values, "comment_limit", "5")
	defaultParam(values, "file_name", "{{TasksID}}")
	if err := requireParam(values, "url"); err != nil {
		return nil, err
	}
	spiderID := "reddit_comment_by-url"
	spiderName := "reddit.com"
	fileName := values["file_name"]
	if fileName == "" {
		fileName = "{{TasksID}}"
	}
	delete(values, "spider_id")
	delete(values, "file_name")
	spiderParameters, err := json.Marshal([]map[string]string{values})
	if err != nil {
		return nil, err
	}
	payload := map[string]string{
		"spider_name":       spiderName,
		"spider_id":         spiderID,
		"spider_parameters": string(spiderParameters),
		"spider_errors":     "true",
		"file_name":         fileName,
	}
	return s.client.doForm(ctx, s.client.scraperBaseURL, "/builder?platform=1", payload)
}

type RedditPostsRequest struct {
	SpiderID   string `json:"spider_id,omitempty"`
	URL        string `json:"url,omitempty"`
	Keyword    string `json:"keyword,omitempty"`
	NumOfPosts string `json:"num_of_posts,omitempty"`
	SortBy     string `json:"sort_by,omitempty"`
	SortByTime string `json:"sort_by_time,omitempty"`
	FileName   string `json:"file_name,omitempty"`
}

func (r RedditPostsRequest) params() map[string]string {
	values := map[string]string{}
	addStringParam(values, "spider_id", r.SpiderID)
	addStringParam(values, "url", r.URL)
	addStringParam(values, "keyword", r.Keyword)
	addStringParam(values, "num_of_posts", r.NumOfPosts)
	addStringParam(values, "sort_by", r.SortBy)
	addStringParam(values, "sort_by_time", r.SortByTime)
	addStringParam(values, "file_name", r.FileName)
	return values
}

func (s *RedditService) Posts(ctx context.Context, req RedditPostsRequest) (*RawResponse, error) {
	values := req.params()
	defaultParam(values, "spider_id", "reddit_posts_by-url")
	defaultParam(values, "file_name", "{{TasksID}}")
	if err := requireParam(values, "spider_id"); err != nil {
		return nil, err
	}
	spiderID := values["spider_id"]
	if spiderID == "" {
		spiderID = "reddit_posts_by-url"
	}
	spiderName := "reddit.com"
	fileName := values["file_name"]
	if fileName == "" {
		fileName = "{{TasksID}}"
	}
	delete(values, "spider_id")
	delete(values, "file_name")
	spiderParameters, err := json.Marshal([]map[string]string{values})
	if err != nil {
		return nil, err
	}
	payload := map[string]string{
		"spider_name":       spiderName,
		"spider_id":         spiderID,
		"spider_parameters": string(spiderParameters),
		"spider_errors":     "true",
		"file_name":         fileName,
	}
	return s.client.doForm(ctx, s.client.scraperBaseURL, "/builder?platform=1", payload)
}

type RequestWebUnlockerRequest struct {
	URL            string `json:"url,omitempty"`
	Type           string `json:"type,omitempty"`
	JsRender       string `json:"js_render,omitempty"`
	BlockResources string `json:"block_resources,omitempty"`
	CleanContent   string `json:"clean_content,omitempty"`
	Country        string `json:"country,omitempty"`
	Headers        string `json:"headers,omitempty"`
	Cookies        string `json:"cookies,omitempty"`
	Wait           string `json:"wait,omitempty"`
	WaitFor        string `json:"wait_for,omitempty"`
	FollowRedirect string `json:"follow_redirect,omitempty"`
}

func (r RequestWebUnlockerRequest) params() map[string]string {
	values := map[string]string{}
	addStringParam(values, "url", r.URL)
	addStringParam(values, "type", r.Type)
	addStringParam(values, "js_render", r.JsRender)
	addStringParam(values, "block_resources", r.BlockResources)
	addStringParam(values, "clean_content", r.CleanContent)
	addStringParam(values, "country", r.Country)
	addStringParam(values, "headers", r.Headers)
	addStringParam(values, "cookies", r.Cookies)
	addStringParam(values, "wait", r.Wait)
	addStringParam(values, "wait_for", r.WaitFor)
	addStringParam(values, "follow_redirect", r.FollowRedirect)
	return values
}

func (s *WebUnlockerService) Request(ctx context.Context, req RequestWebUnlockerRequest) (*RawResponse, error) {
	values := req.params()
	defaultParam(values, "type", "html")
	defaultParam(values, "js_render", "True")
	defaultParam(values, "country", "us")
	defaultParam(values, "follow_redirect", "True")
	if err := requireParam(values, "url"); err != nil {
		return nil, err
	}
	values["isjson"] = "1"
	return s.client.doJSON(ctx, s.client.webUnlockerBaseURL, "/request", values)
}

type TikTokCommentRequest struct {
	URL         string `json:"url,omitempty"`
	PageTurning string `json:"page_turning,omitempty"`
	FileName    string `json:"file_name,omitempty"`
}

func (r TikTokCommentRequest) params() map[string]string {
	values := map[string]string{}
	addStringParam(values, "url", r.URL)
	addStringParam(values, "page_turning", r.PageTurning)
	addStringParam(values, "file_name", r.FileName)
	return values
}

func (s *TikTokService) Comment(ctx context.Context, req TikTokCommentRequest) (*RawResponse, error) {
	values := req.params()
	defaultParam(values, "url", "https://www.tiktok.com/@heymrcat/video/7216019547806092550")
	defaultParam(values, "page_turning", "1")
	defaultParam(values, "file_name", "{{TasksID}}")
	if err := requireParam(values, "url"); err != nil {
		return nil, err
	}
	if err := requireParam(values, "page_turning"); err != nil {
		return nil, err
	}
	spiderID := "tiktok_comment_by-url"
	spiderName := "tiktok.com"
	fileName := values["file_name"]
	if fileName == "" {
		fileName = "{{TasksID}}"
	}
	delete(values, "spider_id")
	delete(values, "file_name")
	spiderParameters, err := json.Marshal([]map[string]string{values})
	if err != nil {
		return nil, err
	}
	payload := map[string]string{
		"spider_name":       spiderName,
		"spider_id":         spiderID,
		"spider_parameters": string(spiderParameters),
		"spider_errors":     "true",
		"file_name":         fileName,
	}
	return s.client.doForm(ctx, s.client.scraperBaseURL, "/builder?platform=1", payload)
}

type TikTokPostsRequest struct {
	URL        string `json:"url,omitempty"`
	NumOfPosts string `json:"num_of_posts,omitempty"`
	FileName   string `json:"file_name,omitempty"`
}

func (r TikTokPostsRequest) params() map[string]string {
	values := map[string]string{}
	addStringParam(values, "url", r.URL)
	addStringParam(values, "num_of_posts", r.NumOfPosts)
	addStringParam(values, "file_name", r.FileName)
	return values
}

func (s *TikTokService) Posts(ctx context.Context, req TikTokPostsRequest) (*RawResponse, error) {
	values := req.params()
	defaultParam(values, "url", "https://www.tiktok.com/discover/dog")
	defaultParam(values, "num_of_posts", "5")
	defaultParam(values, "file_name", "{{TasksID}}")
	if err := requireParam(values, "url"); err != nil {
		return nil, err
	}
	spiderID := "tiktok_posts_by-listurl"
	spiderName := "tiktok.com"
	fileName := values["file_name"]
	if fileName == "" {
		fileName = "{{TasksID}}"
	}
	delete(values, "spider_id")
	delete(values, "file_name")
	spiderParameters, err := json.Marshal([]map[string]string{values})
	if err != nil {
		return nil, err
	}
	payload := map[string]string{
		"spider_name":       spiderName,
		"spider_id":         spiderID,
		"spider_parameters": string(spiderParameters),
		"spider_errors":     "true",
		"file_name":         fileName,
	}
	return s.client.doForm(ctx, s.client.scraperBaseURL, "/builder?platform=1", payload)
}

type TikTokProfilesRequest struct {
	SpiderID    string `json:"spider_id,omitempty"`
	URL         string `json:"url,omitempty"`
	SearchURL   string `json:"search_url,omitempty"`
	Country     string `json:"country,omitempty"`
	PageTurning string `json:"page_turning,omitempty"`
	FileName    string `json:"file_name,omitempty"`
}

func (r TikTokProfilesRequest) params() map[string]string {
	values := map[string]string{}
	addStringParam(values, "spider_id", r.SpiderID)
	addStringParam(values, "url", r.URL)
	addStringParam(values, "search_url", r.SearchURL)
	addStringParam(values, "country", r.Country)
	addStringParam(values, "page_turning", r.PageTurning)
	addStringParam(values, "file_name", r.FileName)
	return values
}

func (s *TikTokService) Profiles(ctx context.Context, req TikTokProfilesRequest) (*RawResponse, error) {
	values := req.params()
	defaultParam(values, "spider_id", "tiktok_profiles_by-url")
	defaultParam(values, "file_name", "{{TasksID}}")
	if err := requireParam(values, "spider_id"); err != nil {
		return nil, err
	}
	spiderID := values["spider_id"]
	if spiderID == "" {
		spiderID = "tiktok_profiles_by-url"
	}
	spiderName := "tiktok.com"
	fileName := values["file_name"]
	if fileName == "" {
		fileName = "{{TasksID}}"
	}
	delete(values, "spider_id")
	delete(values, "file_name")
	spiderParameters, err := json.Marshal([]map[string]string{values})
	if err != nil {
		return nil, err
	}
	payload := map[string]string{
		"spider_name":       spiderName,
		"spider_id":         spiderID,
		"spider_parameters": string(spiderParameters),
		"spider_errors":     "true",
		"file_name":         fileName,
	}
	return s.client.doForm(ctx, s.client.scraperBaseURL, "/builder?platform=1", payload)
}

type TikTokShopRequest struct {
	URL      string `json:"url,omitempty"`
	FileName string `json:"file_name,omitempty"`
}

func (r TikTokShopRequest) params() map[string]string {
	values := map[string]string{}
	addStringParam(values, "url", r.URL)
	addStringParam(values, "file_name", r.FileName)
	return values
}

func (s *TikTokService) Shop(ctx context.Context, req TikTokShopRequest) (*RawResponse, error) {
	values := req.params()
	defaultParam(values, "url", "https://www.tiktok.com/shop/pdp/long-sleeve-crew-neck-tee-3-pack-by-galaxy-by-harvic-cotton-blend/1729461570693075200")
	defaultParam(values, "file_name", "{{TasksID}}")
	if err := requireParam(values, "url"); err != nil {
		return nil, err
	}
	spiderID := "tiktok_shop_by-url"
	spiderName := "tiktok.com"
	fileName := values["file_name"]
	if fileName == "" {
		fileName = "{{TasksID}}"
	}
	delete(values, "spider_id")
	delete(values, "file_name")
	spiderParameters, err := json.Marshal([]map[string]string{values})
	if err != nil {
		return nil, err
	}
	payload := map[string]string{
		"spider_name":       spiderName,
		"spider_id":         spiderID,
		"spider_parameters": string(spiderParameters),
		"spider_errors":     "true",
		"file_name":         fileName,
	}
	return s.client.doForm(ctx, s.client.scraperBaseURL, "/builder?platform=1", payload)
}

type TwitterPostRequest struct {
	URL      string `json:"url,omitempty"`
	FileName string `json:"file_name,omitempty"`
}

func (r TwitterPostRequest) params() map[string]string {
	values := map[string]string{}
	addStringParam(values, "url", r.URL)
	addStringParam(values, "file_name", r.FileName)
	return values
}

func (s *TwitterService) Post(ctx context.Context, req TwitterPostRequest) (*RawResponse, error) {
	values := req.params()
	defaultParam(values, "url", "https://x.com/elonmusk")
	defaultParam(values, "file_name", "{{TasksID}}")
	if err := requireParam(values, "url"); err != nil {
		return nil, err
	}
	spiderID := "twitter_post_by-profileurl"
	spiderName := "x.com"
	fileName := values["file_name"]
	if fileName == "" {
		fileName = "{{TasksID}}"
	}
	delete(values, "spider_id")
	delete(values, "file_name")
	spiderParameters, err := json.Marshal([]map[string]string{values})
	if err != nil {
		return nil, err
	}
	payload := map[string]string{
		"spider_name":       spiderName,
		"spider_id":         spiderID,
		"spider_parameters": string(spiderParameters),
		"spider_errors":     "true",
		"file_name":         fileName,
	}
	return s.client.doForm(ctx, s.client.scraperBaseURL, "/builder?platform=1", payload)
}

type TwitterProfileRequest struct {
	SpiderID string `json:"spider_id,omitempty"`
	URL      string `json:"url,omitempty"`
	UserName string `json:"user_name,omitempty"`
	FileName string `json:"file_name,omitempty"`
}

func (r TwitterProfileRequest) params() map[string]string {
	values := map[string]string{}
	addStringParam(values, "spider_id", r.SpiderID)
	addStringParam(values, "url", r.URL)
	addStringParam(values, "user_name", r.UserName)
	addStringParam(values, "file_name", r.FileName)
	return values
}

func (s *TwitterService) Profile(ctx context.Context, req TwitterProfileRequest) (*RawResponse, error) {
	values := req.params()
	defaultParam(values, "spider_id", "twitter_profile_by-profileurl")
	defaultParam(values, "file_name", "{{TasksID}}")
	if err := requireParam(values, "spider_id"); err != nil {
		return nil, err
	}
	spiderID := values["spider_id"]
	if spiderID == "" {
		spiderID = "twitter_profile_by-profileurl"
	}
	spiderName := "x.com"
	fileName := values["file_name"]
	if fileName == "" {
		fileName = "{{TasksID}}"
	}
	delete(values, "spider_id")
	delete(values, "file_name")
	spiderParameters, err := json.Marshal([]map[string]string{values})
	if err != nil {
		return nil, err
	}
	payload := map[string]string{
		"spider_name":       spiderName,
		"spider_id":         spiderID,
		"spider_parameters": string(spiderParameters),
		"spider_errors":     "true",
		"file_name":         fileName,
	}
	return s.client.doForm(ctx, s.client.scraperBaseURL, "/builder?platform=1", payload)
}

type WalmartProductRequest struct {
	SpiderID      string `json:"spider_id,omitempty"`
	URL           string `json:"url,omitempty"`
	CategoryURL   string `json:"category_url,omitempty"`
	Sku           string `json:"sku,omitempty"`
	Keyword       string `json:"keyword,omitempty"`
	Domain        string `json:"domain,omitempty"`
	AllVariations string `json:"all_variations,omitempty"`
	PageTurning   string `json:"page_turning,omitempty"`
	FileName      string `json:"file_name,omitempty"`
}

func (r WalmartProductRequest) params() map[string]string {
	values := map[string]string{}
	addStringParam(values, "spider_id", r.SpiderID)
	addStringParam(values, "url", r.URL)
	addStringParam(values, "category_url", r.CategoryURL)
	addStringParam(values, "sku", r.Sku)
	addStringParam(values, "keyword", r.Keyword)
	addStringParam(values, "domain", r.Domain)
	addStringParam(values, "all_variations", r.AllVariations)
	addStringParam(values, "page_turning", r.PageTurning)
	addStringParam(values, "file_name", r.FileName)
	return values
}

func (s *WalmartService) Product(ctx context.Context, req WalmartProductRequest) (*RawResponse, error) {
	values := req.params()
	defaultParam(values, "spider_id", "walmart_product_by-url")
	defaultParam(values, "file_name", "{{TasksID}}")
	if err := requireParam(values, "spider_id"); err != nil {
		return nil, err
	}
	spiderID := values["spider_id"]
	if spiderID == "" {
		spiderID = "walmart_product_by-url"
	}
	spiderName := "walmart.com"
	fileName := values["file_name"]
	if fileName == "" {
		fileName = "{{TasksID}}"
	}
	delete(values, "spider_id")
	delete(values, "file_name")
	spiderParameters, err := json.Marshal([]map[string]string{values})
	if err != nil {
		return nil, err
	}
	payload := map[string]string{
		"spider_name":       spiderName,
		"spider_id":         spiderID,
		"spider_parameters": string(spiderParameters),
		"spider_errors":     "true",
		"file_name":         fileName,
	}
	return s.client.doForm(ctx, s.client.scraperBaseURL, "/builder?platform=1", payload)
}

type YandexSearchRequest struct {
	Text         string `json:"text,omitempty"`
	JSON         string `json:"json,omitempty"`
	YandexDomain string `json:"yandex_domain,omitempty"`
	Lang         string `json:"lang,omitempty"`
	Lr           string `json:"lr,omitempty"`
	P            string `json:"p,omitempty"`
	FamilyMode   string `json:"family_mode,omitempty"`
	FixTypo      string `json:"fix_typo,omitempty"`
	GroupsOnPage string `json:"groups_on_page,omitempty"`
	NoCache      string `json:"no_cache,omitempty"`
}

func (r YandexSearchRequest) params() map[string]string {
	values := map[string]string{}
	addStringParam(values, "text", r.Text)
	addStringParam(values, "json", r.JSON)
	addStringParam(values, "yandex_domain", r.YandexDomain)
	addStringParam(values, "lang", r.Lang)
	addStringParam(values, "lr", r.Lr)
	addStringParam(values, "p", r.P)
	addStringParam(values, "family_mode", r.FamilyMode)
	addStringParam(values, "fix_typo", r.FixTypo)
	addStringParam(values, "groups_on_page", r.GroupsOnPage)
	addStringParam(values, "no_cache", r.NoCache)
	return values
}

func (s *YandexService) Search(ctx context.Context, req YandexSearchRequest) (*RawResponse, error) {
	values := req.params()
	defaultParam(values, "json", "1")
	defaultParam(values, "yandex_domain", "yandex.com")
	defaultParam(values, "lang", "en")
	defaultParam(values, "p", "0")
	defaultParam(values, "family_mode", "1")
	defaultParam(values, "fix_typo", "true")
	defaultParam(values, "groups_on_page", "10")
	defaultParam(values, "no_cache", "false")
	if err := requireParam(values, "text"); err != nil {
		return nil, err
	}
	if err := requireParam(values, "json"); err != nil {
		return nil, err
	}
	values["engine"] = "yandex"
	return s.client.doForm(ctx, s.client.serpBaseURL, "/request", values)
}

type YouTubeAudioRequest struct {
	URL               string `json:"url,omitempty"`
	SubtitlesLanguage string `json:"subtitles_language,omitempty"`
	SelectedOnly      string `json:"selected_only,omitempty"`
	Kilohertz         string `json:"kilohertz,omitempty"`
	IsSubtitles       string `json:"is_subtitles,omitempty"`
	AudioFormat       string `json:"audio_format,omitempty"`
	Bitrate           string `json:"bitrate,omitempty"`
	FileName          string `json:"file_name,omitempty"`
}

func (r YouTubeAudioRequest) params() map[string]string {
	values := map[string]string{}
	addStringParam(values, "url", r.URL)
	addStringParam(values, "subtitles_language", r.SubtitlesLanguage)
	addStringParam(values, "selected_only", r.SelectedOnly)
	addStringParam(values, "kilohertz", r.Kilohertz)
	addStringParam(values, "is_subtitles", r.IsSubtitles)
	addStringParam(values, "audio_format", r.AudioFormat)
	addStringParam(values, "bitrate", r.Bitrate)
	addStringParam(values, "file_name", r.FileName)
	return values
}

func (s *YouTubeService) Audio(ctx context.Context, req YouTubeAudioRequest) (*RawResponse, error) {
	values := req.params()
	defaultParam(values, "url", "https://www.youtube.com/watch?v=_SdpvpvVrLY")
	defaultParam(values, "subtitles_language", "ab")
	defaultParam(values, "selected_only", "false")
	defaultParam(values, "kilohertz", "<=48000")
	defaultParam(values, "is_subtitles", "false")
	defaultParam(values, "audio_format", "opus")
	defaultParam(values, "bitrate", "<=320")
	defaultParam(values, "file_name", "{{TasksID}}")
	if err := requireParam(values, "url"); err != nil {
		return nil, err
	}
	spiderID := "youtube_audio_by-url"
	spiderName := "youtube.com"
	fileName := values["file_name"]
	if fileName == "" {
		fileName = "{{TasksID}}"
	}
	delete(values, "spider_id")
	delete(values, "file_name")
	spiderParameters, err := json.Marshal([]map[string]string{values})
	if err != nil {
		return nil, err
	}
	payload := map[string]string{
		"spider_name":       spiderName,
		"spider_id":         spiderID,
		"spider_parameters": string(spiderParameters),
		"spider_errors":     "true",
		"file_name":         fileName,
	}
	return s.client.doForm(ctx, s.client.scraperBaseURL, "/builder?platform=1", payload)
}

type YouTubeCommentRequest struct {
	VideoID       string `json:"video_id,omitempty"`
	LoadReplies   string `json:"load_replies,omitempty"`
	NumOfComments string `json:"num_of_comments,omitempty"`
	FileName      string `json:"file_name,omitempty"`
}

func (r YouTubeCommentRequest) params() map[string]string {
	values := map[string]string{}
	addStringParam(values, "video_id", r.VideoID)
	addStringParam(values, "load_replies", r.LoadReplies)
	addStringParam(values, "num_of_comments", r.NumOfComments)
	addStringParam(values, "file_name", r.FileName)
	return values
}

func (s *YouTubeService) Comment(ctx context.Context, req YouTubeCommentRequest) (*RawResponse, error) {
	values := req.params()
	defaultParam(values, "video_id", "8RePenzQH80")
	defaultParam(values, "load_replies", "10")
	defaultParam(values, "num_of_comments", "Top comments")
	defaultParam(values, "file_name", "{{TasksID}}")
	if err := requireParam(values, "video_id"); err != nil {
		return nil, err
	}
	if err := requireParam(values, "load_replies"); err != nil {
		return nil, err
	}
	if err := requireParam(values, "num_of_comments"); err != nil {
		return nil, err
	}
	spiderID := "youtube_comment_by-id"
	spiderName := "youtube.com"
	fileName := values["file_name"]
	if fileName == "" {
		fileName = "{{TasksID}}"
	}
	delete(values, "spider_id")
	delete(values, "file_name")
	spiderParameters, err := json.Marshal([]map[string]string{values})
	if err != nil {
		return nil, err
	}
	payload := map[string]string{
		"spider_name":       spiderName,
		"spider_id":         spiderID,
		"spider_parameters": string(spiderParameters),
		"spider_errors":     "true",
		"file_name":         fileName,
	}
	return s.client.doForm(ctx, s.client.scraperBaseURL, "/builder?platform=1", payload)
}

type YouTubeProductRequest struct {
	VideoID           string `json:"video_id,omitempty"`
	SubtitlesLanguage string `json:"subtitles_language,omitempty"`
	SubtitlesType     string `json:"subtitles_type,omitempty"`
	SelectedOnly      string `json:"selected_only,omitempty"`
	FileName          string `json:"file_name,omitempty"`
}

func (r YouTubeProductRequest) params() map[string]string {
	values := map[string]string{}
	addStringParam(values, "video_id", r.VideoID)
	addStringParam(values, "subtitles_language", r.SubtitlesLanguage)
	addStringParam(values, "subtitles_type", r.SubtitlesType)
	addStringParam(values, "selected_only", r.SelectedOnly)
	addStringParam(values, "file_name", r.FileName)
	return values
}

func (s *YouTubeService) Product(ctx context.Context, req YouTubeProductRequest) (*RawResponse, error) {
	values := req.params()
	defaultParam(values, "video_id", "8RePenzQH80")
	defaultParam(values, "subtitles_language", "ab")
	defaultParam(values, "subtitles_type", "auto_generated")
	defaultParam(values, "selected_only", "false")
	defaultParam(values, "file_name", "{{TasksID}}")
	if err := requireParam(values, "video_id"); err != nil {
		return nil, err
	}
	spiderID := "youtube_product_by-id"
	spiderName := "youtube.com"
	fileName := values["file_name"]
	if fileName == "" {
		fileName = "{{TasksID}}"
	}
	delete(values, "spider_id")
	delete(values, "file_name")
	spiderParameters, err := json.Marshal([]map[string]string{values})
	if err != nil {
		return nil, err
	}
	payload := map[string]string{
		"spider_name":       spiderName,
		"spider_id":         spiderID,
		"spider_parameters": string(spiderParameters),
		"spider_errors":     "true",
		"file_name":         fileName,
	}
	return s.client.doForm(ctx, s.client.scraperBaseURL, "/builder?platform=1", payload)
}

type YouTubeProfilesRequest struct {
	SpiderID    string `json:"spider_id,omitempty"`
	Keyword     string `json:"keyword,omitempty"`
	PageTurning string `json:"page_turning,omitempty"`
	URL         string `json:"url,omitempty"`
	FileName    string `json:"file_name,omitempty"`
}

func (r YouTubeProfilesRequest) params() map[string]string {
	values := map[string]string{}
	addStringParam(values, "spider_id", r.SpiderID)
	addStringParam(values, "keyword", r.Keyword)
	addStringParam(values, "page_turning", r.PageTurning)
	addStringParam(values, "url", r.URL)
	addStringParam(values, "file_name", r.FileName)
	return values
}

func (s *YouTubeService) Profiles(ctx context.Context, req YouTubeProfilesRequest) (*RawResponse, error) {
	values := req.params()
	defaultParam(values, "spider_id", "youtube_profiles_by-keyword")
	defaultParam(values, "file_name", "{{TasksID}}")
	if err := requireParam(values, "spider_id"); err != nil {
		return nil, err
	}
	spiderID := values["spider_id"]
	if spiderID == "" {
		spiderID = "youtube_profiles_by-keyword"
	}
	spiderName := "youtube.com"
	fileName := values["file_name"]
	if fileName == "" {
		fileName = "{{TasksID}}"
	}
	delete(values, "spider_id")
	delete(values, "file_name")
	spiderParameters, err := json.Marshal([]map[string]string{values})
	if err != nil {
		return nil, err
	}
	payload := map[string]string{
		"spider_name":       spiderName,
		"spider_id":         spiderID,
		"spider_parameters": string(spiderParameters),
		"spider_errors":     "true",
		"file_name":         fileName,
	}
	return s.client.doForm(ctx, s.client.scraperBaseURL, "/builder?platform=1", payload)
}

type YouTubeTranscriptRequest struct {
	VideoID           string `json:"video_id,omitempty"`
	SubtitlesLanguage string `json:"subtitles_language,omitempty"`
	SubtitlesType     string `json:"subtitles_type,omitempty"`
	SelectedOnly      string `json:"selected_only,omitempty"`
	FileName          string `json:"file_name,omitempty"`
}

func (r YouTubeTranscriptRequest) params() map[string]string {
	values := map[string]string{}
	addStringParam(values, "video_id", r.VideoID)
	addStringParam(values, "subtitles_language", r.SubtitlesLanguage)
	addStringParam(values, "subtitles_type", r.SubtitlesType)
	addStringParam(values, "selected_only", r.SelectedOnly)
	addStringParam(values, "file_name", r.FileName)
	return values
}

func (s *YouTubeService) Transcript(ctx context.Context, req YouTubeTranscriptRequest) (*RawResponse, error) {
	values := req.params()
	defaultParam(values, "video_id", "8RePenzQH80")
	defaultParam(values, "subtitles_language", "ab")
	defaultParam(values, "subtitles_type", "auto_generated")
	defaultParam(values, "selected_only", "false")
	defaultParam(values, "file_name", "{{TasksID}}")
	if err := requireParam(values, "video_id"); err != nil {
		return nil, err
	}
	spiderID := "youtube_transcript_by-id"
	spiderName := "youtube.com"
	fileName := values["file_name"]
	if fileName == "" {
		fileName = "{{TasksID}}"
	}
	delete(values, "spider_id")
	delete(values, "file_name")
	spiderParameters, err := json.Marshal([]map[string]string{values})
	if err != nil {
		return nil, err
	}
	payload := map[string]string{
		"spider_name":       spiderName,
		"spider_id":         spiderID,
		"spider_parameters": string(spiderParameters),
		"spider_errors":     "true",
		"file_name":         fileName,
	}
	return s.client.doForm(ctx, s.client.scraperBaseURL, "/builder?platform=1", payload)
}

type YouTubeVideoRequest struct {
	URL               string `json:"url,omitempty"`
	SubtitlesLanguage string `json:"subtitles_language,omitempty"`
	SelectedOnly      string `json:"selected_only,omitempty"`
	Resolution        string `json:"resolution,omitempty"`
	VideoCodec        string `json:"video_codec,omitempty"`
	AudioFormat       string `json:"audio_format,omitempty"`
	Bitrate           string `json:"bitrate,omitempty"`
	FileName          string `json:"file_name,omitempty"`
}

func (r YouTubeVideoRequest) params() map[string]string {
	values := map[string]string{}
	addStringParam(values, "url", r.URL)
	addStringParam(values, "subtitles_language", r.SubtitlesLanguage)
	addStringParam(values, "selected_only", r.SelectedOnly)
	addStringParam(values, "resolution", r.Resolution)
	addStringParam(values, "video_codec", r.VideoCodec)
	addStringParam(values, "audio_format", r.AudioFormat)
	addStringParam(values, "bitrate", r.Bitrate)
	addStringParam(values, "file_name", r.FileName)
	return values
}

func (s *YouTubeService) Video(ctx context.Context, req YouTubeVideoRequest) (*RawResponse, error) {
	values := req.params()
	defaultParam(values, "url", "https://www.youtube.com/watch?v=_SdpvpvVrLY")
	defaultParam(values, "subtitles_language", "ab")
	defaultParam(values, "selected_only", "false")
	defaultParam(values, "resolution", "<=360p")
	defaultParam(values, "video_codec", "vp9")
	defaultParam(values, "audio_format", "opus")
	defaultParam(values, "bitrate", "<=320")
	defaultParam(values, "file_name", "{{TasksID}}")
	if err := requireParam(values, "url"); err != nil {
		return nil, err
	}
	spiderID := "youtube_video_by-url"
	spiderName := "youtube.com"
	fileName := values["file_name"]
	if fileName == "" {
		fileName = "{{TasksID}}"
	}
	delete(values, "spider_id")
	delete(values, "file_name")
	spiderParameters, err := json.Marshal([]map[string]string{values})
	if err != nil {
		return nil, err
	}
	payload := map[string]string{
		"spider_name":       spiderName,
		"spider_id":         spiderID,
		"spider_parameters": string(spiderParameters),
		"spider_errors":     "true",
		"file_name":         fileName,
	}
	return s.client.doForm(ctx, s.client.scraperBaseURL, "/builder?platform=1", payload)
}

type YouTubeVideoPostRequest struct {
	SpiderID      string `json:"spider_id,omitempty"`
	URL           string `json:"url,omitempty"`
	OrderBy       string `json:"order_by,omitempty"`
	StartIndex    string `json:"start_index,omitempty"`
	NumOfPosts    string `json:"num_of_posts,omitempty"`
	KeywordSearch string `json:"keyword_search,omitempty"`
	Features      string `json:"features,omitempty"`
	Type          string `json:"type,omitempty"`
	Duration      string `json:"duration,omitempty"`
	UploadDate    string `json:"upload_date,omitempty"`
	Hashtag       string `json:"hashtag,omitempty"`
	Keyword       string `json:"keyword,omitempty"`
	AllTabs       string `json:"all_tabs,omitempty"`
	FileName      string `json:"file_name,omitempty"`
}

func (r YouTubeVideoPostRequest) params() map[string]string {
	values := map[string]string{}
	addStringParam(values, "spider_id", r.SpiderID)
	addStringParam(values, "url", r.URL)
	addStringParam(values, "order_by", r.OrderBy)
	addStringParam(values, "start_index", r.StartIndex)
	addStringParam(values, "num_of_posts", r.NumOfPosts)
	addStringParam(values, "keyword_search", r.KeywordSearch)
	addStringParam(values, "features", r.Features)
	addStringParam(values, "type", r.Type)
	addStringParam(values, "duration", r.Duration)
	addStringParam(values, "upload_date", r.UploadDate)
	addStringParam(values, "hashtag", r.Hashtag)
	addStringParam(values, "keyword", r.Keyword)
	addStringParam(values, "all_tabs", r.AllTabs)
	addStringParam(values, "file_name", r.FileName)
	return values
}

func (s *YouTubeService) VideoPost(ctx context.Context, req YouTubeVideoPostRequest) (*RawResponse, error) {
	values := req.params()
	defaultParam(values, "spider_id", "youtube_video-post_by-url")
	defaultParam(values, "file_name", "{{TasksID}}")
	if values["spider_id"] == "youtube_video-post_by-url" {
		defaultParam(values, "url", "https://www.youtube.com/@stephcurry/videos")
		defaultParam(values, "order_by", "Latest")
		defaultParam(values, "start_index", "1")
		defaultParam(values, "num_of_posts", "5")
	}
	if err := requireParam(values, "spider_id"); err != nil {
		return nil, err
	}
	spiderID := values["spider_id"]
	if spiderID == "" {
		spiderID = "youtube_video-post_by-url"
	}
	spiderName := "youtube.com"
	fileName := values["file_name"]
	if fileName == "" {
		fileName = "{{TasksID}}"
	}
	delete(values, "spider_id")
	delete(values, "file_name")
	spiderParameters, err := json.Marshal([]map[string]string{values})
	if err != nil {
		return nil, err
	}
	payload := map[string]string{
		"spider_name":       spiderName,
		"spider_id":         spiderID,
		"spider_parameters": string(spiderParameters),
		"spider_errors":     "true",
		"file_name":         fileName,
	}
	return s.client.doForm(ctx, s.client.scraperBaseURL, "/builder?platform=1", payload)
}

type ZillowProductRequest struct {
	KeywordsLocation string `json:"keywords-location,omitempty"`
	Listingcategory  string `json:"listingCategory,omitempty"`
	Hometype         string `json:"HomeType,omitempty"`
	DaysOnZillow     string `json:"days_on_zillow,omitempty"`
	Maximum          string `json:"maximum,omitempty"`
	FileName         string `json:"file_name,omitempty"`
}

func (r ZillowProductRequest) params() map[string]string {
	values := map[string]string{}
	addStringParam(values, "keywords-location", r.KeywordsLocation)
	addStringParam(values, "listingCategory", r.Listingcategory)
	addStringParam(values, "HomeType", r.Hometype)
	addStringParam(values, "days_on_zillow", r.DaysOnZillow)
	addStringParam(values, "maximum", r.Maximum)
	addStringParam(values, "file_name", r.FileName)
	return values
}

func (s *ZillowService) Product(ctx context.Context, req ZillowProductRequest) (*RawResponse, error) {
	values := req.params()
	defaultParam(values, "keywords-location", "South Bend")
	defaultParam(values, "listingCategory", "For Rent")
	defaultParam(values, "HomeType", "Houses")
	defaultParam(values, "days_on_zillow", "Any")
	defaultParam(values, "maximum", "10")
	defaultParam(values, "file_name", "{{TasksID}}")
	if err := requireParam(values, "keywords-location"); err != nil {
		return nil, err
	}
	spiderID := "zillow_product_by-filter"
	spiderName := "zillow.com"
	fileName := values["file_name"]
	if fileName == "" {
		fileName = "{{TasksID}}"
	}
	delete(values, "spider_id")
	delete(values, "file_name")
	spiderParameters, err := json.Marshal([]map[string]string{values})
	if err != nil {
		return nil, err
	}
	payload := map[string]string{
		"spider_name":       spiderName,
		"spider_id":         spiderID,
		"spider_parameters": string(spiderParameters),
		"spider_errors":     "true",
		"file_name":         fileName,
	}
	return s.client.doForm(ctx, s.client.scraperBaseURL, "/builder?platform=1", payload)
}

type ToolParamSpec struct {
	Name     string
	Type     string
	Required bool
	Default  string
	Min      string
	Max      string
}
type ToolSpec struct {
	Name        string
	Group       string
	Method      string
	ServiceKind string
	Engine      string
	SpiderName  string
	SpiderID    string
	Params      []ToolParamSpec
}

var ToolSpecs = []ToolSpec{
	{Name: "airbnb_product", Group: "Airbnb", Method: "Product", ServiceKind: "scraper_builder", Engine: "", SpiderName: "airbnb.com", SpiderID: "airbnb_product_by-searchurl", Params: []ToolParamSpec{
		{Name: "searchurl", Type: "string", Required: true, Default: "https://www.airbnb.com/s/Greece/homes?query=Greece&refinement_paths%5B%5D=%2Fhomes&place_id=ChIJY2xxEcdKWxMRHS2a3HUXOjY&flexible_trip_lengths%5B%5D=one_week&monthly_start_date=2025-03-01&monthly_length=3&monthly_end_date=2025-06-01&search_mode=regular_search&price_filter_input_type=0&channel=EXPLORE&date_picker_type=calendar&source=structured_search_input_header&search_type=filter_change&price_filter_num_nights=5&flexible_date_search_filter_type=1", Min: "", Max: ""},
		{Name: "country", Type: "string", Required: false, Default: "HK", Min: "", Max: ""},
		{Name: "file_name", Type: "string", Required: false, Default: "{{TasksID}}", Min: "", Max: ""},
	}},
	{Name: "amazon_comment", Group: "Amazon", Method: "Comment", ServiceKind: "scraper_builder", Engine: "", SpiderName: "amazon.com", SpiderID: "amazon_comment_by-url", Params: []ToolParamSpec{
		{Name: "url", Type: "string", Required: true, Default: "https://www.amazon.com/HISDERN-Checkered-Handkerchief-Classic-Necktie/dp/B0BRXPR726", Min: "", Max: ""},
		{Name: "file_name", Type: "string", Required: false, Default: "{{TasksID}}", Min: "", Max: ""},
	}},
	{Name: "amazon_global_product", Group: "Amazon", Method: "GlobalProduct", ServiceKind: "scraper_builder", Engine: "", SpiderName: "amazon.com", SpiderID: "amazon_global-product_by-url", Params: []ToolParamSpec{
		{Name: "spider_id", Type: "string", Required: true, Default: "amazon_global-product_by-url", Min: "", Max: ""},
		{Name: "url", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "maximum", Type: "string", Required: false, Default: "5", Min: "", Max: ""},
		{Name: "sort_by", Type: "string", Required: false, Default: "畅销排行", Min: "", Max: ""},
		{Name: "get_sponsored", Type: "string", Required: false, Default: "true", Min: "", Max: ""},
		{Name: "keyword", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "domain", Type: "string", Required: false, Default: "https://www.amazon.com", Min: "", Max: ""},
		{Name: "brands", Type: "string", Required: false, Default: "Adidas", Min: "", Max: ""},
		{Name: "page_turning", Type: "string", Required: false, Default: "2", Min: "", Max: ""},
		{Name: "lowest_price", Type: "string", Required: false, Default: "20", Min: "", Max: ""},
		{Name: "highest_price", Type: "string", Required: false, Default: "50", Min: "", Max: ""},
		{Name: "file_name", Type: "string", Required: false, Default: "{{TasksID}}", Min: "", Max: ""},
	}},
	{Name: "amazon_product", Group: "Amazon", Method: "Product", ServiceKind: "scraper_builder", Engine: "", SpiderName: "amazon.com", SpiderID: "amazon_product_by-asin", Params: []ToolParamSpec{
		{Name: "spider_id", Type: "string", Required: true, Default: "amazon_product_by-asin", Min: "", Max: ""},
		{Name: "asin", Type: "string", Required: false, Default: "B0BZYCJK89", Min: "", Max: ""},
		{Name: "url", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "category_url", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "keyword", Type: "string", Required: false, Default: "coffee", Min: "", Max: ""},
		{Name: "page_turning", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "lowest_price", Type: "string", Required: false, Default: "20", Min: "", Max: ""},
		{Name: "highest_price", Type: "string", Required: false, Default: "50", Min: "", Max: ""},
		{Name: "sort_by", Type: "string", Required: false, Default: "畅销排行", Min: "", Max: ""},
		{Name: "collect_subcategories", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "zip_code", Type: "string", Required: false, Default: "94107", Min: "", Max: ""},
		{Name: "file_name", Type: "string", Required: false, Default: "{{TasksID}}", Min: "", Max: ""},
	}},
	{Name: "amazon_product_list", Group: "Amazon", Method: "ProductList", ServiceKind: "scraper_builder", Engine: "", SpiderName: "amazon.com", SpiderID: "amazon_product-list_by-keywords-domain", Params: []ToolParamSpec{
		{Name: "keyword", Type: "string", Required: true, Default: "X-box", Min: "", Max: ""},
		{Name: "domain", Type: "string", Required: true, Default: "https://www.amazon.com/", Min: "", Max: ""},
		{Name: "page_turning", Type: "string", Required: false, Default: "1", Min: "", Max: ""},
		{Name: "file_name", Type: "string", Required: false, Default: "{{TasksID}}", Min: "", Max: ""},
	}},
	{Name: "amazon_seller", Group: "Amazon", Method: "Seller", ServiceKind: "scraper_builder", Engine: "", SpiderName: "amazon.com", SpiderID: "amazon_seller_by-url", Params: []ToolParamSpec{
		{Name: "url", Type: "string", Required: true, Default: "https://www.amazon.com/sp?ie=UTF8&seller=ADZ7LD48GVFQJ&asin=B07H56J7K1&ref_=dp_merchant_link&isAmazonFulfilled=1", Min: "", Max: ""},
		{Name: "file_name", Type: "string", Required: false, Default: "{{TasksID}}", Min: "", Max: ""},
	}},
	{Name: "bing_images", Group: "Bing", Method: "Images", ServiceKind: "serp", Engine: "bing_images", SpiderName: "", SpiderID: "", Params: []ToolParamSpec{
		{Name: "q", Type: "string", Required: true, Default: "Pizza", Min: "", Max: ""},
		{Name: "json", Type: "string", Required: true, Default: "1", Min: "", Max: ""},
		{Name: "mkt", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "cc", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "first", Type: "string", Required: false, Default: "0", Min: "", Max: ""},
		{Name: "count", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "imagesize", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "color2", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "photo", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "aspect", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "face", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "age", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "license", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "no_cache", Type: "string", Required: false, Default: "false", Min: "", Max: ""},
	}},
	{Name: "bing_maps", Group: "Bing", Method: "Maps", ServiceKind: "serp", Engine: "bing_maps", SpiderName: "", SpiderID: "", Params: []ToolParamSpec{
		{Name: "q", Type: "string", Required: true, Default: "", Min: "", Max: ""},
		{Name: "json", Type: "string", Required: true, Default: "1", Min: "", Max: ""},
		{Name: "cp", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "setlang", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "place_id", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "first", Type: "string", Required: false, Default: "0", Min: "", Max: ""},
		{Name: "count", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "no_cache", Type: "string", Required: false, Default: "false", Min: "", Max: ""},
	}},
	{Name: "bing_news", Group: "Bing", Method: "News", ServiceKind: "serp", Engine: "bing_news", SpiderName: "", SpiderID: "", Params: []ToolParamSpec{
		{Name: "q", Type: "string", Required: true, Default: "Pizza", Min: "", Max: ""},
		{Name: "json", Type: "string", Required: true, Default: "1", Min: "", Max: ""},
		{Name: "mkt", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "cc", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "first", Type: "string", Required: false, Default: "0", Min: "", Max: ""},
		{Name: "count", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "qft", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "safeSearch", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "no_cache", Type: "string", Required: false, Default: "false", Min: "", Max: ""},
	}},
	{Name: "bing_search", Group: "Bing", Method: "Search", ServiceKind: "serp", Engine: "bing", SpiderName: "", SpiderID: "", Params: []ToolParamSpec{
		{Name: "q", Type: "string", Required: true, Default: "Pizza", Min: "", Max: ""},
		{Name: "json", Type: "string", Required: true, Default: "1", Min: "", Max: ""},
		{Name: "location", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "lat", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "lon", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "mkt", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "cc", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "first", Type: "string", Required: false, Default: "0", Min: "", Max: ""},
		{Name: "safeSearch", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "filters", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "no_cache", Type: "string", Required: false, Default: "false", Min: "", Max: ""},
	}},
	{Name: "bing_shopping", Group: "Bing", Method: "Shopping", ServiceKind: "serp", Engine: "bing_shopping", SpiderName: "", SpiderID: "", Params: []ToolParamSpec{
		{Name: "q", Type: "string", Required: true, Default: "Pizza", Min: "", Max: ""},
		{Name: "json", Type: "string", Required: true, Default: "1", Min: "", Max: ""},
		{Name: "mkt", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "cc", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "efirst", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "filters", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "no_cache", Type: "string", Required: false, Default: "false", Min: "", Max: ""},
	}},
	{Name: "bing_videos", Group: "Bing", Method: "Videos", ServiceKind: "serp", Engine: "bing_videos", SpiderName: "", SpiderID: "", Params: []ToolParamSpec{
		{Name: "q", Type: "string", Required: true, Default: "Pizza", Min: "", Max: ""},
		{Name: "json", Type: "string", Required: true, Default: "1", Min: "", Max: ""},
		{Name: "mkt", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "cc", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "setlang", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "first", Type: "string", Required: false, Default: "0", Min: "", Max: ""},
		{Name: "length", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "date", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "resolution", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "source_site", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "price", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "no_cache", Type: "string", Required: false, Default: "false", Min: "", Max: ""},
	}},
	{Name: "booking_hotellist", Group: "Booking", Method: "Hotellist", ServiceKind: "scraper_builder", Engine: "", SpiderName: "booking.com", SpiderID: "booking_hotellist_by-url", Params: []ToolParamSpec{
		{Name: "url", Type: "string", Required: true, Default: "https://www.booking.com/hotel/gb/westlands-of-pitlochry.en-gb.html#tab-main", Min: "", Max: ""},
		{Name: "file_name", Type: "string", Required: false, Default: "{{TasksID}}", Min: "", Max: ""},
	}},
	{Name: "crunchbase_company", Group: "Crunchbase", Method: "Company", ServiceKind: "scraper_builder", Engine: "", SpiderName: "crunchbase.com", SpiderID: "crunchbase_company_by-url", Params: []ToolParamSpec{
		{Name: "spider_id", Type: "string", Required: true, Default: "crunchbase_company_by-url", Min: "", Max: ""},
		{Name: "url", Type: "string", Required: false, Default: "https://www.crunchbase.com/organization/aisci", Min: "", Max: ""},
		{Name: "keyword", Type: "string", Required: false, Default: "NetBooster", Min: "", Max: ""},
		{Name: "file_name", Type: "string", Required: false, Default: "{{TasksID}}", Min: "", Max: ""},
	}},
	{Name: "duckduckgo_search", Group: "DuckDuckGo", Method: "Search", ServiceKind: "serp", Engine: "duckduckgo", SpiderName: "", SpiderID: "", Params: []ToolParamSpec{
		{Name: "q", Type: "string", Required: true, Default: "Pizza", Min: "", Max: ""},
		{Name: "json", Type: "string", Required: true, Default: "1", Min: "", Max: ""},
		{Name: "kl", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "search_assist", Type: "string", Required: false, Default: "false", Min: "", Max: ""},
		{Name: "safe", Type: "string", Required: false, Default: "-1", Min: "", Max: ""},
		{Name: "df", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "start", Type: "string", Required: false, Default: "0", Min: "", Max: ""},
		{Name: "m", Type: "string", Required: false, Default: "10", Min: "", Max: ""},
		{Name: "no_cache", Type: "string", Required: false, Default: "false", Min: "", Max: ""},
	}},
	{Name: "ebay_info", Group: "EBay", Method: "Info", ServiceKind: "scraper_builder", Engine: "", SpiderName: "ebay.com", SpiderID: "ebay_ebay_by-url", Params: []ToolParamSpec{
		{Name: "spider_id", Type: "string", Required: true, Default: "ebay_ebay_by-url", Min: "", Max: ""},
		{Name: "url", Type: "string", Required: false, Default: "https://www.ebay.com/itm/296197468977?itmmeta=01HRWJ04NFHYT9AX0XZB8F18G1&hash=item44f6beb331%3Ag%3ADEQAAOSw3CxlhTJ%7E&_trkparms=%2526rpp_cid%253D6523c97b0b7882040b9472b6", Min: "", Max: ""},
		{Name: "keywords", Type: "string", Required: false, Default: "baby toys", Min: "", Max: ""},
		{Name: "count", Type: "string", Required: false, Default: "60", Min: "", Max: ""},
		{Name: "file_name", Type: "string", Required: false, Default: "{{TasksID}}", Min: "", Max: ""},
	}},
	{Name: "facebook_comment", Group: "Facebook", Method: "Comment", ServiceKind: "scraper_builder", Engine: "", SpiderName: "facebook.com", SpiderID: "facebook_comment_by-comments-url", Params: []ToolParamSpec{
		{Name: "url", Type: "string", Required: true, Default: "https://www.facebook.com/share/p/1K6xfHFkrK/", Min: "", Max: ""},
		{Name: "get_all_replies", Type: "string", Required: false, Default: "True", Min: "", Max: ""},
		{Name: "limit_records", Type: "string", Required: false, Default: "10", Min: "", Max: ""},
		{Name: "comments_sort", Type: "string", Required: false, Default: "所有评论", Min: "", Max: ""},
		{Name: "file_name", Type: "string", Required: false, Default: "{{TasksID}}", Min: "", Max: ""},
	}},
	{Name: "facebook_event", Group: "Facebook", Method: "Event", ServiceKind: "scraper_builder", Engine: "", SpiderName: "facebook.com", SpiderID: "facebook_event_by-eventlist-url", Params: []ToolParamSpec{
		{Name: "spider_id", Type: "string", Required: true, Default: "facebook_event_by-eventlist-url", Min: "", Max: ""},
		{Name: "url", Type: "string", Required: true, Default: "https://www.facebook.com/nohoclub/events", Min: "", Max: ""},
		{Name: "file_name", Type: "string", Required: false, Default: "{{TasksID}}", Min: "", Max: ""},
	}},
	{Name: "facebook_post", Group: "Facebook", Method: "Post", ServiceKind: "scraper_builder", Engine: "", SpiderName: "facebook.com", SpiderID: "facebook_post_by-posts-url", Params: []ToolParamSpec{
		{Name: "url", Type: "string", Required: true, Default: "https://www.facebook.com/permalink.php?story_fbid=pfbid0gNjZBhqCxSqj9xJS5aygNwqFqNEM2fYbTFKKbsvvGdEfTgFyAYWSckvkEHPqAE7gl&id=61574926580533&rdid=86oaujwNGCCdPLfj#", Min: "", Max: ""},
		{Name: "file_name", Type: "string", Required: false, Default: "{{TasksID}}", Min: "", Max: ""},
	}},
	{Name: "facebook_profile", Group: "Facebook", Method: "Profile", ServiceKind: "scraper_builder", Engine: "", SpiderName: "facebook.com", SpiderID: "facebook_profile_by-profiles-url", Params: []ToolParamSpec{
		{Name: "url", Type: "string", Required: true, Default: "https://www.facebook.com/MayeMusk", Min: "", Max: ""},
		{Name: "file_name", Type: "string", Required: false, Default: "{{TasksID}}", Min: "", Max: ""},
	}},
	{Name: "github_repository", Group: "GitHub", Method: "Repository", ServiceKind: "scraper_builder", Engine: "", SpiderName: "github.com", SpiderID: "github_repository_by-repo-url", Params: []ToolParamSpec{
		{Name: "spider_id", Type: "string", Required: true, Default: "github_repository_by-repo-url", Min: "", Max: ""},
		{Name: "repo_url", Type: "string", Required: false, Default: "https://github.com/TheAlgorithms/Python", Min: "", Max: ""},
		{Name: "search_url", Type: "string", Required: false, Default: "https://github.com/search?q=ML&type=repositories", Min: "", Max: ""},
		{Name: "url", Type: "string", Required: false, Default: "https://github.com/TheAlgorithms/Python/blob/master/divide_and_conquer/power.py", Min: "", Max: ""},
		{Name: "page_turning", Type: "string", Required: false, Default: "1", Min: "", Max: ""},
		{Name: "max_num", Type: "string", Required: false, Default: "15", Min: "", Max: ""},
		{Name: "file_name", Type: "string", Required: false, Default: "{{TasksID}}", Min: "", Max: ""},
	}},
	{Name: "glassdoor_company", Group: "Glassdoor", Method: "Company", ServiceKind: "scraper_builder", Engine: "", SpiderName: "glassdoor.com", SpiderID: "glassdoor_company_by-url", Params: []ToolParamSpec{
		{Name: "spider_id", Type: "string", Required: true, Default: "glassdoor_company_by-url", Min: "", Max: ""},
		{Name: "url", Type: "string", Required: false, Default: "https://www.glassdoor.co.uk/Overview/Working-at-Apple-EI_IE1138.11,16.htm", Min: "", Max: ""},
		{Name: "location", Type: "string", Required: false, Default: "United States", Min: "", Max: ""},
		{Name: "company_name", Type: "string", Required: false, Default: "Tesla", Min: "", Max: ""},
		{Name: "industries", Type: "string", Required: false, Default: "Information Technology", Min: "", Max: ""},
		{Name: "Job title", Type: "string", Required: false, Default: "Data", Min: "", Max: ""},
		{Name: "search_url", Type: "string", Required: false, Default: "https://www.glassdoor.com/Search/results.htm?keyword=Apple", Min: "", Max: ""},
		{Name: "max_search_results", Type: "string", Required: false, Default: "5", Min: "", Max: ""},
		{Name: "file_name", Type: "string", Required: false, Default: "{{TasksID}}", Min: "", Max: ""},
	}},
	{Name: "glassdoor_joblistings", Group: "Glassdoor", Method: "Joblistings", ServiceKind: "scraper_builder", Engine: "", SpiderName: "glassdoor.com", SpiderID: "glassdoor_joblistings_by-url", Params: []ToolParamSpec{
		{Name: "spider_id", Type: "string", Required: true, Default: "glassdoor_joblistings_by-url", Min: "", Max: ""},
		{Name: "url", Type: "string", Required: false, Default: "https://www.glassdoor.com/Job/new-york-data-analyst-jobs-SRCH_IL.0,8_IC1132348_KO9,21.htm", Min: "", Max: ""},
		{Name: "keyword", Type: "string", Required: false, Default: "data analyst", Min: "", Max: ""},
		{Name: "location", Type: "string", Required: false, Default: "New York", Min: "", Max: ""},
		{Name: "country", Type: "string", Required: false, Default: "US", Min: "", Max: ""},
		{Name: "file_name", Type: "string", Required: false, Default: "{{TasksID}}", Min: "", Max: ""},
	}},
	{Name: "google_ai_mode", Group: "Google", Method: "AIMode", ServiceKind: "serp", Engine: "google_ai_mode", SpiderName: "", SpiderID: "", Params: []ToolParamSpec{
		{Name: "q", Type: "string", Required: true, Default: "pizza", Min: "", Max: ""},
		{Name: "json", Type: "string", Required: true, Default: "1", Min: "", Max: ""},
		{Name: "location", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "uule", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "no_cache", Type: "string", Required: false, Default: "false", Min: "", Max: ""},
		{Name: "gl", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "hl", Type: "string", Required: false, Default: "", Min: "", Max: ""},
	}},
	{Name: "google_finance", Group: "Google", Method: "Finance", ServiceKind: "serp", Engine: "google_finance", SpiderName: "", SpiderID: "", Params: []ToolParamSpec{
		{Name: "q", Type: "string", Required: true, Default: "", Min: "", Max: ""},
		{Name: "json", Type: "string", Required: true, Default: "1", Min: "", Max: ""},
		{Name: "hl", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "window", Type: "string", Required: false, Default: "1D", Min: "", Max: ""},
		{Name: "no_cache", Type: "string", Required: false, Default: "false", Min: "", Max: ""},
	}},
	{Name: "google_flights", Group: "Google", Method: "Flights", ServiceKind: "serp", Engine: "google_flights", SpiderName: "", SpiderID: "", Params: []ToolParamSpec{
		{Name: "departure_id", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "arrival_id", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "json", Type: "string", Required: true, Default: "1", Min: "", Max: ""},
		{Name: "gl", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "hl", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "currency", Type: "string", Required: false, Default: "USD", Min: "", Max: ""},
		{Name: "type", Type: "string", Required: false, Default: "1", Min: "", Max: ""},
		{Name: "outbound_date", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "return_date", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "travel_class", Type: "string", Required: false, Default: "1", Min: "", Max: ""},
		{Name: "multi_city_json", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "show_hidden", Type: "string", Required: false, Default: "false", Min: "", Max: ""},
		{Name: "exclude_basic", Type: "string", Required: false, Default: "false", Min: "", Max: ""},
		{Name: "deep_search", Type: "string", Required: false, Default: "false", Min: "", Max: ""},
		{Name: "adults", Type: "string", Required: false, Default: "1", Min: "", Max: ""},
		{Name: "children", Type: "string", Required: false, Default: "0", Min: "", Max: ""},
		{Name: "infants_in_seat", Type: "string", Required: false, Default: "0", Min: "", Max: ""},
		{Name: "infants_on_lap", Type: "string", Required: false, Default: "0", Min: "", Max: ""},
		{Name: "sort_by", Type: "string", Required: false, Default: "1", Min: "", Max: ""},
		{Name: "stops", Type: "string", Required: false, Default: "0", Min: "", Max: ""},
		{Name: "exclude_airlines", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "include_airlines", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "bags", Type: "string", Required: false, Default: "0", Min: "", Max: ""},
		{Name: "max_price", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "outbound_times", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "return_times", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "emissions", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "layover_duration", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "exclude_conns", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "max_duration", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "departure_token", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "no_cache", Type: "string", Required: false, Default: "false", Min: "", Max: ""},
	}},
	{Name: "google_hotels", Group: "Google", Method: "Hotels", ServiceKind: "serp", Engine: "google_hotels", SpiderName: "", SpiderID: "", Params: []ToolParamSpec{
		{Name: "q", Type: "string", Required: true, Default: "pizza", Min: "", Max: ""},
		{Name: "json", Type: "string", Required: true, Default: "1", Min: "", Max: ""},
		{Name: "hl", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "gl", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "currency", Type: "string", Required: false, Default: "USD", Min: "", Max: ""},
		{Name: "check_in_date", Type: "string", Required: true, Default: "", Min: "", Max: ""},
		{Name: "check_out_date", Type: "string", Required: true, Default: "", Min: "", Max: ""},
		{Name: "adults", Type: "string", Required: false, Default: "2", Min: "", Max: ""},
		{Name: "children", Type: "string", Required: false, Default: "0", Min: "", Max: ""},
		{Name: "children_ages", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "sort_by", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "min_price", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "max_price", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "property_types", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "amenities", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "rating", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "brands", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "hotel_class", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "free_cancellation", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "special_offers", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "eco_certified", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "vacation_rentals", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "bedrooms", Type: "string", Required: false, Default: "0", Min: "", Max: ""},
		{Name: "bathrooms", Type: "string", Required: false, Default: "0", Min: "", Max: ""},
		{Name: "next_page_token", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "no_cache", Type: "string", Required: false, Default: "false", Min: "", Max: ""},
		{Name: "property_token", Type: "string", Required: false, Default: "", Min: "", Max: ""},
	}},
	{Name: "google_images", Group: "Google", Method: "Images", ServiceKind: "serp", Engine: "google_images", SpiderName: "", SpiderID: "", Params: []ToolParamSpec{
		{Name: "q", Type: "string", Required: true, Default: "pizza", Min: "", Max: ""},
		{Name: "json", Type: "string", Required: true, Default: "1", Min: "", Max: ""},
		{Name: "google_domain", Type: "string", Required: false, Default: "google.com", Min: "", Max: ""},
		{Name: "gl", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "hl", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "cr", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "lr", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "location", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "uule", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "lat", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "lon", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "radius", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "start", Type: "string", Required: false, Default: "0", Min: "", Max: ""},
		{Name: "tbm", Type: "string", Required: false, Default: "isch", Min: "", Max: ""},
		{Name: "ludocid", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "lsig", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "kgmid", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "si", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "ibp", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "uds", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "tbs", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "safe", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "nfpr", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "filter", Type: "string", Required: false, Default: "1", Min: "", Max: ""},
		{Name: "device", Type: "string", Required: false, Default: "desktop", Min: "", Max: ""},
		{Name: "render_js", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "no_cache", Type: "string", Required: false, Default: "false", Min: "", Max: ""},
		{Name: "ai_overview", Type: "string", Required: false, Default: "", Min: "", Max: ""},
	}},
	{Name: "google_jobs", Group: "Google", Method: "Jobs", ServiceKind: "serp", Engine: "google_jobs", SpiderName: "", SpiderID: "", Params: []ToolParamSpec{
		{Name: "q", Type: "string", Required: true, Default: "", Min: "", Max: ""},
		{Name: "json", Type: "string", Required: true, Default: "1", Min: "", Max: ""},
		{Name: "google_domain", Type: "string", Required: false, Default: "google.com", Min: "", Max: ""},
		{Name: "gl", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "hl", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "location", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "uule", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "next_page_token", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "chips", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "lrad", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "ltype", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "uds", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "no_cache", Type: "string", Required: false, Default: "false", Min: "", Max: ""},
	}},
	{Name: "google_lens", Group: "Google", Method: "Lens", ServiceKind: "serp", Engine: "google_lens", SpiderName: "", SpiderID: "", Params: []ToolParamSpec{
		{Name: "url", Type: "string", Required: true, Default: "", Min: "", Max: ""},
		{Name: "json", Type: "string", Required: true, Default: "1", Min: "", Max: ""},
		{Name: "hl", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "country", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "type", Type: "string", Required: false, Default: "all", Min: "", Max: ""},
		{Name: "q", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "safe", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "no_cache", Type: "string", Required: false, Default: "false", Min: "", Max: ""},
	}},
	{Name: "google_local", Group: "Google", Method: "Local", ServiceKind: "serp", Engine: "google_local", SpiderName: "", SpiderID: "", Params: []ToolParamSpec{
		{Name: "q", Type: "string", Required: true, Default: "", Min: "", Max: ""},
		{Name: "json", Type: "string", Required: true, Default: "1", Min: "", Max: ""},
		{Name: "google_domain", Type: "string", Required: false, Default: "google.com", Min: "", Max: ""},
		{Name: "gl", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "hl", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "location", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "uule", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "start", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "ludocid", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "tbs", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "no_cache", Type: "string", Required: false, Default: "false", Min: "", Max: ""},
	}},
	{Name: "google_map_comment", Group: "Google", Method: "MapComment", ServiceKind: "scraper_builder", Engine: "", SpiderName: "google.com", SpiderID: "google_comment_by-url", Params: []ToolParamSpec{
		{Name: "url", Type: "string", Required: true, Default: "https://www.google.com/maps/place/Waterfront+Botanical+Gardens/@38.2630366,-85.7288454,15z/data=!4m8!3m7!1s0x8869731e16a7bdbd:0x2f5d238fefed7ca1!8m2!3d38.2632837!4d-85.7239738!9m1!1b1!16s%2Fg%2F11c709xzzx?hl=en&entry=ttu", Min: "", Max: ""},
		{Name: "days_limit", Type: "string", Required: false, Default: "20", Min: "", Max: ""},
		{Name: "file_name", Type: "string", Required: false, Default: "{{TasksID}}", Min: "", Max: ""},
	}},
	{Name: "google_map_details", Group: "Google", Method: "MapDetails", ServiceKind: "scraper_builder", Engine: "", SpiderName: "google.com", SpiderID: "google_map-details_by-url", Params: []ToolParamSpec{
		{Name: "spider_id", Type: "string", Required: true, Default: "google_map-details_by-url", Min: "", Max: ""},
		{Name: "url", Type: "string", Required: false, Default: "https://www.google.com/maps/place/Pizza+Inn+Magdeburg/data=!4m7!3m6!1s0x47a5f50c083530a3:0xfdba8746b538141!8m2!3d52.1263086!4d11.6094743!16s%2Fg%2F11kqmtk3dt!19sChIJozA1CAz1pUcRQYFTa3So2w8?authuser=0&hl=en&rclk=1", Min: "", Max: ""},
		{Name: "CID", Type: "string", Required: false, Default: "2476046430038551731", Min: "", Max: ""},
		{Name: "keyword", Type: "string", Required: false, Default: "pizza", Min: "", Max: ""},
		{Name: "country", Type: "string", Required: false, Default: "United States", Min: "", Max: ""},
		{Name: "lat", Type: "string", Required: false, Default: "38", Min: "", Max: ""},
		{Name: "long", Type: "string", Required: false, Default: "77", Min: "", Max: ""},
		{Name: "zoom_level", Type: "string", Required: false, Default: "20", Min: "", Max: ""},
		{Name: "place_id", Type: "string", Required: false, Default: "ChIJ3S-JXmauEmsRUcIaWtf4MzE", Min: "", Max: ""},
		{Name: "file_name", Type: "string", Required: false, Default: "{{TasksID}}", Min: "", Max: ""},
	}},
	{Name: "google_maps", Group: "Google", Method: "Maps", ServiceKind: "serp", Engine: "google_maps", SpiderName: "", SpiderID: "", Params: []ToolParamSpec{
		{Name: "q", Type: "string", Required: true, Default: "", Min: "", Max: ""},
		{Name: "json", Type: "string", Required: true, Default: "1", Min: "", Max: ""},
		{Name: "ll", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "location", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "lat", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "lon", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "z", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "m", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "nearby", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "google_domain", Type: "string", Required: false, Default: "google.com", Min: "", Max: ""},
		{Name: "hl", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "gl", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "start", Type: "string", Required: false, Default: "0", Min: "", Max: ""},
		{Name: "type", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "data", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "place_id", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "data_cid", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "no_cache", Type: "string", Required: false, Default: "false", Min: "", Max: ""},
	}},
	{Name: "google_news", Group: "Google", Method: "News", ServiceKind: "serp", Engine: "google_news", SpiderName: "", SpiderID: "", Params: []ToolParamSpec{
		{Name: "q", Type: "string", Required: true, Default: "pizza", Min: "", Max: ""},
		{Name: "json", Type: "string", Required: true, Default: "1", Min: "", Max: ""},
		{Name: "no_cache", Type: "string", Required: false, Default: "false", Min: "", Max: ""},
		{Name: "gl", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "hl", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "topic_token", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "kgmid", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "publication_token", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "section_token", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "story_token", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "so", Type: "string", Required: false, Default: "0", Min: "", Max: ""},
	}},
	{Name: "google_patents", Group: "Google", Method: "Patents", ServiceKind: "serp", Engine: "google_patents", SpiderName: "", SpiderID: "", Params: []ToolParamSpec{
		{Name: "q", Type: "string", Required: true, Default: "", Min: "", Max: ""},
		{Name: "json", Type: "string", Required: true, Default: "1", Min: "", Max: ""},
		{Name: "page", Type: "string", Required: false, Default: "1", Min: "", Max: ""},
		{Name: "num", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "sort", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "clustered", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "dups", Type: "string", Required: false, Default: "family", Min: "", Max: ""},
		{Name: "patents", Type: "string", Required: false, Default: "true", Min: "", Max: ""},
		{Name: "scholar", Type: "string", Required: false, Default: "false", Min: "", Max: ""},
		{Name: "before", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "after", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "inventor", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "assignee", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "country", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "language", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "status", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "type", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "litigation", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "no_cache", Type: "string", Required: false, Default: "false", Min: "", Max: ""},
	}},
	{Name: "google_play", Group: "Google", Method: "Play", ServiceKind: "serp", Engine: "google_play", SpiderName: "", SpiderID: "", Params: []ToolParamSpec{
		{Name: "q", Type: "string", Required: true, Default: "", Min: "", Max: ""},
		{Name: "json", Type: "string", Required: true, Default: "1", Min: "", Max: ""},
		{Name: "hl", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "gl", Type: "string", Required: false, Default: "us", Min: "", Max: ""},
		{Name: "apps_category", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "next_page_token", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "section_page_token", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "chart", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "see_more_token", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "store_device", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "age", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "no_cache", Type: "string", Required: false, Default: "false", Min: "", Max: ""},
	}},
	{Name: "google_play_store_information", Group: "Google", Method: "PlayStoreInformation", ServiceKind: "scraper_builder", Engine: "", SpiderName: "play.google.com", SpiderID: "google-play-store_information_by-url", Params: []ToolParamSpec{
		{Name: "app_url", Type: "string", Required: true, Default: "https://play.google.com/store/apps/details?id=com.linkedin.android", Min: "", Max: ""},
		{Name: "file_name", Type: "string", Required: false, Default: "{{TasksID}}", Min: "", Max: ""},
	}},
	{Name: "google_play_store_reviews", Group: "Google", Method: "PlayStoreReviews", ServiceKind: "scraper_builder", Engine: "", SpiderName: "play.google.com", SpiderID: "google-play-store_reviews_by-url", Params: []ToolParamSpec{
		{Name: "app_url", Type: "string", Required: true, Default: "https://play.google.com/store/apps/details?id=com.linkedin.android", Min: "", Max: ""},
		{Name: "file_name", Type: "string", Required: false, Default: "{{TasksID}}", Min: "", Max: ""},
	}},
	{Name: "google_scholar", Group: "Google", Method: "Scholar", ServiceKind: "serp", Engine: "google_scholar", SpiderName: "", SpiderID: "", Params: []ToolParamSpec{
		{Name: "q", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "json", Type: "string", Required: true, Default: "1", Min: "", Max: ""},
		{Name: "hl", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "lr", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "start", Type: "string", Required: false, Default: "0", Min: "", Max: ""},
		{Name: "num", Type: "string", Required: false, Default: "10", Min: "", Max: ""},
		{Name: "cites", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "as_ylo", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "as_yhi", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "scisbd", Type: "string", Required: false, Default: "0", Min: "", Max: ""},
		{Name: "cluster", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "as_sdt", Type: "string", Required: false, Default: "0", Min: "", Max: ""},
		{Name: "safe", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "filter", Type: "string", Required: false, Default: "1", Min: "", Max: ""},
		{Name: "as_vis", Type: "string", Required: false, Default: "0", Min: "", Max: ""},
		{Name: "as_rr", Type: "string", Required: false, Default: "0", Min: "", Max: ""},
		{Name: "no_cache", Type: "string", Required: false, Default: "false", Min: "", Max: ""},
	}},
	{Name: "google_search", Group: "Google", Method: "Search", ServiceKind: "serp", Engine: "google", SpiderName: "", SpiderID: "", Params: []ToolParamSpec{
		{Name: "q", Type: "string", Required: true, Default: "pizza", Min: "", Max: ""},
		{Name: "json", Type: "string", Required: true, Default: "1", Min: "", Max: ""},
		{Name: "google_domain", Type: "string", Required: false, Default: "google.com", Min: "", Max: ""},
		{Name: "gl", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "hl", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "cr", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "lr", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "location", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "uule", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "start", Type: "string", Required: false, Default: "0", Min: "", Max: ""},
		{Name: "tbs", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "safe", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "nfpr", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "filter", Type: "string", Required: false, Default: "1", Min: "", Max: ""},
		{Name: "device", Type: "string", Required: false, Default: "desktop", Min: "", Max: ""},
		{Name: "render_js", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "no_cache", Type: "string", Required: false, Default: "false", Min: "", Max: ""},
		{Name: "ai_overview", Type: "string", Required: false, Default: "", Min: "", Max: ""},
	}},
	{Name: "google_shopping", Group: "Google", Method: "Shopping", ServiceKind: "serp", Engine: "google_shopping", SpiderName: "", SpiderID: "", Params: []ToolParamSpec{
		{Name: "q", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "json", Type: "string", Required: true, Default: "1", Min: "", Max: ""},
		{Name: "google_domain", Type: "string", Required: false, Default: "google.com", Min: "", Max: ""},
		{Name: "gl", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "hl", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "location", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "uule", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "start", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "shoprs", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "min_price", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "max_price", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "sort_by", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "free_shipping", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "on_sale", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "small_business", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "no_cache", Type: "string", Required: false, Default: "false", Min: "", Max: ""},
	}},
	{Name: "google_shopping_info", Group: "Google", Method: "ShoppingInfo", ServiceKind: "scraper_builder", Engine: "", SpiderName: "google.com", SpiderID: "google_shopping_by-keywords", Params: []ToolParamSpec{
		{Name: "keyword", Type: "string", Required: true, Default: "iphone", Min: "", Max: ""},
		{Name: "file_name", Type: "string", Required: false, Default: "{{TasksID}}", Min: "", Max: ""},
	}},
	{Name: "google_trends", Group: "Google", Method: "Trends", ServiceKind: "serp", Engine: "google_trends", SpiderName: "", SpiderID: "", Params: []ToolParamSpec{
		{Name: "q", Type: "string", Required: true, Default: "pizza", Min: "", Max: ""},
		{Name: "json", Type: "string", Required: true, Default: "1", Min: "", Max: ""},
		{Name: "hl", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "geo", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "region", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "data_type", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "tz", Type: "string", Required: false, Default: "420", Min: "", Max: ""},
		{Name: "cat", Type: "string", Required: false, Default: "0", Min: "", Max: ""},
		{Name: "gprop", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "date", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "csv", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "include_low_search_volume", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "no_cache", Type: "string", Required: false, Default: "false", Min: "", Max: ""},
	}},
	{Name: "google_videos", Group: "Google", Method: "Videos", ServiceKind: "serp", Engine: "google_videos", SpiderName: "", SpiderID: "", Params: []ToolParamSpec{
		{Name: "q", Type: "string", Required: true, Default: "", Min: "", Max: ""},
		{Name: "json", Type: "string", Required: true, Default: "1", Min: "", Max: ""},
		{Name: "google_domain", Type: "string", Required: false, Default: "google.com", Min: "", Max: ""},
		{Name: "gl", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "hl", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "location", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "uule", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "start", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "tbs", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "no_cache", Type: "string", Required: false, Default: "false", Min: "", Max: ""},
		{Name: "lr", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "safe", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "nfpr", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "filter", Type: "string", Required: false, Default: "1", Min: "", Max: ""},
	}},
	{Name: "indeed_companies_info", Group: "Indeed", Method: "CompaniesInfo", ServiceKind: "scraper_builder", Engine: "", SpiderName: "indeed.com", SpiderID: "indeed_companies-info_by-company-list-url", Params: []ToolParamSpec{
		{Name: "spider_id", Type: "string", Required: true, Default: "indeed_companies-info_by-company-list-url", Min: "", Max: ""},
		{Name: "company_list_url", Type: "string", Required: false, Default: "https://www.indeed.com/companies/browse-companies", Min: "", Max: ""},
		{Name: "keyword", Type: "string", Required: false, Default: "openai", Min: "", Max: ""},
		{Name: "industry", Type: "string", Required: false, Default: "Accounting & Tax", Min: "", Max: ""},
		{Name: "state", Type: "string", Required: false, Default: "Alabama - 60 companies", Min: "", Max: ""},
		{Name: "company_url", Type: "string", Required: false, Default: "https://www.indeed.com/cmp/Allstate-Insurance", Min: "", Max: ""},
		{Name: "file_name", Type: "string", Required: false, Default: "{{TasksID}}", Min: "", Max: ""},
	}},
	{Name: "indeed_job_listings", Group: "Indeed", Method: "JobListings", ServiceKind: "scraper_builder", Engine: "", SpiderName: "indeed.com", SpiderID: "indeed_job-listings_by-job-url", Params: []ToolParamSpec{
		{Name: "job_url", Type: "string", Required: true, Default: "https://fr.indeed.com/viewjob?jk=55b3e5dfa0c2ff66", Min: "", Max: ""},
		{Name: "file_name", Type: "string", Required: false, Default: "{{TasksID}}", Min: "", Max: ""},
	}},
	{Name: "instagram_comment", Group: "Instagram", Method: "Comment", ServiceKind: "scraper_builder", Engine: "", SpiderName: "instagram.com", SpiderID: "ins_comment_by-posturl", Params: []ToolParamSpec{
		{Name: "posturl", Type: "string", Required: true, Default: "https://www.instagram.com/cats_of_instagram/reel/C4GLo_eLO2e/", Min: "", Max: ""},
		{Name: "file_name", Type: "string", Required: false, Default: "{{TasksID}}", Min: "", Max: ""},
	}},
	{Name: "instagram_profiles", Group: "Instagram", Method: "Profiles", ServiceKind: "scraper_builder", Engine: "", SpiderName: "instagram.com", SpiderID: "ins_profiles_by-username", Params: []ToolParamSpec{
		{Name: "spider_id", Type: "string", Required: true, Default: "ins_profiles_by-username", Min: "", Max: ""},
		{Name: "username", Type: "string", Required: false, Default: "zoobarcelona", Min: "", Max: ""},
		{Name: "profileurl", Type: "string", Required: false, Default: "https://www.instagram.com/cats_of_world_/", Min: "", Max: ""},
		{Name: "file_name", Type: "string", Required: false, Default: "{{TasksID}}", Min: "", Max: ""},
	}},
	{Name: "instagram_reel", Group: "Instagram", Method: "Reel", ServiceKind: "scraper_builder", Engine: "", SpiderName: "instagram.com", SpiderID: "ins_reel_by-url", Params: []ToolParamSpec{
		{Name: "spider_id", Type: "string", Required: true, Default: "ins_reel_by-url", Min: "", Max: ""},
		{Name: "url", Type: "string", Required: true, Default: "https://www.instagram.com/reel/C5Rdyj_q7YN/", Min: "", Max: ""},
		{Name: "num_of_posts", Type: "string", Required: false, Default: "10", Min: "", Max: ""},
		{Name: "posts_to_not_include", Type: "string", Required: false, Default: "DP861NijuwE", Min: "", Max: ""},
		{Name: "start_date", Type: "string", Required: false, Default: "01-28-2025", Min: "", Max: ""},
		{Name: "end_date", Type: "string", Required: false, Default: "01-28-2026", Min: "", Max: ""},
		{Name: "file_name", Type: "string", Required: false, Default: "{{TasksID}}", Min: "", Max: ""},
	}},
	{Name: "linkedin_company_information", Group: "LinkedIn", Method: "CompanyInformation", ServiceKind: "scraper_builder", Engine: "", SpiderName: "linkedin.com", SpiderID: "linkedin_company_information_by-url", Params: []ToolParamSpec{
		{Name: "url", Type: "string", Required: true, Default: "https://www.linkedin.com/company/dynamo-software", Min: "", Max: ""},
		{Name: "file_name", Type: "string", Required: false, Default: "{{TasksID}}", Min: "", Max: ""},
	}},
	{Name: "linkedin_job_listings_information", Group: "LinkedIn", Method: "JobListingsInformation", ServiceKind: "scraper_builder", Engine: "", SpiderName: "linkedin.com", SpiderID: "linkedin_job_listings_information_by-job-listing-url", Params: []ToolParamSpec{
		{Name: "spider_id", Type: "string", Required: true, Default: "linkedin_job_listings_information_by-job-listing-url", Min: "", Max: ""},
		{Name: "job_listing_url", Type: "string", Required: false, Default: "https://www.linkedin.com/jobs/reddit-inc.-jobs-worldwide?f_C=150573", Min: "", Max: ""},
		{Name: "job_url", Type: "string", Required: false, Default: "https://www.linkedin.com/jobs/view/senior-client-partner-large-customer-sales-gaming-at-reddit-inc-4303761033?position=10&pageNum=0&refId=kHRQtl6Ws14VG9y3UloI4w%3D%3D&trackingId=%2Fb2esqHHEjp1FoEkC8PfuQ%3D%3D", Min: "", Max: ""},
		{Name: "location", Type: "string", Required: false, Default: "New York", Min: "", Max: ""},
		{Name: "keyword", Type: "string", Required: false, Default: "product manager", Min: "", Max: ""},
		{Name: "file_name", Type: "string", Required: false, Default: "{{TasksID}}", Min: "", Max: ""},
	}},
	{Name: "reddit_comment", Group: "Reddit", Method: "Comment", ServiceKind: "scraper_builder", Engine: "", SpiderName: "reddit.com", SpiderID: "reddit_comment_by-url", Params: []ToolParamSpec{
		{Name: "url", Type: "string", Required: true, Default: "https://www.reddit.com/r/datascience/comments/1cmnf0m/comment/l32204i/?utm_source=share&utm_medium=web3x&utm_name=web3xcss&utm_term=1&utm_content=share_button", Min: "", Max: ""},
		{Name: "days_back", Type: "string", Required: false, Default: "10", Min: "", Max: ""},
		{Name: "comment_limit", Type: "string", Required: false, Default: "5", Min: "", Max: ""},
		{Name: "file_name", Type: "string", Required: false, Default: "{{TasksID}}", Min: "", Max: ""},
	}},
	{Name: "reddit_posts", Group: "Reddit", Method: "Posts", ServiceKind: "scraper_builder", Engine: "", SpiderName: "reddit.com", SpiderID: "reddit_posts_by-url", Params: []ToolParamSpec{
		{Name: "spider_id", Type: "string", Required: true, Default: "reddit_posts_by-url", Min: "", Max: ""},
		{Name: "url", Type: "string", Required: false, Default: "https://www.reddit.com/r/battlefield2042/comments/1cmqs1d/official_update_on_the_next_battlefield_game/", Min: "", Max: ""},
		{Name: "keyword", Type: "string", Required: false, Default: "datascience", Min: "", Max: ""},
		{Name: "num_of_posts", Type: "string", Required: false, Default: "10", Min: "", Max: ""},
		{Name: "sort_by", Type: "string", Required: false, Default: "Rising", Min: "", Max: ""},
		{Name: "sort_by_time", Type: "string", Required: false, Default: "Now", Min: "", Max: ""},
		{Name: "file_name", Type: "string", Required: false, Default: "{{TasksID}}", Min: "", Max: ""},
	}},
	{Name: "request_web_unlocker", Group: "WebUnlocker", Method: "Request", ServiceKind: "web_unlocker", Engine: "", SpiderName: "", SpiderID: "", Params: []ToolParamSpec{
		{Name: "url", Type: "string", Required: true, Default: "", Min: "", Max: ""},
		{Name: "type", Type: "string", Required: false, Default: "html", Min: "", Max: ""},
		{Name: "js_render", Type: "string", Required: false, Default: "True", Min: "", Max: ""},
		{Name: "block_resources", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "clean_content", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "country", Type: "string", Required: false, Default: "us", Min: "", Max: ""},
		{Name: "headers", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "cookies", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "wait", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "wait_for", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "follow_redirect", Type: "string", Required: false, Default: "True", Min: "", Max: ""},
	}},
	{Name: "tiktok_comment", Group: "TikTok", Method: "Comment", ServiceKind: "scraper_builder", Engine: "", SpiderName: "tiktok.com", SpiderID: "tiktok_comment_by-url", Params: []ToolParamSpec{
		{Name: "url", Type: "string", Required: true, Default: "https://www.tiktok.com/@heymrcat/video/7216019547806092550", Min: "", Max: ""},
		{Name: "page_turning", Type: "string", Required: true, Default: "1", Min: "", Max: ""},
		{Name: "file_name", Type: "string", Required: false, Default: "{{TasksID}}", Min: "", Max: ""},
	}},
	{Name: "tiktok_posts", Group: "TikTok", Method: "Posts", ServiceKind: "scraper_builder", Engine: "", SpiderName: "tiktok.com", SpiderID: "tiktok_posts_by-listurl", Params: []ToolParamSpec{
		{Name: "url", Type: "string", Required: true, Default: "https://www.tiktok.com/discover/dog", Min: "", Max: ""},
		{Name: "num_of_posts", Type: "string", Required: false, Default: "5", Min: "", Max: ""},
		{Name: "file_name", Type: "string", Required: false, Default: "{{TasksID}}", Min: "", Max: ""},
	}},
	{Name: "tiktok_profiles", Group: "TikTok", Method: "Profiles", ServiceKind: "scraper_builder", Engine: "", SpiderName: "tiktok.com", SpiderID: "tiktok_profiles_by-url", Params: []ToolParamSpec{
		{Name: "spider_id", Type: "string", Required: true, Default: "tiktok_profiles_by-url", Min: "", Max: ""},
		{Name: "url", Type: "string", Required: false, Default: "https://www.tiktok.com/@fofimdmell", Min: "", Max: ""},
		{Name: "search_url", Type: "string", Required: false, Default: "https://www.tiktok.com/explore?lang=en", Min: "", Max: ""},
		{Name: "country", Type: "string", Required: false, Default: "us", Min: "", Max: ""},
		{Name: "page_turning", Type: "string", Required: false, Default: "1", Min: "", Max: ""},
		{Name: "file_name", Type: "string", Required: false, Default: "{{TasksID}}", Min: "", Max: ""},
	}},
	{Name: "tiktok_shop", Group: "TikTok", Method: "Shop", ServiceKind: "scraper_builder", Engine: "", SpiderName: "tiktok.com", SpiderID: "tiktok_shop_by-url", Params: []ToolParamSpec{
		{Name: "url", Type: "string", Required: true, Default: "https://www.tiktok.com/shop/pdp/long-sleeve-crew-neck-tee-3-pack-by-galaxy-by-harvic-cotton-blend/1729461570693075200", Min: "", Max: ""},
		{Name: "file_name", Type: "string", Required: false, Default: "{{TasksID}}", Min: "", Max: ""},
	}},
	{Name: "twitter_post", Group: "Twitter", Method: "Post", ServiceKind: "scraper_builder", Engine: "", SpiderName: "x.com", SpiderID: "twitter_post_by-profileurl", Params: []ToolParamSpec{
		{Name: "url", Type: "string", Required: true, Default: "https://x.com/elonmusk", Min: "", Max: ""},
		{Name: "file_name", Type: "string", Required: false, Default: "{{TasksID}}", Min: "", Max: ""},
	}},
	{Name: "twitter_profile", Group: "Twitter", Method: "Profile", ServiceKind: "scraper_builder", Engine: "", SpiderName: "x.com", SpiderID: "twitter_profile_by-profileurl", Params: []ToolParamSpec{
		{Name: "spider_id", Type: "string", Required: true, Default: "twitter_profile_by-profileurl", Min: "", Max: ""},
		{Name: "url", Type: "string", Required: false, Default: "https://x.com/elonmusk", Min: "", Max: ""},
		{Name: "user_name", Type: "string", Required: false, Default: "elonmusk", Min: "", Max: ""},
		{Name: "file_name", Type: "string", Required: false, Default: "{{TasksID}}", Min: "", Max: ""},
	}},
	{Name: "walmart_product", Group: "Walmart", Method: "Product", ServiceKind: "scraper_builder", Engine: "", SpiderName: "walmart.com", SpiderID: "walmart_product_by-url", Params: []ToolParamSpec{
		{Name: "spider_id", Type: "string", Required: true, Default: "walmart_product_by-url", Min: "", Max: ""},
		{Name: "url", Type: "string", Required: false, Default: "https://www.walmart.com/ip/HI-CHEW-Stand-Up-Pouch-Getaway-Mix-11-65oz/12284762931?athAsset=eyJhdGhjcGlkIjoiMTIyODQ3NjI5MzEiLCJhdGhzdGlkIjoiQ1MwNTV+Q1MwMDR+Q1MwOTgiLCJhdGhlZSI6eyJhIjoyNy44NCwiYiI6Mjk1MS40MSwidyI6MC4wMDk0MjcxMjc3OTA0NzcxMjMsImwiOjAuNX0sImF0aHBvc2IiOiI4IiwiYXRoYW5jaWQiOiIxMDE2NDUwNzU1IiwiYXRocmsiOjAuMH0%3D&athena=true&adsRedirect=true", Min: "", Max: ""},
		{Name: "category_url", Type: "string", Required: false, Default: "https://www.walmart.com/shop/deals/food/", Min: "", Max: ""},
		{Name: "sku", Type: "string", Required: false, Default: "439179861", Min: "", Max: ""},
		{Name: "keyword", Type: "string", Required: false, Default: "leggins", Min: "", Max: ""},
		{Name: "domain", Type: "string", Required: false, Default: "https://www.walmart.com/", Min: "", Max: ""},
		{Name: "all_variations", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "page_turning", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "file_name", Type: "string", Required: false, Default: "{{TasksID}}", Min: "", Max: ""},
	}},
	{Name: "yandex_search", Group: "Yandex", Method: "Search", ServiceKind: "serp", Engine: "yandex", SpiderName: "", SpiderID: "", Params: []ToolParamSpec{
		{Name: "text", Type: "string", Required: true, Default: "", Min: "", Max: ""},
		{Name: "json", Type: "string", Required: true, Default: "1", Min: "", Max: ""},
		{Name: "yandex_domain", Type: "string", Required: false, Default: "yandex.com", Min: "", Max: ""},
		{Name: "lang", Type: "string", Required: false, Default: "en", Min: "", Max: ""},
		{Name: "lr", Type: "string", Required: false, Default: "", Min: "", Max: ""},
		{Name: "p", Type: "string", Required: false, Default: "0", Min: "", Max: ""},
		{Name: "family_mode", Type: "string", Required: false, Default: "1", Min: "", Max: ""},
		{Name: "fix_typo", Type: "string", Required: false, Default: "true", Min: "", Max: ""},
		{Name: "groups_on_page", Type: "string", Required: false, Default: "10", Min: "", Max: ""},
		{Name: "no_cache", Type: "string", Required: false, Default: "false", Min: "", Max: ""},
	}},
	{Name: "youtube_audio", Group: "YouTube", Method: "Audio", ServiceKind: "scraper_builder", Engine: "", SpiderName: "youtube.com", SpiderID: "youtube_audio_by-url", Params: []ToolParamSpec{
		{Name: "url", Type: "string", Required: true, Default: "https://www.youtube.com/watch?v=_SdpvpvVrLY", Min: "", Max: ""},
		{Name: "subtitles_language", Type: "string", Required: false, Default: "ab", Min: "", Max: ""},
		{Name: "selected_only", Type: "string", Required: false, Default: "false", Min: "", Max: ""},
		{Name: "kilohertz", Type: "string", Required: false, Default: "<=48000", Min: "", Max: ""},
		{Name: "is_subtitles", Type: "string", Required: false, Default: "false", Min: "", Max: ""},
		{Name: "audio_format", Type: "string", Required: false, Default: "opus", Min: "", Max: ""},
		{Name: "bitrate", Type: "string", Required: false, Default: "<=320", Min: "", Max: ""},
		{Name: "file_name", Type: "string", Required: false, Default: "{{TasksID}}", Min: "", Max: ""},
	}},
	{Name: "youtube_comment", Group: "YouTube", Method: "Comment", ServiceKind: "scraper_builder", Engine: "", SpiderName: "youtube.com", SpiderID: "youtube_comment_by-id", Params: []ToolParamSpec{
		{Name: "video_id", Type: "string", Required: true, Default: "8RePenzQH80", Min: "", Max: ""},
		{Name: "load_replies", Type: "string", Required: true, Default: "10", Min: "", Max: ""},
		{Name: "num_of_comments", Type: "string", Required: true, Default: "Top comments", Min: "", Max: ""},
		{Name: "file_name", Type: "string", Required: false, Default: "{{TasksID}}", Min: "", Max: ""},
	}},
	{Name: "youtube_product", Group: "YouTube", Method: "Product", ServiceKind: "scraper_builder", Engine: "", SpiderName: "youtube.com", SpiderID: "youtube_product_by-id", Params: []ToolParamSpec{
		{Name: "video_id", Type: "string", Required: true, Default: "8RePenzQH80", Min: "", Max: ""},
		{Name: "subtitles_language", Type: "string", Required: false, Default: "ab", Min: "", Max: ""},
		{Name: "subtitles_type", Type: "string", Required: false, Default: "auto_generated", Min: "", Max: ""},
		{Name: "selected_only", Type: "string", Required: false, Default: "false", Min: "", Max: ""},
		{Name: "file_name", Type: "string", Required: false, Default: "{{TasksID}}", Min: "", Max: ""},
	}},
	{Name: "youtube_profiles", Group: "YouTube", Method: "Profiles", ServiceKind: "scraper_builder", Engine: "", SpiderName: "youtube.com", SpiderID: "youtube_profiles_by-keyword", Params: []ToolParamSpec{
		{Name: "spider_id", Type: "string", Required: true, Default: "youtube_profiles_by-keyword", Min: "", Max: ""},
		{Name: "keyword", Type: "string", Required: false, Default: "MrBeast", Min: "", Max: ""},
		{Name: "page_turning", Type: "string", Required: false, Default: "1", Min: "", Max: ""},
		{Name: "url", Type: "string", Required: false, Default: "https://www.youtube.com/@mrbeast", Min: "", Max: ""},
		{Name: "file_name", Type: "string", Required: false, Default: "{{TasksID}}", Min: "", Max: ""},
	}},
	{Name: "youtube_transcript", Group: "YouTube", Method: "Transcript", ServiceKind: "scraper_builder", Engine: "", SpiderName: "youtube.com", SpiderID: "youtube_transcript_by-id", Params: []ToolParamSpec{
		{Name: "video_id", Type: "string", Required: true, Default: "8RePenzQH80", Min: "", Max: ""},
		{Name: "subtitles_language", Type: "string", Required: false, Default: "ab", Min: "", Max: ""},
		{Name: "subtitles_type", Type: "string", Required: false, Default: "auto_generated", Min: "", Max: ""},
		{Name: "selected_only", Type: "string", Required: false, Default: "false", Min: "", Max: ""},
		{Name: "file_name", Type: "string", Required: false, Default: "{{TasksID}}", Min: "", Max: ""},
	}},
	{Name: "youtube_video", Group: "YouTube", Method: "Video", ServiceKind: "scraper_builder", Engine: "", SpiderName: "youtube.com", SpiderID: "youtube_video_by-url", Params: []ToolParamSpec{
		{Name: "url", Type: "string", Required: true, Default: "https://www.youtube.com/watch?v=_SdpvpvVrLY", Min: "", Max: ""},
		{Name: "subtitles_language", Type: "string", Required: false, Default: "ab", Min: "", Max: ""},
		{Name: "selected_only", Type: "string", Required: false, Default: "false", Min: "", Max: ""},
		{Name: "resolution", Type: "string", Required: false, Default: "<=360p", Min: "", Max: ""},
		{Name: "video_codec", Type: "string", Required: false, Default: "vp9", Min: "", Max: ""},
		{Name: "audio_format", Type: "string", Required: false, Default: "opus", Min: "", Max: ""},
		{Name: "bitrate", Type: "string", Required: false, Default: "<=320", Min: "", Max: ""},
		{Name: "file_name", Type: "string", Required: false, Default: "{{TasksID}}", Min: "", Max: ""},
	}},
	{Name: "youtube_video_post", Group: "YouTube", Method: "VideoPost", ServiceKind: "scraper_builder", Engine: "", SpiderName: "youtube.com", SpiderID: "youtube_video-post_by-url", Params: []ToolParamSpec{
		{Name: "spider_id", Type: "string", Required: true, Default: "youtube_video-post_by-url", Min: "", Max: ""},
		{Name: "url", Type: "string", Required: false, Default: "https://www.youtube.com/@stephcurry/videos", Min: "", Max: ""},
		{Name: "order_by", Type: "string", Required: false, Default: "Latest", Min: "", Max: ""},
		{Name: "start_index", Type: "string", Required: false, Default: "1", Min: "", Max: ""},
		{Name: "num_of_posts", Type: "string", Required: false, Default: "5", Min: "", Max: ""},
		{Name: "keyword_search", Type: "string", Required: false, Default: "popular music", Min: "", Max: ""},
		{Name: "features", Type: "string", Required: false, Default: "All", Min: "", Max: ""},
		{Name: "type", Type: "string", Required: false, Default: "Video", Min: "", Max: ""},
		{Name: "duration", Type: "string", Required: false, Default: "Under 3 minutes", Min: "", Max: ""},
		{Name: "upload_date", Type: "string", Required: false, Default: "上一小时", Min: "", Max: ""},
		{Name: "hashtag", Type: "string", Required: false, Default: "shopping", Min: "", Max: ""},
		{Name: "keyword", Type: "string", Required: false, Default: "top videos", Min: "", Max: ""},
		{Name: "all_tabs", Type: "string", Required: false, Default: "true", Min: "", Max: ""},
		{Name: "file_name", Type: "string", Required: false, Default: "{{TasksID}}", Min: "", Max: ""},
	}},
	{Name: "zillow_product", Group: "Zillow", Method: "Product", ServiceKind: "scraper_builder", Engine: "", SpiderName: "zillow.com", SpiderID: "zillow_product_by-filter", Params: []ToolParamSpec{
		{Name: "keywords-location", Type: "string", Required: true, Default: "South Bend", Min: "", Max: ""},
		{Name: "listingCategory", Type: "string", Required: false, Default: "For Rent", Min: "", Max: ""},
		{Name: "HomeType", Type: "string", Required: false, Default: "Houses", Min: "", Max: ""},
		{Name: "days_on_zillow", Type: "string", Required: false, Default: "Any", Min: "", Max: ""},
		{Name: "maximum", Type: "string", Required: false, Default: "10", Min: "", Max: ""},
		{Name: "file_name", Type: "string", Required: false, Default: "{{TasksID}}", Min: "", Max: ""},
	}},
}
