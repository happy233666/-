package service

import (
	"gin/a/业务/dao"
	"gin/a/业务/model"
)

func Creatuser(u model.User) (err error) {
	err = dao.Insertuser(u)
	return err
}
func SearchuserByID(id string) (err error, u model.User) {
	u, err = dao.Searchid(id)
	return
}
func ChangeuserPasswordByID(id string, password string) {
	dao.ChangeDB(id, password)
}
