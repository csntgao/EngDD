
var Page_CommitState_Execute_Click=function(){
	$.getJSON( "http://"+ServerIP+":"+ServerPort+ "/case/photo", 
	  {
		  phone:MyMobile,
		  action:"commit",
		  acc:acctexts[i]
		  
		  
	  },
	  function(data,textStatus){
        if (data.Suc){
			
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

var Page_CommitState_PageShow=function()
{
	var s='您选择的事故类型：<BR/>';
	s+=acctexts[i];
	$("#Page_CommitState_Msg").html(s);
}


var Page_CommitState_Init=function()
{
	
	
	$( "#page_CommitState" ).bind( "pageshow", Page_CommitState_PageShow);
	$( "#Page_CommitState_Execute" ).bind( "click", Page_CommitState_Execute_Click);
	//$( "#Page_Confirm_No" ).bind( "click", Page_Confirm_No_Click);
	

};