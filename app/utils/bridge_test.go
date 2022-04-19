package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const bridgeStatusApi = "https://bridge.qa.davionlabs.com/bridge/bridge_status"

func TestGetBridgeStatus(t *testing.T) {
	status, err := GetBridgeStatus(bridgeStatusApi)
	require.NoError(t, err)
	require.Equal(t, 1, status)
}

func TestRetryGetBridgeStatus(t *testing.T) {
	status, err := RetryGetBridgeStatus(bridgeStatusApi)
	require.NoError(t, err)
	require.Equal(t, 1, status)
}
