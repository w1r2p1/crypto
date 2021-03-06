package binance

import (
	"crypto/sha256"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/meeDamian/crypto"
	"github.com/meeDamian/crypto/utils"
)

func privateRequest(c crypto.Credentials, method, url2 string, params map[string]string) (response *http.Response, err error) {
	v := url.Values{}
	for key, val := range params {
		v.Add(key, val)
	}

	if method == http.MethodGet {
		query := url.Values{}
		query.Add("timestamp", fmt.Sprintf("%d", time.Now().Unix()*1000))
		encoded := query.Encode()
		sign := utils.HmacSign(sha256.New, encoded, c.Secret)

		url2 = fmt.Sprintf("%s?%s&signature=%s", url2, encoded, sign)
	}

	req, err := http.NewRequest(method, url2, strings.NewReader(v.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("X-MBX-APIKEY", c.Key)

	return utils.NetClient().Do(req)
}
