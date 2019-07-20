# swan-miniprogram
[![Build Status](https://travis-ci.org/solarhell/mina.svg?branch=master)](https://travis-ci.org/solarhell/mina)
百度智能小程序 golang sdk


## done
登录
AccessToken(需持久化 防止超过请求限制)
数据解密

## todo
模板消息


## usage

### 登录
```go
package main

import (
	swan "github.com/solarhell/swan-miniprogram"
)

func main() {
	ui, err := swan.Login("appKey", "appSecret", "code")
	...
}
```
