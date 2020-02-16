package adsinsights

import (
	"net/http"

	"github.com/mazti/facebook-go-business-sdk/sdk"
	"github.com/mazti/facebook-go-business-sdk/sdk/adsactionstats"
)

const (
	endpoint = "/insights"
)

// Entity AdsInsight https://developers.facebook.com/docs/marketing-api/reference/ad-account/insights
type Entity struct {
	request *sdk.APIRequest

	AccountCurrency                           string                  `json:"account_currency"`
	AccountId                                 string                  `json:"account_id"`
	AccountName                               string                  `json:"account_name"`
	ActionValues                              []adsactionstats.Entity `json:"action_values"`
	Actions                                   []adsactionstats.Entity `json:"actions"`
	ActionsPerImpression                      string                  `json:"actions_per_impression"`
	ActionsResults                            adsactionstats.Entity   `json:"actions_results"`
	ActivityRecency                           string                  `json:"activity_recency"`
	AdBidType                                 string                  `json:"ad_bid_type"`
	AdBidValue                                string                  `json:"ad_bid_value"`
	AdClickActions                            []adsactionstats.Entity `json:"ad_click_actions"`
	AdDelivery                                string                  `json:"ad_delivery"`
	AdFormatAsset                             string                  `json:"ad_format_asset"`
	AdId                                      string                  `json:"ad_id"`
	AdImpressionActions                       []adsactionstats.Entity `json:"ad_impression_actions"`
	AdName                                    string                  `json:"ad_name"`
	AdsetBidType                              string                  `json:"adset_bid_type"`
	AdsetBidValue                             string                  `json:"adset_bid_value"`
	AdsetBudgetType                           string                  `json:"adset_budget_type"`
	AdsetBudgetValue                          string                  `json:"adset_budget_value"`
	AdsetDelivery                             string                  `json:"adset_delivery"`
	AdsetEnd                                  string                  `json:"adset_end"`
	AdsetId                                   string                  `json:"adset_id"`
	AdsetName                                 string                  `json:"adset_name"`
	AdsetStart                                string                  `json:"adset_start"`
	Age                                       string                  `json:"age"`
	AgeTargeting                              string                  `json:"age_targeting"`
	AmountInCatalogCurrency                   []adsactionstats.Entity `json:"amount_in_catalog_currency"`
	AppStoreClicks                            string                  `json:"app_store_clicks"`
	AttentionEventsPerImpression              string                  `json:"attention_events_per_impression"`
	AttentionEventsUnqPerReach                string                  `json:"attention_events_unq_per_reach"`
	AuctionBid                                string                  `json:"auction_bid"`
	AuctionCompetitiveness                    string                  `json:"auction_competitiveness"`
	AuctionMaxCompetitorBid                   string                  `json:"auction_max_competitor_bid"`
	BodyAsset                                 interface{}             `json:"body_asset"`
	BuyingType                                string                  `json:"buying_type"`
	CallToActionAsset                         interface{}             `json:"call_to_action_asset"`
	CallToActionClicks                        string                  `json:"call_to_action_clicks"`
	CampaignDelivery                          string                  `json:"campaign_delivery"`
	CampaignEnd                               string                  `json:"campaign_end"`
	CampaignId                                string                  `json:"campaign_id"`
	CampaignName                              string                  `json:"campaign_name"`
	CampaignStart                             string                  `json:"campaign_start"`
	CancelSubscriptionActions                 []adsactionstats.Entity `json:"cancel_subscription_actions"`
	CanvasAvgViewPercent                      string                  `json:"canvas_avg_view_percent"`
	CanvasAvgViewTime                         string                  `json:"canvas_avg_view_time"`
	CardViews                                 string                  `json:"card_views"`
	CatalogSegmentActions                     []adsactionstats.Entity `json:"catalog_segment_actions"`
	CatalogSegmentValueInCatalogCurrency      []adsactionstats.Entity `json:"catalog_segment_value_in_catalog_currency"`
	CatalogSegmentValueMobilePurchaseRoas     []adsactionstats.Entity `json:"catalog_segment_value_mobile_purchase_roas"`
	CatalogSegmentValueWebsitePurchaseRoas    []adsactionstats.Entity `json:"catalog_segment_value_website_purchase_roas"`
	Clicks                                    string                  `json:"clicks"`
	ConditionalTimeSpentMsOver10sActions      []adsactionstats.Entity `json:"conditional_time_spent_ms_over_10s_actions"`
	ConditionalTimeSpentMsOver15sActions      []adsactionstats.Entity `json:"conditional_time_spent_ms_over_15s_actions"`
	ConditionalTimeSpentMsOver2sActions       []adsactionstats.Entity `json:"conditional_time_spent_ms_over_2s_actions"`
	ConditionalTimeSpentMsOver3sActions       []adsactionstats.Entity `json:"conditional_time_spent_ms_over_3s_actions"`
	ConditionalTimeSpentMsOver6sActions       []adsactionstats.Entity `json:"conditional_time_spent_ms_over_6s_actions"`
	ContactActions                            []adsactionstats.Entity `json:"contact_actions"`
	ContactValue                              []adsactionstats.Entity `json:"contact_value"`
	ConversionRateRanking                     string                  `json:"conversion_rate_ranking"`
	ConversionValues                          []adsactionstats.Entity `json:"conversion_values"`
	Conversions                               []adsactionstats.Entity `json:"conversions"`
	CostPer15SecVideoView                     []adsactionstats.Entity `json:"cost_per_15_sec_video_view"`
	CostPer2SecContinuousVideoView            []adsactionstats.Entity `json:"cost_per_2_sec_continuous_video_view"`
	CostPerActionResult                       adsactionstats.Entity   `json:"cost_per_action_result"`
	CostPerActionType                         []adsactionstats.Entity `json:"cost_per_action_type"`
	CostPerAdClick                            []adsactionstats.Entity `json:"cost_per_ad_click"`
	CostPerCompletedVideoView                 []adsactionstats.Entity `json:"cost_per_completed_video_view"`
	CostPerContact                            []adsactionstats.Entity `json:"cost_per_contact"`
	CostPerConversion                         []adsactionstats.Entity `json:"cost_per_conversion"`
	CostPerCustomizeProduct                   []adsactionstats.Entity `json:"cost_per_customize_product"`
	CostPerDdaCountbyConvs                    string                  `json:"cost_per_dda_countby_convs"`
	CostPerDonate                             []adsactionstats.Entity `json:"cost_per_donate"`
	CostPerDwell                              string                  `json:"cost_per_dwell"`
	CostPerDwell3Sec                          string                  `json:"cost_per_dwell_3_sec"`
	CostPerDwell5Sec                          string                  `json:"cost_per_dwell_5_sec"`
	CostPerDwell7Sec                          string                  `json:"cost_per_dwell_7_sec"`
	CostPerEstimatedAdRecallers               string                  `json:"cost_per_estimated_ad_recallers"`
	CostPerFindLocation                       []adsactionstats.Entity `json:"cost_per_find_location"`
	CostPerInlineLinkClick                    string                  `json:"cost_per_inline_link_click"`
	CostPerInlinePostEngagement               string                  `json:"cost_per_inline_post_engagement"`
	CostPerOneThousandAdImpression            []adsactionstats.Entity `json:"cost_per_one_thousand_ad_impression"`
	CostPerOutboundClick                      []adsactionstats.Entity `json:"cost_per_outbound_click"`
	CostPerSchedule                           []adsactionstats.Entity `json:"cost_per_schedule"`
	CostPerStartTrial                         []adsactionstats.Entity `json:"cost_per_start_trial"`
	CostPerSubmitApplication                  []adsactionstats.Entity `json:"cost_per_submit_application"`
	CostPerSubscribe                          []adsactionstats.Entity `json:"cost_per_subscribe"`
	CostPerThruplay                           []adsactionstats.Entity `json:"cost_per_thruplay"`
	CostPerTotalAction                        string                  `json:"cost_per_total_action"`
	CostPerUniqueActionType                   []adsactionstats.Entity `json:"cost_per_unique_action_type"`
	CostPerUniqueClick                        string                  `json:"cost_per_unique_click"`
	CostPerUniqueConversion                   []adsactionstats.Entity `json:"cost_per_unique_conversion"`
	CostPerUniqueInlineLinkClick              string                  `json:"cost_per_unique_inline_link_click"`
	CostPerUniqueOutboundClick                []adsactionstats.Entity `json:"cost_per_unique_outbound_click"`
	Country                                   string                  `json:"country"`
	Cpc                                       string                  `json:"cpc"`
	Cpm                                       string                  `json:"cpm"`
	Cpp                                       string                  `json:"cpp"`
	CreatedTime                               string                  `json:"created_time"`
	CreativeFingerprint                       string                  `json:"creative_fingerprint"`
	Ctr                                       string                  `json:"ctr"`
	CustomizeProductActions                   []adsactionstats.Entity `json:"customize_product_actions"`
	CustomizeProductValue                     []adsactionstats.Entity `json:"customize_product_value"`
	DateStart                                 string                  `json:"date_start"`
	DateStop                                  string                  `json:"date_stop"`
	DdaCountbyConvs                           string                  `json:"dda_countby_convs"`
	Deduping1stSourceRatio                    string                  `json:"deduping_1st_source_ratio"`
	Deduping2ndSourceRatio                    string                  `json:"deduping_2nd_source_ratio"`
	Deduping3rdSourceRatio                    string                  `json:"deduping_3rd_source_ratio"`
	DedupingRatio                             string                  `json:"deduping_ratio"`
	DeeplinkClicks                            string                  `json:"deeplink_clicks"`
	DescriptionAsset                          interface{}             `json:"description_asset"`
	DevicePlatform                            string                  `json:"device_platform"`
	Dma                                       string                  `json:"dma"`
	DonateActions                             []adsactionstats.Entity `json:"donate_actions"`
	DonateValue                               []adsactionstats.Entity `json:"donate_value"`
	Dwell3Sec                                 string                  `json:"dwell_3_sec"`
	Dwell5Sec                                 string                  `json:"dwell_5_sec"`
	Dwell7Sec                                 string                  `json:"dwell_7_sec"`
	DwellRate                                 string                  `json:"dwell_rate"`
	EarnedImpression                          string                  `json:"earned_impression"`
	EngagementRateRanking                     string                  `json:"engagement_rate_ranking"`
	EstimatedAdRecallRate                     string                  `json:"estimated_ad_recall_rate"`
	EstimatedAdRecallRateLowerBound           string                  `json:"estimated_ad_recall_rate_lower_bound"`
	EstimatedAdRecallRateUpperBound           string                  `json:"estimated_ad_recall_rate_upper_bound"`
	EstimatedAdRecallers                      string                  `json:"estimated_ad_recallers"`
	EstimatedAdRecallersLowerBound            string                  `json:"estimated_ad_recallers_lower_bound"`
	EstimatedAdRecallersUpperBound            string                  `json:"estimated_ad_recallers_upper_bound"`
	FindLocationActions                       []adsactionstats.Entity `json:"find_location_actions"`
	FindLocationValue                         []adsactionstats.Entity `json:"find_location_value"`
	Frequency                                 string                  `json:"frequency"`
	FrequencyValue                            string                  `json:"frequency_value"`
	FullViewImpressions                       string                  `json:"full_view_impressions"`
	FullViewReach                             string                  `json:"full_view_reach"`
	Gender                                    string                  `json:"gender"`
	GenderTargeting                           string                  `json:"gender_targeting"`
	HourlyStatsAggregatedByAdvertiserTimeZone string                  `json:"hourly_stats_aggregated_by_advertiser_time_zone"`
	HourlyStatsAggregatedByAudienceTimeZone   string                  `json:"hourly_stats_aggregated_by_audience_time_zone"`
	ImageAsset                                interface{}             `json:"image_asset"`
	ImpressionDevice                          string                  `json:"impression_device"`
	Impressions                               string                  `json:"impressions"`
	ImpressionsAutoRefresh                    string                  `json:"impressions_auto_refresh"`
	ImpressionsGross                          string                  `json:"impressions_gross"`
	InlineLinkClickCtr                        string                  `json:"inline_link_click_ctr"`
	InlineLinkClicks                          string                  `json:"inline_link_clicks"`
	InlinePostEngagement                      string                  `json:"inline_post_engagement"`
	InstantExperienceClicksToOpen             string                  `json:"instant_experience_clicks_to_open"`
	InstantExperienceClicksToStart            string                  `json:"instant_experience_clicks_to_start"`
	InstantExperienceOutboundClicks           string                  `json:"instant_experience_outbound_clicks"`
	InteractiveComponentTap                   []adsactionstats.Entity `json:"interactive_component_tap"`
	Labels                                    string                  `json:"labels"`
	LinkUrlAsset                              interface{}             `json:"link_url_asset"`
	Location                                  string                  `json:"location"`
	MediaAsset                                interface{}             `json:"media_asset"`
	MobileAppPurchaseRoas                     []adsactionstats.Entity `json:"mobile_app_purchase_roas"`
	NewsfeedAvgPosition                       string                  `json:"newsfeed_avg_position"`
	NewsfeedClicks                            string                  `json:"newsfeed_clicks"`
	NewsfeedImpressions                       string                  `json:"newsfeed_impressions"`
	Objective                                 string                  `json:"objective"`
	OptimizationGoal                          string                  `json:"optimization_goal"`
	OutboundClicks                            []adsactionstats.Entity `json:"outbound_clicks"`
	OutboundClicksCtr                         []adsactionstats.Entity `json:"outbound_clicks_ctr"`
	PerformanceIndicator                      string                  `json:"performance_indicator"`
	PlacePageId                               string                  `json:"place_page_id"`
	PlacePageName                             string                  `json:"place_page_name"`
	Placement                                 string                  `json:"placement"`
	PlatformPosition                          string                  `json:"platform_position"`
	ProductId                                 string                  `json:"product_id"`
	PublisherPlatform                         string                  `json:"publisher_platform"`
	PurchaseRoas                              []adsactionstats.Entity `json:"purchase_roas"`
	QualityRanking                            string                  `json:"quality_ranking"`
	QualityScoreEctr                          string                  `json:"quality_score_ectr"`
	QualityScoreEcvr                          string                  `json:"quality_score_ecvr"`
	QualityScoreOrganic                       string                  `json:"quality_score_organic"`
	Reach                                     string                  `json:"reach"`
	RecurringSubscriptionPaymentActions       []adsactionstats.Entity `json:"recurring_subscription_payment_actions"`
	Region                                    string                  `json:"region"`
	RuleAsset                                 interface{}             `json:"rule_asset"`
	ScheduleActions                           []adsactionstats.Entity `json:"schedule_actions"`
	ScheduleValue                             []adsactionstats.Entity `json:"schedule_value"`
	SocialSpend                               string                  `json:"social_spend"`
	Spend                                     string                  `json:"spend"`
	StartTrialActions                         []adsactionstats.Entity `json:"start_trial_actions"`
	StartTrialValue                           []adsactionstats.Entity `json:"start_trial_value"`
	SubmitApplicationActions                  []adsactionstats.Entity `json:"submit_application_actions"`
	SubmitApplicationValue                    []adsactionstats.Entity `json:"submit_application_value"`
	SubscribeActions                          []adsactionstats.Entity `json:"subscribe_actions"`
	SubscribeValue                            []adsactionstats.Entity `json:"subscribe_value"`
	ThumbStops                                string                  `json:"thumb_stops"`
	TitleAsset                                interface{}             `json:"title_asset"`
	TodaySpend                                string                  `json:"today_spend"`
	TotalActionValue                          string                  `json:"total_action_value"`
	TotalActions                              string                  `json:"total_actions"`
	TotalUniqueActions                        string                  `json:"total_unique_actions"`
	UniqueActions                             []adsactionstats.Entity `json:"unique_actions"`
	UniqueClicks                              string                  `json:"unique_clicks"`
	UniqueConversions                         []adsactionstats.Entity `json:"unique_conversions"`
	UniqueCtr                                 string                  `json:"unique_ctr"`
	UniqueImpressions                         string                  `json:"unique_impressions"`
	UniqueInlineLinkClickCtr                  string                  `json:"unique_inline_link_click_ctr"`
	UniqueInlineLinkClicks                    string                  `json:"unique_inline_link_clicks"`
	UniqueLinkClicksCtr                       string                  `json:"unique_link_clicks_ctr"`
	UniqueOutboundClicks                      []adsactionstats.Entity `json:"unique_outbound_clicks"`
	UniqueOutboundClicksCtr                   []adsactionstats.Entity `json:"unique_outbound_clicks_ctr"`
	UniqueVideoContinuous2SecWatchedActions   []adsactionstats.Entity `json:"unique_video_continuous_2_sec_watched_actions"`
	UniqueVideoView15Sec                      []adsactionstats.Entity `json:"unique_video_view_15_sec"`
	UpdatedTime                               string                  `json:"updated_time"`
	Video15SecWatchedActions                  []adsactionstats.Entity `json:"video_15_sec_watched_actions"`
	Video30SecWatchedActions                  []adsactionstats.Entity `json:"video_30_sec_watched_actions"`
	VideoAsset                                interface{}             `json:"video_asset"`
	VideoAvgTimeWatchedActions                []adsactionstats.Entity `json:"video_avg_time_watched_actions"`
	VideoCompleteWatchedActions               []adsactionstats.Entity `json:"video_complete_watched_actions"`
	VideoCompletedViewOr15sPassedActions      []adsactionstats.Entity `json:"video_completed_view_or_15s_passed_actions"`
	VideoContinuous2SecWatchedActions         []adsactionstats.Entity `json:"video_continuous_2_sec_watched_actions"`
	VideoP100WatchedActions                   []adsactionstats.Entity `json:"video_p100_watched_actions"`
	VideoP25WatchedActions                    []adsactionstats.Entity `json:"video_p25_watched_actions"`
	VideoP50WatchedActions                    []adsactionstats.Entity `json:"video_p50_watched_actions"`
	VideoP75WatchedActions                    []adsactionstats.Entity `json:"video_p75_watched_actions"`
	VideoP95WatchedActions                    []adsactionstats.Entity `json:"video_p95_watched_actions"`
	VideoPlayActions                          []adsactionstats.Entity `json:"video_play_actions"`
	VideoPlayCurveActions                     []interface{}           `json:"video_play_curve_actions"`
	VideoPlayRetention0To15sActions           []interface{}           `json:"video_play_retention_0_to_15s_actions"`
	VideoPlayRetention20To60sActions          []interface{}           `json:"video_play_retention_20_to_60s_actions"`
	VideoPlayRetentionGraphActions            []interface{}           `json:"video_play_retention_graph_actions"`
	VideoThruplayWatchedActions               []adsactionstats.Entity `json:"video_thruplay_watched_actions"`
	VideoTimeWatchedActions                   []adsactionstats.Entity `json:"video_time_watched_actions"`
	WebsiteClicks                             string                  `json:"website_clicks"`
	WebsiteCtr                                []adsactionstats.Entity `json:"website_ctr"`
	WebsitePurchaseRoas                       []adsactionstats.Entity `json:"website_purchase_roas"`
	WishBid                                   string                  `json:"wish_bid"`
}

func CreateAPIRequestGet(id string, context *sdk.APIContext) *sdk.APIRequest {
	return sdk.NewAPIRequest(
		context,
		id,
		endpoint,
		http.MethodGet,
		sdk.ParamNames(params),
	)
}

func ParseResponse(rawResp sdk.APIResponse) (resp []Entity, err error) {
	request := rawResp.GetRequest()
	context := request.Context
	nodeList, ok := rawResp.(*sdk.APINodeList)
	if !ok {
		return nil, sdk.UnsupportedResponse
	}
	err = nodeList.Unmarshal(&resp)
	if err != nil {
		context.Log(err)
		return
	}
	for i := 0; i < len(resp); i++ {
		resp[i].SetRequest(request)
	}
	return
}
