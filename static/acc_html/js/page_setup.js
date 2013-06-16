// JavaScript Document
var Page_Setup_OK_Click=function(event,ui){
	MyMobile=$("#MyPhoneNo").val();
	ServerIP=$("#ServerIP").val();
	ServerPort=$("#ServerPort").val();
	
};

var Page_Setup_PageShow=function(event,ui){
	$("#MyPhoneNo").val(MyMobile);
	$("#ServerIP").val(ServerIP);
	$("#ServerPort").val(ServerPort);
	
};
	

var Page_Setup_Init=function()
{
	
	$( "#Page_Setup_OK" ).bind( "click", Page_Setup_OK_Click);
	$( "#MyPhoneNo" ).bind( "onchange", Page_Setup_OK_Click);
	$( "#page_Setup" ).bind( "pagehide", Page_Setup_OK_Click);
	$( "#page_Setup" ).bind( "pageshow", Page_Setup_PageShow);
	
	
	

};