// wordControl
package controllers

import (
	"WordSociety/wordlib"
	"github.com/astaxie/beego"
)

type WordController struct {
	beego.Controller
}

func (this *WordController) Get() {
	wordlib.WordDB.Open()
	defer wordlib.WordDB.Close()
	w := this.Ctx.Params[":word"]
	wo := wordlib.WordDB.GetWordObject(w)
	if wo != nil {
		wo.UpdateBaseInfo()
		this.Data["Word"] = *wo
		this.TplNames = "word.tpl"
	} else {
		this.Ctx.WriteString("Sorry，字库中未收录：" + w)
	}

}
