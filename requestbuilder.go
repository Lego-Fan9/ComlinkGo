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

func (c *Comlink) post(endpoint string, payload any) (*http.Response, error) {
	var err error

	var headers map[string]string

	if c.DoHMAC {
		headers, err = c.Sign(endpoint, payload)
		if err != nil {
			return nil, fmt.Errorf("%w: %w", ErrInvalidHMAC, err)
		}
	}

	jsonBytes, err := json.Marshal(payload)
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
