package service

import (
	"gin/a/业务/dao"
	"gin/a/业务/model"
)

func CreatWarehouse(m model.Warehouse) (err error) {
	err = dao.Insertwarehouse(m)
	return err
}
func SearchiteamByWID(wid string) (err error, m model.Warehouse) {
	m, err = dao.Searchwid(wid)
	return
}
func ChangeuserIteamByWID(wid string, iid string) {
	dao.Changewarehouse(wid, iid)
}
func ChangeuserIwByWID(wid string, iid string, wid2 string) {
	dao.Changeiw(wid, iid, wid2)
}
