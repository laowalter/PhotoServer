$(function(){

    var list = ["/data/album/Jiangjunying_iphone6/100APPLE/IMG_0768.JPG",
                "/data/album/2008/2008-338.jpg",
                "/data/album/Jiangjunying_iphone6/102APPLE/IMG_2990.JPG",
                "/data/album/Jiangjunying_iphone6/102APPLE/IMG_2965.JPG",
                "/data/album/Jiangjunying_iphone6/102APPLE/IMG_2939.JPG"];

    var index = 0;


    //解释下面的循环：
    //getPhotos首先调入list[0]，在getPhoto中ajax的成功之后，回调successCallback,
    //在回调函数中，把index+1 再次调用getPhoto[index+1]，重启继续上面的循环。

    getPhoto(list[index]);

    function successCallback(result){
            $.when(result).done(function(result){
                if ($("#curr").length==0){
                    $(".slide").append('<img height="300px" id="curr" src="data:image/jpeg;base64,'+ result.message + '" alt="'+list[index] +'">');
                }else{
                    $(".slide").append('<img height="300px" id="next" src="data:image/jpeg;base64,'+ result.message + '" alt="'+list[index] +'">');
                    $("#curr").attr("id", "prev");
                    $("#prev").slideDown("slow");
                    $("#prev").remove();
                    $("#next").attr("id", "curr");
                }
            });
            sleep(5000)
            index= index + 1;
            if (index<list.length){
                getPhoto(list[index]);
            }else{
                return
            }
    }
    
    function getPhoto(file){ // Async mode, success -> successCallback
        var data = {"path": file};
        console.log(file);
        $.ajax({ 
            type: "post", 
            url: "/slide", 
            contentType: 'application/json;charset=utf-8',
            dataType: 'json',
            data: JSON.stringify(data),
            async: true,
            success:function(result){successCallback(result)},
            error: function(xhr, status){alert("file not found");}
        }); 
    }

    function sleep(miliseconds) {
        var currentTime = new Date().getTime();
            while (currentTime + miliseconds >= new Date().getTime()) {
        }
    }
});
