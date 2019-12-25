function toogleBoard() {
    var savedComments = []
    var len = 0
    if (localStorage.getItem("savedComments") != null) {
        savedComments = JSON.parse(localStorage.savedComments)
        len = savedComments.length
    }
    gotoUrl = "/api/getComment?czcookie=" + localStorage.czcookie + "&saved=" + len;
    $.ajax({
        type: "GET",
        url: gotoUrl,
        data: "",
        dataType: "JSON",
        async: false,
        success: function (result) {
            var recievedComments = result["comments"];
            if (recievedComments != null) {
                savedComments = savedComments.concat(recievedComments);
                localStorage.savedLen = savedComments.length;
                localStorage.savedComments = JSON.stringify(savedComments);
            }
            // console.log(JSON.stringify(savedComments));
            $("#boardModal").text("");
            for (j = savedComments.length - 1; j >= 0; j--) {
                $("#boardModal").append(formatComment(savedComments[j]));
            }
            $("#exampleModalLong").modal();
        }
    })
}
function sendComment() {
    gotoUrl = "/api/comment?czcookie=" + localStorage.czcookie;
    $.ajax({
        type: "POST",
        url: gotoUrl,
        data: { textareaContents: $('#textareaContents').val() },//第一个textareaContents值文本框的Id，用jquery方式的AJAX提交，这个Id必须要
        dataType: "text",
        success: function (data) {
            $("#textareaContents").val("");
            toogleBoard();
        }
    });
}
function formatComment(obj) {
    return obj.createTime + "&nbsp;&nbsp;&nbsp;&nbsp" + obj.username + "</br>&nbsp;&nbsp;&nbsp;&nbsp" + obj.content + "</br>";
}

function uploadPublicPhotoes() {
    $("#addPublicPhotoesFileInput").click();
}

function uploadPrivatePhotoes() {
    $("#addPrivatePhotoesFileInput").click();
}

function uploadPublicVoice() {
    $("#addPublicVoiceFileInput").click();
}

function uploadPublicPhotosToBack(e) {
    var formData = new FormData();
    for (var i = 0; i < e.files.length; i++) {
        formData.append(i, e.files[i]);
    }
    gotoUrl = "/api/uploadPublicPhotoes?czcookie=" + localStorage.czcookie;
    $.ajax({
        type: "post",
        url: gotoUrl,
        processData: false,
        contentType: false,
        data: formData,
        dataType: "json",
        success: function (data) {
            if (data.resp == "success") {
                alert("all public images upload success")
            }
            else {
                alert("public images upload failured")
            }
        }
    });
}
function uploadPublicVoiceToBack(e) {
    var formData = new FormData();
    for (var i = 0; i < e.files.length; i++) {
        formData.append(i, e.files[i]);
    }
    gotoUrl = "/api/uploadPublicVoice?czcookie=" + localStorage.czcookie;
    $.ajax({
        type: "post",
        url: gotoUrl,
        processData: false,
        contentType: false,
        data: formData,
        dataType: "json",
        success: function (data) {
            if (data.resp == "success") {
                alert("all public voice upload success")
            }
            else {
                alert("public voice upload failured")
            }
        }
    });
}
function gallery() {
    gotoUrl = "/gallery?czcookie=" + localStorage.czcookie;
    window.location.href = gotoUrl;
}

function index() {
    gotoUrl = "/index?czcookie=" + localStorage.czcookie;
    window.location.href = gotoUrl;
}

function passwordp() {
    gotoUrl = "/password?czcookie=" + localStorage.czcookie;
    window.location.href = gotoUrl;
}