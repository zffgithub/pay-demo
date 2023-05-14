package system

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

type CustomerInfo struct {
	Id         int
	CName      string
	CPhone     string
	SumMoney   string
	Lesson     string
	Ip         string
	CreateTime string
	UpdateTime string
}

const CUSTOMER_INFO = "customer_info"

func GetRoleLenByMap(params map[string]string) int {
	o := orm.NewOrm()
	qs := o.QueryTable(CUSTOMER_INFO)
	for k, v := range params {
		if len(v) > 0 {
			qs = qs.Filter(k, v)
		}
	}
	cnt, err := qs.Count()
	if err != nil {
		logs.Error("get role len by map fail: ", err)
	}
	return int(cnt)
}

func InsertItem(roleInfo CustomerInfo) bool {
	o := orm.NewOrm()
	_, err := o.Insert(&roleInfo)
	if err != nil {
		logs.Error("insert role fail: ", err)
		return false
	}
	return true
}

// func DeleteRoleByRoleUid(roleUid string) bool {
// 	o := orm.NewOrm()
// 	_, err := o.QueryTable(CUSTOMER_INFO).Filter("role_uid", roleUid).Delete()
// 	if err != nil {
// 		logs.Error("delete role by role uid fail: ", err)
// 		return false
// 	}
// 	return true
// }

// func UpdateRoleInfo(roleInfo CustomerInfo) bool {
// 	o := orm.NewOrm()
// 	_, err := o.Update(&roleInfo)

// 	if err != nil {
// 		logs.Error("update role info fail: ", err)
// 		return false
// 	}
// 	return true
// }
