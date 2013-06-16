// JavaScript Document
	var i=21
	var acctexts=new Array()
	acctexts[0]="追尾";
	acctexts[1]="逆行";
	acctexts[2]="倒车";
	acctexts[3]="溜车";
	acctexts[4]="开关车门";
	acctexts[5]="违反交通信号的（之一“闯红灯”）";
	acctexts[6]="违反交通信号的（之二“掉头”）";
	acctexts[7]="违反交通信号的（之三“驶入专用车道”） ";
	acctexts[8]="违反交通信号的（之四“未按指示方向通行”）";
	acctexts[9]="违反交通信号的（之五“驶入禁行道路”）";
	acctexts[10]="违反交通信号的（之六“未服从交通警察指挥”）";
	acctexts[11]="未按规定让行的（之一“未让特种车”）";
	acctexts[12]="未按规定让行的（之二“变更车道”）"; 
	acctexts[13]="未按规定让行的（之三“障碍路段未让行”）";
	acctexts[14]="未按规定让行的（之四“障碍路段未让行”）";
	acctexts[15]="未按规定让行的（之五“坡路未让行”）";  
	acctexts[16]="未按规定让行的（之六“坡路未让行”）";
	acctexts[17]="未按规定让行的（之七“狭窄山路未让行”）"; 
	acctexts[18]="未按规定让行的（之八“超正在左转弯车”）";
	acctexts[19]="未按规定让行的（之九“超正在掉头车”）";

acctexts[20]="未按规定让行的（之十“超越正在超车的车”）"; 
acctexts[21]="未按规定让行的（之十一“会车时超车”）";
acctexts[22]="未按规定让行的（之十二“交叉路口超车”）";   
acctexts[23]="未按规定让行的（之十三“掉头时未让行”）";
acctexts[24]="未按规定让行的（之十四“进主路的车未让主路内的车”）";
acctexts[25]="未按规定让行的（之十五“辅路车未让出主路的车”）";
acctexts[26]="未按规定让行的（之十六“进出或穿越道路的车未让行”）";
acctexts[27]="未按规定让行的（之十七“进出停车泊位的车未让行”）";
acctexts[28]="未按规定让行的（之十八“进环路的车未让行”）";
acctexts[29]="未按规定让行的（之十九“有灯路口未让先被放行的车”）";

acctexts[30]="未按规定让行的（之二十“无灯路口未按标志标线让行”）";  
acctexts[31]="未按规定让行的（之二十一“无信号路口未让右侧道路来车先行”）";  
acctexts[32]="未按规定让行的（之 二十二“无信号路口转弯车未让直行车”）";
acctexts[33]="未按规定让行的（之二十三“无信号路口相对右转未让左转”）"; 
acctexts[34]="未按规定让行的（之二十四“未避让作业车辆”）";  
acctexts[35]="未按规定让行的（之二十五“有信号灯的路口右转未让被放行车辆”）  ";
acctexts[36]="未按规定让行的（之二十六“有信号灯的路口转弯未让直行的放行车辆”） "; 
acctexts[37]="依法应负全部责任的其他情形（之一“肇事逃逸”）";
acctexts[38]="依法应负全部责任的其他情形（之二“装载的飘散货物遗洒”）";
acctexts[39]="依法应负全部责任的其他情形（之三“未按规定装载”）";
acctexts[40]="依法应负全部责任的其他情形（之四“刮撞依法停放的车辆”）";
acctexts[41]="依法应负全部责任的其他情形（之五“右侧超车”）";
acctexts[42]="依法应负全部责任的其他情形（之六“单方事故”）";

var prior_click=function(event, ui) {
  		    
			if (i>21) 
			{
				i--;
				var s="img/05-"+i+".png";
				$("#Page_ChooseType_Msg").html("<BR/>"+(i-20)+"."+acctexts[i-21]+"<BR/>");
				$("#img").attr("src",s);
			}
  		};
		
var next_click=function(event, ui) {
  		    
			if (i<63)
			{
				i++;
				var s="img/05-"+i+".png";
				$("#Page_ChooseType_Msg").html("<BR/>"+(i-20)+"."+acctexts[i-21]+"<BR/>");
				$("#img").attr("src",s);
			}
};

var Page_ChooseType_DoRefresh=function(){
	$.getJSON(  "http://"+ServerIP+":"+ServerPort+"/case/photo",
			  {
				  phone:MyMobile,
				  action:"init"
				  
			  },
			  function(data,textStatus){
        		if (data.Suc){
					
					$("#Page_ChooseType_Content").show();
					$("#Page_ChooseType_Msg").html('<BR/>1.追尾<BR/>');
				}else
				{
					$( "#Page_ChooseType_Msg" ).text(data.Msg);
					$("#Page_ChooseType_Content").hide();
				}
			  }
			);
	
	
}
var Page_ChooseType_PageShow=function(event,ui){
	if (MyMobile==""){
		$("#Page_ChooseType_Content").hide();
		$("#Page_ChooseType_Msg").html('<BR/>您需要先设置您的手机号码，<BR/>点击右上角的设置按钮<BR/><BR/>');
	}else{
		Page_ChooseType_DoRefresh();
		
	
	}
}

function Page_ButtonSelectType_Click(){

}



var Page_ChooseType_Init=function()
{
	$( "#prior" ).bind( "click", prior_click);
	$( "#next" ).bind( "click", next_click);
	//$( "#Page_ButtonSelectType" ).bind( "click", Page_ButtonSelectType_Click);
	
	$( "#page_ChooseType" ).bind( "pageshow", Page_ChooseType_PageShow);
	
}