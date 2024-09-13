package linebot_test

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/fitm-elite/elebs/packages/linebot"
)

func TestWithMessagingApi(t *testing.T) {
	os.Setenv("LINE_CHANNEL_ACCESS_TOKEN", "test_token")
	defer os.Unsetenv("LINE_CHANNEL_ACCESS_TOKEN")

	properties := &linebot.Properties{}
	option := linebot.WithMessagingApi()

	err := option(properties)
	require.NoError(t, err)

	assert.NotNil(t, properties.MessagingApi())
}

func TestWithMessagingApiBlob(t *testing.T) {
	os.Setenv("LINE_CHANNEL_ACCESS_TOKEN", "test_token")
	defer os.Unsetenv("LINE_CHANNEL_ACCESS_TOKEN")

	properties := &linebot.Properties{}
	option := linebot.WithMessagingApiBlob()

	err := option(properties)
	require.NoError(t, err)

	assert.NotNil(t, properties.MessagingApiBlob())
}

func TestNew_Success(t *testing.T) {
	os.Setenv("LINE_CHANNEL_ACCESS_TOKEN", "test_token")
	defer os.Unsetenv("LINE_CHANNEL_ACCESS_TOKEN")

	messagingApi, messagingApiBlob, err := linebot.New(
		linebot.WithMessagingApi(), linebot.WithMessagingApiBlob(),
	)

	require.NoError(t, err)
	assert.NotNil(t, messagingApi)
	assert.NotNil(t, messagingApiBlob)
}

func TestNew_ErrorMessagingApiNil(t *testing.T) {
	_, _, err := linebot.New()

	assert.ErrorIs(t, err, linebot.ErrMessagingApiNil)
}

func TestNew_ErrorInOptionFunc(t *testing.T) {
	faultyOption := func(properties *linebot.Properties) error {
		return errors.New("mock error")
	}

	_, _, err := linebot.New(faultyOption)

	require.Error(t, err)
	assert.EqualError(t, err, "mock error")
}
