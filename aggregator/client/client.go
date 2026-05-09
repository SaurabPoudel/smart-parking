package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/SaurabPoudel/smart-parking/types"
)

type Client struct {
	EndPoint string
}

func NewClient(endpoint string) *Client {
	return &Client{
		EndPoint: endpoint,
	}
}

func (c *Client) AggregateInvoice(invoice types.Invoice) error {
	b, err := json.Marshal(invoice)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", c.EndPoint, bytes.NewReader(b))
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("The service responded with a non 200 status code %d", resp.StatusCode)
	}

	return nil
}
