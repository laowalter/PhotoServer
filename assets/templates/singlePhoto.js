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
})
