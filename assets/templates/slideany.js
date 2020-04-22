$(function(){
    //解释下面的循环：
    //getPhotos首先调入list[0]，在getPhoto中ajax的成功之后，回调successCallback,
    //在回调函数中，把index+1 再次调用getPhoto[index+1]，重启继续上面的循环。

    var index =0
    getPhoto();
    function successCallback(result){
        $.when(result).done(function(result){
            if ($("#curr").length==0){
                $(".slide").append('<img id="curr" class="imgfullscreen" src="data:image/jpeg;base64,'+ result.message + '" alt="'+ index +'">');
            }else{
                $(".slide").append('<img id="next" class="imgfullscreen" src="data:image/jpeg;base64,'+ result.message + '" alt="'+ index +'">');
                $("#curr").attr("id", "prev");
                //$("#prev").slideDown("slow");
                $("#prev").remove();
                $("#next").attr("id", "curr");
            }
        });
        //sleep(20000)
        getPhoto();
    }
    
    function getPhoto(){ // Async mode, success -> successCallback
        var data = {"path": "any"};
        $.ajax({ 
            type: "post", 
            url: "/slideany", 
            contentType: 'application/json;charset=utf-8',
            dataType: 'json',
            data: JSON.stringify(data),
            async: true,
            success:function(result){
                //Dont know how to set delay.
                setTimeout(function(){
                    //alert("Boom!");
                }, 32000);
                successCallback(result)
            },
            error: function(xhr, status){alert("file not found");}
        }); 
    }

    function sleep(miliseconds) {
        var currentTime = new Date().getTime();
            while (currentTime + miliseconds >= new Date().getTime()) {
        }
    }
});
