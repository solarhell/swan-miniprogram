package swan_miniprogram

import (
	"errors"
	"github.com/imroc/req"
)

func Login(appKey, appSecret, code string) (lr LoginResponse, err error) {
	api, err := CodeToURL(appKey, appSecret, code)
	if err != nil {
		return lr, err
	}

	r, err := req.Post(api)
	if err != nil {
		return lr, err
	}

	if r.Response().StatusCode != 200 {
		return lr, ErrConnectBaiduServer
	}

	err = r.ToJSON(&lr)
	if err != nil {
		return lr, err
	}
	if lr.Error != "" {
		return lr, errors.New(lr.ErrorDescription)
	}

	return lr, nil
}
