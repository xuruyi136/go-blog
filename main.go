package main

import (
	"fmt"
	"go-blog/common"
	"strings"
)

/* type IndexData struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

func IsODD(num int) bool {
	return num%2 == 0
}
func GetNextName(strs []string, index int) string {
	return strs[index+1]
}
func Date(layout string) string {
	return time.Now().Format(layout)
} */

/* func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("hellow blog"))
	var IndexData IndexData
	IndexData.Title = "go博客"
	IndexData.Desc = "入门教程"
	jsonstr, _ := json.Marshal(IndexData)
	w.Write(jsonstr)
} */

/* func index(w http.ResponseWriter, r *http.Request) {
	t := template.New("index.html")
	//1. 拿到当前的路径
	path := config.Cfg.System.CurrentDir
	//访问博客首页模板的时候，因为有多个模板的嵌套，解析文件的时候，需要将其涉及到的所有模板都进行解析
	home := path + "/template/home.html"
	header := path + "/template/layout/header.html"
	footer := path + "/template/layout/footer.html"
	personal := path + "/template/layout/personal.html"
	post := path + "/template/layout/post-list.html"
	pagination := path + "/template/layout/pagination.html"
	t.Funcs(template.FuncMap{"isODD": IsODD, "getNextName": GetNextName, "date": Date})
	t, err := t.ParseFiles(path+"/template/index.html", home, header, footer, personal, post, pagination)
	if err != nil {
		log.Println(err)
	}
	//页面上涉及到的所有的数据，必须有定义
	var categorys = []models.Category{
		{
			Cid:  1,
			Name: "go",
		},
	}
	var posts = []models.PostMore{
		{
			Pid:          1,
			Title:        "go博客",
			Content:      "内容",
			UserName:     "码神",
			ViewCount:    123,
			CreateAt:     "2022-02-20",
			CategoryId:   1,
			CategoryName: "go",
			Type:         0,
		},
	}
	//https: //www.cnblogs.com/flipped/p/15637466.html
	var hr = &models.HomeResponse{
		Viewer:    config.Cfg.Viewer,
		Categorys: categorys,
		Posts:     posts,
		Total:     1,
		Page:      1,
		Pages:     []int{1},
		PageEnd:   true,
	}
	t.Execute(w, hr)
} */

func init() {
	//模板加载https://www.bilibili.com/video/BV1VS4y1F7NM?p=23  https://golang-tech-stack.com/post/3833
	common.LoadTemplate()
}
func main() {
	/* server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	//路由
	router.Router()

	//http.HandleFunc("/", index)
	//http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
	f := config.Cfg.Viewer
	fmt.Print(f.UserName) */
	slice := make([]string, 1)
	// fmt.Sprintf("admin"+strings.Join(slice, "%v"), "mszlu")
	fmt.Println(slice)
	fmt.Printf(strings.Join(slice, "%v"), "mszlu")
	//s := "hello world hello world"
	//str := "wo"
	// var s = []string{"11","22","33"}
	// //将s中的⼦串连接成⼀个单独的字符串，⼦串之间⽤str分隔。
	// ret := strings.Join(s,"|")
	// fmt.Println(ret) //11|22|33

}
