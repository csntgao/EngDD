// JavaScript Document

var Coords={
	longitude:0,
    latitude:0
};

var CurDate = new Date(); 

var btnCreateCase_click=function(event, ui){
	
	$.getJSON( "http://"+ServerIP+":"+ServerPort+ "/case/create", 
	  {
		  CaseDateTime:CurDate.format("yyyyMMdhhmmss"),
		  CasePOS_X: Coords.longitude,
		  CasePOS_Y: Coords.latitude,
		  UserA_Phone:$("#Page_Alarm_UserA_Phone").val(),
		  UserA_LicenseNo:$("#Page_Alarm_UserA_LicenseNo").val(),
		  UserA_CarNo:$("#Page_Alarm_UserA_CarNo").val(),
		  UserB_Phone:$("#Page03_UserB_Phone").val(),
		  UserB_LicenseNo:$("#Page03_UserB_LicenseNo").val(),
		  UserB_CarNo:$("#Page03_UserB_CarNo").val()
		  
		  
	  },
	  function(data,textStatus){
        if (data.Suc){
			MyMobile=$("#Page_Alarm_UserA_Phone").val();
			$("#MyPhoneNo").val(MyMobile);
			$("#Page_Alarm_UserA_Phone").val("");
		    $("#Page_Alarm_UserA_LicenseNo").val("");
		  	$("#Page_Alarm_UserA_CarNo").val("");
		  	$("#Page03_UserB_Phone").val("");
		  	$("#Page03_UserB_LicenseNo").val("");
		  	$("#Page03_UserB_CarNo").val("");
			
			$( "#popupMessage" ).popup();
			$( "#popupMessageText" ).text(data.Msg);
			$( "#popupMessage" ).popup("open");
			
			
		
			
		}else{
			$( "#popupMessage" ).popup();
			$( "#popupMessageText" ).text("失败，原因为："+data.Msg);
			$( "#popupMessage" ).popup("open");
		}
			
      });
}

var locationSuccess=function(position){ 
      Coords = position.coords; 
	  var s= "【"+Math.round(Coords.longitude*10000)/10000+","+Math.round(Coords.latitude*10000)/10000+"】"  ;
	  $("#currentPosition1").text("案发地点："+s);
	  $("#currentPosition2").text("案发地点："+s);
	 
	};


var locationFailure=function(position){ 
     
	  $("#currentPosition1").text("案发地点：不可知");
	  $("#currentPosition2").text("案发地点：不可知");
	 
	};
var Page_Alarm_Init=function()
{
	
	$("#currentDateTime1").text("案发时间："+CurDate.format("yyyy年MM月dd日hh小时mm分"));
	$("#currentDateTime2").text("案发时间："+CurDate.format("yyyy年MM月dd日hh小时mm分"));
	$( "#btnCreateCase" ).bind( "click", btnCreateCase_click);
		
		
	if (navigator.geolocation) { 
    	navigator.geolocation.getCurrentPosition(locationSuccess, locationFailure,{ 
       		enableHighAcuracy: true, 
       		timeout: 5000, 
       		maximumAge: 3000 
    	});
	}
	else
	{
		 $("#currentPosition1").text("无GPS信息");
	}

};