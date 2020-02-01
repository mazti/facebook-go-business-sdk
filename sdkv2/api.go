package sdkv2

type FacebookAdsAPI struct {
	debug       bool
	showHeader  bool
	accessToken string
	appSecret   string
	appID       string
	locale      string
}

func (api FacebookAdsAPI) Version() string {
	return "v5.0"
}

func (api FacebookAdsAPI) Graph() string {
	return "https://graph.facebook.com"
}

func (api FacebookAdsAPI) GraphVideo() string {
	return "https://graph-video.facebook.com"
}

func NewFacebookAdsAPI(options ...func(*FacebookAdsAPI)) *FacebookAdsAPI {
	api := &FacebookAdsAPI{
		debug:      false,
		showHeader: false,
	}
	for _, option := range options {
		option(api)
	}
	return api
}

// Constructor options to use for setting value in default constructor

func Debug(debug bool) func(*FacebookAdsAPI) {
	return func(api *FacebookAdsAPI) {
		api.debug = debug
	}
}

func AccessToken(token string) func(*FacebookAdsAPI) {
	return func(api *FacebookAdsAPI) {
		api.accessToken = token
	}
}

func AppSecret(secret string) func(*FacebookAdsAPI) {
	return func(api *FacebookAdsAPI) {
		api.appSecret = secret
	}
}

func AppID(id string) func(*FacebookAdsAPI) {
	return func(api *FacebookAdsAPI) {
		api.appID = id
	}
}
