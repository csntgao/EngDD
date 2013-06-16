package main

import (
	"WordSociety/controllers"
	"fmt"
	"github.com/astaxie/beego"
	"os"
	"strings"
)

func main() {

	var ConfigPath = ""

	//*
	fmt.Println("WebApp WordSociety has started!")
	if strings.Contains(os.Args[0], "gosublime") {
		fmt.Println("WebApp WordSociety running in gosublime")
		ConfigPath = "/Users/figo/go/src/WordSociety/"

	} else {
		fmt.Println("WebApp WordSociety running in normal mode")
	}

	beego.Router("/", &controllers.MainController{})
	beego.Router("/words/:word(.+)", &controllers.WordController{})
	beego.Router("/case/:opearte(.+)", &controllers.CaseController{})
	beego.Router("/workstation/:type(.+)", &controllers.WorkStationController{})
	beego.Router("/upload", &controllers.UploadController{})
	beego.SetStaticPath("/nojam", ConfigPath+"static")
	beego.SetStaticPath("/easyui", ConfigPath+"static/easyui")
	beego.Run()
	/*
	 */

}
