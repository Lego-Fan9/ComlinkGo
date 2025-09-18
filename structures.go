package ComlinkGo

type ClientSpecsPointer struct {
	Platform        *string `json:"platform,omitempty"`
	BundleId        *string `json:"bundleId,omitempty"`
	ExternalVersion *string `json:"externalVersion,omitempty"`
	InternalVersion *string `json:"internalVersion,omitempty"`
	Region          *string `json:"region,omitempty"`
}

type LeaderboardIdPointer struct {
	LeaderboardType *int `json:"leaderboardType,omitempty"`
	MonthOffset     *int `json:"monthOffset,omitempty"`
}

type SearchCriteriaPointer struct {
	MinMemberCount         *int      `json:"minMemberCount,omitempty"`
	MaxMemberCount         *int      `json:"maxMemberCount,omitempty"`
	IncludeInviteOnly      *bool     `json:"includeInviteOnly,omitempty"`
	MinGuildGalacticPower  *int      `json:"minGuildGalacticPower,omitempty"`
	MaxGuildGalacticPower  *int      `json:"maxGuildGalacticPower,omitempty"`
	RecentTbParticipatedIn *[]string `json:"recentTbParticipatedIn,omitempty"`
}

type PayloadPointer struct {
	SearchCriteria                 *SearchCriteriaPointer `json:"searchCriteria,omitempty"`
	LeaderboardId                  *LeaderboardIdPointer  `json:"leaderboardId,omitempty"`
	ClientSpecs                    *ClientSpecsPointer    `json:"clientSpecs,omitempty"`
	Version                        *string                `json:"version,omitempty"`
	IncludePveUnits                *bool                  `json:"includePveUnits,omitempty"`
	DevicePlatform                 *string                `json:"devicePlatform,omitempty"`
	RequestSegment                 *int                   `json:"requestSegment,omitempty"`
	Items                          *string                `json:"items,omitempty"`
	Id                             *string                `json:"id,omitempty"`
	GuildId                        *string                `json:"guildId,omitempty"`
	IncludeRecentGuildActivityInfo *bool                  `json:"includeRecentGuildActivityInfo,omitempty"`
	Count                          *int                   `json:"count,omitempty"`
	FilterType                     *int                   `json:"filterType,omitempty"`
	Name                           *string                `json:"Name,omitempty"`
	StartIndex                     *int                   `json:"startIndex,omitempty"`
	LeaderboardType                *int                   `json:"leaderboardType,omitempty"`
	EventInstanceId                *string                `json:"eventInstanceId,omitempty"`
	GroupId                        *string                `json:"groupId,omitempty"`
	League                         *int                   `json:"league,omitempty"`
	Division                       *int                   `json:"division,omitempty"`
	AllyCode                       *string                `json:"allyCode,omitempty"`
	PlayerId                       *string                `json:"playerId,omitempty"`
	PlayerDetailsOnly              *bool                  `json:"playerDetailsOnly,omitempty"`
}

type RequestBodyPointer struct {
	Payload *PayloadPointer `json:"payload,omitempty"`
	Enums   *bool           `json:"enums,omitempty"`
	Unzip   *bool           `json:"unzip,omitempty"`
}

type ClientSpecs struct {
	Platform        string `json:"platform,omitempty"`
	BundleId        string `json:"bundleId,omitempty"`
	ExternalVersion string `json:"externalVersion,omitempty"`
	InternalVersion string `json:"internalVersion,omitempty"`
	Region          string `json:"region,omitempty"`
}

type LeaderboardId struct {
	LeaderboardType int `json:"leaderboardType,omitempty"`
	MonthOffset     int `json:"monthOffset,omitempty"`
}

type SearchCriteria struct {
	MinMemberCount         int      `json:"minMemberCount,omitempty"`
	MaxMemberCount         int      `json:"maxMemberCount,omitempty"`
	IncludeInviteOnly      bool     `json:"includeInviteOnly,omitempty"`
	MinGuildGalacticPower  int      `json:"minGuildGalacticPower,omitempty"`
	MaxGuildGalacticPower  int      `json:"maxGuildGalacticPower,omitempty"`
	RecentTbParticipatedIn []string `json:"recentTbParticipatedIn,omitempty"`
}

type Payload struct {
	SearchCriteria                 SearchCriteria `json:"searchCriteria,omitempty"`
	LeaderboardId                  LeaderboardId  `json:"leaderboardId,omitempty"`
	ClientSpecs                    ClientSpecs    `json:"clientSpecs,omitempty"`
	Version                        string         `json:"version,omitempty"`
	IncludePveUnits                bool           `json:"includePveUnits,omitempty"`
	DevicePlatform                 string         `json:"devicePlatform,omitempty"`
	RequestSegment                 int            `json:"requestSegment,omitempty"`
	Items                          string         `json:"items,omitempty"`
	Id                             string         `json:"id,omitempty"`
	GuildId                        string         `json:"guildId,omitempty"`
	IncludeRecentGuildActivityInfo bool           `json:"includeRecentGuildActivityInfo,omitempty"`
	Count                          int            `json:"count,omitempty"`
	FilterType                     int            `json:"filterType,omitempty"`
	Name                           string         `json:"Name,omitempty"`
	StartIndex                     int            `json:"startIndex,omitempty"`
	LeaderboardType                int            `json:"leaderboardType,omitempty"`
	EventInstanceId                string         `json:"eventInstanceId,omitempty"`
	GroupId                        string         `json:"groupId,omitempty"`
	League                         int            `json:"league,omitempty"`
	Division                       int            `json:"division,omitempty"`
	AllyCode                       string         `json:"allyCode,omitempty"`
	PlayerId                       string         `json:"playerId,omitempty"`
	PlayerDetailsOnly              bool           `json:"playerDetailsOnly,omitempty"`
}

type RequestBody struct {
	Payload Payload `json:"payload,omitempty"`
	Enums   bool    `json:"enums,omitempty"`
	Unzip   bool    `json:"unzip,omitempty"`
}
