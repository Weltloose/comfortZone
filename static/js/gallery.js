$(document).ready(function () {
    gotoUrl = "/api/getPublicPhotoesInfo?czcookie=" + localStorage.czcookie
    $.ajax({
        type: "GET",
        url: gotoUrl,
        data: "",
        dataType: "JSON",
        async: false,
        success: function (result) {
            photoesNum = result.data.length
            for (var i = 0; i < photoesNum; i++) {
                console.log(result.data[i].idName)
                $("#galleryContent").append(' <div class=col-lg-3 col-md-4 col-6"> \
                <div href="#lb-gallery3-t" data-slide-to="' + i + '" data-toggle="modal" class= "d-block mb-4 h-100"> \
                <img class="img-fluid img-thumbnail lazyload" data-original="' + result.data[i].idName + '" alt="" > \
                </div> \
            </div > ');
                if (i == 0) {
                    $("#galleryModalContent").append('<div class="carousel-item active text-center "><img  src="' + result.data[i].idName + '" alt="" class="carouselImg"></div>')
                } else {
                    $("#galleryModalContent").append('<div class="carousel-item text-center "><img  src="' + result.data[i].idName + '" alt=""  class="carouselImg"></div>')
                }
            }
        }
    })
})

$(function () {
    $("img").lazyload({ effect: "fadeIn" });
});

$(document).ready(function () {
    height = screen.height * 0.7
    console.log(screen.height)
    console.log(height)
    $(".carouselImg").css("height", height)
    $(".carouselImg").css("width", "auto")
}); 