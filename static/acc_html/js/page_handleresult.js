

function Page_HandleResult_Confirm_Click()
{
	$.getJSON(  "http://"+ServerIP+":"+ServerPort+"/case/responseresult",
			  {
				  phone:MyMobile,
				  action:"acknowledge"
			  },
			  function(data,textStatus){
        		Page_Status_DoRefresh();
			  }
			);
}

function Page_HandleResult_Refuse_Click()
{
	$.getJSON(  "http://"+ServerIP+":"+ServerPort+"/case/responseresult",
			  {
				  phone:MyMobile,
				  action:"refuse"
			  },
			  function(data,textStatus){
        		Page_Status_DoRefresh();
			  }
			);
}



var Page_HandleResult_Init=function()
{
	
	
	$( "#Page_HandleResult_ButtonConfirm" ).bind( "click", Page_HandleResult_Confirm_Click);
	$( "#Page_HandleResult_ButtonRefuse" ).bind( "click", Page_HandleResult_Refuse_Click);

};