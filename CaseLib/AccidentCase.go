// AccidentCase
package CaseLib

import (
	"labix.org/v2/mgo/bson"
)

type GridCase struct {
	CaseNo   string "caseNo"
	DateTime string "dateTime"
	APhone   string "aPhone"
	BPhone   string "bPhone"
}

type UserData struct {
	AccIndex       int    "accindex"
	PhotoCommitted bool   "photoCommitted"
	Duty           int    "duty"
	Response       string "response"
	Pict0          string
	Pict1          string
	Pict2          string
	Pict3          string
	Pict4          string
}

type AccidentCase struct {
	Id              bson.ObjectId "_id"
	CaseNo          string        "CaseNo"
	CaseDateTime    string
	CasePOS_X       string
	CasePOS_Y       string
	UserA_Phone     string
	UserA_LicenseNo string
	UserA_CarNo     string
	UserB_Phone     string
	UserB_LicenseNo string
	UserB_CarNo     string
	State           string
	UserA           UserData
	UserB           UserData
	Finished        bool
}

func (this *AccidentCase) ToGridCase() GridCase {
	var c GridCase
	c.CaseNo = this.CaseNo
	c.DateTime = this.CaseDateTime
	c.APhone = this.UserA_Phone
	c.BPhone = this.UserB_Phone
	return c
}
func (this *AccidentCase) Description() string {
	s := "案件 编 号：" + this.CaseNo + "\r\n"
	s += "报案人电话：" + this.UserA_Phone + "\r\n"
	s += "报案人驾照：" + this.UserA_LicenseNo + "\r\n"
	s += "报案人车牌：" + this.UserA_CarNo + "\r\n"
	s += "涉案人电话：" + this.UserB_Phone + "\r\n"
	s += "涉案人驾照：" + this.UserB_LicenseNo + "\r\n"
	s += "涉案人车牌：" + this.UserB_CarNo + "\r\n"
	s += "当前 状 态：" + this.State + "\r\n"

	return s
}
