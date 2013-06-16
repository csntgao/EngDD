// caseControl

package controllers

import (
	"WordSociety/models"
	"encoding/json"
	"github.com/astaxie/beego"
)

type CaseController struct {
	beego.Controller
}

func (this *CaseController) Get() {
	o := this.Ctx.Params[":opearte"]
	/*
		s := ""
		for k, _ := range this.Input() {
			s = s + "key:" + k + " " + "value:" + this.Input().Get(k)
		}

		this.Ctx.WriteString(o + " " + s)
	*/
	if o == "create" {
		var res Response
		res.Suc, res.Msg = models.CreateCase(this.Input())
		b, _ := json.Marshal(res)
		this.Ctx.WriteString(string(b))

	} else if o == "confirm" {
		var res Response
		res.Suc, res.Msg = models.ConfirmCase(this.Input())
		b, _ := json.Marshal(res)
		this.Ctx.WriteString(string(b))
	} else if o == "status" {
		var res Response
		res.Suc, res.Msg, res.ResultType, res.Data = models.QueryCase(this.Input())
		b, _ := json.Marshal(res)
		this.Ctx.WriteString(string(b))

	} else if o == "photo" {
		var res Response
		res.Suc, res.Msg = models.TakePhoto(this.Input())
		b, _ := json.Marshal(res)
		this.Ctx.WriteString(string(b))
	} else if o == "responseresult" {
		var res Response
		res.Suc, res.Msg = models.ResponseResult(this.Input())
		b, _ := json.Marshal(res)
		this.Ctx.WriteString(string(b))

	}

}
