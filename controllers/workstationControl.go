package controllers

import (
	"WordSociety/models"
	//"WordSociety/CaseLib"
	"encoding/json"
	"github.com/astaxie/beego"
	"strconv"
)

type WorkStationController struct {
	beego.Controller
}

func (this *WorkStationController) Get() {
	o := this.Ctx.Params[":type"]
	/*
		s := ""
		for k, _ := range this.Input() {
			s = s + "key:" + k + " " + "value:" + this.Input().Get(k)
		}

		this.Ctx.WriteString(o + " " + s)
	*/
	if o == "list" {
		var f = this.Input().Get("f")

		var res Response
		res.Suc, res.Msg, res.Data = models.GetGridCaseListByFilter(f)

		b, _ := json.Marshal(res)
		this.Ctx.WriteString(string(b))

	} else if o == "detail" {
		var caseno = this.Input().Get("caseno")
		var res Response
		res.Suc, res.Msg, res.Data = models.GetCaseByNo(caseno)
		b, _ := json.Marshal(res)
		this.Ctx.WriteString(string(b))
	} else if o == "setduty" {
		var caseno = this.Input().Get("caseno")
		var value = this.Input().Get("duty")
		v, _ := strconv.Atoi(value)
		var res Response
		res.Suc, res.Msg = models.SetDuty(caseno, v)
		b, _ := json.Marshal(res)
		this.Ctx.WriteString(string(b))
	}

}
