package models

import (
	"fmt"
	"gorm.io/gorm"
	"log"
	"miniChat/utils"
	"time"
)

type UserBasic struct {
	gorm.Model           //统一模型
	Name          string `form:"name" json:"name"`
	Password      string `form:"password" json:"password"`
	Phone         string `form:"phone" valid:"matches(^1[3-9]{1}\\d{9}$)" json:"phone"`
	Email         string `form:"email" valid:"email" json:"email"`
	Identity      string `form:"identity" json:"identity"` //唯一标识
	ClientIp      string
	ClientPort    string //客户窗口
	Salt          string
	LoginTime     time.Time //登录状况
	HeartbeatTime time.Time //心跳
	LoginOutTime  time.Time `gorm:"column:login_out_time" json:"login_out_time"` //登录时间
	IsLogout      bool      //下线状态
	DeviceInfo    string    //设备信息
	Err           string    //0失败
	HeadPortrait  string
}

func (t *UserBasic) TableName() string {
	return "user_basic"
}
func (u *UserBasic) FirstName() *UserBasic {
	user := &UserBasic{}
	utils.MDb.Where("name = ?", u.Name).First(user)
	return user
}
func (u *UserBasic) FirstPhone() *UserBasic {
	user := &UserBasic{}
	utils.MDb.Where("phone = ?", u.Phone).First(user)
	return user
}
func (u *UserBasic) FirstEmail() *UserBasic {
	user := &UserBasic{}
	utils.MDb.Where("email = ？", u.Email).First(user)
	return user
}
func (u *UserBasic) FirstNameAndPassword() *UserBasic {
	user := &UserBasic{}
	utils.MDb.Where("name = ? and password = ?", u.Name, u.Password).First(user)
	return user
}
func (u *UserBasic) QueryAndIdentity() bool {
	if utils.MDb.Where("name = ? and identity = ?", u.Name, u.Identity).First(u).RowsAffected > 0 {
		return true
	}
	return false
}

func GetUserList() []*UserBasic {
	data := make([]*UserBasic, 10)
	utils.MDb.Find(&data)
	for _, v := range data {
		fmt.Println(v)
	}
	return data
}
func CreateUser(user *UserBasic) *gorm.DB {
	tx := utils.MDb.Begin()
	db := tx.Create(&user)
	if db.Error != nil {
		tx.Rollback()
		log.Println("失败回滚！")
		return db
	}
	tx.Commit()
	return db
}
func DeleteUser(user *UserBasic) *gorm.DB {
	tx := utils.MDb.Begin()
	db := tx.Unscoped().Delete(&user)
	if db.Error != nil {
		tx.Rollback()
		log.Println("失败回滚！")
		return db
	}
	tx.Commit()
	return db
}

// UpdateUser 第一个选择id，第二个选择更新的字段 覆盖式更新字段
func UpdateUser(user *UserBasic, i interface{}) *gorm.DB {
	tx := utils.MDb.Begin()
	db := tx.Model(&user).Updates(i)
	if db.Error != nil {
		tx.Rollback()
		log.Println("失败回滚:", db.Error)
		return db
	}
	tx.Commit()
	return db
}

type State struct {
	UserName string
	StateId  string
}

func (s *State) Create() {
	tx := utils.MDb.Begin()
	db := tx.Create(&s)
	db.Commit()
}
func (s *State) Delete() {
	utils.MDb.Where("user_name = ?", s.UserName).Delete(&s)
}
func (s *State) Query() bool {
	if utils.MDb.Where("user_name = ? and state_id = ?", s.UserName, s.StateId).First(&s).RowsAffected > 0 {
		return true
	}
	return false
}
func (s *State) Update() bool {
	tx := utils.MDb.Begin()
	db := tx.Model(&s).Updates(&s)
	if db.Error != nil {
		tx.Rollback()
		log.Println("失败回滚！")
		return false
	}
	tx.Commit()
	return true
}
