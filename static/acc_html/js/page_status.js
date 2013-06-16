
var Page_Status_DoRefresh=function()
{
	$.getJSON(  "http://"+ServerIP+":"+ServerPort+"/case/status",
			  {
				  phone:MyMobile
			  },
			  function(data,textStatus){
        		if (data.Suc){
					var str=data.Msg;
					if (data.ResultType=="normal")
					{
						window.location.href="#page_Status"
						var reg=new RegExp("\r\n","g");
						str= str.replace(reg,"<BR/>");
					  	$( "#Page_Status_Msg" ).html(str);
					} else if (data.ResultType=="handleresult"){
						$("#Page_HandleResult_CaseNo").html("案件号："+data.Data.CaseNo+"<BR/>");
						$("#Page_HandleResult_UserA").html("驾驶员："+data.Data.UserA_Phone+" 车号 "
																	+data.Data.UserA_CarNo+" "
																	+data.Data.UserA.Duty
																	+"%的责任 "
																	+data.Data.UserA.Response
																	+"<BR/>");
						$("#Page_HandleResult_UserB").html("驾驶员："+data.Data.UserB_Phone+" 车号 "
																	+data.Data.UserB_CarNo+" "
																	+data.Data.UserB.Duty
																	+"%的责任 "
																	+data.Data.UserB.Response
																	+"<BR/>");


						window.location.href="#page_HandleResult"
					} else if (data.ResultType=="normalfinished"){
						window.location.href="#page_Finished"
					} else if (data.ResultType=="abnormalfinished"){
						window.location.href="#page_Abfinished"
					}

				}else
				{
					//$( "#Page_Status_Msg" ).text("无法从服务器端取得数据！");
					$( "#Page_Status_Msg" ).html(data.Msg);
				}
			  }
			);
	
}
	
	
	

var Page_Status_PageShow=function(event,ui){
	if (MyMobile==""){
		
		window.location.href="#page_Setup";
	}
	else
	{
		Page_Status_DoRefresh()
	}
	
};





var Page_Status_Init=function()
{
	
	
	$( "#page_Status" ).bind( "pageshow", Page_Status_PageShow);
	$( "#Page_Status_Refresh" ).bind( "click", Page_Status_DoRefresh);

};