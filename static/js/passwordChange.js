function changePassword() {
    if ($("#inputNewpassword").val() != $("#inputNewpasswordConfirmed").val()) {
        $('#alertBar').show();
        $("#alertMsg").html("password not the same");
        $("#alertBar").alert();
        setTimeout("$('#alertBar').hide();", 1500);
        return
    }
    gotoUrl = "/api/changePassword?czcookie=" + localStorage.czcookie;
    $.ajax({
        type: "PATCH",
        url: gotoUrl,
        data: { passwd: $('#inputNewpassword').val() },
        datatype: "text",
        async: false,
        success: function (data) {
            $('#alertBar').show();
            $("#alertMsg").html(data);
            $("#alertBar").alert();
            setTimeout("$('#alertBar').hide();", 1500);
        }
    });
}