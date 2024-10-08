package linebot

import (
	"errors"
	"os"

	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
)

// LineBot errors.
var (
	ErrMessagingApiNil = errors.New("messagingApi and messagingApiBlob cannot be nil")
)

// Properties represents a LineBot properties.
type Properties struct {
	messagingApi *messaging_api.MessagingApiAPI
}

// OptionFunc represents a function that configures a LineBot properties.
type OptionFunc func(properties *Properties) error

// WithMessagingApi sets the messaging API for the LineBot properties.
func WithMessagingApi() OptionFunc {
	return func(properties *Properties) error {
		messagingApi, err := messaging_api.NewMessagingApiAPI(os.Getenv("LINE_CHANNEL_ACCESS_TOKEN"))
		if err != nil {
			return err
		}
		properties.messagingApi = messagingApi

		return nil
	}
}

// MessagingApi returns the messaging API.
func (p *Properties) MessagingApi() *messaging_api.MessagingApiAPI {
	return p.messagingApi
}

// New creates a new LineBot client.
func New(options ...OptionFunc) (*messaging_api.MessagingApiAPI, error) {
	var err error

	properties := &Properties{}
	for _, option := range options {
		if err = option(properties); err != nil {
			return nil, err
		}
	}

	if properties.messagingApi == nil {
		return nil, ErrMessagingApiNil
	}

	return properties.messagingApi, nil
}
