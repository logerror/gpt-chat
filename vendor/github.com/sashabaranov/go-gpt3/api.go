package gogpt

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Client is OpenAI GPT-3 API client.
type Client struct {
	config ClientConfig
}

// NewClient creates new OpenAI API client.
func NewClient(authToken string) *Client {
	config := DefaultConfig(authToken)
	return &Client{config}
}

// NewClientWithConfig creates new OpenAI API client for specified config.
func NewClientWithConfig(config ClientConfig) *Client {
	return &Client{config}
}

// NewOrgClient creates new OpenAI API client for specified Organization ID.
//
// Deprecated: Please use NewClientWithConfig.
func NewOrgClient(authToken, org string) *Client {
	config := DefaultConfig(authToken)
	config.OrgID = org
	return &Client{config}
}

func (c *Client) sendRequest(req *http.Request, v interface{}) error {
	req.Header.Set("Accept", "application/json; charset=utf-8")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.config.authToken))

	// Check whether Content-Type is already set, Upload Files API requires
	// Content-Type == multipart/form-data
	contentType := req.Header.Get("Content-Type")
	if contentType == "" {
		req.Header.Set("Content-Type", "application/json; charset=utf-8")
	}

	if len(c.config.OrgID) > 0 {
		req.Header.Set("OpenAI-Organization", c.config.OrgID)
	}

	res, err := c.config.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		var errRes ErrorResponse
		err = json.NewDecoder(res.Body).Decode(&errRes)
		if err != nil || errRes.Error == nil {
			reqErr := RequestError{
				StatusCode: res.StatusCode,
				Err:        err,
			}
			return fmt.Errorf("error, %w", &reqErr)
		}
		errRes.Error.StatusCode = res.StatusCode
		return fmt.Errorf("error, status code: %d, message: %w", res.StatusCode, errRes.Error)
	}

	if v != nil {
		if err = json.NewDecoder(res.Body).Decode(&v); err != nil {
			return err
		}
	}

	return nil
}

func (c *Client) fullURL(suffix string) string {
	return fmt.Sprintf("%s%s", c.config.BaseURL, suffix)
}