// handleCase
package models

import (
	"WordSociety/CaseLib"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

func TakePhoto_Init(m string) (b bool, msg string) {
	CaseLib.CaseDB.Open()
	defer CaseLib.CaseDB.Close()
	cases := CaseLib.CaseDB.GetCasesByPhone(m)
	for _, ca := range cases {
		if ca.State == "已经确认，等待拍照" {
			return true, ca.State
		}
	}
	return false, "当前没有需要拍照的案件！"

}

func TakePhoto_Commit(m string, r url.Values) (b bool, msg string) {
	CaseLib.CaseDB.Open()
	defer CaseLib.CaseDB.Close()
	cases := CaseLib.CaseDB.GetCasesByPhone(m)
	for _, ca := range cases {
		if ca.State == "已经确认，等待拍照" {

			if ca.UserA_Phone == m {
				ca.UserA.AccIndex, _ = strconv.Atoi(r.Get("accindex"))
				ca.UserA.Pict0 = r.Get("p0")
				ca.UserA.Pict1 = r.Get("p1")
				ca.UserA.Pict2 = r.Get("p2")
				ca.UserA.Pict3 = r.Get("p3")
				ca.UserA.Pict4 = r.Get("p4")

				ca.UserA.PhotoCommitted = true

			} else if ca.UserB_Phone == m {
				ca.UserB.AccIndex, _ = strconv.Atoi(r.Get("accindex"))
				ca.UserB.Pict0 = r.Get("p0")
				ca.UserB.Pict1 = r.Get("p1")
				ca.UserB.Pict2 = r.Get("p2")
				ca.UserB.Pict3 = r.Get("p3")
				ca.UserB.Pict4 = r.Get("p4")
				ca.UserB.PhotoCommitted = true
			}
			if ca.UserA.PhotoCommitted && ca.UserB.PhotoCommitted {
				ca.State = "已经拍照，等待审理"
			}
			CaseLib.CaseDB.UpdateCase(&ca)
			return true, "拍照信息已接收！"
		}
	}
	return false, "未找到数据"

}

func TakePhoto(r url.Values) (b bool, msg string) {
	fmt.Println("TakePhotoes")
	m := r.Get("phone")
	action := r.Get("action")
	if action == "init" {
		b, msg = TakePhoto_Init(m)
		return
	}
	if action == "commit" {
		b, msg = TakePhoto_Commit(m, r)
		return
	}
	return false, "未处理的分支Action=" + action

}

func QueryCase(r url.Values) (b bool, msg string, resultType string, ca *CaseLib.AccidentCase) {
	fmt.Println("QueryCase")
	CaseLib.CaseDB.Open()
	defer CaseLib.CaseDB.Close()
	m := r.Get("phone")
	Cases := CaseLib.CaseDB.GetCasesByPhone(m)
	ca = nil
	b = true
	msg = "目前没有与您相关的案件！"
	resultType = "normal"
	for _, c := range Cases {
		ca = &c
		msg = c.Description()
		if c.State == "已审理，待回复" {
			resultType = "handleresult"
		}
		if c.State == "已回应，处理完毕" {
			resultType = "normalfinished"

		} else if c.State == "申请现场" {
			resultType = "abnormalfinished"
		}
		break

	}
	return

}
func ConfirmCase_Refresh(phone string) (b bool, msg string) {

	CaseLib.CaseDB.Open()
	defer CaseLib.CaseDB.Close()

	Cases := CaseLib.CaseDB.GetCasesByPhone(phone)
	if len(Cases) >= 1 {
		for _, c := range Cases {
			if (c.UserB_Phone == phone) && (c.State == "等待确认") {
				b = true
				msg = c.Description()
				return
			}
		}

	}
	b = false
	msg = "当前没有需要您确认的案件！"
	return
}

func CheckEngineNo(EngineNo string, CarNo string) bool {
	r := false
	if strings.ToUpper(EngineNo) == "PASS" {
		r = true
	}
	return r
}
func ConfirmCase_Yes(phone string, engineNo string) (b bool, msg string) {
	fmt.Println("ConfirmCase_Yes")
	CaseLib.CaseDB.Open()
	defer CaseLib.CaseDB.Close()
	fmt.Println(phone)
	fmt.Println(engineNo)

	Cases := CaseLib.CaseDB.GetCasesByPhone(phone)
	if len(Cases) >= 1 {
		for _, c := range Cases {
			if (c.UserB_Phone == phone) && (c.State == "等待确认") {
				if CheckEngineNo(engineNo, c.UserB_CarNo) {
					c.State = "已经确认，等待拍照"
					CaseLib.CaseDB.UpdateCase(&c)
					b = true
					msg = "案件已经确认，可以进行拍照了"
					return
				} else {
					b = false
					msg = "发动机号码和车牌号不对应，请重试！"
					return
				}
			}
		}

	}
	b = false
	msg = "案件不存在，当前没有需要您确认的案件！"
	return

}

func ConfirmCase_No(phone string, engineNo string) (b bool, msg string) {
	fmt.Println("ConfirmCase_Yes")
	CaseLib.CaseDB.Open()
	defer CaseLib.CaseDB.Close()
	fmt.Println(phone)
	fmt.Println(engineNo)

	Cases := CaseLib.CaseDB.GetCasesByPhone(phone)
	if len(Cases) >= 1 {
		for _, c := range Cases {
			if (c.UserB_Phone == phone) && (c.State == "等待确认") {
				if CheckEngineNo(engineNo, c.UserB_CarNo) {
					c.State = "已经拒绝"
					CaseLib.CaseDB.UpdateCase(&c)
					b = true
					msg = "案件已经拒绝"
					return
				} else {
					b = false
					msg = "发动机号码和车牌号不对应，请重试！"
					return
				}
			}
		}

	}
	b = false
	msg = "案件不存在！"
	return

}

func ConfirmCase(r url.Values) (b bool, msg string) {
	fmt.Println("ConfirmCase")
	action := r.Get("action")
	phone := r.Get("phone")
	if action == "refresh" {
		b, msg = ConfirmCase_Refresh(phone)

	}
	if action == "yes" {
		engineNo := r.Get("engineNo")
		b, msg = ConfirmCase_Yes(phone, engineNo)
	}
	if action == "no" {
		engineNo := r.Get("engineNo")
		b, msg = ConfirmCase_No(phone, engineNo)
	}
	return
}

func CreateCase(r url.Values) (b bool, msg string) {
	fmt.Println("CreateCase")
	CaseLib.CaseDB.Open()
	defer CaseLib.CaseDB.Close()
	ACase := CaseLib.CaseDB.CreateNewCase()
	ACase.CaseDateTime = r.Get("CaseDateTime")
	ACase.CasePOS_X = r.Get("CasePOS_X")
	ACase.CasePOS_Y = r.Get("CasePOS_Y")
	ACase.UserA_Phone = r.Get("UserA_Phone")
	ACase.UserA_LicenseNo = r.Get("UserA_LicenseNo")
	ACase.UserA_CarNo = r.Get("UserA_CarNo")
	ACase.UserB_Phone = r.Get("UserB_Phone")
	ACase.UserB_LicenseNo = r.Get("UserB_LicenseNo")
	ACase.UserB_CarNo = r.Get("UserB_CarNo")

	CaseLib.CaseDB.InsertCase(ACase)
	fmt.Println(ACase)
	return true, "报警请求已经接受，正在等待对方进行确认！案件ID为" + ACase.CaseNo

}

func GetGridCaseListByFilter(f string) (b bool, msg string, data []CaseLib.GridCase) {
	fmt.Println("GetGridCaseListByFilter:filter=" + f)
	CaseLib.CaseDB.Open()
	defer CaseLib.CaseDB.Close()
	datas := []CaseLib.GridCase{}
	ads := CaseLib.CaseDB.GetCasesByFilter(f)
	for _, c := range ads {

		datas = append(datas, c.ToGridCase())

	}
	return true, "", datas
}
func GetCaseByNo(no string) (b bool, msg string, acase *CaseLib.AccidentCase) {
	fmt.Println("GetCaseByNo")
	CaseLib.CaseDB.Open()
	defer CaseLib.CaseDB.Close()
	b, acase = CaseLib.CaseDB.GetCaseByNo(no)
	msg = ""
	return
}

func SetDuty(caseno string, duty int) (b bool, msg string) {
	fmt.Printf("SetDuty duty=%d ,caseno=%s\n", duty, caseno)
	fmt.Println("SetDuty duty=" + fmt.Sprintf("%d", duty))
	CaseLib.CaseDB.Open()
	defer CaseLib.CaseDB.Close()
	s, acase := CaseLib.CaseDB.GetCaseByNo(caseno)
	b = s
	if s {
		acase.UserA.Duty = duty
		acase.UserB.Duty = 100 - duty
		acase.State = "已审理，待回复"
		CaseLib.CaseDB.UpdateCase(acase)
		msg = "设置责任成功！"
	} else {
		msg = "不存在指定的案件！"
	}
	return

}

func ResponseResult(r url.Values) (b bool, msg string) {
	m := r.Get("phone")
	a := r.Get("action")
	fmt.Printf("ResponseResult m:%s a:%s\n", m, a)
	CaseLib.CaseDB.Open()
	defer CaseLib.CaseDB.Close()
	cases := CaseLib.CaseDB.GetCasesByPhone(m)
	for _, ca := range cases {
		if ca.State == "已审理，待回复" {
			if ca.UserA_Phone == m {
				if a == "refuse" {
					ca.UserA.Response = "申请现场"
					ca.State = "申请现场"
				} else if a == "acknowledge" {
					ca.UserA.Response = "结果确认"
				}
			} else if ca.UserB_Phone == m {
				if a == "refuse" {
					ca.UserB.Response = "申请现场"
				} else if a == "acknowledge" {
					ca.UserB.Response = "结果确认"
				}
			}
			if ca.UserA.Response == "结果确认" && ca.UserB.Response == "结果确认" {
				ca.State = "已回应，处理完毕"
			}
			CaseLib.CaseDB.UpdateCase(&ca)
			return true, ""
		}
	}
	return false, ""

}
