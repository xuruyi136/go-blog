package service

import (
	"errors"
	"go-blog/dao"
	"go-blog/models"
	"go-blog/utils"
)

func Login(userName, passwd string) (*models.LoginRes, error) {
	passwd = utils.Md5Crypt(passwd, "mszlu")
	// println("d+++++++++++++", userName, passwd)  1254c27e22c2ec22e62855fdf5878027  admin
	user := dao.GetUser(userName, passwd)
	if user == nil {
		return nil, errors.New("账号密码不正确")
	}
	uid := user.Uid

	//生成token,jwt技术进行生成令牌 A加密规则,B数据,过期时间,C对A+B进行加密，密匙
	token, err := utils.Award(&uid)
	if err != nil {
		return nil, errors.New("token未生成")
	}
	var userInfo models.UserInfo
	userInfo.Uid = user.Uid
	userInfo.UserName = user.UserName
	userInfo.Avatar = user.Avatar
	var lr = &models.LoginRes{
		token,
		userInfo,
	}
	return lr, nil
}
