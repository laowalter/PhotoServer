$(function(){

    var selectList = [];
    var height = $("#photolist").outerHeight()/2;
    $(window).on("scroll resize", function(){
        var top = $(window).scrollTop()+ height;
        $("#photolist").offset({top:top});
    })

    $(".thumb").on("click", function(){
        var $this = $(this);
        if ($this.hasClass('clicked')){//dblclick for select
            $this.removeClass('clicked'); 
            $this.toggleClass('selected');
            var index = selectList.indexOf($this.attr("path"));
            if (index>-1){
                selectList.splice(index, 1);
            } else{
                selectList.push($this.attr("path"));
            }
        }else{ //single click for enlarge a single pic
            $this.addClass('clicked');
            setTimeout(function() { 
                if ($this.hasClass('clicked')){
                    $this.removeClass('clicked'); 
                    window.open($this.attr("href"), "_self");
                }
            }, 700);          
        }

        $(function(){
            if (selectList.length>0){
                $("#selectbtn").css("display", "block");
            } else {
                $("#selectbtn").css("display", "none");                   
            }
        });
    });

    $("#delete").on("click", function(){
        if(confirm("Sure to delete these " + selectList.length + " photos?")){
            $.ajax({ 
                type: "post", 
                url: "/delete", 
                contentType: 'application/json',
                dataType: 'text',
                data: selectList.join(','),
                success:function(result){ 
                    location.reload();
                },
                error: function(xhr, status){ alert(xhr.status); }
            }); 
        }else{
            return false;
        }
    });

    $("#addtags").on("click", function(){
        var tag = prompt("Enter your tags, add <space> to split:", "");
        if (tag == null ||tag == "") {
            return;
        } else {
            if(confirm("Sure to Add Tag to these " + selectList.length + " photos?")){
                $.ajax({ 
                    type: "post", 
                    url: "/addtags", 
                    contentType: 'application/json',
                    dataType: 'text',
                    data: tag + "|" + selectList.join(','),
                    success:function(result){ 
                        location.reload();
                    },
                    error: function(xhr, status){ alert(xhr.status); }
                }); 
            }else{
                return false;
            }
        }
    });
})
