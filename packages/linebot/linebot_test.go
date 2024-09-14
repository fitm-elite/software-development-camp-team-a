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
	t.Parallel()

	os.Setenv("LINE_CHANNEL_ACCESS_TOKEN", "test_token")
	defer os.Unsetenv("LINE_CHANNEL_ACCESS_TOKEN")

	properties := &linebot.Properties{}
	option := linebot.WithMessagingApi()

	err := option(properties)
	require.NoError(t, err)

	assert.NotNil(t, properties.MessagingApi())
}

func TestNew_Success(t *testing.T) {
	t.Parallel()

	os.Setenv("LINE_CHANNEL_ACCESS_TOKEN", "1NFRUnP5TIKLsSOrCoqPn+fFEaEo+9wM3iTtc4LB6ylMXb6ylOM0tEXgsjsSUMY0MIPjHxunujolGjaJg+DxLBSpLmQl4+3SEbOyafvrtyuEYsvcN2Ghi6emEbNmx199Pbs5AG05102YC2URNI2pjAdB04t89/1O/w1cDnyilFU=")
	defer os.Unsetenv("LINE_CHANNEL_ACCESS_TOKEN")

	messagingApi, err := linebot.New(
		linebot.WithMessagingApi(),
	)

	require.NoError(t, err)
	assert.NotNil(t, messagingApi)
}

func TestNew_ErrorMessagingApiNil(t *testing.T) {
	t.Parallel()

	_, err := linebot.New()

	assert.ErrorIs(t, err, linebot.ErrMessagingApiNil)
}

func TestNew_ErrorInOptionFunc(t *testing.T) {
	t.Parallel()

	faultyOption := func(properties *linebot.Properties) error {
		return errors.New("mock error")
	}

	_, err := linebot.New(faultyOption)

	require.Error(t, err)
	assert.EqualError(t, err, "mock error")
}
