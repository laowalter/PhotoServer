$(function(){
    $('body').keyup(function(e){
       if(e.keyCode == 8){
           // user has pressed backspace
           alert("back key");
       }
       if(e.keyCode == 32){
         var index = parseInt($("#index").attr("src"),10)+1;
         window.open("single?&index="+index, "_self");
       }
    });

    // (function(){
    //     alert("hi");
    // })();


})

