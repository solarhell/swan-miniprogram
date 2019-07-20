package swan_miniprogram

// https://smartprogram.baidu.com/docs/develop/api/open_log/

// Response 基础数据
type CommonResponse struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

type LoginResponse struct {
	CommonResponse
	Openid     string `json:"openid"`
	SessionKey string `json:"session_key"`
}

type AccessToken struct {
	CommonResponse
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

type Userinfo struct {
	Openid     string `json:"openid"`
	Nickname   string `json:"nickname"`
	HeadimgUrl string `json:"headimgurl"`
	Gender     int    `json:"sex"`
}
