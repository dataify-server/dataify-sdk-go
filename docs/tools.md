# Dataify Tool Manifest

This document is generated from the MCP tool source. Files are UTF-8.

## `airbnb_product`

- SDK: `Airbnb.Product(ctx, dataify.AirbnbProductRequest{...})`
- Service kind: `scraper_builder`
- Spider name: `airbnb.com`
- Default spider ID: `airbnb_product_by-searchurl`

| Parameter | Go Field | Type | Required | Default | Min | Max |
| --- | --- | --- | --- | --- | --- | --- |
| `searchurl` | `Searchurl` | `string` | `true` | `https://www.airbnb.com/s/Greece/homes?query=Greece&refinement_paths%5B%5D=%2Fhomes&place_id=ChIJY2xxEcdKWxMRHS2a3HUXOjY&flexible_trip_lengths%5B%5D=one_week&monthly_start_date=2025-03-01&monthly_length=3&monthly_end_date=2025-06-01&search_mode=regular_search&price_filter_input_type=0&channel=EXPLORE&date_picker_type=calendar&source=structured_search_input_header&search_type=filter_change&price_filter_num_nights=5&flexible_date_search_filter_type=1` | `` | `` |
| `country` | `Country` | `string` | `false` | `HK` | `` | `` |
| `file_name` | `FileName` | `string` | `false` | `{{TasksID}}` | `` | `` |

## `amazon_comment`

- SDK: `Amazon.Comment(ctx, dataify.AmazonCommentRequest{...})`
- Service kind: `scraper_builder`
- Spider name: `amazon.com`
- Default spider ID: `amazon_comment_by-url`

| Parameter | Go Field | Type | Required | Default | Min | Max |
| --- | --- | --- | --- | --- | --- | --- |
| `url` | `URL` | `string` | `true` | `https://www.amazon.com/HISDERN-Checkered-Handkerchief-Classic-Necktie/dp/B0BRXPR726` | `` | `` |
| `file_name` | `FileName` | `string` | `false` | `{{TasksID}}` | `` | `` |

## `amazon_global_product`

- SDK: `Amazon.GlobalProduct(ctx, dataify.AmazonGlobalProductRequest{...})`
- Service kind: `scraper_builder`
- Spider name: `amazon.com`
- Default spider ID: `amazon_global-product_by-url`

| Parameter | Go Field | Type | Required | Default | Min | Max |
| --- | --- | --- | --- | --- | --- | --- |
| `spider_id` | `SpiderID` | `string` | `true` | `amazon_global-product_by-url` | `` | `` |
| `url` | `URL` | `string` | `false` | `` | `` | `` |
| `maximum` | `Maximum` | `string` | `false` | `5` | `` | `` |
| `sort_by` | `SortBy` | `string` | `false` | `畅销排行` | `` | `` |
| `get_sponsored` | `GetSponsored` | `string` | `false` | `true` | `` | `` |
| `keyword` | `Keyword` | `string` | `false` | `` | `` | `` |
| `domain` | `Domain` | `string` | `false` | `https://www.amazon.com` | `` | `` |
| `brands` | `Brands` | `string` | `false` | `Adidas` | `` | `` |
| `page_turning` | `PageTurning` | `string` | `false` | `2` | `` | `` |
| `lowest_price` | `LowestPrice` | `string` | `false` | `20` | `` | `` |
| `highest_price` | `HighestPrice` | `string` | `false` | `50` | `` | `` |
| `file_name` | `FileName` | `string` | `false` | `{{TasksID}}` | `` | `` |

## `amazon_product`

- SDK: `Amazon.Product(ctx, dataify.AmazonProductRequest{...})`
- Service kind: `scraper_builder`
- Spider name: `amazon.com`
- Default spider ID: `amazon_product_by-asin`

| Parameter | Go Field | Type | Required | Default | Min | Max |
| --- | --- | --- | --- | --- | --- | --- |
| `spider_id` | `SpiderID` | `string` | `true` | `amazon_product_by-asin` | `` | `` |
| `asin` | `Asin` | `string` | `false` | `B0BZYCJK89` | `` | `` |
| `url` | `URL` | `string` | `false` | `` | `` | `` |
| `category_url` | `CategoryURL` | `string` | `false` | `` | `` | `` |
| `keyword` | `Keyword` | `string` | `false` | `coffee` | `` | `` |
| `page_turning` | `PageTurning` | `string` | `false` | `` | `` | `` |
| `lowest_price` | `LowestPrice` | `string` | `false` | `20` | `` | `` |
| `highest_price` | `HighestPrice` | `string` | `false` | `50` | `` | `` |
| `sort_by` | `SortBy` | `string` | `false` | `畅销排行` | `` | `` |
| `collect_subcategories` | `CollectSubcategories` | `string` | `false` | `` | `` | `` |
| `zip_code` | `ZipCode` | `string` | `false` | `94107` | `` | `` |
| `file_name` | `FileName` | `string` | `false` | `{{TasksID}}` | `` | `` |

## `amazon_product_list`

- SDK: `Amazon.ProductList(ctx, dataify.AmazonProductListRequest{...})`
- Service kind: `scraper_builder`
- Spider name: `amazon.com`
- Default spider ID: `amazon_product-list_by-keywords-domain`

| Parameter | Go Field | Type | Required | Default | Min | Max |
| --- | --- | --- | --- | --- | --- | --- |
| `keyword` | `Keyword` | `string` | `true` | `X-box` | `` | `` |
| `domain` | `Domain` | `string` | `true` | `https://www.amazon.com/` | `` | `` |
| `page_turning` | `PageTurning` | `string` | `false` | `1` | `` | `` |
| `file_name` | `FileName` | `string` | `false` | `{{TasksID}}` | `` | `` |

## `amazon_seller`

- SDK: `Amazon.Seller(ctx, dataify.AmazonSellerRequest{...})`
- Service kind: `scraper_builder`
- Spider name: `amazon.com`
- Default spider ID: `amazon_seller_by-url`

| Parameter | Go Field | Type | Required | Default | Min | Max |
| --- | --- | --- | --- | --- | --- | --- |
| `url` | `URL` | `string` | `true` | `https://www.amazon.com/sp?ie=UTF8&seller=ADZ7LD48GVFQJ&asin=B07H56J7K1&ref_=dp_merchant_link&isAmazonFulfilled=1` | `` | `` |
| `file_name` | `FileName` | `string` | `false` | `{{TasksID}}` | `` | `` |

## `bing_images`

- SDK: `Bing.Images(ctx, dataify.BingImagesRequest{...})`
- Service kind: `serp`
- Engine: `bing_images`

| Parameter | Go Field | Type | Required | Default | Min | Max |
| --- | --- | --- | --- | --- | --- | --- |
| `q` | `Q` | `string` | `true` | `Pizza` | `` | `` |
| `json` | `JSON` | `string` | `true` | `1` | `` | `` |
| `mkt` | `Mkt` | `string` | `false` | `` | `` | `` |
| `cc` | `Cc` | `string` | `false` | `` | `` | `` |
| `first` | `First` | `string` | `false` | `0` | `` | `` |
| `count` | `Count` | `string` | `false` | `` | `` | `` |
| `imagesize` | `Imagesize` | `string` | `false` | `` | `` | `` |
| `color2` | `Color2` | `string` | `false` | `` | `` | `` |
| `photo` | `Photo` | `string` | `false` | `` | `` | `` |
| `aspect` | `Aspect` | `string` | `false` | `` | `` | `` |
| `face` | `Face` | `string` | `false` | `` | `` | `` |
| `age` | `Age` | `string` | `false` | `` | `` | `` |
| `license` | `License` | `string` | `false` | `` | `` | `` |
| `no_cache` | `NoCache` | `string` | `false` | `false` | `` | `` |

## `bing_maps`

- SDK: `Bing.Maps(ctx, dataify.BingMapsRequest{...})`
- Service kind: `serp`
- Engine: `bing_maps`

| Parameter | Go Field | Type | Required | Default | Min | Max |
| --- | --- | --- | --- | --- | --- | --- |
| `q` | `Q` | `string` | `true` | `` | `` | `` |
| `json` | `JSON` | `string` | `true` | `1` | `` | `` |
| `cp` | `Cp` | `string` | `false` | `` | `` | `` |
| `setlang` | `Setlang` | `string` | `false` | `` | `` | `` |
| `place_id` | `PlaceID` | `string` | `false` | `` | `` | `` |
| `first` | `First` | `string` | `false` | `0` | `` | `` |
| `count` | `Count` | `string` | `false` | `` | `` | `` |
| `no_cache` | `NoCache` | `string` | `false` | `false` | `` | `` |

## `bing_news`

- SDK: `Bing.News(ctx, dataify.BingNewsRequest{...})`
- Service kind: `serp`
- Engine: `bing_news`

| Parameter | Go Field | Type | Required | Default | Min | Max |
| --- | --- | --- | --- | --- | --- | --- |
| `q` | `Q` | `string` | `true` | `Pizza` | `` | `` |
| `json` | `JSON` | `string` | `true` | `1` | `` | `` |
| `mkt` | `Mkt` | `string` | `false` | `` | `` | `` |
| `cc` | `Cc` | `string` | `false` | `` | `` | `` |
| `first` | `First` | `string` | `false` | `0` | `` | `` |
| `count` | `Count` | `string` | `false` | `` | `` | `` |
| `qft` | `Qft` | `string` | `false` | `` | `` | `` |
| `safeSearch` | `Safesearch` | `string` | `false` | `` | `` | `` |
| `no_cache` | `NoCache` | `string` | `false` | `false` | `` | `` |

## `bing_search`

- SDK: `Bing.Search(ctx, dataify.BingSearchRequest{...})`
- Service kind: `serp`
- Engine: `bing`

| Parameter | Go Field | Type | Required | Default | Min | Max |
| --- | --- | --- | --- | --- | --- | --- |
| `q` | `Q` | `string` | `true` | `Pizza` | `` | `` |
| `json` | `JSON` | `string` | `true` | `1` | `` | `` |
| `location` | `Location` | `string` | `false` | `` | `` | `` |
| `lat` | `Lat` | `string` | `false` | `` | `` | `` |
| `lon` | `Lon` | `string` | `false` | `` | `` | `` |
| `mkt` | `Mkt` | `string` | `false` | `` | `` | `` |
| `cc` | `Cc` | `string` | `false` | `` | `` | `` |
| `first` | `First` | `string` | `false` | `0` | `` | `` |
| `safeSearch` | `Safesearch` | `string` | `false` | `` | `` | `` |
| `filters` | `Filters` | `string` | `false` | `` | `` | `` |
| `no_cache` | `NoCache` | `string` | `false` | `false` | `` | `` |

## `bing_shopping`

- SDK: `Bing.Shopping(ctx, dataify.BingShoppingRequest{...})`
- Service kind: `serp`
- Engine: `bing_shopping`

| Parameter | Go Field | Type | Required | Default | Min | Max |
| --- | --- | --- | --- | --- | --- | --- |
| `q` | `Q` | `string` | `true` | `Pizza` | `` | `` |
| `json` | `JSON` | `string` | `true` | `1` | `` | `` |
| `mkt` | `Mkt` | `string` | `false` | `` | `` | `` |
| `cc` | `Cc` | `string` | `false` | `` | `` | `` |
| `efirst` | `Efirst` | `string` | `false` | `` | `` | `` |
| `filters` | `Filters` | `string` | `false` | `` | `` | `` |
| `no_cache` | `NoCache` | `string` | `false` | `false` | `` | `` |

## `bing_videos`

- SDK: `Bing.Videos(ctx, dataify.BingVideosRequest{...})`
- Service kind: `serp`
- Engine: `bing_videos`

| Parameter | Go Field | Type | Required | Default | Min | Max |
| --- | --- | --- | --- | --- | --- | --- |
| `q` | `Q` | `string` | `true` | `Pizza` | `` | `` |
| `json` | `JSON` | `string` | `true` | `1` | `` | `` |
| `mkt` | `Mkt` | `string` | `false` | `` | `` | `` |
| `cc` | `Cc` | `string` | `false` | `` | `` | `` |
| `setlang` | `Setlang` | `string` | `false` | `` | `` | `` |
| `first` | `First` | `string` | `false` | `0` | `` | `` |
| `length` | `Length` | `string` | `false` | `` | `` | `` |
| `date` | `Date` | `string` | `false` | `` | `` | `` |
| `resolution` | `Resolution` | `string` | `false` | `` | `` | `` |
| `source_site` | `SourceSite` | `string` | `false` | `` | `` | `` |
| `price` | `Price` | `string` | `false` | `` | `` | `` |
| `no_cache` | `NoCache` | `string` | `false` | `false` | `` | `` |

## `booking_hotellist`

- SDK: `Booking.Hotellist(ctx, dataify.BookingHotellistRequest{...})`
- Service kind: `scraper_builder`
- Spider name: `booking.com`
- Default spider ID: `booking_hotellist_by-url`

| Parameter | Go Field | Type | Required | Default | Min | Max |
| --- | --- | --- | --- | --- | --- | --- |
| `url` | `URL` | `string` | `true` | `https://www.booking.com/hotel/gb/westlands-of-pitlochry.en-gb.html#tab-main` | `` | `` |
| `file_name` | `FileName` | `string` | `false` | `{{TasksID}}` | `` | `` |

## `crunchbase_company`

- SDK: `Crunchbase.Company(ctx, dataify.CrunchbaseCompanyRequest{...})`
- Service kind: `scraper_builder`
- Spider name: `crunchbase.com`
- Default spider ID: `crunchbase_company_by-url`

| Parameter | Go Field | Type | Required | Default | Min | Max |
| --- | --- | --- | --- | --- | --- | --- |
| `spider_id` | `SpiderID` | `string` | `true` | `crunchbase_company_by-url` | `` | `` |
| `url` | `URL` | `string` | `false` | `https://www.crunchbase.com/organization/aisci` | `` | `` |
| `keyword` | `Keyword` | `string` | `false` | `NetBooster` | `` | `` |
| `file_name` | `FileName` | `string` | `false` | `{{TasksID}}` | `` | `` |

## `duckduckgo_search`

- SDK: `DuckDuckGo.Search(ctx, dataify.DuckDuckGoSearchRequest{...})`
- Service kind: `serp`
- Engine: `duckduckgo`

| Parameter | Go Field | Type | Required | Default | Min | Max |
| --- | --- | --- | --- | --- | --- | --- |
| `q` | `Q` | `string` | `true` | `Pizza` | `` | `` |
| `json` | `JSON` | `string` | `true` | `1` | `` | `` |
| `kl` | `Kl` | `string` | `false` | `` | `` | `` |
| `search_assist` | `SearchAssist` | `string` | `false` | `false` | `` | `` |
| `safe` | `Safe` | `string` | `false` | `-1` | `` | `` |
| `df` | `Df` | `string` | `false` | `` | `` | `` |
| `start` | `Start` | `string` | `false` | `0` | `` | `` |
| `m` | `M` | `string` | `false` | `10` | `` | `` |
| `no_cache` | `NoCache` | `string` | `false` | `false` | `` | `` |

## `ebay_info`

- SDK: `EBay.Info(ctx, dataify.EBayInfoRequest{...})`
- Service kind: `scraper_builder`
- Spider name: `ebay.com`
- Default spider ID: `ebay_ebay_by-url`

| Parameter | Go Field | Type | Required | Default | Min | Max |
| --- | --- | --- | --- | --- | --- | --- |
| `spider_id` | `SpiderID` | `string` | `true` | `ebay_ebay_by-url` | `` | `` |
| `url` | `URL` | `string` | `false` | `https://www.ebay.com/itm/296197468977?itmmeta=01HRWJ04NFHYT9AX0XZB8F18G1&hash=item44f6beb331%3Ag%3ADEQAAOSw3CxlhTJ%7E&_trkparms=%2526rpp_cid%253D6523c97b0b7882040b9472b6` | `` | `` |
| `keywords` | `Keywords` | `string` | `false` | `baby toys` | `` | `` |
| `count` | `Count` | `string` | `false` | `60` | `` | `` |
| `file_name` | `FileName` | `string` | `false` | `{{TasksID}}` | `` | `` |

## `facebook_comment`

- SDK: `Facebook.Comment(ctx, dataify.FacebookCommentRequest{...})`
- Service kind: `scraper_builder`
- Spider name: `facebook.com`
- Default spider ID: `facebook_comment_by-comments-url`

| Parameter | Go Field | Type | Required | Default | Min | Max |
| --- | --- | --- | --- | --- | --- | --- |
| `url` | `URL` | `string` | `true` | `https://www.facebook.com/share/p/1K6xfHFkrK/` | `` | `` |
| `get_all_replies` | `GetAllReplies` | `string` | `false` | `True` | `` | `` |
| `limit_records` | `LimitRecords` | `string` | `false` | `10` | `` | `` |
| `comments_sort` | `CommentsSort` | `string` | `false` | `所有评论` | `` | `` |
| `file_name` | `FileName` | `string` | `false` | `{{TasksID}}` | `` | `` |

## `facebook_event`

- SDK: `Facebook.Event(ctx, dataify.FacebookEventRequest{...})`
- Service kind: `scraper_builder`
- Spider name: `facebook.com`
- Default spider ID: `facebook_event_by-eventlist-url`

| Parameter | Go Field | Type | Required | Default | Min | Max |
| --- | --- | --- | --- | --- | --- | --- |
| `spider_id` | `SpiderID` | `string` | `true` | `facebook_event_by-eventlist-url` | `` | `` |
| `url` | `URL` | `string` | `true` | `https://www.facebook.com/nohoclub/events` | `` | `` |
| `file_name` | `FileName` | `string` | `false` | `{{TasksID}}` | `` | `` |

## `facebook_post`

- SDK: `Facebook.Post(ctx, dataify.FacebookPostRequest{...})`
- Service kind: `scraper_builder`
- Spider name: `facebook.com`
- Default spider ID: `facebook_post_by-posts-url`

| Parameter | Go Field | Type | Required | Default | Min | Max |
| --- | --- | --- | --- | --- | --- | --- |
| `url` | `URL` | `string` | `true` | `https://www.facebook.com/permalink.php?story_fbid=pfbid0gNjZBhqCxSqj9xJS5aygNwqFqNEM2fYbTFKKbsvvGdEfTgFyAYWSckvkEHPqAE7gl&id=61574926580533&rdid=86oaujwNGCCdPLfj#` | `` | `` |
| `file_name` | `FileName` | `string` | `false` | `{{TasksID}}` | `` | `` |

## `facebook_profile`

- SDK: `Facebook.Profile(ctx, dataify.FacebookProfileRequest{...})`
- Service kind: `scraper_builder`
- Spider name: `facebook.com`
- Default spider ID: `facebook_profile_by-profiles-url`

| Parameter | Go Field | Type | Required | Default | Min | Max |
| --- | --- | --- | --- | --- | --- | --- |
| `url` | `URL` | `string` | `true` | `https://www.facebook.com/MayeMusk` | `` | `` |
| `file_name` | `FileName` | `string` | `false` | `{{TasksID}}` | `` | `` |

## `github_repository`

- SDK: `GitHub.Repository(ctx, dataify.GitHubRepositoryRequest{...})`
- Service kind: `scraper_builder`
- Spider name: `github.com`
- Default spider ID: `github_repository_by-repo-url`

| Parameter | Go Field | Type | Required | Default | Min | Max |
| --- | --- | --- | --- | --- | --- | --- |
| `spider_id` | `SpiderID` | `string` | `true` | `github_repository_by-repo-url` | `` | `` |
| `repo_url` | `RepoURL` | `string` | `false` | `https://github.com/TheAlgorithms/Python` | `` | `` |
| `search_url` | `SearchURL` | `string` | `false` | `https://github.com/search?q=ML&type=repositories` | `` | `` |
| `url` | `URL` | `string` | `false` | `https://github.com/TheAlgorithms/Python/blob/master/divide_and_conquer/power.py` | `` | `` |
| `page_turning` | `PageTurning` | `string` | `false` | `1` | `` | `` |
| `max_num` | `MaxNum` | `string` | `false` | `15` | `` | `` |
| `file_name` | `FileName` | `string` | `false` | `{{TasksID}}` | `` | `` |

## `glassdoor_company`

- SDK: `Glassdoor.Company(ctx, dataify.GlassdoorCompanyRequest{...})`
- Service kind: `scraper_builder`
- Spider name: `glassdoor.com`
- Default spider ID: `glassdoor_company_by-url`

| Parameter | Go Field | Type | Required | Default | Min | Max |
| --- | --- | --- | --- | --- | --- | --- |
| `spider_id` | `SpiderID` | `string` | `true` | `glassdoor_company_by-url` | `` | `` |
| `url` | `URL` | `string` | `false` | `https://www.glassdoor.co.uk/Overview/Working-at-Apple-EI_IE1138.11,16.htm` | `` | `` |
| `location` | `Location` | `string` | `false` | `United States` | `` | `` |
| `company_name` | `CompanyName` | `string` | `false` | `Tesla` | `` | `` |
| `industries` | `Industries` | `string` | `false` | `Information Technology` | `` | `` |
| `Job title` | `JobTitle` | `string` | `false` | `Data` | `` | `` |
| `search_url` | `SearchURL` | `string` | `false` | `https://www.glassdoor.com/Search/results.htm?keyword=Apple` | `` | `` |
| `max_search_results` | `MaxSearchResults` | `string` | `false` | `5` | `` | `` |
| `file_name` | `FileName` | `string` | `false` | `{{TasksID}}` | `` | `` |

## `glassdoor_joblistings`

- SDK: `Glassdoor.Joblistings(ctx, dataify.GlassdoorJoblistingsRequest{...})`
- Service kind: `scraper_builder`
- Spider name: `glassdoor.com`
- Default spider ID: `glassdoor_joblistings_by-url`

| Parameter | Go Field | Type | Required | Default | Min | Max |
| --- | --- | --- | --- | --- | --- | --- |
| `spider_id` | `SpiderID` | `string` | `true` | `glassdoor_joblistings_by-url` | `` | `` |
| `url` | `URL` | `string` | `false` | `https://www.glassdoor.com/Job/new-york-data-analyst-jobs-SRCH_IL.0,8_IC1132348_KO9,21.htm` | `` | `` |
| `keyword` | `Keyword` | `string` | `false` | `data analyst` | `` | `` |
| `location` | `Location` | `string` | `false` | `New York` | `` | `` |
| `country` | `Country` | `string` | `false` | `US` | `` | `` |
| `file_name` | `FileName` | `string` | `false` | `{{TasksID}}` | `` | `` |

## `google_ai_mode`

- SDK: `Google.AIMode(ctx, dataify.GoogleAIModeRequest{...})`
- Service kind: `serp`
- Engine: `google_ai_mode`

| Parameter | Go Field | Type | Required | Default | Min | Max |
| --- | --- | --- | --- | --- | --- | --- |
| `q` | `Q` | `string` | `true` | `pizza` | `` | `` |
| `json` | `JSON` | `string` | `true` | `1` | `` | `` |
| `location` | `Location` | `string` | `false` | `` | `` | `` |
| `uule` | `Uule` | `string` | `false` | `` | `` | `` |
| `no_cache` | `NoCache` | `string` | `false` | `false` | `` | `` |
| `gl` | `Gl` | `string` | `false` | `` | `` | `` |
| `hl` | `Hl` | `string` | `false` | `` | `` | `` |

## `google_finance`

- SDK: `Google.Finance(ctx, dataify.GoogleFinanceRequest{...})`
- Service kind: `serp`
- Engine: `google_finance`

| Parameter | Go Field | Type | Required | Default | Min | Max |
| --- | --- | --- | --- | --- | --- | --- |
| `q` | `Q` | `string` | `true` | `` | `` | `` |
| `json` | `JSON` | `string` | `true` | `1` | `` | `` |
| `hl` | `Hl` | `string` | `false` | `` | `` | `` |
| `window` | `Window` | `string` | `false` | `1D` | `` | `` |
| `no_cache` | `NoCache` | `string` | `false` | `false` | `` | `` |

## `google_flights`

- SDK: `Google.Flights(ctx, dataify.GoogleFlightsRequest{...})`
- Service kind: `serp`
- Engine: `google_flights`

| Parameter | Go Field | Type | Required | Default | Min | Max |
| --- | --- | --- | --- | --- | --- | --- |
| `departure_id` | `DepartureID` | `string` | `false` | `` | `` | `` |
| `arrival_id` | `ArrivalID` | `string` | `false` | `` | `` | `` |
| `json` | `JSON` | `string` | `true` | `1` | `` | `` |
| `gl` | `Gl` | `string` | `false` | `` | `` | `` |
| `hl` | `Hl` | `string` | `false` | `` | `` | `` |
| `currency` | `Currency` | `string` | `false` | `USD` | `` | `` |
| `type` | `Type` | `string` | `false` | `1` | `` | `` |
| `outbound_date` | `OutboundDate` | `string` | `false` | `` | `` | `` |
| `return_date` | `ReturnDate` | `string` | `false` | `` | `` | `` |
| `travel_class` | `TravelClass` | `string` | `false` | `1` | `` | `` |
| `multi_city_json` | `MultiCityJSON` | `string` | `false` | `` | `` | `` |
| `show_hidden` | `ShowHidden` | `string` | `false` | `false` | `` | `` |
| `exclude_basic` | `ExcludeBasic` | `string` | `false` | `false` | `` | `` |
| `deep_search` | `DeepSearch` | `string` | `false` | `false` | `` | `` |
| `adults` | `Adults` | `string` | `false` | `1` | `` | `` |
| `children` | `Children` | `string` | `false` | `0` | `` | `` |
| `infants_in_seat` | `InfantsInSeat` | `string` | `false` | `0` | `` | `` |
| `infants_on_lap` | `InfantsOnLap` | `string` | `false` | `0` | `` | `` |
| `sort_by` | `SortBy` | `string` | `false` | `1` | `` | `` |
| `stops` | `Stops` | `string` | `false` | `0` | `` | `` |
| `exclude_airlines` | `ExcludeAirlines` | `string` | `false` | `` | `` | `` |
| `include_airlines` | `IncludeAirlines` | `string` | `false` | `` | `` | `` |
| `bags` | `Bags` | `string` | `false` | `0` | `` | `` |
| `max_price` | `MaxPrice` | `string` | `false` | `` | `` | `` |
| `outbound_times` | `OutboundTimes` | `string` | `false` | `` | `` | `` |
| `return_times` | `ReturnTimes` | `string` | `false` | `` | `` | `` |
| `emissions` | `Emissions` | `string` | `false` | `` | `` | `` |
| `layover_duration` | `LayoverDuration` | `string` | `false` | `` | `` | `` |
| `exclude_conns` | `ExcludeConns` | `string` | `false` | `` | `` | `` |
| `max_duration` | `MaxDuration` | `string` | `false` | `` | `` | `` |
| `departure_token` | `DepartureToken` | `string` | `false` | `` | `` | `` |
| `no_cache` | `NoCache` | `string` | `false` | `false` | `` | `` |

## `google_hotels`

- SDK: `Google.Hotels(ctx, dataify.GoogleHotelsRequest{...})`
- Service kind: `serp`
- Engine: `google_hotels`

| Parameter | Go Field | Type | Required | Default | Min | Max |
| --- | --- | --- | --- | --- | --- | --- |
| `q` | `Q` | `string` | `true` | `pizza` | `` | `` |
| `json` | `JSON` | `string` | `true` | `1` | `` | `` |
| `hl` | `Hl` | `string` | `false` | `` | `` | `` |
| `gl` | `Gl` | `string` | `false` | `` | `` | `` |
| `currency` | `Currency` | `string` | `false` | `USD` | `` | `` |
| `check_in_date` | `CheckInDate` | `string` | `true` | `` | `` | `` |
| `check_out_date` | `CheckOutDate` | `string` | `true` | `` | `` | `` |
| `adults` | `Adults` | `string` | `false` | `2` | `` | `` |
| `children` | `Children` | `string` | `false` | `0` | `` | `` |
| `children_ages` | `ChildrenAges` | `string` | `false` | `` | `` | `` |
| `sort_by` | `SortBy` | `string` | `false` | `` | `` | `` |
| `min_price` | `MinPrice` | `string` | `false` | `` | `` | `` |
| `max_price` | `MaxPrice` | `string` | `false` | `` | `` | `` |
| `property_types` | `PropertyTypes` | `string` | `false` | `` | `` | `` |
| `amenities` | `Amenities` | `string` | `false` | `` | `` | `` |
| `rating` | `Rating` | `string` | `false` | `` | `` | `` |
| `brands` | `Brands` | `string` | `false` | `` | `` | `` |
| `hotel_class` | `HotelClass` | `string` | `false` | `` | `` | `` |
| `free_cancellation` | `FreeCancellation` | `string` | `false` | `` | `` | `` |
| `special_offers` | `SpecialOffers` | `string` | `false` | `` | `` | `` |
| `eco_certified` | `EcoCertified` | `string` | `false` | `` | `` | `` |
| `vacation_rentals` | `VacationRentals` | `string` | `false` | `` | `` | `` |
| `bedrooms` | `Bedrooms` | `string` | `false` | `0` | `` | `` |
| `bathrooms` | `Bathrooms` | `string` | `false` | `0` | `` | `` |
| `next_page_token` | `NextPageToken` | `string` | `false` | `` | `` | `` |
| `no_cache` | `NoCache` | `string` | `false` | `false` | `` | `` |
| `property_token` | `PropertyToken` | `string` | `false` | `` | `` | `` |

## `google_images`

- SDK: `Google.Images(ctx, dataify.GoogleImagesRequest{...})`
- Service kind: `serp`
- Engine: `google_images`

| Parameter | Go Field | Type | Required | Default | Min | Max |
| --- | --- | --- | --- | --- | --- | --- |
| `q` | `Q` | `string` | `true` | `pizza` | `` | `` |
| `json` | `JSON` | `string` | `true` | `1` | `` | `` |
| `google_domain` | `GoogleDomain` | `string` | `false` | `google.com` | `` | `` |
| `gl` | `Gl` | `string` | `false` | `` | `` | `` |
| `hl` | `Hl` | `string` | `false` | `` | `` | `` |
| `cr` | `Cr` | `string` | `false` | `` | `` | `` |
| `lr` | `Lr` | `string` | `false` | `` | `` | `` |
| `location` | `Location` | `string` | `false` | `` | `` | `` |
| `uule` | `Uule` | `string` | `false` | `` | `` | `` |
| `lat` | `Lat` | `string` | `false` | `` | `` | `` |
| `lon` | `Lon` | `string` | `false` | `` | `` | `` |
| `radius` | `Radius` | `string` | `false` | `` | `` | `` |
| `start` | `Start` | `string` | `false` | `0` | `` | `` |
| `tbm` | `Tbm` | `string` | `false` | `isch` | `` | `` |
| `ludocid` | `Ludocid` | `string` | `false` | `` | `` | `` |
| `lsig` | `Lsig` | `string` | `false` | `` | `` | `` |
| `kgmid` | `Kgmid` | `string` | `false` | `` | `` | `` |
| `si` | `Si` | `string` | `false` | `` | `` | `` |
| `ibp` | `Ibp` | `string` | `false` | `` | `` | `` |
| `uds` | `Uds` | `string` | `false` | `` | `` | `` |
| `tbs` | `Tbs` | `string` | `false` | `` | `` | `` |
| `safe` | `Safe` | `string` | `false` | `` | `` | `` |
| `nfpr` | `Nfpr` | `string` | `false` | `` | `` | `` |
| `filter` | `Filter` | `string` | `false` | `1` | `` | `` |
| `device` | `Device` | `string` | `false` | `desktop` | `` | `` |
| `render_js` | `RenderJs` | `string` | `false` | `` | `` | `` |
| `no_cache` | `NoCache` | `string` | `false` | `false` | `` | `` |
| `ai_overview` | `AIOverview` | `string` | `false` | `` | `` | `` |

## `google_jobs`

- SDK: `Google.Jobs(ctx, dataify.GoogleJobsRequest{...})`
- Service kind: `serp`
- Engine: `google_jobs`

| Parameter | Go Field | Type | Required | Default | Min | Max |
| --- | --- | --- | --- | --- | --- | --- |
| `q` | `Q` | `string` | `true` | `` | `` | `` |
| `json` | `JSON` | `string` | `true` | `1` | `` | `` |
| `google_domain` | `GoogleDomain` | `string` | `false` | `google.com` | `` | `` |
| `gl` | `Gl` | `string` | `false` | `` | `` | `` |
| `hl` | `Hl` | `string` | `false` | `` | `` | `` |
| `location` | `Location` | `string` | `false` | `` | `` | `` |
| `uule` | `Uule` | `string` | `false` | `` | `` | `` |
| `next_page_token` | `NextPageToken` | `string` | `false` | `` | `` | `` |
| `chips` | `Chips` | `string` | `false` | `` | `` | `` |
| `lrad` | `Lrad` | `string` | `false` | `` | `` | `` |
| `ltype` | `Ltype` | `string` | `false` | `` | `` | `` |
| `uds` | `Uds` | `string` | `false` | `` | `` | `` |
| `no_cache` | `NoCache` | `string` | `false` | `false` | `` | `` |

## `google_lens`

- SDK: `Google.Lens(ctx, dataify.GoogleLensRequest{...})`
- Service kind: `serp`
- Engine: `google_lens`

| Parameter | Go Field | Type | Required | Default | Min | Max |
| --- | --- | --- | --- | --- | --- | --- |
| `url` | `URL` | `string` | `true` | `` | `` | `` |
| `json` | `JSON` | `string` | `true` | `1` | `` | `` |
| `hl` | `Hl` | `string` | `false` | `` | `` | `` |
| `country` | `Country` | `string` | `false` | `` | `` | `` |
| `type` | `Type` | `string` | `false` | `all` | `` | `` |
| `q` | `Q` | `string` | `false` | `` | `` | `` |
| `safe` | `Safe` | `string` | `false` | `` | `` | `` |
| `no_cache` | `NoCache` | `string` | `false` | `false` | `` | `` |

## `google_local`

- SDK: `Google.Local(ctx, dataify.GoogleLocalRequest{...})`
- Service kind: `serp`
- Engine: `google_local`

| Parameter | Go Field | Type | Required | Default | Min | Max |
| --- | --- | --- | --- | --- | --- | --- |
| `q` | `Q` | `string` | `true` | `` | `` | `` |
| `json` | `JSON` | `string` | `true` | `1` | `` | `` |
| `google_domain` | `GoogleDomain` | `string` | `false` | `google.com` | `` | `` |
| `gl` | `Gl` | `string` | `false` | `` | `` | `` |
| `hl` | `Hl` | `string` | `false` | `` | `` | `` |
| `location` | `Location` | `string` | `false` | `` | `` | `` |
| `uule` | `Uule` | `string` | `false` | `` | `` | `` |
| `start` | `Start` | `string` | `false` | `` | `` | `` |
| `ludocid` | `Ludocid` | `string` | `false` | `` | `` | `` |
| `tbs` | `Tbs` | `string` | `false` | `` | `` | `` |
| `no_cache` | `NoCache` | `string` | `false` | `false` | `` | `` |

## `google_map_comment`

- SDK: `Google.MapComment(ctx, dataify.GoogleMapCommentRequest{...})`
- Service kind: `scraper_builder`
- Spider name: `google.com`
- Default spider ID: `google_comment_by-url`

| Parameter | Go Field | Type | Required | Default | Min | Max |
| --- | --- | --- | --- | --- | --- | --- |
| `url` | `URL` | `string` | `true` | `https://www.google.com/maps/place/Waterfront+Botanical+Gardens/@38.2630366,-85.7288454,15z/data=!4m8!3m7!1s0x8869731e16a7bdbd:0x2f5d238fefed7ca1!8m2!3d38.2632837!4d-85.7239738!9m1!1b1!16s%2Fg%2F11c709xzzx?hl=en&entry=ttu` | `` | `` |
| `days_limit` | `DaysLimit` | `string` | `false` | `20` | `` | `` |
| `file_name` | `FileName` | `string` | `false` | `{{TasksID}}` | `` | `` |

## `google_map_details`

- SDK: `Google.MapDetails(ctx, dataify.GoogleMapDetailsRequest{...})`
- Service kind: `scraper_builder`
- Spider name: `google.com`
- Default spider ID: `google_map-details_by-url`

| Parameter | Go Field | Type | Required | Default | Min | Max |
| --- | --- | --- | --- | --- | --- | --- |
| `spider_id` | `SpiderID` | `string` | `true` | `google_map-details_by-url` | `` | `` |
| `url` | `URL` | `string` | `false` | `https://www.google.com/maps/place/Pizza+Inn+Magdeburg/data=!4m7!3m6!1s0x47a5f50c083530a3:0xfdba8746b538141!8m2!3d52.1263086!4d11.6094743!16s%2Fg%2F11kqmtk3dt!19sChIJozA1CAz1pUcRQYFTa3So2w8?authuser=0&hl=en&rclk=1` | `` | `` |
| `CID` | `CID` | `string` | `false` | `2476046430038551731` | `` | `` |
| `keyword` | `Keyword` | `string` | `false` | `pizza` | `` | `` |
| `country` | `Country` | `string` | `false` | `United States` | `` | `` |
| `lat` | `Lat` | `string` | `false` | `38` | `` | `` |
| `long` | `Long` | `string` | `false` | `77` | `` | `` |
| `zoom_level` | `ZoomLevel` | `string` | `false` | `20` | `` | `` |
| `place_id` | `PlaceID` | `string` | `false` | `ChIJ3S-JXmauEmsRUcIaWtf4MzE` | `` | `` |
| `file_name` | `FileName` | `string` | `false` | `{{TasksID}}` | `` | `` |

## `google_maps`

- SDK: `Google.Maps(ctx, dataify.GoogleMapsRequest{...})`
- Service kind: `serp`
- Engine: `google_maps`

| Parameter | Go Field | Type | Required | Default | Min | Max |
| --- | --- | --- | --- | --- | --- | --- |
| `q` | `Q` | `string` | `true` | `` | `` | `` |
| `json` | `JSON` | `string` | `true` | `1` | `` | `` |
| `ll` | `Ll` | `string` | `false` | `` | `` | `` |
| `location` | `Location` | `string` | `false` | `` | `` | `` |
| `lat` | `Lat` | `string` | `false` | `` | `` | `` |
| `lon` | `Lon` | `string` | `false` | `` | `` | `` |
| `z` | `Z` | `string` | `false` | `` | `` | `` |
| `m` | `M` | `string` | `false` | `` | `` | `` |
| `nearby` | `Nearby` | `string` | `false` | `` | `` | `` |
| `google_domain` | `GoogleDomain` | `string` | `false` | `google.com` | `` | `` |
| `hl` | `Hl` | `string` | `false` | `` | `` | `` |
| `gl` | `Gl` | `string` | `false` | `` | `` | `` |
| `start` | `Start` | `string` | `false` | `0` | `` | `` |
| `type` | `Type` | `string` | `false` | `` | `` | `` |
| `data` | `Data` | `string` | `false` | `` | `` | `` |
| `place_id` | `PlaceID` | `string` | `false` | `` | `` | `` |
| `data_cid` | `DataCID` | `string` | `false` | `` | `` | `` |
| `no_cache` | `NoCache` | `string` | `false` | `false` | `` | `` |

## `google_news`

- SDK: `Google.News(ctx, dataify.GoogleNewsRequest{...})`
- Service kind: `serp`
- Engine: `google_news`

| Parameter | Go Field | Type | Required | Default | Min | Max |
| --- | --- | --- | --- | --- | --- | --- |
| `q` | `Q` | `string` | `true` | `pizza` | `` | `` |
| `json` | `JSON` | `string` | `true` | `1` | `` | `` |
| `no_cache` | `NoCache` | `string` | `false` | `false` | `` | `` |
| `gl` | `Gl` | `string` | `false` | `` | `` | `` |
| `hl` | `Hl` | `string` | `false` | `` | `` | `` |
| `topic_token` | `TopicToken` | `string` | `false` | `` | `` | `` |
| `kgmid` | `Kgmid` | `string` | `false` | `` | `` | `` |
| `publication_token` | `PublicationToken` | `string` | `false` | `` | `` | `` |
| `section_token` | `SectionToken` | `string` | `false` | `` | `` | `` |
| `story_token` | `StoryToken` | `string` | `false` | `` | `` | `` |
| `so` | `So` | `string` | `false` | `0` | `` | `` |

## `google_patents`

- SDK: `Google.Patents(ctx, dataify.GooglePatentsRequest{...})`
- Service kind: `serp`
- Engine: `google_patents`

| Parameter | Go Field | Type | Required | Default | Min | Max |
| --- | --- | --- | --- | --- | --- | --- |
| `q` | `Q` | `string` | `true` | `` | `` | `` |
| `json` | `JSON` | `string` | `true` | `1` | `` | `` |
| `page` | `Page` | `string` | `false` | `1` | `` | `` |
| `num` | `Num` | `string` | `false` | `` | `` | `` |
| `sort` | `Sort` | `string` | `false` | `` | `` | `` |
| `clustered` | `Clustered` | `string` | `false` | `` | `` | `` |
| `dups` | `Dups` | `string` | `false` | `family` | `` | `` |
| `patents` | `Patents` | `string` | `false` | `true` | `` | `` |
| `scholar` | `Scholar` | `string` | `false` | `false` | `` | `` |
| `before` | `Before` | `string` | `false` | `` | `` | `` |
| `after` | `After` | `string` | `false` | `` | `` | `` |
| `inventor` | `Inventor` | `string` | `false` | `` | `` | `` |
| `assignee` | `Assignee` | `string` | `false` | `` | `` | `` |
| `country` | `Country` | `string` | `false` | `` | `` | `` |
| `language` | `Language` | `string` | `false` | `` | `` | `` |
| `status` | `Status` | `string` | `false` | `` | `` | `` |
| `type` | `Type` | `string` | `false` | `` | `` | `` |
| `litigation` | `Litigation` | `string` | `false` | `` | `` | `` |
| `no_cache` | `NoCache` | `string` | `false` | `false` | `` | `` |

## `google_play`

- SDK: `Google.Play(ctx, dataify.GooglePlayRequest{...})`
- Service kind: `serp`
- Engine: `google_play`

| Parameter | Go Field | Type | Required | Default | Min | Max |
| --- | --- | --- | --- | --- | --- | --- |
| `q` | `Q` | `string` | `true` | `` | `` | `` |
| `json` | `JSON` | `string` | `true` | `1` | `` | `` |
| `hl` | `Hl` | `string` | `false` | `` | `` | `` |
| `gl` | `Gl` | `string` | `false` | `us` | `` | `` |
| `apps_category` | `AppsCategory` | `string` | `false` | `` | `` | `` |
| `next_page_token` | `NextPageToken` | `string` | `false` | `` | `` | `` |
| `section_page_token` | `SectionPageToken` | `string` | `false` | `` | `` | `` |
| `chart` | `Chart` | `string` | `false` | `` | `` | `` |
| `see_more_token` | `SeeMoreToken` | `string` | `false` | `` | `` | `` |
| `store_device` | `StoreDevice` | `string` | `false` | `` | `` | `` |
| `age` | `Age` | `string` | `false` | `` | `` | `` |
| `no_cache` | `NoCache` | `string` | `false` | `false` | `` | `` |

## `google_play_store_information`

- SDK: `Google.PlayStoreInformation(ctx, dataify.GooglePlayStoreInformationRequest{...})`
- Service kind: `scraper_builder`
- Spider name: `play.google.com`
- Default spider ID: `google-play-store_information_by-url`

| Parameter | Go Field | Type | Required | Default | Min | Max |
| --- | --- | --- | --- | --- | --- | --- |
| `app_url` | `AppURL` | `string` | `true` | `https://play.google.com/store/apps/details?id=com.linkedin.android` | `` | `` |
| `file_name` | `FileName` | `string` | `false` | `{{TasksID}}` | `` | `` |

## `google_play_store_reviews`

- SDK: `Google.PlayStoreReviews(ctx, dataify.GooglePlayStoreReviewsRequest{...})`
- Service kind: `scraper_builder`
- Spider name: `play.google.com`
- Default spider ID: `google-play-store_reviews_by-url`

| Parameter | Go Field | Type | Required | Default | Min | Max |
| --- | --- | --- | --- | --- | --- | --- |
| `app_url` | `AppURL` | `string` | `true` | `https://play.google.com/store/apps/details?id=com.linkedin.android` | `` | `` |
| `file_name` | `FileName` | `string` | `false` | `{{TasksID}}` | `` | `` |

## `google_scholar`

- SDK: `Google.Scholar(ctx, dataify.GoogleScholarRequest{...})`
- Service kind: `serp`
- Engine: `google_scholar`

| Parameter | Go Field | Type | Required | Default | Min | Max |
| --- | --- | --- | --- | --- | --- | --- |
| `q` | `Q` | `string` | `false` | `` | `` | `` |
| `json` | `JSON` | `string` | `true` | `1` | `` | `` |
| `hl` | `Hl` | `string` | `false` | `` | `` | `` |
| `lr` | `Lr` | `string` | `false` | `` | `` | `` |
| `start` | `Start` | `string` | `false` | `0` | `` | `` |
| `num` | `Num` | `string` | `false` | `10` | `` | `` |
| `cites` | `Cites` | `string` | `false` | `` | `` | `` |
| `as_ylo` | `AsYlo` | `string` | `false` | `` | `` | `` |
| `as_yhi` | `AsYhi` | `string` | `false` | `` | `` | `` |
| `scisbd` | `Scisbd` | `string` | `false` | `0` | `` | `` |
| `cluster` | `Cluster` | `string` | `false` | `` | `` | `` |
| `as_sdt` | `AsSdt` | `string` | `false` | `0` | `` | `` |
| `safe` | `Safe` | `string` | `false` | `` | `` | `` |
| `filter` | `Filter` | `string` | `false` | `1` | `` | `` |
| `as_vis` | `AsVis` | `string` | `false` | `0` | `` | `` |
| `as_rr` | `AsRr` | `string` | `false` | `0` | `` | `` |
| `no_cache` | `NoCache` | `string` | `false` | `false` | `` | `` |

## `google_search`

- SDK: `Google.Search(ctx, dataify.GoogleSearchRequest{...})`
- Service kind: `serp`
- Engine: `google`

| Parameter | Go Field | Type | Required | Default | Min | Max |
| --- | --- | --- | --- | --- | --- | --- |
| `q` | `Q` | `string` | `true` | `pizza` | `` | `` |
| `json` | `JSON` | `string` | `true` | `1` | `` | `` |
| `google_domain` | `GoogleDomain` | `string` | `false` | `google.com` | `` | `` |
| `gl` | `Gl` | `string` | `false` | `` | `` | `` |
| `hl` | `Hl` | `string` | `false` | `` | `` | `` |
| `cr` | `Cr` | `string` | `false` | `` | `` | `` |
| `lr` | `Lr` | `string` | `false` | `` | `` | `` |
| `location` | `Location` | `string` | `false` | `` | `` | `` |
| `uule` | `Uule` | `string` | `false` | `` | `` | `` |
| `start` | `Start` | `string` | `false` | `0` | `` | `` |
| `tbs` | `Tbs` | `string` | `false` | `` | `` | `` |
| `safe` | `Safe` | `string` | `false` | `` | `` | `` |
| `nfpr` | `Nfpr` | `string` | `false` | `` | `` | `` |
| `filter` | `Filter` | `string` | `false` | `1` | `` | `` |
| `device` | `Device` | `string` | `false` | `desktop` | `` | `` |
| `render_js` | `RenderJs` | `string` | `false` | `` | `` | `` |
| `no_cache` | `NoCache` | `string` | `false` | `false` | `` | `` |
| `ai_overview` | `AIOverview` | `string` | `false` | `` | `` | `` |

## `google_shopping`

- SDK: `Google.Shopping(ctx, dataify.GoogleShoppingRequest{...})`
- Service kind: `serp`
- Engine: `google_shopping`

| Parameter | Go Field | Type | Required | Default | Min | Max |
| --- | --- | --- | --- | --- | --- | --- |
| `q` | `Q` | `string` | `false` | `` | `` | `` |
| `json` | `JSON` | `string` | `true` | `1` | `` | `` |
| `google_domain` | `GoogleDomain` | `string` | `false` | `google.com` | `` | `` |
| `gl` | `Gl` | `string` | `false` | `` | `` | `` |
| `hl` | `Hl` | `string` | `false` | `` | `` | `` |
| `location` | `Location` | `string` | `false` | `` | `` | `` |
| `uule` | `Uule` | `string` | `false` | `` | `` | `` |
| `start` | `Start` | `string` | `false` | `` | `` | `` |
| `shoprs` | `Shoprs` | `string` | `false` | `` | `` | `` |
| `min_price` | `MinPrice` | `string` | `false` | `` | `` | `` |
| `max_price` | `MaxPrice` | `string` | `false` | `` | `` | `` |
| `sort_by` | `SortBy` | `string` | `false` | `` | `` | `` |
| `free_shipping` | `FreeShipping` | `string` | `false` | `` | `` | `` |
| `on_sale` | `OnSale` | `string` | `false` | `` | `` | `` |
| `small_business` | `SmallBusiness` | `string` | `false` | `` | `` | `` |
| `no_cache` | `NoCache` | `string` | `false` | `false` | `` | `` |

## `google_shopping_info`

- SDK: `Google.ShoppingInfo(ctx, dataify.GoogleShoppingInfoRequest{...})`
- Service kind: `scraper_builder`
- Spider name: `google.com`
- Default spider ID: `google_shopping_by-keywords`

| Parameter | Go Field | Type | Required | Default | Min | Max |
| --- | --- | --- | --- | --- | --- | --- |
| `keyword` | `Keyword` | `string` | `true` | `iphone` | `` | `` |
| `file_name` | `FileName` | `string` | `false` | `{{TasksID}}` | `` | `` |

## `google_trends`

- SDK: `Google.Trends(ctx, dataify.GoogleTrendsRequest{...})`
- Service kind: `serp`
- Engine: `google_trends`

| Parameter | Go Field | Type | Required | Default | Min | Max |
| --- | --- | --- | --- | --- | --- | --- |
| `q` | `Q` | `string` | `true` | `pizza` | `` | `` |
| `json` | `JSON` | `string` | `true` | `1` | `` | `` |
| `hl` | `Hl` | `string` | `false` | `` | `` | `` |
| `geo` | `Geo` | `string` | `false` | `` | `` | `` |
| `region` | `Region` | `string` | `false` | `` | `` | `` |
| `data_type` | `DataType` | `string` | `false` | `` | `` | `` |
| `tz` | `Tz` | `string` | `false` | `420` | `` | `` |
| `cat` | `Cat` | `string` | `false` | `0` | `` | `` |
| `gprop` | `Gprop` | `string` | `false` | `` | `` | `` |
| `date` | `Date` | `string` | `false` | `` | `` | `` |
| `csv` | `Csv` | `string` | `false` | `` | `` | `` |
| `include_low_search_volume` | `IncludeLowSearchVolume` | `string` | `false` | `` | `` | `` |
| `no_cache` | `NoCache` | `string` | `false` | `false` | `` | `` |

## `google_videos`

- SDK: `Google.Videos(ctx, dataify.GoogleVideosRequest{...})`
- Service kind: `serp`
- Engine: `google_videos`

| Parameter | Go Field | Type | Required | Default | Min | Max |
| --- | --- | --- | --- | --- | --- | --- |
| `q` | `Q` | `string` | `true` | `` | `` | `` |
| `json` | `JSON` | `string` | `true` | `1` | `` | `` |
| `google_domain` | `GoogleDomain` | `string` | `false` | `google.com` | `` | `` |
| `gl` | `Gl` | `string` | `false` | `` | `` | `` |
| `hl` | `Hl` | `string` | `false` | `` | `` | `` |
| `location` | `Location` | `string` | `false` | `` | `` | `` |
| `uule` | `Uule` | `string` | `false` | `` | `` | `` |
| `start` | `Start` | `string` | `false` | `` | `` | `` |
| `tbs` | `Tbs` | `string` | `false` | `` | `` | `` |
| `no_cache` | `NoCache` | `string` | `false` | `false` | `` | `` |
| `lr` | `Lr` | `string` | `false` | `` | `` | `` |
| `safe` | `Safe` | `string` | `false` | `` | `` | `` |
| `nfpr` | `Nfpr` | `string` | `false` | `` | `` | `` |
| `filter` | `Filter` | `string` | `false` | `1` | `` | `` |

## `indeed_companies_info`

- SDK: `Indeed.CompaniesInfo(ctx, dataify.IndeedCompaniesInfoRequest{...})`
- Service kind: `scraper_builder`
- Spider name: `indeed.com`
- Default spider ID: `indeed_companies-info_by-company-list-url`

| Parameter | Go Field | Type | Required | Default | Min | Max |
| --- | --- | --- | --- | --- | --- | --- |
| `spider_id` | `SpiderID` | `string` | `true` | `indeed_companies-info_by-company-list-url` | `` | `` |
| `company_list_url` | `CompanyListURL` | `string` | `false` | `https://www.indeed.com/companies/browse-companies` | `` | `` |
| `keyword` | `Keyword` | `string` | `false` | `openai` | `` | `` |
| `industry` | `Industry` | `string` | `false` | `Accounting & Tax` | `` | `` |
| `state` | `State` | `string` | `false` | `Alabama - 60 companies` | `` | `` |
| `company_url` | `CompanyURL` | `string` | `false` | `https://www.indeed.com/cmp/Allstate-Insurance` | `` | `` |
| `file_name` | `FileName` | `string` | `false` | `{{TasksID}}` | `` | `` |

## `indeed_job_listings`

- SDK: `Indeed.JobListings(ctx, dataify.IndeedJobListingsRequest{...})`
- Service kind: `scraper_builder`
- Spider name: `indeed.com`
- Default spider ID: `indeed_job-listings_by-job-url`

| Parameter | Go Field | Type | Required | Default | Min | Max |
| --- | --- | --- | --- | --- | --- | --- |
| `job_url` | `JobURL` | `string` | `true` | `https://fr.indeed.com/viewjob?jk=55b3e5dfa0c2ff66` | `` | `` |
| `file_name` | `FileName` | `string` | `false` | `{{TasksID}}` | `` | `` |

## `instagram_comment`

- SDK: `Instagram.Comment(ctx, dataify.InstagramCommentRequest{...})`
- Service kind: `scraper_builder`
- Spider name: `instagram.com`
- Default spider ID: `ins_comment_by-posturl`

| Parameter | Go Field | Type | Required | Default | Min | Max |
| --- | --- | --- | --- | --- | --- | --- |
| `posturl` | `Posturl` | `string` | `true` | `https://www.instagram.com/cats_of_instagram/reel/C4GLo_eLO2e/` | `` | `` |
| `file_name` | `FileName` | `string` | `false` | `{{TasksID}}` | `` | `` |

## `instagram_profiles`

- SDK: `Instagram.Profiles(ctx, dataify.InstagramProfilesRequest{...})`
- Service kind: `scraper_builder`
- Spider name: `instagram.com`
- Default spider ID: `ins_profiles_by-username`

| Parameter | Go Field | Type | Required | Default | Min | Max |
| --- | --- | --- | --- | --- | --- | --- |
| `spider_id` | `SpiderID` | `string` | `true` | `ins_profiles_by-username` | `` | `` |
| `username` | `Username` | `string` | `false` | `zoobarcelona` | `` | `` |
| `profileurl` | `Profileurl` | `string` | `false` | `https://www.instagram.com/cats_of_world_/` | `` | `` |
| `file_name` | `FileName` | `string` | `false` | `{{TasksID}}` | `` | `` |

## `instagram_reel`

- SDK: `Instagram.Reel(ctx, dataify.InstagramReelRequest{...})`
- Service kind: `scraper_builder`
- Spider name: `instagram.com`
- Default spider ID: `ins_reel_by-url`

| Parameter | Go Field | Type | Required | Default | Min | Max |
| --- | --- | --- | --- | --- | --- | --- |
| `spider_id` | `SpiderID` | `string` | `true` | `ins_reel_by-url` | `` | `` |
| `url` | `URL` | `string` | `true` | `https://www.instagram.com/reel/C5Rdyj_q7YN/` | `` | `` |
| `num_of_posts` | `NumOfPosts` | `string` | `false` | `10` | `` | `` |
| `posts_to_not_include` | `PostsToNotInclude` | `string` | `false` | `DP861NijuwE` | `` | `` |
| `start_date` | `StartDate` | `string` | `false` | `01-28-2025` | `` | `` |
| `end_date` | `EndDate` | `string` | `false` | `01-28-2026` | `` | `` |
| `file_name` | `FileName` | `string` | `false` | `{{TasksID}}` | `` | `` |

## `linkedin_company_information`

- SDK: `LinkedIn.CompanyInformation(ctx, dataify.LinkedInCompanyInformationRequest{...})`
- Service kind: `scraper_builder`
- Spider name: `linkedin.com`
- Default spider ID: `linkedin_company_information_by-url`

| Parameter | Go Field | Type | Required | Default | Min | Max |
| --- | --- | --- | --- | --- | --- | --- |
| `url` | `URL` | `string` | `true` | `https://www.linkedin.com/company/dynamo-software` | `` | `` |
| `file_name` | `FileName` | `string` | `false` | `{{TasksID}}` | `` | `` |

## `linkedin_job_listings_information`

- SDK: `LinkedIn.JobListingsInformation(ctx, dataify.LinkedInJobListingsInformationRequest{...})`
- Service kind: `scraper_builder`
- Spider name: `linkedin.com`
- Default spider ID: `linkedin_job_listings_information_by-job-listing-url`

| Parameter | Go Field | Type | Required | Default | Min | Max |
| --- | --- | --- | --- | --- | --- | --- |
| `spider_id` | `SpiderID` | `string` | `true` | `linkedin_job_listings_information_by-job-listing-url` | `` | `` |
| `job_listing_url` | `JobListingURL` | `string` | `false` | `https://www.linkedin.com/jobs/reddit-inc.-jobs-worldwide?f_C=150573` | `` | `` |
| `job_url` | `JobURL` | `string` | `false` | `https://www.linkedin.com/jobs/view/senior-client-partner-large-customer-sales-gaming-at-reddit-inc-4303761033?position=10&pageNum=0&refId=kHRQtl6Ws14VG9y3UloI4w%3D%3D&trackingId=%2Fb2esqHHEjp1FoEkC8PfuQ%3D%3D` | `` | `` |
| `location` | `Location` | `string` | `false` | `New York` | `` | `` |
| `keyword` | `Keyword` | `string` | `false` | `product manager` | `` | `` |
| `file_name` | `FileName` | `string` | `false` | `{{TasksID}}` | `` | `` |

## `reddit_comment`

- SDK: `Reddit.Comment(ctx, dataify.RedditCommentRequest{...})`
- Service kind: `scraper_builder`
- Spider name: `reddit.com`
- Default spider ID: `reddit_comment_by-url`

| Parameter | Go Field | Type | Required | Default | Min | Max |
| --- | --- | --- | --- | --- | --- | --- |
| `url` | `URL` | `string` | `true` | `https://www.reddit.com/r/datascience/comments/1cmnf0m/comment/l32204i/?utm_source=share&utm_medium=web3x&utm_name=web3xcss&utm_term=1&utm_content=share_button` | `` | `` |
| `days_back` | `DaysBack` | `string` | `false` | `10` | `` | `` |
| `comment_limit` | `CommentLimit` | `string` | `false` | `5` | `` | `` |
| `file_name` | `FileName` | `string` | `false` | `{{TasksID}}` | `` | `` |

## `reddit_posts`

- SDK: `Reddit.Posts(ctx, dataify.RedditPostsRequest{...})`
- Service kind: `scraper_builder`
- Spider name: `reddit.com`
- Default spider ID: `reddit_posts_by-url`

| Parameter | Go Field | Type | Required | Default | Min | Max |
| --- | --- | --- | --- | --- | --- | --- |
| `spider_id` | `SpiderID` | `string` | `true` | `reddit_posts_by-url` | `` | `` |
| `url` | `URL` | `string` | `false` | `https://www.reddit.com/r/battlefield2042/comments/1cmqs1d/official_update_on_the_next_battlefield_game/` | `` | `` |
| `keyword` | `Keyword` | `string` | `false` | `datascience` | `` | `` |
| `num_of_posts` | `NumOfPosts` | `string` | `false` | `10` | `` | `` |
| `sort_by` | `SortBy` | `string` | `false` | `Rising` | `` | `` |
| `sort_by_time` | `SortByTime` | `string` | `false` | `Now` | `` | `` |
| `file_name` | `FileName` | `string` | `false` | `{{TasksID}}` | `` | `` |

## `request_web_unlocker`

- SDK: `WebUnlocker.Request(ctx, dataify.RequestWebUnlockerRequest{...})`
- Service kind: `web_unlocker`

| Parameter | Go Field | Type | Required | Default | Min | Max |
| --- | --- | --- | --- | --- | --- | --- |
| `url` | `URL` | `string` | `true` | `` | `` | `` |
| `type` | `Type` | `string` | `false` | `html` | `` | `` |
| `js_render` | `JsRender` | `string` | `false` | `True` | `` | `` |
| `block_resources` | `BlockResources` | `string` | `false` | `` | `` | `` |
| `clean_content` | `CleanContent` | `string` | `false` | `` | `` | `` |
| `country` | `Country` | `string` | `false` | `us` | `` | `` |
| `headers` | `Headers` | `string` | `false` | `` | `` | `` |
| `cookies` | `Cookies` | `string` | `false` | `` | `` | `` |
| `wait` | `Wait` | `string` | `false` | `` | `` | `` |
| `wait_for` | `WaitFor` | `string` | `false` | `` | `` | `` |
| `follow_redirect` | `FollowRedirect` | `string` | `false` | `True` | `` | `` |

## `tiktok_comment`

- SDK: `TikTok.Comment(ctx, dataify.TikTokCommentRequest{...})`
- Service kind: `scraper_builder`
- Spider name: `tiktok.com`
- Default spider ID: `tiktok_comment_by-url`

| Parameter | Go Field | Type | Required | Default | Min | Max |
| --- | --- | --- | --- | --- | --- | --- |
| `url` | `URL` | `string` | `true` | `https://www.tiktok.com/@heymrcat/video/7216019547806092550` | `` | `` |
| `page_turning` | `PageTurning` | `string` | `true` | `1` | `` | `` |
| `file_name` | `FileName` | `string` | `false` | `{{TasksID}}` | `` | `` |

## `tiktok_posts`

- SDK: `TikTok.Posts(ctx, dataify.TikTokPostsRequest{...})`
- Service kind: `scraper_builder`
- Spider name: `tiktok.com`
- Default spider ID: `tiktok_posts_by-listurl`

| Parameter | Go Field | Type | Required | Default | Min | Max |
| --- | --- | --- | --- | --- | --- | --- |
| `url` | `URL` | `string` | `true` | `https://www.tiktok.com/discover/dog` | `` | `` |
| `num_of_posts` | `NumOfPosts` | `string` | `false` | `5` | `` | `` |
| `file_name` | `FileName` | `string` | `false` | `{{TasksID}}` | `` | `` |

## `tiktok_profiles`

- SDK: `TikTok.Profiles(ctx, dataify.TikTokProfilesRequest{...})`
- Service kind: `scraper_builder`
- Spider name: `tiktok.com`
- Default spider ID: `tiktok_profiles_by-url`

| Parameter | Go Field | Type | Required | Default | Min | Max |
| --- | --- | --- | --- | --- | --- | --- |
| `spider_id` | `SpiderID` | `string` | `true` | `tiktok_profiles_by-url` | `` | `` |
| `url` | `URL` | `string` | `false` | `https://www.tiktok.com/@fofimdmell` | `` | `` |
| `search_url` | `SearchURL` | `string` | `false` | `https://www.tiktok.com/explore?lang=en` | `` | `` |
| `country` | `Country` | `string` | `false` | `us` | `` | `` |
| `page_turning` | `PageTurning` | `string` | `false` | `1` | `` | `` |
| `file_name` | `FileName` | `string` | `false` | `{{TasksID}}` | `` | `` |

## `tiktok_shop`

- SDK: `TikTok.Shop(ctx, dataify.TikTokShopRequest{...})`
- Service kind: `scraper_builder`
- Spider name: `tiktok.com`
- Default spider ID: `tiktok_shop_by-url`

| Parameter | Go Field | Type | Required | Default | Min | Max |
| --- | --- | --- | --- | --- | --- | --- |
| `url` | `URL` | `string` | `true` | `https://www.tiktok.com/shop/pdp/long-sleeve-crew-neck-tee-3-pack-by-galaxy-by-harvic-cotton-blend/1729461570693075200` | `` | `` |
| `file_name` | `FileName` | `string` | `false` | `{{TasksID}}` | `` | `` |

## `twitter_post`

- SDK: `Twitter.Post(ctx, dataify.TwitterPostRequest{...})`
- Service kind: `scraper_builder`
- Spider name: `x.com`
- Default spider ID: `twitter_post_by-profileurl`

| Parameter | Go Field | Type | Required | Default | Min | Max |
| --- | --- | --- | --- | --- | --- | --- |
| `url` | `URL` | `string` | `true` | `https://x.com/elonmusk` | `` | `` |
| `file_name` | `FileName` | `string` | `false` | `{{TasksID}}` | `` | `` |

## `twitter_profile`

- SDK: `Twitter.Profile(ctx, dataify.TwitterProfileRequest{...})`
- Service kind: `scraper_builder`
- Spider name: `x.com`
- Default spider ID: `twitter_profile_by-profileurl`

| Parameter | Go Field | Type | Required | Default | Min | Max |
| --- | --- | --- | --- | --- | --- | --- |
| `spider_id` | `SpiderID` | `string` | `true` | `twitter_profile_by-profileurl` | `` | `` |
| `url` | `URL` | `string` | `false` | `https://x.com/elonmusk` | `` | `` |
| `user_name` | `UserName` | `string` | `false` | `elonmusk` | `` | `` |
| `file_name` | `FileName` | `string` | `false` | `{{TasksID}}` | `` | `` |

## `walmart_product`

- SDK: `Walmart.Product(ctx, dataify.WalmartProductRequest{...})`
- Service kind: `scraper_builder`
- Spider name: `walmart.com`
- Default spider ID: `walmart_product_by-url`

| Parameter | Go Field | Type | Required | Default | Min | Max |
| --- | --- | --- | --- | --- | --- | --- |
| `spider_id` | `SpiderID` | `string` | `true` | `walmart_product_by-url` | `` | `` |
| `url` | `URL` | `string` | `false` | `https://www.walmart.com/ip/HI-CHEW-Stand-Up-Pouch-Getaway-Mix-11-65oz/12284762931?athAsset=eyJhdGhjcGlkIjoiMTIyODQ3NjI5MzEiLCJhdGhzdGlkIjoiQ1MwNTV+Q1MwMDR+Q1MwOTgiLCJhdGhlZSI6eyJhIjoyNy44NCwiYiI6Mjk1MS40MSwidyI6MC4wMDk0MjcxMjc3OTA0NzcxMjMsImwiOjAuNX0sImF0aHBvc2IiOiI4IiwiYXRoYW5jaWQiOiIxMDE2NDUwNzU1IiwiYXRocmsiOjAuMH0%3D&athena=true&adsRedirect=true` | `` | `` |
| `category_url` | `CategoryURL` | `string` | `false` | `https://www.walmart.com/shop/deals/food/` | `` | `` |
| `sku` | `Sku` | `string` | `false` | `439179861` | `` | `` |
| `keyword` | `Keyword` | `string` | `false` | `leggins` | `` | `` |
| `domain` | `Domain` | `string` | `false` | `https://www.walmart.com/` | `` | `` |
| `all_variations` | `AllVariations` | `string` | `false` | `` | `` | `` |
| `page_turning` | `PageTurning` | `string` | `false` | `` | `` | `` |
| `file_name` | `FileName` | `string` | `false` | `{{TasksID}}` | `` | `` |

## `yandex_search`

- SDK: `Yandex.Search(ctx, dataify.YandexSearchRequest{...})`
- Service kind: `serp`
- Engine: `yandex`

| Parameter | Go Field | Type | Required | Default | Min | Max |
| --- | --- | --- | --- | --- | --- | --- |
| `text` | `Text` | `string` | `true` | `` | `` | `` |
| `json` | `JSON` | `string` | `true` | `1` | `` | `` |
| `yandex_domain` | `YandexDomain` | `string` | `false` | `yandex.com` | `` | `` |
| `lang` | `Lang` | `string` | `false` | `en` | `` | `` |
| `lr` | `Lr` | `string` | `false` | `` | `` | `` |
| `p` | `P` | `string` | `false` | `0` | `` | `` |
| `family_mode` | `FamilyMode` | `string` | `false` | `1` | `` | `` |
| `fix_typo` | `FixTypo` | `string` | `false` | `true` | `` | `` |
| `groups_on_page` | `GroupsOnPage` | `string` | `false` | `10` | `` | `` |
| `no_cache` | `NoCache` | `string` | `false` | `false` | `` | `` |

## `youtube_audio`

- SDK: `YouTube.Audio(ctx, dataify.YouTubeAudioRequest{...})`
- Service kind: `scraper_builder`
- Spider name: `youtube.com`
- Default spider ID: `youtube_audio_by-url`

| Parameter | Go Field | Type | Required | Default | Min | Max |
| --- | --- | --- | --- | --- | --- | --- |
| `url` | `URL` | `string` | `true` | `https://www.youtube.com/watch?v=_SdpvpvVrLY` | `` | `` |
| `subtitles_language` | `SubtitlesLanguage` | `string` | `false` | `ab` | `` | `` |
| `selected_only` | `SelectedOnly` | `string` | `false` | `false` | `` | `` |
| `kilohertz` | `Kilohertz` | `string` | `false` | `<=48000` | `` | `` |
| `is_subtitles` | `IsSubtitles` | `string` | `false` | `false` | `` | `` |
| `audio_format` | `AudioFormat` | `string` | `false` | `opus` | `` | `` |
| `bitrate` | `Bitrate` | `string` | `false` | `<=320` | `` | `` |
| `file_name` | `FileName` | `string` | `false` | `{{TasksID}}` | `` | `` |

## `youtube_comment`

- SDK: `YouTube.Comment(ctx, dataify.YouTubeCommentRequest{...})`
- Service kind: `scraper_builder`
- Spider name: `youtube.com`
- Default spider ID: `youtube_comment_by-id`

| Parameter | Go Field | Type | Required | Default | Min | Max |
| --- | --- | --- | --- | --- | --- | --- |
| `video_id` | `VideoID` | `string` | `true` | `8RePenzQH80` | `` | `` |
| `load_replies` | `LoadReplies` | `string` | `true` | `10` | `` | `` |
| `num_of_comments` | `NumOfComments` | `string` | `true` | `Top comments` | `` | `` |
| `file_name` | `FileName` | `string` | `false` | `{{TasksID}}` | `` | `` |

## `youtube_product`

- SDK: `YouTube.Product(ctx, dataify.YouTubeProductRequest{...})`
- Service kind: `scraper_builder`
- Spider name: `youtube.com`
- Default spider ID: `youtube_product_by-id`

| Parameter | Go Field | Type | Required | Default | Min | Max |
| --- | --- | --- | --- | --- | --- | --- |
| `video_id` | `VideoID` | `string` | `true` | `8RePenzQH80` | `` | `` |
| `subtitles_language` | `SubtitlesLanguage` | `string` | `false` | `ab` | `` | `` |
| `subtitles_type` | `SubtitlesType` | `string` | `false` | `auto_generated` | `` | `` |
| `selected_only` | `SelectedOnly` | `string` | `false` | `false` | `` | `` |
| `file_name` | `FileName` | `string` | `false` | `{{TasksID}}` | `` | `` |

## `youtube_profiles`

- SDK: `YouTube.Profiles(ctx, dataify.YouTubeProfilesRequest{...})`
- Service kind: `scraper_builder`
- Spider name: `youtube.com`
- Default spider ID: `youtube_profiles_by-keyword`

| Parameter | Go Field | Type | Required | Default | Min | Max |
| --- | --- | --- | --- | --- | --- | --- |
| `spider_id` | `SpiderID` | `string` | `true` | `youtube_profiles_by-keyword` | `` | `` |
| `keyword` | `Keyword` | `string` | `false` | `MrBeast` | `` | `` |
| `page_turning` | `PageTurning` | `string` | `false` | `1` | `` | `` |
| `url` | `URL` | `string` | `false` | `https://www.youtube.com/@mrbeast` | `` | `` |
| `file_name` | `FileName` | `string` | `false` | `{{TasksID}}` | `` | `` |

## `youtube_transcript`

- SDK: `YouTube.Transcript(ctx, dataify.YouTubeTranscriptRequest{...})`
- Service kind: `scraper_builder`
- Spider name: `youtube.com`
- Default spider ID: `youtube_transcript_by-id`

| Parameter | Go Field | Type | Required | Default | Min | Max |
| --- | --- | --- | --- | --- | --- | --- |
| `video_id` | `VideoID` | `string` | `true` | `8RePenzQH80` | `` | `` |
| `subtitles_language` | `SubtitlesLanguage` | `string` | `false` | `ab` | `` | `` |
| `subtitles_type` | `SubtitlesType` | `string` | `false` | `auto_generated` | `` | `` |
| `selected_only` | `SelectedOnly` | `string` | `false` | `false` | `` | `` |
| `file_name` | `FileName` | `string` | `false` | `{{TasksID}}` | `` | `` |

## `youtube_video`

- SDK: `YouTube.Video(ctx, dataify.YouTubeVideoRequest{...})`
- Service kind: `scraper_builder`
- Spider name: `youtube.com`
- Default spider ID: `youtube_video_by-url`

| Parameter | Go Field | Type | Required | Default | Min | Max |
| --- | --- | --- | --- | --- | --- | --- |
| `url` | `URL` | `string` | `true` | `https://www.youtube.com/watch?v=_SdpvpvVrLY` | `` | `` |
| `subtitles_language` | `SubtitlesLanguage` | `string` | `false` | `ab` | `` | `` |
| `selected_only` | `SelectedOnly` | `string` | `false` | `false` | `` | `` |
| `resolution` | `Resolution` | `string` | `false` | `<=360p` | `` | `` |
| `video_codec` | `VideoCodec` | `string` | `false` | `vp9` | `` | `` |
| `audio_format` | `AudioFormat` | `string` | `false` | `opus` | `` | `` |
| `bitrate` | `Bitrate` | `string` | `false` | `<=320` | `` | `` |
| `file_name` | `FileName` | `string` | `false` | `{{TasksID}}` | `` | `` |

## `youtube_video_post`

- SDK: `YouTube.VideoPost(ctx, dataify.YouTubeVideoPostRequest{...})`
- Service kind: `scraper_builder`
- Spider name: `youtube.com`
- Default spider ID: `youtube_video-post_by-url`

| Parameter | Go Field | Type | Required | Default | Min | Max |
| --- | --- | --- | --- | --- | --- | --- |
| `spider_id` | `SpiderID` | `string` | `true` | `youtube_video-post_by-url` | `` | `` |
| `url` | `URL` | `string` | `false` | `https://www.youtube.com/@stephcurry/videos` | `` | `` |
| `order_by` | `OrderBy` | `string` | `false` | `Latest` | `` | `` |
| `start_index` | `StartIndex` | `string` | `false` | `1` | `` | `` |
| `num_of_posts` | `NumOfPosts` | `string` | `false` | `5` | `` | `` |
| `keyword_search` | `KeywordSearch` | `string` | `false` | `popular music` | `` | `` |
| `features` | `Features` | `string` | `false` | `All` | `` | `` |
| `type` | `Type` | `string` | `false` | `Video` | `` | `` |
| `duration` | `Duration` | `string` | `false` | `Under 3 minutes` | `` | `` |
| `upload_date` | `UploadDate` | `string` | `false` | `上一小时` | `` | `` |
| `hashtag` | `Hashtag` | `string` | `false` | `shopping` | `` | `` |
| `keyword` | `Keyword` | `string` | `false` | `top videos` | `` | `` |
| `all_tabs` | `AllTabs` | `string` | `false` | `true` | `` | `` |
| `file_name` | `FileName` | `string` | `false` | `{{TasksID}}` | `` | `` |

## `zillow_product`

- SDK: `Zillow.Product(ctx, dataify.ZillowProductRequest{...})`
- Service kind: `scraper_builder`
- Spider name: `zillow.com`
- Default spider ID: `zillow_product_by-filter`

| Parameter | Go Field | Type | Required | Default | Min | Max |
| --- | --- | --- | --- | --- | --- | --- |
| `keywords-location` | `KeywordsLocation` | `string` | `true` | `South Bend` | `` | `` |
| `listingCategory` | `Listingcategory` | `string` | `false` | `For Rent` | `` | `` |
| `HomeType` | `Hometype` | `string` | `false` | `Houses` | `` | `` |
| `days_on_zillow` | `DaysOnZillow` | `string` | `false` | `Any` | `` | `` |
| `maximum` | `Maximum` | `string` | `false` | `10` | `` | `` |
| `file_name` | `FileName` | `string` | `false` | `{{TasksID}}` | `` | `` |

