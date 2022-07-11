package utils

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	packettypes "github.com/teleport-network/teleport/x/xibc/core/packet/types"

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

type Response struct {
	Code    int        `json:"code"`
	Message string     `json:"message"`
	Data    RespStatus `json:"data"`
}

type RespStatus struct {
	Status  int8   `json:"status"`
	Message string `json:"message"`
}

const (
	Failed int8 = iota
	Success
	TooSmall
	Overflow
)

func BridgeTimeLimitCheck(packets []packettypes.Packet, api string) (*Response, error) {
	b, _ := json.Marshal(packets)
	res, err := http.Post(api, "application/json", bytes.NewBuffer(b))
	if err != nil {
		return nil, errors.ErrBridgeConn
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.ErrBridgeConn
	}

	if res.StatusCode == 200 {
		var response *Response
		err = json.Unmarshal(body, &response)
		if err != nil {
			return nil, errors.ErrUnmarshal
		}
		return response, nil
	}
	return nil, errors.ErrBridgeConn
}
