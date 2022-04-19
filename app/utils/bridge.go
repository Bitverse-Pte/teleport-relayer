package utils

import (
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/avast/retry-go"

	"github.com/teleport-network/teleport-relayer/app/types/errors"
)

func GetBridgeStatus(api string) (int, error) {
	resp, err := http.Get(api)
	if err != nil {
		return -1, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return -1, err
	}
	if resp.StatusCode == 200 {
		status, err := strconv.Atoi(string(body))
		if err != nil {
			return status, err
		}
		return status, nil
	}
	return -1, errors.ErrBridgeConn
}

func RetryGetBridgeStatus(api string) (status int, err error) {
	status = -1
	retryableFunc := func() error {
		status, err = GetBridgeStatus(api)
		if err != nil {
			return err
		}
		return nil
	}

	retryIfFunc := func(err error) bool {
		return true
	}

	onRetryFunc := func(n uint, err error) {
		time.Sleep(5 * time.Second)
	}

	err = retry.Do(
		retryableFunc,
		retry.Attempts(3),
		retry.RetryIf(retryIfFunc),
		retry.OnRetry(onRetryFunc),
	)
	return
}
