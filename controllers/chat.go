package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"miniChat/models"
	"strconv"
)

func LinkMan(c *gin.Context) {
	user := &models.UserBasic{}
	c.Bind(user)
	if !user.QueryAndIdentity() {
		c.JSON(200, gin.H{
			"code":    "0",
			"message": "系统错误！",
		})
		return
	}
	r := &models.Relation{
		A:    user.Name,
		B:    "",
		Type: 1,
	}
	rs, db := r.QueryAll()
	if db.RowsAffected == 0 {
		c.JSON(200, gin.H{
			"code":    "1",
			"message": "暂无联系人！",
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    "1",
		"message": strconv.FormatInt(db.RowsAffected, 10),
		"data":    rs,
	})
	return
}

// LinkManPage 所有联系人
func LinkManPage(c *gin.Context) {
	user := &models.UserBasic{}
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		log.Println(err)
	}
	c.Bind(user)
	if !user.QueryAndIdentity() {
		c.JSON(200, gin.H{
			"code":    "0",
			"message": "系统错误！",
		})
		return
	}
	r := &models.Relation{
		A:    user.Name,
		B:    "",
		Type: 1,
	}
	rs, db := r.QueryPage(page, 20)
	_, pages := r.QueryAll()
	var pagesI int64
	if pages.RowsAffected%20 == 0 {
		pagesI = pages.RowsAffected / 20
	} else {
		pagesI = pages.RowsAffected / 21
	}
	if db.RowsAffected == 0 {
		c.JSON(200, gin.H{
			"code":    "1",
			"message": "暂无联系人！",
		})
		return
	}
	c.JSON(200, gin.H{
		"code":  "1",
		"data":  rs,
		"pages": strconv.FormatInt(pagesI, 10),
		"all":   strconv.FormatInt(pages.RowsAffected, 10),
	})
	return
}
