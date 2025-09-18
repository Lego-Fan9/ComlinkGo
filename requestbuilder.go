package ComlinkGo

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5" //nolint:gosec
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func (c *Comlink) Sign(endpoint string, payload any) (map[string]string, error) {
	reqTime := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)

	headers := map[string]string{
		"X-Date": reqTime,
	}

	mac := hmac.New(sha256.New, []byte(c.HMAC.SecretKey))
	mac.Write([]byte(reqTime))
	mac.Write([]byte(http.MethodPost))
	mac.Write([]byte(endpoint))

	var payloadBytes []byte

	var err error

	if payload != nil {
		payloadBytes, err = json.Marshal(payload)
		if err != nil {
			return nil, err
		}
	} else {
		payloadBytes = []byte("{}")
	}

	md5Sum := md5.Sum(payloadBytes) //nolint:gosec
	md5Hex := hex.EncodeToString(md5Sum[:])

	mac.Write([]byte(md5Hex))

	signature := hex.EncodeToString(mac.Sum(nil))

	headers["Authorization"] = fmt.Sprintf("HMAC-SHA256 Credential=%s,Signature=%s", c.HMAC.AccessKey, signature)

	return headers, nil
}

func (c *Comlink) post(endpoint string, payload RequestBody) (*http.Response, error) {
	var err error

	var headers map[string]string

	convertedPayload := convertRequestBody(payload)

	if c.DoHMAC {
		headers, err = c.Sign(endpoint, convertedPayload)
		if err != nil {
			return nil, fmt.Errorf("%w: %w", ErrInvalidHMAC, err)
		}
	}

	jsonBytes, err := json.Marshal(convertedPayload)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrInvalidBody, err)
	}

	body := bytes.NewBuffer(jsonBytes)

	req, err := http.NewRequest(http.MethodPost, c.ComlinkURL.String()+endpoint, body)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrUnknownComlink, err)
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	req.Header.Set("Content-Type", "application/json")

	return c.HttpClient.DoWithRetry(req)
}

func PtrIfNotZero[T comparable](v T) *T {
	var zero T
	if v == zero {
		return nil
	}

	return &v
}

func PtrIfNotEmpty[T any](s []T) *[]T {
	if len(s) == 0 {
		return nil
	}

	return &s
}

//nolint:cyclop,funlen // The complexity is fine
func convertRequestBody(payload RequestBody) RequestBodyPointer {
	var response RequestBodyPointer

	if payload.Payload.SearchCriteria.MinMemberCount != 0 ||
		payload.Payload.SearchCriteria.MaxMemberCount != 0 ||
		payload.Payload.SearchCriteria.IncludeInviteOnly ||
		len(payload.Payload.SearchCriteria.RecentTbParticipatedIn) > 0 {
		response.Payload.SearchCriteria = &SearchCriteriaPointer{}
		response.Payload.SearchCriteria.MinMemberCount = PtrIfNotZero(payload.Payload.SearchCriteria.MinMemberCount)
		response.Payload.SearchCriteria.MaxMemberCount = PtrIfNotZero(payload.Payload.SearchCriteria.MaxMemberCount)
		response.Payload.SearchCriteria.IncludeInviteOnly = PtrIfNotZero(payload.Payload.SearchCriteria.IncludeInviteOnly)
		response.Payload.SearchCriteria.MinGuildGalacticPower = PtrIfNotZero(payload.Payload.SearchCriteria.MinGuildGalacticPower)
		response.Payload.SearchCriteria.MaxGuildGalacticPower = PtrIfNotZero(payload.Payload.SearchCriteria.MaxGuildGalacticPower)
		response.Payload.SearchCriteria.RecentTbParticipatedIn = PtrIfNotEmpty(payload.Payload.SearchCriteria.RecentTbParticipatedIn)
	}

	if payload.Payload.LeaderboardId.LeaderboardType != 0 ||
		payload.Payload.LeaderboardId.MonthOffset != 0 {
		response.Payload.LeaderboardId = &LeaderboardIdPointer{}
		response.Payload.LeaderboardId.LeaderboardType = PtrIfNotZero(payload.Payload.LeaderboardId.LeaderboardType)
		response.Payload.LeaderboardId.MonthOffset = PtrIfNotZero(payload.Payload.LeaderboardId.MonthOffset)
	}

	if payload.Payload.ClientSpecs.Platform != "" ||
		payload.Payload.ClientSpecs.BundleId != "" ||
		payload.Payload.ClientSpecs.ExternalVersion != "" ||
		payload.Payload.ClientSpecs.InternalVersion != "" ||
		payload.Payload.ClientSpecs.Region != "" {
		response.Payload.ClientSpecs = &ClientSpecsPointer{}
		response.Payload.ClientSpecs.Platform = PtrIfNotZero(payload.Payload.ClientSpecs.Platform)
		response.Payload.ClientSpecs.BundleId = PtrIfNotZero(payload.Payload.ClientSpecs.BundleId)
		response.Payload.ClientSpecs.ExternalVersion = PtrIfNotZero(payload.Payload.ClientSpecs.ExternalVersion)
		response.Payload.ClientSpecs.InternalVersion = PtrIfNotZero(payload.Payload.ClientSpecs.InternalVersion)
		response.Payload.ClientSpecs.Region = PtrIfNotZero(payload.Payload.ClientSpecs.Region)
	}

	if doesPayloadPointerNeedToExist(payload) {
		response.Payload = &PayloadPointer{}
		response.Payload.Version = PtrIfNotZero(payload.Payload.Version)
		response.Payload.IncludePveUnits = PtrIfNotZero(payload.Payload.IncludePveUnits)
		response.Payload.DevicePlatform = PtrIfNotZero(payload.Payload.DevicePlatform)
		response.Payload.RequestSegment = PtrIfNotZero(payload.Payload.RequestSegment)
		response.Payload.Items = PtrIfNotZero(payload.Payload.Items)
		response.Payload.Id = PtrIfNotZero(payload.Payload.Id)
		response.Payload.GuildId = PtrIfNotZero(payload.Payload.GuildId)
		response.Payload.IncludeRecentGuildActivityInfo = PtrIfNotZero(payload.Payload.IncludeRecentGuildActivityInfo)
		response.Payload.Count = PtrIfNotZero(payload.Payload.Count)
		response.Payload.FilterType = PtrIfNotZero(payload.Payload.FilterType)
		response.Payload.Name = PtrIfNotZero(payload.Payload.Name)
		response.Payload.StartIndex = PtrIfNotZero(payload.Payload.StartIndex)
		response.Payload.LeaderboardType = PtrIfNotZero(payload.Payload.LeaderboardType)
		response.Payload.EventInstanceId = PtrIfNotZero(payload.Payload.EventInstanceId)
		response.Payload.GroupId = PtrIfNotZero(payload.Payload.GroupId)
		response.Payload.League = PtrIfNotZero(payload.Payload.League)
		response.Payload.Division = PtrIfNotZero(payload.Payload.Division)
		response.Payload.AllyCode = PtrIfNotZero(payload.Payload.AllyCode)
		response.Payload.PlayerId = PtrIfNotZero(payload.Payload.PlayerId)
		response.Payload.PlayerDetailsOnly = PtrIfNotZero(payload.Payload.PlayerDetailsOnly)
	}

	response.Enums = PtrIfNotZero(payload.Enums)
	response.Unzip = PtrIfNotZero(payload.Unzip)

	return response
}

//nolint:cyclop,funlen,nolintlint // The complexity is fine. nolintlint is there because of a bug in golangci
func doesPayloadPointerNeedToExist(payload RequestBody) bool {
	if payload.Payload.Version != "" ||
		payload.Payload.IncludePveUnits ||
		payload.Payload.DevicePlatform != "" ||
		payload.Payload.RequestSegment != 0 ||
		payload.Payload.Items != "" ||
		payload.Payload.Id != "" ||
		payload.Payload.GuildId != "" ||
		payload.Payload.IncludeRecentGuildActivityInfo ||
		payload.Payload.Count != 0 ||
		payload.Payload.FilterType != 0 ||
		payload.Payload.Name != "" ||
		payload.Payload.StartIndex != 0 ||
		payload.Payload.LeaderboardType != 0 ||
		payload.Payload.EventInstanceId != "" ||
		payload.Payload.GroupId != "" ||
		payload.Payload.League != 0 ||
		payload.Payload.Division != 0 ||
		payload.Payload.AllyCode != "" ||
		payload.Payload.PlayerId != "" ||
		payload.Payload.PlayerDetailsOnly {
		return true
	}

	return false
}
