// CaseDatabase
package CaseLib

import (
	"fmt"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"strings"
	"time"
)

type CaseDatabase struct {
	MongoURL       string
	CaseCollection *mgo.Collection
	session        *mgo.Session
	LastDate       string
	LastID         int
}

func (this *CaseDatabase) alert(err error, msg string) {
	if err != nil {
		fmt.Println(msg)
		panic(err)
	}
}

func (this *CaseDatabase) Open() {
	var err error
	this.session, err = mgo.Dial("127.0.0.1")
	this.alert(err, "数据库打开错误！")
	this.CaseCollection = this.session.DB("Cases").C("Cases")

}

func (this *CaseDatabase) Close() {
	this.session.Close()

}

func (this *CaseDatabase) CreateNewCase() *AccidentCase {
	var accident AccidentCase
	s := time.Now().Format("20060102")
	if CaseDB.LastDate == s {
		CaseDB.LastID++
	} else {
		CaseDB.LastDate = s
		CaseDB.LastID = 1
	}

	s = fmt.Sprintf("%4d", CaseDB.LastID)
	s = strings.Replace(s, " ", "0", -1)

	accident.Id = bson.NewObjectId()
	accident.CaseNo = CaseDB.LastDate + s
	accident.Finished = false
	accident.State = "等待确认"

	return &accident

}

func (this *CaseDatabase) InsertCase(ACase *AccidentCase) {
	this.CaseCollection.Insert(ACase)

}

func (this *CaseDatabase) UpdateCase(ACase *AccidentCase) {
	this.CaseCollection.UpdateId(ACase.Id, ACase)

}

func (this *CaseDatabase) FindCases(query interface{}, result interface{}) error {
	err := this.CaseCollection.Find(query).All(result)
	//err := this.CaseCollection.Find(&bson.M{}).All(&Aresult)

	//fmt.Println(Aresult)

	this.alert(err, "FindCases错误！")
	return err

}

func (this *CaseDatabase) GetCasesByFilter(f string) (c []AccidentCase) {
	cases := []AccidentCase{}
	if f == "all" {
		this.FindCases(&bson.M{}, &cases)
	} else if f == "created" {
		this.FindCases(&bson.M{"state": "等待确认"}, &cases)
	} else if f == "confirmed" {
		this.FindCases(&bson.M{"state": "已经确认，等待拍照"}, &cases)
	} else if f == "phototaken" {
		this.FindCases(&bson.M{"state": "已经拍照，等待审理"}, &cases)
	} else if f == "handled" {
		this.FindCases(&bson.M{"state": "已审理，待回复"}, &cases)
	} else if f == "spot" {
		this.FindCases(&bson.M{"state": "申请现场"}, &cases)
	} else if f == "finished" {
		this.FindCases(&bson.M{"state": "已回应，处理完毕"}, &cases)
	}

	return cases

}

func (this *CaseDatabase) GetCaseByNo(no string) (b bool, acase *AccidentCase) {
	cases := []AccidentCase{}
	this.FindCases(&bson.M{"CaseNo": no}, &cases)

	b = len(cases) > 0
	if b {
		acase = &cases[0]

	} else {
		acase = nil
	}
	return
}
func (this *CaseDatabase) GetCasesByPhone(s string) (c []AccidentCase) {

	//var acase *AccidentCase
	//var su bool
	cases := []AccidentCase{}
	this.FindCases(&bson.M{"$or": []bson.M{bson.M{"usera_phone": s}, bson.M{"userb_phone": s}}}, &cases)

	/*
			if len(cases) > 0 {
			su = true
			acase = &cases[0]

		} else {
			su = false
			acase = nil
		}

		return su, acase
	*/
	return cases

}

var CaseDB CaseDatabase

func Init() {
	CaseDB.LastID = 0
	CaseDB.LastDate = time.Now().Format("20060102")

}
