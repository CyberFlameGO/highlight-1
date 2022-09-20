// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type Account struct {
	ID                   int        `json:"id"`
	Name                 string     `json:"name"`
	SessionCountCur      int        `json:"session_count_cur"`
	ViewCountCur         int        `json:"view_count_cur"`
	SessionCountPrev     int        `json:"session_count_prev"`
	ViewCountPrev        int        `json:"view_count_prev"`
	SessionCountPrevPrev int        `json:"session_count_prev_prev"`
	SessionLimit         int        `json:"session_limit"`
	PaidPrev             int        `json:"paid_prev"`
	PaidPrevPrev         int        `json:"paid_prev_prev"`
	Email                string     `json:"email"`
	SubscriptionStart    *time.Time `json:"subscription_start"`
	PlanTier             string     `json:"plan_tier"`
	UnlimitedMembers     bool       `json:"unlimited_members"`
	StripeCustomerID     string     `json:"stripe_customer_id"`
	MemberCount          int        `json:"member_count"`
	MemberLimit          *int       `json:"member_limit"`
}

type AccountDetails struct {
	ID                   int                     `json:"id"`
	Name                 string                  `json:"name"`
	SessionCountPerMonth []*NamedCount           `json:"session_count_per_month"`
	SessionCountPerDay   []*NamedCount           `json:"session_count_per_day"`
	StripeCustomerID     string                  `json:"stripe_customer_id"`
	Members              []*AccountDetailsMember `json:"members"`
}

type AccountDetailsMember struct {
	ID         int        `json:"id"`
	Name       string     `json:"name"`
	Email      string     `json:"email"`
	LastActive *time.Time `json:"last_active"`
}

type AdminAboutYouDetails struct {
	FirstName          string  `json:"first_name"`
	LastName           string  `json:"last_name"`
	UserDefinedRole    string  `json:"user_defined_role"`
	UserDefinedPersona string  `json:"user_defined_persona"`
	Referral           string  `json:"referral"`
	Phone              *string `json:"phone"`
}

type AverageSessionLength struct {
	Length float64 `json:"length"`
}

type BillingDetails struct {
	Plan               *Plan `json:"plan"`
	Meter              int64 `json:"meter"`
	MembersMeter       int64 `json:"membersMeter"`
	SessionsOutOfQuota int64 `json:"sessionsOutOfQuota"`
}

type CategoryHistogramBucket struct {
	Category string `json:"category"`
	Count    int    `json:"count"`
}

type CategoryHistogramPayload struct {
	Buckets []*CategoryHistogramBucket `json:"buckets"`
}

type DashboardDefinition struct {
	ID                int                      `json:"id"`
	UpdatedAt         time.Time                `json:"updated_at"`
	ProjectID         int                      `json:"project_id"`
	Name              string                   `json:"name"`
	Metrics           []*DashboardMetricConfig `json:"metrics"`
	LastAdminToEditID *int                     `json:"last_admin_to_edit_id"`
	Layout            *string                  `json:"layout"`
	IsDefault         *bool                    `json:"is_default"`
}

type DashboardMetricConfig struct {
	Name                     string                   `json:"name"`
	Description              string                   `json:"description"`
	ComponentType            *MetricViewComponentType `json:"component_type"`
	MaxGoodValue             *float64                 `json:"max_good_value"`
	MaxNeedsImprovementValue *float64                 `json:"max_needs_improvement_value"`
	PoorValue                *float64                 `json:"poor_value"`
	Units                    *string                  `json:"units"`
	HelpArticle              *string                  `json:"help_article"`
	ChartType                *DashboardChartType      `json:"chart_type"`
	Aggregator               *MetricAggregator        `json:"aggregator"`
	MinValue                 *float64                 `json:"min_value"`
	MinPercentile            *float64                 `json:"min_percentile"`
	MaxValue                 *float64                 `json:"max_value"`
	MaxPercentile            *float64                 `json:"max_percentile"`
	Filters                  []*MetricTagFilter       `json:"filters"`
	Groups                   []string                 `json:"groups"`
}

type DashboardMetricConfigInput struct {
	Name                     string                   `json:"name"`
	Description              string                   `json:"description"`
	ComponentType            *MetricViewComponentType `json:"component_type"`
	MaxGoodValue             *float64                 `json:"max_good_value"`
	MaxNeedsImprovementValue *float64                 `json:"max_needs_improvement_value"`
	PoorValue                *float64                 `json:"poor_value"`
	Units                    *string                  `json:"units"`
	HelpArticle              *string                  `json:"help_article"`
	ChartType                *DashboardChartType      `json:"chart_type"`
	Aggregator               *MetricAggregator        `json:"aggregator"`
	MinValue                 *float64                 `json:"min_value"`
	MinPercentile            *float64                 `json:"min_percentile"`
	MaxValue                 *float64                 `json:"max_value"`
	MaxPercentile            *float64                 `json:"max_percentile"`
	Filters                  []*MetricTagFilterInput  `json:"filters"`
	Groups                   []string                 `json:"groups"`
}

type DashboardParamsInput struct {
	DateRange         *DateRangeInput         `json:"date_range"`
	ResolutionMinutes *int                    `json:"resolution_minutes"`
	Timezone          *string                 `json:"timezone"`
	Units             *string                 `json:"units"`
	Aggregator        *MetricAggregator       `json:"aggregator"`
	Filters           []*MetricTagFilterInput `json:"filters"`
	Groups            []string                `json:"groups"`
}

type DashboardPayload struct {
	Date       string            `json:"date"`
	Value      float64           `json:"value"`
	Aggregator *MetricAggregator `json:"aggregator"`
	Group      *string           `json:"group"`
}

type DateHistogramBucketSize struct {
	CalendarInterval OpenSearchCalendarInterval `json:"calendar_interval"`
	Multiple         int                        `json:"multiple"`
}

type DateHistogramOptions struct {
	BucketSize *DateHistogramBucketSize `json:"bucket_size"`
	TimeZone   string                   `json:"time_zone"`
	Bounds     *DateRangeInput          `json:"bounds"`
}

type DateRangeInput struct {
	StartDate *time.Time `json:"start_date"`
	EndDate   *time.Time `json:"end_date"`
}

type EnhancedUserDetailsResult struct {
	ID      *int          `json:"id"`
	Name    *string       `json:"name"`
	Avatar  *string       `json:"avatar"`
	Bio     *string       `json:"bio"`
	Socials []*SocialLink `json:"socials"`
	Email   *string       `json:"email"`
}

type ErrorDistributionItem struct {
	Name  string `json:"name"`
	Value int64  `json:"value"`
}

type ErrorMetadata struct {
	ErrorID         int        `json:"error_id"`
	SessionID       int        `json:"session_id"`
	SessionSecureID string     `json:"session_secure_id"`
	Environment     *string    `json:"environment"`
	Timestamp       *time.Time `json:"timestamp"`
	Os              *string    `json:"os"`
	Browser         *string    `json:"browser"`
	VisitedURL      *string    `json:"visited_url"`
	Fingerprint     string     `json:"fingerprint"`
	Identifier      *string    `json:"identifier"`
	UserProperties  *string    `json:"user_properties"`
	RequestID       *string    `json:"request_id"`
	Payload         *string    `json:"payload"`
}

type ErrorSearchParamsInput struct {
	DateRange  *DateRangeInput `json:"date_range"`
	Os         *string         `json:"os"`
	Browser    *string         `json:"browser"`
	VisitedURL *string         `json:"visited_url"`
	State      *ErrorState     `json:"state"`
	Event      *string         `json:"event"`
	Type       *string         `json:"type"`
	Query      *string         `json:"query"`
}

type ErrorTrace struct {
	FileName     *string `json:"fileName"`
	LineNumber   *int    `json:"lineNumber"`
	FunctionName *string `json:"functionName"`
	ColumnNumber *int    `json:"columnNumber"`
	Error        *string `json:"error"`
	LineContent  *string `json:"lineContent"`
	LinesBefore  *string `json:"linesBefore"`
	LinesAfter   *string `json:"linesAfter"`
}

type HistogramBucket struct {
	Bucket     float64 `json:"bucket"`
	RangeStart float64 `json:"range_start"`
	RangeEnd   float64 `json:"range_end"`
	Count      int     `json:"count"`
}

type HistogramParamsInput struct {
	DateRange     *DateRangeInput         `json:"date_range"`
	Buckets       *int                    `json:"buckets"`
	MinValue      *float64                `json:"min_value"`
	MinPercentile *float64                `json:"min_percentile"`
	MaxValue      *float64                `json:"max_value"`
	MaxPercentile *float64                `json:"max_percentile"`
	Units         *string                 `json:"units"`
	Filters       []*MetricTagFilterInput `json:"filters"`
}

type HistogramPayload struct {
	Buckets []*HistogramBucket `json:"buckets"`
	Min     float64            `json:"min"`
	Max     float64            `json:"max"`
}

type Invoice struct {
	AmountDue    *int64     `json:"amountDue"`
	AmountPaid   *int64     `json:"amountPaid"`
	AttemptCount *int64     `json:"attemptCount"`
	Date         *time.Time `json:"date"`
	URL          *string    `json:"url"`
	Status       *string    `json:"status"`
}

type LengthRangeInput struct {
	Min *float64 `json:"min"`
	Max *float64 `json:"max"`
}

type LinearTeam struct {
	TeamID string `json:"team_id"`
	Name   string `json:"name"`
	Key    string `json:"key"`
}

type MetricPreview struct {
	Date  time.Time `json:"date"`
	Value float64   `json:"value"`
}

type MetricTagFilter struct {
	Tag   string            `json:"tag"`
	Op    MetricTagFilterOp `json:"op"`
	Value string            `json:"value"`
}

type MetricTagFilterInput struct {
	Tag   string            `json:"tag"`
	Op    MetricTagFilterOp `json:"op"`
	Value string            `json:"value"`
}

type NamedCount struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

type NetworkHistogramParamsInput struct {
	LookbackDays *int                     `json:"lookback_days"`
	Attribute    *NetworkRequestAttribute `json:"attribute"`
}

type NewUsersCount struct {
	Count int64 `json:"count"`
}

type Plan struct {
	Type         PlanType             `json:"type"`
	Interval     SubscriptionInterval `json:"interval"`
	Quota        int                  `json:"quota"`
	MembersLimit *int                 `json:"membersLimit"`
}

type RageClickEventForProject struct {
	Identifier      string `json:"identifier"`
	SessionSecureID string `json:"session_secure_id"`
	TotalClicks     int    `json:"total_clicks"`
	UserProperties  string `json:"user_properties"`
}

type ReferrerTablePayload struct {
	Host    string  `json:"host"`
	Count   int     `json:"count"`
	Percent float64 `json:"percent"`
}

type S3File struct {
	Key *string `json:"key"`
}

type SanitizedAdmin struct {
	ID       int     `json:"id"`
	Name     *string `json:"name"`
	Email    string  `json:"email"`
	PhotoURL *string `json:"photo_url"`
}

type SanitizedAdminInput struct {
	ID    int     `json:"id"`
	Name  *string `json:"name"`
	Email string  `json:"email"`
}

type SanitizedSlackChannel struct {
	WebhookChannel   *string `json:"webhook_channel"`
	WebhookChannelID *string `json:"webhook_channel_id"`
}

type SanitizedSlackChannelInput struct {
	WebhookChannelName *string `json:"webhook_channel_name"`
	WebhookChannelID   *string `json:"webhook_channel_id"`
}

type SearchParamsInput struct {
	UserProperties          []*UserPropertyInput `json:"user_properties"`
	ExcludedProperties      []*UserPropertyInput `json:"excluded_properties"`
	TrackProperties         []*UserPropertyInput `json:"track_properties"`
	ExcludedTrackProperties []*UserPropertyInput `json:"excluded_track_properties"`
	Environments            []*string            `json:"environments"`
	AppVersions             []*string            `json:"app_versions"`
	DateRange               *DateRangeInput      `json:"date_range"`
	LengthRange             *LengthRangeInput    `json:"length_range"`
	Os                      *string              `json:"os"`
	Browser                 *string              `json:"browser"`
	DeviceID                *string              `json:"device_id"`
	VisitedURL              *string              `json:"visited_url"`
	Referrer                *string              `json:"referrer"`
	Identified              *bool                `json:"identified"`
	HideViewed              *bool                `json:"hide_viewed"`
	FirstTime               *bool                `json:"first_time"`
	ShowLiveSessions        *bool                `json:"show_live_sessions"`
	Query                   *string              `json:"query"`
}

type SessionCommentTagInput struct {
	ID   *int   `json:"id"`
	Name string `json:"name"`
}

type SlackSyncResponse struct {
	Success               bool `json:"success"`
	NewChannelsAddedCount int  `json:"newChannelsAddedCount"`
}

type SocialLink struct {
	Type SocialType `json:"type"`
	Link *string    `json:"link"`
}

type SubscriptionDetails struct {
	BaseAmount      int64    `json:"baseAmount"`
	DiscountPercent float64  `json:"discountPercent"`
	DiscountAmount  int64    `json:"discountAmount"`
	LastInvoice     *Invoice `json:"lastInvoice"`
}

type TopSegmentsPayload struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	SessionCount int    `json:"session_count"`
}

type TopUsersPayload struct {
	ID                   int     `json:"id"`
	Identifier           string  `json:"identifier"`
	TotalActiveTime      int     `json:"total_active_time"`
	ActiveTimePercentage float64 `json:"active_time_percentage"`
	UserProperties       string  `json:"user_properties"`
}

type TrackPropertyInput struct {
	ID    *int   `json:"id"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

type User struct {
	ID int `json:"id"`
}

type UserFingerprintCount struct {
	Count int64 `json:"count"`
}

type UserPropertyInput struct {
	ID    *int   `json:"id"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

type DashboardChartType string

const (
	DashboardChartTypeTimeline    DashboardChartType = "Timeline"
	DashboardChartTypeTimelineBar DashboardChartType = "TimelineBar"
	DashboardChartTypeHistogram   DashboardChartType = "Histogram"
)

var AllDashboardChartType = []DashboardChartType{
	DashboardChartTypeTimeline,
	DashboardChartTypeTimelineBar,
	DashboardChartTypeHistogram,
}

func (e DashboardChartType) IsValid() bool {
	switch e {
	case DashboardChartTypeTimeline, DashboardChartTypeTimelineBar, DashboardChartTypeHistogram:
		return true
	}
	return false
}

func (e DashboardChartType) String() string {
	return string(e)
}

func (e *DashboardChartType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = DashboardChartType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid DashboardChartType", str)
	}
	return nil
}

func (e DashboardChartType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ErrorState string

const (
	ErrorStateOpen     ErrorState = "OPEN"
	ErrorStateResolved ErrorState = "RESOLVED"
	ErrorStateIgnored  ErrorState = "IGNORED"
)

var AllErrorState = []ErrorState{
	ErrorStateOpen,
	ErrorStateResolved,
	ErrorStateIgnored,
}

func (e ErrorState) IsValid() bool {
	switch e {
	case ErrorStateOpen, ErrorStateResolved, ErrorStateIgnored:
		return true
	}
	return false
}

func (e ErrorState) String() string {
	return string(e)
}

func (e *ErrorState) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ErrorState(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ErrorState", str)
	}
	return nil
}

func (e ErrorState) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type IntegrationType string

const (
	IntegrationTypeSlack  IntegrationType = "Slack"
	IntegrationTypeLinear IntegrationType = "Linear"
	IntegrationTypeZapier IntegrationType = "Zapier"
	IntegrationTypeFront  IntegrationType = "Front"
)

var AllIntegrationType = []IntegrationType{
	IntegrationTypeSlack,
	IntegrationTypeLinear,
	IntegrationTypeZapier,
	IntegrationTypeFront,
}

func (e IntegrationType) IsValid() bool {
	switch e {
	case IntegrationTypeSlack, IntegrationTypeLinear, IntegrationTypeZapier, IntegrationTypeFront:
		return true
	}
	return false
}

func (e IntegrationType) String() string {
	return string(e)
}

func (e *IntegrationType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = IntegrationType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid IntegrationType", str)
	}
	return nil
}

func (e IntegrationType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type MetricAggregator string

const (
	MetricAggregatorAvg   MetricAggregator = "Avg"
	MetricAggregatorP50   MetricAggregator = "P50"
	MetricAggregatorP75   MetricAggregator = "P75"
	MetricAggregatorP90   MetricAggregator = "P90"
	MetricAggregatorP95   MetricAggregator = "P95"
	MetricAggregatorP99   MetricAggregator = "P99"
	MetricAggregatorMax   MetricAggregator = "Max"
	MetricAggregatorCount MetricAggregator = "Count"
	MetricAggregatorSum   MetricAggregator = "Sum"
)

var AllMetricAggregator = []MetricAggregator{
	MetricAggregatorAvg,
	MetricAggregatorP50,
	MetricAggregatorP75,
	MetricAggregatorP90,
	MetricAggregatorP95,
	MetricAggregatorP99,
	MetricAggregatorMax,
	MetricAggregatorCount,
	MetricAggregatorSum,
}

func (e MetricAggregator) IsValid() bool {
	switch e {
	case MetricAggregatorAvg, MetricAggregatorP50, MetricAggregatorP75, MetricAggregatorP90, MetricAggregatorP95, MetricAggregatorP99, MetricAggregatorMax, MetricAggregatorCount, MetricAggregatorSum:
		return true
	}
	return false
}

func (e MetricAggregator) String() string {
	return string(e)
}

func (e *MetricAggregator) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = MetricAggregator(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid MetricAggregator", str)
	}
	return nil
}

func (e MetricAggregator) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type MetricTagFilterOp string

const (
	MetricTagFilterOpEquals   MetricTagFilterOp = "equals"
	MetricTagFilterOpContains MetricTagFilterOp = "contains"
)

var AllMetricTagFilterOp = []MetricTagFilterOp{
	MetricTagFilterOpEquals,
	MetricTagFilterOpContains,
}

func (e MetricTagFilterOp) IsValid() bool {
	switch e {
	case MetricTagFilterOpEquals, MetricTagFilterOpContains:
		return true
	}
	return false
}

func (e MetricTagFilterOp) String() string {
	return string(e)
}

func (e *MetricTagFilterOp) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = MetricTagFilterOp(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid MetricTagFilterOp", str)
	}
	return nil
}

func (e MetricTagFilterOp) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type MetricViewComponentType string

const (
	MetricViewComponentTypeKeyPerformanceGauge MetricViewComponentType = "KeyPerformanceGauge"
	MetricViewComponentTypeSessionCountChart   MetricViewComponentType = "SessionCountChart"
	MetricViewComponentTypeErrorCountChart     MetricViewComponentType = "ErrorCountChart"
	MetricViewComponentTypeReferrersTable      MetricViewComponentType = "ReferrersTable"
	MetricViewComponentTypeActiveUsersTable    MetricViewComponentType = "ActiveUsersTable"
	MetricViewComponentTypeRageClicksTable     MetricViewComponentType = "RageClicksTable"
	MetricViewComponentTypeTopRoutesTable      MetricViewComponentType = "TopRoutesTable"
	MetricViewComponentTypeTopSegmentsTable    MetricViewComponentType = "TopSegmentsTable"
)

var AllMetricViewComponentType = []MetricViewComponentType{
	MetricViewComponentTypeKeyPerformanceGauge,
	MetricViewComponentTypeSessionCountChart,
	MetricViewComponentTypeErrorCountChart,
	MetricViewComponentTypeReferrersTable,
	MetricViewComponentTypeActiveUsersTable,
	MetricViewComponentTypeRageClicksTable,
	MetricViewComponentTypeTopRoutesTable,
	MetricViewComponentTypeTopSegmentsTable,
}

func (e MetricViewComponentType) IsValid() bool {
	switch e {
	case MetricViewComponentTypeKeyPerformanceGauge, MetricViewComponentTypeSessionCountChart, MetricViewComponentTypeErrorCountChart, MetricViewComponentTypeReferrersTable, MetricViewComponentTypeActiveUsersTable, MetricViewComponentTypeRageClicksTable, MetricViewComponentTypeTopRoutesTable, MetricViewComponentTypeTopSegmentsTable:
		return true
	}
	return false
}

func (e MetricViewComponentType) String() string {
	return string(e)
}

func (e *MetricViewComponentType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = MetricViewComponentType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid MetricViewComponentType", str)
	}
	return nil
}

func (e MetricViewComponentType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type NetworkRequestAttribute string

const (
	NetworkRequestAttributeMethod           NetworkRequestAttribute = "method"
	NetworkRequestAttributeInitiatorType    NetworkRequestAttribute = "initiator_type"
	NetworkRequestAttributeURL              NetworkRequestAttribute = "url"
	NetworkRequestAttributeBodySize         NetworkRequestAttribute = "body_size"
	NetworkRequestAttributeResponseSize     NetworkRequestAttribute = "response_size"
	NetworkRequestAttributeStatus           NetworkRequestAttribute = "status"
	NetworkRequestAttributeLatency          NetworkRequestAttribute = "latency"
	NetworkRequestAttributeRequestID        NetworkRequestAttribute = "request_id"
	NetworkRequestAttributeGraphqlOperation NetworkRequestAttribute = "graphql_operation"
)

var AllNetworkRequestAttribute = []NetworkRequestAttribute{
	NetworkRequestAttributeMethod,
	NetworkRequestAttributeInitiatorType,
	NetworkRequestAttributeURL,
	NetworkRequestAttributeBodySize,
	NetworkRequestAttributeResponseSize,
	NetworkRequestAttributeStatus,
	NetworkRequestAttributeLatency,
	NetworkRequestAttributeRequestID,
	NetworkRequestAttributeGraphqlOperation,
}

func (e NetworkRequestAttribute) IsValid() bool {
	switch e {
	case NetworkRequestAttributeMethod, NetworkRequestAttributeInitiatorType, NetworkRequestAttributeURL, NetworkRequestAttributeBodySize, NetworkRequestAttributeResponseSize, NetworkRequestAttributeStatus, NetworkRequestAttributeLatency, NetworkRequestAttributeRequestID, NetworkRequestAttributeGraphqlOperation:
		return true
	}
	return false
}

func (e NetworkRequestAttribute) String() string {
	return string(e)
}

func (e *NetworkRequestAttribute) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = NetworkRequestAttribute(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid NetworkRequestAttribute", str)
	}
	return nil
}

func (e NetworkRequestAttribute) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type OpenSearchCalendarInterval string

const (
	OpenSearchCalendarIntervalMinute  OpenSearchCalendarInterval = "minute"
	OpenSearchCalendarIntervalHour    OpenSearchCalendarInterval = "hour"
	OpenSearchCalendarIntervalDay     OpenSearchCalendarInterval = "day"
	OpenSearchCalendarIntervalWeek    OpenSearchCalendarInterval = "week"
	OpenSearchCalendarIntervalMonth   OpenSearchCalendarInterval = "month"
	OpenSearchCalendarIntervalQuarter OpenSearchCalendarInterval = "quarter"
	OpenSearchCalendarIntervalYear    OpenSearchCalendarInterval = "year"
)

var AllOpenSearchCalendarInterval = []OpenSearchCalendarInterval{
	OpenSearchCalendarIntervalMinute,
	OpenSearchCalendarIntervalHour,
	OpenSearchCalendarIntervalDay,
	OpenSearchCalendarIntervalWeek,
	OpenSearchCalendarIntervalMonth,
	OpenSearchCalendarIntervalQuarter,
	OpenSearchCalendarIntervalYear,
}

func (e OpenSearchCalendarInterval) IsValid() bool {
	switch e {
	case OpenSearchCalendarIntervalMinute, OpenSearchCalendarIntervalHour, OpenSearchCalendarIntervalDay, OpenSearchCalendarIntervalWeek, OpenSearchCalendarIntervalMonth, OpenSearchCalendarIntervalQuarter, OpenSearchCalendarIntervalYear:
		return true
	}
	return false
}

func (e OpenSearchCalendarInterval) String() string {
	return string(e)
}

func (e *OpenSearchCalendarInterval) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OpenSearchCalendarInterval(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid OpenSearchCalendarInterval", str)
	}
	return nil
}

func (e OpenSearchCalendarInterval) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type PlanType string

const (
	PlanTypeFree       PlanType = "Free"
	PlanTypeBasic      PlanType = "Basic"
	PlanTypeStartup    PlanType = "Startup"
	PlanTypeEnterprise PlanType = "Enterprise"
)

var AllPlanType = []PlanType{
	PlanTypeFree,
	PlanTypeBasic,
	PlanTypeStartup,
	PlanTypeEnterprise,
}

func (e PlanType) IsValid() bool {
	switch e {
	case PlanTypeFree, PlanTypeBasic, PlanTypeStartup, PlanTypeEnterprise:
		return true
	}
	return false
}

func (e PlanType) String() string {
	return string(e)
}

func (e *PlanType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PlanType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PlanType", str)
	}
	return nil
}

func (e PlanType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type SessionCommentType string

const (
	SessionCommentTypeAdmin    SessionCommentType = "Admin"
	SessionCommentTypeFeedback SessionCommentType = "FEEDBACK"
)

var AllSessionCommentType = []SessionCommentType{
	SessionCommentTypeAdmin,
	SessionCommentTypeFeedback,
}

func (e SessionCommentType) IsValid() bool {
	switch e {
	case SessionCommentTypeAdmin, SessionCommentTypeFeedback:
		return true
	}
	return false
}

func (e SessionCommentType) String() string {
	return string(e)
}

func (e *SessionCommentType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SessionCommentType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SessionCommentType", str)
	}
	return nil
}

func (e SessionCommentType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type SessionLifecycle string

const (
	SessionLifecycleAll       SessionLifecycle = "All"
	SessionLifecycleLive      SessionLifecycle = "Live"
	SessionLifecycleCompleted SessionLifecycle = "Completed"
)

var AllSessionLifecycle = []SessionLifecycle{
	SessionLifecycleAll,
	SessionLifecycleLive,
	SessionLifecycleCompleted,
}

func (e SessionLifecycle) IsValid() bool {
	switch e {
	case SessionLifecycleAll, SessionLifecycleLive, SessionLifecycleCompleted:
		return true
	}
	return false
}

func (e SessionLifecycle) String() string {
	return string(e)
}

func (e *SessionLifecycle) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SessionLifecycle(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SessionLifecycle", str)
	}
	return nil
}

func (e SessionLifecycle) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type SocialType string

const (
	SocialTypeGithub   SocialType = "Github"
	SocialTypeLinkedIn SocialType = "LinkedIn"
	SocialTypeTwitter  SocialType = "Twitter"
	SocialTypeFacebook SocialType = "Facebook"
	SocialTypeSite     SocialType = "Site"
)

var AllSocialType = []SocialType{
	SocialTypeGithub,
	SocialTypeLinkedIn,
	SocialTypeTwitter,
	SocialTypeFacebook,
	SocialTypeSite,
}

func (e SocialType) IsValid() bool {
	switch e {
	case SocialTypeGithub, SocialTypeLinkedIn, SocialTypeTwitter, SocialTypeFacebook, SocialTypeSite:
		return true
	}
	return false
}

func (e SocialType) String() string {
	return string(e)
}

func (e *SocialType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SocialType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SocialType", str)
	}
	return nil
}

func (e SocialType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type SubscriptionInterval string

const (
	SubscriptionIntervalMonthly SubscriptionInterval = "Monthly"
	SubscriptionIntervalAnnual  SubscriptionInterval = "Annual"
)

var AllSubscriptionInterval = []SubscriptionInterval{
	SubscriptionIntervalMonthly,
	SubscriptionIntervalAnnual,
}

func (e SubscriptionInterval) IsValid() bool {
	switch e {
	case SubscriptionIntervalMonthly, SubscriptionIntervalAnnual:
		return true
	}
	return false
}

func (e SubscriptionInterval) String() string {
	return string(e)
}

func (e *SubscriptionInterval) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SubscriptionInterval(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SubscriptionInterval", str)
	}
	return nil
}

func (e SubscriptionInterval) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
