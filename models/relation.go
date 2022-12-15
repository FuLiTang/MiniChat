package models

import (
	"gorm.io/gorm"
	"log"
	"miniChat/utils"
)

type Relation struct {
	gorm.Model
	A           string `json:"a"`
	B           string `json:"b"`
	Type        int    //0：还不是好友 1：好友状态，3：B删了A，4：A删了B，5：B拉黑了A，6：A拉黑了B，7：A发起好友申请，8：B发起了好友申请
	AGroup      string
	BGroup      string
	Err         int        //0：失败，1：成功，2：已创建，....
	ImageBase64 string     `json:"imageBase64"`
	BUser       *UserBasic `gorm:"-"`
}

func (t *Relation) Query() bool {
	r := &Relation{}
	if utils.MDb.Where("A = ? AND B = ? ", t.A, t.B).First(r).RowsAffected > 0 {
		return true
	}
	if utils.MDb.Where("B = ? AND A = ? ", t.A, t.B).First(r).RowsAffected > 0 {
		return true
	}
	return false

}
func (t *Relation) Create() error {
	//if t.Query().Err != 0 {
	//	return &Relation{
	//		Err: 2,
	//	}
	//}
	tx := utils.MDb.Begin()
	db := tx.Create(t)
	if db.Error != nil {
		tx.Rollback()
		log.Println(db.Error)
		return db.Error
	}
	tx.Commit()
	return nil
}

// QueryAll 关系状态type查询所有
func (t *Relation) QueryAll() ([]*Relation, *gorm.DB) {
	var r []*Relation
	allR := utils.MDb.Where("a = ? and type = ?", t.A, t.Type).Find(&r)
	return r, allR
}
func (t *Relation) QueryPage(i, p int) ([]*Relation, *gorm.DB) {
	var r []*Relation
	return r, utils.MDb.Where("a = ? and type = ?", t.A, t.Type).Offset((i - 1) * p).Limit(p).Find(&r)
}
