package icobench

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	"io/ioutil"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

const (
	APIv1       = "https://icobench.com/api/v1/"
	ErrICObench = "icobench client"
)

var defaultHTTPClient = http.Client{
	Timeout: time.Second * 10,
	Transport: &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
	},
}

type Client struct {
	PrivateKey []byte
	PublicKey  string
	APIURL     string
	HTTPClient http.Client
}

func (c Client) AllICO(filters AllICORequest) (*AllICOResponse, error) {
	url := c.APIURL + "icos/all"
	data, err := filters.MarshalJSON()
	if err != nil {
		return nil, errors.Wrap(err, ErrICObench)
	}
	req, err := c.initRequest(url, data)
	if err != nil {
		return nil, errors.Wrap(err, ErrICObench)
	}
	resp, err := c.makeRequest(req)
	if err != nil {
		return nil, errors.Wrap(err, ErrICObench)
	}
	icos := AllICOResponse{}
	if err := icos.UnmarshalJSON(resp); err != nil {
		return nil, errors.Wrap(err, ErrICObench)
	}
	return &icos, nil
}

func (c Client) Trending() (*TrendingResponse, error) {
	url := c.APIURL + "icos/trending"
	req, err := c.initRequest(url, []byte{})
	if err != nil {
		return nil, errors.Wrap(err, ErrICObench)
	}
	resp, err := c.makeRequest(req)
	if err != nil {
		return nil, errors.Wrap(err, ErrICObench)
	}
	icos := TrendingResponse{}
	if err := icos.UnmarshalJSON(resp); err != nil {
		return nil, errors.Wrap(err, ErrICObench)
	}
	return &icos, nil
}

func (c Client) ICODetails(ID int64) (*DetailedICO, error) {
	url := c.APIURL + "ico/" + strconv.Itoa(int(ID))
	req, err := c.initRequest(url, []byte{})
	if err != nil {
		return nil, errors.Wrap(err, ErrICObench)
	}
	resp, err := c.makeRequest(req)
	if err != nil {
		return nil, errors.Wrap(err, ErrICObench)
	}
	ico := DetailedICO{}
	if err := ico.UnmarshalJSON(resp); err != nil {
		return nil, errors.Wrap(err, ErrICObench)
	}
	return &ico, nil
}

func (c Client) sign(data []byte) string {
	mac := hmac.New(sha512.New384, c.PrivateKey)
	mac.Write(data)
	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}

func (c Client) initRequest(url string, data []byte) (*http.Request, error) {
	buf := bytes.NewReader(data)
	req, err := http.NewRequest("POST", url, buf)
	if err != nil {
		return nil, errors.Wrap(err, ErrICObench)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-ICObench-Key", c.PublicKey)
	req.Header.Set("X-ICObench-Sig", c.sign(data))
	return req, nil
}

func (c Client) makeRequest(req *http.Request) ([]byte, error) {
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, ErrICObench)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, ErrICObench)
	}
	return body, nil
}

func NewClient(private, public string) *Client {
	return &Client{
		PrivateKey: []byte(private),
		PublicKey:  public,
		APIURL:     APIv1,
		HTTPClient: defaultHTTPClient,
	}
}
