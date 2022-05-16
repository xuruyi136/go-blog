package router

import (
	"go-blog/api"
	"go-blog/views"
	"net/http"
)

func Router() {
	//1. 页面  views 2. api 数据（json） 3. 静态资源
	http.HandleFunc("/", views.HTML.Index)
	// http://127.0.0.1:8080/c/1 1参数 分类id
	http.HandleFunc("/c/", views.HTML.Category)
	http.HandleFunc("/login", views.HTML.Login)
	// http://127.0.0.1:8080/p/1.html详情页
	http.HandleFunc("/p/", views.HTML.Detail)
	//发布写作页
	http.HandleFunc("/writing", views.HTML.Writing)

	http.HandleFunc("/pigeonhole", views.HTML.Pigeonhole)

	//http://127.0.0.1:8080/api/v1/post/29
	http.HandleFunc("/api/v1/post/", api.API.GetPost)

	http.HandleFunc("/api/v1/qiniu/token", api.API.QiniuToken)

	// http://127.0.0.1:8080/api/v1/post/search?val=
	http.HandleFunc("/api/v1/post/search", api.API.SearchPost)
	http.HandleFunc("/api/v1/post", api.API.SaveAndUpdatePost)
	http.HandleFunc("/api/v1/login", api.API.Login)
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))
}
