	var captureSuccess=function(mediaFiles) {  
        var i, len;  
        for (i = 0, len = mediaFiles.length; i < len; i += 1) {  
           //业务逻辑  
            navigator.notification.alert(mediaFiles[i].fullPath + " " +mediaFiles[i].name);  
        }         
    }  
  
    //captureAudio方法执行失败后回调函数  
    var  captureError=function (error) {  
        var msg = 'capture 发生错误: ' + error.code;  
        navigator.notification.alert(msg, null, 'Uh oh!');  
    }  
  
    var  captureImage=function() {  
        // limit 拍照的张数  
        navigator.device.capture.captureImage(captureSuccess, captureError, {limit: 1});  
    }  
	var  Page_TakePhotos_ButtonTakePhoto_Click=function(event,ui)
	{
		captureImage();
	}
	var  Page_TakePhotos_ButtonIgnore_Click=function(event,ui){
		if (j<pictexts.length-1)
		{
			j++;
			$("#lbdesp").text(pictexts[j]);
				//$("#img").attr("src",s);
		}else{
				window.location.href="#page_CommitState"
		}
	
	}



var j=0
	


	var pictexts=new Array()
	pictexts[0]="请按照图示进行拍照，[车头45°]";
	pictexts[1]="请按照图示进行拍照，[车尾45°]";
	pictexts[2]="请按照图示进行拍照，[全瞰图]";
	pictexts[3]="请按照图示进行拍照，[局部图]";
	pictexts[4]="自己认为需要补充的图";
var Page_TakePhotos_PageShow=function(){
	j=0;
	$("#lbdesp").text(pictexts[j]);
}
	
 
var Page_TakePhotos_Init=function(){
		$( "#Page_TakePhotos_ButtonIgnore" ).bind( "click", Page_TakePhotos_ButtonIgnore_Click);
		$( "#Page_TakePhotos_ButtonTakePhoto" ).bind( "click", Page_TakePhotos_ButtonTakePhoto_Click);
		$( "#page_TakePhotos" ).bind( "pageshow", Page_TakePhotos_PageShow);
		
		$( "#resultrefresh" ).bind( "click", function(event, ui) {
  		    
			i++;
			if (i % 2 ==0)
			{
				
				window.location="#page09";
				
			}
			else
			{
				window.location="#page10";
				
			}	
	
  		});
		
		$( "#resultconfirm" ).bind( "click", function(event, ui) {
  		    
			window.location="#page11";
				
	
  		});

}
	
		
		