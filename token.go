package swan_miniprogram

import (
	"errors"
	"github.com/imroc/req"
)

func GetAccessToken(appKey, appSecret string) (ak AccessToken, err error) {
	api, err := TokenURL(appKey, appSecret)
	if err != nil {
		return ak, err
	}

	r, err := req.Get(api)
	if err != nil {
		return ak, err
	}

	if r.Response().StatusCode != 200 {
		return ak, ErrConnectBaiduServer
	}

	err = r.ToJSON(&ak)
	if err != nil {
		return ak, err
	}

	if ak.Error != "" {
		return ak, errors.New(ak.ErrorDescription)
	}

	return ak, nil
}
