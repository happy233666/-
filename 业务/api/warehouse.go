package api

import (
	"database/sql"
	"gin/a/业务/model"
	"gin/a/业务/service"
	"gin/a/业务/util"
	"github.com/gin-gonic/gin"
	"log"
)

func NewWarehouse(c *gin.Context) {
	wid := c.PostForm("wid")
	iid := c.PostForm("iid")
	err := service.CreatWarehouse(model.Warehouse{
		Wid: wid,
		Iid: iid,
	})
	if wid == "" || iid == "" {
		util.RespParamerror(c)
		return
	}
	if err != nil {
		util.RestpInternalErr(c)
		return
	}
	util.Respok(c)
}
func ChangeWarehouse(c *gin.Context) {
	wid := c.PostForm("wid")
	iid := c.PostForm("iid")
	err, m := service.SearchiteamByWID(wid)
	if wid == "" || iid == "" {
		util.RespParamerror(c)
		return
	}
	if err != nil {
		if err == sql.ErrNoRows {
			util.Normalerr(c, 3, "物品不存在")
		} else {
			log.Printf("search user error: %v", err)
			util.RestpInternalErr(c)
			return
		}
		return
	}
	if m.Wid != wid {
		util.Normalerr(c, 6, "没有")
		return
	}
	service.ChangeuserIteamByWID(wid, iid)
	util.Respok(c)
}
func DeleteIteam(c *gin.Context) {
	wid := c.PostForm("wid")
	iid := c.PostForm("iid")
	err, m := service.SearchiteamByWID(wid)
	if wid == "" || iid == "" {
		util.RespParamerror(c)
		return
	}
	if err != nil {
		if err == sql.ErrNoRows {
			util.Normalerr(c, 3, "物品不存在")
		} else {
			log.Printf("search user error: %v", err)
			util.RestpInternalErr(c)
			return
		}
		return
	}
	if m.Wid != wid {
		util.Normalerr(c, 6, "没有")
		return
	}
	service.ChangeuserIteamByWID(wid, "nil")
	util.Respok(c)
}
func ChangeIteam(c *gin.Context) {
	wid := c.PostForm("wid")
	iid := c.PostForm("iid")
	wid2 := c.PostForm("wid2")
	err, m := service.SearchiteamByWID(wid)
	if wid == "" || iid == "" {
		util.RespParamerror(c)
		return
	}
	if err != nil {
		if err == sql.ErrNoRows {
			util.Normalerr(c, 3, "物品不存在")
		} else {
			log.Printf("search user error: %v", err)
			util.RestpInternalErr(c)
			return
		}
		return
	}
	if m.Wid != wid {
		util.Normalerr(c, 6, "没有")
		return
	}
	service.ChangeuserIwByWID(wid, iid, wid2)
	util.Respok(c)
}
