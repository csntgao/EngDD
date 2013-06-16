<!DOCTYPE html> 



<html> 

<head> 
    <title>{{.Word.WordName}}八维联系法记单词</title> 
	
    <meta name="viewport" content="width=device-width, initial-scale=1"> 


	
	<link rel="stylesheet" href="/static/jqm/jquery.mobile.structure-1.3.1.min.css" />
    <link rel="stylesheet" href="/static/jqm/jquery.mobile.theme-1.3.1.min.css" />
    
	<script src="/static/jqm/jquery-1.7.1.min.js"></script>
    <script src="/static/jqm/jquery.mobile-1.3.1.min.js"></script>
   

<script>
function searchword(word)
{
     //window.navigate("/words/"+word);
	 window.location.href="/words/"+word;
};


$(document).ready(function(){
	    $( "#searchbutton" ).bind( "click", function(event, ui) {
			s="/words/"+$( "#input" ).val();
			//alert(s);
			window.location.href="/words/"+$( "#input" ).val();
		})
	})
		
		
  	

</script>



</head> 


<body> 

<div data-role="page" id="page1" data-position="fixed">

    <div data-role="header">
       <h1>八维联系法记单词</h1>
    </div>

    <div data-role="content">
	<form action="">
		<div data-role="fieldcontain">
            <div class="ui-grid-a">
				<div class="ui-block-a"><input name="" id="input" placeholder="" value="{{.Word.WordName}}" type=		                       "search"  data-mini="true"></div>
			    <div class="ui-block-b"><a data-role="button" id="searchbutton">查查看</a></div>
            </div>   
		</div>
    </form>
	<div data-role="collapsible-set">
            <div data-role="collapsible" data-collapsed="false">
                <h3>
                    基本：{{.Word.WordName}}
                </h3>
				<h5>
				 美：[{{.Word.APronounce}}] 英：[{{.Word.BPronounce}}]
				 </h5>
				
				
				{{with .Word.Meanings}}
                	{{range .}}
					</h5>
                   			{{.Property}}
							{{with .Meaning}}
                			{{range .}}
							{{.}};
							{{end}}
							{{end}}
					<BR/>		
					</h5>
                		
						
                	{{end}}
                	{{end}}
					
            </div>
            <div data-role="collapsible" data-collapsed="true">
                <h3>
                    拼写接近:<a href=/words/{{.Word.SpellSimilarWord}}>{{.Word.SpellSimilarWord}}</a>
                </h3>
				<ul data-role="listview" data-divider-theme="b" data-inset="true">
            		{{with .Word.SpellLikes}}
                	{{range .}}
                   		<li data-theme="c">
                		<a href="/words/{{.Item}}" data-transition="slide">
                    		{{.Item}}
                		</a>
            			</li>
						
                	{{end}}
                	{{end}}

					
            
        		</ul>
				
			</div>
			<div data-role="collapsible" data-collapsed="true">
                <h3>
                    读音接近:<a href=/words/{{.Word.PronounceSimilarWord}}>{{.Word.PronounceSimilarWord}} </a>
                </h3>
					<ul data-role="listview" data-divider-theme="b" data-inset="true">
            		{{with .Word.PronounceLikes}}
                	{{range .}}
                   		<li data-theme="c">
                		<a href="/words/{{.Item}}" data-transition="slide">
                    		{{.Item}}
                		</a>
            			</li>
						
                	{{end}}
                	{{end}}

					
            
        		</ul>
			
            </div>
			<div data-role="collapsible" data-collapsed="true">
                <h3>
				意思接近:<a href=/words/{{.Word.MeaningSimilarWord}}>{{.Word.MeaningSimilarWord}}</a>
                   
                </h3>
				<ul data-role="listview" data-divider-theme="b" data-inset="true">
            		{{with .Word.MeaningLikes}}
                	{{range .}}
                   		<li data-theme="c">
                		<a href="/words/{{.Item}}" data-transition="slide">
                    		{{.Item}}
                		</a>
            			</li>
						
                	{{end}}
                	{{end}}

					
            
        		</ul>
            </div>
			<div data-role="collapsible" data-collapsed="true">
                <h3>
				构词法：
                   
                </h3>
				前缀：<BR/>
				字根：<BR/>
				后缀：<BR/>
				派生词：<BR/>
            </div>
			<div data-role="collapsible" data-collapsed="true">
                <h3>
				分类：
                </h3>
			</div>
			<div data-role="collapsible" data-collapsed="true">
                <h3>
				短语:
                </h3>
				
				{{with .Word.Phrases}}
                {{range .}}
					    <h5>{{.Phrase}}<BR/></h5>
						{{with .Meanings}}
                		{{range .}}
					    <h5>{{.Meaning}}<BR/></h5>
						{{if `.ESample!=""`}}<h5>{{.ESample}}<BR/></h5>{{end}}
						{{if `.CSample==""`}}<h5>{{.CSample}}<BR/></h5>{{end}}
						{{end}}
                		{{end}}
			
				{{end}}
                {{end}}
			</div>
			<div data-role="collapsible" data-collapsed="true">
                <h3>
				例句：
				</h3>
				
					{{with .Word.Sentences}}
                	{{range .}}
					    <h5>{{.EnglishSentence}}<BR/>
                		{{.ChineseSentence}}  源自{{.Source}}</h5>
					{{end}}
                	{{end}}
				
			</div>
			
        </div>
      
       </p>

    </div>

    <div data-role="footer" data-position="fixed">
       <h1>FOOTER</h1>
    </div>
</div>



</body>
</html>

