package controllers

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"miniChat/models"
	"miniChat/utils"
	"strconv"
)

func RegisterNameAndPassword(c *gin.Context) {
	us := &models.UserBasic{}
	c.Bind(us)
	user := us.FirstName()
	if user.Name != us.Name {
		c.JSON(200, gin.H{
			"code":    "0",
			"message": "用户名密码错误！",
		})
		return
	}
	if !utils.ValidPassword(us.Password, user.Salt, user.Password) {
		c.JSON(200, gin.H{
			"code":    "0",
			"message": "用户名密码错误！",
		})
		return
	}
	//token登录
	user.Identity = utils.SaltTime()
	up := &models.UserBasic{ //更新选择
		Identity: user.Identity,
	}
	if models.UpdateUser(user, up).RowsAffected == 0 {
		c.JSON(200, gin.H{
			"code":    "0",
			"message": "系统出现错误！",
		})
		return
	}
	user.Password = ""
	fmt.Println(user)
	s := &models.State{UserName: user.Name, StateId: user.Identity}
	s.Create()
	c.JSON(200, gin.H{
		"code":    "1", //0失败，1成功
		"message": "登陆成功！",
		"data":    user,
	})
	return
}

func CreateUser(c *gin.Context) {
	user := &models.UserBasic{}
	c.Bind(user)
	rePassword := c.PostForm("rePassword")
	if user.Name == "" || user.Password == "" || rePassword == "" {
		c.JSON(200, gin.H{
			"message": "注册信息不能为空！",
		})
		return
	}
	if user.Password != rePassword {
		c.JSON(-1, gin.H{
			"message": "两次密码不一致",
		})
		return
	}
	if user.FirstName().Name == user.Name {
		c.JSON(200, gin.H{
			"message": "用户名已被注册",
		})
		return
	}
	user.Salt = fmt.Sprintf("%06d", rand.Int31())
	user.Identity = utils.SaltTime()
	user.Password = utils.MakePassword(user.Password, user.Salt) //加密
	if models.CreateUser(user).RowsAffected > 0 {
		c.JSON(200, gin.H{
			"message": "新增用户成功",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "新增用户失败！",
	})
}
func DeleteUser(c *gin.Context) {
	user := &models.UserBasic{}
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(200, gin.H{
			"message": "用户ID不能为空！",
		})
		return
	}
	user.ID = uint(id)
	if models.DeleteUser(user).RowsAffected > 0 {
		c.JSON(200, gin.H{
			"message": "删除用户成功！",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "删除用户失败！",
	})
	return
}
func UpdateUser(c *gin.Context) {
	user := &models.UserBasic{}
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(200, gin.H{
			"message": "用户ID不能为空！",
		})
		return
	}
	user.ID = uint(id)
	if err := c.Bind(user); err != nil {
		log.Println(err)
	}
	if _, err := govalidator.ValidateStruct(user); err != nil { //格式验证
		c.JSON(200, gin.H{
			"message": "修改参数不匹配！",
		})
		return
	}
	up := &models.UserBasic{
		Name:     user.Name,
		Password: user.Password,
		Phone:    user.Phone,
		Email:    user.Email,
	}
	if models.UpdateUser(user, up).RowsAffected > 0 {
		c.JSON(200, gin.H{
			"message": "修改用户信息成功！",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "修改用户信息失败！",
	})
	return

}
