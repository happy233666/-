package dao

import (
	"fmt"
	"gin/a/业务/model"
)

func Insertwarehouse(m model.Warehouse) (err error) {
	if m.Wid == "" || m.Iid == "" {
		return nil
	}
	_, err = DB.Exec("insert into warehouse(wid,iid)values (?,?)", m.Wid, m.Iid)
	return err
}
func Searchwid(wid string) (m model.Warehouse, err error) {
	row := DB.QueryRow("select wid,iid from warehouse where wid = ?", wid)
	if err = row.Err(); row.Err() != nil {
		return
	}
	err = row.Scan(&m.Wid, &m.Iid)
	return
}
func Changewarehouse(wid string, iid string) {
	sql := "update warehouse set iid =? where wid = ?"
	res, err := DB.Exec(sql, iid, wid)
	if err != nil {
		fmt.Println("exec failed,", err)
		return
	}
	row, err := res.RowsAffected()
	if err != nil {
		fmt.Println("rows failed", err)
	}

	fmt.Println("update succ:", row)

}
func Changeiw(wid string, iid string, wid2 string) {
	sql := "update warehouse set iid =? where wid = ?"
	res, err := DB.Exec(sql, iid, wid2)
	if err != nil {
		fmt.Println("exec failed,", err)
		return
	}
	row, err := res.RowsAffected()
	if err != nil {
		fmt.Println("rows failed", err)
	}

	fmt.Println("update succ:", row)

}
