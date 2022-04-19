package dto

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_relay(t *testing.T) {
	params := url.Values{"hash": {"A0760BC19BD644CFCE7834DAF337D70A3C796CABD8F3A23954997E7FC415DA4B"}}
	resp, err := http.PostForm("http://localhost:8080/relayer/teleport/relay", params)
	require.NoError(t, err)

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))
}
