$(function(){
    $(".imgfullscreen").on("click", function(){
        var $this = $(this);
        if ($this.hasClass('clicked')){//dblclick for select
            $this.removeClass('clicked'); 
            //Double Click codes begin here
            $("#overlay").css("display", function(){
                return $("#overlay").css("display") == "none" ? "block":"none";
            });

        }else{ //single click for enlarge a single pic
            $this.addClass('clicked');
            setTimeout(function() { 
                if ($this.hasClass('clicked')){
                    $this.removeClass('clicked'); 
                    //Single Click codes begin here
                    parent.history.back();
                }
            }, 700);          
        }
    });

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
})
