package fields

type UserField string

const (
	UserFieldCreatedAt         UserField = "created_at"
	UserFieldDescription       UserField = "description"
	UserFieldEntities          UserField = "entities"
	UserFieldID                UserField = "id"
	UserFieldLocation          UserField = "location"
	UserFieldName              UserField = "name"
	UserFieldPinnedTweetID     UserField = "pinned_tweet_id"
	UserFieldProfileImageUrl   UserField = "profile_image_url"
	UserFieldProfileBannerUrl  UserField = "profile_banner_url"
	UserFieldProtected         UserField = "protected"
	UserFieldPublicMetrics     UserField = "public_metrics"
	UserFieldUrl               UserField = "url"
	UserFieldUsername          UserField = "username"
	UserFieldVerified          UserField = "verified"
	UserFieldWithheld          UserField = "withheld"
	UserFieldMostRecentTweetID UserField = "most_recent_tweet_id"
	UserFieldAffiliation       UserField = "affiliation"
	UserFieldConnectionStatus  UserField = "connection_status"
	UserFieldReceivesYourDm    UserField = "receives_your_dm"
	UserFieldSubscriptionType  UserField = "subscription_type"
	UserFieldVerifiedType      UserField = "verified_type"
)

func (f UserField) String() string {
	return string(f)
}

type UserFieldList []UserField

func (fl UserFieldList) FieldsName() string {
	return "user.fields"
}

func (fl UserFieldList) Values() []string {
	if fl == nil {
		return []string{}
	}

	s := []string{}
	for _, f := range fl {
		s = append(s, f.String())
	}

	return s
}
