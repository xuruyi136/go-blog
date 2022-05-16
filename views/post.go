package views

import (
	"errors"
	"go-blog/common"
	"go-blog/service"
	"net/http"
	"strconv"
	"strings"
)

func (*HTMLApi) Detail(w http.ResponseWriter, r *http.Request) {
	detail := common.Template.Detail
	path := r.URL.Path

	pIdStr := strings.TrimPrefix(path, "/p/") //7.html去除后缀
	pIdStr = strings.TrimSuffix(pIdStr, ".html")
	pid, err := strconv.Atoi(pIdStr) //转换int
	if err != nil {
		detail.WriteErr(w, errors.New("不识别此请求路径"))
		return
	}

	postRes, err := service.GetPostDetaile(pid)
	if err != nil {
		detail.WriteErr(w, errors.New("查询出错"))
		return
	}
	detail.WriteData(w, postRes)
}
