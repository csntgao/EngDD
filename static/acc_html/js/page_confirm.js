
var Page_Confirm_Yes_Click=function(event,ui){
	$.getJSON(  "http://"+ServerIP+":"+ServerPort+"/case/confirm",
			  {
				  phone:MyMobile,
				  action:"yes",
				  engineNo:$("#Page_Confirm_EngineNo").val()
			  },
			  function(data,textStatus){
        		if (data.Suc){
					var str=data.Msg;
					var reg=new RegExp("\r\n","g");
					str= str.replace(reg,"<BR/>");
					$( "#Page_Confirm_Msg" ).html(str);
					$("#Page_Confirm_OpeateDiv").hide();
				}else
				{
					$( "#Page_Confirm_Msg" ).text(data.Msg);
					$("#Page_Confirm_OpeateDiv").hide();
					
				}
			  }
			);
}
var Page_Confirm_No_Click=function(event,ui){
	$.getJSON(  "http://"+ServerIP+":"+ServerPort+"/case/confirm",
			  {
				  phone:MyMobile,
				  action:"no",
				  engineNo:$("#Page_Confirm_EngineNo").val()
				  
			  },
			  function(data,textStatus){
        		if (data.Suc){
					var str=data.Msg;
					var reg=new RegExp("\r\n","g");
					str= str.replace(reg,"<BR/>");
					$( "#Page_Confirm_Msg" ).html(str);
					$("#Page_Confirm_OpeateDiv").hide();
				}else
				{
					$( "#Page_Confirm_Msg" ).text(data.Msg);
					$("#Page_Confirm_OpeateDiv").hide();
				}
			  }
			);
}
var Page_Confirm_Refresh=function(event,ui){
	$.getJSON(  "http://"+ServerIP+":"+ServerPort+"/case/confirm",
			  {
				  phone:MyMobile,
				  action:"refresh",
			  },
			  function(data,textStatus){
        		if (data.Suc){
					var str=data.Msg;
					var reg=new RegExp("\r\n","g");
					str= str.replace(reg,"<BR/>");
					$( "#Page_Confirm_Msg" ).html(str);
					$("#Page_Confirm_OpeateDiv").show();
				}else
				{
					$( "#Page_Confirm_Msg" ).text(data.Msg);
					$("#Page_Confirm_OpeateDiv").hide();
				}
			  }
			);
}

var Page_Confirm_PageShow=function(event,ui){
	if (MyMobile==""){
		$("#Page_Confirm_OpeateDiv").hide();
		//window.location.href="#page_Setup";
	}
	else
	{
		Page_Confirm_Refresh()
	}
	
};





var Page_Confirm_Init=function()
{
	
	
	$( "#page_Confirm" ).bind( "pageshow", Page_Confirm_PageShow);
	$( "#Page_Confirm_Yes" ).bind( "click", Page_Confirm_Yes_Click);
	$( "#Page_Confirm_No" ).bind( "click", Page_Confirm_No_Click);
	

};