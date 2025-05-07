package bunnyrequestrest

type DrmProvider struct {
	Enabled                   bool    `json:"Enabled"`
	CertificateId             *string `json:"CertificateId"`
	CertificateExpirationDate *string `json:"CertificateExpirationDate"`
	Provider                  *string `json:"Provider"`
	SdOnlyForL3               bool    `json:"SdOnlyForL3,omitempty"`
	MinClientSecurityLevel    *int    `json:"MinClientSecurityLevel,omitempty"`
}

type AppleFairPlayDrm struct {
	Enabled                   bool    `json:"Enabled"`
	CertificateId             *string `json:"CertificateId"`
	CertificateExpirationDate *string `json:"CertificateExpirationDate"`
	Provider                  *string `json:"Provider"`
}

type CreateVideoLibraryResponse struct {
	Id                                      int              `json:"Id"`
	Name                                    string           `json:"Name"`
	VideoCount                              int              `json:"VideoCount"`
	TrafficUsage                            int              `json:"TrafficUsage"`
	StorageUsage                            int              `json:"StorageUsage"`
	DateCreated                             string           `json:"DateCreated"`
	ReplicationRegions                      []string         `json:"ReplicationRegions"`
	ApiKey                                  string           `json:"ApiKey"`
	ReadOnlyApiKey                          string           `json:"ReadOnlyApiKey"`
	HasWatermark                            bool             `json:"HasWatermark"`
	WatermarkPositionLeft                   int              `json:"WatermarkPositionLeft"`
	WatermarkPositionTop                    int              `json:"WatermarkPositionTop"`
	WatermarkWidth                          int              `json:"WatermarkWidth"`
	PullZoneId                              int              `json:"PullZoneId"`
	StorageZoneId                           int              `json:"StorageZoneId"`
	WatermarkHeight                         int              `json:"WatermarkHeight"`
	EnabledResolutions                      string           `json:"EnabledResolutions"`
	ViAiPublisherId                         *string          `json:"ViAiPublisherId"`
	VastTagUrl                              *string          `json:"VastTagUrl"`
	WebhookUrl                              *string          `json:"WebhookUrl"`
	CaptionsFontSize                        int              `json:"CaptionsFontSize"`
	CaptionsFontColor                       string           `json:"CaptionsFontColor"`
	CaptionsBackground                      string           `json:"CaptionsBackground"`
	UILanguage                              string           `json:"UILanguage"`
	AllowEarlyPlay                          bool             `json:"AllowEarlyPlay"`
	PlayerTokenAuthenticationEnabled        bool             `json:"PlayerTokenAuthenticationEnabled"`
	AllowedReferrers                        []string         `json:"AllowedReferrers"`
	BlockedReferrers                        []string         `json:"BlockedReferrers"`
	BlockNoneReferrer                       bool             `json:"BlockNoneReferrer"`
	EnableMP4Fallback                       bool             `json:"EnableMP4Fallback"`
	KeepOriginalFiles                       bool             `json:"KeepOriginalFiles"`
	AllowDirectPlay                         bool             `json:"AllowDirectPlay"`
	EnableDRM                               bool             `json:"EnableDRM"`
	DrmVersion                              int              `json:"DrmVersion"`
	AppleFairPlayDrm                        AppleFairPlayDrm `json:"AppleFairPlayDrm"`
	GoogleWidevineDrm                       DrmProvider      `json:"GoogleWidevineDrm"`
	Bitrate240p                             int              `json:"Bitrate240p"`
	Bitrate360p                             int              `json:"Bitrate360p"`
	Bitrate480p                             int              `json:"Bitrate480p"`
	Bitrate720p                             int              `json:"Bitrate720p"`
	Bitrate1080p                            int              `json:"Bitrate1080p"`
	Bitrate1440p                            int              `json:"Bitrate1440p"`
	Bitrate2160p                            int              `json:"Bitrate2160p"`
	ApiAccessKey                            *string          `json:"ApiAccessKey"`
	ShowHeatmap                             bool             `json:"ShowHeatmap"`
	EnableContentTagging                    bool             `json:"EnableContentTagging"`
	PullZoneType                            int              `json:"PullZoneType"`
	CustomHTML                              *string          `json:"CustomHTML"`
	Controls                                string           `json:"Controls"`
	PlayerKeyColor                          string           `json:"PlayerKeyColor"`
	FontFamily                              string           `json:"FontFamily"`
	WatermarkVersion                        int              `json:"WatermarkVersion"`
	EnableTranscribing                      bool             `json:"EnableTranscribing"`
	EnableTranscribingTitleGeneration       bool             `json:"EnableTranscribingTitleGeneration"`
	EnableTranscribingDescriptionGeneration bool             `json:"EnableTranscribingDescriptionGeneration"`
	TranscribingCaptionLanguages            []string         `json:"TranscribingCaptionLanguages"`
	RememberPlayerPosition                  bool             `json:"RememberPlayerPosition"`
	EnableMultiAudioTrackSupport            bool             `json:"EnableMultiAudioTrackSupport"`
	UseSeparateAudioStream                  bool             `json:"UseSeparateAudioStream"`
	JitEncodingEnabled                      bool             `json:"JitEncodingEnabled"`
	EncodingTier                            int              `json:"EncodingTier"`
	OutputCodecs                            string           `json:"OutputCodecs"`
	DrmBasePriceOverride                    float32          `json:"DrmBasePriceOverride"`
	DrmCostPerLicenseOverride               *int             `json:"DrmCostPerLicenseOverride"`
	TranscribingPriceOverride               *int             `json:"TranscribingPriceOverride"`
	PremiumEncodingPriceOverride            *int             `json:"PremiumEncodingPriceOverride"`
	MonthlyChargesTranscribing              float32          `json:"MonthlyChargesTranscribing"`
	MonthlyChargesPremiumEncoding           float32          `json:"MonthlyChargesPremiumEncoding"`
	MonthlyChargesEnterpriseDrm             float32          `json:"MonthlyChargesEnterpriseDrm"`
}
