package ComlinkGo

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"

	"github.com/Lego-Fan9/ComlinkGo/httpclient"
)

var (
	ErrMissingComlinkURL   = errors.New("missing required field \"ComlinkURL\"")
	ErrMalformedComlinkURL = errors.New("failed to parse ComlinkURL")
	ErrUnknownComlink      = errors.New("failed to make comlink request")
	ErrBadStatusCode       = errors.New("got a bad status code from comlink")
	ErrInvalidHMAC         = errors.New("failed HMAC")
	ErrInvalidBody         = errors.New("failed to form body")
)

type HMACSettings struct {
	AccessKey string
	SecretKey string
}

type ComlinkSettings struct {
	ComlinkURL string
	HMAC       HMACSettings
	Ctx        context.Context
	Wg         *sync.WaitGroup
}

type Comlink struct {
	ComlinkURL *url.URL
	DoHMAC     bool
	HMAC       struct {
		AccessKey string
		SecretKey string
	}
	HttpClient *httpclient.HTTPClient
	Ctx        context.Context
	Wg         *sync.WaitGroup
}

func GetComlink(settings *ComlinkSettings) (*Comlink, error) {
	var err error

	if settings.Ctx == nil {
		settings.Ctx = context.Background()
	}

	if settings.Wg == nil {
		settings.Wg = &sync.WaitGroup{}
	}

	var comlink Comlink
	comlink.Ctx = settings.Ctx
	comlink.Wg = settings.Wg

	if settings.ComlinkURL == "" {
		return nil, fmt.Errorf("%w", ErrMissingComlinkURL)
	}

	comlink.ComlinkURL, err = url.ParseRequestURI(settings.ComlinkURL)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrMalformedComlinkURL, err)
	}

	if settings.HMAC.AccessKey != "" && settings.HMAC.SecretKey != "" {
		comlink.HMAC.AccessKey = settings.HMAC.AccessKey
		comlink.HMAC.SecretKey = settings.HMAC.SecretKey
		comlink.DoHMAC = true
	} else {
		comlink.DoHMAC = false
	}

	comlink.HttpClient = httpclient.Init(comlink.Ctx, comlink.Wg)

	return &comlink, nil
}

func handleResp(resp *http.Response, err error) (map[string]any, error) {
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrUnknownComlink, err)
	}

	defer func() {
		_, _ = io.Copy(io.Discard, resp.Body)
		resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		comlinkError, err := ComlinkErrorHandler(resp)
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("%w: Code: %s Message: %s", ErrBadStatusCode, comlinkError.Code, comlinkError.Message)
	}

	var response map[string]any

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrUnknownComlink, err)
	}

	return response, nil
}

type ComlinkError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func ComlinkErrorHandler(resp *http.Response) (ComlinkError, error) {
	var comlinkError ComlinkError

	err := json.NewDecoder(resp.Body).Decode(&comlinkError)
	if err != nil {
		return ComlinkError{}, fmt.Errorf("%w: %w", ErrUnknownComlink, err)
	}

	return comlinkError, nil
}

func (c *Comlink) Enums() (map[string]any, error) {
	return handleResp(c.EnumsRaw()) //nolint:bodyclose // Handled by handleResp()
}

func (c *Comlink) EnumsRaw() (*http.Response, error) {
	return c.HttpClient.Get(c.ComlinkURL.String() + "/enums")
}

func (c *Comlink) GameData(payload RequestBody) (map[string]any, error) {
	return handleResp(c.GameDataRaw(payload)) //nolint:bodyclose // Handled by handleResp()
}

func (c *Comlink) GameDataRaw(payload RequestBody) (*http.Response, error) {
	return c.post("/data", payload)
}

func (c *Comlink) Metadata(payload RequestBody) (map[string]any, error) {
	return handleResp(c.MetadataRaw(payload)) //nolint:bodyclose // Handled by handleResp()
}

func (c *Comlink) MetadataRaw(payload RequestBody) (*http.Response, error) {
	return c.post("/metadata", payload)
}

func (c *Comlink) Localization(payload RequestBody) (map[string]any, error) {
	return handleResp(c.LocalizationRaw(payload)) //nolint:bodyclose // Handled by handleResp()
}

func (c *Comlink) LocalizationRaw(payload RequestBody) (*http.Response, error) {
	return c.post("/localization", payload)
}

func (c *Comlink) GetEvents(payload RequestBody) (map[string]any, error) {
	return handleResp(c.GetEventsRaw(payload)) //nolint:bodyclose // Handled by handleResp()
}

func (c *Comlink) GetEventsRaw(payload RequestBody) (*http.Response, error) {
	return c.post("/GetEvents", payload)
}

func (c *Comlink) Guild(payload RequestBody) (map[string]any, error) {
	return handleResp(c.GuildRaw(payload)) //nolint:bodyclose // Handled by handleResp()
}

func (c *Comlink) GuildRaw(payload RequestBody) (*http.Response, error) {
	return c.post("/Guild", payload)
}
func (c *Comlink) GetGuildLeaderboard(payload RequestBody) (map[string]any, error) {
	return handleResp(c.GetGuildLeaderboardRaw(payload)) //nolint:bodyclose // Handled by handleResp()
}

func (c *Comlink) GetGuildLeaderboardRaw(payload RequestBody) (*http.Response, error) {
	return c.post("/getGuildLeaderboard", payload)
}
func (c *Comlink) GetGuilds(payload RequestBody) (map[string]any, error) {
	return handleResp(c.GetGuildsRaw(payload)) //nolint:bodyclose // Handled by handleResp()
}

func (c *Comlink) GetGuildsRaw(payload RequestBody) (*http.Response, error) {
	return c.post("/getGuilds", payload)
}

func (c *Comlink) GetLeaderboard(payload RequestBody) (map[string]any, error) {
	return handleResp(c.GetLeaderboardRaw(payload)) //nolint:bodyclose // Handled by handleResp()
}

func (c *Comlink) GetLeaderboardRaw(payload RequestBody) (*http.Response, error) {
	return c.post("/getLeaderboard", payload)
}

func (c *Comlink) Player(payload RequestBody) (map[string]any, error) {
	return handleResp(c.PlayerRaw(payload)) //nolint:bodyclose // Handled by handleResp()
}

func (c *Comlink) PlayerRaw(payload RequestBody) (*http.Response, error) {
	return c.post("/player", payload)
}

func (c *Comlink) PlayerArena(payload RequestBody) (map[string]any, error) {
	return handleResp(c.PlayerArenaRaw(payload)) //nolint:bodyclose // Handled by handleResp()
}

func (c *Comlink) PlayerArenaRaw(payload RequestBody) (*http.Response, error) {
	return c.post("/playerArena", payload)
}
