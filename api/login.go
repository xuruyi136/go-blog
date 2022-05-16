package api

import (
	"go-blog/common"
	"go-blog/service"
	"net/http"
)

func (*Api) Login(w http.ResponseWriter, r *http.Request) {
	//接收用户名和密码
	params := common.GetRequestJsonParam(r)
	username := params["username"].(string)
	passwd := params["passwd"].(string)
	// fmt.Println(username, passwd)
	loginRes, err := service.Login(username, passwd)
	if err != nil {
		common.Error(w, err)
		return
	}
	common.Success(w, loginRes)
}
