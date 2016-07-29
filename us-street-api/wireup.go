package us_street

import (
	"net/http"
	"net/url"

	"fmt"
	"time"

	"bitbucket.org/smartystreets/smartystreets-go-sdk"
	"bitbucket.org/smartystreets/smartystreets-go-sdk/internal/sdk"
)

// ClientBuilder is responsible for accepting credentials and other configuration options to combine
// all components necessary to assemble a fully functional Client for use in an application.
type ClientBuilder struct {
	credential smarty_sdk.Credential
	baseURL    string
	retries    int
	timeout    time.Duration
}

// NewClientBuilder creates a new client builder, ready to receive calls to its chain-able methods.
func NewClientBuilder() *ClientBuilder {
	return &ClientBuilder{
		credential: &sdk.NopCredential{},
		timeout:    time.Second * 10,
	}
}

// WithSecretKeyCredential allows the caller to set the authID and authToken for use with the client.
// In all but very few cases calling this method with a valid authID and authToken is required.
func (b *ClientBuilder) WithSecretKeyCredential(authID, authToken string) *ClientBuilder {
	b.credential = &smarty_sdk.SecretKeyCredential{AuthID: authID, AuthToken: authToken}
	return b
}

// WithSecretKeyCredential allows the caller to specify the url that the client will use.
// In all but very few use cases the default value is sufficient and this method should not be called.
func (b *ClientBuilder) WithCustomBaseURL(uri string) *ClientBuilder {
	_, err := url.Parse(uri)
	if err != nil {
		panic(fmt.Sprint("Could not parse provided address:", err.Error()))
	}
	b.baseURL = uri
	return b
}

// WithMaxRetry allows the caller to specify the number of times an API request will be resent in the
// case of network errors or unexpected results.
func (b *ClientBuilder) WithMaxRetry(retries int) *ClientBuilder {
	if retries < 0 {
		panic(fmt.Sprintf("Please provide a non-negative number of retry attempts (you supplied %d).", retries))
	}
	b.retries = retries
	return b
}

// WithTimeout allows the caller to specify the timeout for all API requests.
func (b *ClientBuilder) WithTimeout(duration time.Duration) *ClientBuilder {
	if duration < 0 {
		panic(fmt.Sprintf("Please provide a non-negative duration (you supplied %s).", duration.String()))
	}
	b.timeout = duration
	return b
}

// Builds the client using the provided configuration details provided by other methods on the ClientBuilder.
func (b *ClientBuilder) Build() *Client {
	client := http.Client{Timeout: b.timeout}
	retryClient := sdk.NewRetryClient(client, b.retries)
	signingClient := sdk.NewSigningClient(retryClient, b.credential)
	sender := sdk.NewHTTPSender(signingClient)
	return NewClient(sender)
}
