package models

import "go-blog/config"

//首页数据
type HomeResponse struct {
	config.Viewer
	Categorys []Category
	Posts     []PostMore
	Total     int
	Page      int
	Pages     []int
	PageEnd   bool
}