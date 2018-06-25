package icobench

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net"
	"net/http"
	"strconv"
	"time"
)

const apiv1 = "https://icobench.com/api/v1/"

var defaultHTTPClient = http.Client{
	Timeout: time.Second * 10,
	Transport: &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
	},
}

// Client represents convenient way to access ICObench data.
// Must be initialized with NewClient method to set private and public keys.
type Client struct {
	privateKey []byte
	publicKey  string
	baseURL    string
	HTTPClient *http.Client
}

// NewClient returns ICObench client with specified public and private keys.
// HTTP client could be overrided directly.
func NewClient(private, public string) *Client {
	return &Client{
		privateKey: []byte(private),
		publicKey:  public,
		baseURL:    apiv1,
		HTTPClient: &defaultHTTPClient,
	}
}

// Search returns up to 12 ICOs per page with the number of pages based on the search criteria
// and the filters provided. Use empty SearchRequest to get response from All endpoint.
// The response provides some basic information about the ICOs.
func (c Client) Search(ctx context.Context, filters SearchRequest) (*SearchResponse, error) {
	url := c.baseURL + "icos/all"
	data := bytes.Buffer{}
	err := json.NewEncoder(&data).Encode(filters)
	if err != nil {
		return nil, err
	}
	req, err := c.initRequest(url, &data)
	if err != nil {
		return nil, err
	}
	resp, err := c.makeRequest(req.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	icos := SearchResponse{}
	if err := json.Unmarshal(resp, &icos); err != nil {
		return nil, err
	}
	return &icos, nil
}

// Trending returns up to 8 ICOs that are currently "Hot and Trending" on ICObench.
// The response provides some basic information about ICOs.
func (c Client) Trending(ctx context.Context) (*TrendingResponse, error) {
	url := c.baseURL + "icos/trending"
	req, err := c.initRequest(url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.makeRequest(req.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	icos := TrendingResponse{}
	if err := json.Unmarshal(resp, &icos); err != nil {
		return nil, err
	}
	return &icos, nil
}

// Profile returns detailed information on the ICO
func (c Client) Profile(ctx context.Context, id int) (*ProfileResponse, error) {
	url := c.baseURL + "ico/" + strconv.Itoa(id)
	req, err := c.initRequest(url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.makeRequest(req.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	ico := ProfileResponse{}
	if err := json.Unmarshal(resp, &ico); err != nil {
		return nil, err
	}
	return &ico, nil
}

// Filters returns all available filters
func (c Client) Filters(ctx context.Context) (*FiltersResponse, error) {
	url := c.baseURL + "icos/filters"
	req, err := c.initRequest(url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.makeRequest(req.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	ico := FiltersResponse{}
	if err := json.Unmarshal(resp, &ico); err != nil {
		return nil, err
	}
	return &ico, nil
}

// Ratings returns all ICOs that have received rating for either ICO profile or
// by experts along with their URLs and logos.
func (c Client) Ratings(ctx context.Context) (*RatingsResponse, error) {
	url := c.baseURL + "icos/ratings"
	req, err := c.initRequest(url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.makeRequest(req.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	ico := RatingsResponse{}
	if err := json.Unmarshal(resp, &ico); err != nil {
		return nil, err
	}
	return &ico, nil
}

func (c Client) sign(data []byte) string {
	mac := hmac.New(sha512.New384, c.privateKey)
	mac.Write(data)
	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}

func (c Client) initRequest(url string, data *bytes.Buffer) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodPost, url, data)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-ICObench-Key", c.publicKey)
	req.Header.Set("X-ICObench-Sig", c.sign(data.Bytes()))
	return req, nil
}

func (c Client) makeRequest(req *http.Request) ([]byte, error) {
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
