$(function(){

    var list = ["/data/album/Jiangjunying_iphone6/100APPLE/IMG_0768.JPG",
                "/data/album/2008/2008-338.jpg",
                "/data/album/Jiangjunying_iphone6/102APPLE/IMG_2990.JPG",
                "/data/album/Jiangjunying_iphone6/102APPLE/IMG_2965.JPG",
                "/data/album/Jiangjunying_iphone6/102APPLE/IMG_2939.JPG"];

    var index = 0
    $("#slide").on("click", function(){
            getPhoto(list[index]);
    });

    function successCallback(result){
        index= index + 1;
        getPhoto(list[index]);
        $.when(result).done(function(result){
            console.log(result);
            $("#curr").attr("src", "data:image/jpeg;base64, " + result.message);
        });
    }
    
    function getPhoto(file){ // Async mode, success -> successCallback
        var data = {"path": file};
        $.ajax({ 
            type: "post", 
            url: "/slide", 
            contentType: 'application/json;charset=utf-8',
            dataType: 'json',
            data: JSON.stringify(data),
            async: true,
            success:function(result){successCallback(result)},
            error: function(xhr, status){ alert(xhr.status); }
        }); 
    }
});
