package main

import (
	"Auto-NCO-ZJGSU/service"
	"github.com/mizuki1412/go-core-kit/class/exception"
	"github.com/mizuki1412/go-core-kit/init/initkit"
	"github.com/mizuki1412/go-core-kit/service/configkit"
	"github.com/spf13/cast"
	"time"
)

func main() {
	initkit.LoadConfig()
	users, err := cast.ToSliceE(configkit.Get("id", ""))
	if err != nil {
		panic(exception.New(err.Error()))
	}
	for _, v := range users {
		user := cast.ToStringMapString(v)
		cookie := service.GetCookie(user["name"], user["psswd"], user["userAgent"], user["home"])
		service.Push2WeChat(user["wechatPushKey"], user["name"], user["trueName"], service.PostInfo(cookie, user["userAgent"]))
		time.Sleep(time.Second * 1)
	}
}
