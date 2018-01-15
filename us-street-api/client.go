package street

import (
	"encoding/json"

	"github.com/smartystreets/smartystreets-go-sdk"
)

// Client is responsible for sending batches of addresses to the us-street-api.
type Client struct {
	sender sdk.RequestSender
}

// NewClient creates a client with the provided sender.
func NewClient(sender sdk.RequestSender) *Client {
	return &Client{sender: sender}
}

// SendBatch sends the batch of inputs, populating the output for each input if the batch was successful.
func (c *Client) SendBatch(batch *Batch) error {
	if batch == nil || batch.Length() == 0 {
		return nil
	} else if response, err := c.sender.Send(batch.buildRequest()); err != nil {
		return err
	} else {
		return deserializeResponse(response, batch)
	}
}

func deserializeResponse(response []byte, batch *Batch) error {
	var candidates []*Candidate
	err := json.Unmarshal(response, &candidates)
	if err == nil {
		batch.attach(candidates)
	}
	return err
}
