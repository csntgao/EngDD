package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"labix.org/v2/mgo/bson"
)

type UploadController struct {
	beego.Controller
}

func (this *UploadController) Get() {
	s := `<!DOCTYPE html> 
	<html> 
		<body>
			<form  id="uploadForm"  method="post" enctype="MULTIPART/FORM-DATA" >
             	<input type="file" name="the_file" id="the_file" style="width: 1000px; margin-left:20px;" />
        		<input type="submit" value="上传" style="width: 60px;"     />
        	</form>
		</body>
	</html>`
	this.Ctx.WriteString(s)
	//this.TplNames = "upload.tpl"
}
func (this *UploadController) Post() {
	_, t, _ := this.GetFile("the_file")
	s := bson.NewObjectId().Hex() + "__" + t.Filename
	this.SaveToFile("the_file", "/Users/figo/go/src/WordSociety/static/upload/"+s)
	var r Response
	r.Suc = true
	r.Msg = "transfer Finished"
	r.Data = s
	b, _ := json.Marshal(r)
	this.Ctx.WriteString(string(b))

}
